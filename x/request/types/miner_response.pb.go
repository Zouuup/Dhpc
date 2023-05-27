// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: decent/request/miner_response.proto

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

type MinerResponse struct {
	UUID        string `protobuf:"bytes,1,opt,name=uUID,proto3" json:"uUID,omitempty"`
	RequestUUID string `protobuf:"bytes,2,opt,name=requestUUID,proto3" json:"requestUUID,omitempty"`
	Hash        string `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	Answer      int32  `protobuf:"varint,4,opt,name=answer,proto3" json:"answer,omitempty"`
	Creator     string `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
	Dataused    string `protobuf:"bytes,6,opt,name=dataused,proto3" json:"dataused,omitempty"`
}

func (m *MinerResponse) Reset()         { *m = MinerResponse{} }
func (m *MinerResponse) String() string { return proto.CompactTextString(m) }
func (*MinerResponse) ProtoMessage()    {}
func (*MinerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5e8810e07127561, []int{0}
}
func (m *MinerResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MinerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MinerResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MinerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MinerResponse.Merge(m, src)
}
func (m *MinerResponse) XXX_Size() int {
	return m.Size()
}
func (m *MinerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MinerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MinerResponse proto.InternalMessageInfo

func (m *MinerResponse) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *MinerResponse) GetRequestUUID() string {
	if m != nil {
		return m.RequestUUID
	}
	return ""
}

func (m *MinerResponse) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *MinerResponse) GetAnswer() int32 {
	if m != nil {
		return m.Answer
	}
	return 0
}

func (m *MinerResponse) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MinerResponse) GetDataused() string {
	if m != nil {
		return m.Dataused
	}
	return ""
}

func init() {
	proto.RegisterType((*MinerResponse)(nil), "decent.request.MinerResponse")
}

func init() {
	proto.RegisterFile("decent/request/miner_response.proto", fileDescriptor_d5e8810e07127561)
}

var fileDescriptor_d5e8810e07127561 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0x49, 0x4d, 0x4e,
	0xcd, 0x2b, 0xd1, 0x2f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0xd1, 0xcf, 0xcd, 0xcc, 0x4b, 0x2d,
	0x8a, 0x2f, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x83, 0x28, 0xd2, 0x83, 0x2a, 0x52, 0x5a, 0xce, 0xc8, 0xc5, 0xeb, 0x0b, 0x52, 0x18, 0x04,
	0x55, 0x27, 0x24, 0xc4, 0xc5, 0x52, 0x1a, 0xea, 0xe9, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19,
	0x04, 0x66, 0x0b, 0x29, 0x70, 0x71, 0x43, 0x35, 0x84, 0x82, 0xa4, 0x98, 0xc0, 0x52, 0xc8, 0x42,
	0x20, 0x5d, 0x19, 0x89, 0xc5, 0x19, 0x12, 0xcc, 0x10, 0x5d, 0x20, 0xb6, 0x90, 0x18, 0x17, 0x5b,
	0x62, 0x5e, 0x71, 0x79, 0x6a, 0x91, 0x04, 0x8b, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x94, 0x27, 0x24,
	0xc1, 0xc5, 0x9e, 0x5c, 0x94, 0x9a, 0x58, 0x92, 0x5f, 0x24, 0xc1, 0x0a, 0x56, 0x0e, 0xe3, 0x0a,
	0x49, 0x71, 0x71, 0xa4, 0x24, 0x96, 0x24, 0x96, 0x16, 0xa7, 0xa6, 0x48, 0xb0, 0x81, 0xa5, 0xe0,
	0x7c, 0x27, 0x83, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71,
	0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x12, 0x73, 0x81,
	0x78, 0xbc, 0x02, 0xee, 0xf5, 0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0, 0x97, 0x8d, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x66, 0xb6, 0xd6, 0x19, 0x01, 0x00, 0x00,
}

func (m *MinerResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MinerResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MinerResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Dataused) > 0 {
		i -= len(m.Dataused)
		copy(dAtA[i:], m.Dataused)
		i = encodeVarintMinerResponse(dAtA, i, uint64(len(m.Dataused)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintMinerResponse(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Answer != 0 {
		i = encodeVarintMinerResponse(dAtA, i, uint64(m.Answer))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintMinerResponse(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.RequestUUID) > 0 {
		i -= len(m.RequestUUID)
		copy(dAtA[i:], m.RequestUUID)
		i = encodeVarintMinerResponse(dAtA, i, uint64(len(m.RequestUUID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.UUID) > 0 {
		i -= len(m.UUID)
		copy(dAtA[i:], m.UUID)
		i = encodeVarintMinerResponse(dAtA, i, uint64(len(m.UUID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMinerResponse(dAtA []byte, offset int, v uint64) int {
	offset -= sovMinerResponse(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MinerResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.UUID)
	if l > 0 {
		n += 1 + l + sovMinerResponse(uint64(l))
	}
	l = len(m.RequestUUID)
	if l > 0 {
		n += 1 + l + sovMinerResponse(uint64(l))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovMinerResponse(uint64(l))
	}
	if m.Answer != 0 {
		n += 1 + sovMinerResponse(uint64(m.Answer))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovMinerResponse(uint64(l))
	}
	l = len(m.Dataused)
	if l > 0 {
		n += 1 + l + sovMinerResponse(uint64(l))
	}
	return n
}

func sovMinerResponse(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMinerResponse(x uint64) (n int) {
	return sovMinerResponse(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MinerResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMinerResponse
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
			return fmt.Errorf("proto: MinerResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MinerResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinerResponse
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
				return ErrInvalidLengthMinerResponse
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMinerResponse
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestUUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinerResponse
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
				return ErrInvalidLengthMinerResponse
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMinerResponse
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestUUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinerResponse
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
				return ErrInvalidLengthMinerResponse
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMinerResponse
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Answer", wireType)
			}
			m.Answer = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinerResponse
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Answer |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinerResponse
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
				return ErrInvalidLengthMinerResponse
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMinerResponse
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dataused", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinerResponse
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
				return ErrInvalidLengthMinerResponse
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMinerResponse
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Dataused = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMinerResponse(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMinerResponse
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
func skipMinerResponse(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMinerResponse
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
					return 0, ErrIntOverflowMinerResponse
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
					return 0, ErrIntOverflowMinerResponse
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
				return 0, ErrInvalidLengthMinerResponse
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMinerResponse
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMinerResponse
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMinerResponse        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMinerResponse          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMinerResponse = fmt.Errorf("proto: unexpected end of group")
)
