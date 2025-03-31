// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bitoroprotocol/clob/streaming.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	types1 "github.com/Bitoro-Network/chain/protocol/x/prices/types"
	types "github.com/Bitoro-Network/chain/protocol/x/subaccounts/types"
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

// StagedFinalizeBlockEvent is an event staged during `FinalizeBlock`.
type StagedFinalizeBlockEvent struct {
	// Contains one of StreamOrderbookFill, StreamSubaccountUpdate.
	//
	// Types that are valid to be assigned to Event:
	//	*StagedFinalizeBlockEvent_OrderFill
	//	*StagedFinalizeBlockEvent_SubaccountUpdate
	//	*StagedFinalizeBlockEvent_OrderbookUpdate
	//	*StagedFinalizeBlockEvent_PriceUpdate
	Event isStagedFinalizeBlockEvent_Event `protobuf_oneof:"event"`
}

func (m *StagedFinalizeBlockEvent) Reset()         { *m = StagedFinalizeBlockEvent{} }
func (m *StagedFinalizeBlockEvent) String() string { return proto.CompactTextString(m) }
func (*StagedFinalizeBlockEvent) ProtoMessage()    {}
func (*StagedFinalizeBlockEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_cecf6ffcf2554dee, []int{0}
}
func (m *StagedFinalizeBlockEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StagedFinalizeBlockEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StagedFinalizeBlockEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StagedFinalizeBlockEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StagedFinalizeBlockEvent.Merge(m, src)
}
func (m *StagedFinalizeBlockEvent) XXX_Size() int {
	return m.Size()
}
func (m *StagedFinalizeBlockEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_StagedFinalizeBlockEvent.DiscardUnknown(m)
}

var xxx_messageInfo_StagedFinalizeBlockEvent proto.InternalMessageInfo

type isStagedFinalizeBlockEvent_Event interface {
	isStagedFinalizeBlockEvent_Event()
	MarshalTo([]byte) (int, error)
	Size() int
}

type StagedFinalizeBlockEvent_OrderFill struct {
	OrderFill *StreamOrderbookFill `protobuf:"bytes,1,opt,name=order_fill,json=orderFill,proto3,oneof" json:"order_fill,omitempty"`
}
type StagedFinalizeBlockEvent_SubaccountUpdate struct {
	SubaccountUpdate *types.StreamSubaccountUpdate `protobuf:"bytes,2,opt,name=subaccount_update,json=subaccountUpdate,proto3,oneof" json:"subaccount_update,omitempty"`
}
type StagedFinalizeBlockEvent_OrderbookUpdate struct {
	OrderbookUpdate *StreamOrderbookUpdate `protobuf:"bytes,3,opt,name=orderbook_update,json=orderbookUpdate,proto3,oneof" json:"orderbook_update,omitempty"`
}
type StagedFinalizeBlockEvent_PriceUpdate struct {
	PriceUpdate *types1.StreamPriceUpdate `protobuf:"bytes,4,opt,name=price_update,json=priceUpdate,proto3,oneof" json:"price_update,omitempty"`
}

func (*StagedFinalizeBlockEvent_OrderFill) isStagedFinalizeBlockEvent_Event()        {}
func (*StagedFinalizeBlockEvent_SubaccountUpdate) isStagedFinalizeBlockEvent_Event() {}
func (*StagedFinalizeBlockEvent_OrderbookUpdate) isStagedFinalizeBlockEvent_Event()  {}
func (*StagedFinalizeBlockEvent_PriceUpdate) isStagedFinalizeBlockEvent_Event()      {}

func (m *StagedFinalizeBlockEvent) GetEvent() isStagedFinalizeBlockEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *StagedFinalizeBlockEvent) GetOrderFill() *StreamOrderbookFill {
	if x, ok := m.GetEvent().(*StagedFinalizeBlockEvent_OrderFill); ok {
		return x.OrderFill
	}
	return nil
}

func (m *StagedFinalizeBlockEvent) GetSubaccountUpdate() *types.StreamSubaccountUpdate {
	if x, ok := m.GetEvent().(*StagedFinalizeBlockEvent_SubaccountUpdate); ok {
		return x.SubaccountUpdate
	}
	return nil
}

