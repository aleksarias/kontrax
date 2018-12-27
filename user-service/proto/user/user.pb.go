// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

package go_micro_srv_user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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
	Id                   string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName            string           `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	MiddleName           string           `protobuf:"bytes,3,opt,name=middle_name,json=middleName,proto3" json:"middle_name,omitempty"`
	LastName             string           `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email                string           `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Password             string           `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	Phones               []*Phone         `protobuf:"bytes,7,rep,name=phones,proto3" json:"phones,omitempty"`
	Addresses            []*PostalAddress `protobuf:"bytes,8,rep,name=addresses,proto3" json:"addresses,omitempty"`
	Type                 string           `protobuf:"bytes,9,opt,name=type,proto3" json:"type,omitempty"`
	Role                 string           `protobuf:"bytes,10,opt,name=role,proto3" json:"role,omitempty"`
	Organization         *Organization    `protobuf:"bytes,11,opt,name=organization,proto3" json:"organization,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{0}
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

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetMiddleName() string {
	if m != nil {
		return m.MiddleName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetPhones() []*Phone {
	if m != nil {
		return m.Phones
	}
	return nil
}

func (m *User) GetAddresses() []*PostalAddress {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *User) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *User) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *User) GetOrganization() *Organization {
	if m != nil {
		return m.Organization
	}
	return nil
}

type Organization struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PhoneNumber          string   `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Organization) Reset()         { *m = Organization{} }
func (m *Organization) String() string { return proto.CompactTextString(m) }
func (*Organization) ProtoMessage()    {}
func (*Organization) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{1}
}

func (m *Organization) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Organization.Unmarshal(m, b)
}
func (m *Organization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Organization.Marshal(b, m, deterministic)
}
func (m *Organization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Organization.Merge(m, src)
}
func (m *Organization) XXX_Size() int {
	return xxx_messageInfo_Organization.Size(m)
}
func (m *Organization) XXX_DiscardUnknown() {
	xxx_messageInfo_Organization.DiscardUnknown(m)
}

var xxx_messageInfo_Organization proto.InternalMessageInfo

func (m *Organization) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Organization) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Organization) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

type PostalAddress struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	City                 string   `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	State                string   `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	Zip                  string   `protobuf:"bytes,5,opt,name=zip,proto3" json:"zip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostalAddress) Reset()         { *m = PostalAddress{} }
func (m *PostalAddress) String() string { return proto.CompactTextString(m) }
func (*PostalAddress) ProtoMessage()    {}
func (*PostalAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{2}
}

func (m *PostalAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostalAddress.Unmarshal(m, b)
}
func (m *PostalAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostalAddress.Marshal(b, m, deterministic)
}
func (m *PostalAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostalAddress.Merge(m, src)
}
func (m *PostalAddress) XXX_Size() int {
	return xxx_messageInfo_PostalAddress.Size(m)
}
func (m *PostalAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_PostalAddress.DiscardUnknown(m)
}

var xxx_messageInfo_PostalAddress proto.InternalMessageInfo

func (m *PostalAddress) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PostalAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *PostalAddress) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *PostalAddress) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *PostalAddress) GetZip() string {
	if m != nil {
		return m.Zip
	}
	return ""
}

type Phone struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PhoneNumber          string   `protobuf:"bytes,2,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Ext                  string   `protobuf:"bytes,3,opt,name=ext,proto3" json:"ext,omitempty"`
	PhoneType            string   `protobuf:"bytes,4,opt,name=phone_type,json=phoneType,proto3" json:"phone_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Phone) Reset()         { *m = Phone{} }
func (m *Phone) String() string { return proto.CompactTextString(m) }
func (*Phone) ProtoMessage()    {}
func (*Phone) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{3}
}

func (m *Phone) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Phone.Unmarshal(m, b)
}
func (m *Phone) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Phone.Marshal(b, m, deterministic)
}
func (m *Phone) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Phone.Merge(m, src)
}
func (m *Phone) XXX_Size() int {
	return xxx_messageInfo_Phone.Size(m)
}
func (m *Phone) XXX_DiscardUnknown() {
	xxx_messageInfo_Phone.DiscardUnknown(m)
}

var xxx_messageInfo_Phone proto.InternalMessageInfo

func (m *Phone) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Phone) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *Phone) GetExt() string {
	if m != nil {
		return m.Ext
	}
	return ""
}

func (m *Phone) GetPhoneType() string {
	if m != nil {
		return m.PhoneType
	}
	return ""
}

type Response struct {
	Created              bool     `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	User                 *User    `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Users                []*User  `protobuf:"bytes,3,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{4}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Response) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type GetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{5}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*User)(nil), "go.micro.srv.user.User")
	proto.RegisterType((*Organization)(nil), "go.micro.srv.user.Organization")
	proto.RegisterType((*PostalAddress)(nil), "go.micro.srv.user.PostalAddress")
	proto.RegisterType((*Phone)(nil), "go.micro.srv.user.Phone")
	proto.RegisterType((*Response)(nil), "go.micro.srv.user.Response")
	proto.RegisterType((*GetRequest)(nil), "go.micro.srv.user.GetRequest")
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor_9b283a848145d6b7) }

