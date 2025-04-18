// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bitoroprotocol/blocktime/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

// MsgUpdateDowntimeParams is the Msg/UpdateDowntimeParams request type.
type MsgUpdateDowntimeParams struct {
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// Defines the parameters to update. All parameters must be supplied.
	Params DowntimeParams `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
}

func (m *MsgUpdateDowntimeParams) Reset()         { *m = MsgUpdateDowntimeParams{} }
func (m *MsgUpdateDowntimeParams) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateDowntimeParams) ProtoMessage()    {}
func (*MsgUpdateDowntimeParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_300c0e7978c8aea8, []int{0}
}
func (m *MsgUpdateDowntimeParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateDowntimeParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateDowntimeParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateDowntimeParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateDowntimeParams.Merge(m, src)
}
func (m *MsgUpdateDowntimeParams) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateDowntimeParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateDowntimeParams.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateDowntimeParams proto.InternalMessageInfo

func (m *MsgUpdateDowntimeParams) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgUpdateDowntimeParams) GetParams() DowntimeParams {
	if m != nil {
		return m.Params
	}
	return DowntimeParams{}
}

// MsgUpdateDowntimeParamsResponse is the Msg/UpdateDowntimeParams response
// type.
type MsgUpdateDowntimeParamsResponse struct {
}

func (m *MsgUpdateDowntimeParamsResponse) Reset()         { *m = MsgUpdateDowntimeParamsResponse{} }
func (m *MsgUpdateDowntimeParamsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateDowntimeParamsResponse) ProtoMessage()    {}
func (*MsgUpdateDowntimeParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_300c0e7978c8aea8, []int{1}
}
func (m *MsgUpdateDowntimeParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateDowntimeParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateDowntimeParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateDowntimeParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateDowntimeParamsResponse.Merge(m, src)
}
func (m *MsgUpdateDowntimeParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateDowntimeParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateDowntimeParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateDowntimeParamsResponse proto.InternalMessageInfo

// MsgUpdateSynchronyParams is the Msg/UpdateSynchronyParams request type.
type MsgUpdateSynchronyParams struct {
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// Defines the parameters to update. All parameters must be supplied.
	Params SynchronyParams `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
}

func (m *MsgUpdateSynchronyParams) Reset()         { *m = MsgUpdateSynchronyParams{} }
func (m *MsgUpdateSynchronyParams) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateSynchronyParams) ProtoMessage()    {}
func (*MsgUpdateSynchronyParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_300c0e7978c8aea8, []int{2}
}
func (m *MsgUpdateSynchronyParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateSynchronyParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateSynchronyParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateSynchronyParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateSynchronyParams.Merge(m, src)
}
func (m *MsgUpdateSynchronyParams) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateSynchronyParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateSynchronyParams.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateSynchronyParams proto.InternalMessageInfo

func (m *MsgUpdateSynchronyParams) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgUpdateSynchronyParams) GetParams() SynchronyParams {
	if m != nil {
		return m.Params
	}
	return SynchronyParams{}
}

// MsgUpdateSynchronyParamsResponse is the Msg/UpdateSynchronyParams response
// type.
type MsgUpdateSynchronyParamsResponse struct {
}

func (m *MsgUpdateSynchronyParamsResponse) Reset()         { *m = MsgUpdateSynchronyParamsResponse{} }
func (m *MsgUpdateSynchronyParamsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateSynchronyParamsResponse) ProtoMessage()    {}
func (*MsgUpdateSynchronyParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_300c0e7978c8aea8, []int{3}
}
func (m *MsgUpdateSynchronyParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateSynchronyParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateSynchronyParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateSynchronyParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateSynchronyParamsResponse.Merge(m, src)
}
func (m *MsgUpdateSynchronyParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateSynchronyParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateSynchronyParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateSynchronyParamsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgUpdateDowntimeParams)(nil), "bitoroprotocol.blocktime.MsgUpdateDowntimeParams")
	proto.RegisterType((*MsgUpdateDowntimeParamsResponse)(nil), "bitoroprotocol.blocktime.MsgUpdateDowntimeParamsResponse")
	proto.RegisterType((*MsgUpdateSynchronyParams)(nil), "bitoroprotocol.blocktime.MsgUpdateSynchronyParams")
	proto.RegisterType((*MsgUpdateSynchronyParamsResponse)(nil), "bitoroprotocol.blocktime.MsgUpdateSynchronyParamsResponse")
}

