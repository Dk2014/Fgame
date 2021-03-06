// Code generated by protoc-gen-go. DO NOT EDIT.
// source: base.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	base.proto

It has these top-level messages:
	LogMessage
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LogMessage struct {
	LogName          *string `protobuf:"bytes,1,req,name=logName" json:"logName,omitempty"`
	Content          []byte  `protobuf:"bytes,2,req,name=content" json:"content,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LogMessage) Reset()                    { *m = LogMessage{} }
func (m *LogMessage) String() string            { return proto.CompactTextString(m) }
func (*LogMessage) ProtoMessage()               {}
func (*LogMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LogMessage) GetLogName() string {
	if m != nil && m.LogName != nil {
		return *m.LogName
	}
	return ""
}

func (m *LogMessage) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func init() {
	proto.RegisterType((*LogMessage)(nil), "pb.LogMessage")
}

func init() { proto.RegisterFile("base.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 86 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4a, 0x2c, 0x4e,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xd2, 0xe3, 0xe2, 0xf2, 0xc9,
	0x4f, 0xf7, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0xe2, 0xe7, 0x62, 0xcf, 0xc9, 0x4f, 0xf7,
	0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd2, 0xe0, 0x04, 0x09, 0x24, 0xe7, 0xe7, 0x95, 0xa4,
	0xe6, 0x95, 0x48, 0x30, 0x29, 0x30, 0x69, 0xf0, 0x00, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x05,
	0xe8, 0xd9, 0x40, 0x00, 0x00, 0x00,
}
