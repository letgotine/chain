// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cdp/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the cdp module's genesis state.
type GenesisState struct {
	Params                    Params                    `protobuf:"bytes,1,opt,name=params,proto3" json:"params" yaml:"params"`
	Cdps                      []Cdp                     `protobuf:"bytes,2,rep,name=cdps,proto3" json:"cdps" yaml:"cdps"`
	Deposits                  []Deposit                 `protobuf:"bytes,3,rep,name=deposits,proto3" json:"deposits" yaml:"deposits"`
	StartingCdpId             uint64                    `protobuf:"varint,4,opt,name=starting_cdp_id,json=startingCdpId,proto3" json:"starting_cdp_id,omitempty" yaml:"starting_cdp_id"`
	DebtDenom                 string                    `protobuf:"bytes,5,opt,name=debt_denom,json=debtDenom,proto3" json:"debt_denom,omitempty" yaml:"debt_denom"`
	GovDenom                  string                    `protobuf:"bytes,6,opt,name=gov_denom,json=govDenom,proto3" json:"gov_denom,omitempty" yaml:"gov_denom"`
	PreviousAccumulationTimes []GenesisAccumulationTime `protobuf:"bytes,7,rep,name=previous_accumulation_times,json=previousAccumulationTimes,proto3" json:"previous_accumulation_times" yaml:"previous_accumulation_times"`
	TotalPrincipals           []GenesisTotalPrincipal   `protobuf:"bytes,8,rep,name=total_principals,json=totalPrincipals,proto3" json:"total_principals" yaml:"total_principals"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_6792b73c0d9625b5, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetCdps() []Cdp {
	if m != nil {
		return m.Cdps
	}
	return nil
}

func (m *GenesisState) GetDeposits() []Deposit {
	if m != nil {
		return m.Deposits
	}
	return nil
}

func (m *GenesisState) GetStartingCdpId() uint64 {
	if m != nil {
		return m.StartingCdpId
	}
	return 0
}

func (m *GenesisState) GetDebtDenom() string {
	if m != nil {
		return m.DebtDenom
	}
	return ""
}

func (m *GenesisState) GetGovDenom() string {
	if m != nil {
		return m.GovDenom
	}
	return ""
}

func (m *GenesisState) GetPreviousAccumulationTimes() []GenesisAccumulationTime {
	if m != nil {
		return m.PreviousAccumulationTimes
	}
	return nil
}

func (m *GenesisState) GetTotalPrincipals() []GenesisTotalPrincipal {
	if m != nil {
		return m.TotalPrincipals
	}
	return nil
}

type GenesisAccumulationTime struct {
	CollateralType           string                                 `protobuf:"bytes,1,opt,name=collateral_type,json=collateralType,proto3" json:"collateral_type,omitempty" yaml:"collateral_type"`
	PreviousAccumulationTime time.Time                              `protobuf:"bytes,2,opt,name=previous_accumulation_time,json=previousAccumulationTime,proto3,stdtime" json:"previous_accumulation_time" yaml:"previous_accumulation_time"`
	InterestFactor           github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=interest_factor,json=interestFactor,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"interest_factor" yaml:"interest_factor"`
}

func (m *GenesisAccumulationTime) Reset()         { *m = GenesisAccumulationTime{} }
func (m *GenesisAccumulationTime) String() string { return proto.CompactTextString(m) }
func (*GenesisAccumulationTime) ProtoMessage()    {}
func (*GenesisAccumulationTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_6792b73c0d9625b5, []int{1}
}
func (m *GenesisAccumulationTime) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisAccumulationTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisAccumulationTime.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisAccumulationTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisAccumulationTime.Merge(m, src)
}
func (m *GenesisAccumulationTime) XXX_Size() int {
	return m.Size()
}
func (m *GenesisAccumulationTime) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisAccumulationTime.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisAccumulationTime proto.InternalMessageInfo

func (m *GenesisAccumulationTime) GetCollateralType() string {
	if m != nil {
		return m.CollateralType
	}
	return ""
}

func (m *GenesisAccumulationTime) GetPreviousAccumulationTime() time.Time {
	if m != nil {
		return m.PreviousAccumulationTime
	}
	return time.Time{}
}

type GenesisTotalPrincipal struct {
	CollateralType string                                 `protobuf:"bytes,1,opt,name=collateral_type,json=collateralType,proto3" json:"collateral_type,omitempty" yaml:"collateral_type"`
	TotalPrincipal github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=total_principal,json=totalPrincipal,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total_principal" yaml:"total_principal"`
}

