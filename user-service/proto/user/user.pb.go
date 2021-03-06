// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

package user

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
	OrganizationId       string           `protobuf:"bytes,11,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
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

func (m *User) GetOrganizationId() string {
	if m != nil {
		return m.OrganizationId
	}
	return ""
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
	CustomerId           string   `protobuf:"bytes,6,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
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

func (m *PostalAddress) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

type Phone struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PhoneNumber          string   `protobuf:"bytes,2,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Ext                  string   `protobuf:"bytes,3,opt,name=ext,proto3" json:"ext,omitempty"`
	PhoneType            string   `protobuf:"bytes,4,opt,name=phone_type,json=phoneType,proto3" json:"phone_type,omitempty"`
	CustomerId           string   `protobuf:"bytes,5,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
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

func (m *Phone) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

type Token struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	CustomerId           string   `protobuf:"bytes,3,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Errors               []*Error `protobuf:"bytes,4,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{4}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Token) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Token) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Token) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{5}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type Response struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Users                []*User  `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	Errors               []*Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	Token                *Token   `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{6}
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

func (m *Response) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *Response) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b283a848145d6b7, []int{7}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func init() {
	proto.RegisterType((*User)(nil), "user.User")
	proto.RegisterType((*Organization)(nil), "user.Organization")
	proto.RegisterType((*PostalAddress)(nil), "user.PostalAddress")
	proto.RegisterType((*Phone)(nil), "user.Phone")
	proto.RegisterType((*Token)(nil), "user.Token")
	proto.RegisterType((*Error)(nil), "user.Error")
	proto.RegisterType((*Response)(nil), "user.Response")
	proto.RegisterType((*Request)(nil), "user.Request")
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor_9b283a848145d6b7) }

