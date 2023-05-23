// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: decent/request/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e764b566425bfc21, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e764b566425bfc21, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryGetMinerResponseRequest struct {
	Index string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
}

func (m *QueryGetMinerResponseRequest) Reset()         { *m = QueryGetMinerResponseRequest{} }
func (m *QueryGetMinerResponseRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetMinerResponseRequest) ProtoMessage()    {}
func (*QueryGetMinerResponseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e764b566425bfc21, []int{2}
}
func (m *QueryGetMinerResponseRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetMinerResponseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetMinerResponseRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetMinerResponseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetMinerResponseRequest.Merge(m, src)
}
func (m *QueryGetMinerResponseRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetMinerResponseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetMinerResponseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetMinerResponseRequest proto.InternalMessageInfo

func (m *QueryGetMinerResponseRequest) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

type QueryGetMinerResponseResponse struct {
	MinerResponse MinerResponse `protobuf:"bytes,1,opt,name=minerResponse,proto3" json:"minerResponse"`
}

func (m *QueryGetMinerResponseResponse) Reset()         { *m = QueryGetMinerResponseResponse{} }
func (m *QueryGetMinerResponseResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetMinerResponseResponse) ProtoMessage()    {}
func (*QueryGetMinerResponseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e764b566425bfc21, []int{3}
}
func (m *QueryGetMinerResponseResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetMinerResponseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetMinerResponseResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetMinerResponseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetMinerResponseResponse.Merge(m, src)
}
func (m *QueryGetMinerResponseResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetMinerResponseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetMinerResponseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetMinerResponseResponse proto.InternalMessageInfo

func (m *QueryGetMinerResponseResponse) GetMinerResponse() MinerResponse {
	if m != nil {
		return m.MinerResponse
	}
	return MinerResponse{}
}

type QueryAllMinerResponseRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllMinerResponseRequest) Reset()         { *m = QueryAllMinerResponseRequest{} }
func (m *QueryAllMinerResponseRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllMinerResponseRequest) ProtoMessage()    {}
func (*QueryAllMinerResponseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e764b566425bfc21, []int{4}
}
func (m *QueryAllMinerResponseRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllMinerResponseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllMinerResponseRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllMinerResponseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllMinerResponseRequest.Merge(m, src)
}
func (m *QueryAllMinerResponseRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllMinerResponseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllMinerResponseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllMinerResponseRequest proto.InternalMessageInfo

func (m *QueryAllMinerResponseRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllMinerResponseResponse struct {
	MinerResponse []MinerResponse     `protobuf:"bytes,1,rep,name=minerResponse,proto3" json:"minerResponse"`
	Pagination    *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllMinerResponseResponse) Reset()         { *m = QueryAllMinerResponseResponse{} }
func (m *QueryAllMinerResponseResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllMinerResponseResponse) ProtoMessage()    {}
func (*QueryAllMinerResponseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e764b566425bfc21, []int{5}
}
func (m *QueryAllMinerResponseResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllMinerResponseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllMinerResponseResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllMinerResponseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllMinerResponseResponse.Merge(m, src)
}
func (m *QueryAllMinerResponseResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllMinerResponseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllMinerResponseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllMinerResponseResponse proto.InternalMessageInfo

func (m *QueryAllMinerResponseResponse) GetMinerResponse() []MinerResponse {
	if m != nil {
		return m.MinerResponse
	}
	return nil
}

func (m *QueryAllMinerResponseResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "decent.request.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "decent.request.QueryParamsResponse")
	proto.RegisterType((*QueryGetMinerResponseRequest)(nil), "decent.request.QueryGetMinerResponseRequest")
	proto.RegisterType((*QueryGetMinerResponseResponse)(nil), "decent.request.QueryGetMinerResponseResponse")
	proto.RegisterType((*QueryAllMinerResponseRequest)(nil), "decent.request.QueryAllMinerResponseRequest")
	proto.RegisterType((*QueryAllMinerResponseResponse)(nil), "decent.request.QueryAllMinerResponseResponse")
}

func init() { proto.RegisterFile("decent/request/query.proto", fileDescriptor_e764b566425bfc21) }

var fileDescriptor_e764b566425bfc21 = []byte{
	// 484 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0x33, 0xad, 0x0d, 0x38, 0x52, 0x91, 0x31, 0x84, 0xb2, 0xb6, 0x63, 0x99, 0x42, 0x2c,
	0xa2, 0x33, 0xb6, 0xf6, 0x0b, 0xb4, 0x88, 0x45, 0x44, 0xa8, 0x39, 0x7a, 0x91, 0x49, 0x3b, 0x2e,
	0x2b, 0x9b, 0x99, 0xcd, 0xce, 0x54, 0x5a, 0xc4, 0x8b, 0x9f, 0x40, 0xf0, 0xe0, 0xc5, 0x6f, 0xd1,
	0x2f, 0xd1, 0x63, 0xc1, 0x8b, 0x27, 0x91, 0xc4, 0x0f, 0x22, 0x99, 0x79, 0xd1, 0xcc, 0x36, 0x1b,
	0xa5, 0xb7, 0x24, 0xef, 0xff, 0x7f, 0xef, 0xf7, 0xf6, 0xfd, 0xb3, 0x38, 0x39, 0x52, 0x87, 0x4a,
	0x3b, 0x51, 0xaa, 0xc1, 0xb1, 0xb2, 0x4e, 0x0c, 0x8e, 0x55, 0x79, 0xca, 0x8b, 0xd2, 0x38, 0x43,
	0x6e, 0x86, 0x1a, 0x87, 0x5a, 0xd2, 0x4a, 0x4d, 0x6a, 0x7c, 0x49, 0x8c, 0x3f, 0x05, 0x55, 0xb2,
	0x9a, 0x1a, 0x93, 0xe6, 0x4a, 0xc8, 0x22, 0x13, 0x52, 0x6b, 0xe3, 0xa4, 0xcb, 0x8c, 0xb6, 0x50,
	0xbd, 0x7f, 0x68, 0x6c, 0xdf, 0x58, 0xd1, 0x93, 0x56, 0x85, 0xe6, 0xe2, 0xdd, 0x56, 0x4f, 0x39,
	0xb9, 0x25, 0x0a, 0x99, 0x66, 0xda, 0x8b, 0x41, 0x7b, 0xa7, 0xc2, 0x52, 0xc8, 0x52, 0xf6, 0x27,
	0x8d, 0x36, 0x2a, 0xc5, 0x7e, 0xa6, 0x55, 0xf9, 0xba, 0x54, 0xb6, 0x30, 0xda, 0xaa, 0x20, 0x62,
	0x2d, 0x4c, 0x5e, 0x8e, 0x67, 0x1c, 0x78, 0x67, 0x37, 0x48, 0xd9, 0x73, 0x7c, 0x3b, 0xfa, 0x35,
	0x58, 0xc8, 0x0e, 0x6e, 0x86, 0x09, 0x2b, 0x68, 0x1d, 0x6d, 0xde, 0xd8, 0x6e, 0xf3, 0x78, 0x5f,
	0x1e, 0xf4, 0x7b, 0xd7, 0xce, 0x7f, 0xdc, 0x6d, 0x74, 0x41, 0xcb, 0x76, 0xf0, 0xaa, 0x6f, 0xb6,
	0xaf, 0xdc, 0x8b, 0x31, 0xc2, 0xa4, 0x1d, 0x0c, 0x23, 0x2d, 0xbc, 0x94, 0xe9, 0x23, 0x75, 0xe2,
	0x9b, 0x5e, 0xef, 0x86, 0x2f, 0xec, 0x2d, 0x5e, 0xab, 0x71, 0x01, 0xcc, 0x33, 0xbc, 0xdc, 0x9f,
	0x2e, 0x00, 0xd3, 0x5a, 0x95, 0x29, 0x72, 0x03, 0x5a, 0xec, 0x64, 0x6f, 0x80, 0x70, 0x37, 0xcf,
	0x67, 0x12, 0x3e, 0xc5, 0xf8, 0xef, 0xa3, 0x87, 0x39, 0x1d, 0x1e, 0xee, 0xc4, 0xc7, 0x77, 0xe2,
	0x21, 0x04, 0x70, 0x27, 0x7e, 0x20, 0xd3, 0x89, 0xb7, 0x3b, 0xe5, 0x64, 0x67, 0x08, 0x96, 0xba,
	0x3c, 0xa8, 0x7e, 0xa9, 0xc5, 0xab, 0x2d, 0x45, 0xf6, 0x23, 0xe8, 0x05, 0x0f, 0x7d, 0xef, 0x9f,
	0xd0, 0xc0, 0x33, 0x65, 0xdd, 0x3e, 0x5b, 0xc4, 0x4b, 0x9e, 0x9a, 0x0c, 0x70, 0x33, 0x5c, 0x98,
	0xb0, 0x2a, 0xd0, 0xe5, 0x10, 0x25, 0x1b, 0x73, 0x35, 0xf0, 0xe8, 0xe9, 0xc7, 0x6f, 0xbf, 0x3e,
	0x2f, 0xac, 0x90, 0xb6, 0x78, 0x32, 0x2b, 0xca, 0xe4, 0x2b, 0xc2, 0xcb, 0xd1, 0xb2, 0xe4, 0xc1,
	0xcc, 0xb6, 0x35, 0xe1, 0x4a, 0x1e, 0xfe, 0xa7, 0x1a, 0x70, 0xb8, 0xc7, 0xd9, 0x24, 0x9d, 0x2a,
	0x4e, 0xfc, 0xe7, 0x11, 0xef, 0x7d, 0x48, 0x3f, 0x90, 0x2f, 0x08, 0xdf, 0x8a, 0x3a, 0xed, 0xe6,
	0x79, 0x0d, 0x61, 0x4d, 0xb8, 0x6a, 0x08, 0xeb, 0x12, 0xc2, 0x3a, 0x9e, 0x70, 0x9d, 0xd0, 0xf9,
	0x84, 0x7b, 0x8f, 0xce, 0x87, 0x14, 0x5d, 0x0c, 0x29, 0xfa, 0x39, 0xa4, 0xe8, 0xd3, 0x88, 0x36,
	0x2e, 0x46, 0xb4, 0xf1, 0x7d, 0x44, 0x1b, 0xaf, 0xda, 0x60, 0x3c, 0xf9, 0x63, 0x75, 0xa7, 0x85,
	0xb2, 0xbd, 0xa6, 0x7f, 0x23, 0x3c, 0xfe, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x68, 0x2b, 0x7e,
	0xe1, 0x04, 0x00, 0x00,
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
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of MinerResponse items.
	MinerResponse(ctx context.Context, in *QueryGetMinerResponseRequest, opts ...grpc.CallOption) (*QueryGetMinerResponseResponse, error)
	MinerResponseAll(ctx context.Context, in *QueryAllMinerResponseRequest, opts ...grpc.CallOption) (*QueryAllMinerResponseResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/decent.request.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) MinerResponse(ctx context.Context, in *QueryGetMinerResponseRequest, opts ...grpc.CallOption) (*QueryGetMinerResponseResponse, error) {
	out := new(QueryGetMinerResponseResponse)
	err := c.cc.Invoke(ctx, "/decent.request.Query/MinerResponse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) MinerResponseAll(ctx context.Context, in *QueryAllMinerResponseRequest, opts ...grpc.CallOption) (*QueryAllMinerResponseResponse, error) {
	out := new(QueryAllMinerResponseResponse)
	err := c.cc.Invoke(ctx, "/decent.request.Query/MinerResponseAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of MinerResponse items.
	MinerResponse(context.Context, *QueryGetMinerResponseRequest) (*QueryGetMinerResponseResponse, error)
	MinerResponseAll(context.Context, *QueryAllMinerResponseRequest) (*QueryAllMinerResponseResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) MinerResponse(ctx context.Context, req *QueryGetMinerResponseRequest) (*QueryGetMinerResponseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MinerResponse not implemented")
}
func (*UnimplementedQueryServer) MinerResponseAll(ctx context.Context, req *QueryAllMinerResponseRequest) (*QueryAllMinerResponseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MinerResponseAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/decent.request.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_MinerResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetMinerResponseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).MinerResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/decent.request.Query/MinerResponse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).MinerResponse(ctx, req.(*QueryGetMinerResponseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_MinerResponseAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllMinerResponseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).MinerResponseAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/decent.request.Query/MinerResponseAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).MinerResponseAll(ctx, req.(*QueryAllMinerResponseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "decent.request.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "MinerResponse",
			Handler:    _Query_MinerResponse_Handler,
		},
		{
			MethodName: "MinerResponseAll",
			Handler:    _Query_MinerResponseAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "decent/request/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *QueryGetMinerResponseRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetMinerResponseRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetMinerResponseRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetMinerResponseResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetMinerResponseResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetMinerResponseResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.MinerResponse.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryAllMinerResponseRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllMinerResponseRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllMinerResponseRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *QueryAllMinerResponseResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllMinerResponseResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllMinerResponseResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.MinerResponse) > 0 {
		for iNdEx := len(m.MinerResponse) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MinerResponse[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryGetMinerResponseRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetMinerResponseResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MinerResponse.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllMinerResponseRequest) Size() (n int) {
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

func (m *QueryAllMinerResponseResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.MinerResponse) > 0 {
		for _, e := range m.MinerResponse {
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
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryGetMinerResponseRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetMinerResponseRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetMinerResponseRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
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
			m.Index = string(dAtA[iNdEx:postIndex])
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
func (m *QueryGetMinerResponseResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGetMinerResponseResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetMinerResponseResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinerResponse", wireType)
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
			if err := m.MinerResponse.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryAllMinerResponseRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllMinerResponseRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllMinerResponseRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryAllMinerResponseResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAllMinerResponseResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllMinerResponseResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinerResponse", wireType)
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
			m.MinerResponse = append(m.MinerResponse, MinerResponse{})
			if err := m.MinerResponse[len(m.MinerResponse)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
