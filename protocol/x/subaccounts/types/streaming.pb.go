// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bitoroprotocol/subaccounts/streaming.proto

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

// StreamSubaccountUpdate provides information on a subaccount update. Used in
// the full node GRPC stream.
type StreamSubaccountUpdate struct {
	SubaccountId *SubaccountId `protobuf:"bytes,1,opt,name=subaccount_id,json=subaccountId,proto3" json:"subaccount_id,omitempty"`
	// updated_perpetual_positions will each be for unique perpetuals.
	UpdatedPerpetualPositions []*SubaccountPerpetualPosition `protobuf:"bytes,2,rep,name=updated_perpetual_positions,json=updatedPerpetualPositions,proto3" json:"updated_perpetual_positions,omitempty"`
	// updated_asset_positions will each be for unique assets.
	UpdatedAssetPositions []*SubaccountAssetPosition `protobuf:"bytes,3,rep,name=updated_asset_positions,json=updatedAssetPositions,proto3" json:"updated_asset_positions,omitempty"`
	// Snapshot indicates if the response is from a snapshot of the subaccount.
	// All updates should be ignored until snapshot is received.
	// If the snapshot is true, then all previous entries should be
	// discarded and the subaccount should be resynced.
	// For a snapshot subaccount update, the `updated_perpetual_positions` and
	// `updated_asset_positions` fields will contain the full state of the
	// subaccount.
	Snapshot bool `protobuf:"varint,4,opt,name=snapshot,proto3" json:"snapshot,omitempty"`
}

func (m *StreamSubaccountUpdate) Reset()         { *m = StreamSubaccountUpdate{} }
func (m *StreamSubaccountUpdate) String() string { return proto.CompactTextString(m) }
func (*StreamSubaccountUpdate) ProtoMessage()    {}
func (*StreamSubaccountUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6cf3092946c3c13, []int{0}
}
func (m *StreamSubaccountUpdate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StreamSubaccountUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StreamSubaccountUpdate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StreamSubaccountUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamSubaccountUpdate.Merge(m, src)
}
func (m *StreamSubaccountUpdate) XXX_Size() int {
	return m.Size()
}
func (m *StreamSubaccountUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamSubaccountUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_StreamSubaccountUpdate proto.InternalMessageInfo

func (m *StreamSubaccountUpdate) GetSubaccountId() *SubaccountId {
	if m != nil {
		return m.SubaccountId
	}
	return nil
}

func (m *StreamSubaccountUpdate) GetUpdatedPerpetualPositions() []*SubaccountPerpetualPosition {
	if m != nil {
		return m.UpdatedPerpetualPositions
	}
	return nil
}

func (m *StreamSubaccountUpdate) GetUpdatedAssetPositions() []*SubaccountAssetPosition {
	if m != nil {
		return m.UpdatedAssetPositions
	}
	return nil
}

func (m *StreamSubaccountUpdate) GetSnapshot() bool {
	if m != nil {
		return m.Snapshot
	}
	return false
}

// SubaccountPerpetualPosition provides information on a subaccount's updated
// perpetual positions.
type SubaccountPerpetualPosition struct {
	// The `Id` of the `Perpetual`.
	PerpetualId uint32 `protobuf:"varint,1,opt,name=perpetual_id,json=perpetualId,proto3" json:"perpetual_id,omitempty"`
	// The size of the position in base quantums. Negative means short.
	Quantums int64 `protobuf:"varint,2,opt,name=quantums,proto3" json:"quantums,omitempty"`
}

func (m *SubaccountPerpetualPosition) Reset()         { *m = SubaccountPerpetualPosition{} }
func (m *SubaccountPerpetualPosition) String() string { return proto.CompactTextString(m) }
func (*SubaccountPerpetualPosition) ProtoMessage()    {}
func (*SubaccountPerpetualPosition) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6cf3092946c3c13, []int{1}
}
func (m *SubaccountPerpetualPosition) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SubaccountPerpetualPosition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SubaccountPerpetualPosition.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SubaccountPerpetualPosition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubaccountPerpetualPosition.Merge(m, src)
}
func (m *SubaccountPerpetualPosition) XXX_Size() int {
	return m.Size()
}
func (m *SubaccountPerpetualPosition) XXX_DiscardUnknown() {
	xxx_messageInfo_SubaccountPerpetualPosition.DiscardUnknown(m)
}

