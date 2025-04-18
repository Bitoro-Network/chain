// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bitoroprotocol/vault/share.proto

package types

import (
	fmt "fmt"
	github_com_bitoro_network_chain_protocol_dtypes "github.com/bitoro-network/chain/protocol/dtypes"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// NumShares represents the number of shares.
type NumShares struct {
	// Number of shares.
	NumShares github_com_bitoro_network_chain_protocol_dtypes.SerializableInt `protobuf:"bytes,2,opt,name=num_shares,json=numShares,proto3,customtype=github.com/bitoro-network/chain/protocol/dtypes.SerializableInt" json:"num_shares"`
}

func (m *NumShares) Reset()         { *m = NumShares{} }
func (m *NumShares) String() string { return proto.CompactTextString(m) }
func (*NumShares) ProtoMessage()    {}
func (*NumShares) Descriptor() ([]byte, []int) {
	return fileDescriptor_0db9de1ef13fedaf, []int{0}
}
func (m *NumShares) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NumShares) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NumShares.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NumShares) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NumShares.Merge(m, src)
}
func (m *NumShares) XXX_Size() int {
	return m.Size()
}
func (m *NumShares) XXX_DiscardUnknown() {
	xxx_messageInfo_NumShares.DiscardUnknown(m)
}

var xxx_messageInfo_NumShares proto.InternalMessageInfo

// OwnerShare is a type for owner shares.
type OwnerShare struct {
	Owner  string    `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Shares NumShares `protobuf:"bytes,2,opt,name=shares,proto3" json:"shares"`
}

func (m *OwnerShare) Reset()         { *m = OwnerShare{} }
func (m *OwnerShare) String() string { return proto.CompactTextString(m) }
func (*OwnerShare) ProtoMessage()    {}
func (*OwnerShare) Descriptor() ([]byte, []int) {
	return fileDescriptor_0db9de1ef13fedaf, []int{1}
}
func (m *OwnerShare) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OwnerShare) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OwnerShare.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OwnerShare) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OwnerShare.Merge(m, src)
}
func (m *OwnerShare) XXX_Size() int {
	return m.Size()
}
func (m *OwnerShare) XXX_DiscardUnknown() {
	xxx_messageInfo_OwnerShare.DiscardUnknown(m)
}

var xxx_messageInfo_OwnerShare proto.InternalMessageInfo

func (m *OwnerShare) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *OwnerShare) GetShares() NumShares {
	if m != nil {
		return m.Shares
	}
	return NumShares{}
}

// OwnerShareUnlocks stores share unlocks for an owner.
type OwnerShareUnlocks struct {
	// Address of the owner of below shares.
	OwnerAddress string `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	// All share unlocks.
	ShareUnlocks []ShareUnlock `protobuf:"bytes,2,rep,name=share_unlocks,json=shareUnlocks,proto3" json:"share_unlocks"`
}

func (m *OwnerShareUnlocks) Reset()         { *m = OwnerShareUnlocks{} }
func (m *OwnerShareUnlocks) String() string { return proto.CompactTextString(m) }
func (*OwnerShareUnlocks) ProtoMessage()    {}
func (*OwnerShareUnlocks) Descriptor() ([]byte, []int) {
	return fileDescriptor_0db9de1ef13fedaf, []int{2}
}
func (m *OwnerShareUnlocks) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OwnerShareUnlocks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OwnerShareUnlocks.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OwnerShareUnlocks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OwnerShareUnlocks.Merge(m, src)
}
func (m *OwnerShareUnlocks) XXX_Size() int {
	return m.Size()
}
func (m *OwnerShareUnlocks) XXX_DiscardUnknown() {
	xxx_messageInfo_OwnerShareUnlocks.DiscardUnknown(m)
}

var xxx_messageInfo_OwnerShareUnlocks proto.InternalMessageInfo

func (m *OwnerShareUnlocks) GetOwnerAddress() string {
	if m != nil {
		return m.OwnerAddress
	}
	return ""
}

func (m *OwnerShareUnlocks) GetShareUnlocks() []ShareUnlock {
	if m != nil {
		return m.ShareUnlocks
	}
	return nil
}

// ShareUnlock stores a single instance of `shares` number of shares
// unlocking at block height `unlock_block_height`.
type ShareUnlock struct {
	// Number of shares to unlock.
	Shares NumShares `protobuf:"bytes,1,opt,name=shares,proto3" json:"shares"`
	// Block height at which above shares unlock.
	UnlockBlockHeight uint32 `protobuf:"varint,2,opt,name=unlock_block_height,json=unlockBlockHeight,proto3" json:"unlock_block_height,omitempty"`
}

func (m *ShareUnlock) Reset()         { *m = ShareUnlock{} }
func (m *ShareUnlock) String() string { return proto.CompactTextString(m) }
func (*ShareUnlock) ProtoMessage()    {}
func (*ShareUnlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_0db9de1ef13fedaf, []int{3}
}
func (m *ShareUnlock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ShareUnlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ShareUnlock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ShareUnlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShareUnlock.Merge(m, src)
}
func (m *ShareUnlock) XXX_Size() int {
	return m.Size()
}
func (m *ShareUnlock) XXX_DiscardUnknown() {
	xxx_messageInfo_ShareUnlock.DiscardUnknown(m)
}

