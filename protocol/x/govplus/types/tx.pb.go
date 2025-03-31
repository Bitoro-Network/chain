// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bitoroprotocol/govplus/tx.proto

package types

import (
	context "context"
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_bitoroprotocol_chain_protocol_dtypes "github.com/Bitoro-Network/chain/protocol/dtypes"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MsgSlashValidator is the Msg/SlashValidator request type.
type MsgSlashValidator struct {
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// Consensus address of the validator to slash
	ValidatorAddress string `protobuf:"bytes,2,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	// Colloquially, the height at which the validator is deemed to have
	// misbehaved. In practice, this is the height used to determine the targets
	// of the slash. For example, undelegating after this height will not escape
	// slashing. This height should be set to a recent height at the time of the
	// proposal to prevent delegators from undelegating during the vote period.
	// i.e. infraction_height <= proposal submission height.
	//
	// NB: At the time this message is applied, this height must have occured
	// equal to or less than an unbonding period in the past in order for the
	// slash to be effective.
	// i.e. time(proposal pass height) - time(infraction_height) < unbonding
	// period
	InfractionHeight uint32 `protobuf:"varint,3,opt,name=infraction_height,json=infractionHeight,proto3" json:"infraction_height,omitempty"`
	// Tokens of the validator at the specified height. Used to compute the slash
	// amount. The x/staking HistoricalInfo query endpoint can be used to find
	// this.
	TokensAtInfractionHeight github_com_bitoroprotocol_chain_protocol_dtypes.SerializableInt `protobuf:"bytes,4,opt,name=tokens_at_infraction_height,json=tokensAtInfractionHeight,proto3,customtype=github.com/Bitoro-Network/chain/protocol/dtypes.SerializableInt" json:"tokens_at_infraction_height"`
	// Multiplier for how much of the validator's stake should be slashed.
	// slash_factor * tokens_at_infraction_height = tokens slashed
	SlashFactor cosmossdk_io_math.LegacyDec `protobuf:"bytes,5,opt,name=slash_factor,json=slashFactor,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"slash_factor"`
}

func (m *MsgSlashValidator) Reset()         { *m = MsgSlashValidator{} }
func (m *MsgSlashValidator) String() string { return proto.CompactTextString(m) }
func (*MsgSlashValidator) ProtoMessage()    {}
func (*MsgSlashValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_9394a0d94676770f, []int{0}
}
func (m *MsgSlashValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSlashValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSlashValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSlashValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSlashValidator.Merge(m, src)
}
func (m *MsgSlashValidator) XXX_Size() int {
	return m.Size()
}
func (m *MsgSlashValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSlashValidator.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSlashValidator proto.InternalMessageInfo

func (m *MsgSlashValidator) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgSlashValidator) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}

func (m *MsgSlashValidator) GetInfractionHeight() uint32 {
	if m != nil {
		return m.InfractionHeight
	}
	return 0
}

// MsgSlashValidatorResponse is the Msg/SlashValidator response type.
type MsgSlashValidatorResponse struct {
}

func (m *MsgSlashValidatorResponse) Reset()         { *m = MsgSlashValidatorResponse{} }
func (m *MsgSlashValidatorResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSlashValidatorResponse) ProtoMessage()    {}
func (*MsgSlashValidatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9394a0d94676770f, []int{1}
}
func (m *MsgSlashValidatorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSlashValidatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSlashValidatorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSlashValidatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSlashValidatorResponse.Merge(m, src)
}
func (m *MsgSlashValidatorResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSlashValidatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSlashValidatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSlashValidatorResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSlashValidator)(nil), "bitoroprotocol.govplus.MsgSlashValidator")
	proto.RegisterType((*MsgSlashValidatorResponse)(nil), "bitoroprotocol.govplus.MsgSlashValidatorResponse")
}

func init() { proto.RegisterFile("bitoroprotocol/govplus/tx.proto", fileDescriptor_9394a0d94676770f) }

