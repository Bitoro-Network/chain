// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bitoroprotocol/clob/process_proposer_matches_events.proto

package types

import (
	fmt "fmt"
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

// ProcessProposerMatchesEvents is used for communicating which events occurred
// in the last block that require updating the state of the memclob in the
// Commit blocker. It contains information about the following state updates:
// - Long term order IDs that were placed in the last block.
// - Stateful order IDs that were expired in the last block.
// - Order IDs that were filled in the last block.
// - Stateful cancellations order IDs that were placed in the last block.
// - Stateful order IDs forcefully removed in the last block.
// - Conditional order IDs triggered in the last block.
// - Conditional order IDs placed, but not triggered in the last block.
// - The height of the block in which the events occurred.
type ProcessProposerMatchesEvents struct {
	PlacedLongTermOrderIds                  []OrderId `protobuf:"bytes,1,rep,name=placed_long_term_order_ids,json=placedLongTermOrderIds,proto3" json:"placed_long_term_order_ids"` // Deprecated: Do not use.
	ExpiredStatefulOrderIds                 []OrderId `protobuf:"bytes,2,rep,name=expired_stateful_order_ids,json=expiredStatefulOrderIds,proto3" json:"expired_stateful_order_ids"`
	OrderIdsFilledInLastBlock               []OrderId `protobuf:"bytes,3,rep,name=order_ids_filled_in_last_block,json=orderIdsFilledInLastBlock,proto3" json:"order_ids_filled_in_last_block"`
	PlacedStatefulCancellationOrderIds      []OrderId `protobuf:"bytes,4,rep,name=placed_stateful_cancellation_order_ids,json=placedStatefulCancellationOrderIds,proto3" json:"placed_stateful_cancellation_order_ids"` // Deprecated: Do not use.
	RemovedStatefulOrderIds                 []OrderId `protobuf:"bytes,5,rep,name=removed_stateful_order_ids,json=removedStatefulOrderIds,proto3" json:"removed_stateful_order_ids"`
	ConditionalOrderIdsTriggeredInLastBlock []OrderId `protobuf:"bytes,6,rep,name=conditional_order_ids_triggered_in_last_block,json=conditionalOrderIdsTriggeredInLastBlock,proto3" json:"conditional_order_ids_triggered_in_last_block"`
	PlacedConditionalOrderIds               []OrderId `protobuf:"bytes,7,rep,name=placed_conditional_order_ids,json=placedConditionalOrderIds,proto3" json:"placed_conditional_order_ids"` // Deprecated: Do not use.
	BlockHeight                             uint32    `protobuf:"varint,8,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
}

func (m *ProcessProposerMatchesEvents) Reset()         { *m = ProcessProposerMatchesEvents{} }
func (m *ProcessProposerMatchesEvents) String() string { return proto.CompactTextString(m) }
func (*ProcessProposerMatchesEvents) ProtoMessage()    {}
func (*ProcessProposerMatchesEvents) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcc6afea29e9513f, []int{0}
}
func (m *ProcessProposerMatchesEvents) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProcessProposerMatchesEvents) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProcessProposerMatchesEvents.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProcessProposerMatchesEvents) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessProposerMatchesEvents.Merge(m, src)
}
func (m *ProcessProposerMatchesEvents) XXX_Size() int {
	return m.Size()
}
func (m *ProcessProposerMatchesEvents) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessProposerMatchesEvents.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessProposerMatchesEvents proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *ProcessProposerMatchesEvents) GetPlacedLongTermOrderIds() []OrderId {
	if m != nil {
		return m.PlacedLongTermOrderIds
	}
	return nil
}

func (m *ProcessProposerMatchesEvents) GetExpiredStatefulOrderIds() []OrderId {
	if m != nil {
		return m.ExpiredStatefulOrderIds
	}
	return nil
}

func (m *ProcessProposerMatchesEvents) GetOrderIdsFilledInLastBlock() []OrderId {
	if m != nil {
		return m.OrderIdsFilledInLastBlock
	}
	return nil
}

// Deprecated: Do not use.
func (m *ProcessProposerMatchesEvents) GetPlacedStatefulCancellationOrderIds() []OrderId {
	if m != nil {
		return m.PlacedStatefulCancellationOrderIds
	}
	return nil
}

func (m *ProcessProposerMatchesEvents) GetRemovedStatefulOrderIds() []OrderId {
	if m != nil {
		return m.RemovedStatefulOrderIds
	}
	return nil
}

func (m *ProcessProposerMatchesEvents) GetConditionalOrderIdsTriggeredInLastBlock() []OrderId {
	if m != nil {
		return m.ConditionalOrderIdsTriggeredInLastBlock
	}
	return nil
}

// Deprecated: Do not use.
func (m *ProcessProposerMatchesEvents) GetPlacedConditionalOrderIds() []OrderId {
	if m != nil {
		return m.PlacedConditionalOrderIds
	}
	return nil
}

func (m *ProcessProposerMatchesEvents) GetBlockHeight() uint32 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*ProcessProposerMatchesEvents)(nil), "bitoroprotocol.clob.ProcessProposerMatchesEvents")
}

func init() {
	proto.RegisterFile("bitoroprotocol/clob/process_proposer_matches_events.proto", fileDescriptor_bcc6afea29e9513f)
}

var fileDescriptor_bcc6afea29e9513f = []byte{
	// 462 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x63, 0x5a, 0x02, 0xda, 0xc2, 0xc5, 0x20, 0x70, 0xa3, 0xc8, 0x2d, 0x3d, 0x40, 0x2f,
	0xb5, 0x25, 0x10, 0x07, 0xae, 0xa9, 0x40, 0x54, 0x2a, 0x6a, 0x15, 0x7a, 0xe2, 0xb2, 0xac, 0xd7,
	0x53, 0x7b, 0xd5, 0xf5, 0x8e, 0xb5, 0xbb, 0x2d, 0xe1, 0xca, 0x13, 0xf0, 0x2c, 0x3c, 0x45, 0x8f,
	0x3d, 0x72, 0x42, 0x28, 0x79, 0x11, 0x64, 0xef, 0x26, 0x32, 0x4a, 0x0e, 0xc9, 0xcd, 0x1a, 0xcf,
	0x7e, 0xdf, 0xcc, 0x2f, 0x0d, 0x79, 0x97, 0x09, 0x8b, 0x1a, 0x6b, 0x8d, 0x16, 0x39, 0xca, 0x94,
	0x4b, 0xcc, 0xd2, 0x5a, 0x23, 0x07, 0x63, 0x68, 0xad, 0xb1, 0x46, 0x03, 0x9a, 0x56, 0xcc, 0xf2,
	0x12, 0x0c, 0x85, 0x1b, 0x50, 0xd6, 0x24, 0x6d, 0x77, 0xf8, 0xe4, 0xff, 0xa7, 0x49, 0xf3, 0x74,
	0xf0, 0xb4, 0xc0, 0xc2, 0x95, 0xd2, 0xe6, 0xcb, 0xb5, 0x0e, 0xf6, 0x56, 0x59, 0x50, 0xe7, 0xa0,
	0x5d, 0xc3, 0xc1, 0xaf, 0x3e, 0x19, 0x9e, 0x3b, 0xeb, 0xb9, 0x97, 0x7e, 0x72, 0xce, 0xf7, 0xad,
	0x32, 0xfc, 0x4a, 0x06, 0xb5, 0x64, 0x1c, 0x72, 0x2a, 0x51, 0x15, 0xd4, 0x82, 0xae, 0x68, 0x0b,
	0xa0, 0x22, 0x37, 0x51, 0xb0, 0xbf, 0x75, 0xb8, 0xf3, 0x7a, 0x98, 0xac, 0x98, 0x28, 0x39, 0x6b,
	0xba, 0x4e, 0xf2, 0x51, 0xff, 0xf6, 0xcf, 0x5e, 0x2f, 0x0a, 0xc6, 0xcf, 0x1c, 0xe7, 0x14, 0x55,
	0x71, 0x01, 0xba, 0xf2, 0xbf, 0x4d, 0x48, 0xc9, 0x00, 0x26, 0xb5, 0xd0, 0x90, 0x53, 0x63, 0x99,
	0x85, 0xcb, 0x6b, 0xd9, 0x31, 0xdc, 0x5b, 0xc3, 0xb0, 0xdd, 0x18, 0xc6, 0xcf, 0x3d, 0xe5, 0xb3,
	0x87, 0x2c, 0x04, 0x40, 0xe2, 0x05, 0x8f, 0x5e, 0x0a, 0x29, 0x21, 0xa7, 0x42, 0x51, 0xc9, 0x8c,
	0xa5, 0x99, 0x44, 0x7e, 0x15, 0x6d, 0xad, 0x2d, 0xd9, 0x45, 0x4f, 0xfd, 0xd0, 0x72, 0x4e, 0xd4,
	0x29, 0x33, 0x76, 0xd4, 0x40, 0xc2, 0x09, 0x79, 0xe9, 0x93, 0x5a, 0xac, 0xc1, 0x99, 0xe2, 0x20,
	0x25, 0xb3, 0x02, 0x55, 0x67, 0xa7, 0xed, 0x0d, 0x52, 0x3b, 0x70, 0xcc, 0xf9, 0x52, 0xc7, 0x1d,
	0x62, 0x37, 0x41, 0x0d, 0x15, 0xde, 0xac, 0x4e, 0xf0, 0xfe, 0xfa, 0x09, 0x7a, 0xca, 0x52, 0x82,
	0x3f, 0x02, 0x72, 0xc4, 0x51, 0xe5, 0xa2, 0xd1, 0xb2, 0x0e, 0x9c, 0x5a, 0x2d, 0x8a, 0x02, 0xf4,
	0x52, 0xa2, 0xfd, 0xb5, 0xa5, 0xaf, 0x3a, 0xe0, 0xb9, 0xf0, 0x62, 0x4e, 0xed, 0xe6, 0x0b, 0x64,
	0xe8, 0xf3, 0x5d, 0x39, 0x4a, 0xf4, 0x60, 0x83, 0x54, 0x77, 0x1d, 0xe9, 0x78, 0x59, 0x1d, 0xbe,
	0x20, 0x8f, 0xda, 0x15, 0x68, 0x09, 0xa2, 0x28, 0x6d, 0xf4, 0x70, 0x3f, 0x38, 0x7c, 0x3c, 0xde,
	0x69, 0x6b, 0x1f, 0xdb, 0xd2, 0xe8, 0xec, 0x76, 0x1a, 0x07, 0x77, 0xd3, 0x38, 0xf8, 0x3b, 0x8d,
	0x83, 0x9f, 0xb3, 0xb8, 0x77, 0x37, 0x8b, 0x7b, 0xbf, 0x67, 0x71, 0xef, 0xcb, 0xdb, 0x42, 0xd8,
	0xf2, 0x3a, 0x4b, 0x38, 0x56, 0xa9, 0x9b, 0xe3, 0x48, 0x81, 0xfd, 0x86, 0xfa, 0x2a, 0xe5, 0x25,
	0x13, 0x2a, 0x5d, 0x5c, 0xe2, 0xc4, 0xdd, 0xa2, 0xfd, 0x5e, 0x83, 0xc9, 0xfa, 0x6d, 0xf9, 0xcd,
	0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x63, 0x9e, 0x80, 0x15, 0x04, 0x00, 0x00,
}

func (m *ProcessProposerMatchesEvents) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProcessProposerMatchesEvents) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProcessProposerMatchesEvents) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlockHeight != 0 {
		i = encodeVarintProcessProposerMatchesEvents(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x40
	}
	if len(m.PlacedConditionalOrderIds) > 0 {
		for iNdEx := len(m.PlacedConditionalOrderIds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PlacedConditionalOrderIds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProcessProposerMatchesEvents(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.ConditionalOrderIdsTriggeredInLastBlock) > 0 {
		for iNdEx := len(m.ConditionalOrderIdsTriggeredInLastBlock) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ConditionalOrderIdsTriggeredInLastBlock[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProcessProposerMatchesEvents(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.RemovedStatefulOrderIds) > 0 {
		for iNdEx := len(m.RemovedStatefulOrderIds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RemovedStatefulOrderIds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProcessProposerMatchesEvents(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.PlacedStatefulCancellationOrderIds) > 0 {
		for iNdEx := len(m.PlacedStatefulCancellationOrderIds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PlacedStatefulCancellationOrderIds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProcessProposerMatchesEvents(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.OrderIdsFilledInLastBlock) > 0 {
		for iNdEx := len(m.OrderIdsFilledInLastBlock) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OrderIdsFilledInLastBlock[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProcessProposerMatchesEvents(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.ExpiredStatefulOrderIds) > 0 {
		for iNdEx := len(m.ExpiredStatefulOrderIds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ExpiredStatefulOrderIds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProcessProposerMatchesEvents(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.PlacedLongTermOrderIds) > 0 {
		for iNdEx := len(m.PlacedLongTermOrderIds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PlacedLongTermOrderIds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProcessProposerMatchesEvents(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintProcessProposerMatchesEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovProcessProposerMatchesEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProcessProposerMatchesEvents) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PlacedLongTermOrderIds) > 0 {
		for _, e := range m.PlacedLongTermOrderIds {
			l = e.Size()
			n += 1 + l + sovProcessProposerMatchesEvents(uint64(l))
		}
	}
	if len(m.ExpiredStatefulOrderIds) > 0 {
		for _, e := range m.ExpiredStatefulOrderIds {
			l = e.Size()
			n += 1 + l + sovProcessProposerMatchesEvents(uint64(l))
		}
	}
	if len(m.OrderIdsFilledInLastBlock) > 0 {
		for _, e := range m.OrderIdsFilledInLastBlock {
			l = e.Size()
			n += 1 + l + sovProcessProposerMatchesEvents(uint64(l))
		}
	}
	if len(m.PlacedStatefulCancellationOrderIds) > 0 {
		for _, e := range m.PlacedStatefulCancellationOrderIds {
			l = e.Size()
			n += 1 + l + sovProcessProposerMatchesEvents(uint64(l))
		}
	}
	if len(m.RemovedStatefulOrderIds) > 0 {
		for _, e := range m.RemovedStatefulOrderIds {
			l = e.Size()
			n += 1 + l + sovProcessProposerMatchesEvents(uint64(l))
		}
	}
	if len(m.ConditionalOrderIdsTriggeredInLastBlock) > 0 {
		for _, e := range m.ConditionalOrderIdsTriggeredInLastBlock {
			l = e.Size()
			n += 1 + l + sovProcessProposerMatchesEvents(uint64(l))
		}
	}
	if len(m.PlacedConditionalOrderIds) > 0 {
		for _, e := range m.PlacedConditionalOrderIds {
			l = e.Size()
			n += 1 + l + sovProcessProposerMatchesEvents(uint64(l))
		}
	}
	if m.BlockHeight != 0 {
		n += 1 + sovProcessProposerMatchesEvents(uint64(m.BlockHeight))
	}
	return n
}

func sovProcessProposerMatchesEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProcessProposerMatchesEvents(x uint64) (n int) {
	return sovProcessProposerMatchesEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProcessProposerMatchesEvents) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProcessProposerMatchesEvents
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
			return fmt.Errorf("proto: ProcessProposerMatchesEvents: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProcessProposerMatchesEvents: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlacedLongTermOrderIds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProcessProposerMatchesEvents
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
				return ErrInvalidLengthProcessProposerMatchesEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProcessProposerMatchesEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlacedLongTermOrderIds = append(m.PlacedLongTermOrderIds, OrderId{})
			if err := m.PlacedLongTermOrderIds[len(m.PlacedLongTermOrderIds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiredStatefulOrderIds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProcessProposerMatchesEvents
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
				return ErrInvalidLengthProcessProposerMatchesEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProcessProposerMatchesEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExpiredStatefulOrderIds = append(m.ExpiredStatefulOrderIds, OrderId{})
			if err := m.ExpiredStatefulOrderIds[len(m.ExpiredStatefulOrderIds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderIdsFilledInLastBlock", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProcessProposerMatchesEvents
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
				return ErrInvalidLengthProcessProposerMatchesEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProcessProposerMatchesEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrderIdsFilledInLastBlock = append(m.OrderIdsFilledInLastBlock, OrderId{})
			if err := m.OrderIdsFilledInLastBlock[len(m.OrderIdsFilledInLastBlock)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlacedStatefulCancellationOrderIds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProcessProposerMatchesEvents
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
				return ErrInvalidLengthProcessProposerMatchesEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProcessProposerMatchesEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlacedStatefulCancellationOrderIds = append(m.PlacedStatefulCancellationOrderIds, OrderId{})
			if err := m.PlacedStatefulCancellationOrderIds[len(m.PlacedStatefulCancellationOrderIds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemovedStatefulOrderIds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProcessProposerMatchesEvents
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
				return ErrInvalidLengthProcessProposerMatchesEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProcessProposerMatchesEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RemovedStatefulOrderIds = append(m.RemovedStatefulOrderIds, OrderId{})
			if err := m.RemovedStatefulOrderIds[len(m.RemovedStatefulOrderIds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConditionalOrderIdsTriggeredInLastBlock", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProcessProposerMatchesEvents
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
				return ErrInvalidLengthProcessProposerMatchesEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProcessProposerMatchesEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConditionalOrderIdsTriggeredInLastBlock = append(m.ConditionalOrderIdsTriggeredInLastBlock, OrderId{})
			if err := m.ConditionalOrderIdsTriggeredInLastBlock[len(m.ConditionalOrderIdsTriggeredInLastBlock)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlacedConditionalOrderIds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProcessProposerMatchesEvents
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
				return ErrInvalidLengthProcessProposerMatchesEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProcessProposerMatchesEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlacedConditionalOrderIds = append(m.PlacedConditionalOrderIds, OrderId{})
			if err := m.PlacedConditionalOrderIds[len(m.PlacedConditionalOrderIds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProcessProposerMatchesEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProcessProposerMatchesEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProcessProposerMatchesEvents
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
func skipProcessProposerMatchesEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProcessProposerMatchesEvents
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
					return 0, ErrIntOverflowProcessProposerMatchesEvents
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
					return 0, ErrIntOverflowProcessProposerMatchesEvents
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
				return 0, ErrInvalidLengthProcessProposerMatchesEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProcessProposerMatchesEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProcessProposerMatchesEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProcessProposerMatchesEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProcessProposerMatchesEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProcessProposerMatchesEvents = fmt.Errorf("proto: unexpected end of group")
)