var xxx_messageInfo_ShareUnlock proto.InternalMessageInfo

func (m *ShareUnlock) GetShares() NumShares {
	if m != nil {
		return m.Shares
	}
	return NumShares{}
}

func (m *ShareUnlock) GetUnlockBlockHeight() uint32 {
	if m != nil {
		return m.UnlockBlockHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*NumShares)(nil), "bitoroprotocol.vault.NumShares")
	proto.RegisterType((*OwnerShare)(nil), "bitoroprotocol.vault.OwnerShare")
	proto.RegisterType((*OwnerShareUnlocks)(nil), "bitoroprotocol.vault.OwnerShareUnlocks")
	proto.RegisterType((*ShareUnlock)(nil), "bitoroprotocol.vault.ShareUnlock")
}

func init() { proto.RegisterFile("bitoroprotocol/vault/share.proto", fileDescriptor_0db9de1ef13fedaf) }

var fileDescriptor_0db9de1ef13fedaf = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xbf, 0x6e, 0xdb, 0x30,
	0x10, 0xc6, 0x45, 0xb7, 0x35, 0x60, 0xda, 0x1e, 0xac, 0x7a, 0x50, 0x3d, 0xc8, 0xaa, 0xba, 0x78,
	0x31, 0x05, 0xb8, 0x40, 0xb7, 0xa2, 0xa8, 0x96, 0xb6, 0x40, 0xd1, 0x16, 0x32, 0xba, 0x74, 0x11,
	0x24, 0x99, 0x95, 0x04, 0x4b, 0xa4, 0x41, 0x52, 0x71, 0xfe, 0xcd, 0x99, 0xf3, 0x30, 0x79, 0x08,
	0x8f, 0x46, 0xa6, 0x20, 0x83, 0x11, 0xd8, 0x2f, 0x12, 0x88, 0x54, 0x1c, 0x05, 0xf0, 0x10, 0x64,
	0x21, 0x74, 0x77, 0x1f, 0xef, 0xfb, 0x1d, 0x4f, 0xd0, 0x0a, 0x53, 0x41, 0x19, 0x5d, 0x30, 0x2a,
	0x68, 0x44, 0x33, 0xe7, 0x28, 0x28, 0x32, 0xe1, 0xf0, 0x24, 0x60, 0x18, 0xc9, 0xa4, 0xde, 0x7f,
	0xaa, 0x40, 0x52, 0x31, 0xe8, 0xc7, 0x34, 0x56, 0x39, 0xa7, 0xfc, 0x52, 0xda, 0xc1, 0xbb, 0x88,
	0xf2, 0x9c, 0x72, 0x5f, 0x15, 0x54, 0xa0, 0x4a, 0x36, 0x87, 0xad, 0x5f, 0x45, 0x3e, 0x2d, 0x1b,
	0x73, 0xfd, 0x3f, 0x84, 0xa4, 0xc8, 0x7d, 0x69, 0xc3, 0x8d, 0x86, 0x05, 0x46, 0x1d, 0xf7, 0xdb,
	0x6a, 0x33, 0xd4, 0x6e, 0x37, 0xc3, 0x2f, 0x71, 0x2a, 0x92, 0x22, 0x44, 0x11, 0xcd, 0x1d, 0x65,
	0x3d, 0x26, 0x58, 0x2c, 0x29, 0x9b, 0x3b, 0x51, 0x12, 0xa4, 0xc4, 0xd9, 0xb3, 0xce, 0xc4, 0xc9,
	0x02, 0x73, 0x34, 0xc5, 0x2c, 0x0d, 0xb2, 0xf4, 0x34, 0x08, 0x33, 0xfc, 0x83, 0x08, 0xaf, 0x45,
	0x1e, 0x7c, 0xec, 0x33, 0x08, 0x7f, 0x2f, 0x09, 0x66, 0x32, 0xd4, 0x11, 0x7c, 0x43, 0xcb, 0xc8,
	0x00, 0x16, 0x18, 0xb5, 0x5c, 0xe3, 0xfa, 0x6a, 0xdc, 0xaf, 0x18, 0xbf, 0xce, 0x66, 0x0c, 0x73,
	0x3e, 0x15, 0x2c, 0x25, 0xb1, 0xa7, 0x64, 0xfa, 0x67, 0xd8, 0xac, 0x11, 0xb6, 0x27, 0x43, 0x74,
	0xe8, 0x29, 0xd0, 0x7e, 0x2c, 0xf7, 0x75, 0x39, 0x82, 0x57, 0x5d, 0xb2, 0x2f, 0x00, 0xec, 0x3d,
	0xba, 0xff, 0x25, 0x19, 0x8d, 0xe6, 0x5c, 0xff, 0x00, 0xbb, 0xb2, 0xbb, 0x1f, 0x28, 0x4b, 0x05,
	0xe3, 0x75, 0x64, 0xb2, 0xc2, 0xd0, 0x7f, 0xc2, 0xae, 0x6c, 0xe2, 0x17, 0xea, 0x96, 0xd1, 0xb0,
	0x5e, 0x8d, 0xda, 0x93, 0xf7, 0x87, 0x01, 0x6a, 0xfd, 0x2b, 0x84, 0x0e, 0xaf, 0x59, 0xda, 0xe7,
	0xb0, 0x5d, 0x93, 0xd4, 0xc6, 0x02, 0x2f, 0x18, 0x4b, 0x47, 0xf0, 0xad, 0xa2, 0xf2, 0x43, 0x79,
	0x26, 0x38, 0x8d, 0x13, 0x21, 0x9f, 0xa8, 0xeb, 0xf5, 0x54, 0xc9, 0x2d, 0x8f, 0xef, 0xb2, 0xe0,
	0xfe, 0x59, 0x6d, 0x4d, 0xb0, 0xde, 0x9a, 0xe0, 0x6e, 0x6b, 0x82, 0xcb, 0x9d, 0xa9, 0xad, 0x77,
	0xa6, 0x76, 0xb3, 0x33, 0xb5, 0x7f, 0x9f, 0x9e, 0xbd, 0xe9, 0xe3, 0xea, 0xbf, 0x94, 0x0b, 0x0f,
	0x9b, 0x32, 0xff, 0xf1, 0x3e, 0x00, 0x00, 0xff, 0xff, 0xeb, 0xa9, 0xce, 0xdb, 0xbc, 0x02, 0x00,
	0x00,
}

