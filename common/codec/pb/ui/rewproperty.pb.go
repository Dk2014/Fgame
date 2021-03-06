// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rewproperty.proto

package ui

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RewProperty struct {
	Exp              *int32 `protobuf:"varint,1,req,name=exp" json:"exp,omitempty"`
	ExpPoint         *int32 `protobuf:"varint,2,req,name=expPoint" json:"expPoint,omitempty"`
	Silver           *int32 `protobuf:"varint,3,req,name=silver" json:"silver,omitempty"`
	Gold             *int32 `protobuf:"varint,4,req,name=gold" json:"gold,omitempty"`
	BindGold         *int32 `protobuf:"varint,5,req,name=bindGold" json:"bindGold,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RewProperty) Reset()                    { *m = RewProperty{} }
func (m *RewProperty) String() string            { return proto.CompactTextString(m) }
func (*RewProperty) ProtoMessage()               {}
func (*RewProperty) Descriptor() ([]byte, []int) { return fileDescriptor89, []int{0} }

func (m *RewProperty) GetExp() int32 {
	if m != nil && m.Exp != nil {
		return *m.Exp
	}
	return 0
}

func (m *RewProperty) GetExpPoint() int32 {
	if m != nil && m.ExpPoint != nil {
		return *m.ExpPoint
	}
	return 0
}

func (m *RewProperty) GetSilver() int32 {
	if m != nil && m.Silver != nil {
		return *m.Silver
	}
	return 0
}

func (m *RewProperty) GetGold() int32 {
	if m != nil && m.Gold != nil {
		return *m.Gold
	}
	return 0
}

func (m *RewProperty) GetBindGold() int32 {
	if m != nil && m.BindGold != nil {
		return *m.BindGold
	}
	return 0
}

func init() {
	proto.RegisterType((*RewProperty)(nil), "ui.RewProperty")
}

func init() { proto.RegisterFile("rewproperty.proto", fileDescriptor89) }

var fileDescriptor89 = []byte{
	// 119 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x4a, 0x2d, 0x2f,
	0x28, 0xca, 0x2f, 0x48, 0x2d, 0x2a, 0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a,
	0xcd, 0x54, 0x8a, 0xe1, 0xe2, 0x0e, 0x4a, 0x2d, 0x0f, 0x80, 0x4a, 0x08, 0x71, 0x73, 0x31, 0xa7,
	0x56, 0x14, 0x48, 0x30, 0x2a, 0x30, 0x69, 0xb0, 0x0a, 0x09, 0x70, 0x71, 0xa4, 0x56, 0x14, 0x04,
	0xe4, 0x67, 0xe6, 0x95, 0x48, 0x30, 0x81, 0x45, 0xf8, 0xb8, 0xd8, 0x8a, 0x33, 0x73, 0xca, 0x52,
	0x8b, 0x24, 0x98, 0xc1, 0x7c, 0x1e, 0x2e, 0x96, 0xf4, 0xfc, 0x9c, 0x14, 0x09, 0x16, 0x98, 0xfa,
	0xa4, 0xcc, 0xbc, 0x14, 0x77, 0x90, 0x08, 0x2b, 0x48, 0x04, 0x10, 0x00, 0x00, 0xff, 0xff, 0xed,
	0x1a, 0x18, 0xb5, 0x75, 0x00, 0x00, 0x00,
}
