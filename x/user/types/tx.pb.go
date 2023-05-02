// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: decent/user/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type MsgDepositToken struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Deposit string `protobuf:"bytes,2,opt,name=deposit,proto3" json:"deposit,omitempty"`
}

func (m *MsgDepositToken) Reset()         { *m = MsgDepositToken{} }
func (m *MsgDepositToken) String() string { return proto.CompactTextString(m) }
func (*MsgDepositToken) ProtoMessage()    {}
func (*MsgDepositToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_1719cb30b3bc4bbd, []int{0}
}
func (m *MsgDepositToken) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDepositToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDepositToken.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDepositToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDepositToken.Merge(m, src)
}
func (m *MsgDepositToken) XXX_Size() int {
	return m.Size()
}
func (m *MsgDepositToken) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDepositToken.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDepositToken proto.InternalMessageInfo

func (m *MsgDepositToken) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgDepositToken) GetDeposit() string {
	if m != nil {
		return m.Deposit
	}
	return ""
}

type MsgDepositTokenResponse struct {
}

func (m *MsgDepositTokenResponse) Reset()         { *m = MsgDepositTokenResponse{} }
func (m *MsgDepositTokenResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDepositTokenResponse) ProtoMessage()    {}
func (*MsgDepositTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1719cb30b3bc4bbd, []int{1}
}
func (m *MsgDepositTokenResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDepositTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDepositTokenResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDepositTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDepositTokenResponse.Merge(m, src)
}
func (m *MsgDepositTokenResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDepositTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDepositTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDepositTokenResponse proto.InternalMessageInfo

type MsgWithdrawToken struct {
	Creator  string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Withdraw string `protobuf:"bytes,2,opt,name=withdraw,proto3" json:"withdraw,omitempty"`
}

func (m *MsgWithdrawToken) Reset()         { *m = MsgWithdrawToken{} }
func (m *MsgWithdrawToken) String() string { return proto.CompactTextString(m) }
func (*MsgWithdrawToken) ProtoMessage()    {}
func (*MsgWithdrawToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_1719cb30b3bc4bbd, []int{2}
}
func (m *MsgWithdrawToken) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWithdrawToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWithdrawToken.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWithdrawToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithdrawToken.Merge(m, src)
}
func (m *MsgWithdrawToken) XXX_Size() int {
	return m.Size()
}
func (m *MsgWithdrawToken) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithdrawToken.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithdrawToken proto.InternalMessageInfo

func (m *MsgWithdrawToken) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgWithdrawToken) GetWithdraw() string {
	if m != nil {
		return m.Withdraw
	}
	return ""
}

type MsgWithdrawTokenResponse struct {
}

