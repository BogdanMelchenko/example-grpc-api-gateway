// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/microservice/micro.proto

package example

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

type MicroMessage struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MicroMessage) Reset()         { *m = MicroMessage{} }
func (m *MicroMessage) String() string { return proto.CompactTextString(m) }
func (*MicroMessage) ProtoMessage()    {}
func (*MicroMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_f480b4ae1042f49c, []int{0}
}

func (m *MicroMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MicroMessage.Unmarshal(m, b)
}
func (m *MicroMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MicroMessage.Marshal(b, m, deterministic)
}
func (m *MicroMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MicroMessage.Merge(m, src)
}
func (m *MicroMessage) XXX_Size() int {
	return xxx_messageInfo_MicroMessage.Size(m)
}
func (m *MicroMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_MicroMessage.DiscardUnknown(m)
}

var xxx_messageInfo_MicroMessage proto.InternalMessageInfo

func (m *MicroMessage) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*MicroMessage)(nil), "example.MicroMessage")
}

func init() { proto.RegisterFile("proto/microservice/micro.proto", fileDescriptor_f480b4ae1042f49c) }

var fileDescriptor_f480b4ae1042f49c = []byte{
	// 128 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0xcd, 0x4c, 0x2e, 0xca, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x85, 0x70,
	0xf4, 0xc0, 0x12, 0x42, 0xec, 0xa9, 0x15, 0x89, 0xb9, 0x05, 0x39, 0xa9, 0x4a, 0x2a, 0x5c, 0x3c,
	0xbe, 0x20, 0x71, 0xdf, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0x54, 0x21, 0x11, 0x2e, 0xd6, 0xb2, 0xc4,
	0x9c, 0xd2, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x08, 0xc7, 0xc8, 0x1b, 0xaa, 0x2a,
	0x18, 0x62, 0x94, 0x90, 0x35, 0x17, 0xa7, 0x7b, 0x6a, 0x49, 0x70, 0x49, 0x51, 0x66, 0x5e, 0xba,
	0x90, 0xa8, 0x1e, 0xd4, 0x30, 0x3d, 0x64, 0x93, 0xa4, 0xb0, 0x0b, 0x2b, 0x31, 0x24, 0xb1, 0x81,
	0x9d, 0x60, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x60, 0x9b, 0xd4, 0x22, 0xa4, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MicroServiceClient is the client API for MicroService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MicroServiceClient interface {
	GetString(ctx context.Context, in *MicroMessage, opts ...grpc.CallOption) (*MicroMessage, error)
}

type microServiceClient struct {
	cc *grpc.ClientConn
}

func NewMicroServiceClient(cc *grpc.ClientConn) MicroServiceClient {
	return &microServiceClient{cc}
}

func (c *microServiceClient) GetString(ctx context.Context, in *MicroMessage, opts ...grpc.CallOption) (*MicroMessage, error) {
	out := new(MicroMessage)
	err := c.cc.Invoke(ctx, "/example.MicroService/GetString", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MicroServiceServer is the server API for MicroService service.
type MicroServiceServer interface {
	GetString(context.Context, *MicroMessage) (*MicroMessage, error)
}

func RegisterMicroServiceServer(s *grpc.Server, srv MicroServiceServer) {
	s.RegisterService(&_MicroService_serviceDesc, srv)
}

func _MicroService_GetString_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MicroMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroServiceServer).GetString(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.MicroService/GetString",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroServiceServer).GetString(ctx, req.(*MicroMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _MicroService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.MicroService",
	HandlerType: (*MicroServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetString",
			Handler:    _MicroService_GetString_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/microservice/micro.proto",
}