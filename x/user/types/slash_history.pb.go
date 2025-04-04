// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dhpc/user/slash_history.proto

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

type SlashHistory struct {
	Index    string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Datetime int32  `protobuf:"varint,2,opt,name=datetime,proto3" json:"datetime,omitempty"`
	Amount   int32  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Level    int32  `protobuf:"varint,4,opt,name=level,proto3" json:"level,omitempty"`
}

func (m *SlashHistory) Reset()         { *m = SlashHistory{} }
func (m *SlashHistory) String() string { return proto.CompactTextString(m) }
func (*SlashHistory) ProtoMessage()    {}
func (*SlashHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbedc0e929ecf38f, []int{0}
}
func (m *SlashHistory) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SlashHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SlashHistory.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SlashHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SlashHistory.Merge(m, src)
}
func (m *SlashHistory) XXX_Size() int {
	return m.Size()
}
func (m *SlashHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_SlashHistory.DiscardUnknown(m)
}

var xxx_messageInfo_SlashHistory proto.InternalMessageInfo

func (m *SlashHistory) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *SlashHistory) GetDatetime() int32 {
	if m != nil {
		return m.Datetime
	}
	return 0
}

func (m *SlashHistory) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *SlashHistory) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func init() {
	proto.RegisterType((*SlashHistory)(nil), "dhpc.user.SlashHistory")
}

func init() { proto.RegisterFile("dhpc/user/slash_history.proto", fileDescriptor_cbedc0e929ecf38f) }

var fileDescriptor_cbedc0e929ecf38f = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0xc9, 0x28, 0x48,
	0xd6, 0x2f, 0x2d, 0x4e, 0x2d, 0xd2, 0x2f, 0xce, 0x49, 0x2c, 0xce, 0x88, 0xcf, 0xc8, 0x2c, 0x2e,
	0xc9, 0x2f, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x04, 0x49, 0xeb, 0x81, 0xa4,
	0x95, 0xf2, 0xb8, 0x78, 0x82, 0x41, 0x2a, 0x3c, 0x20, 0x0a, 0x84, 0x44, 0xb8, 0x58, 0x33, 0xf3,
	0x52, 0x52, 0x2b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x20, 0x1c, 0x21, 0x29, 0x2e, 0x8e,
	0x94, 0xc4, 0x92, 0xd4, 0x92, 0xcc, 0xdc, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x38,
	0x5f, 0x48, 0x8c, 0x8b, 0x2d, 0x31, 0x37, 0xbf, 0x34, 0xaf, 0x44, 0x82, 0x19, 0x2c, 0x03, 0xe5,
	0x81, 0x4c, 0xca, 0x49, 0x2d, 0x4b, 0xcd, 0x91, 0x60, 0x01, 0x0b, 0x43, 0x38, 0x4e, 0x0e, 0x27,
	0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c,
	0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x96, 0x9e, 0x59, 0x92, 0x51, 0x9a,
	0xa4, 0x97, 0x9c, 0x9f, 0xab, 0xef, 0x92, 0x51, 0x90, 0xec, 0x9c, 0x91, 0x98, 0x99, 0x07, 0x66,
	0xe9, 0x57, 0x40, 0xbc, 0x52, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0xf6, 0x83, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0x25, 0x3c, 0xa5, 0x89, 0xe4, 0x00, 0x00, 0x00,
}

func (m *SlashHistory) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SlashHistory) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SlashHistory) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Level != 0 {
		i = encodeVarintSlashHistory(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x20
	}
	if m.Amount != 0 {
		i = encodeVarintSlashHistory(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x18
	}
	if m.Datetime != 0 {
		i = encodeVarintSlashHistory(dAtA, i, uint64(m.Datetime))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintSlashHistory(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSlashHistory(dAtA []byte, offset int, v uint64) int {
	offset -= sovSlashHistory(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SlashHistory) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovSlashHistory(uint64(l))
	}
	if m.Datetime != 0 {
		n += 1 + sovSlashHistory(uint64(m.Datetime))
	}
	if m.Amount != 0 {
		n += 1 + sovSlashHistory(uint64(m.Amount))
	}
	if m.Level != 0 {
		n += 1 + sovSlashHistory(uint64(m.Level))
	}
	return n
}

func sovSlashHistory(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSlashHistory(x uint64) (n int) {
	return sovSlashHistory(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SlashHistory) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSlashHistory
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
			return fmt.Errorf("proto: SlashHistory: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SlashHistory: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashHistory
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
				return ErrInvalidLengthSlashHistory
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSlashHistory
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Datetime", wireType)
			}
			m.Datetime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashHistory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Datetime |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashHistory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Level", wireType)
			}
			m.Level = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSlashHistory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Level |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSlashHistory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSlashHistory
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
func skipSlashHistory(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSlashHistory
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
					return 0, ErrIntOverflowSlashHistory
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
					return 0, ErrIntOverflowSlashHistory
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
				return 0, ErrInvalidLengthSlashHistory
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSlashHistory
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSlashHistory
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSlashHistory        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSlashHistory          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSlashHistory = fmt.Errorf("proto: unexpected end of group")
)
