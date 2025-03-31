// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bitoroprotocol/revshare/query.proto

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

// Queries for the default market mapper revenue share params
type QueryMarketMapperRevenueShareParams struct {
}

func (m *QueryMarketMapperRevenueShareParams) Reset()         { *m = QueryMarketMapperRevenueShareParams{} }
func (m *QueryMarketMapperRevenueShareParams) String() string { return proto.CompactTextString(m) }
func (*QueryMarketMapperRevenueShareParams) ProtoMessage()    {}
func (*QueryMarketMapperRevenueShareParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_13d50c6e3048e744, []int{0}
}
func (m *QueryMarketMapperRevenueShareParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryMarketMapperRevenueShareParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryMarketMapperRevenueShareParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryMarketMapperRevenueShareParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryMarketMapperRevenueShareParams.Merge(m, src)
}
func (m *QueryMarketMapperRevenueShareParams) XXX_Size() int {
	return m.Size()
}
func (m *QueryMarketMapperRevenueShareParams) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryMarketMapperRevenueShareParams.DiscardUnknown(m)
}

var xxx_messageInfo_QueryMarketMapperRevenueShareParams proto.InternalMessageInfo

// Response type for QueryMarketMapperRevenueShareParams
type QueryMarketMapperRevenueShareParamsResponse struct {
	Params MarketMapperRevenueShareParams `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryMarketMapperRevenueShareParamsResponse) Reset() {
	*m = QueryMarketMapperRevenueShareParamsResponse{}
}
func (m *QueryMarketMapperRevenueShareParamsResponse) String() string {
	return proto.CompactTextString(m)
}
func (*QueryMarketMapperRevenueShareParamsResponse) ProtoMessage() {}
func (*QueryMarketMapperRevenueShareParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_13d50c6e3048e744, []int{1}
}
func (m *QueryMarketMapperRevenueShareParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryMarketMapperRevenueShareParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryMarketMapperRevenueShareParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryMarketMapperRevenueShareParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryMarketMapperRevenueShareParamsResponse.Merge(m, src)
}
func (m *QueryMarketMapperRevenueShareParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryMarketMapperRevenueShareParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryMarketMapperRevenueShareParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryMarketMapperRevenueShareParamsResponse proto.InternalMessageInfo

func (m *QueryMarketMapperRevenueShareParamsResponse) GetParams() MarketMapperRevenueShareParams {
	if m != nil {
		return m.Params
	}
	return MarketMapperRevenueShareParams{}
}

// Queries market mapper revenue share details for a specific market
type QueryMarketMapperRevShareDetails struct {
	MarketId uint32 `protobuf:"varint,1,opt,name=market_id,json=marketId,proto3" json:"market_id,omitempty"`
}

func (m *QueryMarketMapperRevShareDetails) Reset()         { *m = QueryMarketMapperRevShareDetails{} }
func (m *QueryMarketMapperRevShareDetails) String() string { return proto.CompactTextString(m) }
func (*QueryMarketMapperRevShareDetails) ProtoMessage()    {}
func (*QueryMarketMapperRevShareDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_13d50c6e3048e744, []int{2}
}
func (m *QueryMarketMapperRevShareDetails) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryMarketMapperRevShareDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryMarketMapperRevShareDetails.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryMarketMapperRevShareDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryMarketMapperRevShareDetails.Merge(m, src)
}
func (m *QueryMarketMapperRevShareDetails) XXX_Size() int {
	return m.Size()
}
func (m *QueryMarketMapperRevShareDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryMarketMapperRevShareDetails.DiscardUnknown(m)
}

var xxx_messageInfo_QueryMarketMapperRevShareDetails proto.InternalMessageInfo

func (m *QueryMarketMapperRevShareDetails) GetMarketId() uint32 {
	if m != nil {
		return m.MarketId
	}
	return 0
}

// Response type for QueryMarketMapperRevShareDetails
type QueryMarketMapperRevShareDetailsResponse struct {
	Details MarketMapperRevShareDetails `protobuf:"bytes,1,opt,name=details,proto3" json:"details"`
}

func (m *QueryMarketMapperRevShareDetailsResponse) Reset() {
	*m = QueryMarketMapperRevShareDetailsResponse{}
}
func (m *QueryMarketMapperRevShareDetailsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryMarketMapperRevShareDetailsResponse) ProtoMessage()    {}
func (*QueryMarketMapperRevShareDetailsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_13d50c6e3048e744, []int{3}
}
func (m *QueryMarketMapperRevShareDetailsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryMarketMapperRevShareDetailsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryMarketMapperRevShareDetailsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryMarketMapperRevShareDetailsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryMarketMapperRevShareDetailsResponse.Merge(m, src)
}
func (m *QueryMarketMapperRevShareDetailsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryMarketMapperRevShareDetailsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryMarketMapperRevShareDetailsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryMarketMapperRevShareDetailsResponse proto.InternalMessageInfo

func (m *QueryMarketMapperRevShareDetailsResponse) GetDetails() MarketMapperRevShareDetails {
	if m != nil {
		return m.Details
	}
	return MarketMapperRevShareDetails{}
}

// Queries unconditional revenue share details
type QueryUnconditionalRevShareConfig struct {
}

func (m *QueryUnconditionalRevShareConfig) Reset()         { *m = QueryUnconditionalRevShareConfig{} }
func (m *QueryUnconditionalRevShareConfig) String() string { return proto.CompactTextString(m) }
func (*QueryUnconditionalRevShareConfig) ProtoMessage()    {}
func (*QueryUnconditionalRevShareConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_13d50c6e3048e744, []int{4}
}
func (m *QueryUnconditionalRevShareConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryUnconditionalRevShareConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryUnconditionalRevShareConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryUnconditionalRevShareConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUnconditionalRevShareConfig.Merge(m, src)
}
func (m *QueryUnconditionalRevShareConfig) XXX_Size() int {
	return m.Size()
}
func (m *QueryUnconditionalRevShareConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUnconditionalRevShareConfig.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUnconditionalRevShareConfig proto.InternalMessageInfo

// Response type for QueryUnconditionalRevShareConfig
type QueryUnconditionalRevShareConfigResponse struct {
	Config UnconditionalRevShareConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config"`
}

