// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: decent/request/request_record.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type RequestRecord struct {
	Index   string           `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	UUID    string           `protobuf:"bytes,2,opt,name=uUID,proto3" json:"uUID,omitempty"`
	TXhash  string           `protobuf:"bytes,3,opt,name=tXhash,proto3" json:"tXhash,omitempty"`
	Network string           `protobuf:"bytes,4,opt,name=network,proto3" json:"network,omitempty"`
	From    string           `protobuf:"bytes,5,opt,name=from,proto3" json:"from,omitempty"`
	Address string           `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	Score   int32            `protobuf:"varint,7,opt,name=score,proto3" json:"score,omitempty"`
	Oracle  string           `protobuf:"bytes,8,opt,name=oracle,proto3" json:"oracle,omitempty"`
	Block   int32            `protobuf:"varint,9,opt,name=block,proto3" json:"block,omitempty"`
	Stage   int32            `protobuf:"varint,10,opt,name=stage,proto3" json:"stage,omitempty"`
	Miners  []*MinerResponse `protobuf:"bytes,11,rep,name=miners,proto3" json:"miners,omitempty"`
	Creator string           `protobuf:"bytes,12,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *RequestRecord) Reset()         { *m = RequestRecord{} }
func (m *RequestRecord) String() string { return proto.CompactTextString(m) }
func (*RequestRecord) ProtoMessage()    {}
func (*RequestRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_5dd4093d0ee40417, []int{0}
}
func (m *RequestRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestRecord.Merge(m, src)
}
func (m *RequestRecord) XXX_Size() int {
	return m.Size()
}
func (m *RequestRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestRecord.DiscardUnknown(m)
}

var xxx_messageInfo_RequestRecord proto.InternalMessageInfo

func (m *RequestRecord) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *RequestRecord) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *RequestRecord) GetTXhash() string {
	if m != nil {
		return m.TXhash
	}
	return ""
}

func (m *RequestRecord) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *RequestRecord) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *RequestRecord) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *RequestRecord) GetScore() int32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *RequestRecord) GetOracle() string {
	if m != nil {
		return m.Oracle
	}
	return ""
}

func (m *RequestRecord) GetBlock() int32 {
	if m != nil {
		return m.Block
	}
	return 0
}

func (m *RequestRecord) GetStage() int32 {
	if m != nil {
		return m.Stage
	}
	return 0
}

func (m *RequestRecord) GetMiners() []*MinerResponse {
	if m != nil {
		return m.Miners
	}
	return nil
}

func (m *RequestRecord) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*RequestRecord)(nil), "decent.request.RequestRecord")
}

func init() {
	proto.RegisterFile("decent/request/request_record.proto", fileDescriptor_5dd4093d0ee40417)
}

var fileDescriptor_5dd4093d0ee40417 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0xc6, 0xbb, 0xfd, 0xb3, 0xb5, 0xa9, 0x7a, 0x08, 0x52, 0x06, 0xc1, 0xa5, 0xe8, 0xa5, 0xa7,
	0x55, 0x14, 0x5f, 0x40, 0x7a, 0xf1, 0xe0, 0x65, 0x41, 0x10, 0x2f, 0x65, 0x9b, 0x1d, 0x6d, 0x69,
	0xbb, 0xa9, 0x93, 0x14, 0xeb, 0x5b, 0xf8, 0x50, 0x1e, 0x3c, 0xf6, 0xe8, 0x51, 0xba, 0x2f, 0x22,
	0xc9, 0xc4, 0x82, 0x9e, 0x92, 0x5f, 0xe6, 0xcb, 0x37, 0xc3, 0x37, 0xe2, 0xac, 0x40, 0x85, 0xa5,
	0x3d, 0x27, 0x7c, 0x59, 0xa1, 0xd9, 0x9d, 0x23, 0x42, 0xa5, 0xa9, 0x48, 0x97, 0xa4, 0xad, 0x96,
	0x87, 0x2c, 0x4a, 0x43, 0xf1, 0xf8, 0xff, 0xa7, 0xc5, 0xb4, 0x44, 0x1a, 0x11, 0x9a, 0xa5, 0x2e,
	0x0d, 0xf2, 0xa7, 0xd3, 0x8f, 0xba, 0x38, 0xc8, 0x58, 0x90, 0x79, 0x33, 0x79, 0x24, 0x5a, 0xd3,
	0xb2, 0xc0, 0x35, 0x44, 0xfd, 0x68, 0xd0, 0xc9, 0x18, 0xa4, 0x14, 0xcd, 0xd5, 0xfd, 0xed, 0x10,
	0xea, 0xfe, 0xd1, 0xdf, 0x65, 0x4f, 0xc4, 0xf6, 0x61, 0x92, 0x9b, 0x09, 0x34, 0xfc, 0x6b, 0x20,
	0x09, 0xa2, 0x5d, 0xa2, 0x7d, 0xd5, 0x34, 0x83, 0xa6, 0x2f, 0xfc, 0xa2, 0x73, 0x79, 0x22, 0xbd,
	0x80, 0x16, 0xbb, 0xb8, 0xbb, 0x53, 0xe7, 0x45, 0x41, 0x68, 0x0c, 0xc4, 0xac, 0x0e, 0xe8, 0x26,
	0x31, 0x4a, 0x13, 0x42, 0xbb, 0x1f, 0x0d, 0x5a, 0x19, 0x83, 0xeb, 0xaa, 0x29, 0x57, 0x73, 0x84,
	0x3d, 0xee, 0xca, 0xe4, 0xd4, 0xe3, 0xb9, 0x56, 0x33, 0xe8, 0xb0, 0xda, 0x83, 0xf7, 0xb0, 0xf9,
	0x33, 0x82, 0x08, 0x1e, 0x0e, 0xe4, 0xb5, 0x88, 0x7d, 0x1a, 0x06, 0xba, 0xfd, 0xc6, 0xa0, 0x7b,
	0x79, 0x92, 0xfe, 0xcd, 0x2e, 0xbd, 0x73, 0xd5, 0x2c, 0x44, 0x95, 0x05, 0xb1, 0x1b, 0x55, 0x11,
	0xe6, 0x56, 0x13, 0xec, 0xf3, 0xa8, 0x01, 0x6f, 0x2e, 0x3e, 0xb7, 0x49, 0xb4, 0xd9, 0x26, 0xd1,
	0xf7, 0x36, 0x89, 0xde, 0xab, 0xa4, 0xb6, 0xa9, 0x92, 0xda, 0x57, 0x95, 0xd4, 0x1e, 0x7b, 0x43,
	0xde, 0xc2, 0x7a, 0xb7, 0x07, 0xfb, 0xb6, 0x44, 0x33, 0x8e, 0x7d, 0xfe, 0x57, 0x3f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xb8, 0xee, 0x1b, 0x14, 0xdb, 0x01, 0x00, 0x00,
}

func (m *RequestRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintRequestRecord(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.Miners) > 0 {
		for iNdEx := len(m.Miners) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Miners[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRequestRecord(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x5a
		}
	}
	if m.Stage != 0 {
		i = encodeVarintRequestRecord(dAtA, i, uint64(m.Stage))
		i--
		dAtA[i] = 0x50
	}
	if m.Block != 0 {
		i = encodeVarintRequestRecord(dAtA, i, uint64(m.Block))
		i--
		dAtA[i] = 0x48
	}
	if len(m.Oracle) > 0 {
		i -= len(m.Oracle)
		copy(dAtA[i:], m.Oracle)
		i = encodeVarintRequestRecord(dAtA, i, uint64(len(m.Oracle)))
		i--
		dAtA[i] = 0x42
	}
	if m.Score != 0 {
		i = encodeVarintRequestRecord(dAtA, i, uint64(m.Score))
		i--
		dAtA[i] = 0x38
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintRequestRecord(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintRequestRecord(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Network) > 0 {
		i -= len(m.Network)
		copy(dAtA[i:], m.Network)
		i = encodeVarintRequestRecord(dAtA, i, uint64(len(m.Network)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.TXhash) > 0 {
		i -= len(m.TXhash)
		copy(dAtA[i:], m.TXhash)
		i = encodeVarintRequestRecord(dAtA, i, uint64(len(m.TXhash)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.UUID) > 0 {
		i -= len(m.UUID)
		copy(dAtA[i:], m.UUID)
		i = encodeVarintRequestRecord(dAtA, i, uint64(len(m.UUID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintRequestRecord(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRequestRecord(dAtA []byte, offset int, v uint64) int {
	offset -= sovRequestRecord(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RequestRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovRequestRecord(uint64(l))
	}
	l = len(m.UUID)
	if l > 0 {
		n += 1 + l + sovRequestRecord(uint64(l))
	}
	l = len(m.TXhash)
	if l > 0 {
		n += 1 + l + sovRequestRecord(uint64(l))
	}
	l = len(m.Network)
	if l > 0 {
		n += 1 + l + sovRequestRecord(uint64(l))
	}
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovRequestRecord(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovRequestRecord(uint64(l))
	}
	if m.Score != 0 {
		n += 1 + sovRequestRecord(uint64(m.Score))
	}
	l = len(m.Oracle)
	if l > 0 {
		n += 1 + l + sovRequestRecord(uint64(l))
	}
	if m.Block != 0 {
		n += 1 + sovRequestRecord(uint64(m.Block))
	}
	if m.Stage != 0 {
		n += 1 + sovRequestRecord(uint64(m.Stage))
	}
	if len(m.Miners) > 0 {
		for _, e := range m.Miners {
			l = e.Size()
			n += 1 + l + sovRequestRecord(uint64(l))
		}
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovRequestRecord(uint64(l))
	}
	return n
}

func sovRequestRecord(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRequestRecord(x uint64) (n int) {
	return sovRequestRecord(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RequestRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRequestRecord
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
			return fmt.Errorf("proto: RequestRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestRecord
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
				return ErrInvalidLengthRequestRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequestRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestRecord
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
				return ErrInvalidLengthRequestRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequestRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TXhash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestRecord
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
				return ErrInvalidLengthRequestRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequestRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TXhash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Network", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestRecord
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
				return ErrInvalidLengthRequestRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequestRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Network = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestRecord
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
				return ErrInvalidLengthRequestRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequestRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestRecord
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
				return ErrInvalidLengthRequestRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequestRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Score", wireType)
			}
			m.Score = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestRecord
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
				return fmt.Errorf("proto: wrong wireType = %d for field Oracle", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestRecord
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
				return ErrInvalidLengthRequestRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequestRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Oracle = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			m.Block = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Block |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stage", wireType)
			}
			m.Stage = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestRecord
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Stage |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Miners", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestRecord
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
				return ErrInvalidLengthRequestRecord
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRequestRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Miners = append(m.Miners, &MinerResponse{})
			if err := m.Miners[len(m.Miners)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequestRecord
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
				return ErrInvalidLengthRequestRecord
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequestRecord
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRequestRecord(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRequestRecord
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
func skipRequestRecord(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRequestRecord
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
					return 0, ErrIntOverflowRequestRecord
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
					return 0, ErrIntOverflowRequestRecord
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
				return 0, ErrInvalidLengthRequestRecord
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRequestRecord
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRequestRecord
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRequestRecord        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRequestRecord          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRequestRecord = fmt.Errorf("proto: unexpected end of group")
)
