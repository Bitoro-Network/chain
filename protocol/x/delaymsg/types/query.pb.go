// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bitoroprotocol/delaymsg/query.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// QueryNextDelayedMessageIdRequest is the request type for the
// NextDelayedMessageId RPC method.
type QueryNextDelayedMessageIdRequest struct {
}

func (m *QueryNextDelayedMessageIdRequest) Reset()         { *m = QueryNextDelayedMessageIdRequest{} }
func (m *QueryNextDelayedMessageIdRequest) String() string { return proto.CompactTextString(m) }
func (*QueryNextDelayedMessageIdRequest) ProtoMessage()    {}
func (*QueryNextDelayedMessageIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_60f81d4d7b29defa, []int{0}
}
func (m *QueryNextDelayedMessageIdRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryNextDelayedMessageIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryNextDelayedMessageIdRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryNextDelayedMessageIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryNextDelayedMessageIdRequest.Merge(m, src)
}
func (m *QueryNextDelayedMessageIdRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryNextDelayedMessageIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryNextDelayedMessageIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryNextDelayedMessageIdRequest proto.InternalMessageInfo

// QueryNextDelayedMessageIdResponse is the response type for the
// NextDelayedMessageId RPC method.
type QueryNextDelayedMessageIdResponse struct {
	NextDelayedMessageId uint32 `protobuf:"varint,1,opt,name=next_delayed_message_id,json=nextDelayedMessageId,proto3" json:"next_delayed_message_id,omitempty"`
}

func (m *QueryNextDelayedMessageIdResponse) Reset()         { *m = QueryNextDelayedMessageIdResponse{} }
func (m *QueryNextDelayedMessageIdResponse) String() string { return proto.CompactTextString(m) }
func (*QueryNextDelayedMessageIdResponse) ProtoMessage()    {}
func (*QueryNextDelayedMessageIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_60f81d4d7b29defa, []int{1}
}
func (m *QueryNextDelayedMessageIdResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryNextDelayedMessageIdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryNextDelayedMessageIdResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryNextDelayedMessageIdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryNextDelayedMessageIdResponse.Merge(m, src)
}
func (m *QueryNextDelayedMessageIdResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryNextDelayedMessageIdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryNextDelayedMessageIdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryNextDelayedMessageIdResponse proto.InternalMessageInfo

func (m *QueryNextDelayedMessageIdResponse) GetNextDelayedMessageId() uint32 {
	if m != nil {
		return m.NextDelayedMessageId
	}
	return 0
}

// QueryMessageRequest is the request type for the Message RPC method.
type QueryMessageRequest struct {
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryMessageRequest) Reset()         { *m = QueryMessageRequest{} }
func (m *QueryMessageRequest) String() string { return proto.CompactTextString(m) }
func (*QueryMessageRequest) ProtoMessage()    {}
func (*QueryMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_60f81d4d7b29defa, []int{2}
}
func (m *QueryMessageRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryMessageRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryMessageRequest.Merge(m, src)
}
func (m *QueryMessageRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryMessageRequest proto.InternalMessageInfo

func (m *QueryMessageRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

// QueryGetMessageResponse is the response type for the Message RPC method.
type QueryMessageResponse struct {
	Message *DelayedMessage `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *QueryMessageResponse) Reset()         { *m = QueryMessageResponse{} }
func (m *QueryMessageResponse) String() string { return proto.CompactTextString(m) }
func (*QueryMessageResponse) ProtoMessage()    {}
func (*QueryMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_60f81d4d7b29defa, []int{3}
}
func (m *QueryMessageResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryMessageResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryMessageResponse.Merge(m, src)
}
func (m *QueryMessageResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryMessageResponse proto.InternalMessageInfo

func (m *QueryMessageResponse) GetMessage() *DelayedMessage {
	if m != nil {
		return m.Message
	}
	return nil
}

// QueryBlockMessageIdsRequest is the request type for the BlockMessageIds
// RPC method.
type QueryBlockMessageIdsRequest struct {
	BlockHeight uint32 `protobuf:"varint,1,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
}

func (m *QueryBlockMessageIdsRequest) Reset()         { *m = QueryBlockMessageIdsRequest{} }
func (m *QueryBlockMessageIdsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryBlockMessageIdsRequest) ProtoMessage()    {}
func (*QueryBlockMessageIdsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_60f81d4d7b29defa, []int{4}
}
func (m *QueryBlockMessageIdsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBlockMessageIdsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBlockMessageIdsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBlockMessageIdsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBlockMessageIdsRequest.Merge(m, src)
}
func (m *QueryBlockMessageIdsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryBlockMessageIdsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBlockMessageIdsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBlockMessageIdsRequest proto.InternalMessageInfo

func (m *QueryBlockMessageIdsRequest) GetBlockHeight() uint32 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

// QueryGetBlockMessageIdsResponse is the response type for the BlockMessageIds
// RPC method.
type QueryBlockMessageIdsResponse struct {
	MessageIds []uint32 `protobuf:"varint,1,rep,packed,name=message_ids,json=messageIds,proto3" json:"message_ids,omitempty"`
}

func (m *QueryBlockMessageIdsResponse) Reset()         { *m = QueryBlockMessageIdsResponse{} }
func (m *QueryBlockMessageIdsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryBlockMessageIdsResponse) ProtoMessage()    {}
func (*QueryBlockMessageIdsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_60f81d4d7b29defa, []int{5}
}
func (m *QueryBlockMessageIdsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBlockMessageIdsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBlockMessageIdsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBlockMessageIdsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBlockMessageIdsResponse.Merge(m, src)
}
func (m *QueryBlockMessageIdsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryBlockMessageIdsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBlockMessageIdsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBlockMessageIdsResponse proto.InternalMessageInfo

func (m *QueryBlockMessageIdsResponse) GetMessageIds() []uint32 {
	if m != nil {
		return m.MessageIds
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryNextDelayedMessageIdRequest)(nil), "bitoroprotocol.delaymsg.QueryNextDelayedMessageIdRequest")
	proto.RegisterType((*QueryNextDelayedMessageIdResponse)(nil), "bitoroprotocol.delaymsg.QueryNextDelayedMessageIdResponse")
	proto.RegisterType((*QueryMessageRequest)(nil), "bitoroprotocol.delaymsg.QueryMessageRequest")
	proto.RegisterType((*QueryMessageResponse)(nil), "bitoroprotocol.delaymsg.QueryMessageResponse")
	proto.RegisterType((*QueryBlockMessageIdsRequest)(nil), "bitoroprotocol.delaymsg.QueryBlockMessageIdsRequest")
	proto.RegisterType((*QueryBlockMessageIdsResponse)(nil), "bitoroprotocol.delaymsg.QueryBlockMessageIdsResponse")
}

func init() { proto.RegisterFile("bitoroprotocol/delaymsg/query.proto", fileDescriptor_60f81d4d7b29defa) }

var fileDescriptor_60f81d4d7b29defa = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4f, 0x6b, 0x13, 0x41,
	0x14, 0xcf, 0xc4, 0x3f, 0x85, 0x17, 0xab, 0x30, 0x46, 0x2c, 0x6b, 0x59, 0x93, 0x91, 0x4a, 0xb4,
	0xb8, 0x03, 0xa9, 0xa2, 0x14, 0xa1, 0x12, 0x3c, 0xe8, 0x41, 0xc1, 0x20, 0x08, 0xbd, 0x84, 0x4d,
	0x66, 0xd8, 0x0c, 0x26, 0x3b, 0xdb, 0xce, 0x44, 0x76, 0x29, 0xbd, 0xf8, 0x09, 0x44, 0xfd, 0x28,
	0x9e, 0x3d, 0x7b, 0x2c, 0x78, 0xf1, 0x28, 0x89, 0x1f, 0x44, 0x3a, 0x3b, 0xd9, 0x64, 0xc3, 0x26,
	0x9a, 0xdb, 0xf2, 0xe6, 0xf7, 0x6f, 0xdf, 0xfb, 0x41, 0x9d, 0x25, 0x2c, 0x8e, 0x8e, 0xa5, 0x96,
	0x3d, 0x39, 0xa0, 0x8c, 0x0f, 0xfc, 0x64, 0xa8, 0x02, 0x7a, 0x34, 0xe2, 0xc7, 0x89, 0x67, 0xe6,
	0xf8, 0xc6, 0x3c, 0xc4, 0x9b, 0x42, 0x9c, 0xed, 0x40, 0xca, 0x60, 0xc0, 0xa9, 0x1f, 0x09, 0xea,
	0x87, 0xa1, 0xd4, 0xbe, 0x16, 0x32, 0x54, 0x29, 0xc9, 0xd9, 0x2d, 0xd6, 0x35, 0x1f, 0x9c, 0x75,
	0x86, 0x5c, 0x29, 0x3f, 0xe0, 0x29, 0x98, 0x10, 0xa8, 0xbd, 0x39, 0x37, 0x7c, 0xcd, 0x63, 0xfd,
	0x3c, 0x45, 0xbc, 0x4a, 0x01, 0x2f, 0x59, 0x9b, 0x1f, 0x8d, 0xb8, 0xd2, 0xe4, 0x10, 0xea, 0x2b,
	0x30, 0x2a, 0x92, 0xa1, 0xe2, 0xf8, 0x11, 0xdc, 0x0c, 0x79, 0xac, 0x3b, 0x0b, 0x36, 0x1d, 0xc1,
	0xb6, 0x50, 0x0d, 0x35, 0x36, 0xdb, 0xd5, 0xb0, 0x80, 0x4e, 0x76, 0xe0, 0xba, 0xd1, 0xb6, 0x13,
	0x6b, 0x89, 0xaf, 0x42, 0x39, 0x23, 0x96, 0x05, 0x23, 0xef, 0xa0, 0x9a, 0x87, 0x59, 0xd7, 0x03,
	0xd8, 0xb0, 0x46, 0x06, 0x5c, 0x69, 0xee, 0x78, 0x85, 0x2b, 0xf3, 0xf2, 0xc6, 0xed, 0x29, 0x8b,
	0x3c, 0x83, 0x5b, 0x46, 0xb8, 0x35, 0x90, 0xbd, 0xf7, 0x59, 0x2c, 0x35, 0xcd, 0x51, 0x87, 0x2b,
	0xdd, 0xf3, 0x97, 0x4e, 0x9f, 0x8b, 0xa0, 0xaf, 0x6d, 0xa2, 0x8a, 0x99, 0xbd, 0x30, 0x23, 0x72,
	0x00, 0xdb, 0xc5, 0x0a, 0x36, 0xe2, 0x6d, 0xa8, 0xcc, 0x76, 0xa1, 0xb6, 0x50, 0xed, 0x42, 0x63,
	0xb3, 0x0d, 0xc3, 0x0c, 0xd8, 0xfc, 0x7a, 0x11, 0x2e, 0x19, 0x05, 0xfc, 0x0d, 0x41, 0xb5, 0x68,
	0xc9, 0xf8, 0xf1, 0x92, 0xbf, 0xfa, 0xd7, 0xe9, 0x9c, 0x27, 0xeb, 0x13, 0xd3, 0xd8, 0xe4, 0xde,
	0xc7, 0x9f, 0x7f, 0xbe, 0x94, 0xef, 0xe0, 0x3a, 0xcd, 0xd5, 0xe9, 0xc3, 0xc3, 0x59, 0xa3, 0xcc,
	0xbd, 0x05, 0xc3, 0x9f, 0x11, 0x6c, 0x58, 0x01, 0x7c, 0x7f, 0x95, 0x61, 0xfe, 0xc8, 0xce, 0xee,
	0x7f, 0x61, 0x6d, 0x1e, 0xcf, 0xe4, 0x69, 0xe0, 0xbb, 0xcb, 0xf3, 0xd8, 0x9d, 0xd2, 0x13, 0xc1,
	0x4e, 0xf1, 0x77, 0x04, 0xd7, 0x16, 0x4e, 0x82, 0x9b, 0xab, 0x0c, 0x8b, 0x1b, 0xe0, 0xec, 0xad,
	0xc5, 0xb1, 0x61, 0x5b, 0x26, 0xec, 0x53, 0xbc, 0xbf, 0x3c, 0xac, 0xa9, 0x10, 0x9d, 0x6b, 0x06,
	0x3d, 0x99, 0x6f, 0xda, 0x69, 0xeb, 0xed, 0x8f, 0xb1, 0x8b, 0xce, 0xc6, 0x2e, 0xfa, 0x3d, 0x76,
	0xd1, 0xa7, 0x89, 0x5b, 0x3a, 0x9b, 0xb8, 0xa5, 0x5f, 0x13, 0xb7, 0x74, 0xb8, 0x1f, 0x08, 0xdd,
	0x1f, 0x75, 0xbd, 0x9e, 0x1c, 0x2e, 0xea, 0x3f, 0xe8, 0xf5, 0x7d, 0x11, 0xd2, 0x6c, 0x12, 0xcf,
	0x0c, 0x75, 0x12, 0x71, 0xd5, 0xbd, 0x6c, 0x9e, 0xf6, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x74,
	0x37, 0xc1, 0x6d, 0x7d, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Queries the next DelayedMessage's id.
	NextDelayedMessageId(ctx context.Context, in *QueryNextDelayedMessageIdRequest, opts ...grpc.CallOption) (*QueryNextDelayedMessageIdResponse, error)
	// Queries the DelayedMessage by id.
	Message(ctx context.Context, in *QueryMessageRequest, opts ...grpc.CallOption) (*QueryMessageResponse, error)
	// Queries the DelayedMessages at a given block height.
	BlockMessageIds(ctx context.Context, in *QueryBlockMessageIdsRequest, opts ...grpc.CallOption) (*QueryBlockMessageIdsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) NextDelayedMessageId(ctx context.Context, in *QueryNextDelayedMessageIdRequest, opts ...grpc.CallOption) (*QueryNextDelayedMessageIdResponse, error) {
	out := new(QueryNextDelayedMessageIdResponse)
	err := c.cc.Invoke(ctx, "/bitoroprotocol.delaymsg.Query/NextDelayedMessageId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Message(ctx context.Context, in *QueryMessageRequest, opts ...grpc.CallOption) (*QueryMessageResponse, error) {
	out := new(QueryMessageResponse)
	err := c.cc.Invoke(ctx, "/bitoroprotocol.delaymsg.Query/Message", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) BlockMessageIds(ctx context.Context, in *QueryBlockMessageIdsRequest, opts ...grpc.CallOption) (*QueryBlockMessageIdsResponse, error) {
	out := new(QueryBlockMessageIdsResponse)
	err := c.cc.Invoke(ctx, "/bitoroprotocol.delaymsg.Query/BlockMessageIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries the next DelayedMessage's id.
	NextDelayedMessageId(context.Context, *QueryNextDelayedMessageIdRequest) (*QueryNextDelayedMessageIdResponse, error)
	// Queries the DelayedMessage by id.
	Message(context.Context, *QueryMessageRequest) (*QueryMessageResponse, error)
	// Queries the DelayedMessages at a given block height.
	BlockMessageIds(context.Context, *QueryBlockMessageIdsRequest) (*QueryBlockMessageIdsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) NextDelayedMessageId(ctx context.Context, req *QueryNextDelayedMessageIdRequest) (*QueryNextDelayedMessageIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NextDelayedMessageId not implemented")
}
func (*UnimplementedQueryServer) Message(ctx context.Context, req *QueryMessageRequest) (*QueryMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Message not implemented")
}
func (*UnimplementedQueryServer) BlockMessageIds(ctx context.Context, req *QueryBlockMessageIdsRequest) (*QueryBlockMessageIdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockMessageIds not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_NextDelayedMessageId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryNextDelayedMessageIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).NextDelayedMessageId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bitoroprotocol.delaymsg.Query/NextDelayedMessageId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).NextDelayedMessageId(ctx, req.(*QueryNextDelayedMessageIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Message_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Message(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bitoroprotocol.delaymsg.Query/Message",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Message(ctx, req.(*QueryMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_BlockMessageIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBlockMessageIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).BlockMessageIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bitoroprotocol.delaymsg.Query/BlockMessageIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).BlockMessageIds(ctx, req.(*QueryBlockMessageIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bitoroprotocol.delaymsg.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NextDelayedMessageId",
			Handler:    _Query_NextDelayedMessageId_Handler,
		},
		{
			MethodName: "Message",
			Handler:    _Query_Message_Handler,
		},
		{
			MethodName: "BlockMessageIds",
			Handler:    _Query_BlockMessageIds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bitoroprotocol/delaymsg/query.proto",
}

func (m *QueryNextDelayedMessageIdRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryNextDelayedMessageIdRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryNextDelayedMessageIdRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryNextDelayedMessageIdResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryNextDelayedMessageIdResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryNextDelayedMessageIdResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NextDelayedMessageId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.NextDelayedMessageId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryMessageRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryMessageRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryMessageRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryMessageResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryMessageResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryMessageResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Message != nil {
		{
			size, err := m.Message.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryBlockMessageIdsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBlockMessageIdsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBlockMessageIdsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlockHeight != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryBlockMessageIdsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBlockMessageIdsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBlockMessageIdsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MessageIds) > 0 {
		dAtA3 := make([]byte, len(m.MessageIds)*10)
		var j2 int
		for _, num := range m.MessageIds {
			for num >= 1<<7 {
				dAtA3[j2] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j2++
			}
			dAtA3[j2] = uint8(num)
			j2++
		}
		i -= j2
		copy(dAtA[i:], dAtA3[:j2])
		i = encodeVarintQuery(dAtA, i, uint64(j2))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryNextDelayedMessageIdRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryNextDelayedMessageIdResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NextDelayedMessageId != 0 {
		n += 1 + sovQuery(uint64(m.NextDelayedMessageId))
	}
	return n
}

func (m *QueryMessageRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovQuery(uint64(m.Id))
	}
	return n
}

func (m *QueryMessageResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Message != nil {
		l = m.Message.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryBlockMessageIdsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockHeight != 0 {
		n += 1 + sovQuery(uint64(m.BlockHeight))
	}
	return n
}

func (m *QueryBlockMessageIdsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.MessageIds) > 0 {
		l = 0
		for _, e := range m.MessageIds {
			l += sovQuery(uint64(e))
		}
		n += 1 + sovQuery(uint64(l)) + l
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryNextDelayedMessageIdRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryNextDelayedMessageIdRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryNextDelayedMessageIdRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryNextDelayedMessageIdResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryNextDelayedMessageIdResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryNextDelayedMessageIdResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextDelayedMessageId", wireType)
			}
			m.NextDelayedMessageId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextDelayedMessageId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryMessageRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryMessageRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryMessageRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryMessageResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryMessageResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryMessageResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Message == nil {
				m.Message = &DelayedMessage{}
			}
			if err := m.Message.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryBlockMessageIdsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryBlockMessageIdsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBlockMessageIdsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryBlockMessageIdsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryBlockMessageIdsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBlockMessageIdsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowQuery
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.MessageIds = append(m.MessageIds, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowQuery
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthQuery
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthQuery
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.MessageIds) == 0 {
					m.MessageIds = make([]uint32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowQuery
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.MessageIds = append(m.MessageIds, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageIds", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
