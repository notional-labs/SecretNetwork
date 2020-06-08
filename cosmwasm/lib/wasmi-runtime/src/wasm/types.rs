// use crate::cosmwasm::types::CosmosMsg;

use crate::crypto::{AESKey, PublicKey, SIVEncryptable};
use crate::wasm::io::calc_encryption_key;
use enclave_ffi_types::EnclaveError;
use log::*;
use serde::{Deserialize, Serialize};

pub type IoNonce = [u8; 32];

#[derive(Serialize, Deserialize)]
pub struct SecretMessage {
    pub nonce: IoNonce,
    pub user_public_key: PublicKey,
    pub msg: Vec<u8>,
}

impl SecretMessage {
    pub fn decrypt(&self) -> Result<Vec<u8>, EnclaveError> {
        let key = self.encryption_key();

        // pass
        let msg = key
            .decrypt_siv(self.msg.as_slice(), &vec![&[]])
            .map_err(|err| {
                error!(
                    "handle() got an error while trying to decrypt the msg: {}",
                    err
                );
                EnclaveError::DecryptionError
            })?;

        Ok(msg)
    }

    pub fn encryption_key(&self) -> AESKey {
        calc_encryption_key(&self.nonce, &self.user_public_key)
    }

    pub fn from_slice(msg: &[u8]) -> Result<Self, EnclaveError> {
        // 32 bytes of AD
        // 33 bytes of secp256k1 compressed public key
        // 16+ bytes of encrypted data

        if msg.len() < 82 {
            error!(
                "Encrypted message length {:?} is too short. Cannot parse",
                msg.len()
            );
            return Err(EnclaveError::DecryptionError);
        };

        let mut nonce = [0u8; 32];
        nonce.copy_from_slice(&msg[0..32]);

        let mut user_pubkey = [0u8; 32];
        user_pubkey.copy_from_slice(&msg[32..64]);

        Ok(SecretMessage {
            nonce,
            user_public_key: user_pubkey,
            msg: msg[64..].to_vec(),
        })
    }

    pub fn to_slice(&self) -> Vec<u8> {
        let mut packed_msg: Vec<u8> = self.nonce.to_vec();
        packed_msg.extend_from_slice(&self.user_public_key);
        packed_msg.extend_from_slice(self.msg.as_slice());
        packed_msg
    }
}

#[cfg(feature = "test")]
pub mod tests {

    use super::*;
    use crate::crypto::{key_manager, AESKey, SIVEncryptable, Seed};

    // todo: fix test vectors to actually work
    fn test_new_from_slice() {
        let nonce = [0u8; 32];
        let user_public_key = [0u8; 32];
        let msg = "{\"ok\": \"{\"balance\": \"108\"}\"}";

        let mut slice = nonce.to_vec();
        slice.extend_from_slice(&user_public_key);
        slice.extend_from_slice(msg.as_bytes());

        let secret_msg = SecretMessage {
            nonce,
            user_public_key,
            msg: msg.as_bytes().to_vec(),
        };

        let msg_from_slice = SecretMessage::from_slice(&slice);

        assert_eq!(secret_msg, msg_from_slice);
    }

    // todo: fix test vectors to actually work
    fn test_msg_decrypt() {
        let seed = Seed::new()?;

        if let Err(e) = key_manager.set_consensus_seed(seed) {
            fail!("Failed to set seed")
        }

        let nonce = [0u8; 32];
        let user_public_key = [0u8; 32];

        let msg = "{\"ok\": \"{\"balance\": \"108\"}\"}";
        let ket = calc_encryption_key(&nonce, &user_public_key);

        let encrypted_msg = key.encrypt_siv(msg.as_bytes(), &[&[]]);

        let secret_msg = SecretMessage {
            nonce,
            user_public_key,
            msg: encrypted_msg,
        };

        let decrypted_msg = secret_msg.decrypt()?;

        assert_eq!(decrypted_msg, msg)
    }
}