func (m *QueryUnconditionalRevShareConfigResponse) Reset() {
	*m = QueryUnconditionalRevShareConfigResponse{}
}
func (m *QueryUnconditionalRevShareConfigResponse) String() string { return proto.CompactTextString(m) }
func (*QueryUnconditionalRevShareConfigResponse) ProtoMessage()    {}
func (*QueryUnconditionalRevShareConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_13d50c6e3048e744, []int{5}
}
func (m *QueryUnconditionalRevShareConfigResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryUnconditionalRevShareConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryUnconditionalRevShareConfigResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryUnconditionalRevShareConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUnconditionalRevShareConfigResponse.Merge(m, src)
}
func (m *QueryUnconditionalRevShareConfigResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryUnconditionalRevShareConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUnconditionalRevShareConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUnconditionalRevShareConfigResponse proto.InternalMessageInfo

func (m *QueryUnconditionalRevShareConfigResponse) GetConfig() UnconditionalRevShareConfig {
	if m != nil {
		return m.Config
	}
	return UnconditionalRevShareConfig{}
}

func init() {
	proto.RegisterType((*QueryMarketMapperRevenueShareParams)(nil), "bitoroprotocol.revshare.QueryMarketMapperRevenueShareParams")
	proto.RegisterType((*QueryMarketMapperRevenueShareParamsResponse)(nil), "bitoroprotocol.revshare.QueryMarketMapperRevenueShareParamsResponse")
	proto.RegisterType((*QueryMarketMapperRevShareDetails)(nil), "bitoroprotocol.revshare.QueryMarketMapperRevShareDetails")
	proto.RegisterType((*QueryMarketMapperRevShareDetailsResponse)(nil), "bitoroprotocol.revshare.QueryMarketMapperRevShareDetailsResponse")
	proto.RegisterType((*QueryUnconditionalRevShareConfig)(nil), "bitoroprotocol.revshare.QueryUnconditionalRevShareConfig")
	proto.RegisterType((*QueryUnconditionalRevShareConfigResponse)(nil), "bitoroprotocol.revshare.QueryUnconditionalRevShareConfigResponse")
}

func init() { proto.RegisterFile("bitoroprotocol/revshare/query.proto", fileDescriptor_13d50c6e3048e744) }

var fileDescriptor_13d50c6e3048e744 = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x4f, 0x6b, 0xd4, 0x40,
	0x18, 0xc6, 0x77, 0xc4, 0xae, 0x3a, 0xe2, 0x65, 0x50, 0x90, 0x5d, 0x89, 0x75, 0x54, 0x28, 0x88,
	0x19, 0x59, 0xad, 0x85, 0x82, 0x14, 0xd6, 0x5e, 0x14, 0x0a, 0x35, 0xd5, 0x8b, 0x97, 0x65, 0x9a,
	0x8c, 0xd9, 0xc1, 0xdd, 0x99, 0x38, 0xc9, 0x86, 0x2e, 0xfe, 0x39, 0xf4, 0x13, 0x08, 0x7e, 0x02,
	0xbf, 0x4d, 0x6f, 0x16, 0xbc, 0x78, 0x52, 0xd9, 0xf8, 0x41, 0x24, 0x93, 0x49, 0x9a, 0x42, 0x3a,
	0x4d, 0x7b, 0x1b, 0xde, 0x79, 0xde, 0xe7, 0x7d, 0x7e, 0x99, 0x37, 0xf0, 0x4e, 0x30, 0x0f, 0xf6,
	0x22, 0x25, 0x13, 0xe9, 0xcb, 0x09, 0x51, 0x2c, 0x8d, 0xc7, 0x54, 0x31, 0xf2, 0x61, 0xc6, 0xd4,
	0xdc, 0xd5, 0x75, 0x74, 0xa3, 0x2e, 0x71, 0x4b, 0x49, 0xef, 0x7a, 0x28, 0x43, 0xa9, 0xcb, 0x24,
	0x3f, 0x15, 0xe2, 0xde, 0xad, 0x50, 0xca, 0x70, 0xc2, 0x08, 0x8d, 0x38, 0xa1, 0x42, 0xc8, 0x84,
	0x26, 0x5c, 0x8a, 0xd8, 0xdc, 0xe2, 0xe6, 0x69, 0x11, 0x55, 0x74, 0x5a, 0x6a, 0xee, 0x35, 0x6b,
	0xca, 0x43, 0xa1, 0xc2, 0xf7, 0xe1, 0xdd, 0x57, 0x79, 0xc6, 0x2d, 0xaa, 0xde, 0xb3, 0x64, 0x8b,
	0x46, 0x11, 0x53, 0x1e, 0x4b, 0x99, 0x98, 0xb1, 0x9d, 0x5c, 0xb6, 0xad, 0x2d, 0xf1, 0x3e, 0x80,
	0x0f, 0x5a, 0xe8, 0x3c, 0x16, 0x47, 0x52, 0xc4, 0x0c, 0xed, 0xc0, 0x6e, 0x11, 0xe6, 0x26, 0x58,
	0x06, 0x2b, 0x57, 0x07, 0xab, 0x6e, 0x23, 0xbc, 0x6b, 0xb7, 0x1b, 0x5e, 0x3c, 0xf8, 0x7d, 0xbb,
	0xe3, 0x19, 0x2b, 0xbc, 0x01, 0x97, 0x9b, 0x32, 0xe8, 0x86, 0x4d, 0x96, 0x50, 0x3e, 0x89, 0x51,
	0x1f, 0x5e, 0x99, 0xea, 0xeb, 0x11, 0x0f, 0xf4, 0xec, 0x6b, 0xde, 0xe5, 0xa2, 0xf0, 0x22, 0xc0,
	0x5f, 0xe0, 0xca, 0x69, 0x06, 0x15, 0x81, 0x07, 0x2f, 0x05, 0x45, 0xc9, 0x20, 0x0c, 0xda, 0x21,
	0xd4, 0xcd, 0x4c, 0xfe, 0xd2, 0x08, 0x63, 0x03, 0xf0, 0x46, 0xf8, 0x52, 0x04, 0x3c, 0x7f, 0x51,
	0x3a, 0x29, 0x7b, 0x9e, 0x4b, 0xf1, 0x8e, 0x87, 0xf8, 0x93, 0xc9, 0x68, 0xd1, 0x54, 0x19, 0xb7,
	0x61, 0xd7, 0xd7, 0x95, 0x53, 0x22, 0x5a, 0xbc, 0xca, 0x4f, 0x5c, 0xf8, 0x0c, 0xbe, 0x2f, 0xc1,
	0x25, 0x3d, 0x1e, 0xfd, 0x01, 0xd0, 0xb1, 0xbf, 0x0e, 0x5a, 0x3f, 0x61, 0x5c, 0x8b, 0x45, 0xe9,
	0x0d, 0xcf, 0xdf, 0x5b, 0xe2, 0xe3, 0x67, 0xfb, 0x3f, 0xff, 0x7d, 0xbb, 0xb0, 0x86, 0x56, 0x49,
	0xf3, 0xaa, 0x9b, 0x45, 0x98, 0x6a, 0x9f, 0x91, 0x62, 0xe9, 0x48, 0xd7, 0x47, 0xc5, 0x3a, 0xa1,
	0x0c, 0xc0, 0xbe, 0x6d, 0x95, 0xd6, 0xce, 0x10, 0xb1, 0xde, 0xd8, 0xdb, 0x38, 0x67, 0x63, 0x05,
	0xf6, 0x52, 0x83, 0x6d, 0xa2, 0xe1, 0x19, 0xc1, 0xcc, 0x9e, 0x91, 0x8f, 0xd5, 0x2f, 0xf0, 0x19,
	0xfd, 0x00, 0xb0, 0x6f, 0x79, 0x7f, 0x3b, 0xa5, 0xa5, 0xd1, 0x4e, 0xd9, 0x62, 0x7b, 0xf1, 0x53,
	0x4d, 0xf9, 0x08, 0xb9, 0x27, 0x50, 0xce, 0xea, 0x1e, 0x47, 0x94, 0xc3, 0xd7, 0x07, 0x0b, 0x07,
	0x1c, 0x2e, 0x1c, 0xf0, 0x77, 0xe1, 0x80, 0xaf, 0x99, 0xd3, 0x39, 0xcc, 0x9c, 0xce, 0xaf, 0xcc,
	0xe9, 0xbc, 0x5d, 0x0f, 0x79, 0x32, 0x9e, 0xed, 0xba, 0xbe, 0x9c, 0x1e, 0xf7, 0x4c, 0x9f, 0x3c,
	0xf4, 0xc7, 0x94, 0x0b, 0x52, 0x55, 0xf6, 0x8e, 0xe6, 0x24, 0xf3, 0x88, 0xc5, 0xbb, 0x5d, 0x7d,
	0xf5, 0xf8, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x34, 0xdf, 0xbf, 0xc9, 0x05, 0x00, 0x00,
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
	// MarketMapperRevenueShareParams queries the revenue share params for the
	// market mapper
	MarketMapperRevenueShareParams(ctx context.Context, in *QueryMarketMapperRevenueShareParams, opts ...grpc.CallOption) (*QueryMarketMapperRevenueShareParamsResponse, error)
	// Queries market mapper revenue share details for a specific market
	MarketMapperRevShareDetails(ctx context.Context, in *QueryMarketMapperRevShareDetails, opts ...grpc.CallOption) (*QueryMarketMapperRevShareDetailsResponse, error)
	// Queries unconditional revenue share config
	UnconditionalRevShareConfig(ctx context.Context, in *QueryUnconditionalRevShareConfig, opts ...grpc.CallOption) (*QueryUnconditionalRevShareConfigResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) MarketMapperRevenueShareParams(ctx context.Context, in *QueryMarketMapperRevenueShareParams, opts ...grpc.CallOption) (*QueryMarketMapperRevenueShareParamsResponse, error) {
	out := new(QueryMarketMapperRevenueShareParamsResponse)
	err := c.cc.Invoke(ctx, "/bitoroprotocol.revshare.Query/MarketMapperRevenueShareParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) MarketMapperRevShareDetails(ctx context.Context, in *QueryMarketMapperRevShareDetails, opts ...grpc.CallOption) (*QueryMarketMapperRevShareDetailsResponse, error) {
	out := new(QueryMarketMapperRevShareDetailsResponse)
	err := c.cc.Invoke(ctx, "/bitoroprotocol.revshare.Query/MarketMapperRevShareDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) UnconditionalRevShareConfig(ctx context.Context, in *QueryUnconditionalRevShareConfig, opts ...grpc.CallOption) (*QueryUnconditionalRevShareConfigResponse, error) {
	out := new(QueryUnconditionalRevShareConfigResponse)
	err := c.cc.Invoke(ctx, "/bitoroprotocol.revshare.Query/UnconditionalRevShareConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// MarketMapperRevenueShareParams queries the revenue share params for the
	// market mapper
	MarketMapperRevenueShareParams(context.Context, *QueryMarketMapperRevenueShareParams) (*QueryMarketMapperRevenueShareParamsResponse, error)
	// Queries market mapper revenue share details for a specific market
	MarketMapperRevShareDetails(context.Context, *QueryMarketMapperRevShareDetails) (*QueryMarketMapperRevShareDetailsResponse, error)
	// Queries unconditional revenue share config
	UnconditionalRevShareConfig(context.Context, *QueryUnconditionalRevShareConfig) (*QueryUnconditionalRevShareConfigResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) MarketMapperRevenueShareParams(ctx context.Context, req *QueryMarketMapperRevenueShareParams) (*QueryMarketMapperRevenueShareParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarketMapperRevenueShareParams not implemented")
}
func (*UnimplementedQueryServer) MarketMapperRevShareDetails(ctx context.Context, req *QueryMarketMapperRevShareDetails) (*QueryMarketMapperRevShareDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarketMapperRevShareDetails not implemented")
}
func (*UnimplementedQueryServer) UnconditionalRevShareConfig(ctx context.Context, req *QueryUnconditionalRevShareConfig) (*QueryUnconditionalRevShareConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnconditionalRevShareConfig not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_MarketMapperRevenueShareParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryMarketMapperRevenueShareParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).MarketMapperRevenueShareParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bitoroprotocol.revshare.Query/MarketMapperRevenueShareParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).MarketMapperRevenueShareParams(ctx, req.(*QueryMarketMapperRevenueShareParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_MarketMapperRevShareDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryMarketMapperRevShareDetails)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).MarketMapperRevShareDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bitoroprotocol.revshare.Query/MarketMapperRevShareDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).MarketMapperRevShareDetails(ctx, req.(*QueryMarketMapperRevShareDetails))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_UnconditionalRevShareConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryUnconditionalRevShareConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).UnconditionalRevShareConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bitoroprotocol.revshare.Query/UnconditionalRevShareConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).UnconditionalRevShareConfig(ctx, req.(*QueryUnconditionalRevShareConfig))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bitoroprotocol.revshare.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MarketMapperRevenueShareParams",
			Handler:    _Query_MarketMapperRevenueShareParams_Handler,
		},
		{
			MethodName: "MarketMapperRevShareDetails",
			Handler:    _Query_MarketMapperRevShareDetails_Handler,
		},
		{
			MethodName: "UnconditionalRevShareConfig",
			Handler:    _Query_UnconditionalRevShareConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bitoroprotocol/revshare/query.proto",
}

func (m *QueryMarketMapperRevenueShareParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryMarketMapperRevenueShareParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryMarketMapperRevenueShareParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryMarketMapperRevenueShareParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryMarketMapperRevenueShareParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryMarketMapperRevenueShareParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryMarketMapperRevShareDetails) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryMarketMapperRevShareDetails) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryMarketMapperRevShareDetails) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MarketId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.MarketId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryMarketMapperRevShareDetailsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryMarketMapperRevShareDetailsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryMarketMapperRevShareDetailsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Details.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryUnconditionalRevShareConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryUnconditionalRevShareConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryUnconditionalRevShareConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryUnconditionalRevShareConfigResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryUnconditionalRevShareConfigResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryUnconditionalRevShareConfigResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Config.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
func (m *QueryMarketMapperRevenueShareParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryMarketMapperRevenueShareParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryMarketMapperRevShareDetails) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MarketId != 0 {
		n += 1 + sovQuery(uint64(m.MarketId))
	}
	return n
}

func (m *QueryMarketMapperRevShareDetailsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Details.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryUnconditionalRevShareConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryUnconditionalRevShareConfigResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Config.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryMarketMapperRevenueShareParams) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryMarketMapperRevenueShareParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryMarketMapperRevenueShareParams: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryMarketMapperRevenueShareParamsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryMarketMapperRevenueShareParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryMarketMapperRevenueShareParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryMarketMapperRevShareDetails) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryMarketMapperRevShareDetails: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryMarketMapperRevShareDetails: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarketId", wireType)
			}
			m.MarketId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MarketId |= uint32(b&0x7F) << shift
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
func (m *QueryMarketMapperRevShareDetailsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryMarketMapperRevShareDetailsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryMarketMapperRevShareDetailsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Details", wireType)
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
			if err := m.Details.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryUnconditionalRevShareConfig) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryUnconditionalRevShareConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryUnconditionalRevShareConfig: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryUnconditionalRevShareConfigResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryUnconditionalRevShareConfigResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryUnconditionalRevShareConfigResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
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
			if err := m.Config.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
