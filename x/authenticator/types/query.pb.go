// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/authenticator/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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
	return fileDescriptor_4042cdb0ec4643b6, []int{0}
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
	return fileDescriptor_4042cdb0ec4643b6, []int{1}
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

// MsgGetAuthenticatorsRequest defines the Msg/GetAuthenticators request type.
type GetAuthenticatorsRequest struct {
	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (m *GetAuthenticatorsRequest) Reset()         { *m = GetAuthenticatorsRequest{} }
func (m *GetAuthenticatorsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAuthenticatorsRequest) ProtoMessage()    {}
func (*GetAuthenticatorsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4042cdb0ec4643b6, []int{2}
}
func (m *GetAuthenticatorsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetAuthenticatorsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetAuthenticatorsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetAuthenticatorsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAuthenticatorsRequest.Merge(m, src)
}
func (m *GetAuthenticatorsRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetAuthenticatorsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAuthenticatorsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAuthenticatorsRequest proto.InternalMessageInfo

func (m *GetAuthenticatorsRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

// MsgGetAuthenticatorsResponse defines the Msg/GetAuthenticators response type.
type GetAuthenticatorsResponse struct {
	AccountAuthenticators []*AccountAuthenticator `protobuf:"bytes,1,rep,name=account_authenticators,json=accountAuthenticators,proto3" json:"account_authenticators,omitempty"`
}

func (m *GetAuthenticatorsResponse) Reset()         { *m = GetAuthenticatorsResponse{} }
func (m *GetAuthenticatorsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAuthenticatorsResponse) ProtoMessage()    {}
func (*GetAuthenticatorsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4042cdb0ec4643b6, []int{3}
}
func (m *GetAuthenticatorsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetAuthenticatorsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetAuthenticatorsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetAuthenticatorsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAuthenticatorsResponse.Merge(m, src)
}
func (m *GetAuthenticatorsResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetAuthenticatorsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAuthenticatorsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAuthenticatorsResponse proto.InternalMessageInfo

func (m *GetAuthenticatorsResponse) GetAccountAuthenticators() []*AccountAuthenticator {
	if m != nil {
		return m.AccountAuthenticators
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "osmosis.authenticator.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "osmosis.authenticator.QueryParamsResponse")
	proto.RegisterType((*GetAuthenticatorsRequest)(nil), "osmosis.authenticator.GetAuthenticatorsRequest")
	proto.RegisterType((*GetAuthenticatorsResponse)(nil), "osmosis.authenticator.GetAuthenticatorsResponse")
}

func init() { proto.RegisterFile("osmosis/authenticator/query.proto", fileDescriptor_4042cdb0ec4643b6) }

var fileDescriptor_4042cdb0ec4643b6 = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x31, 0x6b, 0xdb, 0x40,
	0x1c, 0xc5, 0x25, 0xb7, 0x75, 0xe9, 0x79, 0xea, 0xd5, 0x2e, 0xae, 0xa8, 0xe5, 0x56, 0x50, 0x70,
	0x5d, 0xaa, 0xab, 0xdc, 0x42, 0x31, 0x9d, 0xec, 0xa5, 0x6b, 0xa2, 0x6c, 0x59, 0xc2, 0x49, 0x39,
	0x64, 0x81, 0xa5, 0x93, 0x75, 0x27, 0x13, 0x13, 0x42, 0x20, 0x43, 0xe6, 0x40, 0x3e, 0x48, 0xbe,
	0x43, 0x26, 0x8f, 0x86, 0x2c, 0x99, 0x42, 0xb0, 0xf3, 0x41, 0x82, 0x4f, 0xe7, 0x10, 0xc5, 0x92,
	0xf1, 0x26, 0x9d, 0x7e, 0xef, 0xff, 0xde, 0xff, 0x9d, 0xc0, 0x57, 0xca, 0x02, 0xca, 0x7c, 0x86,
	0x70, 0xc2, 0x07, 0x24, 0xe4, 0xbe, 0x8b, 0x39, 0x8d, 0xd1, 0x28, 0x21, 0xf1, 0xc4, 0x8c, 0x62,
	0xca, 0x29, 0xac, 0x49, 0xc4, 0xcc, 0x20, 0x5a, 0xd5, 0xa3, 0x1e, 0x15, 0x04, 0x5a, 0x3e, 0xa5,
	0xb0, 0xf6, 0xd9, 0xa3, 0xd4, 0x1b, 0x12, 0x84, 0x23, 0x1f, 0xe1, 0x30, 0xa4, 0x1c, 0x73, 0x9f,
	0x86, 0x4c, 0x7e, 0x6d, 0xbb, 0x62, 0x16, 0x72, 0x30, 0x23, 0xa9, 0x07, 0x1a, 0x5b, 0x0e, 0xe1,
	0xd8, 0x42, 0x11, 0xf6, 0xfc, 0x50, 0xc0, 0x92, 0x35, 0xf2, 0x93, 0x45, 0x38, 0xc6, 0x01, 0xdb,
	0xcc, 0x04, 0xf4, 0x90, 0x0c, 0x25, 0x63, 0x54, 0x01, 0xdc, 0x5d, 0x3a, 0xed, 0x08, 0xa1, 0x4d,
	0x46, 0x09, 0x61, 0xdc, 0xb0, 0xc1, 0x87, 0xcc, 0x29, 0x8b, 0x68, 0xc8, 0x08, 0xfc, 0x07, 0xca,
	0xa9, 0x41, 0x5d, 0xfd, 0xa2, 0xb6, 0x2a, 0x9d, 0x86, 0x99, 0xbb, 0xbc, 0x99, 0xca, 0xfa, 0xaf,
	0xa7, 0x77, 0x4d, 0xc5, 0x96, 0x12, 0xe3, 0x0f, 0xa8, 0xff, 0x27, 0xbc, 0xf7, 0x1c, 0x5c, 0xf9,
	0xc1, 0x3a, 0x78, 0x8b, 0x5d, 0x97, 0x26, 0x21, 0x17, 0x93, 0xdf, 0xd9, 0xab, 0x57, 0xe3, 0x14,
	0x7c, 0xca, 0x51, 0xc9, 0x3c, 0x0e, 0xf8, 0x28, 0xb9, 0x83, 0x4c, 0x80, 0x65, 0xbe, 0x57, 0xad,
	0x4a, 0xe7, 0x47, 0x41, 0xbe, 0x5e, 0x2a, 0xca, 0x4c, 0xb5, 0x6b, 0x38, 0xe7, 0x94, 0x75, 0xae,
	0x4b, 0xe0, 0x8d, 0xe8, 0x02, 0x9e, 0xab, 0xa0, 0x9c, 0x6e, 0x06, 0xbf, 0x17, 0x0c, 0x5e, 0xaf,
	0x52, 0x6b, 0x6f, 0x83, 0xa6, 0xfb, 0x18, 0xdf, 0xce, 0x6e, 0x1e, 0x2e, 0x4b, 0x4d, 0xd8, 0x40,
	0x9b, 0x6e, 0x17, 0x5e, 0xa9, 0xe0, 0xfd, 0x5a, 0x29, 0x10, 0x15, 0x18, 0x15, 0x95, 0xae, 0xfd,
	0xda, 0x5e, 0x20, 0xf3, 0xfd, 0x15, 0xf9, 0x2c, 0x88, 0x0a, 0xf2, 0x65, 0x2f, 0x01, 0x1d, 0xcb,
	0x46, 0x4f, 0xfa, 0x7b, 0xd3, 0xb9, 0xae, 0xce, 0xe6, 0xba, 0x7a, 0x3f, 0xd7, 0xd5, 0x8b, 0x85,
	0xae, 0xcc, 0x16, 0xba, 0x72, 0xbb, 0xd0, 0x95, 0xfd, 0xae, 0xe7, 0xf3, 0x41, 0xe2, 0x98, 0x2e,
	0x0d, 0x56, 0x43, 0x7f, 0x0e, 0xb1, 0xc3, 0x9e, 0x1c, 0xc6, 0x56, 0x17, 0x1d, 0xbd, 0xf0, 0xe1,
	0x93, 0x88, 0x30, 0xa7, 0x2c, 0xfe, 0xe0, 0xdf, 0x8f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x96, 0xaa,
	0x08, 0xf4, 0xa5, 0x03, 0x00, 0x00,
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
	GetAuthenticators(ctx context.Context, in *GetAuthenticatorsRequest, opts ...grpc.CallOption) (*GetAuthenticatorsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/osmosis.authenticator.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetAuthenticators(ctx context.Context, in *GetAuthenticatorsRequest, opts ...grpc.CallOption) (*GetAuthenticatorsResponse, error) {
	out := new(GetAuthenticatorsResponse)
	err := c.cc.Invoke(ctx, "/osmosis.authenticator.Query/GetAuthenticators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	GetAuthenticators(context.Context, *GetAuthenticatorsRequest) (*GetAuthenticatorsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) GetAuthenticators(ctx context.Context, req *GetAuthenticatorsRequest) (*GetAuthenticatorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthenticators not implemented")
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
		FullMethod: "/osmosis.authenticator.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetAuthenticators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthenticatorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetAuthenticators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.authenticator.Query/GetAuthenticators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetAuthenticators(ctx, req.(*GetAuthenticatorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "osmosis.authenticator.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "GetAuthenticators",
			Handler:    _Query_GetAuthenticators_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "osmosis/authenticator/query.proto",
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

func (m *GetAuthenticatorsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetAuthenticatorsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetAuthenticatorsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetAuthenticatorsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetAuthenticatorsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetAuthenticatorsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AccountAuthenticators) > 0 {
		for iNdEx := len(m.AccountAuthenticators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AccountAuthenticators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *GetAuthenticatorsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *GetAuthenticatorsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AccountAuthenticators) > 0 {
		for _, e := range m.AccountAuthenticators {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
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
func (m *GetAuthenticatorsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetAuthenticatorsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetAuthenticatorsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
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
			m.Account = string(dAtA[iNdEx:postIndex])
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
func (m *GetAuthenticatorsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetAuthenticatorsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetAuthenticatorsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountAuthenticators", wireType)
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
			m.AccountAuthenticators = append(m.AccountAuthenticators, &AccountAuthenticator{})
			if err := m.AccountAuthenticators[len(m.AccountAuthenticators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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