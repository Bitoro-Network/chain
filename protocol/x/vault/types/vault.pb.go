// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bitoroprotocol/vault/vault.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

// VaultType represents different types of vaults.
type VaultType int32

const (
	// Default value, invalid and unused.
	VaultType_VAULT_TYPE_UNSPECIFIED VaultType = 0
	// Vault is associated with a CLOB pair.
	VaultType_VAULT_TYPE_CLOB VaultType = 1
)

var VaultType_name = map[int32]string{
	0: "VAULT_TYPE_UNSPECIFIED",
	1: "VAULT_TYPE_CLOB",
}

var VaultType_value = map[string]int32{
	"VAULT_TYPE_UNSPECIFIED": 0,
	"VAULT_TYPE_CLOB":        1,
}

func (x VaultType) String() string {
	return proto.EnumName(VaultType_name, int32(x))
}

func (VaultType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9eca6a79511a8e91, []int{0}
}

// VaultStatus represents the status of a vault.
type VaultStatus int32

const (
	// Default value, invalid and unused.
	VaultStatus_VAULT_STATUS_UNSPECIFIED VaultStatus = 0
	// Don’t place orders. Does not count toward global vault balances.
	// A vault can only be set to this status if its equity is non-positive.
	VaultStatus_VAULT_STATUS_DEACTIVATED VaultStatus = 1
	// Don’t place orders. Does count towards global vault balances.
	VaultStatus_VAULT_STATUS_STAND_BY VaultStatus = 2
	// Places orders on both sides of the book.
	VaultStatus_VAULT_STATUS_QUOTING VaultStatus = 3
	// Only place orders that close the position.
	VaultStatus_VAULT_STATUS_CLOSE_ONLY VaultStatus = 4
)

var VaultStatus_name = map[int32]string{
	0: "VAULT_STATUS_UNSPECIFIED",
	1: "VAULT_STATUS_DEACTIVATED",
	2: "VAULT_STATUS_STAND_BY",
	3: "VAULT_STATUS_QUOTING",
	4: "VAULT_STATUS_CLOSE_ONLY",
}

var VaultStatus_value = map[string]int32{
	"VAULT_STATUS_UNSPECIFIED": 0,
	"VAULT_STATUS_DEACTIVATED": 1,
	"VAULT_STATUS_STAND_BY":    2,
	"VAULT_STATUS_QUOTING":     3,
	"VAULT_STATUS_CLOSE_ONLY":  4,
}

func (x VaultStatus) String() string {
	return proto.EnumName(VaultStatus_name, int32(x))
}

func (VaultStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9eca6a79511a8e91, []int{1}
}

