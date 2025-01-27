// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: secret/registration/v1beta1/types.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_scrtlabs_SecretNetwork_x_registration_remote_attestation "github.com/scrtlabs/SecretNetwork/x/registration/remote_attestation"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type SeedConfig struct {
	MasterCert   string `protobuf:"bytes,1,opt,name=master_cert,json=masterCert,proto3" json:"pk"`
	EncryptedKey string `protobuf:"bytes,2,opt,name=encrypted_key,json=encryptedKey,proto3" json:"encKey"`
}

func (m *SeedConfig) Reset()         { *m = SeedConfig{} }
func (m *SeedConfig) String() string { return proto.CompactTextString(m) }
func (*SeedConfig) ProtoMessage()    {}
func (*SeedConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3db05f1d182f4de, []int{0}
}
func (m *SeedConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SeedConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SeedConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SeedConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SeedConfig.Merge(m, src)
}
func (m *SeedConfig) XXX_Size() int {
	return m.Size()
}
func (m *SeedConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SeedConfig.DiscardUnknown(m)
}

var xxx_messageInfo_SeedConfig proto.InternalMessageInfo

type RegistrationNodeInfo struct {
	Certificate   github_com_scrtlabs_SecretNetwork_x_registration_remote_attestation.Certificate `protobuf:"bytes,1,opt,name=certificate,proto3,casttype=github.com/scrtlabs/SecretNetwork/x/registration/remote_attestation.Certificate" json:"certificate,omitempty"`
	EncryptedSeed []byte                                                                          `protobuf:"bytes,2,opt,name=encrypted_seed,json=encryptedSeed,proto3" json:"encrypted_seed,omitempty"`
}

func (m *RegistrationNodeInfo) Reset()         { *m = RegistrationNodeInfo{} }
func (m *RegistrationNodeInfo) String() string { return proto.CompactTextString(m) }
func (*RegistrationNodeInfo) ProtoMessage()    {}
func (*RegistrationNodeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3db05f1d182f4de, []int{1}
}
func (m *RegistrationNodeInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegistrationNodeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegistrationNodeInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegistrationNodeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistrationNodeInfo.Merge(m, src)
}
func (m *RegistrationNodeInfo) XXX_Size() int {
	return m.Size()
}
func (m *RegistrationNodeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistrationNodeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RegistrationNodeInfo proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SeedConfig)(nil), "secret.registration.v1beta1.SeedConfig")
	proto.RegisterType((*RegistrationNodeInfo)(nil), "secret.registration.v1beta1.RegistrationNodeInfo")
}

func init() {
	proto.RegisterFile("secret/registration/v1beta1/types.proto", fileDescriptor_f3db05f1d182f4de)
}

