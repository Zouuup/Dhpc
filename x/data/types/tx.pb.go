// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dhpc/data/tx.proto

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

type MsgAddData struct {
	Creator       string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Address       string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Owner         string `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	Network       string `protobuf:"bytes,4,opt,name=network,proto3" json:"network,omitempty"`
	Event         string `protobuf:"bytes,5,opt,name=event,proto3" json:"event,omitempty"`
	Blockvalidity uint64 `protobuf:"varint,6,opt,name=blockvalidity,proto3" json:"blockvalidity,omitempty"`
	Score         int32  `protobuf:"varint,7,opt,name=score,proto3" json:"score,omitempty"`
	Hash          string `protobuf:"bytes,8,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *MsgAddData) Reset()         { *m = MsgAddData{} }
func (m *MsgAddData) String() string { return proto.CompactTextString(m) }
func (*MsgAddData) ProtoMessage()    {}
func (*MsgAddData) Descriptor() ([]byte, []int) {
	return fileDescriptor_acbdb61c8414c235, []int{0}
}
func (m *MsgAddData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddData.Merge(m, src)
}
func (m *MsgAddData) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddData) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddData.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddData proto.InternalMessageInfo

func (m *MsgAddData) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgAddData) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgAddData) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *MsgAddData) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *MsgAddData) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *MsgAddData) GetBlockvalidity() uint64 {
	if m != nil {
		return m.Blockvalidity
	}
	return 0
}

func (m *MsgAddData) GetScore() int32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *MsgAddData) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type MsgAddDataResponse struct {
}

