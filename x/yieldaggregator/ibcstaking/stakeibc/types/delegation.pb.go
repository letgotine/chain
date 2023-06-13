// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stakeibc/delegation.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type Delegation struct {
	DelegateAcctAddress string     `protobuf:"bytes,1,opt,name=delegateAcctAddress,proto3" json:"delegateAcctAddress,omitempty"`
	Validator           *Validator `protobuf:"bytes,2,opt,name=validator,proto3" json:"validator,omitempty"`
	Amt                 int64      `protobuf:"varint,3,opt,name=amt,proto3" json:"amt,omitempty"`
}

func (m *Delegation) Reset()         { *m = Delegation{} }
func (m *Delegation) String() string { return proto.CompactTextString(m) }
func (*Delegation) ProtoMessage()    {}
func (*Delegation) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c983a13a7d09465, []int{0}
}
func (m *Delegation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Delegation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Delegation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Delegation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Delegation.Merge(m, src)
}
func (m *Delegation) XXX_Size() int {
	return m.Size()
}
func (m *Delegation) XXX_DiscardUnknown() {
	xxx_messageInfo_Delegation.DiscardUnknown(m)
}

var xxx_messageInfo_Delegation proto.InternalMessageInfo

func (m *Delegation) GetDelegateAcctAddress() string {
	if m != nil {
		return m.DelegateAcctAddress
	}
	return ""
}

func (m *Delegation) GetValidator() *Validator {
	if m != nil {
		return m.Validator
	}
	return nil
}

func (m *Delegation) GetAmt() int64 {
	if m != nil {
		return m.Amt
	}
	return 0
}

func init() {
	proto.RegisterType((*Delegation)(nil), "Stridelabs.stride.stakeibc.Delegation")
}

func init() {
	// proto.RegisterFile("stakeibc/delegation.proto", fileDescriptor_4c983a13a7d09465)
}

var fileDescriptor_4c983a13a7d09465 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x2e, 0x49, 0xcc,
	0x4e, 0xcd, 0x4c, 0x4a, 0xd6, 0x4f, 0x49, 0xcd, 0x49, 0x4d, 0x4f, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x0a, 0x2e, 0x29, 0xca, 0x4c, 0x49, 0xcd, 0x49, 0x4c,
	0x2a, 0xd6, 0x2b, 0x06, 0x33, 0xf5, 0x60, 0x8a, 0xa5, 0x24, 0xe0, 0xda, 0xca, 0x12, 0x73, 0x32,
	0x53, 0x12, 0x4b, 0xf2, 0x8b, 0x20, 0xba, 0x94, 0xa6, 0x32, 0x72, 0x71, 0xb9, 0xc0, 0x8d, 0x12,
	0x32, 0xe0, 0x12, 0x86, 0x1a, 0x9c, 0xea, 0x98, 0x9c, 0x5c, 0xe2, 0x98, 0x92, 0x52, 0x94, 0x5a,
	0x5c, 0x2c, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x84, 0x4d, 0x4a, 0xc8, 0x99, 0x8b, 0x13, 0x6e,
	0xa6, 0x04, 0x93, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0xaa, 0x1e, 0x6e, 0xa7, 0xe8, 0x85, 0xc1, 0x14,
	0x07, 0x21, 0xf4, 0x09, 0x09, 0x70, 0x31, 0x27, 0xe6, 0x96, 0x48, 0x30, 0x2b, 0x30, 0x6a, 0x30,
	0x07, 0x81, 0x98, 0x4e, 0x1e, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91,
	0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5,
	0x97, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x0f, 0xb1, 0x47, 0xd7, 0x27,
	0x31, 0xa9, 0x58, 0x1f, 0x62, 0x91, 0x7e, 0x85, 0x3e, 0xdc, 0xaf, 0x25, 0x95, 0x05, 0xa9, 0xc5,
	0x49, 0x6c, 0x60, 0x8f, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x19, 0x6e, 0x0f, 0x3b,
	0x01, 0x00, 0x00,
}

func (m *Delegation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Delegation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Delegation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Amt != 0 {
		i = encodeVarintDelegation(dAtA, i, uint64(m.Amt))
		i--
		dAtA[i] = 0x18
	}
	if m.Validator != nil {
		{
			size, err := m.Validator.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDelegation(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.DelegateAcctAddress) > 0 {
		i -= len(m.DelegateAcctAddress)
		copy(dAtA[i:], m.DelegateAcctAddress)
		i = encodeVarintDelegation(dAtA, i, uint64(len(m.DelegateAcctAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDelegation(dAtA []byte, offset int, v uint64) int {
	offset -= sovDelegation(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Delegation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DelegateAcctAddress)
	if l > 0 {
		n += 1 + l + sovDelegation(uint64(l))
	}
	if m.Validator != nil {
		l = m.Validator.Size()
		n += 1 + l + sovDelegation(uint64(l))
	}
	if m.Amt != 0 {
		n += 1 + sovDelegation(uint64(m.Amt))
	}
	return n
}

func sovDelegation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDelegation(x uint64) (n int) {
	return sovDelegation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Delegation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDelegation
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
			return fmt.Errorf("proto: Delegation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Delegation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegateAcctAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegation
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
				return ErrInvalidLengthDelegation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDelegation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegateAcctAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegation
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
				return ErrInvalidLengthDelegation
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDelegation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Validator == nil {
				m.Validator = &Validator{}
			}
			if err := m.Validator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amt", wireType)
			}
			m.Amt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDelegation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDelegation
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
func skipDelegation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDelegation
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
					return 0, ErrIntOverflowDelegation
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
					return 0, ErrIntOverflowDelegation
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
				return 0, ErrInvalidLengthDelegation
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDelegation
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDelegation
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDelegation        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDelegation          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDelegation = fmt.Errorf("proto: unexpected end of group")
)