var fileDescriptor_9b283a848145d6b7 = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x41, 0x6f, 0xd3, 0x4c,
	0x10, 0xfd, 0x12, 0x3b, 0xa9, 0x3d, 0x4e, 0x3f, 0x95, 0x15, 0x88, 0x55, 0x2b, 0xd4, 0xe0, 0x53,
	0x24, 0x84, 0x41, 0xe1, 0x0c, 0xa2, 0xea, 0x21, 0x12, 0x87, 0x52, 0x19, 0x7a, 0xae, 0xb6, 0xf1,
	0x50, 0x56, 0xd8, 0x5e, 0xb3, 0xbb, 0x29, 0xa4, 0x27, 0x0e, 0xfc, 0x13, 0xfe, 0x28, 0xda, 0x59,
	0x1b, 0x52, 0x9c, 0x5e, 0x72, 0x49, 0x66, 0xde, 0x7b, 0x33, 0xb3, 0xa3, 0xb7, 0x6b, 0x78, 0xd4,
	0x68, 0x65, 0xd5, 0x8b, 0x95, 0x41, 0x4d, 0x3f, 0x19, 0xe5, 0xec, 0xc1, 0xb5, 0xca, 0x2a, 0xb9,
	0xd4, 0x2a, 0x33, 0xfa, 0x26, 0x73, 0x44, 0xfa, 0x33, 0x80, 0xf0, 0xc2, 0xa0, 0x66, 0xff, 0xc3,
	0x50, 0x16, 0x7c, 0x30, 0x1d, 0xcc, 0xe2, 0x7c, 0x28, 0x0b, 0xf6, 0x04, 0xe0, 0x93, 0xd4, 0xc6,
	0x5e, 0xd6, 0xa2, 0x42, 0x3e, 0x24, 0x3c, 0x26, 0xe4, 0x4c, 0x54, 0xc8, 0x8e, 0x21, 0xa9, 0x64,
	0x51, 0x94, 0xe8, 0xf9, 0x80, 0x78, 0xf0, 0x10, 0x09, 0x8e, 0x20, 0x2e, 0x45, 0x57, 0x1e, 0x12,
	0x1d, 0x39, 0x80, 0xc8, 0x87, 0x30, 0xc2, 0x4a, 0xc8, 0x92, 0x8f, 0x88, 0xf0, 0x09, 0x3b, 0x84,
	0xa8, 0x11, 0xc6, 0x7c, 0x53, 0xba, 0xe0, 0x63, 0x5f, 0xd1, 0xe5, 0xec, 0x25, 0x8c, 0x9b, 0xcf,
	0xaa, 0x46, 0xc3, 0xf7, 0xa6, 0xc1, 0x2c, 0x99, 0xf3, 0xac, 0xb7, 0x4b, 0x76, 0xee, 0x04, 0x79,
	0xab, 0x63, 0x6f, 0x20, 0x16, 0x45, 0xa1, 0xd1, 0x18, 0x34, 0x3c, 0xa2, 0xa2, 0xe9, 0xb6, 0x22,
	0x65, 0xac, 0x28, 0x4f, 0xbc, 0x32, 0xff, 0x5b, 0xc2, 0x18, 0x84, 0x76, 0xdd, 0x20, 0x8f, 0xe9,
	0x24, 0x14, 0x3b, 0x4c, 0xab, 0x12, 0x39, 0x78, 0xcc, 0xc5, 0xec, 0x14, 0x26, 0x4a, 0x5f, 0x8b,
	0x5a, 0xde, 0x0a, 0x2b, 0x55, 0xcd, 0x93, 0xe9, 0x60, 0x96, 0xcc, 0x8f, 0xb7, 0x8c, 0x7a, 0xbf,
	0x21, 0xcb, 0xef, 0x14, 0xa5, 0x17, 0x30, 0xd9, 0x64, 0x7b, 0x6e, 0x30, 0x08, 0x37, 0x7c, 0xa0,
	0x98, 0x3d, 0x85, 0x09, 0xad, 0x7a, 0x59, 0xaf, 0xaa, 0x2b, 0xd4, 0xad, 0x07, 0x09, 0x61, 0x67,
	0x04, 0xa5, 0x2b, 0xd8, 0xbf, 0xb3, 0x5f, 0xaf, 0x2f, 0x87, 0xbd, 0x76, 0xe3, 0xb6, 0x75, 0x97,
	0xba, 0x89, 0x4b, 0x69, 0xd7, 0x6d, 0x57, 0x8a, 0x9d, 0x6d, 0xc6, 0x0a, 0xdb, 0xf9, 0xe9, 0x13,
	0x76, 0x00, 0xc1, 0xad, 0x6c, 0x5a, 0x2b, 0x5d, 0x98, 0x7e, 0x81, 0x11, 0x79, 0xd1, 0x1b, 0xf7,
	0xef, 0x91, 0x87, 0xbd, 0x23, 0xbb, 0x6e, 0xf8, 0xdd, 0xb6, 0x63, 0x5d, 0xe8, 0x6e, 0xa2, 0x2f,
	0x22, 0x3b, 0xfc, 0xe8, 0x98, 0x90, 0x8f, 0xeb, 0x06, 0xd3, 0x1f, 0x03, 0x88, 0x72, 0x34, 0x8d,
	0xaa, 0x0d, 0xba, 0x7d, 0x96, 0x1a, 0x85, 0x45, 0x3f, 0x35, 0xca, 0xbb, 0x94, 0x3d, 0x83, 0xd0,
	0x99, 0x40, 0x23, 0x93, 0xf9, 0xe3, 0x2d, 0xf6, 0xb8, 0x67, 0x90, 0x93, 0x88, 0x3d, 0x87, 0x91,
	0xfb, 0x37, 0x3c, 0xa0, 0x7b, 0x73, 0xaf, 0xda, 0xab, 0xd2, 0x09, 0xc0, 0x02, 0x6d, 0x8e, 0x5f,
	0x57, 0x68, 0xec, 0xfc, 0xd7, 0x10, 0x12, 0xc7, 0x7e, 0x40, 0x7d, 0x23, 0x97, 0xc8, 0xde, 0xc2,
	0xf8, 0x94, 0x0e, 0xc1, 0xee, 0xeb, 0x73, 0x78, 0xb4, 0x85, 0xe8, 0x76, 0x4a, 0xff, 0x63, 0xaf,
	0x21, 0x58, 0xa0, 0xdd, 0xb9, 0xfc, 0x1d, 0x1c, 0x2c, 0xd0, 0x3a, 0xe5, 0xc9, 0x9f, 0xdb, 0xbd,
	0x6b, 0xaf, 0x05, 0xec, 0xb7, 0xbd, 0xce, 0xfd, 0x33, 0xdb, 0xb1, 0xd1, 0xd5, 0x98, 0x3e, 0x49,
	0xaf, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x96, 0xbd, 0xdf, 0x6f, 0xab, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserService service

type UserServiceClient interface {
	Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	Get(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	GetUserAddresses(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	GetUserPhones(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
}

type userServiceClient struct {
	c           client.Client
	serviceName string
}

func NewUserServiceClient(serviceName string, c client.Client) UserServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.user"
	}
	return &userServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *userServiceClient) Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Get(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserAddresses(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.GetUserAddresses", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserPhones(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.GetUserPhones", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	Create(context.Context, *User, *Response) error
	Get(context.Context, *User, *Response) error
	GetUserAddresses(context.Context, *User, *Response) error
	GetUserPhones(context.Context, *User, *Response) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&UserService{hdlr}, opts...))
}

type UserService struct {
	UserServiceHandler
}

func (h *UserService) Create(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.Create(ctx, in, out)
}

func (h *UserService) Get(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.Get(ctx, in, out)
}

func (h *UserService) GetUserAddresses(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.GetUserAddresses(ctx, in, out)
}

func (h *UserService) GetUserPhones(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.GetUserPhones(ctx, in, out)
}