func (m *GenesisTotalPrincipal) Reset()         { *m = GenesisTotalPrincipal{} }
func (m *GenesisTotalPrincipal) String() string { return proto.CompactTextString(m) }
func (*GenesisTotalPrincipal) ProtoMessage()    {}
func (*GenesisTotalPrincipal) Descriptor() ([]byte, []int) {
	return fileDescriptor_6792b73c0d9625b5, []int{2}
}
func (m *GenesisTotalPrincipal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisTotalPrincipal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisTotalPrincipal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisTotalPrincipal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisTotalPrincipal.Merge(m, src)
}
func (m *GenesisTotalPrincipal) XXX_Size() int {
	return m.Size()
}
func (m *GenesisTotalPrincipal) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisTotalPrincipal.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisTotalPrincipal proto.InternalMessageInfo

func (m *GenesisTotalPrincipal) GetCollateralType() string {
	if m != nil {
		return m.CollateralType
	}
	return ""
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "botany.cdp.GenesisState")
	proto.RegisterType((*GenesisAccumulationTime)(nil), "botany.cdp.GenesisAccumulationTime")
	proto.RegisterType((*GenesisTotalPrincipal)(nil), "botany.cdp.GenesisTotalPrincipal")
}

func init() { proto.RegisterFile("cdp/genesis.proto", fileDescriptor_6792b73c0d9625b5) }

