// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bitoroprotocol/subaccounts/asset_position.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_bitoroprotocol_chain_protocol_dtypes "github.com/bitoro-network/chain/protocol/dtypes"
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

// AssetPositions define an account’s positions of an `Asset`.
// Therefore they hold any information needed to trade on Spot and Margin.
type AssetPosition struct {
	// The `Id` of the `Asset`.
	AssetId uint32 `protobuf:"varint,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	// The absolute size of the position in base quantums.
	Quantums github_com_bitoroprotocol_chain_protocol_dtypes.SerializableInt `protobuf:"bytes,2,opt,name=quantums,proto3,customtype=github.com/bitoro-network/chain/protocol/dtypes.SerializableInt" json:"quantums"`
	// The `Index` (either `LongIndex` or `ShortIndex`) of the `Asset` the last
	// time this position was settled
	// TODO(DEC-582): pending margin trading being added.
	Index uint64 `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
}

func (m *AssetPosition) Reset()         { *m = AssetPosition{} }
func (m *AssetPosition) String() string { return proto.CompactTextString(m) }
func (*AssetPosition) ProtoMessage()    {}
func (*AssetPosition) Descriptor() ([]byte, []int) {
	return fileDescriptor_21b5b4cb0c7ef07c, []int{0}
}
func (m *AssetPosition) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AssetPosition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AssetPosition.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AssetPosition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetPosition.Merge(m, src)
}
func (m *AssetPosition) XXX_Size() int {
	return m.Size()
}
func (m *AssetPosition) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetPosition.DiscardUnknown(m)
}

var xxx_messageInfo_AssetPosition proto.InternalMessageInfo

func (m *AssetPosition) GetAssetId() uint32 {
	if m != nil {
		return m.AssetId
	}
	return 0
}

func (m *AssetPosition) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func init() {
	proto.RegisterType((*AssetPosition)(nil), "bitoroprotocol.subaccounts.AssetPosition")
}

func init() {
	proto.RegisterFile("bitoroprotocol/subaccounts/asset_position.proto", fileDescriptor_21b5b4cb0c7ef07c)
}

var fileDescriptor_21b5b4cb0c7ef07c = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4d, 0xa9, 0x4c, 0xa9,
	0x28, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x2f, 0x2e, 0x4d, 0x4a, 0x4c, 0x4e, 0xce,
	0x2f, 0xcd, 0x2b, 0x29, 0xd6, 0x4f, 0x2c, 0x2e, 0x4e, 0x2d, 0x89, 0x2f, 0xc8, 0x2f, 0xce, 0x2c,
	0xc9, 0xcc, 0xcf, 0xd3, 0x03, 0xab, 0x11, 0x92, 0x40, 0x56, 0xae, 0x87, 0xa4, 0x5c, 0x4a, 0x24,
	0x3d, 0x3f, 0x3d, 0x1f, 0x2c, 0xa3, 0x0f, 0x62, 0x41, 0xd4, 0x2b, 0x2d, 0x63, 0xe4, 0xe2, 0x75,
	0x04, 0x19, 0x14, 0x00, 0x35, 0x47, 0x48, 0x92, 0x8b, 0x03, 0x62, 0x72, 0x66, 0x8a, 0x04, 0xa3,
	0x02, 0xa3, 0x06, 0x6f, 0x10, 0x3b, 0x98, 0xef, 0x99, 0x22, 0x94, 0xc2, 0xc5, 0x51, 0x58, 0x9a,
	0x98, 0x57, 0x52, 0x9a, 0x5b, 0x2c, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0xe3, 0xe4, 0x71, 0xe2, 0x9e,
	0x3c, 0xc3, 0xad, 0x7b, 0xf2, 0x0e, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9,
	0xfa, 0x28, 0x0e, 0x2e, 0x33, 0xd1, 0x4d, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x87, 0x8b, 0xa4, 0x94,
	0x54, 0x16, 0xa4, 0x16, 0xeb, 0x05, 0xa7, 0x16, 0x65, 0x26, 0xe6, 0x64, 0x56, 0x25, 0x26, 0xe5,
	0xa4, 0x7a, 0xe6, 0x95, 0x04, 0xc1, 0x4d, 0x16, 0x12, 0xe1, 0x62, 0xcd, 0xcc, 0x4b, 0x49, 0xad,
	0x90, 0x60, 0x56, 0x60, 0xd4, 0x60, 0x09, 0x82, 0x70, 0x9c, 0xc2, 0x4f, 0x3c, 0x92, 0x63, 0xbc,
	0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63,
	0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x96, 0x78, 0xbb, 0x2b, 0x50, 0x02, 0x10, 0xec, 0x90, 0x24,
	0x36, 0xb0, 0xac, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x36, 0x84, 0xfb, 0x96, 0x69, 0x01, 0x00,
	0x00,
}

func (m *AssetPosition) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AssetPosition) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AssetPosition) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Index != 0 {
		i = encodeVarintAssetPosition(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.Quantums.Size()
		i -= size
		if _, err := m.Quantums.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAssetPosition(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.AssetId != 0 {
		i = encodeVarintAssetPosition(dAtA, i, uint64(m.AssetId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAssetPosition(dAtA []byte, offset int, v uint64) int {
	offset -= sovAssetPosition(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AssetPosition) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AssetId != 0 {
		n += 1 + sovAssetPosition(uint64(m.AssetId))
	}
	l = m.Quantums.Size()
	n += 1 + l + sovAssetPosition(uint64(l))
	if m.Index != 0 {
		n += 1 + sovAssetPosition(uint64(m.Index))
	}
	return n
}

func sovAssetPosition(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAssetPosition(x uint64) (n int) {
	return sovAssetPosition(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AssetPosition) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAssetPosition
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
			return fmt.Errorf("proto: AssetPosition: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AssetPosition: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetId", wireType)
			}
			m.AssetId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAssetPosition
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quantums", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAssetPosition
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
				return ErrInvalidLengthAssetPosition
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthAssetPosition
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Quantums.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAssetPosition
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAssetPosition(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAssetPosition
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
func skipAssetPosition(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAssetPosition
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
					return 0, ErrIntOverflowAssetPosition
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
					return 0, ErrIntOverflowAssetPosition
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
				return 0, ErrInvalidLengthAssetPosition
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAssetPosition
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAssetPosition
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAssetPosition        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAssetPosition          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAssetPosition = fmt.Errorf("proto: unexpected end of group")
)
