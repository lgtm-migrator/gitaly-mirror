// Code generated by protoc-gen-go. DO NOT EDIT.
// source: namespace.proto

package gitalypb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type AddNamespaceRequest struct {
	StorageName          string   `protobuf:"bytes,1,opt,name=storage_name,json=storageName,proto3" json:"storage_name,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddNamespaceRequest) Reset()         { *m = AddNamespaceRequest{} }
func (m *AddNamespaceRequest) String() string { return proto.CompactTextString(m) }
func (*AddNamespaceRequest) ProtoMessage()    {}
func (*AddNamespaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecb1e126f615f5dd, []int{0}
}

func (m *AddNamespaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddNamespaceRequest.Unmarshal(m, b)
}
func (m *AddNamespaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddNamespaceRequest.Marshal(b, m, deterministic)
}
func (m *AddNamespaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddNamespaceRequest.Merge(m, src)
}
func (m *AddNamespaceRequest) XXX_Size() int {
	return xxx_messageInfo_AddNamespaceRequest.Size(m)
}
func (m *AddNamespaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddNamespaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddNamespaceRequest proto.InternalMessageInfo

func (m *AddNamespaceRequest) GetStorageName() string {
	if m != nil {
		return m.StorageName
	}
	return ""
}

func (m *AddNamespaceRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type RemoveNamespaceRequest struct {
	StorageName          string   `protobuf:"bytes,1,opt,name=storage_name,json=storageName,proto3" json:"storage_name,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveNamespaceRequest) Reset()         { *m = RemoveNamespaceRequest{} }
func (m *RemoveNamespaceRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveNamespaceRequest) ProtoMessage()    {}
func (*RemoveNamespaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecb1e126f615f5dd, []int{1}
}

func (m *RemoveNamespaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveNamespaceRequest.Unmarshal(m, b)
}
func (m *RemoveNamespaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveNamespaceRequest.Marshal(b, m, deterministic)
}
func (m *RemoveNamespaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveNamespaceRequest.Merge(m, src)
}
func (m *RemoveNamespaceRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveNamespaceRequest.Size(m)
}
func (m *RemoveNamespaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveNamespaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveNamespaceRequest proto.InternalMessageInfo

func (m *RemoveNamespaceRequest) GetStorageName() string {
	if m != nil {
		return m.StorageName
	}
	return ""
}

func (m *RemoveNamespaceRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type RenameNamespaceRequest struct {
	StorageName          string   `protobuf:"bytes,1,opt,name=storage_name,json=storageName,proto3" json:"storage_name,omitempty"`
	From                 string   `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To                   string   `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RenameNamespaceRequest) Reset()         { *m = RenameNamespaceRequest{} }
func (m *RenameNamespaceRequest) String() string { return proto.CompactTextString(m) }
func (*RenameNamespaceRequest) ProtoMessage()    {}
func (*RenameNamespaceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecb1e126f615f5dd, []int{2}
}

func (m *RenameNamespaceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RenameNamespaceRequest.Unmarshal(m, b)
}
func (m *RenameNamespaceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RenameNamespaceRequest.Marshal(b, m, deterministic)
}
func (m *RenameNamespaceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RenameNamespaceRequest.Merge(m, src)
}
func (m *RenameNamespaceRequest) XXX_Size() int {
	return xxx_messageInfo_RenameNamespaceRequest.Size(m)
}
func (m *RenameNamespaceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RenameNamespaceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RenameNamespaceRequest proto.InternalMessageInfo

func (m *RenameNamespaceRequest) GetStorageName() string {
	if m != nil {
		return m.StorageName
	}
	return ""
}

func (m *RenameNamespaceRequest) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *RenameNamespaceRequest) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

type NamespaceExistsRequest struct {
	StorageName          string   `protobuf:"bytes,1,opt,name=storage_name,json=storageName,proto3" json:"storage_name,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NamespaceExistsRequest) Reset()         { *m = NamespaceExistsRequest{} }
func (m *NamespaceExistsRequest) String() string { return proto.CompactTextString(m) }
func (*NamespaceExistsRequest) ProtoMessage()    {}
func (*NamespaceExistsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecb1e126f615f5dd, []int{3}
}

func (m *NamespaceExistsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NamespaceExistsRequest.Unmarshal(m, b)
}
func (m *NamespaceExistsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NamespaceExistsRequest.Marshal(b, m, deterministic)
}
func (m *NamespaceExistsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NamespaceExistsRequest.Merge(m, src)
}
func (m *NamespaceExistsRequest) XXX_Size() int {
	return xxx_messageInfo_NamespaceExistsRequest.Size(m)
}
func (m *NamespaceExistsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NamespaceExistsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NamespaceExistsRequest proto.InternalMessageInfo

func (m *NamespaceExistsRequest) GetStorageName() string {
	if m != nil {
		return m.StorageName
	}
	return ""
}

func (m *NamespaceExistsRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type NamespaceExistsResponse struct {
	Exists               bool     `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NamespaceExistsResponse) Reset()         { *m = NamespaceExistsResponse{} }
func (m *NamespaceExistsResponse) String() string { return proto.CompactTextString(m) }
func (*NamespaceExistsResponse) ProtoMessage()    {}
func (*NamespaceExistsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecb1e126f615f5dd, []int{4}
}

func (m *NamespaceExistsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NamespaceExistsResponse.Unmarshal(m, b)
}
func (m *NamespaceExistsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NamespaceExistsResponse.Marshal(b, m, deterministic)
}
func (m *NamespaceExistsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NamespaceExistsResponse.Merge(m, src)
}
func (m *NamespaceExistsResponse) XXX_Size() int {
	return xxx_messageInfo_NamespaceExistsResponse.Size(m)
}
func (m *NamespaceExistsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NamespaceExistsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NamespaceExistsResponse proto.InternalMessageInfo

func (m *NamespaceExistsResponse) GetExists() bool {
	if m != nil {
		return m.Exists
	}
	return false
}

type AddNamespaceResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddNamespaceResponse) Reset()         { *m = AddNamespaceResponse{} }
func (m *AddNamespaceResponse) String() string { return proto.CompactTextString(m) }
func (*AddNamespaceResponse) ProtoMessage()    {}
func (*AddNamespaceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecb1e126f615f5dd, []int{5}
}

func (m *AddNamespaceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddNamespaceResponse.Unmarshal(m, b)
}
func (m *AddNamespaceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddNamespaceResponse.Marshal(b, m, deterministic)
}
func (m *AddNamespaceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddNamespaceResponse.Merge(m, src)
}
func (m *AddNamespaceResponse) XXX_Size() int {
	return xxx_messageInfo_AddNamespaceResponse.Size(m)
}
func (m *AddNamespaceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddNamespaceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddNamespaceResponse proto.InternalMessageInfo

type RemoveNamespaceResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveNamespaceResponse) Reset()         { *m = RemoveNamespaceResponse{} }
func (m *RemoveNamespaceResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveNamespaceResponse) ProtoMessage()    {}
func (*RemoveNamespaceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecb1e126f615f5dd, []int{6}
}

func (m *RemoveNamespaceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveNamespaceResponse.Unmarshal(m, b)
}
func (m *RemoveNamespaceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveNamespaceResponse.Marshal(b, m, deterministic)
}
func (m *RemoveNamespaceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveNamespaceResponse.Merge(m, src)
}
func (m *RemoveNamespaceResponse) XXX_Size() int {
	return xxx_messageInfo_RemoveNamespaceResponse.Size(m)
}
func (m *RemoveNamespaceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveNamespaceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveNamespaceResponse proto.InternalMessageInfo

type RenameNamespaceResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RenameNamespaceResponse) Reset()         { *m = RenameNamespaceResponse{} }
func (m *RenameNamespaceResponse) String() string { return proto.CompactTextString(m) }
func (*RenameNamespaceResponse) ProtoMessage()    {}
func (*RenameNamespaceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecb1e126f615f5dd, []int{7}
}

func (m *RenameNamespaceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RenameNamespaceResponse.Unmarshal(m, b)
}
func (m *RenameNamespaceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RenameNamespaceResponse.Marshal(b, m, deterministic)
}
func (m *RenameNamespaceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RenameNamespaceResponse.Merge(m, src)
}
func (m *RenameNamespaceResponse) XXX_Size() int {
	return xxx_messageInfo_RenameNamespaceResponse.Size(m)
}
func (m *RenameNamespaceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RenameNamespaceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RenameNamespaceResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AddNamespaceRequest)(nil), "gitaly.AddNamespaceRequest")
	proto.RegisterType((*RemoveNamespaceRequest)(nil), "gitaly.RemoveNamespaceRequest")
	proto.RegisterType((*RenameNamespaceRequest)(nil), "gitaly.RenameNamespaceRequest")
	proto.RegisterType((*NamespaceExistsRequest)(nil), "gitaly.NamespaceExistsRequest")
	proto.RegisterType((*NamespaceExistsResponse)(nil), "gitaly.NamespaceExistsResponse")
	proto.RegisterType((*AddNamespaceResponse)(nil), "gitaly.AddNamespaceResponse")
	proto.RegisterType((*RemoveNamespaceResponse)(nil), "gitaly.RemoveNamespaceResponse")
	proto.RegisterType((*RenameNamespaceResponse)(nil), "gitaly.RenameNamespaceResponse")
}

func init() { proto.RegisterFile("namespace.proto", fileDescriptor_ecb1e126f615f5dd) }

var fileDescriptor_ecb1e126f615f5dd = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcd, 0x4e, 0x02, 0x31,
	0x10, 0xc7, 0x43, 0x21, 0x04, 0x47, 0x22, 0xa4, 0x1a, 0x40, 0x34, 0x7e, 0xec, 0x89, 0x8b, 0xbb,
	0x7e, 0x3c, 0x81, 0x26, 0xde, 0x8c, 0x26, 0xcb, 0xcd, 0x98, 0x90, 0x02, 0xe3, 0x4a, 0xc2, 0xd2,
	0xb5, 0xad, 0x44, 0x8f, 0x3e, 0x85, 0xef, 0xea, 0xc9, 0xb4, 0x5d, 0x60, 0x77, 0x29, 0x17, 0xf5,
	0x36, 0xfd, 0xff, 0x67, 0x7f, 0x9d, 0x9d, 0x99, 0x42, 0x63, 0xc6, 0x62, 0x94, 0x09, 0x1b, 0xa1,
	0x9f, 0x08, 0xae, 0x38, 0xad, 0x46, 0x13, 0xc5, 0xa6, 0x1f, 0xdd, 0xba, 0x7c, 0x61, 0x02, 0xc7,
	0x56, 0xf5, 0xee, 0x60, 0xf7, 0x7a, 0x3c, 0xbe, 0x5f, 0xe4, 0x86, 0xf8, 0xfa, 0x86, 0x52, 0xd1,
	0x53, 0xa8, 0x4b, 0xc5, 0x05, 0x8b, 0x70, 0xa0, 0x39, 0x9d, 0xd2, 0x49, 0xa9, 0xb7, 0x15, 0x6e,
	0xa7, 0x9a, 0x4e, 0xa7, 0x14, 0x2a, 0xc6, 0x22, 0xc6, 0x32, 0xb1, 0xf7, 0x00, 0xad, 0x10, 0x63,
	0x3e, 0xc7, 0xff, 0x02, 0x0e, 0x34, 0x50, 0x47, 0xbf, 0x04, 0x3e, 0x0b, 0x1e, 0x2f, 0x80, 0x3a,
	0xa6, 0x3b, 0x40, 0x14, 0xef, 0x94, 0x8d, 0x42, 0x14, 0xd7, 0x15, 0x2f, 0xd1, 0xb7, 0xef, 0x13,
	0xa9, 0xe4, 0x1f, 0x2b, 0xbe, 0x80, 0xf6, 0x1a, 0x50, 0x26, 0x7c, 0x26, 0x91, 0xb6, 0xa0, 0x8a,
	0x46, 0x31, 0xac, 0x5a, 0x98, 0x9e, 0xbc, 0x16, 0xec, 0xe5, 0x67, 0x60, 0xf3, 0xbd, 0x7d, 0x68,
	0xaf, 0x75, 0x33, 0x6b, 0x15, 0xfa, 0x62, 0xad, 0xcb, 0xcf, 0x32, 0x34, 0x97, 0x6a, 0x1f, 0xc5,
	0x7c, 0x32, 0x42, 0xda, 0x87, 0x7a, 0xf6, 0x0a, 0x7a, 0xe0, 0xdb, 0x6d, 0xf0, 0x1d, 0xc3, 0xef,
	0x1e, 0xba, 0xcd, 0xf4, 0xea, 0xda, 0xf7, 0x57, 0xaf, 0x52, 0x2b, 0x35, 0x09, 0x7d, 0x82, 0x46,
	0xa1, 0x3e, 0x7a, 0xb4, 0xf8, 0xd4, 0xbd, 0x06, 0xdd, 0xe3, 0x8d, 0xbe, 0x9b, 0x9e, 0xfb, 0xc5,
	0x2c, 0xdd, 0xb5, 0x13, 0x59, 0xba, 0xb3, 0x37, 0x79, 0x7a, 0x61, 0x4c, 0x2b, 0xba, 0x7b, 0x21,
	0x56, 0xf4, 0x0d, 0xf3, 0x4d, 0xe9, 0xa4, 0x49, 0x6e, 0xce, 0x1f, 0x75, 0xee, 0x94, 0x0d, 0xfd,
	0x11, 0x8f, 0x03, 0x1b, 0x9e, 0x71, 0x11, 0x05, 0x96, 0x10, 0x98, 0xb7, 0x17, 0x44, 0x3c, 0x3d,
	0x27, 0xc3, 0x61, 0xd5, 0x48, 0x57, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x03, 0xa2, 0x6d,
	0xb7, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NamespaceServiceClient is the client API for NamespaceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NamespaceServiceClient interface {
	AddNamespace(ctx context.Context, in *AddNamespaceRequest, opts ...grpc.CallOption) (*AddNamespaceResponse, error)
	RemoveNamespace(ctx context.Context, in *RemoveNamespaceRequest, opts ...grpc.CallOption) (*RemoveNamespaceResponse, error)
	RenameNamespace(ctx context.Context, in *RenameNamespaceRequest, opts ...grpc.CallOption) (*RenameNamespaceResponse, error)
	NamespaceExists(ctx context.Context, in *NamespaceExistsRequest, opts ...grpc.CallOption) (*NamespaceExistsResponse, error)
}

type namespaceServiceClient struct {
	cc *grpc.ClientConn
}

func NewNamespaceServiceClient(cc *grpc.ClientConn) NamespaceServiceClient {
	return &namespaceServiceClient{cc}
}

func (c *namespaceServiceClient) AddNamespace(ctx context.Context, in *AddNamespaceRequest, opts ...grpc.CallOption) (*AddNamespaceResponse, error) {
	out := new(AddNamespaceResponse)
	err := c.cc.Invoke(ctx, "/gitaly.NamespaceService/AddNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespaceServiceClient) RemoveNamespace(ctx context.Context, in *RemoveNamespaceRequest, opts ...grpc.CallOption) (*RemoveNamespaceResponse, error) {
	out := new(RemoveNamespaceResponse)
	err := c.cc.Invoke(ctx, "/gitaly.NamespaceService/RemoveNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespaceServiceClient) RenameNamespace(ctx context.Context, in *RenameNamespaceRequest, opts ...grpc.CallOption) (*RenameNamespaceResponse, error) {
	out := new(RenameNamespaceResponse)
	err := c.cc.Invoke(ctx, "/gitaly.NamespaceService/RenameNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespaceServiceClient) NamespaceExists(ctx context.Context, in *NamespaceExistsRequest, opts ...grpc.CallOption) (*NamespaceExistsResponse, error) {
	out := new(NamespaceExistsResponse)
	err := c.cc.Invoke(ctx, "/gitaly.NamespaceService/NamespaceExists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NamespaceServiceServer is the server API for NamespaceService service.
type NamespaceServiceServer interface {
	AddNamespace(context.Context, *AddNamespaceRequest) (*AddNamespaceResponse, error)
	RemoveNamespace(context.Context, *RemoveNamespaceRequest) (*RemoveNamespaceResponse, error)
	RenameNamespace(context.Context, *RenameNamespaceRequest) (*RenameNamespaceResponse, error)
	NamespaceExists(context.Context, *NamespaceExistsRequest) (*NamespaceExistsResponse, error)
}

// UnimplementedNamespaceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNamespaceServiceServer struct {
}

func (*UnimplementedNamespaceServiceServer) AddNamespace(ctx context.Context, req *AddNamespaceRequest) (*AddNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNamespace not implemented")
}
func (*UnimplementedNamespaceServiceServer) RemoveNamespace(ctx context.Context, req *RemoveNamespaceRequest) (*RemoveNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveNamespace not implemented")
}
func (*UnimplementedNamespaceServiceServer) RenameNamespace(ctx context.Context, req *RenameNamespaceRequest) (*RenameNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenameNamespace not implemented")
}
func (*UnimplementedNamespaceServiceServer) NamespaceExists(ctx context.Context, req *NamespaceExistsRequest) (*NamespaceExistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NamespaceExists not implemented")
}

func RegisterNamespaceServiceServer(s *grpc.Server, srv NamespaceServiceServer) {
	s.RegisterService(&_NamespaceService_serviceDesc, srv)
}

func _NamespaceService_AddNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceServiceServer).AddNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.NamespaceService/AddNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceServiceServer).AddNamespace(ctx, req.(*AddNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NamespaceService_RemoveNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceServiceServer).RemoveNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.NamespaceService/RemoveNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceServiceServer).RemoveNamespace(ctx, req.(*RemoveNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NamespaceService_RenameNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenameNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceServiceServer).RenameNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.NamespaceService/RenameNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceServiceServer).RenameNamespace(ctx, req.(*RenameNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NamespaceService_NamespaceExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NamespaceExistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceServiceServer).NamespaceExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.NamespaceService/NamespaceExists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceServiceServer).NamespaceExists(ctx, req.(*NamespaceExistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NamespaceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gitaly.NamespaceService",
	HandlerType: (*NamespaceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddNamespace",
			Handler:    _NamespaceService_AddNamespace_Handler,
		},
		{
			MethodName: "RemoveNamespace",
			Handler:    _NamespaceService_RemoveNamespace_Handler,
		},
		{
			MethodName: "RenameNamespace",
			Handler:    _NamespaceService_RenameNamespace_Handler,
		},
		{
			MethodName: "NamespaceExists",
			Handler:    _NamespaceService_NamespaceExists_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "namespace.proto",
}