func (m *MsgAddDataResponse) Reset()         { *m = MsgAddDataResponse{} }
func (m *MsgAddDataResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAddDataResponse) ProtoMessage()    {}
func (*MsgAddDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_acbdb61c8414c235, []int{1}
}
func (m *MsgAddDataResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddDataResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddDataResponse.Merge(m, src)
}
func (m *MsgAddDataResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddDataResponse proto.InternalMessageInfo

type MsgDeleteData struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Owner   string `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	Network string `protobuf:"bytes,4,opt,name=network,proto3" json:"network,omitempty"`
}

func (m *MsgDeleteData) Reset()         { *m = MsgDeleteData{} }
func (m *MsgDeleteData) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteData) ProtoMessage()    {}
func (*MsgDeleteData) Descriptor() ([]byte, []int) {
	return fileDescriptor_acbdb61c8414c235, []int{2}
}
func (m *MsgDeleteData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteData.Merge(m, src)
}
func (m *MsgDeleteData) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteData) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteData.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteData proto.InternalMessageInfo

func (m *MsgDeleteData) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgDeleteData) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgDeleteData) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *MsgDeleteData) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

type MsgDeleteDataResponse struct {
}

func (m *MsgDeleteDataResponse) Reset()         { *m = MsgDeleteDataResponse{} }
func (m *MsgDeleteDataResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteDataResponse) ProtoMessage()    {}
func (*MsgDeleteDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_acbdb61c8414c235, []int{3}
}
func (m *MsgDeleteDataResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteDataResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteDataResponse.Merge(m, src)
}
func (m *MsgDeleteDataResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteDataResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgAddData)(nil), "dhpc.data.MsgAddData")
	proto.RegisterType((*MsgAddDataResponse)(nil), "dhpc.data.MsgAddDataResponse")
	proto.RegisterType((*MsgDeleteData)(nil), "dhpc.data.MsgDeleteData")
	proto.RegisterType((*MsgDeleteDataResponse)(nil), "dhpc.data.MsgDeleteDataResponse")
}

func init() { proto.RegisterFile("dhpc/data/tx.proto", fileDescriptor_acbdb61c8414c235) }

var fileDescriptor_acbdb61c8414c235 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x52, 0xbb, 0x6a, 0xf3, 0x30,
	0x14, 0x8e, 0xfe, 0xdc, 0xfe, 0x1c, 0xc8, 0x22, 0x12, 0x2a, 0x02, 0x35, 0xc6, 0x94, 0x92, 0xc9,
	0x86, 0xf6, 0x01, 0x7a, 0xcb, 0xd0, 0x25, 0x8b, 0xc7, 0x6e, 0x8a, 0x75, 0xb0, 0x4d, 0x52, 0xcb,
	0x48, 0x6a, 0x2e, 0x6f, 0x91, 0xc7, 0xea, 0x98, 0xa9, 0x74, 0x2c, 0xc9, 0x8b, 0x14, 0xcb, 0x71,
	0x53, 0x43, 0xbb, 0x76, 0x3b, 0xdf, 0x45, 0x47, 0xe7, 0xe3, 0x1c, 0xa0, 0x22, 0xc9, 0xa3, 0x40,
	0x70, 0xc3, 0x03, 0xb3, 0xf6, 0x73, 0x25, 0x8d, 0xa4, 0xbd, 0x82, 0xf3, 0x0b, 0xce, 0x7b, 0x23,
	0x00, 0x53, 0x1d, 0xdf, 0x09, 0x31, 0xe1, 0x86, 0x53, 0x06, 0xdd, 0x48, 0x21, 0x37, 0x52, 0x31,
	0xe2, 0x92, 0x71, 0x2f, 0xac, 0x60, 0xa1, 0x70, 0x21, 0x14, 0x6a, 0xcd, 0xfe, 0x95, 0xca, 0x11,
	0xd2, 0x01, 0xb4, 0xe5, 0x2a, 0x43, 0xc5, 0x9a, 0x96, 0x2f, 0x41, 0xe1, 0xcf, 0xd0, 0xac, 0xa4,
	0x9a, 0xb3, 0x56, 0xe9, 0x3f, 0xc2, 0xc2, 0x8f, 0x4b, 0xcc, 0x0c, 0x6b, 0x97, 0x7e, 0x0b, 0xe8,
	0x05, 0xf4, 0x67, 0x0b, 0x19, 0xcd, 0x97, 0x7c, 0x91, 0x8a, 0xd4, 0x6c, 0x58, 0xc7, 0x25, 0xe3,
	0x56, 0x58, 0x27, 0x8b, 0xb7, 0x3a, 0x92, 0x0a, 0x59, 0xd7, 0x25, 0xe3, 0x76, 0x58, 0x02, 0x4a,
	0xa1, 0x95, 0x70, 0x9d, 0xb0, 0xff, 0xb6, 0xa1, 0xad, 0xbd, 0x01, 0xd0, 0x53, 0xae, 0x10, 0x75,
	0x2e, 0x33, 0x8d, 0x9e, 0x86, 0xfe, 0x54, 0xc7, 0x13, 0x5c, 0xa0, 0xc1, 0xbf, 0x0a, 0xec, 0x9d,
	0xc1, 0xb0, 0xf6, 0x69, 0x35, 0xcd, 0xd5, 0x96, 0x40, 0x73, 0xaa, 0x63, 0x7a, 0x03, 0xdd, 0x6a,
	0x01, 0x43, 0xff, 0x6b, 0x37, 0xfe, 0x69, 0xfe, 0xd1, 0xf9, 0x8f, 0x74, 0xd5, 0x88, 0x3e, 0x02,
	0x7c, 0xcf, 0x54, 0x37, 0x9f, 0x94, 0x91, 0xfb, 0x9b, 0x52, 0x75, 0xba, 0xbf, 0x7d, 0xdd, 0x3b,
	0x64, 0xb7, 0x77, 0xc8, 0xc7, 0xde, 0x21, 0xdb, 0x83, 0xd3, 0xd8, 0x1d, 0x9c, 0xc6, 0xfb, 0xc1,
	0x69, 0x3c, 0x5d, 0xc6, 0xa9, 0x49, 0x5e, 0x66, 0x7e, 0x24, 0x9f, 0x83, 0x49, 0x92, 0x47, 0x0f,
	0x09, 0x4f, 0x33, 0x5b, 0x05, 0xeb, 0xe3, 0x7d, 0x6d, 0x72, 0xd4, 0xb3, 0x8e, 0xbd, 0xb1, 0xeb,
	0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x66, 0x44, 0x47, 0x21, 0x79, 0x02, 0x00, 0x00,
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
	AddData(ctx context.Context, in *MsgAddData, opts ...grpc.CallOption) (*MsgAddDataResponse, error)
	DeleteData(ctx context.Context, in *MsgDeleteData, opts ...grpc.CallOption) (*MsgDeleteDataResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) AddData(ctx context.Context, in *MsgAddData, opts ...grpc.CallOption) (*MsgAddDataResponse, error) {
	out := new(MsgAddDataResponse)
	err := c.cc.Invoke(ctx, "/dhpc.data.Msg/AddData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteData(ctx context.Context, in *MsgDeleteData, opts ...grpc.CallOption) (*MsgDeleteDataResponse, error) {
	out := new(MsgDeleteDataResponse)
	err := c.cc.Invoke(ctx, "/dhpc.data.Msg/DeleteData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	AddData(context.Context, *MsgAddData) (*MsgAddDataResponse, error)
	DeleteData(context.Context, *MsgDeleteData) (*MsgDeleteDataResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) AddData(ctx context.Context, req *MsgAddData) (*MsgAddDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddData not implemented")
}
func (*UnimplementedMsgServer) DeleteData(ctx context.Context, req *MsgDeleteData) (*MsgDeleteDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteData not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_AddData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dhpc.data.Msg/AddData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddData(ctx, req.(*MsgAddData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dhpc.data.Msg/DeleteData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteData(ctx, req.(*MsgDeleteData))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dhpc.data.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddData",
			Handler:    _Msg_AddData_Handler,
		},
		{
			MethodName: "DeleteData",
			Handler:    _Msg_DeleteData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dhpc/data/tx.proto",
}

func (m *MsgAddData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x42
	}
	if m.Score != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Score))
		i--
		dAtA[i] = 0x38
	}
	if m.Blockvalidity != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Blockvalidity))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Event) > 0 {
		i -= len(m.Event)
		copy(dAtA[i:], m.Event)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Event)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Network) > 0 {
		i -= len(m.Network)
		copy(dAtA[i:], m.Network)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Network)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
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

func (m *MsgAddDataResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddDataResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddDataResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgDeleteData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Network) > 0 {
		i -= len(m.Network)
		copy(dAtA[i:], m.Network)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Network)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
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

func (m *MsgDeleteDataResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteDataResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteDataResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgAddData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Network)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Event)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Blockvalidity != 0 {
		n += 1 + sovTx(uint64(m.Blockvalidity))
	}
	if m.Score != 0 {
		n += 1 + sovTx(uint64(m.Score))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgAddDataResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgDeleteData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Network)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgDeleteDataResponse) Size() (n int) {
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
func (m *MsgAddData) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgAddData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddData: illegal tag %d (wire type %d)", fieldNum, wire)
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
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
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
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
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Event", wireType)
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
			m.Event = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blockvalidity", wireType)
			}
			m.Blockvalidity = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Blockvalidity |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Score", wireType)
			}
			m.Score = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Score |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
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
			m.Hash = string(dAtA[iNdEx:postIndex])
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
func (m *MsgAddDataResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgAddDataResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddDataResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgDeleteData) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgDeleteData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteData: illegal tag %d (wire type %d)", fieldNum, wire)
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
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
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
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
func (m *MsgDeleteDataResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgDeleteDataResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteDataResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
