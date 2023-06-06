// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: eventhook/proposal.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// proposal to add new hook
type ProposalAddHook struct {
	Title           string          `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description     string          `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	EventType       string          `protobuf:"bytes,3,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
	ContractAddress string          `protobuf:"bytes,4,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	Name            string          `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	GitUrl          string          `protobuf:"bytes,6,opt,name=git_url,json=gitUrl,proto3" json:"git_url,omitempty"`
	EventAttributes []*KeyValuePair `protobuf:"bytes,7,rep,name=event_attributes,json=eventAttributes,proto3" json:"event_attributes,omitempty"`
}

func (m *ProposalAddHook) Reset()         { *m = ProposalAddHook{} }
func (m *ProposalAddHook) String() string { return proto.CompactTextString(m) }
func (*ProposalAddHook) ProtoMessage()    {}
func (*ProposalAddHook) Descriptor() ([]byte, []int) {
	return fileDescriptor_66fa6eb051d89299, []int{0}
}
func (m *ProposalAddHook) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProposalAddHook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProposalAddHook.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProposalAddHook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalAddHook.Merge(m, src)
}
func (m *ProposalAddHook) XXX_Size() int {
	return m.Size()
}
func (m *ProposalAddHook) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposalAddHook.DiscardUnknown(m)
}

var xxx_messageInfo_ProposalAddHook proto.InternalMessageInfo

func (m *ProposalAddHook) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ProposalAddHook) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ProposalAddHook) GetEventType() string {
	if m != nil {
		return m.EventType
	}
	return ""
}

func (m *ProposalAddHook) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *ProposalAddHook) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProposalAddHook) GetGitUrl() string {
	if m != nil {
		return m.GitUrl
	}
	return ""
}

func (m *ProposalAddHook) GetEventAttributes() []*KeyValuePair {
	if m != nil {
		return m.EventAttributes
	}
	return nil
}

func init() {
	proto.RegisterType((*ProposalAddHook)(nil), "nftvault.eventhook.ProposalAddHook")
}

func init() { proto.RegisterFile("eventhook/proposal.proto", fileDescriptor_66fa6eb051d89299) }

var fileDescriptor_66fa6eb051d89299 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xcf, 0x6a, 0x2a, 0x31,
	0x18, 0xc5, 0x1d, 0xff, 0x62, 0x5c, 0x28, 0x41, 0xb8, 0xb9, 0x42, 0x87, 0xa1, 0x2b, 0xdb, 0xc5,
	0x0c, 0xb4, 0x4f, 0x60, 0x29, 0xa5, 0xe0, 0x46, 0xa4, 0x76, 0xd1, 0x8d, 0xc4, 0x99, 0x74, 0x0c,
	0x8e, 0xf9, 0x42, 0xf2, 0x8d, 0xd4, 0x87, 0x28, 0xf4, 0xb1, 0xba, 0x74, 0xd9, 0x65, 0xd1, 0x17,
	0x29, 0x66, 0xda, 0xb1, 0xd0, 0xdd, 0xc9, 0xef, 0x9c, 0x9c, 0x90, 0x43, 0x98, 0xd8, 0x08, 0x85,
	0x4b, 0x80, 0x55, 0xa4, 0x0d, 0x68, 0xb0, 0x3c, 0x0b, 0xb5, 0x01, 0x04, 0x4a, 0xd5, 0x33, 0x6e,
	0x78, 0x9e, 0x61, 0x58, 0x46, 0x06, 0xfd, 0x14, 0x52, 0x70, 0x76, 0x74, 0x54, 0x45, 0x72, 0xf0,
	0xff, 0xd4, 0x51, 0xaa, 0xc2, 0x3a, 0x7f, 0xad, 0x92, 0xee, 0xe4, 0xbb, 0x77, 0x94, 0x24, 0xf7,
	0x00, 0x2b, 0xda, 0x27, 0x0d, 0x94, 0x98, 0x09, 0xe6, 0x05, 0xde, 0xb0, 0x3d, 0x2d, 0x0e, 0x34,
	0x20, 0x9d, 0x44, 0xd8, 0xd8, 0x48, 0x8d, 0x12, 0x14, 0xab, 0x3a, 0xef, 0x37, 0xa2, 0x67, 0x84,
	0xb8, 0xfa, 0x39, 0x6e, 0xb5, 0x60, 0x35, 0x17, 0x68, 0x3b, 0xf2, 0xb0, 0xd5, 0x82, 0x5e, 0x90,
	0x5e, 0x0c, 0x0a, 0x0d, 0x8f, 0x71, 0xce, 0x93, 0xc4, 0x08, 0x6b, 0x59, 0xdd, 0x85, 0xba, 0x3f,
	0x7c, 0x54, 0x60, 0x4a, 0x49, 0x5d, 0xf1, 0xb5, 0x60, 0x0d, 0x67, 0x3b, 0x4d, 0xff, 0x91, 0x56,
	0x2a, 0x71, 0x9e, 0x9b, 0x8c, 0x35, 0x1d, 0x6e, 0xa6, 0x12, 0x67, 0x26, 0xa3, 0x63, 0xd2, 0x2b,
	0x9e, 0xe5, 0x88, 0x46, 0x2e, 0x72, 0x14, 0x96, 0xb5, 0x82, 0xda, 0xb0, 0x73, 0x15, 0x84, 0x7f,
	0x27, 0x0a, 0xc7, 0x62, 0xfb, 0xc8, 0xb3, 0x5c, 0x4c, 0xb8, 0x34, 0xd3, 0xae, 0xe3, 0xa3, 0xf2,
	0xe2, 0xcd, 0xed, 0xfb, 0xde, 0xf7, 0x76, 0x7b, 0xdf, 0xfb, 0xdc, 0xfb, 0xde, 0xdb, 0xc1, 0xaf,
	0xec, 0x0e, 0x7e, 0xe5, 0xe3, 0xe0, 0x57, 0x9e, 0x2e, 0x53, 0x89, 0xcb, 0x7c, 0x11, 0xc6, 0xb0,
	0x8e, 0x66, 0x6a, 0xa6, 0xe4, 0x9d, 0x8c, 0xe2, 0x25, 0x97, 0x2a, 0x7a, 0x39, 0xad, 0x1a, 0x1d,
	0xbf, 0x6e, 0x17, 0x4d, 0x37, 0xee, 0xf5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x5f, 0x64,
	0x81, 0xbd, 0x01, 0x00, 0x00,
}

func (m *ProposalAddHook) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProposalAddHook) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProposalAddHook) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EventAttributes) > 0 {
		for iNdEx := len(m.EventAttributes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.EventAttributes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProposal(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.GitUrl) > 0 {
		i -= len(m.GitUrl)
		copy(dAtA[i:], m.GitUrl)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.GitUrl)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.EventType) > 0 {
		i -= len(m.EventType)
		copy(dAtA[i:], m.EventType)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.EventType)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProposalAddHook) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.EventType)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.GitUrl)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if len(m.EventAttributes) > 0 {
		for _, e := range m.EventAttributes {
			l = e.Size()
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	return n
}

func sovProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposal(x uint64) (n int) {
	return sovProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProposalAddHook) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: ProposalAddHook: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProposalAddHook: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EventType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GitUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GitUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventAttributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EventAttributes = append(m.EventAttributes, &KeyValuePair{})
			if err := m.EventAttributes[len(m.EventAttributes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func skipProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
				return 0, ErrInvalidLengthProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposal = fmt.Errorf("proto: unexpected end of group")
)
