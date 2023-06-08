// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dhpc/data/data.proto

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

type Data struct {
	Address       string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Owner         string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Network       string `protobuf:"bytes,3,opt,name=network,proto3" json:"network,omitempty"`
	DateAdded     uint64 `protobuf:"varint,4,opt,name=dateAdded,proto3" json:"dateAdded,omitempty"`
	DateUpdated   uint64 `protobuf:"varint,5,opt,name=dateUpdated,proto3" json:"dateUpdated,omitempty"`
	BlockValidity uint64 `protobuf:"varint,6,opt,name=blockValidity,proto3" json:"blockValidity,omitempty"`
	Event         string `protobuf:"bytes,7,opt,name=event,proto3" json:"event,omitempty"`
	Score         int32  `protobuf:"varint,8,opt,name=score,proto3" json:"score,omitempty"`
	Hit           uint64 `protobuf:"varint,9,opt,name=hit,proto3" json:"hit,omitempty"`
	Hash          string `protobuf:"bytes,10,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_84c539e64da5d140, []int{0}
}
func (m *Data) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Data.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return m.Size()
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Data) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Data) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *Data) GetDateAdded() uint64 {
	if m != nil {
		return m.DateAdded
	}
	return 0
}

func (m *Data) GetDateUpdated() uint64 {
	if m != nil {
		return m.DateUpdated
	}
	return 0
}

func (m *Data) GetBlockValidity() uint64 {
	if m != nil {
		return m.BlockValidity
	}
	return 0
}

func (m *Data) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *Data) GetScore() int32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *Data) GetHit() uint64 {
	if m != nil {
		return m.Hit
	}
	return 0
}

func (m *Data) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*Data)(nil), "dhpc.data.Data")
}

func init() { proto.RegisterFile("dhpc/data/data.proto", fileDescriptor_84c539e64da5d140) }

var fileDescriptor_84c539e64da5d140 = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xb3, 0x6d, 0xd2, 0x9a, 0x11, 0x41, 0x97, 0x1e, 0xe6, 0x20, 0x4b, 0x10, 0x0f, 0x01,
	0xa1, 0x1e, 0x7c, 0x02, 0xa5, 0x4f, 0x10, 0xd0, 0x83, 0xb7, 0x6d, 0x66, 0x21, 0xa1, 0x25, 0x1b,
	0x76, 0x17, 0x6b, 0xdf, 0xc0, 0xa3, 0x8f, 0xe5, 0xb1, 0x47, 0x8f, 0x92, 0xbc, 0x88, 0xec, 0xae,
	0xa2, 0xbd, 0x0c, 0xf3, 0x7d, 0xf9, 0xb3, 0x0c, 0x3f, 0x2c, 0xa8, 0xe9, 0xeb, 0x5b, 0x92, 0x4e,
	0x86, 0xb1, 0xec, 0x8d, 0x76, 0x9a, 0xe7, 0xde, 0x2e, 0xbd, 0xb8, 0x7a, 0x9b, 0x40, 0xba, 0x92,
	0x4e, 0x72, 0x84, 0xb9, 0x24, 0x32, 0xca, 0x5a, 0x64, 0x05, 0x2b, 0xf3, 0xea, 0x17, 0xf9, 0x02,
	0x32, 0xbd, 0xeb, 0x94, 0xc1, 0x49, 0xf0, 0x11, 0x7c, 0xbe, 0x53, 0x6e, 0xa7, 0xcd, 0x06, 0xa7,
	0x31, 0xff, 0x83, 0xfc, 0x12, 0x72, 0x92, 0x4e, 0xdd, 0x13, 0x29, 0xc2, 0xb4, 0x60, 0x65, 0x5a,
	0xfd, 0x09, 0x5e, 0xc0, 0xa9, 0x87, 0xc7, 0xde, 0x4f, 0xc2, 0x2c, 0x7c, 0xff, 0xaf, 0xf8, 0x35,
	0x9c, 0xad, 0xb7, 0xba, 0xde, 0x3c, 0xc9, 0x6d, 0x4b, 0xad, 0xdb, 0xe3, 0x2c, 0x64, 0x8e, 0xa5,
	0xbf, 0x4a, 0xbd, 0xa8, 0xce, 0xe1, 0x3c, 0x5e, 0x15, 0xc0, 0x5b, 0x5b, 0x6b, 0xa3, 0xf0, 0xa4,
	0x60, 0x65, 0x56, 0x45, 0xe0, 0xe7, 0x30, 0x6d, 0x5a, 0x87, 0x79, 0x78, 0xc7, 0xaf, 0x9c, 0x43,
	0xda, 0x48, 0xdb, 0x20, 0x84, 0x9f, 0xc3, 0xfe, 0x70, 0xf3, 0x31, 0x08, 0x76, 0x18, 0x04, 0xfb,
	0x1a, 0x04, 0x7b, 0x1f, 0x45, 0x72, 0x18, 0x45, 0xf2, 0x39, 0x8a, 0xe4, 0xf9, 0x62, 0xe5, 0x5b,
	0x7c, 0x8d, 0x3d, 0xba, 0x7d, 0xaf, 0xec, 0x7a, 0x16, 0x9a, 0xbc, 0xfb, 0x0e, 0x00, 0x00, 0xff,
	0xff, 0x93, 0x3b, 0x7b, 0x5f, 0x61, 0x01, 0x00, 0x00,
}

func (m *Data) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Data) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Data) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintData(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x52
	}
	if m.Hit != 0 {
		i = encodeVarintData(dAtA, i, uint64(m.Hit))
		i--
		dAtA[i] = 0x48
	}
	if m.Score != 0 {
		i = encodeVarintData(dAtA, i, uint64(m.Score))
		i--
		dAtA[i] = 0x40
	}
	if len(m.Event) > 0 {
		i -= len(m.Event)
		copy(dAtA[i:], m.Event)
		i = encodeVarintData(dAtA, i, uint64(len(m.Event)))
		i--
		dAtA[i] = 0x3a
	}
	if m.BlockValidity != 0 {
		i = encodeVarintData(dAtA, i, uint64(m.BlockValidity))
		i--
		dAtA[i] = 0x30
	}
	if m.DateUpdated != 0 {
		i = encodeVarintData(dAtA, i, uint64(m.DateUpdated))
		i--
		dAtA[i] = 0x28
	}
	if m.DateAdded != 0 {
		i = encodeVarintData(dAtA, i, uint64(m.DateAdded))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Network) > 0 {
		i -= len(m.Network)
		copy(dAtA[i:], m.Network)
		i = encodeVarintData(dAtA, i, uint64(len(m.Network)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintData(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintData(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintData(dAtA []byte, offset int, v uint64) int {
	offset -= sovData(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Data) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovData(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovData(uint64(l))
	}
	l = len(m.Network)
	if l > 0 {
		n += 1 + l + sovData(uint64(l))
	}
	if m.DateAdded != 0 {
		n += 1 + sovData(uint64(m.DateAdded))
	}
	if m.DateUpdated != 0 {
		n += 1 + sovData(uint64(m.DateUpdated))
	}
	if m.BlockValidity != 0 {
		n += 1 + sovData(uint64(m.BlockValidity))
	}
	l = len(m.Event)
	if l > 0 {
		n += 1 + l + sovData(uint64(l))
	}
	if m.Score != 0 {
		n += 1 + sovData(uint64(m.Score))
	}
	if m.Hit != 0 {
		n += 1 + sovData(uint64(m.Hit))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovData(uint64(l))
	}
	return n
}

func sovData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozData(x uint64) (n int) {
	return sovData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Data) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowData
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
			return fmt.Errorf("proto: Data: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Data: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
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
				return ErrInvalidLengthData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
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
				return ErrInvalidLengthData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Network", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
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
				return ErrInvalidLengthData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Network = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DateAdded", wireType)
			}
			m.DateAdded = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DateAdded |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DateUpdated", wireType)
			}
			m.DateUpdated = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DateUpdated |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockValidity", wireType)
			}
			m.BlockValidity = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockValidity |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Event", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
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
				return ErrInvalidLengthData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Event = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Score", wireType)
			}
			m.Score = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
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
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hit", wireType)
			}
			m.Hit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Hit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowData
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
				return ErrInvalidLengthData
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthData
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
func skipData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowData
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
					return 0, ErrIntOverflowData
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
					return 0, ErrIntOverflowData
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
				return 0, ErrInvalidLengthData
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupData
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthData
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthData        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowData          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupData = fmt.Errorf("proto: unexpected end of group")
)