func init() { proto.RegisterFile("bitoroprotocol/blocktime/tx.proto", fileDescriptor_300c0e7978c8aea8) }

var fileDescriptor_300c0e7978c8aea8 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0xca, 0x2c, 0xc9,
	0x2f, 0xca, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x4f, 0xca, 0xc9, 0x4f, 0xce,
	0x2e, 0xc9, 0xcc, 0x4d, 0xd5, 0x2f, 0xa9, 0xd0, 0x03, 0x8b, 0x0a, 0x49, 0xa0, 0x2a, 0xd1, 0x83,
	0x2b, 0x91, 0x92, 0x4c, 0xce, 0x2f, 0xce, 0xcd, 0x2f, 0x8e, 0x07, 0x4b, 0xe9, 0x43, 0x38, 0x10,
	0x4d, 0x52, 0xe2, 0x10, 0x9e, 0x7e, 0x6e, 0x71, 0xba, 0x7e, 0x99, 0x21, 0x88, 0x82, 0x4a, 0xa8,
	0xe2, 0xb4, 0xb0, 0x20, 0xb1, 0x28, 0x31, 0x17, 0xa6, 0x5f, 0x24, 0x3d, 0x3f, 0x1d, 0xa2, 0x48,
	0x1f, 0xc4, 0x82, 0x88, 0x2a, 0xad, 0x64, 0xe4, 0x12, 0xf7, 0x2d, 0x4e, 0x0f, 0x2d, 0x48, 0x49,
	0x2c, 0x49, 0x75, 0xc9, 0x2f, 0xcf, 0x03, 0x69, 0x0c, 0x00, 0xeb, 0x13, 0x32, 0xe3, 0xe2, 0x4c,
	0x2c, 0x2d, 0xc9, 0xc8, 0x2f, 0xca, 0x2c, 0xa9, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x74, 0x92,
	0xb8, 0xb4, 0x45, 0x57, 0x04, 0xea, 0x2c, 0xc7, 0x94, 0x94, 0xa2, 0xd4, 0xe2, 0xe2, 0xe0, 0x92,
	0xa2, 0xcc, 0xbc, 0xf4, 0x20, 0x84, 0x52, 0x21, 0x37, 0x2e, 0x36, 0x88, 0xcd, 0x12, 0x4c, 0x0a,
	0x8c, 0x1a, 0xdc, 0x46, 0x1a, 0x7a, 0xb8, 0xfc, 0xab, 0x87, 0x6a, 0xa3, 0x13, 0xcb, 0x89, 0x7b,
	0xf2, 0x0c, 0x41, 0x50, 0xdd, 0x56, 0x7c, 0x4d, 0xcf, 0x37, 0x68, 0x21, 0xcc, 0x55, 0x52, 0xe4,
	0x92, 0xc7, 0xe1, 0xd4, 0xa0, 0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0xa5, 0xd5, 0x8c, 0x5c,
	0x12, 0x70, 0x35, 0xc1, 0x95, 0x79, 0xc9, 0x19, 0x45, 0xf9, 0x79, 0x95, 0x14, 0xfa, 0xc7, 0x1d,
	0xcd, 0x3f, 0x9a, 0xb8, 0xfd, 0x83, 0x66, 0x25, 0x01, 0x0f, 0x29, 0x71, 0x29, 0xe0, 0x72, 0x2c,
	0xcc, 0x47, 0x46, 0xd3, 0x98, 0xb8, 0x98, 0x7d, 0x8b, 0xd3, 0x85, 0x5a, 0x18, 0xb9, 0x44, 0xb0,
	0xc6, 0x92, 0x21, 0x6e, 0xd7, 0xe0, 0x08, 0x2d, 0x29, 0x4b, 0x92, 0xb5, 0xc0, 0x9c, 0x23, 0xd4,
	0xce, 0xc8, 0x25, 0x8a, 0x3d, 0x74, 0x8d, 0x88, 0x30, 0x14, 0x4d, 0x8f, 0x94, 0x15, 0xe9, 0x7a,
	0x60, 0x2e, 0x71, 0x0a, 0x39, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4,
	0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xab,
	0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x88, 0xf9, 0xba, 0x79, 0xa9,
	0x25, 0xe5, 0xf9, 0x45, 0xd9, 0xfa, 0xc9, 0x19, 0x89, 0x99, 0x79, 0xfa, 0xf0, 0xac, 0x52, 0x81,
	0x9c, 0x3b, 0x2b, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x72, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xa6, 0x00, 0xb8, 0xde, 0xc6, 0x03, 0x00, 0x00,
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
	// UpdateDowntimeParams updates the DowntimeParams in state.
	UpdateDowntimeParams(ctx context.Context, in *MsgUpdateDowntimeParams, opts ...grpc.CallOption) (*MsgUpdateDowntimeParamsResponse, error)
	// UpdateSynchronyParams updates the SynchronyParams in state.
	UpdateSynchronyParams(ctx context.Context, in *MsgUpdateSynchronyParams, opts ...grpc.CallOption) (*MsgUpdateSynchronyParamsResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateDowntimeParams(ctx context.Context, in *MsgUpdateDowntimeParams, opts ...grpc.CallOption) (*MsgUpdateDowntimeParamsResponse, error) {
	out := new(MsgUpdateDowntimeParamsResponse)
	err := c.cc.Invoke(ctx, "/bitoroprotocol.blocktime.Msg/UpdateDowntimeParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateSynchronyParams(ctx context.Context, in *MsgUpdateSynchronyParams, opts ...grpc.CallOption) (*MsgUpdateSynchronyParamsResponse, error) {
	out := new(MsgUpdateSynchronyParamsResponse)
	err := c.cc.Invoke(ctx, "/bitoroprotocol.blocktime.Msg/UpdateSynchronyParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// UpdateDowntimeParams updates the DowntimeParams in state.
	UpdateDowntimeParams(context.Context, *MsgUpdateDowntimeParams) (*MsgUpdateDowntimeParamsResponse, error)
	// UpdateSynchronyParams updates the SynchronyParams in state.
	UpdateSynchronyParams(context.Context, *MsgUpdateSynchronyParams) (*MsgUpdateSynchronyParamsResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) UpdateDowntimeParams(ctx context.Context, req *MsgUpdateDowntimeParams) (*MsgUpdateDowntimeParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDowntimeParams not implemented")
}
func (*UnimplementedMsgServer) UpdateSynchronyParams(ctx context.Context, req *MsgUpdateSynchronyParams) (*MsgUpdateSynchronyParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSynchronyParams not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_UpdateDowntimeParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateDowntimeParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateDowntimeParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bitoroprotocol.blocktime.Msg/UpdateDowntimeParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateDowntimeParams(ctx, req.(*MsgUpdateDowntimeParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateSynchronyParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateSynchronyParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateSynchronyParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bitoroprotocol.blocktime.Msg/UpdateSynchronyParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateSynchronyParams(ctx, req.(*MsgUpdateSynchronyParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bitoroprotocol.blocktime.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateDowntimeParams",
			Handler:    _Msg_UpdateDowntimeParams_Handler,
		},
		{
			MethodName: "UpdateSynchronyParams",
			Handler:    _Msg_UpdateSynchronyParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bitoroprotocol/blocktime/tx.proto",
}

func (m *MsgUpdateDowntimeParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateDowntimeParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateDowntimeParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateDowntimeParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateDowntimeParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateDowntimeParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpdateSynchronyParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateSynchronyParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateSynchronyParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateSynchronyParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateSynchronyParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateSynchronyParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgUpdateDowntimeParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Params.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgUpdateDowntimeParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpdateSynchronyParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Params.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgUpdateSynchronyParamsResponse) Size() (n int) {
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
func (m *MsgUpdateDowntimeParams) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateDowntimeParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateDowntimeParams: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgUpdateDowntimeParamsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateDowntimeParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateDowntimeParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgUpdateSynchronyParams) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateSynchronyParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateSynchronyParams: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgUpdateSynchronyParamsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdateSynchronyParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateSynchronyParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
