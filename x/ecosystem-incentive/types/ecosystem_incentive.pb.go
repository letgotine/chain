// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ecosystem-incentive/ecosystem_incentive.proto

package types

import (
	fmt "fmt"
	github_com_UnUniFi_chain_types "github.com/UnUniFi/chain/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type IncentiveUnit struct {
	Id              string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" yaml:"id"`
	SubjectInfoList []SubjectInfo `protobuf:"bytes,2,rep,name=subject_info_list,json=subjectInfoList,proto3" json:"subject_info_list" yaml:"subject_info_lists"`
}

func (m *IncentiveUnit) Reset()         { *m = IncentiveUnit{} }
func (m *IncentiveUnit) String() string { return proto.CompactTextString(m) }
func (*IncentiveUnit) ProtoMessage()    {}
func (*IncentiveUnit) Descriptor() ([]byte, []int) {
	return fileDescriptor_b09f3079d309ee36, []int{0}
}
func (m *IncentiveUnit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IncentiveUnit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IncentiveUnit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IncentiveUnit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncentiveUnit.Merge(m, src)
}
func (m *IncentiveUnit) XXX_Size() int {
	return m.Size()
}
func (m *IncentiveUnit) XXX_DiscardUnknown() {
	xxx_messageInfo_IncentiveUnit.DiscardUnknown(m)
}

var xxx_messageInfo_IncentiveUnit proto.InternalMessageInfo

func (m *IncentiveUnit) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *IncentiveUnit) GetSubjectInfoList() []SubjectInfo {
	if m != nil {
		return m.SubjectInfoList
	}
	return nil
}

type SubjectInfo struct {
	Address github_com_UnUniFi_chain_types.StringAccAddress `protobuf:"bytes,1,opt,name=address,proto3,customtype=github.com/UnUniFi/chain/types.StringAccAddress" json:"address" yaml:"subject_addr"`
	Weight  github_com_cosmos_cosmos_sdk_types.Dec          `protobuf:"bytes,2,opt,name=weight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"weight" yaml:"weight"`
}

func (m *SubjectInfo) Reset()         { *m = SubjectInfo{} }
func (m *SubjectInfo) String() string { return proto.CompactTextString(m) }
func (*SubjectInfo) ProtoMessage()    {}
func (*SubjectInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b09f3079d309ee36, []int{1}
}
func (m *SubjectInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SubjectInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SubjectInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SubjectInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubjectInfo.Merge(m, src)
}
func (m *SubjectInfo) XXX_Size() int {
	return m.Size()
}
func (m *SubjectInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SubjectInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SubjectInfo proto.InternalMessageInfo

type Reward struct {
	SubjectAddr github_com_UnUniFi_chain_types.StringAccAddress `protobuf:"bytes,1,opt,name=subject_addr,json=subjectAddr,proto3,customtype=github.com/UnUniFi/chain/types.StringAccAddress" json:"subject_addr" yaml:"subject_addr"`
	Rewards     github_com_cosmos_cosmos_sdk_types.Coins        `protobuf:"bytes,2,rep,name=rewards,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"rewards" yaml:"rewards"`
}

func (m *Reward) Reset()         { *m = Reward{} }
func (m *Reward) String() string { return proto.CompactTextString(m) }
func (*Reward) ProtoMessage()    {}
func (*Reward) Descriptor() ([]byte, []int) {
	return fileDescriptor_b09f3079d309ee36, []int{2}
}
func (m *Reward) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Reward) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Reward.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Reward) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reward.Merge(m, src)
}
func (m *Reward) XXX_Size() int {
	return m.Size()
}
func (m *Reward) XXX_DiscardUnknown() {
	xxx_messageInfo_Reward.DiscardUnknown(m)
}

var xxx_messageInfo_Reward proto.InternalMessageInfo

func (m *Reward) GetRewards() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Rewards
	}
	return nil
}

func init() {
	proto.RegisterType((*IncentiveUnit)(nil), "ununifi.ecosystemincentive.IncentiveUnit")
	proto.RegisterType((*SubjectInfo)(nil), "ununifi.ecosystemincentive.SubjectInfo")
	proto.RegisterType((*Reward)(nil), "ununifi.ecosystemincentive.Reward")
}

func init() {
	proto.RegisterFile("ecosystem-incentive/ecosystem_incentive.proto", fileDescriptor_b09f3079d309ee36)
}

