// Code generated by protoc-gen-go.
// source: proto/service.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	proto/service.proto

It has these top-level messages:
	CreateUserRequest
	GetUserRequest
	GreetUserRequest
	GreetUserResponse
	User
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/empty"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateUserRequest struct {
	User *User `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *CreateUserRequest) Reset()                    { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string            { return proto1.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()               {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type GetUserRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
}

func (m *GetUserRequest) Reset()                    { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string            { return proto1.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()               {}
func (*GetUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type GreetUserRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Greeting string `protobuf:"bytes,2,opt,name=greeting" json:"greeting,omitempty"`
}

func (m *GreetUserRequest) Reset()                    { *m = GreetUserRequest{} }
func (m *GreetUserRequest) String() string            { return proto1.CompactTextString(m) }
func (*GreetUserRequest) ProtoMessage()               {}
func (*GreetUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GreetUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GreetUserRequest) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

type GreetUserResponse struct {
	Greeting string `protobuf:"bytes,1,opt,name=greeting" json:"greeting,omitempty"`
}

func (m *GreetUserResponse) Reset()                    { *m = GreetUserResponse{} }
func (m *GreetUserResponse) String() string            { return proto1.CompactTextString(m) }
func (*GreetUserResponse) ProtoMessage()               {}
func (*GreetUserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GreetUserResponse) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

type User struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Role     string `protobuf:"bytes,2,opt,name=role" json:"role,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto1.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func init() {
	proto1.RegisterType((*CreateUserRequest)(nil), "gowroc.grpcgateway.grpc.service.CreateUserRequest")
	proto1.RegisterType((*GetUserRequest)(nil), "gowroc.grpcgateway.grpc.service.GetUserRequest")
	proto1.RegisterType((*GreetUserRequest)(nil), "gowroc.grpcgateway.grpc.service.GreetUserRequest")
	proto1.RegisterType((*GreetUserResponse)(nil), "gowroc.grpcgateway.grpc.service.GreetUserResponse")
	proto1.RegisterType((*User)(nil), "gowroc.grpcgateway.grpc.service.User")
}

func init() { proto1.RegisterFile("proto/service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xde, 0x46, 0x51, 0xf7, 0x14, 0x71, 0x11, 0x64, 0xd4, 0x83, 0x12, 0x10, 0x3c, 0xcc, 0x04,
	0x2b, 0x08, 0x5e, 0x15, 0x19, 0x78, 0xf0, 0xb0, 0xe1, 0x65, 0xb7, 0xae, 0x3c, 0x63, 0xa1, 0x6d,
	0x62, 0x92, 0x6e, 0xec, 0x0f, 0xf4, 0xff, 0x92, 0xa4, 0xdd, 0xdc, 0x26, 0xd8, 0x79, 0xca, 0x7b,
	0xbc, 0xef, 0x07, 0x7c, 0x5f, 0xe0, 0x54, 0x69, 0x69, 0x25, 0x37, 0xa8, 0x67, 0x69, 0x82, 0xcc,
	0x6f, 0xe4, 0x42, 0xc8, 0xb9, 0x96, 0x09, 0x13, 0x5a, 0x25, 0x22, 0xb6, 0x38, 0x8f, 0x17, 0x7e,
	0x66, 0x35, 0x2c, 0x3c, 0x17, 0x52, 0x8a, 0x0c, 0xb9, 0x87, 0x4f, 0xcb, 0x77, 0x8e, 0xb9, 0xb2,
	0x8b, 0x8a, 0x4d, 0x5f, 0xa1, 0xf7, 0xa4, 0x31, 0xb6, 0xf8, 0x66, 0x50, 0x8f, 0xf0, 0xb3, 0x44,
	0x63, 0xc9, 0x03, 0x04, 0xa5, 0x41, 0xdd, 0x6f, 0x5f, 0xb6, 0xaf, 0x0f, 0xa3, 0x2b, 0xd6, 0xe0,
	0xc0, 0x3c, 0xd7, 0x53, 0xe8, 0x00, 0x8e, 0x87, 0x68, 0xd7, 0xc5, 0x42, 0x38, 0x70, 0x97, 0x22,
	0xce, 0xd1, 0x0b, 0x76, 0x47, 0xab, 0x9d, 0xbe, 0xc0, 0xc9, 0x50, 0xe3, 0xce, 0x78, 0x77, 0x13,
	0x0e, 0x9f, 0x16, 0xa2, 0xdf, 0xa9, 0x6e, 0xcb, 0x9d, 0x72, 0xe8, 0xad, 0x69, 0x19, 0x25, 0x0b,
	0xb3, 0x49, 0x68, 0x6f, 0x11, 0xee, 0x21, 0x70, 0xd8, 0x3f, 0x0d, 0x09, 0x04, 0x5a, 0x66, 0x58,
	0x9b, 0xf9, 0x39, 0xfa, 0xea, 0xc0, 0xd1, 0x38, 0xcd, 0x55, 0x86, 0x63, 0xd4, 0x33, 0xd4, 0x64,
	0x02, 0xf0, 0x93, 0x21, 0x89, 0x1a, 0xe3, 0xfa, 0x15, 0x78, 0x78, 0xc6, 0xaa, 0x8e, 0xd8, 0xb2,
	0x23, 0xf6, 0xec, 0x3a, 0xa2, 0x2d, 0x92, 0xc0, 0x7e, 0x9d, 0x27, 0xe1, 0x8d, 0xc2, 0x9b, 0xc9,
	0x87, 0xbb, 0x15, 0x47, 0x5b, 0xc4, 0x42, 0x77, 0x15, 0x1d, 0xb9, 0x6d, 0xb6, 0xd9, 0xaa, 0x2c,
	0x8c, 0xfe, 0x43, 0xa9, 0x9a, 0xa1, 0xad, 0x47, 0x36, 0x19, 0x88, 0xd4, 0x7e, 0x94, 0x53, 0x96,
	0xc8, 0x9c, 0x57, 0x0a, 0x3c, 0x47, 0xb4, 0xa5, 0x32, 0xdc, 0xb1, 0x6f, 0x6a, 0xa9, 0xfa, 0xe7,
	0xee, 0xf9, 0xe7, 0xee, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x4d, 0x01, 0x4d, 0xd6, 0x06, 0x03, 0x00,
	0x00,
}