var fileDescriptor_9b283a848145d6b7 = []byte{
	// 633 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xcf, 0x6e, 0xd3, 0x4c,
	0x10, 0x6f, 0x12, 0x3b, 0x8d, 0xc7, 0x4d, 0xbf, 0x6a, 0x3f, 0x90, 0xac, 0x22, 0x20, 0x75, 0x25,
	0x28, 0x42, 0x14, 0x28, 0x67, 0x0e, 0x11, 0x42, 0x55, 0x2f, 0xa5, 0x32, 0x2d, 0xd7, 0xca, 0x8d,
	0x07, 0xba, 0xc2, 0xf1, 0x9a, 0xdd, 0x4d, 0xa1, 0x7d, 0x06, 0x8e, 0x88, 0xc7, 0xe1, 0x11, 0x78,
	0x26, 0x34, 0xb3, 0x6b, 0x30, 0x0e, 0xd0, 0x4b, 0x32, 0xf3, 0xfb, 0xcd, 0xff, 0x99, 0x35, 0xdc,
	0xac, 0xb5, 0xb2, 0xea, 0xf1, 0xc2, 0xa0, 0xe6, 0x9f, 0x5d, 0xd6, 0x45, 0x40, 0x72, 0xfa, 0xbd,
	0x0f, 0xc1, 0x89, 0x41, 0x2d, 0xd6, 0xa1, 0x2f, 0x8b, 0xa4, 0x37, 0xe9, 0xed, 0x44, 0x59, 0x5f,
	0x16, 0xe2, 0x36, 0xc0, 0x5b, 0xa9, 0x8d, 0x3d, 0xad, 0xf2, 0x39, 0x26, 0x7d, 0xc6, 0x23, 0x46,
	0x0e, 0xf3, 0x39, 0x8a, 0xbb, 0x10, 0xcf, 0x65, 0x51, 0x94, 0xe8, 0xf8, 0x01, 0xf3, 0xe0, 0x20,
	0x36, 0xb8, 0x05, 0x51, 0x99, 0x37, 0xee, 0x01, 0xd3, 0x23, 0x02, 0x98, 0xbc, 0x01, 0x21, 0xce,
	0x73, 0x59, 0x26, 0x21, 0x13, 0x4e, 0x11, 0x9b, 0x30, 0xaa, 0x73, 0x63, 0x3e, 0x2a, 0x5d, 0x24,
	0x43, 0xe7, 0xd1, 0xe8, 0x62, 0x1b, 0x86, 0xf5, 0xb9, 0xaa, 0xd0, 0x24, 0xab, 0x93, 0xc1, 0x4e,
	0xbc, 0x17, 0xef, 0x72, 0x2b, 0x47, 0x84, 0x65, 0x9e, 0x12, 0x4f, 0x21, 0xca, 0x8b, 0x42, 0xa3,
	0x31, 0x68, 0x92, 0x11, 0xdb, 0xfd, 0xef, 0xed, 0x94, 0xb1, 0x79, 0x39, 0x75, 0x64, 0xf6, 0xcb,
	0x4a, 0x08, 0x08, 0xec, 0x65, 0x8d, 0x49, 0xc4, 0xf9, 0x58, 0x26, 0x4c, 0xab, 0x12, 0x13, 0x70,
	0x18, 0xc9, 0xe2, 0x3e, 0xfc, 0xa7, 0xf4, 0xbb, 0xbc, 0x92, 0x57, 0xb9, 0x95, 0xaa, 0x3a, 0x95,
	0x45, 0x12, 0x33, 0xbd, 0xde, 0x86, 0x0f, 0x8a, 0xf4, 0x04, 0xd6, 0x5e, 0xb5, 0x90, 0xa5, 0xb9,
	0x0a, 0x08, 0x5a, 0x13, 0x65, 0x59, 0x6c, 0xc1, 0x1a, 0x77, 0x70, 0x5a, 0x2d, 0xe6, 0x67, 0xa8,
	0xfd, 0x34, 0x63, 0xc6, 0x0e, 0x19, 0x4a, 0xbf, 0xf6, 0x60, 0xfc, 0x5b, 0x13, 0x4b, 0x81, 0x13,
	0x58, 0xf5, 0x6d, 0xf9, 0xd8, 0x8d, 0x4a, 0x29, 0x67, 0xd2, 0x5e, 0xfa, 0xb0, 0x2c, 0xd3, 0x06,
	0x8c, 0xcd, 0x6d, 0xb3, 0x1a, 0xa7, 0x88, 0x0d, 0x18, 0x5c, 0xc9, 0xda, 0x6f, 0x85, 0x44, 0xda,
	0xf3, 0x6c, 0x61, 0xac, 0x9a, 0xa3, 0xa6, 0x9e, 0xdd, 0x5a, 0xa0, 0x81, 0x0e, 0x8a, 0xf4, 0x73,
	0x0f, 0x42, 0xde, 0xc2, 0x52, 0x41, 0xdd, 0xae, 0xfa, 0x4b, 0x5d, 0x51, 0x3e, 0xfc, 0x64, 0x7d,
	0x61, 0x24, 0xd2, 0xd9, 0x39, 0x27, 0xde, 0x8a, 0x2b, 0x2e, 0x62, 0xe4, 0x98, 0x56, 0xd3, 0x29,
	0x27, 0x5c, 0x2a, 0xe7, 0x12, 0xc2, 0x63, 0xf5, 0x1e, 0x2b, 0x6a, 0xd0, 0x92, 0xe0, 0x0b, 0x72,
	0x0a, 0xa1, 0x17, 0x79, 0x29, 0x0b, 0x2e, 0x66, 0x94, 0x39, 0xa5, 0x1b, 0x75, 0xd0, 0x8d, 0x4a,
	0xd7, 0x87, 0x5a, 0x2b, 0x6d, 0x92, 0xa0, 0x7d, 0x7d, 0x2f, 0x09, 0xcb, 0x3c, 0x95, 0x3e, 0x87,
	0x90, 0x01, 0x9e, 0xb7, 0x2a, 0x90, 0x33, 0x87, 0x19, 0xcb, 0x62, 0x02, 0x71, 0x81, 0x66, 0xa6,
	0x65, 0x4d, 0x57, 0xd1, 0xcc, 0xa2, 0x05, 0xa5, 0x5f, 0x7a, 0x30, 0xca, 0xd0, 0xd4, 0xaa, 0x32,
	0x28, 0xee, 0x00, 0x3f, 0x4f, 0x0e, 0x11, 0xef, 0x81, 0x4b, 0x47, 0xef, 0x34, 0x63, 0x5c, 0x4c,
	0x20, 0xa4, 0x7f, 0x5a, 0xf5, 0xa0, 0x63, 0xe0, 0x88, 0x56, 0xc9, 0x83, 0xbf, 0x96, 0x2c, 0xb6,
	0x9a, 0x21, 0x05, 0x9c, 0xc7, 0xdb, 0xf0, 0x00, 0xfd, 0xc4, 0xd2, 0x08, 0x56, 0x33, 0xfc, 0xb0,
	0x40, 0x63, 0xf7, 0xbe, 0xf5, 0x21, 0xa6, 0x14, 0xaf, 0x51, 0x5f, 0xc8, 0x19, 0x8a, 0x7b, 0x30,
	0x7c, 0xa1, 0x91, 0xee, 0xa6, 0x95, 0x7f, 0x73, 0xdd, 0xc9, 0x4d, 0x2b, 0xe9, 0x8a, 0xd8, 0x86,
	0xc1, 0x3e, 0xda, 0x6b, 0x8c, 0x1e, 0xc0, 0x70, 0x1f, 0xed, 0xb4, 0x2c, 0xc5, 0xb8, 0xe1, 0x38,
	0xeb, 0x1f, 0x4c, 0x9f, 0xc0, 0xc6, 0x3e, 0x5a, 0x8a, 0x33, 0xfd, 0xf9, 0x8e, 0xff, 0x1d, 0xfc,
	0x11, 0x8c, 0xbd, 0xc7, 0x91, 0xfb, 0x52, 0x5c, 0x57, 0xcb, 0xda, 0x74, 0x61, 0xcf, 0xb1, 0xb2,
	0x72, 0xd6, 0x6d, 0xaf, 0x3d, 0xa3, 0x74, 0x45, 0x3c, 0x84, 0xf1, 0x1b, 0xba, 0xa1, 0xdc, 0xa2,
	0xbb, 0xbb, 0x36, 0xdf, 0x31, 0x3e, 0x1b, 0xf2, 0x97, 0xf7, 0xd9, 0x8f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xae, 0xd8, 0xdb, 0x94, 0x92, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserService service

type UserServiceClient interface {
	Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	Get(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	GetUserAddresses(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	GetUserPhones(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	Authenticate(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error)
	ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error)
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
		serviceName = "user"
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

func (c *userServiceClient) GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.GetAll", in)
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

func (c *userServiceClient) Authenticate(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Authenticate", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.ValidateToken", in)
	out := new(Token)
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
	GetAll(context.Context, *Request, *Response) error
	GetUserAddresses(context.Context, *User, *Response) error
	GetUserPhones(context.Context, *User, *Response) error
	Authenticate(context.Context, *User, *Token) error
	ValidateToken(context.Context, *Token, *Token) error
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

func (h *UserService) GetAll(ctx context.Context, in *Request, out *Response) error {
	return h.UserServiceHandler.GetAll(ctx, in, out)
}

func (h *UserService) GetUserAddresses(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.GetUserAddresses(ctx, in, out)
}

func (h *UserService) GetUserPhones(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.GetUserPhones(ctx, in, out)
}

func (h *UserService) Authenticate(ctx context.Context, in *User, out *Token) error {
	return h.UserServiceHandler.Authenticate(ctx, in, out)
}

func (h *UserService) ValidateToken(ctx context.Context, in *Token, out *Token) error {
	return h.UserServiceHandler.ValidateToken(ctx, in, out)
}