func (m *StagedFinalizeBlockEvent) GetOrderbookUpdate() *StreamOrderbookUpdate {
	if x, ok := m.GetEvent().(*StagedFinalizeBlockEvent_OrderbookUpdate); ok {
		return x.OrderbookUpdate
	}
	return nil
}

func (m *StagedFinalizeBlockEvent) GetPriceUpdate() *types1.StreamPriceUpdate {
	if x, ok := m.GetEvent().(*StagedFinalizeBlockEvent_PriceUpdate); ok {
		return x.PriceUpdate
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StagedFinalizeBlockEvent) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StagedFinalizeBlockEvent_OrderFill)(nil),
		(*StagedFinalizeBlockEvent_SubaccountUpdate)(nil),
		(*StagedFinalizeBlockEvent_OrderbookUpdate)(nil),
		(*StagedFinalizeBlockEvent_PriceUpdate)(nil),
	}
}

func init() {
	proto.RegisterType((*StagedFinalizeBlockEvent)(nil), "bitoroprotocol.clob.StagedFinalizeBlockEvent")
}

func init() { proto.RegisterFile("bitoroprotocol/clob/streaming.proto", fileDescriptor_cecf6ffcf2554dee) }

var fileDescriptor_cecf6ffcf2554dee = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0xa9, 0x4c, 0xa9,
	0x28, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x4f, 0xce, 0xc9, 0x4f, 0xd2, 0x2f, 0x2e,
	0x29, 0x4a, 0x4d, 0xcc, 0xcd, 0xcc, 0x4b, 0xd7, 0x03, 0x8b, 0x0b, 0x09, 0x22, 0x2b, 0xd1, 0x03,
	0x29, 0x91, 0xd2, 0x40, 0xd1, 0x55, 0x5c, 0x9a, 0x94, 0x98, 0x9c, 0x9c, 0x5f, 0x9a, 0x57, 0x52,
	0x8c, 0xae, 0x59, 0x4a, 0x19, 0x45, 0x65, 0x41, 0x51, 0x66, 0x72, 0x2a, 0xa6, 0x22, 0x59, 0x4c,
	0x47, 0x14, 0x96, 0xa6, 0x16, 0x55, 0x42, 0xa4, 0x95, 0x3e, 0x32, 0x71, 0x49, 0x04, 0x97, 0x24,
	0xa6, 0xa7, 0xa6, 0xb8, 0x65, 0xe6, 0x25, 0xe6, 0x64, 0x56, 0xa5, 0x3a, 0xe5, 0xe4, 0x27, 0x67,
	0xbb, 0x96, 0xa5, 0xe6, 0x95, 0x08, 0xb9, 0x73, 0x71, 0xe5, 0x17, 0xa5, 0xa4, 0x16, 0xc5, 0xa7,
	0x65, 0xe6, 0xe4, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0xa9, 0xe9, 0x61, 0x38, 0x59, 0x2f,
	0x18, 0x6c, 0xa7, 0x3f, 0x48, 0x69, 0x52, 0x7e, 0x7e, 0xb6, 0x5b, 0x66, 0x4e, 0x8e, 0x07, 0x43,
	0x10, 0x27, 0x58, 0x2f, 0x88, 0x23, 0x14, 0xcf, 0x25, 0x88, 0xf0, 0x48, 0x7c, 0x69, 0x41, 0x4a,
	0x62, 0x49, 0xaa, 0x04, 0x13, 0xd8, 0x3c, 0x03, 0x54, 0xf3, 0x90, 0xfc, 0x0b, 0x35, 0x36, 0x18,
	0x2e, 0x12, 0x0a, 0xd6, 0xe7, 0xc1, 0x10, 0x24, 0x50, 0x8c, 0x26, 0x26, 0x14, 0xca, 0x25, 0x90,
	0x0f, 0xb3, 0x1e, 0x66, 0x3e, 0x33, 0xd8, 0x7c, 0x0d, 0xc2, 0xee, 0x85, 0x9b, 0xcb, 0x9f, 0x8f,
	0x2a, 0x24, 0xe4, 0xcd, 0xc5, 0x03, 0x0e, 0x56, 0x98, 0x91, 0x2c, 0xd8, 0x82, 0x00, 0x12, 0xf0,
	0x50, 0x43, 0x03, 0x40, 0x1c, 0xb8, 0x81, 0xdc, 0x05, 0x08, 0xae, 0x13, 0x3b, 0x17, 0x6b, 0x2a,
	0x28, 0x58, 0x9d, 0x02, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39,
	0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x2c,
	0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x1f, 0x25, 0xde, 0xca, 0x4c, 0x74,
	0x93, 0x33, 0x12, 0x33, 0xf3, 0xf4, 0xe1, 0x22, 0x15, 0x90, 0xb8, 0x2c, 0xa9, 0x2c, 0x48, 0x2d,
	0x4e, 0x62, 0x03, 0x0b, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x8f, 0x58, 0x41, 0x72,
	0x02, 0x00, 0x00,
}

