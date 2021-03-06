// Code generated by protoc-gen-go.
// source: login.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

login message optional from 1000-1999

It is generated from these files:
	login.proto

It has these top-level messages:
	CGLogin
	GCLogin
	CGPing
	GCPing
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import pb1 "fgame/fgame/gm/gamegm/basic/pb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CGLogin struct {
	PlayerId         *int64  `protobuf:"varint,1,req,name=playerId" json:"playerId,omitempty"`
	Token            *string `protobuf:"bytes,2,req,name=token" json:"token,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CGLogin) Reset()                    { *m = CGLogin{} }
func (m *CGLogin) String() string            { return proto.CompactTextString(m) }
func (*CGLogin) ProtoMessage()               {}
func (*CGLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CGLogin) GetPlayerId() int64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

func (m *CGLogin) GetToken() string {
	if m != nil && m.Token != nil {
		return *m.Token
	}
	return ""
}

type GCLogin struct {
	PlayerId         *int64 `protobuf:"varint,1,req,name=playerId" json:"playerId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *GCLogin) Reset()                    { *m = GCLogin{} }
func (m *GCLogin) String() string            { return proto.CompactTextString(m) }
func (*GCLogin) ProtoMessage()               {}
func (*GCLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GCLogin) GetPlayerId() int64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

type CGPing struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CGPing) Reset()                    { *m = CGPing{} }
func (m *CGPing) String() string            { return proto.CompactTextString(m) }
func (*CGPing) ProtoMessage()               {}
func (*CGPing) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type GCPing struct {
	Now              *int64 `protobuf:"varint,1,req,name=now" json:"now,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *GCPing) Reset()                    { *m = GCPing{} }
func (m *GCPing) String() string            { return proto.CompactTextString(m) }
func (*GCPing) ProtoMessage()               {}
func (*GCPing) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GCPing) GetNow() int64 {
	if m != nil && m.Now != nil {
		return *m.Now
	}
	return 0
}

var E_CgLogin = &proto.ExtensionDesc{
	ExtendedType:  (*pb1.Message)(nil),
	ExtensionType: (*CGLogin)(nil),
	Field:         101,
	Name:          "pb.cgLogin",
	Tag:           "bytes,101,opt,name=cgLogin",
	Filename:      "login.proto",
}

var E_GcLogin = &proto.ExtensionDesc{
	ExtendedType:  (*pb1.Message)(nil),
	ExtensionType: (*GCLogin)(nil),
	Field:         102,
	Name:          "pb.gcLogin",
	Tag:           "bytes,102,opt,name=gcLogin",
	Filename:      "login.proto",
}

var E_CgPing = &proto.ExtensionDesc{
	ExtendedType:  (*pb1.Message)(nil),
	ExtensionType: (*CGPing)(nil),
	Field:         103,
	Name:          "pb.cgPing",
	Tag:           "bytes,103,opt,name=cgPing",
	Filename:      "login.proto",
}

var E_GcPing = &proto.ExtensionDesc{
	ExtendedType:  (*pb1.Message)(nil),
	ExtensionType: (*GCPing)(nil),
	Field:         104,
	Name:          "pb.gcPing",
	Tag:           "bytes,104,opt,name=gcPing",
	Filename:      "login.proto",
}

func init() {
	proto.RegisterType((*CGLogin)(nil), "pb.CGLogin")
	proto.RegisterType((*GCLogin)(nil), "pb.GCLogin")
	proto.RegisterType((*CGPing)(nil), "pb.CGPing")
	proto.RegisterType((*GCPing)(nil), "pb.GCPing")
	proto.RegisterExtension(E_CgLogin)
	proto.RegisterExtension(E_GcLogin)
	proto.RegisterExtension(E_CgPing)
	proto.RegisterExtension(E_GcPing)
}

func init() { proto.RegisterFile("login.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x8f, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x59, 0x87, 0xed, 0x7c, 0xbd, 0x48, 0xf0, 0x50, 0x72, 0x2a, 0x01, 0xa1, 0xa7, 0x16,
	0x76, 0x9c, 0xc7, 0x1c, 0x82, 0xa0, 0x20, 0xf9, 0x06, 0x49, 0x8c, 0xcf, 0xe0, 0x48, 0xe2, 0x3a,
	0x10, 0xbf, 0xbd, 0xa4, 0xc9, 0x1c, 0xa2, 0xb0, 0xdb, 0xff, 0xe5, 0xfd, 0x7e, 0xc9, 0x3f, 0xd0,
	0xee, 0x03, 0x3a, 0x3f, 0xc6, 0x43, 0x38, 0x06, 0x52, 0x45, 0x4d, 0xe9, 0x87, 0x8b, 0xca, 0x4d,
	0x5a, 0xcd, 0xce, 0x4c, 0x51, 0xe7, 0x90, 0xf7, 0xec, 0x1e, 0x1a, 0x2e, 0x1e, 0x93, 0x40, 0x28,
	0x6c, 0xe2, 0x5e, 0x7d, 0xd9, 0xc3, 0xc3, 0x4b, 0xb7, 0xea, 0xab, 0x61, 0x2d, 0x7f, 0x66, 0x72,
	0x0b, 0x57, 0xc7, 0xf0, 0x6e, 0x7d, 0x57, 0xf5, 0xd5, 0x70, 0x2d, 0xf3, 0xc0, 0xee, 0xa0, 0x11,
	0xfc, 0xa2, 0xcc, 0x36, 0x50, 0x73, 0xf1, 0xec, 0x3c, 0x32, 0x0a, 0xb5, 0xe0, 0x29, 0x91, 0x1b,
	0x58, 0xfb, 0xf0, 0x59, 0xd0, 0x14, 0x77, 0x5b, 0x68, 0x0c, 0xe6, 0xcb, 0xda, 0x31, 0xea, 0xf1,
	0xc9, 0xce, 0xb3, 0x42, 0xdb, 0xd9, 0x7e, 0x35, 0xb4, 0xdb, 0xe5, 0xa8, 0x34, 0x95, 0x27, 0x30,
	0x39, 0x68, 0xfe, 0x71, 0x5e, 0xcf, 0x4e, 0x29, 0x28, 0x4f, 0xe0, 0x6e, 0x82, 0xda, 0xe0, 0xd2,
	0xe1, 0x97, 0x82, 0x8b, 0x02, 0xf9, 0x99, 0xb4, 0x96, 0x05, 0x4b, 0x02, 0x9a, 0xbf, 0xc2, 0xdb,
	0x59, 0xc8, 0x7f, 0x92, 0x05, 0xfb, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x0a, 0x5f, 0x88, 0x81,
	0x01, 0x00, 0x00,
}