var xxx_messageInfo_SubaccountPerpetualPosition proto.InternalMessageInfo

func (m *SubaccountPerpetualPosition) GetPerpetualId() uint32 {
	if m != nil {
		return m.PerpetualId
	}
	return 0
}

func (m *SubaccountPerpetualPosition) GetQuantums() int64 {
	if m != nil {
		return m.Quantums
	}
	return 0
}

// SubaccountAssetPosition provides information on a subaccount's updated asset
// positions.
type SubaccountAssetPosition struct {
	// The `Id` of the `Asset`.
	AssetId uint32 `protobuf:"varint,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	// The absolute size of the position in base quantums.
	Quantums uint64 `protobuf:"varint,2,opt,name=quantums,proto3" json:"quantums,omitempty"`
}

func (m *SubaccountAssetPosition) Reset()         { *m = SubaccountAssetPosition{} }
func (m *SubaccountAssetPosition) String() string { return proto.CompactTextString(m) }
func (*SubaccountAssetPosition) ProtoMessage()    {}
func (*SubaccountAssetPosition) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6cf3092946c3c13, []int{2}
}
func (m *SubaccountAssetPosition) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SubaccountAssetPosition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SubaccountAssetPosition.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SubaccountAssetPosition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubaccountAssetPosition.Merge(m, src)
}
func (m *SubaccountAssetPosition) XXX_Size() int {
	return m.Size()
}
func (m *SubaccountAssetPosition) XXX_DiscardUnknown() {
	xxx_messageInfo_SubaccountAssetPosition.DiscardUnknown(m)
}

var xxx_messageInfo_SubaccountAssetPosition proto.InternalMessageInfo

func (m *SubaccountAssetPosition) GetAssetId() uint32 {
	if m != nil {
		return m.AssetId
	}
	return 0
}

func (m *SubaccountAssetPosition) GetQuantums() uint64 {
	if m != nil {
		return m.Quantums
	}
	return 0
}

func init() {
	proto.RegisterType((*StreamSubaccountUpdate)(nil), "bitoroprotocol.subaccounts.StreamSubaccountUpdate")
	proto.RegisterType((*SubaccountPerpetualPosition)(nil), "bitoroprotocol.subaccounts.SubaccountPerpetualPosition")
	proto.RegisterType((*SubaccountAssetPosition)(nil), "bitoroprotocol.subaccounts.SubaccountAssetPosition")
}

func init() {
	proto.RegisterFile("bitoroprotocol/subaccounts/streaming.proto", fileDescriptor_e6cf3092946c3c13)
}

var fileDescriptor_e6cf3092946c3c13 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x86, 0x29, 0x10, 0x25, 0x0b, 0x5c, 0x36, 0x51, 0x0a, 0x24, 0x0d, 0x72, 0x30, 0xf5, 0x60,
	0x1b, 0x51, 0x8f, 0x1e, 0xf4, 0x46, 0xbc, 0x90, 0x12, 0x63, 0x62, 0x4c, 0xc8, 0xd2, 0x6d, 0x60,
	0x13, 0xe8, 0xae, 0xcc, 0xae, 0x81, 0xb7, 0xf0, 0xb1, 0x3c, 0x72, 0xf4, 0x68, 0xe8, 0x8b, 0x18,
	0x4a, 0xd9, 0x96, 0x28, 0x86, 0x5b, 0xff, 0xf9, 0x67, 0xbe, 0xbf, 0x3b, 0x19, 0x64, 0xd3, 0x05,
	0x9d, 0x8b, 0x19, 0x97, 0xdc, 0xe7, 0x13, 0x17, 0xd4, 0x90, 0xf8, 0x3e, 0x57, 0xa1, 0x04, 0x17,
	0xe4, 0x2c, 0x20, 0x53, 0x16, 0x8e, 0x9c, 0xd8, 0xc6, 0x66, 0xb6, 0xd3, 0xc9, 0x74, 0x36, 0x2e,
	0xf6, 0x33, 0xf4, 0xf7, 0x06, 0xd2, 0x8e, 0xf2, 0xe8, 0xb4, 0x1f, 0x83, 0xfb, 0xda, 0x7a, 0x12,
	0x94, 0xc8, 0x00, 0x3f, 0xa2, 0x6a, 0xda, 0x3e, 0x60, 0xd4, 0x34, 0x5a, 0x86, 0x5d, 0xee, 0x9c,
	0x3b, 0xfb, 0x72, 0x9d, 0x14, 0xd1, 0xa5, 0x5e, 0x05, 0x32, 0x0a, 0x2b, 0xd4, 0x54, 0x31, 0x96,
	0x0e, 0x44, 0x30, 0x13, 0x81, 0x54, 0x64, 0x32, 0x10, 0x1c, 0x98, 0x64, 0x3c, 0x04, 0x33, 0xdf,
	0x2a, 0xd8, 0xe5, 0xce, 0xed, 0x21, 0xe8, 0xde, 0x76, 0xbc, 0x97, 0x4c, 0x7b, 0xf5, 0x84, 0xfc,
	0xcb, 0x01, 0xcc, 0x50, 0x6d, 0x1b, 0x4b, 0x00, 0x02, 0x99, 0x89, 0x2c, 0xc4, 0x91, 0x57, 0x87,
	0x44, 0xde, 0xaf, 0x47, 0x75, 0xdc, 0x49, 0x42, 0xdc, 0xa9, 0x02, 0x6e, 0xa0, 0x12, 0x84, 0x44,
	0xc0, 0x98, 0x4b, 0xb3, 0xd8, 0x32, 0xec, 0x92, 0xa7, 0x75, 0xfb, 0x15, 0x35, 0xff, 0x79, 0x00,
	0x3e, 0x43, 0x95, 0x74, 0x29, 0xc9, 0xa2, 0xab, 0x5e, 0x59, 0xd7, 0xba, 0x74, 0x4d, 0x7f, 0x53,
	0x24, 0x94, 0x6a, 0xba, 0x5e, 0x96, 0x61, 0x17, 0x3c, 0xad, 0xdb, 0x3d, 0x54, 0xdb, 0xf3, 0xaf,
	0xb8, 0x8e, 0x4a, 0x9b, 0x77, 0x6b, 0xea, 0x71, 0xac, 0xff, 0x20, 0x16, 0x53, 0xe2, 0xc3, 0xf3,
	0xe7, 0xca, 0x32, 0x96, 0x2b, 0xcb, 0xf8, 0x5e, 0x59, 0xc6, 0x47, 0x64, 0xe5, 0x96, 0x91, 0x95,
	0xfb, 0x8a, 0xac, 0xdc, 0xcb, 0xdd, 0x88, 0xc9, 0xb1, 0x1a, 0x3a, 0x3e, 0x9f, 0xba, 0x3b, 0x57,
	0xf6, 0x7e, 0x73, 0xe9, 0x8f, 0x09, 0x0b, 0x5d, 0x5d, 0x99, 0xef, 0x5c, 0x9e, 0x5c, 0x88, 0x00,
	0x86, 0x47, 0xb1, 0x7b, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x00, 0x93, 0x3a, 0xe6, 0x02,
	0x00, 0x00,
}

func (m *StreamSubaccountUpdate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StreamSubaccountUpdate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StreamSubaccountUpdate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Snapshot {
		i--
		if m.Snapshot {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.UpdatedAssetPositions) > 0 {
		for iNdEx := len(m.UpdatedAssetPositions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UpdatedAssetPositions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStreaming(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.UpdatedPerpetualPositions) > 0 {
		for iNdEx := len(m.UpdatedPerpetualPositions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UpdatedPerpetualPositions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStreaming(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.SubaccountId != nil {
		{
			size, err := m.SubaccountId.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintStreaming(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SubaccountPerpetualPosition) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SubaccountPerpetualPosition) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SubaccountPerpetualPosition) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Quantums != 0 {
		i = encodeVarintStreaming(dAtA, i, uint64(m.Quantums))
		i--
		dAtA[i] = 0x10
	}
	if m.PerpetualId != 0 {
		i = encodeVarintStreaming(dAtA, i, uint64(m.PerpetualId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SubaccountAssetPosition) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SubaccountAssetPosition) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SubaccountAssetPosition) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Quantums != 0 {
		i = encodeVarintStreaming(dAtA, i, uint64(m.Quantums))
		i--
		dAtA[i] = 0x10
	}
	if m.AssetId != 0 {
		i = encodeVarintStreaming(dAtA, i, uint64(m.AssetId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintStreaming(dAtA []byte, offset int, v uint64) int {
	offset -= sovStreaming(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StreamSubaccountUpdate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SubaccountId != nil {
		l = m.SubaccountId.Size()
		n += 1 + l + sovStreaming(uint64(l))
	}
	if len(m.UpdatedPerpetualPositions) > 0 {
		for _, e := range m.UpdatedPerpetualPositions {
			l = e.Size()
			n += 1 + l + sovStreaming(uint64(l))
		}
	}
	if len(m.UpdatedAssetPositions) > 0 {
		for _, e := range m.UpdatedAssetPositions {
			l = e.Size()
			n += 1 + l + sovStreaming(uint64(l))
		}
	}
	if m.Snapshot {
		n += 2
	}
	return n
}

func (m *SubaccountPerpetualPosition) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PerpetualId != 0 {
		n += 1 + sovStreaming(uint64(m.PerpetualId))
	}
	if m.Quantums != 0 {
		n += 1 + sovStreaming(uint64(m.Quantums))
	}
	return n
}

func (m *SubaccountAssetPosition) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AssetId != 0 {
		n += 1 + sovStreaming(uint64(m.AssetId))
	}
	if m.Quantums != 0 {
		n += 1 + sovStreaming(uint64(m.Quantums))
	}
	return n
}

func sovStreaming(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStreaming(x uint64) (n int) {
	return sovStreaming(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StreamSubaccountUpdate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStreaming
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
			return fmt.Errorf("proto: StreamSubaccountUpdate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StreamSubaccountUpdate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubaccountId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreaming
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
				return ErrInvalidLengthStreaming
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStreaming
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SubaccountId == nil {
				m.SubaccountId = &SubaccountId{}
			}
			if err := m.SubaccountId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedPerpetualPositions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreaming
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
				return ErrInvalidLengthStreaming
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStreaming
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UpdatedPerpetualPositions = append(m.UpdatedPerpetualPositions, &SubaccountPerpetualPosition{})
			if err := m.UpdatedPerpetualPositions[len(m.UpdatedPerpetualPositions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAssetPositions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreaming
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
				return ErrInvalidLengthStreaming
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStreaming
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UpdatedAssetPositions = append(m.UpdatedAssetPositions, &SubaccountAssetPosition{})
			if err := m.UpdatedAssetPositions[len(m.UpdatedAssetPositions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Snapshot", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreaming
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Snapshot = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipStreaming(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStreaming
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
func (m *SubaccountPerpetualPosition) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStreaming
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
			return fmt.Errorf("proto: SubaccountPerpetualPosition: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubaccountPerpetualPosition: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PerpetualId", wireType)
			}
			m.PerpetualId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreaming
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PerpetualId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quantums", wireType)
			}
			m.Quantums = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreaming
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Quantums |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStreaming(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStreaming
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
func (m *SubaccountAssetPosition) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStreaming
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
			return fmt.Errorf("proto: SubaccountAssetPosition: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubaccountAssetPosition: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetId", wireType)
			}
			m.AssetId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreaming
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quantums", wireType)
			}
			m.Quantums = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStreaming
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Quantums |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStreaming(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStreaming
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
func skipStreaming(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStreaming
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
					return 0, ErrIntOverflowStreaming
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
					return 0, ErrIntOverflowStreaming
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
				return 0, ErrInvalidLengthStreaming
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStreaming
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStreaming
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStreaming        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStreaming          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStreaming = fmt.Errorf("proto: unexpected end of group")
)