func (m *NumShares) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NumShares) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NumShares) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.NumShares.Size()
		i -= size
		if _, err := m.NumShares.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintShare(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}

func (m *OwnerShare) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OwnerShare) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OwnerShare) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Shares.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintShare(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintShare(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OwnerShareUnlocks) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OwnerShareUnlocks) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OwnerShareUnlocks) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ShareUnlocks) > 0 {
		for iNdEx := len(m.ShareUnlocks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ShareUnlocks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintShare(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.OwnerAddress) > 0 {
		i -= len(m.OwnerAddress)
		copy(dAtA[i:], m.OwnerAddress)
		i = encodeVarintShare(dAtA, i, uint64(len(m.OwnerAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ShareUnlock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ShareUnlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ShareUnlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UnlockBlockHeight != 0 {
		i = encodeVarintShare(dAtA, i, uint64(m.UnlockBlockHeight))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.Shares.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintShare(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintShare(dAtA []byte, offset int, v uint64) int {
	offset -= sovShare(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NumShares) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.NumShares.Size()
	n += 1 + l + sovShare(uint64(l))
	return n
}

func (m *OwnerShare) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovShare(uint64(l))
	}
	l = m.Shares.Size()
	n += 1 + l + sovShare(uint64(l))
	return n
}

func (m *OwnerShareUnlocks) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OwnerAddress)
	if l > 0 {
		n += 1 + l + sovShare(uint64(l))
	}
	if len(m.ShareUnlocks) > 0 {
		for _, e := range m.ShareUnlocks {
			l = e.Size()
			n += 1 + l + sovShare(uint64(l))
		}
	}
	return n
}

func (m *ShareUnlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Shares.Size()
	n += 1 + l + sovShare(uint64(l))
	if m.UnlockBlockHeight != 0 {
		n += 1 + sovShare(uint64(m.UnlockBlockHeight))
	}
	return n
}

func sovShare(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozShare(x uint64) (n int) {
	return sovShare(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NumShares) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShare
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
			return fmt.Errorf("proto: NumShares: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NumShares: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumShares", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShare
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
				return ErrInvalidLengthShare
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthShare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NumShares.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipShare(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthShare
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
func (m *OwnerShare) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShare
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
			return fmt.Errorf("proto: OwnerShare: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OwnerShare: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShare
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
				return ErrInvalidLengthShare
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthShare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Shares", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthShare
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthShare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Shares.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipShare(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthShare
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
func (m *OwnerShareUnlocks) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShare
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
			return fmt.Errorf("proto: OwnerShareUnlocks: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OwnerShareUnlocks: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShare
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
				return ErrInvalidLengthShare
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthShare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OwnerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShareUnlocks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthShare
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthShare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ShareUnlocks = append(m.ShareUnlocks, ShareUnlock{})
			if err := m.ShareUnlocks[len(m.ShareUnlocks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipShare(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthShare
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
func (m *ShareUnlock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShare
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
			return fmt.Errorf("proto: ShareUnlock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ShareUnlock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Shares", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthShare
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthShare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Shares.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnlockBlockHeight", wireType)
			}
			m.UnlockBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UnlockBlockHeight |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipShare(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthShare
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
func skipShare(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowShare
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
					return 0, ErrIntOverflowShare
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
					return 0, ErrIntOverflowShare
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
				return 0, ErrInvalidLengthShare
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupShare
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthShare
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthShare        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowShare          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupShare = fmt.Errorf("proto: unexpected end of group")
)