var fileDescriptor_b09f3079d309ee36 = []byte{
	// 474 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xb1, 0x8e, 0xd3, 0x30,
	0x18, 0xc7, 0x9b, 0x22, 0xf5, 0x74, 0x2e, 0x05, 0x11, 0x18, 0xee, 0x2a, 0x48, 0x0e, 0x0f, 0xd0,
	0xa5, 0xb6, 0x0e, 0x98, 0x58, 0xd0, 0x05, 0x84, 0x74, 0x12, 0xba, 0x21, 0xa7, 0x0a, 0x89, 0xa5,
	0x72, 0x1c, 0x37, 0xfd, 0xa0, 0xb5, 0x4b, 0xec, 0xb6, 0xf4, 0x2d, 0x78, 0x0a, 0x06, 0x9e, 0xe4,
	0xc6, 0x8e, 0x88, 0x21, 0xa0, 0xf6, 0x0d, 0x2a, 0x1e, 0x00, 0xa5, 0x76, 0x4b, 0x11, 0x54, 0x62,
	0x61, 0x4a, 0x62, 0xff, 0xbf, 0xdf, 0xff, 0xff, 0xd9, 0x5f, 0x50, 0x5b, 0x70, 0xa5, 0x67, 0xda,
	0x88, 0x61, 0x1b, 0x24, 0x17, 0xd2, 0xc0, 0x44, 0xd0, 0xed, 0x5a, 0x77, 0xbb, 0x46, 0x46, 0xb9,
	0x32, 0xca, 0x6f, 0x8e, 0xe5, 0x58, 0x42, 0x0f, 0xc8, 0x56, 0xb2, 0x55, 0x34, 0xef, 0x64, 0x2a,
	0x53, 0x6b, 0x19, 0x2d, 0xdf, 0x6c, 0x45, 0xf3, 0x6e, 0xa6, 0x54, 0x36, 0x10, 0x94, 0x8d, 0x80,
	0x32, 0x29, 0x95, 0x61, 0x06, 0x94, 0xd4, 0x6e, 0x37, 0xe0, 0x4a, 0x0f, 0x95, 0xa6, 0x09, 0xd3,
	0x82, 0x4e, 0x4e, 0x13, 0x61, 0xd8, 0x29, 0xe5, 0x0a, 0xa4, 0xdd, 0xc7, 0x9f, 0x3c, 0xd4, 0x38,
	0xdf, 0x38, 0x74, 0x24, 0x18, 0xff, 0x1e, 0xaa, 0x42, 0x7a, 0xe4, 0x9d, 0x78, 0xad, 0xc3, 0xa8,
	0xb1, 0x2a, 0xc2, 0xc3, 0x19, 0x1b, 0x0e, 0x9e, 0x62, 0x48, 0x71, 0x5c, 0x85, 0xd4, 0x9f, 0xa0,
	0x5b, 0x7a, 0x9c, 0xbc, 0x15, 0xdc, 0x74, 0x41, 0xf6, 0x54, 0x77, 0x00, 0xda, 0x1c, 0x55, 0x4f,
	0xae, 0xb5, 0xea, 0x8f, 0x1e, 0x92, 0xfd, 0xe1, 0xc9, 0xa5, 0x2d, 0x3a, 0x97, 0x3d, 0x15, 0xdd,
	0xbf, 0x2a, 0xc2, 0xca, 0xaa, 0x08, 0x8f, 0x2d, 0xfa, 0x0f, 0x9e, 0xc6, 0xf1, 0x4d, 0xfd, 0x4b,
	0xff, 0x0a, 0xb4, 0xc1, 0x73, 0x0f, 0xd5, 0x77, 0x18, 0x7e, 0x1f, 0x1d, 0xb0, 0x34, 0xcd, 0x85,
	0xd6, 0x2e, 0xeb, 0x45, 0x09, 0xfd, 0x5a, 0x84, 0x34, 0x03, 0xd3, 0x1f, 0x27, 0x84, 0xab, 0x21,
	0xed, 0xc8, 0x8e, 0x84, 0x97, 0x40, 0x79, 0x9f, 0x81, 0xa4, 0x66, 0x36, 0x12, 0x9a, 0x5c, 0x9a,
	0x1c, 0x64, 0x76, 0xc6, 0xf9, 0x99, 0x2d, 0x5f, 0x15, 0xe1, 0xed, 0xdf, 0x73, 0x94, 0x5c, 0x1c,
	0x6f, 0xf0, 0xfe, 0x6b, 0x54, 0x9b, 0x0a, 0xc8, 0xfa, 0x65, 0x9b, 0xa5, 0xd1, 0x33, 0x67, 0xf4,
	0x60, 0xc7, 0xc8, 0x9d, 0xb2, 0x7d, 0xb4, 0x75, 0xfa, 0xce, 0x99, 0xbd, 0x10, 0x7c, 0x55, 0x84,
	0x0d, 0xcb, 0xb7, 0x14, 0x1c, 0x3b, 0x1c, 0xfe, 0xe1, 0xa1, 0x5a, 0x2c, 0xa6, 0x2c, 0x4f, 0xfd,
	0xf7, 0xe8, 0xfa, 0xae, 0xfb, 0x7f, 0x6a, 0xa9, 0xee, 0x3e, 0x4b, 0x99, 0x3f, 0x45, 0x07, 0xf9,
	0xda, 0x5c, 0xbb, 0xeb, 0x3b, 0x26, 0x36, 0x3e, 0x29, 0x67, 0x85, 0xb8, 0x59, 0x21, 0xcf, 0x15,
	0xc8, 0x28, 0x72, 0x17, 0x76, 0xc3, 0x52, 0x5d, 0x1d, 0xfe, 0xfc, 0x2d, 0x6c, 0xfd, 0xc3, 0x21,
	0x94, 0x08, 0x1d, 0x6f, 0xdc, 0xa2, 0x8b, 0xab, 0x45, 0xe0, 0xcd, 0x17, 0x81, 0xf7, 0x7d, 0x11,
	0x78, 0x1f, 0x97, 0x41, 0x65, 0xbe, 0x0c, 0x2a, 0x5f, 0x96, 0x41, 0xe5, 0xcd, 0x93, 0xbd, 0x7d,
	0x7e, 0xa0, 0x7f, 0xfb, 0x8d, 0xd6, 0xf8, 0xa4, 0xb6, 0x9e, 0xe4, 0xc7, 0x3f, 0x03, 0x00, 0x00,
	0xff, 0xff, 0xcd, 0x87, 0xc3, 0xb3, 0x6a, 0x03, 0x00, 0x00,
}

