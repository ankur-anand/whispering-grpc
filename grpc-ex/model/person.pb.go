// Code generated by protoc-gen-go. DO NOT EDIT.
// source: person.proto

package model

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type User struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  uint32   `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetAge() uint32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type Users struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Users) Reset()         { *m = Users{} }
func (m *Users) String() string { return proto.CompactTextString(m) }
func (*Users) ProtoMessage()    {}
func (*Users) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{1}
}

func (m *Users) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Users.Unmarshal(m, b)
}
func (m *Users) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Users.Marshal(b, m, deterministic)
}
func (m *Users) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Users.Merge(m, src)
}
func (m *Users) XXX_Size() int {
	return xxx_messageInfo_Users.Size(m)
}
func (m *Users) XXX_DiscardUnknown() {
	xxx_messageInfo_Users.DiscardUnknown(m)
}

var xxx_messageInfo_Users proto.InternalMessageInfo

func (m *Users) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type Void struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Void) Reset()         { *m = Void{} }
func (m *Void) String() string { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()    {}
func (*Void) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{2}
}

func (m *Void) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Void.Unmarshal(m, b)
}
func (m *Void) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Void.Marshal(b, m, deterministic)
}
func (m *Void) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Void.Merge(m, src)
}
func (m *Void) XXX_Size() int {
	return xxx_messageInfo_Void.Size(m)
}
func (m *Void) XXX_DiscardUnknown() {
	xxx_messageInfo_Void.DiscardUnknown(m)
}

var xxx_messageInfo_Void proto.InternalMessageInfo

func init() {
	proto.RegisterType((*User)(nil), "model.User")
	proto.RegisterType((*Users)(nil), "model.Users")
	proto.RegisterType((*Void)(nil), "model.Void")
}

func init() { proto.RegisterFile("person.proto", fileDescriptor_4c9e10cf24b1156d) }

var fileDescriptor_4c9e10cf24b1156d = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x48, 0x2d, 0x2a,
	0xce, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcd, 0xcd, 0x4f, 0x49, 0xcd, 0x51,
	0x72, 0xe2, 0x62, 0x09, 0x2d, 0x4e, 0x2d, 0x12, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0x04, 0xb8, 0x98, 0x13, 0xd3, 0x53, 0x25,
	0x98, 0x14, 0x18, 0x35, 0x78, 0x83, 0x40, 0x4c, 0x21, 0x11, 0x2e, 0xd6, 0xd4, 0xdc, 0xc4, 0xcc,
	0x1c, 0x09, 0x66, 0xb0, 0x32, 0x08, 0x47, 0x49, 0x8b, 0x8b, 0x15, 0x64, 0x46, 0xb1, 0x90, 0x22,
	0x17, 0x6b, 0x29, 0x88, 0x21, 0xc1, 0xa8, 0xc0, 0xac, 0xc1, 0x6d, 0xc4, 0xad, 0x07, 0xb6, 0x43,
	0x0f, 0x24, 0x19, 0x04, 0x91, 0x51, 0x62, 0xe3, 0x62, 0x09, 0xcb, 0xcf, 0x4c, 0x31, 0x0a, 0xe4,
	0x62, 0x77, 0x4f, 0x2d, 0x01, 0x5b, 0xad, 0xcc, 0xc5, 0xe2, 0x93, 0x59, 0x5c, 0x22, 0x04, 0x53,
	0x0e, 0x92, 0x97, 0xe2, 0x41, 0xd2, 0x5b, 0xac, 0xc4, 0x20, 0xa4, 0xc8, 0xc5, 0xec, 0x98, 0x92,
	0x22, 0x84, 0x6c, 0xa4, 0x14, 0xb2, 0x06, 0x25, 0x86, 0x24, 0x36, 0xb0, 0xc7, 0x8c, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x11, 0x57, 0x05, 0x4b, 0xe8, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GetUserClient is the client API for GetUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GetUserClient interface {
	List(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Users, error)
	Add(ctx context.Context, in *User, opts ...grpc.CallOption) (*Void, error)
}

type getUserClient struct {
	cc *grpc.ClientConn
}

func NewGetUserClient(cc *grpc.ClientConn) GetUserClient {
	return &getUserClient{cc}
}

func (c *getUserClient) List(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/model.GetUser/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getUserClient) Add(ctx context.Context, in *User, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/model.GetUser/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetUserServer is the server API for GetUser service.
type GetUserServer interface {
	List(context.Context, *Void) (*Users, error)
	Add(context.Context, *User) (*Void, error)
}

// UnimplementedGetUserServer can be embedded to have forward compatible implementations.
type UnimplementedGetUserServer struct {
}

func (*UnimplementedGetUserServer) List(ctx context.Context, req *Void) (*Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedGetUserServer) Add(ctx context.Context, req *User) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}

func RegisterGetUserServer(s *grpc.Server, srv GetUserServer) {
	s.RegisterService(&_GetUser_serviceDesc, srv)
}

func _GetUser_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetUserServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.GetUser/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetUserServer).List(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _GetUser_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetUserServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.GetUser/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetUserServer).Add(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

var _GetUser_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.GetUser",
	HandlerType: (*GetUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _GetUser_List_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _GetUser_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "person.proto",
}