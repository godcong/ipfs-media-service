// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node.proto

package proto

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type NodeRemoteType int32

const (
	NodeRemoteType_Basic NodeRemoteType = 0
	NodeRemoteType_Retry NodeRemoteType = 1
	NodeRemoteType_Force NodeRemoteType = 2
)

var NodeRemoteType_name = map[int32]string{
	0: "Basic",
	1: "Retry",
	2: "Force",
}

var NodeRemoteType_value = map[string]int32{
	"Basic": 0,
	"Retry": 1,
	"Force": 2,
}

func (x NodeRemoteType) String() string {
	return proto.EnumName(NodeRemoteType_name, int32(x))
}

func (NodeRemoteType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{0}
}

type NodeBackType int32

const (
	NodeBackType_HTTP NodeBackType = 0
	NodeBackType_GRPC NodeBackType = 1
)

var NodeBackType_name = map[int32]string{
	0: "HTTP",
	1: "GRPC",
}

var NodeBackType_value = map[string]int32{
	"HTTP": 0,
	"GRPC": 1,
}

func (x NodeBackType) String() string {
	return proto.EnumName(NodeBackType_name, int32(x))
}

func (NodeBackType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{1}
}

type StatusRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusRequest) Reset()         { *m = StatusRequest{} }
func (m *StatusRequest) String() string { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()    {}
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{0}
}

func (m *StatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusRequest.Unmarshal(m, b)
}
func (m *StatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusRequest.Marshal(b, m, deterministic)
}
func (m *StatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusRequest.Merge(m, src)
}
func (m *StatusRequest) XXX_Size() int {
	return xxx_messageInfo_StatusRequest.Size(m)
}
func (m *StatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatusRequest proto.InternalMessageInfo

func (m *StatusRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type RemoteDownloadRequest struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	ObjectKey            string   `protobuf:"bytes,2,opt,name=objectKey,proto3" json:"objectKey,omitempty"`
	Encrypt              bool     `protobuf:"varint,3,opt,name=encrypt,proto3" json:"encrypt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoteDownloadRequest) Reset()         { *m = RemoteDownloadRequest{} }
func (m *RemoteDownloadRequest) String() string { return proto.CompactTextString(m) }
func (*RemoteDownloadRequest) ProtoMessage()    {}
func (*RemoteDownloadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{1}
}

func (m *RemoteDownloadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoteDownloadRequest.Unmarshal(m, b)
}
func (m *RemoteDownloadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoteDownloadRequest.Marshal(b, m, deterministic)
}
func (m *RemoteDownloadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteDownloadRequest.Merge(m, src)
}
func (m *RemoteDownloadRequest) XXX_Size() int {
	return xxx_messageInfo_RemoteDownloadRequest.Size(m)
}
func (m *RemoteDownloadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteDownloadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteDownloadRequest proto.InternalMessageInfo

func (m *RemoteDownloadRequest) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *RemoteDownloadRequest) GetObjectKey() string {
	if m != nil {
		return m.ObjectKey
	}
	return ""
}

func (m *RemoteDownloadRequest) GetEncrypt() bool {
	if m != nil {
		return m.Encrypt
	}
	return false
}

type NodeReply struct {
	Code                 int32            `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string           `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Detail               *NodeReplyDetail `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *NodeReply) Reset()         { *m = NodeReply{} }
func (m *NodeReply) String() string { return proto.CompactTextString(m) }
func (*NodeReply) ProtoMessage()    {}
func (*NodeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{2}
}

func (m *NodeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeReply.Unmarshal(m, b)
}
func (m *NodeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeReply.Marshal(b, m, deterministic)
}
func (m *NodeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeReply.Merge(m, src)
}
func (m *NodeReply) XXX_Size() int {
	return xxx_messageInfo_NodeReply.Size(m)
}
func (m *NodeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeReply.DiscardUnknown(m)
}

var xxx_messageInfo_NodeReply proto.InternalMessageInfo

func (m *NodeReply) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *NodeReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *NodeReply) GetDetail() *NodeReplyDetail {
	if m != nil {
		return m.Detail
	}
	return nil
}

