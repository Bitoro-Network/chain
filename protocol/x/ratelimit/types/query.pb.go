// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bitoroprotocol/ratelimit/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// ListLimitParamsRequest is a request type of the ListLimitParams RPC method.
type ListLimitParamsRequest struct {
}

func (m *ListLimitParamsRequest) Reset()         { *m = ListLimitParamsRequest{} }
func (m *ListLimitParamsRequest) String() string { return proto.CompactTextString(m) }
func (*ListLimitParamsRequest) ProtoMessage()    {}
func (*ListLimitParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_55408771ecf91c0e, []int{0}
}
func (m *ListLimitParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListLimitParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListLimitParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListLimitParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListLimitParamsRequest.Merge(m, src)
}
func (m *ListLimitParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *ListLimitParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListLimitParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListLimitParamsRequest proto.InternalMessageInfo

// ListLimitParamsResponse is a response type of the ListLimitParams RPC method.
type ListLimitParamsResponse struct {
	LimitParamsList []LimitParams `protobuf:"bytes,1,rep,name=limit_params_list,json=limitParamsList,proto3" json:"limit_params_list"`
}

func (m *ListLimitParamsResponse) Reset()         { *m = ListLimitParamsResponse{} }
func (m *ListLimitParamsResponse) String() string { return proto.CompactTextString(m) }
func (*ListLimitParamsResponse) ProtoMessage()    {}
func (*ListLimitParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_55408771ecf91c0e, []int{1}
}
func (m *ListLimitParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListLimitParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListLimitParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListLimitParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListLimitParamsResponse.Merge(m, src)
}
func (m *ListLimitParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *ListLimitParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListLimitParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListLimitParamsResponse proto.InternalMessageInfo

func (m *ListLimitParamsResponse) GetLimitParamsList() []LimitParams {
	if m != nil {
		return m.LimitParamsList
	}
	return nil
}

// QueryCapacityByDenomRequest is a request type for the CapacityByDenom RPC
// method.
type QueryCapacityByDenomRequest struct {
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (m *QueryCapacityByDenomRequest) Reset()         { *m = QueryCapacityByDenomRequest{} }
func (m *QueryCapacityByDenomRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCapacityByDenomRequest) ProtoMessage()    {}
func (*QueryCapacityByDenomRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_55408771ecf91c0e, []int{2}
}
func (m *QueryCapacityByDenomRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCapacityByDenomRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCapacityByDenomRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCapacityByDenomRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCapacityByDenomRequest.Merge(m, src)
}
func (m *QueryCapacityByDenomRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryCapacityByDenomRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCapacityByDenomRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCapacityByDenomRequest proto.InternalMessageInfo

func (m *QueryCapacityByDenomRequest) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

// QueryCapacityByDenomResponse is a response type of the CapacityByDenom RPC
// method.
type QueryCapacityByDenomResponse struct {
	LimiterCapacityList []LimiterCapacity `protobuf:"bytes,1,rep,name=limiter_capacity_list,json=limiterCapacityList,proto3" json:"limiter_capacity_list"`
}

func (m *QueryCapacityByDenomResponse) Reset()         { *m = QueryCapacityByDenomResponse{} }
func (m *QueryCapacityByDenomResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCapacityByDenomResponse) ProtoMessage()    {}
func (*QueryCapacityByDenomResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_55408771ecf91c0e, []int{3}
}
func (m *QueryCapacityByDenomResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCapacityByDenomResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCapacityByDenomResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCapacityByDenomResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCapacityByDenomResponse.Merge(m, src)
}
func (m *QueryCapacityByDenomResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCapacityByDenomResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCapacityByDenomResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCapacityByDenomResponse proto.InternalMessageInfo

func (m *QueryCapacityByDenomResponse) GetLimiterCapacityList() []LimiterCapacity {
	if m != nil {
		return m.LimiterCapacityList
	}
	return nil
}

// QueryAllPendingSendPacketsRequest is a request type for the
// AllPendingSendPackets RPC
type QueryAllPendingSendPacketsRequest struct {
}

func (m *QueryAllPendingSendPacketsRequest) Reset()         { *m = QueryAllPendingSendPacketsRequest{} }
func (m *QueryAllPendingSendPacketsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllPendingSendPacketsRequest) ProtoMessage()    {}
func (*QueryAllPendingSendPacketsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_55408771ecf91c0e, []int{4}
}
func (m *QueryAllPendingSendPacketsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllPendingSendPacketsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllPendingSendPacketsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllPendingSendPacketsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllPendingSendPacketsRequest.Merge(m, src)
}
func (m *QueryAllPendingSendPacketsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllPendingSendPacketsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllPendingSendPacketsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllPendingSendPacketsRequest proto.InternalMessageInfo

// QueryAllPendingSendPacketsResponse is a response type of the
// AllPendingSendPackets RPC
type QueryAllPendingSendPacketsResponse struct {
	PendingSendPackets []PendingSendPacket `protobuf:"bytes,1,rep,name=pending_send_packets,json=pendingSendPackets,proto3" json:"pending_send_packets"`
}

func (m *QueryAllPendingSendPacketsResponse) Reset()         { *m = QueryAllPendingSendPacketsResponse{} }
func (m *QueryAllPendingSendPacketsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllPendingSendPacketsResponse) ProtoMessage()    {}
func (*QueryAllPendingSendPacketsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_55408771ecf91c0e, []int{5}
}
func (m *QueryAllPendingSendPacketsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllPendingSendPacketsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllPendingSendPacketsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllPendingSendPacketsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllPendingSendPacketsResponse.Merge(m, src)
}
func (m *QueryAllPendingSendPacketsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllPendingSendPacketsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllPendingSendPacketsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllPendingSendPacketsResponse proto.InternalMessageInfo

func (m *QueryAllPendingSendPacketsResponse) GetPendingSendPackets() []PendingSendPacket {
	if m != nil {
		return m.PendingSendPackets
	}
	return nil
}

func init() {
	proto.RegisterType((*ListLimitParamsRequest)(nil), "bitoroprotocol.ratelimit.ListLimitParamsRequest")
	proto.RegisterType((*ListLimitParamsResponse)(nil), "bitoroprotocol.ratelimit.ListLimitParamsResponse")
	proto.RegisterType((*QueryCapacityByDenomRequest)(nil), "bitoroprotocol.ratelimit.QueryCapacityByDenomRequest")
	proto.RegisterType((*QueryCapacityByDenomResponse)(nil), "bitoroprotocol.ratelimit.QueryCapacityByDenomResponse")
	proto.RegisterType((*QueryAllPendingSendPacketsRequest)(nil), "bitoroprotocol.ratelimit.QueryAllPendingSendPacketsRequest")
	proto.RegisterType((*QueryAllPendingSendPacketsResponse)(nil), "bitoroprotocol.ratelimit.QueryAllPendingSendPacketsResponse")
}

func init() {
	proto.RegisterFile("bitoroprotocol/ratelimit/query.proto", fileDescriptor_55408771ecf91c0e)
}

var fileDescriptor_55408771ecf91c0e = []byte{
	// 528 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x4d, 0x6f, 0xd3, 0x30,
	0x18, 0xc7, 0x6b, 0xd8, 0x90, 0x30, 0x87, 0x0a, 0xd3, 0x41, 0x15, 0xa6, 0x30, 0x02, 0x88, 0xa1,
	0x6a, 0x0d, 0xb4, 0x1a, 0x48, 0xbc, 0x1c, 0x28, 0x1c, 0x77, 0x28, 0x05, 0x09, 0x89, 0x4b, 0xe4,
	0xa6, 0x56, 0x66, 0xcd, 0xb5, 0x33, 0xdb, 0x15, 0xf4, 0x0a, 0x27, 0x6e, 0x48, 0x7c, 0x0c, 0xce,
	0x7c, 0x00, 0x6e, 0x3b, 0x0e, 0x71, 0xe1, 0x84, 0x50, 0xcb, 0x07, 0x41, 0x71, 0xbc, 0x90, 0xad,
	0x49, 0xc6, 0x76, 0xa9, 0x6a, 0xfb, 0xff, 0x3c, 0xff, 0xdf, 0xf3, 0x12, 0x78, 0x73, 0x48, 0xb5,
	0x90, 0x22, 0x96, 0x42, 0x8b, 0x50, 0x30, 0x5f, 0x62, 0x4d, 0x18, 0x1d, 0x53, 0xed, 0xef, 0x4e,
	0x88, 0x9c, 0xb6, 0xcd, 0x03, 0x6a, 0x1e, 0x56, 0xb5, 0x33, 0x95, 0xd3, 0x88, 0x44, 0x94, 0xde,
	0xfb, 0xc9, 0xbf, 0x54, 0xef, 0xac, 0x46, 0x42, 0x44, 0x8c, 0xf8, 0x38, 0xa6, 0x3e, 0xe6, 0x5c,
	0x68, 0xac, 0xa9, 0xe0, 0xca, 0xbe, 0xb6, 0x4a, 0x3d, 0xcd, 0x6f, 0x10, 0x63, 0x89, 0xc7, 0x07,
	0xe2, 0xdb, 0xa5, 0xe2, 0x10, 0xc7, 0x38, 0xa4, 0xda, 0x32, 0x3a, 0x9d, 0x52, 0x61, 0x4c, 0xf8,
	0x88, 0xf2, 0x28, 0x50, 0x84, 0x8f, 0x82, 0x18, 0x87, 0x3b, 0x44, 0xa7, 0x31, 0x5e, 0x13, 0x5e,
	0xde, 0xa2, 0x4a, 0x6f, 0x25, 0xb2, 0xbe, 0x71, 0x1d, 0x90, 0xdd, 0x09, 0x51, 0xda, 0x93, 0xf0,
	0xca, 0xc2, 0x8b, 0x8a, 0x05, 0x57, 0x04, 0xbd, 0x86, 0x17, 0xf3, 0x9c, 0x01, 0xa3, 0x4a, 0x37,
	0xc1, 0xda, 0xd9, 0xf5, 0x0b, 0x9d, 0x5b, 0xed, 0xb2, 0x46, 0xb5, 0x73, 0x99, 0x7a, 0x4b, 0x7b,
	0xbf, 0xae, 0xd5, 0x06, 0x75, 0xf6, 0xef, 0x2a, 0xf1, 0xf2, 0xba, 0xf0, 0xea, 0x8b, 0xa4, 0xe9,
	0xcf, 0x6c, 0x61, 0xbd, 0xe9, 0x73, 0xc2, 0xc5, 0xd8, 0x22, 0xa1, 0x06, 0x5c, 0x1e, 0x25, 0xe7,
	0x26, 0x58, 0x03, 0xeb, 0xe7, 0x07, 0xe9, 0xc1, 0xfb, 0x00, 0xe0, 0x6a, 0x71, 0x94, 0xc5, 0x0d,
	0xe1, 0x8a, 0x31, 0x22, 0x32, 0x38, 0xe8, 0x58, 0x1e, 0xf9, 0xce, 0x31, 0xc8, 0x44, 0x66, 0x89,
	0x53, 0xec, 0x4b, 0xec, 0xf0, 0xb5, 0x41, 0xbf, 0x01, 0xaf, 0x1b, 0x88, 0xa7, 0x8c, 0xf5, 0xd3,
	0x6e, 0xbf, 0x24, 0x7c, 0xd4, 0x37, 0xbd, 0xce, 0x7a, 0xfa, 0x11, 0x40, 0xaf, 0x4a, 0x95, 0x01,
	0x37, 0x0a, 0x26, 0xa6, 0x2c, 0x6f, 0xab, 0x9c, 0x77, 0x21, 0xa7, 0x25, 0x46, 0xf1, 0x82, 0x59,
	0xe7, 0xdb, 0x12, 0x5c, 0x36, 0x2c, 0xe8, 0x0b, 0x80, 0xf5, 0x23, 0xa3, 0x46, 0x77, 0xab, 0x9a,
	0x52, 0xb4, 0x2f, 0xce, 0xbd, 0x13, 0x44, 0xa4, 0x75, 0x7a, 0xdd, 0xf7, 0x3f, 0xfe, 0x7c, 0x3e,
	0xb3, 0x81, 0x5a, 0x7e, 0xc5, 0xf7, 0xa0, 0x74, 0x90, 0x5f, 0x36, 0xf4, 0x15, 0xc0, 0xfa, 0x91,
	0x49, 0xa3, 0xcd, 0x72, 0xef, 0x8a, 0x7d, 0x72, 0xee, 0x9f, 0x34, 0xec, 0xff, 0xb9, 0xb3, 0x45,
	0x1b, 0x4e, 0x03, 0xb3, 0xa6, 0xe8, 0x3b, 0x80, 0x2b, 0x85, 0x63, 0x47, 0x8f, 0x8e, 0xc1, 0xa8,
	0x5a, 0x29, 0xe7, 0xf1, 0xe9, 0x82, 0x6d, 0x25, 0x4f, 0x4c, 0x25, 0x0f, 0xd0, 0x66, 0x79, 0x25,
	0x11, 0xd1, 0x01, 0x66, 0x2c, 0x28, 0xd8, 0xc8, 0xde, 0xab, 0xbd, 0x99, 0x0b, 0xf6, 0x67, 0x2e,
	0xf8, 0x3d, 0x73, 0xc1, 0xa7, 0xb9, 0x5b, 0xdb, 0x9f, 0xbb, 0xb5, 0x9f, 0x73, 0xb7, 0xf6, 0xe6,
	0x61, 0x44, 0xf5, 0xf6, 0x64, 0xd8, 0x0e, 0xc5, 0xd8, 0xa6, 0xde, 0xe0, 0x44, 0xbf, 0x15, 0x72,
	0xc7, 0x0f, 0xb7, 0x31, 0xe5, 0x7e, 0xe6, 0xf4, 0x2e, 0xe7, 0xa5, 0xa7, 0x31, 0x51, 0xc3, 0x73,
	0xe6, 0xad, 0xfb, 0x37, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x89, 0xdc, 0x5f, 0x9a, 0x05, 0x00, 0x00,
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
	// List all limit params.
	ListLimitParams(ctx context.Context, in *ListLimitParamsRequest, opts ...grpc.CallOption) (*ListLimitParamsResponse, error)
	// Query capacity by denom.
	CapacityByDenom(ctx context.Context, in *QueryCapacityByDenomRequest, opts ...grpc.CallOption) (*QueryCapacityByDenomResponse, error)
	// Get all pending send packets
	AllPendingSendPackets(ctx context.Context, in *QueryAllPendingSendPacketsRequest, opts ...grpc.CallOption) (*QueryAllPendingSendPacketsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) ListLimitParams(ctx context.Context, in *ListLimitParamsRequest, opts ...grpc.CallOption) (*ListLimitParamsResponse, error) {
	out := new(ListLimitParamsResponse)
	err := c.cc.Invoke(ctx, "/bitoroprotocol.ratelimit.Query/ListLimitParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CapacityByDenom(ctx context.Context, in *QueryCapacityByDenomRequest, opts ...grpc.CallOption) (*QueryCapacityByDenomResponse, error) {
	out := new(QueryCapacityByDenomResponse)
	err := c.cc.Invoke(ctx, "/bitoroprotocol.ratelimit.Query/CapacityByDenom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AllPendingSendPackets(ctx context.Context, in *QueryAllPendingSendPacketsRequest, opts ...grpc.CallOption) (*QueryAllPendingSendPacketsResponse, error) {
	out := new(QueryAllPendingSendPacketsResponse)
	err := c.cc.Invoke(ctx, "/bitoroprotocol.ratelimit.Query/AllPendingSendPackets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// List all limit params.
	ListLimitParams(context.Context, *ListLimitParamsRequest) (*ListLimitParamsResponse, error)
	// Query capacity by denom.
	CapacityByDenom(context.Context, *QueryCapacityByDenomRequest) (*QueryCapacityByDenomResponse, error)
	// Get all pending send packets
	AllPendingSendPackets(context.Context, *QueryAllPendingSendPacketsRequest) (*QueryAllPendingSendPacketsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) ListLimitParams(ctx context.Context, req *ListLimitParamsRequest) (*ListLimitParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLimitParams not implemented")
}
func (*UnimplementedQueryServer) CapacityByDenom(ctx context.Context, req *QueryCapacityByDenomRequest) (*QueryCapacityByDenomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CapacityByDenom not implemented")
}
func (*UnimplementedQueryServer) AllPendingSendPackets(ctx context.Context, req *QueryAllPendingSendPacketsRequest) (*QueryAllPendingSendPacketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllPendingSendPackets not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_ListLimitParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLimitParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ListLimitParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bitoroprotocol.ratelimit.Query/ListLimitParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ListLimitParams(ctx, req.(*ListLimitParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CapacityByDenom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCapacityByDenomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CapacityByDenom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bitoroprotocol.ratelimit.Query/CapacityByDenom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CapacityByDenom(ctx, req.(*QueryCapacityByDenomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AllPendingSendPackets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllPendingSendPacketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AllPendingSendPackets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bitoroprotocol.ratelimit.Query/AllPendingSendPackets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AllPendingSendPackets(ctx, req.(*QueryAllPendingSendPacketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bitoroprotocol.ratelimit.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListLimitParams",
			Handler:    _Query_ListLimitParams_Handler,
		},
		{
			MethodName: "CapacityByDenom",
			Handler:    _Query_CapacityByDenom_Handler,
		},
		{
			MethodName: "AllPendingSendPackets",
			Handler:    _Query_AllPendingSendPackets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bitoroprotocol/ratelimit/query.proto",
}

func (m *ListLimitParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListLimitParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListLimitParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ListLimitParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListLimitParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListLimitParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LimitParamsList) > 0 {
		for iNdEx := len(m.LimitParamsList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LimitParamsList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryCapacityByDenomRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCapacityByDenomRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCapacityByDenomRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryCapacityByDenomResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCapacityByDenomResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCapacityByDenomResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LimiterCapacityList) > 0 {
		for iNdEx := len(m.LimiterCapacityList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LimiterCapacityList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllPendingSendPacketsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllPendingSendPacketsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllPendingSendPacketsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryAllPendingSendPacketsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllPendingSendPacketsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllPendingSendPacketsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PendingSendPackets) > 0 {
		for iNdEx := len(m.PendingSendPackets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PendingSendPackets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
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
func (m *ListLimitParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ListLimitParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.LimitParamsList) > 0 {
		for _, e := range m.LimitParamsList {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryCapacityByDenomRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryCapacityByDenomResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.LimiterCapacityList) > 0 {
		for _, e := range m.LimiterCapacityList {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryAllPendingSendPacketsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryAllPendingSendPacketsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PendingSendPackets) > 0 {
		for _, e := range m.PendingSendPackets {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ListLimitParamsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ListLimitParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListLimitParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *ListLimitParamsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ListLimitParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListLimitParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LimitParamsList", wireType)
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
			m.LimitParamsList = append(m.LimitParamsList, LimitParams{})
			if err := m.LimitParamsList[len(m.LimitParamsList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryCapacityByDenomRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryCapacityByDenomRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCapacityByDenomRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
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
func (m *QueryCapacityByDenomResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryCapacityByDenomResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCapacityByDenomResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LimiterCapacityList", wireType)
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
			m.LimiterCapacityList = append(m.LimiterCapacityList, LimiterCapacity{})
			if err := m.LimiterCapacityList[len(m.LimiterCapacityList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryAllPendingSendPacketsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllPendingSendPacketsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllPendingSendPacketsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryAllPendingSendPacketsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllPendingSendPacketsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllPendingSendPacketsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PendingSendPackets", wireType)
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
			m.PendingSendPackets = append(m.PendingSendPackets, PendingSendPacket{})
			if err := m.PendingSendPackets[len(m.PendingSendPackets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
