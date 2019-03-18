// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dnsmasqmgr.proto

package dnsmasqmgr

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Key int32

const (
	Key_HOSTNAME Key = 0
	Key_MACADDR  Key = 1
	Key_IPADDR   Key = 2
)

var Key_name = map[int32]string{
	0: "HOSTNAME",
	1: "MACADDR",
	2: "IPADDR",
}

var Key_value = map[string]int32{
	"HOSTNAME": 0,
	"MACADDR":  1,
	"IPADDR":   2,
}

func (x Key) String() string {
	return proto.EnumName(Key_name, int32(x))
}

func (Key) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b3815698c51f4a73, []int{0}
}

type Match int32

const (
	Match_NONE    Match = 0
	Match_PARTIAL Match = 1
	Match_FULL    Match = 2
)

var Match_name = map[int32]string{
	0: "NONE",
	1: "PARTIAL",
	2: "FULL",
}

var Match_value = map[string]int32{
	"NONE":    0,
	"PARTIAL": 1,
	"FULL":    2,
}

func (x Match) String() string {
	return proto.EnumName(Match_name, int32(x))
}

func (Match) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b3815698c51f4a73, []int{1}
}

type Error int32

const (
	Error_SUCCESS   Error = 0
	Error_NOTFOUND  Error = 1
	Error_DUPLICATE Error = 2
	Error_MISMATCH  Error = 3
)

var Error_name = map[int32]string{
	0: "SUCCESS",
	1: "NOTFOUND",
	2: "DUPLICATE",
	3: "MISMATCH",
}

var Error_value = map[string]int32{
	"SUCCESS":   0,
	"NOTFOUND":  1,
	"DUPLICATE": 2,
	"MISMATCH":  3,
}

func (x Error) String() string {
	return proto.EnumName(Error_name, int32(x))
}

func (Error) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b3815698c51f4a73, []int{2}
}

type Address struct {
	Hostname             string   `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Macaddr              string   `protobuf:"bytes,2,opt,name=macaddr,proto3" json:"macaddr,omitempty"`
	Ipaddr               string   `protobuf:"bytes,3,opt,name=ipaddr,proto3" json:"ipaddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3815698c51f4a73, []int{0}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *Address) GetMacaddr() string {
	if m != nil {
		return m.Macaddr
	}
	return ""
}

func (m *Address) GetIpaddr() string {
	if m != nil {
		return m.Ipaddr
	}
	return ""
}

type AddressRequest struct {
	Addr                 *Address `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Key                  Key      `protobuf:"varint,2,opt,name=key,proto3,enum=dnsmasqmgr.Key" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddressRequest) Reset()         { *m = AddressRequest{} }
func (m *AddressRequest) String() string { return proto.CompactTextString(m) }
func (*AddressRequest) ProtoMessage()    {}
func (*AddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3815698c51f4a73, []int{1}
}