var fileDescriptor_f3db05f1d182f4de = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x41, 0x4e, 0x32, 0x31,
	0x14, 0xc7, 0xa7, 0x2c, 0x48, 0xbe, 0xc2, 0xe7, 0x62, 0xc2, 0x82, 0x68, 0x52, 0x0c, 0x89, 0xc1,
	0xd5, 0x34, 0xc4, 0x03, 0x98, 0xc0, 0xca, 0x90, 0x60, 0x32, 0xec, 0xdc, 0x90, 0x4e, 0xe7, 0xcd,
	0xd8, 0x00, 0xed, 0xa4, 0x7d, 0xa8, 0x73, 0x0b, 0x8f, 0xe1, 0x01, 0x3c, 0x04, 0x4b, 0x96, 0xae,
	0x88, 0x0e, 0x3b, 0x8e, 0xe0, 0xca, 0x30, 0x63, 0x04, 0x97, 0xee, 0x9a, 0xf4, 0xd7, 0xfe, 0xdf,
	0xef, 0xfd, 0x69, 0xcf, 0x81, 0xb4, 0x80, 0xdc, 0x42, 0xaa, 0x1c, 0x5a, 0x81, 0xca, 0x68, 0xfe,
	0xd0, 0x8f, 0x00, 0x45, 0x9f, 0x63, 0x9e, 0x81, 0x0b, 0x32, 0x6b, 0xd0, 0xf8, 0x67, 0x15, 0x18,
	0x1c, 0x83, 0xc1, 0x37, 0x78, 0xda, 0x4a, 0x4d, 0x6a, 0x4a, 0x8e, 0xef, 0x4f, 0xd5, 0x93, 0x6e,
	0x42, 0xe9, 0x04, 0x20, 0x1e, 0x1a, 0x9d, 0xa8, 0xd4, 0xef, 0xd1, 0xc6, 0x42, 0x38, 0x04, 0x3b,
	0x95, 0x60, 0xb1, 0x4d, 0xce, 0xc9, 0xe5, 0xbf, 0x41, 0x7d, 0xb7, 0xe9, 0xd4, 0xb2, 0x59, 0x48,
	0xab, 0xab, 0x21, 0x58, 0xf4, 0x39, 0xfd, 0x0f, 0x5a, 0xda, 0x3c, 0x43, 0x88, 0xa7, 0x33, 0xc8,
	0xdb, 0xb5, 0x12, 0xa5, 0xbb, 0x4d, 0xa7, 0x0e, 0x5a, 0x8e, 0x20, 0x0f, 0x9b, 0x3f, 0xc0, 0x08,
	0xf2, 0xee, 0x2b, 0xa1, 0xad, 0xf0, 0x68, 0xac, 0xb1, 0x89, 0xe1, 0x46, 0x27, 0xc6, 0x5f, 0xd2,
	0xc6, 0x3e, 0x4b, 0x25, 0x4a, 0x0a, 0x84, 0x32, 0xb2, 0x39, 0x98, 0x7c, 0x6e, 0x3a, 0xb7, 0xa9,
	0xc2, 0xfb, 0x65, 0x14, 0x48, 0xb3, 0xe0, 0x4e, 0x5a, 0x9c, 0x8b, 0xc8, 0xf1, 0x49, 0x29, 0x38,
	0x06, 0x7c, 0x34, 0x76, 0xc6, 0x9f, 0x7e, 0xaf, 0xc4, 0xc2, 0xc2, 0x20, 0x4c, 0x05, 0x22, 0x38,
	0xac, 0xe4, 0x87, 0x87, 0xaf, 0xc3, 0xe3, 0x1c, 0xff, 0x82, 0x9e, 0x1c, 0x04, 0x1c, 0x40, 0x5c,
	0x1a, 0x34, 0xc3, 0x83, 0xd6, 0x7e, 0x2d, 0x03, 0xb1, 0xfa, 0x60, 0xde, 0x4b, 0xc1, 0xc8, 0xaa,
	0x60, 0x64, 0x5d, 0x30, 0xf2, 0x5e, 0x30, 0xf2, 0xbc, 0x65, 0xde, 0x7a, 0xcb, 0xbc, 0xb7, 0x2d,
	0xf3, 0xee, 0xae, 0xff, 0x3c, 0xa6, 0xd2, 0x08, 0x56, 0x8b, 0x79, 0x55, 0x5d, 0x54, 0x2f, 0x8b,
	0xb8, 0xfa, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x8b, 0xe7, 0x2b, 0x46, 0xe6, 0x01, 0x00, 0x00,
}

func (this *SeedConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SeedConfig)
	if !ok {
		that2, ok := that.(SeedConfig)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.MasterCert != that1.MasterCert {
		return false
	}
	if this.EncryptedKey != that1.EncryptedKey {
		return false
	}
	return true
}
func (this *RegistrationNodeInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RegistrationNodeInfo)
	if !ok {
		that2, ok := that.(RegistrationNodeInfo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Certificate, that1.Certificate) {
		return false
	}
	if !bytes.Equal(this.EncryptedSeed, that1.EncryptedSeed) {
		return false
	}
	return true
}
func (m *SeedConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SeedConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SeedConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EncryptedKey) > 0 {
		i -= len(m.EncryptedKey)
		copy(dAtA[i:], m.EncryptedKey)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.EncryptedKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.MasterCert) > 0 {
		i -= len(m.MasterCert)
		copy(dAtA[i:], m.MasterCert)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.MasterCert)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RegistrationNodeInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegistrationNodeInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegistrationNodeInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EncryptedSeed) > 0 {
		i -= len(m.EncryptedSeed)
		copy(dAtA[i:], m.EncryptedSeed)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.EncryptedSeed)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Certificate) > 0 {
		i -= len(m.Certificate)
		copy(dAtA[i:], m.Certificate)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Certificate)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SeedConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MasterCert)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.EncryptedKey)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *RegistrationNodeInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Certificate)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.EncryptedSeed)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SeedConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SeedConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SeedConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MasterCert", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MasterCert = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EncryptedKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EncryptedKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RegistrationNodeInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RegistrationNodeInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegistrationNodeInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Certificate", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Certificate = append(m.Certificate[:0], dAtA[iNdEx:postIndex]...)
			if m.Certificate == nil {
				m.Certificate = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EncryptedSeed", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EncryptedSeed = append(m.EncryptedSeed[:0], dAtA[iNdEx:postIndex]...)
			if m.EncryptedSeed == nil {
				m.EncryptedSeed = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