var fileDescriptor_6792b73c0d9625b5 = []byte{
	// 674 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4f, 0x4f, 0xdb, 0x4e,
	0x10, 0x8d, 0xf9, 0xf7, 0x23, 0xcb, 0x0f, 0x02, 0x06, 0x8a, 0x9b, 0x4a, 0x71, 0xd8, 0x4a, 0x55,
	0x54, 0x09, 0x5b, 0xd0, 0x1e, 0xaa, 0xf6, 0x84, 0x41, 0x2d, 0xdc, 0x90, 0x9b, 0x53, 0x2f, 0xd6,
	0x7a, 0xbd, 0xb8, 0x6e, 0x6d, 0xef, 0xe2, 0xdd, 0x44, 0xe4, 0xdc, 0x43, 0x8f, 0xe5, 0x63, 0x71,
	0x2b, 0xc7, 0xaa, 0x95, 0xd2, 0x0a, 0xbe, 0x41, 0x3e, 0x41, 0xe5, 0x5d, 0x9b, 0x04, 0x8b, 0xa8,
	0xaa, 0xd4, 0x53, 0xbc, 0x33, 0xef, 0xbd, 0x99, 0x79, 0xb3, 0x1b, 0xb0, 0x86, 0x03, 0x66, 0x87,
	0x24, 0x25, 0x3c, 0xe2, 0x16, 0xcb, 0xa8, 0xa0, 0x3a, 0xf0, 0xa9, 0x40, 0xe9, 0xc0, 0xc2, 0x01,
	0x6b, 0x6e, 0x84, 0x34, 0xa4, 0x32, 0x6c, 0xe7, 0x5f, 0x0a, 0xd1, 0x34, 0x43, 0x4a, 0xc3, 0x98,
	0xd8, 0xf2, 0xe4, 0xf7, 0x4e, 0x6d, 0x11, 0x25, 0x84, 0x0b, 0x94, 0xb0, 0x02, 0xd0, 0xc2, 0x94,
	0x27, 0x94, 0xdb, 0x3e, 0xe2, 0xc4, 0xee, 0xef, 0xfa, 0x44, 0xa0, 0x5d, 0x1b, 0xd3, 0x28, 0x2d,
	0xf2, 0xcb, 0x79, 0x55, 0x1c, 0x14, 0x70, 0xf8, 0x69, 0x1e, 0xfc, 0xff, 0x46, 0xf5, 0xf0, 0x56,
	0x20, 0x41, 0xf4, 0x7d, 0xb0, 0xc0, 0x50, 0x86, 0x12, 0x6e, 0x68, 0x6d, 0xad, 0xb3, 0xb4, 0xa7,
	0x5b, 0xe3, 0x9e, 0xac, 0x13, 0x99, 0x71, 0x36, 0x2f, 0x87, 0x66, 0x6d, 0x34, 0x34, 0x97, 0x07,
	0x28, 0x89, 0x5f, 0x42, 0x85, 0x87, 0x6e, 0x41, 0xd4, 0x5f, 0x80, 0x39, 0x1c, 0x30, 0x6e, 0xcc,
	0xb4, 0x67, 0x3b, 0x4b, 0x7b, 0x8d, 0x49, 0x81, 0x83, 0x80, 0x39, 0xeb, 0x05, 0x7b, 0x49, 0xb1,
	0x73, 0x28, 0x74, 0x25, 0x43, 0x3f, 0x02, 0x8b, 0x01, 0x61, 0x94, 0x47, 0x82, 0x1b, 0xb3, 0x92,
	0xbd, 0x3e, 0xc9, 0x3e, 0x54, 0x39, 0x67, 0xab, 0x50, 0x68, 0x28, 0x85, 0x92, 0x02, 0xdd, 0x5b,
	0xb6, 0xee, 0x80, 0x06, 0x17, 0x28, 0x13, 0x51, 0x1a, 0x7a, 0x38, 0x60, 0x5e, 0x14, 0x18, 0x73,
	0x6d, 0xad, 0x33, 0xe7, 0x34, 0x47, 0x43, 0xf3, 0x81, 0xe2, 0x55, 0x00, 0xd0, 0x5d, 0x2e, 0x23,
	0x07, 0x01, 0x3b, 0x0e, 0xf4, 0xe7, 0x00, 0x04, 0xc4, 0x17, 0x5e, 0x40, 0x52, 0x9a, 0x18, 0xf3,
	0x6d, 0xad, 0x53, 0x77, 0x36, 0x47, 0x43, 0x73, 0xad, 0x2c, 0x5b, 0xe6, 0xa0, 0x5b, 0xcf, 0x0f,
	0x87, 0xf9, 0xb7, 0xbe, 0x0b, 0xea, 0x21, 0xed, 0x17, 0xa4, 0x05, 0x49, 0xda, 0x18, 0x0d, 0xcd,
	0x55, 0x45, 0xba, 0x4d, 0x41, 0x77, 0x31, 0xa4, 0x7d, 0x45, 0xf9, 0xa2, 0x81, 0x47, 0x2c, 0x23,
	0xfd, 0x88, 0xf6, 0xb8, 0x87, 0x30, 0xee, 0x25, 0xbd, 0x18, 0x89, 0x88, 0xa6, 0x9e, 0xdc, 0xae,
	0xf1, 0x9f, 0xb4, 0xe2, 0xf1, 0xa4, 0x15, 0xc5, 0xce, 0xf6, 0x27, 0xc0, 0xdd, 0x28, 0x21, 0xce,
	0xd3, 0xc2, 0x1a, 0x58, 0xac, 0x66, 0xba, 0x2a, 0x74, 0x1f, 0x96, 0xd9, 0xaa, 0x0a, 0xd7, 0x13,
	0xb0, 0x2a, 0xa8, 0x40, 0xb1, 0xc7, 0xb2, 0x28, 0xc5, 0x11, 0x43, 0x31, 0x37, 0x16, 0x65, 0x17,
	0xdb, 0xf7, 0x74, 0xd1, 0xcd, 0xa1, 0x27, 0x25, 0xd2, 0x31, 0x8b, 0x1e, 0xb6, 0x54, 0x0f, 0x55,
	0x21, 0xe8, 0x36, 0xc4, 0x1d, 0x02, 0x87, 0x3f, 0x66, 0xc0, 0xd6, 0x94, 0x89, 0xf4, 0x03, 0xd0,
	0xc0, 0x34, 0x8e, 0x91, 0x20, 0x19, 0x8a, 0x3d, 0x31, 0x60, 0x44, 0xde, 0xcc, 0xfa, 0xe4, 0x26,
	0x2b, 0x00, 0xe8, 0xae, 0x8c, 0x23, 0xdd, 0x01, 0x23, 0xfa, 0x67, 0x0d, 0x34, 0xa7, 0x7b, 0x61,
	0xcc, 0xc8, 0xab, 0xde, 0xb4, 0xd4, 0xe3, 0xb2, 0xca, 0xc7, 0x65, 0x75, 0xcb, 0xc7, 0xe5, 0xec,
	0x14, 0x33, 0x6d, 0xff, 0xc9, 0x57, 0x78, 0xf1, 0xd3, 0xd4, 0x5c, 0x63, 0x9a, 0xb5, 0xfa, 0x19,
	0x68, 0x44, 0xa9, 0x20, 0x19, 0xe1, 0xc2, 0x3b, 0x45, 0x58, 0xd0, 0xcc, 0x98, 0x95, 0xe3, 0x1c,
	0xe5, 0x15, 0xbe, 0x0f, 0xcd, 0x27, 0x61, 0x24, 0xde, 0xf7, 0x7c, 0x0b, 0xd3, 0xc4, 0x2e, 0xde,
	0xb2, 0xfa, 0xd9, 0xe1, 0xc1, 0x47, 0x3b, 0x1f, 0x8f, 0x5b, 0x87, 0x04, 0x8f, 0x87, 0xaf, 0xc8,
	0x41, 0x77, 0xa5, 0x8c, 0xbc, 0x56, 0x81, 0xaf, 0x1a, 0xd8, 0xbc, 0x77, 0x53, 0xff, 0xc6, 0xdb,
	0x33, 0xd0, 0xa8, 0xac, 0x58, 0xfa, 0xf9, 0x77, 0x13, 0x1d, 0xa7, 0x62, 0x5c, 0xb2, 0x22, 0x07,
	0xdd, 0x95, 0xbb, 0x17, 0xc6, 0x79, 0x75, 0x79, 0xdd, 0xd2, 0xae, 0xae, 0x5b, 0xda, 0xaf, 0xeb,
	0x96, 0x76, 0x71, 0xd3, 0xaa, 0x5d, 0xdd, 0xb4, 0x6a, 0xdf, 0x6e, 0x5a, 0xb5, 0x77, 0xdb, 0x13,
	0xb5, 0x62, 0x9c, 0x92, 0xc4, 0xfe, 0xc0, 0x06, 0xe7, 0xf6, 0x79, 0xfe, 0x97, 0xa7, 0x4a, 0xf9,
	0x0b, 0x72, 0xbd, 0xcf, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x1b, 0xc1, 0xd7, 0x80, 0x05,
	0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TotalPrincipals) > 0 {
		for iNdEx := len(m.TotalPrincipals) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TotalPrincipals[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.PreviousAccumulationTimes) > 0 {
		for iNdEx := len(m.PreviousAccumulationTimes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PreviousAccumulationTimes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.GovDenom) > 0 {
		i -= len(m.GovDenom)
		copy(dAtA[i:], m.GovDenom)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.GovDenom)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.DebtDenom) > 0 {
		i -= len(m.DebtDenom)
		copy(dAtA[i:], m.DebtDenom)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.DebtDenom)))
		i--
		dAtA[i] = 0x2a
	}
	if m.StartingCdpId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.StartingCdpId))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Deposits) > 0 {
		for iNdEx := len(m.Deposits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Deposits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Cdps) > 0 {
		for iNdEx := len(m.Cdps) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Cdps[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *GenesisAccumulationTime) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisAccumulationTime) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisAccumulationTime) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.InterestFactor.Size()
		i -= size
		if _, err := m.InterestFactor.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.PreviousAccumulationTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.PreviousAccumulationTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintGenesis(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x12
	if len(m.CollateralType) > 0 {
		i -= len(m.CollateralType)
		copy(dAtA[i:], m.CollateralType)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.CollateralType)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GenesisTotalPrincipal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisTotalPrincipal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisTotalPrincipal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TotalPrincipal.Size()
		i -= size
		if _, err := m.TotalPrincipal.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.CollateralType) > 0 {
		i -= len(m.CollateralType)
		copy(dAtA[i:], m.CollateralType)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.CollateralType)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Cdps) > 0 {
		for _, e := range m.Cdps {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Deposits) > 0 {
		for _, e := range m.Deposits {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.StartingCdpId != 0 {
		n += 1 + sovGenesis(uint64(m.StartingCdpId))
	}
	l = len(m.DebtDenom)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.GovDenom)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.PreviousAccumulationTimes) > 0 {
		for _, e := range m.PreviousAccumulationTimes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.TotalPrincipals) > 0 {
		for _, e := range m.TotalPrincipals {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GenesisAccumulationTime) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CollateralType)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.PreviousAccumulationTime)
	n += 1 + l + sovGenesis(uint64(l))
	l = m.InterestFactor.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *GenesisTotalPrincipal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CollateralType)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.TotalPrincipal.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cdps", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cdps = append(m.Cdps, Cdp{})
			if err := m.Cdps[len(m.Cdps)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Deposits = append(m.Deposits, Deposit{})
			if err := m.Deposits[len(m.Deposits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartingCdpId", wireType)
			}
			m.StartingCdpId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartingCdpId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DebtDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DebtDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GovDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GovDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousAccumulationTimes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PreviousAccumulationTimes = append(m.PreviousAccumulationTimes, GenesisAccumulationTime{})
			if err := m.PreviousAccumulationTimes[len(m.PreviousAccumulationTimes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalPrincipals", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TotalPrincipals = append(m.TotalPrincipals, GenesisTotalPrincipal{})
			if err := m.TotalPrincipals[len(m.TotalPrincipals)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisAccumulationTime) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisAccumulationTime: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisAccumulationTime: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollateralType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousAccumulationTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.PreviousAccumulationTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InterestFactor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InterestFactor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisTotalPrincipal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisTotalPrincipal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisTotalPrincipal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollateralType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalPrincipal", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalPrincipal.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)