func (m *AddressRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressRequest.Unmarshal(m, b)
}
func (m *AddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressRequest.Marshal(b, m, deterministic)
}
func (m *AddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressRequest.Merge(m, src)
}
func (m *AddressRequest) XXX_Size() int {
	return xxx_messageInfo_AddressRequest.Size(m)
}
func (m *AddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddressRequest proto.InternalMessageInfo

func (m *AddressRequest) GetAddr() *Address {
	if m != nil {
		return m.Addr
	}
	return nil
}

func (m *AddressRequest) GetKey() Key {
	if m != nil {
		return m.Key
	}
	return Key_HOSTNAME
}

type AddressReply struct {
	Addr                 *Address `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Match                Match    `protobuf:"varint,2,opt,name=match,proto3,enum=dnsmasqmgr.Match" json:"match,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddressReply) Reset()         { *m = AddressReply{} }
func (m *AddressReply) String() string { return proto.CompactTextString(m) }
func (*AddressReply) ProtoMessage()    {}
func (*AddressReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3815698c51f4a73, []int{2}
}

func (m *AddressReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressReply.Unmarshal(m, b)
}
func (m *AddressReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressReply.Marshal(b, m, deterministic)
}
func (m *AddressReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressReply.Merge(m, src)
}
func (m *AddressReply) XXX_Size() int {
	return xxx_messageInfo_AddressReply.Size(m)
}
func (m *AddressReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressReply.DiscardUnknown(m)
}

var xxx_messageInfo_AddressReply proto.InternalMessageInfo

func (m *AddressReply) GetAddr() *Address {
	if m != nil {
		return m.Addr
	}
	return nil
}

func (m *AddressReply) GetMatch() Match {
	if m != nil {
		return m.Match
	}
	return Match_NONE
}

type StatusReply struct {
	Addr                 *Address `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Code                 Error    `protobuf:"varint,2,opt,name=code,proto3,enum=dnsmasqmgr.Error" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusReply) Reset()         { *m = StatusReply{} }
func (m *StatusReply) String() string { return proto.CompactTextString(m) }
func (*StatusReply) ProtoMessage()    {}
func (*StatusReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3815698c51f4a73, []int{3}
}

func (m *StatusReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusReply.Unmarshal(m, b)
}
func (m *StatusReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusReply.Marshal(b, m, deterministic)
}
func (m *StatusReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusReply.Merge(m, src)
}
func (m *StatusReply) XXX_Size() int {
	return xxx_messageInfo_StatusReply.Size(m)
}
func (m *StatusReply) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusReply.DiscardUnknown(m)
}

var xxx_messageInfo_StatusReply proto.InternalMessageInfo

func (m *StatusReply) GetAddr() *Address {
	if m != nil {
		return m.Addr
	}
	return nil
}

func (m *StatusReply) GetCode() Error {
	if m != nil {
		return m.Code
	}
	return Error_SUCCESS
}

func init() {
	proto.RegisterEnum("dnsmasqmgr.Key", Key_name, Key_value)
	proto.RegisterEnum("dnsmasqmgr.Match", Match_name, Match_value)
	proto.RegisterEnum("dnsmasqmgr.Error", Error_name, Error_value)
	proto.RegisterType((*Address)(nil), "dnsmasqmgr.Address")
	proto.RegisterType((*AddressRequest)(nil), "dnsmasqmgr.AddressRequest")
	proto.RegisterType((*AddressReply)(nil), "dnsmasqmgr.AddressReply")
	proto.RegisterType((*StatusReply)(nil), "dnsmasqmgr.StatusReply")
}

func init() { proto.RegisterFile("dnsmasqmgr.proto", fileDescriptor_b3815698c51f4a73) }

var fileDescriptor_b3815698c51f4a73 = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xdd, 0x6e, 0xda, 0x30,
	0x18, 0x86, 0x09, 0xff, 0xfd, 0x28, 0xd4, 0xf3, 0xa4, 0x0d, 0x55, 0x9a, 0xb4, 0x45, 0x9a, 0x5a,
	0xa1, 0x89, 0x03, 0x76, 0x01, 0x93, 0x9b, 0x40, 0x8b, 0x9a, 0x84, 0x28, 0x09, 0xda, 0xc9, 0x26,
	0xcd, 0x25, 0x16, 0x74, 0x6d, 0x30, 0xd8, 0x66, 0x52, 0x6e, 0x79, 0x57, 0x31, 0xd9, 0x0d, 0x2c,
	0x07, 0x1c, 0x30, 0xed, 0x2c, 0x9f, 0x9f, 0x57, 0xcf, 0xeb, 0x38, 0x31, 0xa0, 0x74, 0x2d, 0x33,
	0x2a, 0xb7, 0xd9, 0x52, 0x0c, 0x37, 0x82, 0x2b, 0x8e, 0xe1, 0xef, 0x8a, 0xfd, 0x15, 0x5a, 0x24,
	0x4d, 0x05, 0x93, 0x12, 0x5f, 0x42, 0x7b, 0xc5, 0xa5, 0x5a, 0xd3, 0x8c, 0xf5, 0xad, 0xf7, 0xd6,
	0xf5, 0x59, 0x74, 0x98, 0x71, 0x1f, 0x5a, 0x19, 0x5d, 0xd0, 0x34, 0x15, 0xfd, 0xaa, 0x41, 0xfb,
	0x11, 0xbf, 0x81, 0xe6, 0xe3, 0xc6, 0x80, 0x9a, 0x01, 0xc5, 0x64, 0x7f, 0x83, 0x5e, 0x21, 0x8e,
	0xd8, 0x76, 0xc7, 0xa4, 0xc2, 0x57, 0x50, 0x37, 0x39, 0xed, 0xee, 0x8c, 0x5e, 0x0f, 0x4b, 0xfb,
	0xda, 0x27, 0x4d, 0x00, 0x7f, 0x80, 0xda, 0x13, 0xcb, 0x4d, 0x51, 0x6f, 0x74, 0x51, 0xce, 0xdd,
	0xb3, 0x3c, 0xd2, 0xcc, 0xfe, 0x01, 0xe7, 0x07, 0xfb, 0xe6, 0x39, 0x3f, 0xdd, 0x7d, 0x05, 0x8d,
	0x8c, 0xaa, 0xc5, 0xaa, 0xb0, 0xbf, 0x2a, 0x27, 0x7d, 0x0d, 0xa2, 0x17, 0x6e, 0x7f, 0x87, 0x4e,
	0xac, 0xa8, 0xda, 0xfd, 0x6b, 0xc1, 0x47, 0xa8, 0x2f, 0x78, 0xca, 0x8e, 0xf9, 0xc7, 0x42, 0x70,
	0x11, 0x19, 0x3c, 0xf8, 0x04, 0xb5, 0x7b, 0x96, 0xe3, 0x73, 0x68, 0xdf, 0xcd, 0xe2, 0x24, 0x20,
	0xfe, 0x18, 0x55, 0x70, 0x07, 0x5a, 0x3e, 0x71, 0x88, 0xeb, 0x46, 0xc8, 0xc2, 0x00, 0xcd, 0x69,
	0x68, 0x9e, 0xab, 0x83, 0x6b, 0x68, 0x98, 0xcd, 0xe1, 0x36, 0xd4, 0x83, 0x59, 0x50, 0x64, 0x43,
	0x12, 0x25, 0x53, 0xe2, 0x21, 0x4b, 0x2f, 0x4f, 0xe6, 0x9e, 0x87, 0xaa, 0x83, 0x2f, 0xd0, 0x30,
	0x35, 0x9a, 0xc7, 0x73, 0xc7, 0x19, 0xc7, 0x31, 0xaa, 0xe8, 0x9a, 0x60, 0x96, 0x4c, 0x66, 0xf3,
	0xc0, 0x45, 0x16, 0xee, 0xc2, 0x99, 0x3b, 0x0f, 0xbd, 0xa9, 0x43, 0x92, 0x31, 0xaa, 0x6a, 0xe8,
	0x4f, 0x63, 0x9f, 0x24, 0xce, 0x1d, 0xaa, 0x8d, 0x7e, 0x5b, 0xd0, 0x73, 0x83, 0xd8, 0xa7, 0x72,
	0xeb, 0xd3, 0x35, 0x5d, 0x32, 0x81, 0x6f, 0xa1, 0x57, 0x7c, 0xc3, 0xc3, 0xaf, 0x72, 0xec, 0xfd,
	0x5f, 0x22, 0x97, 0x6f, 0xcb, 0xac, 0x74, 0x84, 0x76, 0x05, 0x4f, 0xa0, 0xeb, 0xb2, 0x67, 0xa6,
	0xd8, 0x7f, 0x7a, 0x6e, 0xa1, 0xeb, 0x71, 0xfe, 0xb4, 0xdb, 0x9c, 0xe2, 0xe9, 0x1f, 0x65, 0x46,
	0x74, 0x33, 0x82, 0x77, 0x0b, 0x9e, 0x0d, 0x97, 0x8f, 0x6a, 0xb5, 0x7b, 0x18, 0x66, 0xfc, 0x27,
	0xfd, 0xc5, 0x64, 0x29, 0x7f, 0x73, 0xb1, 0x3f, 0x8a, 0xa5, 0x08, 0xf5, 0xdd, 0x09, 0xad, 0x87,
	0xa6, 0xb9, 0x44, 0x9f, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x04, 0x36, 0x7d, 0x1c, 0x58, 0x03,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DNSMasqManagerClient is the client API for DNSMasqManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DNSMasqManagerClient interface {
	RequestAddress(ctx context.Context, in *AddressRequest, opts ...grpc.CallOption) (*StatusReply, error)
	DeleteAddress(ctx context.Context, in *AddressRequest, opts ...grpc.CallOption) (*StatusReply, error)
	LookupAddress(ctx context.Context, in *AddressRequest, opts ...grpc.CallOption) (*AddressReply, error)
}

type dNSMasqManagerClient struct {
	cc *grpc.ClientConn
}

func NewDNSMasqManagerClient(cc *grpc.ClientConn) DNSMasqManagerClient {
	return &dNSMasqManagerClient{cc}
}

func (c *dNSMasqManagerClient) RequestAddress(ctx context.Context, in *AddressRequest, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/dnsmasqmgr.DNSMasqManager/RequestAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dNSMasqManagerClient) DeleteAddress(ctx context.Context, in *AddressRequest, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/dnsmasqmgr.DNSMasqManager/DeleteAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dNSMasqManagerClient) LookupAddress(ctx context.Context, in *AddressRequest, opts ...grpc.CallOption) (*AddressReply, error) {
	out := new(AddressReply)
	err := c.cc.Invoke(ctx, "/dnsmasqmgr.DNSMasqManager/LookupAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DNSMasqManagerServer is the server API for DNSMasqManager service.
type DNSMasqManagerServer interface {
	RequestAddress(context.Context, *AddressRequest) (*StatusReply, error)
	DeleteAddress(context.Context, *AddressRequest) (*StatusReply, error)
	LookupAddress(context.Context, *AddressRequest) (*AddressReply, error)
}

func RegisterDNSMasqManagerServer(s *grpc.Server, srv DNSMasqManagerServer) {
	s.RegisterService(&_DNSMasqManager_serviceDesc, srv)
}

func _DNSMasqManager_RequestAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSMasqManagerServer).RequestAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dnsmasqmgr.DNSMasqManager/RequestAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSMasqManagerServer).RequestAddress(ctx, req.(*AddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DNSMasqManager_DeleteAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSMasqManagerServer).DeleteAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dnsmasqmgr.DNSMasqManager/DeleteAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSMasqManagerServer).DeleteAddress(ctx, req.(*AddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DNSMasqManager_LookupAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSMasqManagerServer).LookupAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dnsmasqmgr.DNSMasqManager/LookupAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSMasqManagerServer).LookupAddress(ctx, req.(*AddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DNSMasqManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dnsmasqmgr.DNSMasqManager",
	HandlerType: (*DNSMasqManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestAddress",
			Handler:    _DNSMasqManager_RequestAddress_Handler,
		},
		{
			MethodName: "DeleteAddress",
			Handler:    _DNSMasqManager_DeleteAddress_Handler,
		},
		{
			MethodName: "LookupAddress",
			Handler:    _DNSMasqManager_LookupAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dnsmasqmgr.proto",
}