func (m *IncentiveUnit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IncentiveUnit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IncentiveUnit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SubjectInfoList) > 0 {
		for iNdEx := len(m.SubjectInfoList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SubjectInfoList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEcosystemIncentive(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintEcosystemIncentive(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SubjectInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SubjectInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SubjectInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Weight.Size()
		i -= size
		if _, err := m.Weight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEcosystemIncentive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Address.Size()
		i -= size
		if _, err := m.Address.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEcosystemIncentive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Reward) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Reward) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Reward) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Rewards) > 0 {
		for iNdEx := len(m.Rewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Rewards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEcosystemIncentive(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size := m.SubjectAddr.Size()
		i -= size
		if _, err := m.SubjectAddr.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEcosystemIncentive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintEcosystemIncentive(dAtA []byte, offset int, v uint64) int {
	offset -= sovEcosystemIncentive(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IncentiveUnit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovEcosystemIncentive(uint64(l))
	}
	if len(m.SubjectInfoList) > 0 {
		for _, e := range m.SubjectInfoList {
			l = e.Size()
			n += 1 + l + sovEcosystemIncentive(uint64(l))
		}
	}
	return n
}

func (m *SubjectInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Address.Size()
	n += 1 + l + sovEcosystemIncentive(uint64(l))
	l = m.Weight.Size()
	n += 1 + l + sovEcosystemIncentive(uint64(l))
	return n
}

func (m *Reward) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.SubjectAddr.Size()
	n += 1 + l + sovEcosystemIncentive(uint64(l))
	if len(m.Rewards) > 0 {
		for _, e := range m.Rewards {
			l = e.Size()
			n += 1 + l + sovEcosystemIncentive(uint64(l))
		}
	}
	return n
}

func sovEcosystemIncentive(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEcosystemIncentive(x uint64) (n int) {
	return sovEcosystemIncentive(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IncentiveUnit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEcosystemIncentive
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
			return fmt.Errorf("proto: IncentiveUnit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IncentiveUnit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEcosystemIncentive
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
				return ErrInvalidLengthEcosystemIncentive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEcosystemIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubjectInfoList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEcosystemIncentive
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
				return ErrInvalidLengthEcosystemIncentive
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEcosystemIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubjectInfoList = append(m.SubjectInfoList, SubjectInfo{})
			if err := m.SubjectInfoList[len(m.SubjectInfoList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEcosystemIncentive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEcosystemIncentive
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
func (m *SubjectInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEcosystemIncentive
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
			return fmt.Errorf("proto: SubjectInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubjectInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEcosystemIncentive
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
				return ErrInvalidLengthEcosystemIncentive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEcosystemIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Address.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEcosystemIncentive
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
				return ErrInvalidLengthEcosystemIncentive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEcosystemIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Weight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEcosystemIncentive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEcosystemIncentive
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
func (m *Reward) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEcosystemIncentive
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
			return fmt.Errorf("proto: Reward: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Reward: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubjectAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEcosystemIncentive
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
				return ErrInvalidLengthEcosystemIncentive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEcosystemIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SubjectAddr.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEcosystemIncentive
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
				return ErrInvalidLengthEcosystemIncentive
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEcosystemIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rewards = append(m.Rewards, types.Coin{})
			if err := m.Rewards[len(m.Rewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEcosystemIncentive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEcosystemIncentive
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
func skipEcosystemIncentive(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEcosystemIncentive
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
					return 0, ErrIntOverflowEcosystemIncentive
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
					return 0, ErrIntOverflowEcosystemIncentive
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
				return 0, ErrInvalidLengthEcosystemIncentive
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEcosystemIncentive
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEcosystemIncentive
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEcosystemIncentive        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEcosystemIncentive          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEcosystemIncentive = fmt.Errorf("proto: unexpected end of group")
)