func (m *StagedFinalizeBlockEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StagedFinalizeBlockEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StagedFinalizeBlockEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Event != nil {
		{
			size := m.Event.Size()
			i -= size
			if _, err := m.Event.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *StagedFinalizeBlockEvent_OrderFill) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StagedFinalizeBlockEvent_OrderFill) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.OrderFill != nil {
		{
			size, err := m.OrderFill.MarshalToSizedBuffer(dAtA[:i])
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
func (m *StagedFinalizeBlockEvent_SubaccountUpdate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StagedFinalizeBlockEvent_SubaccountUpdate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.SubaccountUpdate != nil {
		{
			size, err := m.SubaccountUpdate.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintStreaming(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *StagedFinalizeBlockEvent_OrderbookUpdate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StagedFinalizeBlockEvent_OrderbookUpdate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.OrderbookUpdate != nil {
		{
			size, err := m.OrderbookUpdate.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintStreaming(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *StagedFinalizeBlockEvent_PriceUpdate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StagedFinalizeBlockEvent_PriceUpdate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.PriceUpdate != nil {
		{
			size, err := m.PriceUpdate.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintStreaming(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
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
func (m *StagedFinalizeBlockEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Event != nil {
		n += m.Event.Size()
	}
	return n
}

func (m *StagedFinalizeBlockEvent_OrderFill) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OrderFill != nil {
		l = m.OrderFill.Size()
		n += 1 + l + sovStreaming(uint64(l))
	}
	return n
}
func (m *StagedFinalizeBlockEvent_SubaccountUpdate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SubaccountUpdate != nil {
		l = m.SubaccountUpdate.Size()
		n += 1 + l + sovStreaming(uint64(l))
	}
	return n
}
func (m *StagedFinalizeBlockEvent_OrderbookUpdate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OrderbookUpdate != nil {
		l = m.OrderbookUpdate.Size()
		n += 1 + l + sovStreaming(uint64(l))
	}
	return n
}
func (m *StagedFinalizeBlockEvent_PriceUpdate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PriceUpdate != nil {
		l = m.PriceUpdate.Size()
		n += 1 + l + sovStreaming(uint64(l))
	}
	return n
}

func sovStreaming(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStreaming(x uint64) (n int) {
	return sovStreaming(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StagedFinalizeBlockEvent) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: StagedFinalizeBlockEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StagedFinalizeBlockEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderFill", wireType)
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
			v := &StreamOrderbookFill{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Event = &StagedFinalizeBlockEvent_OrderFill{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubaccountUpdate", wireType)
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
			v := &types.StreamSubaccountUpdate{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Event = &StagedFinalizeBlockEvent_SubaccountUpdate{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderbookUpdate", wireType)
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
			v := &StreamOrderbookUpdate{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Event = &StagedFinalizeBlockEvent_OrderbookUpdate{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceUpdate", wireType)
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
			v := &types1.StreamPriceUpdate{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Event = &StagedFinalizeBlockEvent_PriceUpdate{v}
			iNdEx = postIndex
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
