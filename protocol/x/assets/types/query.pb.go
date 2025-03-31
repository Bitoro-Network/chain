// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bitoroprotocol/assets/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
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

// Queries an Asset by id.
type QueryAssetRequest struct {
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryAssetRequest) Reset()         { *m = QueryAssetRequest{} }
func (m *QueryAssetRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAssetRequest) ProtoMessage()    {}
func (*QueryAssetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e6c21d5bfb3fef3, []int{0}
}
func (m *QueryAssetRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAssetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAssetRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAssetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAssetRequest.Merge(m, src)
}
func (m *QueryAssetRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAssetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAssetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAssetRequest proto.InternalMessageInfo

func (m *QueryAssetRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

// QueryAssetResponse is response type for the Asset RPC method.
type QueryAssetResponse struct {
	Asset Asset `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset"`
}

func (m *QueryAssetResponse) Reset()         { *m = QueryAssetResponse{} }
func (m *QueryAssetResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAssetResponse) ProtoMessage()    {}
func (*QueryAssetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e6c21d5bfb3fef3, []int{1}
}
func (m *QueryAssetResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAssetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAssetResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAssetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAssetResponse.Merge(m, src)
}
func (m *QueryAssetResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAssetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAssetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAssetResponse proto.InternalMessageInfo

func (m *QueryAssetResponse) GetAsset() Asset {
	if m != nil {
		return m.Asset
	}
	return Asset{}
}

// Queries a list of Asset items.
type QueryAllAssetsRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllAssetsRequest) Reset()         { *m = QueryAllAssetsRequest{} }
func (m *QueryAllAssetsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllAssetsRequest) ProtoMessage()    {}
func (*QueryAllAssetsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e6c21d5bfb3fef3, []int{2}
}
func (m *QueryAllAssetsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllAssetsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllAssetsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllAssetsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllAssetsRequest.Merge(m, src)
}
func (m *QueryAllAssetsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllAssetsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllAssetsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllAssetsRequest proto.InternalMessageInfo

func (m *QueryAllAssetsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryAllAssetsResponse is response type for the AllAssets RPC method.
type QueryAllAssetsResponse struct {
	Asset      []Asset             `protobuf:"bytes,1,rep,name=asset,proto3" json:"asset"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllAssetsResponse) Reset()         { *m = QueryAllAssetsResponse{} }
func (m *QueryAllAssetsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllAssetsResponse) ProtoMessage()    {}
func (*QueryAllAssetsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e6c21d5bfb3fef3, []int{3}
}
func (m *QueryAllAssetsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllAssetsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllAssetsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllAssetsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllAssetsResponse.Merge(m, src)
}
func (m *QueryAllAssetsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllAssetsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllAssetsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllAssetsResponse proto.InternalMessageInfo

func (m *QueryAllAssetsResponse) GetAsset() []Asset {
	if m != nil {
		return m.Asset
	}
	return nil
}

func (m *QueryAllAssetsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryAssetRequest)(nil), "bitoroprotocol.assets.QueryAssetRequest")
	proto.RegisterType((*QueryAssetResponse)(nil), "bitoroprotocol.assets.QueryAssetResponse")
	proto.RegisterType((*QueryAllAssetsRequest)(nil), "bitoroprotocol.assets.QueryAllAssetsRequest")
	proto.RegisterType((*QueryAllAssetsResponse)(nil), "bitoroprotocol.assets.QueryAllAssetsResponse")
}

func init() { proto.RegisterFile("bitoroprotocol/assets/query.proto", fileDescriptor_8e6c21d5bfb3fef3) }

var fileDescriptor_8e6c21d5bfb3fef3 = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xcf, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x33, 0xd1, 0x0a, 0x8e, 0x28, 0x38, 0xfe, 0x40, 0x42, 0x49, 0x35, 0x42, 0x2b, 0x15,
	0x67, 0x68, 0x15, 0xf1, 0x6a, 0x0f, 0x7a, 0xf1, 0xa0, 0x39, 0x7a, 0x91, 0x49, 0x32, 0xa4, 0x03,
	0x69, 0x26, 0xed, 0x4c, 0x4b, 0x8b, 0x78, 0xd0, 0x93, 0x47, 0x41, 0x10, 0xfc, 0x8f, 0x7a, 0x2c,
	0xec, 0x65, 0x4f, 0xcb, 0xd2, 0xee, 0x1f, 0xb2, 0x64, 0x26, 0xed, 0xa6, 0xbf, 0xb6, 0xbb, 0xa7,
	0x84, 0xf7, 0xbe, 0xef, 0xfb, 0xfd, 0xbc, 0x79, 0xb0, 0x16, 0x4d, 0xa2, 0x71, 0x36, 0x10, 0x4a,
	0x84, 0x22, 0x21, 0x54, 0x4a, 0xa6, 0x24, 0xe9, 0x0f, 0xd9, 0x60, 0x82, 0x75, 0x15, 0x3d, 0x28,
	0x0b, 0xb0, 0x11, 0x38, 0x0f, 0x63, 0x11, 0x0b, 0x5d, 0x24, 0xf9, 0x9f, 0x91, 0x3a, 0xd5, 0x58,
	0x88, 0x38, 0x61, 0x84, 0x66, 0x9c, 0xd0, 0x34, 0x15, 0x8a, 0x2a, 0x2e, 0x52, 0x59, 0x74, 0x9b,
	0xa1, 0x90, 0x3d, 0x21, 0x49, 0x40, 0x25, 0x33, 0x09, 0x64, 0xd4, 0x0a, 0x98, 0xa2, 0x2d, 0x92,
	0xd1, 0x98, 0xa7, 0x5a, 0x5c, 0x68, 0x77, 0x52, 0xe9, 0x8f, 0x11, 0x78, 0xcf, 0xe1, 0xfd, 0x2f,
	0xb9, 0xc5, 0xfb, 0xbc, 0xe6, 0xb3, 0xfe, 0x90, 0x49, 0x85, 0xee, 0x41, 0x9b, 0x47, 0x4f, 0xc0,
	0x53, 0xf0, 0xe2, 0xae, 0x6f, 0xf3, 0xc8, 0xfb, 0x04, 0x51, 0x59, 0x24, 0x33, 0x91, 0x4a, 0x86,
	0xde, 0xc2, 0x8a, 0x76, 0xd2, 0xc2, 0x3b, 0x6d, 0x07, 0xef, 0x58, 0x10, 0xeb, 0x91, 0xce, 0xcd,
	0xe9, 0x49, 0xcd, 0xf2, 0x8d, 0xdc, 0xfb, 0x06, 0x1f, 0x19, 0xb7, 0x24, 0xd1, 0x5d, 0xb9, 0x8c,
	0xfd, 0x00, 0xe1, 0xc5, 0x02, 0x85, 0x6b, 0x1d, 0x9b, 0x6d, 0x71, 0xbe, 0x2d, 0x36, 0xef, 0x59,
	0x6c, 0x8b, 0x3f, 0xd3, 0x98, 0x15, 0xb3, 0x7e, 0x69, 0xd2, 0xfb, 0x0f, 0xe0, 0xe3, 0xcd, 0x84,
	0x6d, 0xe6, 0x1b, 0xd7, 0x60, 0x46, 0x1f, 0xd7, 0xd0, 0x6c, 0x8d, 0xd6, 0x38, 0x88, 0x66, 0x42,
	0xcb, 0x6c, 0xed, 0x7f, 0x36, 0xac, 0x68, 0x36, 0xf4, 0x13, 0xc0, 0x8a, 0x4e, 0x42, 0xf5, 0x9d,
	0x14, 0x5b, 0x67, 0x71, 0x1a, 0x07, 0x75, 0x26, 0xd0, 0x6b, 0xfc, 0x3a, 0x3a, 0xfb, 0x6b, 0x3f,
	0x43, 0x35, 0xb2, 0xf7, 0xfc, 0xe4, 0x3b, 0x8f, 0x7e, 0xa0, 0xdf, 0x00, 0xde, 0x5e, 0x3d, 0x12,
	0x6a, 0x5e, 0xe2, 0xbf, 0x71, 0x2b, 0xe7, 0xe5, 0x95, 0xb4, 0x05, 0x8f, 0xa7, 0x79, 0xaa, 0xc8,
	0xd9, 0xcf, 0xd3, 0xf1, 0xa7, 0x73, 0x17, 0xcc, 0xe6, 0x2e, 0x38, 0x9d, 0xbb, 0xe0, 0xcf, 0xc2,
	0xb5, 0x66, 0x0b, 0xd7, 0x3a, 0x5e, 0xb8, 0xd6, 0xd7, 0x77, 0x31, 0x57, 0xdd, 0x61, 0x80, 0x43,
	0xd1, 0x5b, 0x9f, 0x1f, 0xbd, 0x79, 0x15, 0x76, 0x29, 0x4f, 0xc9, 0xaa, 0x32, 0x5e, 0x7a, 0xaa,
	0x49, 0xc6, 0x64, 0x70, 0x4b, 0x37, 0x5e, 0x9f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x98, 0xa2, 0x50,
	0xbb, 0x9c, 0x03, 0x00, 0x00,
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
	// Queries a Asset by id.
	Asset(ctx context.Context, in *QueryAssetRequest, opts ...grpc.CallOption) (*QueryAssetResponse, error)
	// Queries a list of Asset items.
	AllAssets(ctx context.Context, in *QueryAllAssetsRequest, opts ...grpc.CallOption) (*QueryAllAssetsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Asset(ctx context.Context, in *QueryAssetRequest, opts ...grpc.CallOption) (*QueryAssetResponse, error) {
	out := new(QueryAssetResponse)
	err := c.cc.Invoke(ctx, "/bitoroprotocol.assets.Query/Asset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AllAssets(ctx context.Context, in *QueryAllAssetsRequest, opts ...grpc.CallOption) (*QueryAllAssetsResponse, error) {
	out := new(QueryAllAssetsResponse)
	err := c.cc.Invoke(ctx, "/bitoroprotocol.assets.Query/AllAssets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries a Asset by id.
	Asset(context.Context, *QueryAssetRequest) (*QueryAssetResponse, error)
	// Queries a list of Asset items.
	AllAssets(context.Context, *QueryAllAssetsRequest) (*QueryAllAssetsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Asset(ctx context.Context, req *QueryAssetRequest) (*QueryAssetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Asset not implemented")
}
func (*UnimplementedQueryServer) AllAssets(ctx context.Context, req *QueryAllAssetsRequest) (*QueryAllAssetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllAssets not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Asset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Asset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bitoroprotocol.assets.Query/Asset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Asset(ctx, req.(*QueryAssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AllAssets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllAssetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AllAssets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bitoroprotocol.assets.Query/AllAssets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AllAssets(ctx, req.(*QueryAllAssetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bitoroprotocol.assets.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Asset",
			Handler:    _Query_Asset_Handler,
		},
		{
			MethodName: "AllAssets",
			Handler:    _Query_AllAssets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bitoroprotocol/assets/query.proto",
}

func (m *QueryAssetRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAssetRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAssetRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *QueryAssetResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAssetResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAssetResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Asset.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryAllAssetsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllAssetsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllAssetsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryAllAssetsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllAssetsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllAssetsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Asset) > 0 {
		for iNdEx := len(m.Asset) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Asset[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryAssetRequest) Size() (n int) {
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

func (m *QueryAssetResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Asset.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllAssetsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllAssetsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Asset) > 0 {
		for _, e := range m.Asset {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryAssetRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAssetRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAssetRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryAssetResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAssetResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAssetResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset", wireType)
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
			if err := m.Asset.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryAllAssetsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllAssetsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllAssetsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryAllAssetsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllAssetsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllAssetsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset", wireType)
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
			m.Asset = append(m.Asset, Asset{})
			if err := m.Asset[len(m.Asset)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
