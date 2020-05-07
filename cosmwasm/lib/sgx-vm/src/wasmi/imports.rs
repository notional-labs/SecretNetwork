//! This file should be autogenerated based on the headers created from the .edl file.

use enclave_ffi_types::{Ctx, EnclaveBuffer, HandleResult, InitResult, KeyGenResult, QueryResult};
use sgx_types::*;

extern "C" {
    /// Copy a buffer into the enclave memory space, and receive an opaque pointer to it.
    pub fn ecall_allocate(
        eid: sgx_enclave_id_t,
        retval: *mut EnclaveBuffer,
        buffer: *const u8,
        length: usize,
    ) -> sgx_status_t;

    pub fn ecall_init_seed(
        eid: sgx_enclave_id_t,
        retval: *mut sgx_status_t,
        public_key: *const u8,
        public_key_len: u32,
        encrypted_seed: *const u8,
        encrypted_seed_len: u32
    ) -> sgx_status_t;

    /// Trigger the init method in a wasm contract
    pub fn ecall_init(
        eid: sgx_enclave_id_t,
        retval: *mut InitResult,
        context: Ctx,
        gas_limit: u64,
        contract: *const u8,
        contract_len: usize,
        env: *const u8,
        env_len: usize,
        msg: *const u8,
        msg_len: usize,
    ) -> sgx_status_t;

    /// Trigger a handle method in a wasm contract
    pub fn ecall_handle(
        eid: sgx_enclave_id_t,
        retval: *mut HandleResult,
        context: Ctx,
        gas_limit: u64,
        contract: *const u8,
        contract_len: usize,
        env: *const u8,
        env_len: usize,
        msg: *const u8,
        msg_len: usize,
    ) -> sgx_status_t;

    /// Trigger a query method in a wasm contract
    pub fn ecall_query(
        eid: sgx_enclave_id_t,
        retval: *mut QueryResult,
        context: Ctx,
        gas_limit: u64,
        contract: *const u8,
        contract_len: usize,
        msg: *const u8,
        msg_len: usize,
    ) -> sgx_status_t;

    /// Trigger a key generation method in a wasm contract
    pub fn ecall_key_gen(eid: sgx_enclave_id_t, retval: *mut KeyGenResult) -> sgx_status_t;

    pub fn ecall_get_registration_quote(eid: sgx_enclave_id_t, retval: *mut sgx_status_t,
                                        target_info: *const sgx_target_info_t, report: *mut sgx_report_t) -> sgx_status_t;

    pub fn ecall_get_attestation_report(eid: sgx_enclave_id_t, retval: *mut sgx_status_t);

    // pub fn ocall_sgx_init_quote ( ret_val : *mut sgx_status_t,
    //                               ret_ti  : *mut sgx_target_info_t,
    //                               ret_gid : *mut sgx_epid_group_id_t) -> sgx_status_t;
    // pub fn ocall_get_ias_socket ( ret_val : *mut sgx_status_t,
    //                               ret_fd  : *mut i32) -> sgx_status_t;
    //
    // pub fn ocall_get_quote (ret_val            : *mut sgx_status_t,
    //                         p_sigrl            : *const u8,
    //                         sigrl_len          : u32,
    //                         p_report           : *const sgx_report_t,
    //                         quote_type         : sgx_quote_sign_type_t,
    //                         p_spid             : *const sgx_spid_t,
    //                         p_nonce            : *const sgx_quote_nonce_t,
    //                         p_qe_report        : *mut sgx_report_t,
    //                         p_quote            : *mut u8,
    //                         maxlen             : u32,
    //                         p_quote_len        : *mut u32) -> sgx_status_t;
    // pub fn sgx_init_quote(p_target_info: *mut sgx_target_info_t, p_gid: *mut sgx_epid_group_id_t) -> sgx_status_t;

    // pub fn sgx_calc_quote_size(p_sig_rl: *const uint8_t, sig_rl_size: uint32_t, p_quote_size: *mut uint32_t) -> sgx_status_t;

    // pub fn sgx_get_quote(p_report: *const sgx_report_t, quote_type: sgx_quote_sign_type_t,
    //                      p_spid: *const sgx_spid_t, p_nonce: *const sgx_quote_nonce_t, p_sig_rl: *const uint8_t,
    //                      sig_rl_size: uint32_t, p_qe_report: *mut sgx_report_t, p_quote: *mut sgx_quote_t,
    //                      quote_size: uint32_t) -> sgx_status_t;
}