type NodeReplyDetail struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Json                 string   `protobuf:"bytes,2,opt,name=json,proto3" json:"json,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeReplyDetail) Reset()         { *m = NodeReplyDetail{} }
func (m *NodeReplyDetail) String() string { return proto.CompactTextString(m) }
func (*NodeReplyDetail) ProtoMessage()    {}
func (*NodeReplyDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{3}
}

func (m *NodeReplyDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeReplyDetail.Unmarshal(m, b)
}
func (m *NodeReplyDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeReplyDetail.Marshal(b, m, deterministic)
}
func (m *NodeReplyDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeReplyDetail.Merge(m, src)
}
func (m *NodeReplyDetail) XXX_Size() int {
	return xxx_messageInfo_NodeReplyDetail.Size(m)
}
func (m *NodeReplyDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeReplyDetail.DiscardUnknown(m)
}

var xxx_messageInfo_NodeReplyDetail proto.InternalMessageInfo

func (m *NodeReplyDetail) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *NodeReplyDetail) GetJson() string {
	if m != nil {
		return m.Json
	}
	return ""
}

func init() {
	proto.RegisterEnum("proto.NodeRemoteType", NodeRemoteType_name, NodeRemoteType_value)
	proto.RegisterEnum("proto.NodeBackType", NodeBackType_name, NodeBackType_value)
	proto.RegisterType((*StatusRequest)(nil), "proto.StatusRequest")
	proto.RegisterType((*RemoteDownloadRequest)(nil), "proto.RemoteDownloadRequest")
	proto.RegisterType((*NodeReply)(nil), "proto.NodeReply")
	proto.RegisterType((*NodeReplyDetail)(nil), "proto.NodeReplyDetail")
}

func init() { proto.RegisterFile("node.proto", fileDescriptor_0c843d59d2d938e7) }

var fileDescriptor_0c843d59d2d938e7 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0xdd, 0x4e, 0xc2, 0x30,
	0x14, 0x66, 0x13, 0x90, 0x1d, 0x14, 0x97, 0x46, 0xcd, 0x62, 0x48, 0x24, 0xbb, 0x22, 0x5c, 0x2c,
	0x11, 0xe3, 0x0b, 0xcc, 0x45, 0x25, 0x26, 0x66, 0x29, 0xdc, 0x9b, 0xd1, 0x9e, 0xcc, 0x21, 0xac,
	0x73, 0x2b, 0x9a, 0x3d, 0x80, 0xef, 0x6d, 0xda, 0x0d, 0x05, 0xe4, 0x6a, 0x5f, 0xcf, 0xf7, 0xb3,
	0x7e, 0x3d, 0x00, 0xa9, 0xe0, 0xe8, 0x65, 0xb9, 0x90, 0x82, 0xb4, 0xf4, 0xc7, 0xbd, 0x86, 0xd3,
	0xa9, 0x8c, 0xe4, 0xba, 0xa0, 0xf8, 0xb1, 0xc6, 0x42, 0x92, 0x1e, 0x98, 0x09, 0x77, 0x8c, 0x81,
	0x31, 0xb4, 0xa8, 0x99, 0x70, 0xf7, 0x15, 0x2e, 0x28, 0xae, 0x84, 0xc4, 0x40, 0x7c, 0xa5, 0x4b,
	0x11, 0xf1, 0x2d, 0xe1, 0x24, 0xd8, 0x08, 0x27, 0x01, 0xe9, 0x83, 0x25, 0xe6, 0x0b, 0x64, 0xf2,
	0x19, 0x4b, 0xc7, 0xd4, 0xe3, 0xbf, 0x01, 0x71, 0xe0, 0x18, 0x53, 0x96, 0x97, 0x99, 0x74, 0x8e,
	0x06, 0xc6, 0xb0, 0x43, 0x37, 0x47, 0x37, 0x01, 0xeb, 0x45, 0x70, 0xa4, 0x98, 0x2d, 0x4b, 0x42,
	0xa0, 0xc9, 0x04, 0x47, 0x1d, 0xdb, 0xa2, 0x1a, 0x2b, 0xeb, 0x0a, 0x8b, 0x22, 0x8a, 0xb1, 0x8e,
	0xdd, 0x1c, 0x89, 0x07, 0x6d, 0x8e, 0x32, 0x4a, 0x96, 0x3a, 0xb3, 0x3b, 0xbe, 0xac, 0xba, 0x79,
	0xbf, 0x79, 0x81, 0x66, 0x69, 0xad, 0x72, 0xef, 0xe0, 0x6c, 0x8f, 0xfa, 0xd7, 0x82, 0x40, 0x73,
	0x51, 0x88, 0xb4, 0xfe, 0x93, 0xc6, 0xa3, 0x1b, 0xe8, 0x55, 0x36, 0xf5, 0x0c, 0xb3, 0x32, 0x43,
	0x62, 0x41, 0xcb, 0x8f, 0x8a, 0x84, 0xd9, 0x0d, 0x05, 0x29, 0xca, 0xbc, 0xb4, 0x0d, 0x05, 0x1f,
	0x44, 0xce, 0xd0, 0x36, 0x47, 0x2e, 0x9c, 0x28, 0x8b, 0x1f, 0xb1, 0x77, 0x6d, 0xe8, 0x40, 0xf3,
	0x69, 0x36, 0x0b, 0xed, 0x86, 0x42, 0x8f, 0x34, 0xbc, 0xb7, 0x8d, 0xf1, 0xb7, 0x01, 0x5d, 0x25,
	0x9a, 0x62, 0xfe, 0x99, 0x30, 0x24, 0x3e, 0xf4, 0x76, 0x5f, 0x9a, 0xf4, 0xeb, 0x3e, 0x07, 0x17,
	0x70, 0x65, 0xef, 0xb7, 0x75, 0x1b, 0x64, 0x0c, 0xed, 0x6a, 0x9d, 0xe4, 0xbc, 0x66, 0x77, 0xb6,
	0x7b, 0xc8, 0xe3, 0x7b, 0xe0, 0x30, 0xb1, 0xf2, 0xe2, 0x44, 0xbe, 0xad, 0xe7, 0x5e, 0x2c, 0x38,
	0x13, 0x69, 0x5c, 0xe9, 0x7c, 0x7b, 0xeb, 0x82, 0xa1, 0x9a, 0x84, 0xc6, 0xbc, 0xad, 0xa9, 0xdb,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x65, 0xe2, 0xb5, 0x70, 0x4e, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NodeServiceClient is the client API for NodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeServiceClient interface {
	RemoteDownload(ctx context.Context, in *RemoteDownloadRequest, opts ...grpc.CallOption) (*NodeReply, error)
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*NodeReply, error)
}

type nodeServiceClient struct {
	cc *grpc.ClientConn
}

func NewNodeServiceClient(cc *grpc.ClientConn) NodeServiceClient {
	return &nodeServiceClient{cc}
}

func (c *nodeServiceClient) RemoteDownload(ctx context.Context, in *RemoteDownloadRequest, opts ...grpc.CallOption) (*NodeReply, error) {
	out := new(NodeReply)
	err := c.cc.Invoke(ctx, "/proto.NodeService/RemoteDownload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*NodeReply, error) {
	out := new(NodeReply)
	err := c.cc.Invoke(ctx, "/proto.NodeService/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServiceServer is the server API for NodeService service.
type NodeServiceServer interface {
	RemoteDownload(context.Context, *RemoteDownloadRequest) (*NodeReply, error)
	Status(context.Context, *StatusRequest) (*NodeReply, error)
}

func RegisterNodeServiceServer(s *grpc.Server, srv NodeServiceServer) {
	s.RegisterService(&_NodeService_serviceDesc, srv)
}

func _NodeService_RemoteDownload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoteDownloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).RemoteDownload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NodeService/RemoteDownload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).RemoteDownload(ctx, req.(*RemoteDownloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NodeService/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.NodeService",
	HandlerType: (*NodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RemoteDownload",
			Handler:    _NodeService_RemoteDownload_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _NodeService_Status_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node.proto",
}