var fileDescriptor_9394a0d94676770f = []byte{
	// 483 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xbf, 0x6f, 0xd3, 0x4e,
	0x1c, 0xcd, 0x7d, 0xf3, 0x05, 0xa9, 0x47, 0xa9, 0x1a, 0x2b, 0x12, 0x6e, 0xa2, 0xba, 0xa1, 0x0b,
	0x51, 0x51, 0x7c, 0xe2, 0x87, 0x2a, 0xc1, 0x44, 0xa3, 0x0a, 0xb5, 0x12, 0x65, 0x70, 0x24, 0x24,
	0x58, 0xac, 0xcb, 0xf9, 0x7a, 0x3e, 0x6a, 0xfb, 0x13, 0x7c, 0x17, 0x2b, 0x61, 0x64, 0x61, 0xe5,
	0xcf, 0x60, 0xec, 0x90, 0x95, 0xbd, 0x63, 0x95, 0x09, 0x31, 0x54, 0x28, 0x19, 0xfa, 0x6f, 0xa0,
	0x9c, 0xdd, 0x86, 0x26, 0x0c, 0x5d, 0x6c, 0xdf, 0x7b, 0x9f, 0xf7, 0x9e, 0xf5, 0x3e, 0x36, 0xde,
	0x0c, 0x86, 0xc1, 0xa0, 0x97, 0x82, 0x06, 0x06, 0x11, 0x11, 0x90, 0xf5, 0xa2, 0xbe, 0x22, 0x7a,
	0xe0, 0x1a, 0xcc, 0xaa, 0xfe, 0x4d, 0xbb, 0x05, 0x5d, 0xdb, 0x60, 0xa0, 0x62, 0x50, 0xbe, 0x21,
	0x48, 0x7e, 0xc8, 0x05, 0xb5, 0x07, 0xf9, 0x89, 0xc4, 0x4a, 0x90, 0xec, 0xc9, 0xec, 0x56, 0x10,
	0x55, 0x01, 0x02, 0x72, 0xc1, 0xec, 0xa9, 0x40, 0x2b, 0x34, 0x96, 0x09, 0x10, 0x73, 0xcd, 0xa1,
	0xed, 0x1f, 0x65, 0x5c, 0x39, 0x52, 0xa2, 0x13, 0x51, 0x15, 0xbe, 0xa3, 0x91, 0x0c, 0xa8, 0x86,
	0xd4, 0xda, 0xc5, 0x2b, 0xb4, 0xaf, 0x43, 0x48, 0xa5, 0x1e, 0xda, 0xa8, 0x81, 0x9a, 0x2b, 0x6d,
	0x7b, 0x3c, 0x6a, 0x55, 0x8b, 0xf0, 0xbd, 0x20, 0x48, 0xb9, 0x52, 0x1d, 0x9d, 0xca, 0x44, 0x78,
	0xf3, 0x51, 0xeb, 0x2d, 0xae, 0x64, 0x57, 0x26, 0x3e, 0xcd, 0xa7, 0xec, 0xff, 0x8c, 0xfe, 0xe1,
	0x78, 0xd4, 0xda, 0x2c, 0xf4, 0xd7, 0x41, 0x37, 0x8d, 0xd6, 0xb3, 0x05, 0xdc, 0x7a, 0x8c, 0x2b,
	0x32, 0x39, 0x4e, 0x29, 0xd3, 0x12, 0x12, 0x3f, 0xe4, 0x52, 0x84, 0xda, 0x2e, 0x37, 0x50, 0xf3,
	0xbe, 0xb7, 0x3e, 0x27, 0x0e, 0x0c, 0x6e, 0x7d, 0x45, 0xb8, 0xae, 0xe1, 0x84, 0x27, 0xca, 0xa7,
	0xda, 0x5f, 0xd6, 0xfd, 0xdf, 0x40, 0xcd, 0xd5, 0xf6, 0xc1, 0xd9, 0xc5, 0x56, 0xe9, 0xd7, 0xc5,
	0xd6, 0x2b, 0x21, 0x75, 0xd8, 0xef, 0xba, 0x0c, 0x62, 0x72, 0x63, 0x2b, 0xd9, 0xf3, 0x16, 0x0b,
	0xa9, 0x4c, 0xc8, 0x35, 0x12, 0xe8, 0x61, 0x8f, 0x2b, 0xb7, 0xc3, 0x53, 0x49, 0x23, 0xf9, 0x99,
	0x76, 0x23, 0x7e, 0x98, 0x68, 0xcf, 0xce, 0xc3, 0xf6, 0xf4, 0xe1, 0xe2, 0x9b, 0xbc, 0xc7, 0xab,
	0x6a, 0x56, 0xa8, 0x7f, 0x4c, 0x99, 0x86, 0xd4, 0xbe, 0x63, 0x1a, 0xd8, 0x2d, 0x92, 0xeb, 0x79,
	0x0b, 0x2a, 0x38, 0x71, 0x25, 0x90, 0x98, 0xea, 0xd0, 0x7d, 0xc3, 0x05, 0x65, 0xc3, 0x7d, 0xce,
	0xc6, 0xa3, 0x16, 0x2e, 0x4a, 0xda, 0xe7, 0xec, 0xfb, 0xe5, 0xe9, 0x0e, 0xf2, 0xee, 0x19, 0xaf,
	0xd7, 0xc6, 0xea, 0xe5, 0xda, 0x97, 0xcb, 0xd3, 0x9d, 0x79, 0xe3, 0xdb, 0x75, 0xbc, 0xb1, 0xb4,
	0x3e, 0x8f, 0xab, 0x1e, 0x24, 0x8a, 0x3f, 0xfd, 0x84, 0xcb, 0x47, 0x4a, 0x58, 0x1f, 0xf1, 0xda,
	0xc2, 0x7e, 0x1f, 0xb9, 0xff, 0xfa, 0xd2, 0xdc, 0x25, 0xa7, 0x1a, 0xb9, 0xe5, 0xe0, 0x55, 0x64,
	0xbb, 0x73, 0x36, 0x71, 0xd0, 0xf9, 0xc4, 0x41, 0xbf, 0x27, 0x0e, 0xfa, 0x36, 0x75, 0x4a, 0xe7,
	0x53, 0xa7, 0xf4, 0x73, 0xea, 0x94, 0x3e, 0xbc, 0xb8, 0x7d, 0xe1, 0x83, 0xf9, 0xaf, 0x31, 0x6b,
	0xbe, 0x7b, 0xd7, 0x30, 0xcf, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xf7, 0x95, 0xbe, 0x42, 0x3f,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// SlashValidator is exposed to allow slashing of a misbehaving validator via
	// governance.
	SlashValidator(ctx context.Context, in *MsgSlashValidator, opts ...grpc.CallOption) (*MsgSlashValidatorResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SlashValidator(ctx context.Context, in *MsgSlashValidator, opts ...grpc.CallOption) (*MsgSlashValidatorResponse, error) {
	out := new(MsgSlashValidatorResponse)
	err := c.cc.Invoke(ctx, "/bitoroprotocol.govplus.Msg/SlashValidator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// SlashValidator is exposed to allow slashing of a misbehaving validator via
	// governance.
	SlashValidator(context.Context, *MsgSlashValidator) (*MsgSlashValidatorResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SlashValidator(ctx context.Context, req *MsgSlashValidator) (*MsgSlashValidatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SlashValidator not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SlashValidator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSlashValidator)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SlashValidator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bitoroprotocol.govplus.Msg/SlashValidator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SlashValidator(ctx, req.(*MsgSlashValidator))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bitoroprotocol.govplus.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SlashValidator",
			Handler:    _Msg_SlashValidator_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bitoroprotocol/govplus/tx.proto",
}

func (m *MsgSlashValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSlashValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSlashValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.SlashFactor.Size()
		i -= size
		if _, err := m.SlashFactor.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.TokensAtInfractionHeight.Size()
		i -= size
		if _, err := m.TokensAtInfractionHeight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.InfractionHeight != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.InfractionHeight))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSlashValidatorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSlashValidatorResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSlashValidatorResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSlashValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.InfractionHeight != 0 {
		n += 1 + sovTx(uint64(m.InfractionHeight))
	}
	l = m.TokensAtInfractionHeight.Size()
	n += 1 + l + sovTx(uint64(l))
	l = m.SlashFactor.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgSlashValidatorResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSlashValidator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSlashValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSlashValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InfractionHeight", wireType)
			}
			m.InfractionHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InfractionHeight |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokensAtInfractionHeight", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TokensAtInfractionHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashFactor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SlashFactor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSlashValidatorResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSlashValidatorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSlashValidatorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
