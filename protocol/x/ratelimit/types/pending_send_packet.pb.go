// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bitoroprotocol/ratelimit/pending_send_packet.proto

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

// PendingSendPacket contains the channel_id and sequence pair to identify a
// pending packet
type PendingSendPacket struct {
	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	Sequence  uint64 `protobuf:"varint,2,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (m *PendingSendPacket) Reset()         { *m = PendingSendPacket{} }
func (m *PendingSendPacket) String() string { return proto.CompactTextString(m) }
func (*PendingSendPacket) ProtoMessage()    {}
func (*PendingSendPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_549681474be80f53, []int{0}
}
func (m *PendingSendPacket) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PendingSendPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PendingSendPacket.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PendingSendPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PendingSendPacket.Merge(m, src)
}
func (m *PendingSendPacket) XXX_Size() int {
	return m.Size()
}
func (m *PendingSendPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_PendingSendPacket.DiscardUnknown(m)
}

var xxx_messageInfo_PendingSendPacket proto.InternalMessageInfo

func (m *PendingSendPacket) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *PendingSendPacket) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func init() {
	proto.RegisterType((*PendingSendPacket)(nil), "bitoroprotocol.ratelimit.PendingSendPacket")
}

func init() {
	proto.RegisterFile("bitoroprotocol/ratelimit/pending_send_packet.proto", fileDescriptor_549681474be80f53)
}

var fileDescriptor_549681474be80f53 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4a, 0xca, 0x2c, 0xc9,
	0x2f, 0xca, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x2f, 0x4a, 0x2c, 0x49, 0xcd,
	0xc9, 0xcc, 0xcd, 0x2c, 0xd1, 0x2f, 0x48, 0xcd, 0x4b, 0xc9, 0xcc, 0x4b, 0x8f, 0x2f, 0x4e, 0xcd,
	0x4b, 0x89, 0x2f, 0x48, 0x4c, 0xce, 0x4e, 0x2d, 0xd1, 0x03, 0x2b, 0x13, 0x92, 0x40, 0xd5, 0xa3,
	0x07, 0xd7, 0xa3, 0xe4, 0xc7, 0x25, 0x18, 0x00, 0xd1, 0x16, 0x9c, 0x9a, 0x97, 0x12, 0x00, 0xd6,
	0x24, 0x24, 0xcb, 0xc5, 0x95, 0x9c, 0x91, 0x98, 0x97, 0x97, 0x9a, 0x13, 0x9f, 0x99, 0x22, 0xc1,
	0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0xc4, 0x09, 0x15, 0xf1, 0x4c, 0x11, 0x92, 0xe2, 0xe2, 0x28, 0x4e,
	0x2d, 0x2c, 0x4d, 0xcd, 0x4b, 0x4e, 0x95, 0x60, 0x52, 0x60, 0xd4, 0x60, 0x09, 0x82, 0xf3, 0x9d,
	0x42, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f,
	0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x2a, 0x3d, 0xb3, 0x24,
	0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x1f, 0xe2, 0x1c, 0xdd, 0xbc, 0xd4, 0x92, 0xf2, 0xfc,
	0xa2, 0x6c, 0xfd, 0xe4, 0x8c, 0xc4, 0xcc, 0x3c, 0x7d, 0xb8, 0x8f, 0x2a, 0x90, 0xfc, 0x54, 0x52,
	0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0x96, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x23, 0xf1,
	0xd4, 0xc8, 0xfc, 0x00, 0x00, 0x00,
}

func (m *PendingSendPacket) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PendingSendPacket) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PendingSendPacket) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sequence != 0 {
		i = encodeVarintPendingSendPacket(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintPendingSendPacket(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPendingSendPacket(dAtA []byte, offset int, v uint64) int {
	offset -= sovPendingSendPacket(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PendingSendPacket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovPendingSendPacket(uint64(l))
	}
	if m.Sequence != 0 {
		n += 1 + sovPendingSendPacket(uint64(m.Sequence))
	}
	return n
}

func sovPendingSendPacket(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPendingSendPacket(x uint64) (n int) {
	return sovPendingSendPacket(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PendingSendPacket) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPendingSendPacket
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
			return fmt.Errorf("proto: PendingSendPacket: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PendingSendPacket: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPendingSendPacket
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
				return ErrInvalidLengthPendingSendPacket
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPendingSendPacket
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPendingSendPacket
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPendingSendPacket(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPendingSendPacket
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
func skipPendingSendPacket(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPendingSendPacket
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
					return 0, ErrIntOverflowPendingSendPacket
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
					return 0, ErrIntOverflowPendingSendPacket
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
				return 0, ErrInvalidLengthPendingSendPacket
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPendingSendPacket
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPendingSendPacket
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPendingSendPacket        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPendingSendPacket          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPendingSendPacket = fmt.Errorf("proto: unexpected end of group")
)
