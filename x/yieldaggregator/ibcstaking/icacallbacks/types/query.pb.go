// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: icacallbacks/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	proto "github.com/cosmos/gogoproto/proto"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5823c9776c03825e, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5823c9776c03825e, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryGetCallbackDataRequest struct {
	CallbackKey string `protobuf:"bytes,1,opt,name=callbackKey,proto3" json:"callbackKey,omitempty"`
}

func (m *QueryGetCallbackDataRequest) Reset()         { *m = QueryGetCallbackDataRequest{} }
func (m *QueryGetCallbackDataRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetCallbackDataRequest) ProtoMessage()    {}
func (*QueryGetCallbackDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5823c9776c03825e, []int{2}
}
func (m *QueryGetCallbackDataRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetCallbackDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetCallbackDataRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetCallbackDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetCallbackDataRequest.Merge(m, src)
}
func (m *QueryGetCallbackDataRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetCallbackDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetCallbackDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetCallbackDataRequest proto.InternalMessageInfo

func (m *QueryGetCallbackDataRequest) GetCallbackKey() string {
	if m != nil {
		return m.CallbackKey
	}
	return ""
}

type QueryGetCallbackDataResponse struct {
	CallbackData CallbackData `protobuf:"bytes,1,opt,name=callbackData,proto3" json:"callbackData"`
}

func (m *QueryGetCallbackDataResponse) Reset()         { *m = QueryGetCallbackDataResponse{} }
func (m *QueryGetCallbackDataResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetCallbackDataResponse) ProtoMessage()    {}
func (*QueryGetCallbackDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5823c9776c03825e, []int{3}
}
func (m *QueryGetCallbackDataResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetCallbackDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetCallbackDataResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetCallbackDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetCallbackDataResponse.Merge(m, src)
}
func (m *QueryGetCallbackDataResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetCallbackDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetCallbackDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetCallbackDataResponse proto.InternalMessageInfo

func (m *QueryGetCallbackDataResponse) GetCallbackData() CallbackData {
	if m != nil {
		return m.CallbackData
	}
	return CallbackData{}
}

type QueryAllCallbackDataRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllCallbackDataRequest) Reset()         { *m = QueryAllCallbackDataRequest{} }
func (m *QueryAllCallbackDataRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllCallbackDataRequest) ProtoMessage()    {}
func (*QueryAllCallbackDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5823c9776c03825e, []int{4}
}
func (m *QueryAllCallbackDataRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllCallbackDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllCallbackDataRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllCallbackDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllCallbackDataRequest.Merge(m, src)
}
func (m *QueryAllCallbackDataRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllCallbackDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllCallbackDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllCallbackDataRequest proto.InternalMessageInfo

func (m *QueryAllCallbackDataRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllCallbackDataResponse struct {
	CallbackData []CallbackData      `protobuf:"bytes,1,rep,name=callbackData,proto3" json:"callbackData"`
	Pagination   *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllCallbackDataResponse) Reset()         { *m = QueryAllCallbackDataResponse{} }
func (m *QueryAllCallbackDataResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllCallbackDataResponse) ProtoMessage()    {}
func (*QueryAllCallbackDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5823c9776c03825e, []int{5}
}
func (m *QueryAllCallbackDataResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllCallbackDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllCallbackDataResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllCallbackDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllCallbackDataResponse.Merge(m, src)
}
func (m *QueryAllCallbackDataResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllCallbackDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllCallbackDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllCallbackDataResponse proto.InternalMessageInfo

func (m *QueryAllCallbackDataResponse) GetCallbackData() []CallbackData {
	if m != nil {
		return m.CallbackData
	}
	return nil
}

func (m *QueryAllCallbackDataResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "stridelabs.stride.icacallbacks.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "stridelabs.stride.icacallbacks.QueryParamsResponse")
	proto.RegisterType((*QueryGetCallbackDataRequest)(nil), "stridelabs.stride.icacallbacks.QueryGetCallbackDataRequest")
	proto.RegisterType((*QueryGetCallbackDataResponse)(nil), "stridelabs.stride.icacallbacks.QueryGetCallbackDataResponse")
	proto.RegisterType((*QueryAllCallbackDataRequest)(nil), "stridelabs.stride.icacallbacks.QueryAllCallbackDataRequest")
	proto.RegisterType((*QueryAllCallbackDataResponse)(nil), "stridelabs.stride.icacallbacks.QueryAllCallbackDataResponse")
}

func init() {
	// proto.RegisterFile("icacallbacks/query.proto", fileDescriptor_5823c9776c03825e)
}

var fileDescriptor_5823c9776c03825e = []byte{
	// 519 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xc1, 0x6a, 0x14, 0x31,
	0x18, 0xc7, 0x37, 0xb5, 0x2e, 0x98, 0x16, 0x84, 0xd8, 0x43, 0x5d, 0xcb, 0xb8, 0xe4, 0x60, 0x55,
	0x34, 0xb1, 0x5b, 0xf0, 0x62, 0x51, 0x5a, 0x17, 0x7b, 0xb0, 0x87, 0xba, 0x82, 0x07, 0x3d, 0xc8,
	0x37, 0xd3, 0x30, 0x0e, 0x66, 0x27, 0xd3, 0x4d, 0xb6, 0xb8, 0x88, 0x17, 0x9f, 0x40, 0xf0, 0x0d,
	0x7c, 0x11, 0x2f, 0x0a, 0xbd, 0x59, 0xf0, 0xe2, 0x49, 0x64, 0xd7, 0x07, 0x91, 0x4d, 0xb2, 0x9a,
	0xa1, 0x53, 0x77, 0x2b, 0xde, 0x42, 0xe6, 0xfb, 0xff, 0xbf, 0xff, 0x2f, 0xf9, 0x32, 0x78, 0x39,
	0x4b, 0x20, 0x01, 0x29, 0x63, 0x48, 0x5e, 0x6a, 0xbe, 0xdf, 0x17, 0xbd, 0x01, 0x2b, 0x7a, 0xca,
	0x28, 0x12, 0x69, 0xd3, 0xcb, 0xf6, 0x84, 0x84, 0x58, 0x33, 0xb7, 0x64, 0x61, 0x6d, 0x63, 0x29,
	0x55, 0xa9, 0xb2, 0xa5, 0x7c, 0xbc, 0x72, 0xaa, 0xc6, 0x4a, 0xaa, 0x54, 0x2a, 0x05, 0x87, 0x22,
	0xe3, 0x90, 0xe7, 0xca, 0x80, 0xc9, 0x54, 0xae, 0xfd, 0xd7, 0xeb, 0x89, 0xd2, 0x5d, 0xa5, 0x79,
	0x0c, 0x5a, 0xb8, 0x66, 0xfc, 0x60, 0x2d, 0x16, 0x06, 0xd6, 0x78, 0x01, 0x69, 0x96, 0xdb, 0x62,
	0x5f, 0x7b, 0xb1, 0x94, 0xac, 0x80, 0x1e, 0x74, 0x27, 0x36, 0xcd, 0xd2, 0xa7, 0xc9, 0xea, 0xf9,
	0x1e, 0x18, 0x70, 0x15, 0x74, 0x09, 0x93, 0x47, 0x63, 0xfb, 0x5d, 0x2b, 0xeb, 0x88, 0xfd, 0xbe,
	0xd0, 0x86, 0x3e, 0xc3, 0x17, 0x4a, 0xbb, 0xba, 0x50, 0xb9, 0x16, 0xa4, 0x8d, 0xeb, 0xce, 0x7e,
	0x19, 0x35, 0xd1, 0xd5, 0x85, 0xd6, 0x15, 0xf6, 0x77, 0x74, 0xe6, 0xf4, 0x5b, 0xf3, 0x87, 0xdf,
	0x2f, 0xd7, 0x3a, 0x5e, 0x4b, 0xef, 0xe1, 0x4b, 0xd6, 0x7c, 0x5b, 0x98, 0xfb, 0xbe, 0xb2, 0x0d,
	0x06, 0x7c, 0x6f, 0xd2, 0xc4, 0x0b, 0x13, 0x83, 0x87, 0x62, 0x60, 0x3b, 0x9d, 0xeb, 0x84, 0x5b,
	0xf4, 0x00, 0xaf, 0x54, 0x1b, 0xf8, 0x98, 0x4f, 0xf0, 0x62, 0x12, 0xec, 0xfb, 0xb0, 0x37, 0xa6,
	0x85, 0x0d, 0xbd, 0x7c, 0xe4, 0x92, 0x0f, 0x15, 0x3e, 0xf8, 0xa6, 0x94, 0x55, 0xc1, 0x1f, 0x60,
	0xfc, 0xe7, 0x6e, 0x7e, 0x9f, 0x90, 0xbb, 0x48, 0x36, 0xbe, 0x48, 0xe6, 0xa6, 0xc6, 0x5f, 0x24,
	0xdb, 0x85, 0x54, 0x78, 0x6d, 0x27, 0x50, 0xd2, 0x8f, 0xc8, 0xf3, 0x1d, 0xeb, 0x73, 0x22, 0xdf,
	0x99, 0xff, 0xc1, 0x47, 0xb6, 0x4b, 0x00, 0x73, 0x16, 0x60, 0x75, 0x2a, 0x80, 0x0b, 0x15, 0x12,
	0xb4, 0x3e, 0xcf, 0xe3, 0xb3, 0x96, 0x80, 0x7c, 0x40, 0xb8, 0xee, 0x86, 0x80, 0xb4, 0xa6, 0xe5,
	0x3b, 0x3e, 0x87, 0x8d, 0xf5, 0x53, 0x69, 0x5c, 0x12, 0xca, 0xdf, 0x7e, 0xfd, 0xf9, 0x7e, 0xee,
	0x1a, 0x59, 0xe5, 0x8f, 0xad, 0xe2, 0xe6, 0x0e, 0xc4, 0x9a, 0x3b, 0x35, 0xaf, 0x78, 0x2b, 0xe4,
	0x0b, 0xc2, 0x8b, 0xe1, 0xe1, 0x90, 0x3b, 0x33, 0xb5, 0xad, 0x9e, 0xdf, 0xc6, 0xc6, 0xbf, 0x89,
	0x7d, 0xf8, 0xb6, 0x0d, 0x7f, 0x97, 0x6c, 0x4c, 0x0d, 0x5f, 0x7a, 0xcd, 0xfc, 0x75, 0xf0, 0x40,
	0xde, 0x90, 0x4f, 0x08, 0x9f, 0x0f, 0xed, 0x37, 0xa5, 0x9c, 0x11, 0xaa, 0x7a, 0xb6, 0x67, 0x84,
	0x3a, 0x61, 0x60, 0xe9, 0x6d, 0x0b, 0x75, 0x8b, 0xb0, 0xd3, 0x41, 0x6d, 0xed, 0x1c, 0x0e, 0x23,
	0x74, 0x34, 0x8c, 0xd0, 0x8f, 0x61, 0x84, 0xde, 0x8d, 0xa2, 0xda, 0xd1, 0x28, 0xaa, 0x7d, 0x1b,
	0x45, 0xb5, 0xa7, 0xad, 0x34, 0x33, 0x2f, 0xfa, 0x31, 0x4b, 0x54, 0xb7, 0xca, 0xf3, 0x55, 0xd9,
	0xd5, 0x0c, 0x0a, 0xa1, 0xe3, 0xba, 0xfd, 0xe3, 0xad, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xe9,
	0xe0, 0xba, 0xfb, 0xca, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a CallbackData by index.
	CallbackData(ctx context.Context, in *QueryGetCallbackDataRequest, opts ...grpc.CallOption) (*QueryGetCallbackDataResponse, error)
	// Queries a list of CallbackData items.
	CallbackDataAll(ctx context.Context, in *QueryAllCallbackDataRequest, opts ...grpc.CallOption) (*QueryAllCallbackDataResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/stridelabs.stride.icacallbacks.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CallbackData(ctx context.Context, in *QueryGetCallbackDataRequest, opts ...grpc.CallOption) (*QueryGetCallbackDataResponse, error) {
	out := new(QueryGetCallbackDataResponse)
	err := c.cc.Invoke(ctx, "/stridelabs.stride.icacallbacks.Query/CallbackData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CallbackDataAll(ctx context.Context, in *QueryAllCallbackDataRequest, opts ...grpc.CallOption) (*QueryAllCallbackDataResponse, error) {
	out := new(QueryAllCallbackDataResponse)
	err := c.cc.Invoke(ctx, "/stridelabs.stride.icacallbacks.Query/CallbackDataAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a CallbackData by index.
	CallbackData(context.Context, *QueryGetCallbackDataRequest) (*QueryGetCallbackDataResponse, error)
	// Queries a list of CallbackData items.
	CallbackDataAll(context.Context, *QueryAllCallbackDataRequest) (*QueryAllCallbackDataResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) CallbackData(ctx context.Context, req *QueryGetCallbackDataRequest) (*QueryGetCallbackDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CallbackData not implemented")
}
func (*UnimplementedQueryServer) CallbackDataAll(ctx context.Context, req *QueryAllCallbackDataRequest) (*QueryAllCallbackDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CallbackDataAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stridelabs.stride.icacallbacks.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CallbackData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetCallbackDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CallbackData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stridelabs.stride.icacallbacks.Query/CallbackData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CallbackData(ctx, req.(*QueryGetCallbackDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CallbackDataAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllCallbackDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CallbackDataAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stridelabs.stride.icacallbacks.Query/CallbackDataAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CallbackDataAll(ctx, req.(*QueryAllCallbackDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stridelabs.stride.icacallbacks.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "CallbackData",
			Handler:    _Query_CallbackData_Handler,
		},
		{
			MethodName: "CallbackDataAll",
			Handler:    _Query_CallbackDataAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "icacallbacks/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryGetCallbackDataRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetCallbackDataRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetCallbackDataRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CallbackKey) > 0 {
		i -= len(m.CallbackKey)
		copy(dAtA[i:], m.CallbackKey)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.CallbackKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetCallbackDataResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetCallbackDataResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetCallbackDataResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.CallbackData.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAllCallbackDataRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllCallbackDataRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllCallbackDataRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllCallbackDataResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllCallbackDataResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllCallbackDataResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.CallbackData) > 0 {
		for iNdEx := len(m.CallbackData) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CallbackData[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryGetCallbackDataRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CallbackKey)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetCallbackDataResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.CallbackData.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllCallbackDataRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllCallbackDataResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.CallbackData) > 0 {
		for _, e := range m.CallbackData {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetCallbackDataRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetCallbackDataRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetCallbackDataRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CallbackKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CallbackKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetCallbackDataResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetCallbackDataResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetCallbackDataResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CallbackData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CallbackData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllCallbackDataRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllCallbackDataRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllCallbackDataRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllCallbackDataResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllCallbackDataResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllCallbackDataResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CallbackData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CallbackData = append(m.CallbackData, CallbackData{})
			if err := m.CallbackData[len(m.CallbackData)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