// VaultId uniquely identifies a vault by its type and number.
type VaultId struct {
	// Type of the vault.
	Type VaultType `protobuf:"varint,1,opt,name=type,proto3,enum=bitoroprotocol.vault.VaultType" json:"type,omitempty"`
	// Unique ID of the vault within above type.
	Number uint32 `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
}

func (m *VaultId) Reset()         { *m = VaultId{} }
func (m *VaultId) String() string { return proto.CompactTextString(m) }
func (*VaultId) ProtoMessage()    {}
func (*VaultId) Descriptor() ([]byte, []int) {
	return fileDescriptor_9eca6a79511a8e91, []int{0}
}
func (m *VaultId) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VaultId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VaultId.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VaultId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VaultId.Merge(m, src)
}
func (m *VaultId) XXX_Size() int {
	return m.Size()
}
func (m *VaultId) XXX_DiscardUnknown() {
	xxx_messageInfo_VaultId.DiscardUnknown(m)
}

var xxx_messageInfo_VaultId proto.InternalMessageInfo

func (m *VaultId) GetType() VaultType {
	if m != nil {
		return m.Type
	}
	return VaultType_VAULT_TYPE_UNSPECIFIED
}

func (m *VaultId) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func init() {
	proto.RegisterEnum("bitoroprotocol.vault.VaultType", VaultType_name, VaultType_value)
	proto.RegisterEnum("bitoroprotocol.vault.VaultStatus", VaultStatus_name, VaultStatus_value)
	proto.RegisterType((*VaultId)(nil), "bitoroprotocol.vault.VaultId")
}

func init() { proto.RegisterFile("bitoroprotocol/vault/vault.proto", fileDescriptor_9eca6a79511a8e91) }

var fileDescriptor_9eca6a79511a8e91 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x48, 0xca, 0x2c, 0xc9,
	0x2f, 0xca, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x2f, 0x4b, 0x2c, 0xcd, 0x29,
	0x81, 0x90, 0x7a, 0x60, 0x41, 0x21, 0x11, 0x54, 0x15, 0x7a, 0x60, 0x39, 0xa5, 0x30, 0x2e, 0xf6,
	0x30, 0x10, 0xc3, 0x33, 0x45, 0xc8, 0x98, 0x8b, 0xa5, 0xa4, 0xb2, 0x20, 0x55, 0x82, 0x51, 0x81,
	0x51, 0x83, 0xcf, 0x48, 0x5e, 0x0f, 0x9b, 0x7a, 0x3d, 0xb0, 0xe2, 0x90, 0xca, 0x82, 0xd4, 0x20,
	0xb0, 0x62, 0x21, 0x31, 0x2e, 0xb6, 0xbc, 0xd2, 0xdc, 0xa4, 0xd4, 0x22, 0x09, 0x26, 0x05, 0x46,
	0x0d, 0xde, 0x20, 0x28, 0x4f, 0xcb, 0x86, 0x8b, 0x13, 0xae, 0x54, 0x48, 0x8a, 0x4b, 0x2c, 0xcc,
	0x31, 0xd4, 0x27, 0x24, 0x3e, 0x24, 0x32, 0xc0, 0x35, 0x3e, 0xd4, 0x2f, 0x38, 0xc0, 0xd5, 0xd9,
	0xd3, 0xcd, 0xd3, 0xd5, 0x45, 0x80, 0x41, 0x48, 0x98, 0x8b, 0x1f, 0x49, 0xce, 0xd9, 0xc7, 0xdf,
	0x49, 0x80, 0x51, 0x6b, 0x36, 0x23, 0x17, 0x37, 0x58, 0x7b, 0x70, 0x49, 0x62, 0x49, 0x69, 0xb1,
	0x90, 0x0c, 0x97, 0x04, 0x44, 0x51, 0x70, 0x88, 0x63, 0x48, 0x68, 0x30, 0x9a, 0x11, 0xe8, 0xb2,
	0x2e, 0xae, 0x8e, 0xce, 0x21, 0x9e, 0x61, 0x8e, 0x21, 0xae, 0x2e, 0x02, 0x8c, 0x42, 0x92, 0x5c,
	0xa2, 0x28, 0xb2, 0xc1, 0x21, 0x8e, 0x7e, 0x2e, 0xf1, 0x4e, 0x91, 0x02, 0x4c, 0x42, 0x12, 0x5c,
	0x22, 0x28, 0x52, 0x81, 0xa1, 0xfe, 0x21, 0x9e, 0x7e, 0xee, 0x02, 0xcc, 0x42, 0xd2, 0x5c, 0xe2,
	0x28, 0x32, 0xce, 0x3e, 0xfe, 0xc1, 0xae, 0xf1, 0xfe, 0x7e, 0x3e, 0x91, 0x02, 0x2c, 0x4e, 0x01,
	0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72,
	0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x65, 0x96, 0x9e, 0x59, 0x92, 0x51,
	0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x0f, 0x09, 0x3e, 0xdd, 0xbc, 0xd4, 0x92, 0xf2, 0xfc, 0xa2,
	0x6c, 0xfd, 0xe4, 0x8c, 0xc4, 0xcc, 0x3c, 0x7d, 0x78, 0xfc, 0x54, 0x40, 0x63, 0x08, 0x14, 0x88,
	0xc5, 0x49, 0x6c, 0x60, 0x71, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x08, 0x6f, 0x90, 0x41,
	0xc6, 0x01, 0x00, 0x00,
}

func (m *VaultId) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VaultId) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VaultId) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Number != 0 {
		i = encodeVarintVault(dAtA, i, uint64(m.Number))
		i--
		dAtA[i] = 0x10
	}
	if m.Type != 0 {
		i = encodeVarintVault(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintVault(dAtA []byte, offset int, v uint64) int {
	offset -= sovVault(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *VaultId) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovVault(uint64(m.Type))
	}
	if m.Number != 0 {
		n += 1 + sovVault(uint64(m.Number))
	}
	return n
}

func sovVault(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVault(x uint64) (n int) {
	return sovVault(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VaultId) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVault
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
			return fmt.Errorf("proto: VaultId: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VaultId: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= VaultType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Number", wireType)
			}
			m.Number = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Number |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipVault(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVault
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
func skipVault(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVault
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
					return 0, ErrIntOverflowVault
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
					return 0, ErrIntOverflowVault
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
				return 0, ErrInvalidLengthVault
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVault
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVault
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVault        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVault          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVault = fmt.Errorf("proto: unexpected end of group")
)
