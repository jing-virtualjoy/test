// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/alanchchen/grpc-lb-istio/api/identity.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type WhoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WhoRequest) Reset()         { *m = WhoRequest{} }
func (m *WhoRequest) String() string { return proto.CompactTextString(m) }
func (*WhoRequest) ProtoMessage()    {}
func (*WhoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_identity_4d7288c83f5dd9f6, []int{0}
}
func (m *WhoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WhoRequest.Unmarshal(m, b)
}
func (m *WhoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WhoRequest.Marshal(b, m, deterministic)
}
func (dst *WhoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WhoRequest.Merge(dst, src)
}
func (m *WhoRequest) XXX_Size() int {
	return xxx_messageInfo_WhoRequest.Size(m)
}
func (m *WhoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WhoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WhoRequest proto.InternalMessageInfo

type WhoResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WhoResponse) Reset()         { *m = WhoResponse{} }
func (m *WhoResponse) String() string { return proto.CompactTextString(m) }
func (*WhoResponse) ProtoMessage()    {}
func (*WhoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_identity_4d7288c83f5dd9f6, []int{1}
}
func (m *WhoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WhoResponse.Unmarshal(m, b)
}
func (m *WhoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WhoResponse.Marshal(b, m, deterministic)
}
func (dst *WhoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WhoResponse.Merge(dst, src)
}
func (m *WhoResponse) XXX_Size() int {
	return xxx_messageInfo_WhoResponse.Size(m)
}
func (m *WhoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WhoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WhoResponse proto.InternalMessageInfo

func (m *WhoResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*WhoRequest)(nil), "api.WhoRequest")
	proto.RegisterType((*WhoResponse)(nil), "api.WhoResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IdentityClient is the client API for Identity service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IdentityClient interface {
	Who(ctx context.Context, in *WhoRequest, opts ...grpc.CallOption) (*WhoResponse, error)
}

type identityClient struct {
	cc *grpc.ClientConn
}

func NewIdentityClient(cc *grpc.ClientConn) IdentityClient {
	return &identityClient{cc}
}

func (c *identityClient) Who(ctx context.Context, in *WhoRequest, opts ...grpc.CallOption) (*WhoResponse, error) {
	out := new(WhoResponse)
	err := c.cc.Invoke(ctx, "/api.Identity/Who", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentityServer is the server API for Identity service.
type IdentityServer interface {
	Who(context.Context, *WhoRequest) (*WhoResponse, error)
}

func RegisterIdentityServer(s *grpc.Server, srv IdentityServer) {
	s.RegisterService(&_Identity_serviceDesc, srv)
}

func _Identity_Who_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WhoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).Who(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Identity/Who",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).Who(ctx, req.(*WhoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Identity_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Identity",
	HandlerType: (*IdentityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Who",
			Handler:    _Identity_Who_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/alanchchen/grpc-lb-istio/api/identity.proto",
}

func init() {
	proto.RegisterFile("github.com/alanchchen/grpc-lb-istio/api/identity.proto", fileDescriptor_identity_4d7288c83f5dd9f6)
}

var fileDescriptor_identity_4d7288c83f5dd9f6 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4b, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcc, 0x49, 0xcc, 0x4b, 0xce, 0x48, 0xce, 0x48,
	0xcd, 0xd3, 0x4f, 0x2f, 0x2a, 0x48, 0xd6, 0xcd, 0x49, 0xd2, 0xcd, 0x2c, 0x2e, 0xc9, 0xcc, 0xd7,
	0x4f, 0x2c, 0xc8, 0xd4, 0xcf, 0x4c, 0x49, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9, 0xd4, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x94, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0x05,
	0xcb, 0x27, 0xe6, 0xe5, 0xe5, 0x97, 0x24, 0x96, 0x64, 0xe6, 0xe7, 0x15, 0x43, 0x94, 0x28, 0xf1,
	0x70, 0x71, 0x85, 0x67, 0xe4, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x28, 0x29, 0x72, 0x71,
	0x83, 0x79, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9,
	0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x91, 0x27, 0x17, 0x87, 0x27, 0xd4, 0x16,
	0x21, 0x5b, 0x2e, 0xe6, 0xf0, 0x8c, 0x7c, 0x21, 0x7e, 0xbd, 0xc4, 0x82, 0x4c, 0x3d, 0x84, 0x31,
	0x52, 0x02, 0x08, 0x01, 0x88, 0x49, 0x4a, 0xc2, 0x4d, 0x97, 0x9f, 0x4c, 0x66, 0xe2, 0x15, 0xe2,
	0x00, 0x3b, 0xa3, 0x3c, 0x23, 0xdf, 0x8a, 0x51, 0xcb, 0x89, 0x35, 0x0a, 0xe4, 0xc0, 0x24, 0x36,
	0xb0, 0x4b, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x24, 0xe8, 0x2e, 0xf3, 0xe6, 0x00, 0x00,
	0x00,
}