func (m *MsgWithdrawTokenResponse) Reset()         { *m = MsgWithdrawTokenResponse{} }
func (m *MsgWithdrawTokenResponse) String() string { return proto.CompactTextString(m) }
func (*MsgWithdrawTokenResponse) ProtoMessage()    {}
func (*MsgWithdrawTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1719cb30b3bc4bbd, []int{3}
}
func (m *MsgWithdrawTokenResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgWithdrawTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgWithdrawTokenResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgWithdrawTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgWithdrawTokenResponse.Merge(m, src)
}
func (m *MsgWithdrawTokenResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgWithdrawTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgWithdrawTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgWithdrawTokenResponse proto.InternalMessageInfo

type MsgAddLinkedRequester struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Network string `protobuf:"bytes,2,opt,name=network,proto3" json:"network,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *MsgAddLinkedRequester) Reset()         { *m = MsgAddLinkedRequester{} }
func (m *MsgAddLinkedRequester) String() string { return proto.CompactTextString(m) }
func (*MsgAddLinkedRequester) ProtoMessage()    {}
func (*MsgAddLinkedRequester) Descriptor() ([]byte, []int) {
	return fileDescriptor_1719cb30b3bc4bbd, []int{4}
}
func (m *MsgAddLinkedRequester) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddLinkedRequester) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddLinkedRequester.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddLinkedRequester) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddLinkedRequester.Merge(m, src)
}
func (m *MsgAddLinkedRequester) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddLinkedRequester) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddLinkedRequester.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddLinkedRequester proto.InternalMessageInfo

func (m *MsgAddLinkedRequester) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgAddLinkedRequester) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *MsgAddLinkedRequester) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type MsgAddLinkedRequesterResponse struct {
}

func (m *MsgAddLinkedRequesterResponse) Reset()         { *m = MsgAddLinkedRequesterResponse{} }
func (m *MsgAddLinkedRequesterResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAddLinkedRequesterResponse) ProtoMessage()    {}
func (*MsgAddLinkedRequesterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1719cb30b3bc4bbd, []int{5}
}
func (m *MsgAddLinkedRequesterResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddLinkedRequesterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddLinkedRequesterResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddLinkedRequesterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddLinkedRequesterResponse.Merge(m, src)
}
func (m *MsgAddLinkedRequesterResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddLinkedRequesterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddLinkedRequesterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddLinkedRequesterResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgDepositToken)(nil), "decent.user.MsgDepositToken")
	proto.RegisterType((*MsgDepositTokenResponse)(nil), "decent.user.MsgDepositTokenResponse")
	proto.RegisterType((*MsgWithdrawToken)(nil), "decent.user.MsgWithdrawToken")
	proto.RegisterType((*MsgWithdrawTokenResponse)(nil), "decent.user.MsgWithdrawTokenResponse")
	proto.RegisterType((*MsgAddLinkedRequester)(nil), "decent.user.MsgAddLinkedRequester")
	proto.RegisterType((*MsgAddLinkedRequesterResponse)(nil), "decent.user.MsgAddLinkedRequesterResponse")
}

func init() { proto.RegisterFile("decent/user/tx.proto", fileDescriptor_1719cb30b3bc4bbd) }

var fileDescriptor_1719cb30b3bc4bbd = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x6d, 0x5a, 0xf0, 0x63, 0x54, 0x94, 0xa8, 0x18, 0x83, 0x5d, 0x65, 0x51, 0x10, 0xc1, 0x14,
	0xf4, 0x17, 0x28, 0x15, 0x3c, 0x98, 0x4b, 0x50, 0x04, 0x6f, 0xb5, 0x3b, 0xc4, 0x52, 0xc8, 0xc6,
	0x9d, 0x2d, 0xad, 0xff, 0xc2, 0x93, 0xbf, 0xc9, 0x63, 0x8f, 0x1e, 0xa5, 0xf9, 0x23, 0x12, 0xb3,
	0x09, 0x26, 0x96, 0xf4, 0xf8, 0xf6, 0xbd, 0x79, 0x6f, 0xe7, 0x31, 0xb0, 0x23, 0xb0, 0x8f, 0x91,
	0xee, 0x8c, 0x08, 0x55, 0x47, 0x4f, 0xbc, 0x58, 0x49, 0x2d, 0xed, 0xb5, 0xec, 0xd5, 0x4b, 0x5f,
	0xf9, 0x0d, 0x6c, 0xfa, 0x14, 0x76, 0x31, 0x96, 0x34, 0xd0, 0xf7, 0x72, 0x88, 0x91, 0xed, 0xc0,
	0x72, 0x5f, 0x61, 0x4f, 0x4b, 0xe5, 0x58, 0x47, 0xd6, 0xe9, 0x6a, 0x90, 0xc3, 0x94, 0x11, 0x99,
	0xd2, 0x69, 0x66, 0x8c, 0x81, 0x7c, 0x1f, 0xf6, 0x2a, 0x36, 0x01, 0x52, 0x2c, 0x23, 0x42, 0x7e,
	0x0b, 0x5b, 0x3e, 0x85, 0x8f, 0x03, 0xfd, 0x22, 0x54, 0x6f, 0xbc, 0x28, 0xc2, 0x85, 0x95, 0xb1,
	0x91, 0x9a, 0x8c, 0x02, 0x73, 0x17, 0x9c, 0xaa, 0x53, 0x91, 0x82, 0xb0, 0xeb, 0x53, 0x78, 0x25,
	0xc4, 0xdd, 0x20, 0x1a, 0xa2, 0x08, 0xf0, 0x75, 0x84, 0xa4, 0x51, 0xd5, 0x6f, 0x13, 0xa1, 0x1e,
	0x4b, 0x35, 0xcc, 0xb7, 0x31, 0x30, 0x65, 0x7a, 0x42, 0x28, 0x24, 0x72, 0x5a, 0x19, 0x63, 0x20,
	0x3f, 0x84, 0xf6, 0xdc, 0x98, 0xfc, 0x1f, 0x17, 0x1f, 0x4d, 0x68, 0xf9, 0x14, 0xda, 0x01, 0xac,
	0x97, 0x4a, 0x3d, 0xf0, 0xfe, 0xb4, 0xee, 0x55, 0xba, 0x72, 0x8f, 0xeb, 0xd8, 0xdc, 0xdb, 0x7e,
	0x80, 0x8d, 0x72, 0x8d, 0xed, 0xea, 0x58, 0x89, 0x76, 0x4f, 0x6a, 0xe9, 0xc2, 0x56, 0x80, 0x3d,
	0xa7, 0x37, 0x5e, 0x1d, 0xfe, 0xaf, 0x71, 0xcf, 0x16, 0x6b, 0xf2, 0x94, 0xeb, 0xf3, 0xcf, 0x19,
	0xb3, 0xa6, 0x33, 0x66, 0x7d, 0xcf, 0x98, 0xf5, 0x9e, 0xb0, 0xc6, 0x34, 0x61, 0x8d, 0xaf, 0x84,
	0x35, 0x9e, 0xb6, 0xbb, 0xd9, 0x95, 0x4e, 0xcc, 0x9d, 0xbe, 0xc5, 0x48, 0xcf, 0x4b, 0xbf, 0xb7,
	0x7a, 0xf9, 0x13, 0x00, 0x00, 0xff, 0xff, 0x0d, 0xfb, 0x85, 0x2a, 0xc3, 0x02, 0x00, 0x00,
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
	DepositToken(ctx context.Context, in *MsgDepositToken, opts ...grpc.CallOption) (*MsgDepositTokenResponse, error)
	WithdrawToken(ctx context.Context, in *MsgWithdrawToken, opts ...grpc.CallOption) (*MsgWithdrawTokenResponse, error)
	AddLinkedRequester(ctx context.Context, in *MsgAddLinkedRequester, opts ...grpc.CallOption) (*MsgAddLinkedRequesterResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) DepositToken(ctx context.Context, in *MsgDepositToken, opts ...grpc.CallOption) (*MsgDepositTokenResponse, error) {
	out := new(MsgDepositTokenResponse)
	err := c.cc.Invoke(ctx, "/decent.user.Msg/DepositToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) WithdrawToken(ctx context.Context, in *MsgWithdrawToken, opts ...grpc.CallOption) (*MsgWithdrawTokenResponse, error) {
	out := new(MsgWithdrawTokenResponse)
	err := c.cc.Invoke(ctx, "/decent.user.Msg/WithdrawToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AddLinkedRequester(ctx context.Context, in *MsgAddLinkedRequester, opts ...grpc.CallOption) (*MsgAddLinkedRequesterResponse, error) {
	out := new(MsgAddLinkedRequesterResponse)
	err := c.cc.Invoke(ctx, "/decent.user.Msg/AddLinkedRequester", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	DepositToken(context.Context, *MsgDepositToken) (*MsgDepositTokenResponse, error)
	WithdrawToken(context.Context, *MsgWithdrawToken) (*MsgWithdrawTokenResponse, error)
	AddLinkedRequester(context.Context, *MsgAddLinkedRequester) (*MsgAddLinkedRequesterResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) DepositToken(ctx context.Context, req *MsgDepositToken) (*MsgDepositTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DepositToken not implemented")
}
func (*UnimplementedMsgServer) WithdrawToken(ctx context.Context, req *MsgWithdrawToken) (*MsgWithdrawTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithdrawToken not implemented")
}
func (*UnimplementedMsgServer) AddLinkedRequester(ctx context.Context, req *MsgAddLinkedRequester) (*MsgAddLinkedRequesterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddLinkedRequester not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_DepositToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDepositToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DepositToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/decent.user.Msg/DepositToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DepositToken(ctx, req.(*MsgDepositToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_WithdrawToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgWithdrawToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).WithdrawToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/decent.user.Msg/WithdrawToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).WithdrawToken(ctx, req.(*MsgWithdrawToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AddLinkedRequester_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddLinkedRequester)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddLinkedRequester(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/decent.user.Msg/AddLinkedRequester",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddLinkedRequester(ctx, req.(*MsgAddLinkedRequester))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "decent.user.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DepositToken",
			Handler:    _Msg_DepositToken_Handler,
		},
		{
			MethodName: "WithdrawToken",
			Handler:    _Msg_WithdrawToken_Handler,
		},
		{
			MethodName: "AddLinkedRequester",
			Handler:    _Msg_AddLinkedRequester_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "decent/user/tx.proto",
}

func (m *MsgDepositToken) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDepositToken) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDepositToken) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Deposit) > 0 {
		i -= len(m.Deposit)
		copy(dAtA[i:], m.Deposit)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Deposit)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDepositTokenResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDepositTokenResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDepositTokenResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgWithdrawToken) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWithdrawToken) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWithdrawToken) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Withdraw) > 0 {
		i -= len(m.Withdraw)
		copy(dAtA[i:], m.Withdraw)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Withdraw)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgWithdrawTokenResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgWithdrawTokenResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgWithdrawTokenResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgAddLinkedRequester) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddLinkedRequester) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddLinkedRequester) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Network) > 0 {
		i -= len(m.Network)
		copy(dAtA[i:], m.Network)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Network)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAddLinkedRequesterResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddLinkedRequesterResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddLinkedRequesterResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgDepositToken) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Deposit)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgDepositTokenResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgWithdrawToken) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Withdraw)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgWithdrawTokenResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgAddLinkedRequester) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Network)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgAddLinkedRequesterResponse) Size() (n int) {
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
func (m *MsgDepositToken) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgDepositToken: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDepositToken: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposit", wireType)
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
			m.Deposit = string(dAtA[iNdEx:postIndex])
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
func (m *MsgDepositTokenResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgDepositTokenResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDepositTokenResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgWithdrawToken) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgWithdrawToken: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWithdrawToken: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Withdraw", wireType)
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
			m.Withdraw = string(dAtA[iNdEx:postIndex])
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
func (m *MsgWithdrawTokenResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgWithdrawTokenResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgWithdrawTokenResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgAddLinkedRequester) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgAddLinkedRequester: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddLinkedRequester: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Network", wireType)
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
			m.Network = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
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
func (m *MsgAddLinkedRequesterResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgAddLinkedRequesterResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddLinkedRequesterResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
