// Code generated by protoc-gen-go. DO NOT EDIT.
// source: xuechi.proto

package ui

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CSXueChiGet struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSXueChiGet) Reset()                    { *m = CSXueChiGet{} }
func (m *CSXueChiGet) String() string            { return proto.CompactTextString(m) }
func (*CSXueChiGet) ProtoMessage()               {}
func (*CSXueChiGet) Descriptor() ([]byte, []int) { return fileDescriptor138, []int{0} }

type SCXueChiGet struct {
	BloodLine        *int32 `protobuf:"varint,1,req,name=bloodLine" json:"bloodLine,omitempty"`
	Blood            *int64 `protobuf:"varint,2,req,name=blood" json:"blood,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCXueChiGet) Reset()                    { *m = SCXueChiGet{} }
func (m *SCXueChiGet) String() string            { return proto.CompactTextString(m) }
func (*SCXueChiGet) ProtoMessage()               {}
func (*SCXueChiGet) Descriptor() ([]byte, []int) { return fileDescriptor138, []int{1} }

func (m *SCXueChiGet) GetBloodLine() int32 {
	if m != nil && m.BloodLine != nil {
		return *m.BloodLine
	}
	return 0
}

func (m *SCXueChiGet) GetBlood() int64 {
	if m != nil && m.Blood != nil {
		return *m.Blood
	}
	return 0
}

type CSXueChiBloodLine struct {
	BloodLine        *int32 `protobuf:"varint,1,req,name=bloodLine" json:"bloodLine,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSXueChiBloodLine) Reset()                    { *m = CSXueChiBloodLine{} }
func (m *CSXueChiBloodLine) String() string            { return proto.CompactTextString(m) }
func (*CSXueChiBloodLine) ProtoMessage()               {}
func (*CSXueChiBloodLine) Descriptor() ([]byte, []int) { return fileDescriptor138, []int{2} }

func (m *CSXueChiBloodLine) GetBloodLine() int32 {
	if m != nil && m.BloodLine != nil {
		return *m.BloodLine
	}
	return 0
}

type SCXueChiBloodLine struct {
	BloodLine        *int32 `protobuf:"varint,1,req,name=bloodLine" json:"bloodLine,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCXueChiBloodLine) Reset()                    { *m = SCXueChiBloodLine{} }
func (m *SCXueChiBloodLine) String() string            { return proto.CompactTextString(m) }
func (*SCXueChiBloodLine) ProtoMessage()               {}
func (*SCXueChiBloodLine) Descriptor() ([]byte, []int) { return fileDescriptor138, []int{3} }

func (m *SCXueChiBloodLine) GetBloodLine() int32 {
	if m != nil && m.BloodLine != nil {
		return *m.BloodLine
	}
	return 0
}

type CSXueChiAutoBuy struct {
	ItemId           *int32 `protobuf:"varint,1,req,name=itemId" json:"itemId,omitempty"`
	ItemNum          *int32 `protobuf:"varint,2,req,name=itemNum" json:"itemNum,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSXueChiAutoBuy) Reset()                    { *m = CSXueChiAutoBuy{} }
func (m *CSXueChiAutoBuy) String() string            { return proto.CompactTextString(m) }
func (*CSXueChiAutoBuy) ProtoMessage()               {}
func (*CSXueChiAutoBuy) Descriptor() ([]byte, []int) { return fileDescriptor138, []int{4} }

func (m *CSXueChiAutoBuy) GetItemId() int32 {
	if m != nil && m.ItemId != nil {
		return *m.ItemId
	}
	return 0
}

func (m *CSXueChiAutoBuy) GetItemNum() int32 {
	if m != nil && m.ItemNum != nil {
		return *m.ItemNum
	}
	return 0
}

type SCXueChiAutoBuy struct {
	ItemId           *int32 `protobuf:"varint,1,req,name=itemId" json:"itemId,omitempty"`
	ItemNum          *int32 `protobuf:"varint,2,req,name=itemNum" json:"itemNum,omitempty"`
	AddBlood         *int64 `protobuf:"varint,3,req,name=addBlood" json:"addBlood,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCXueChiAutoBuy) Reset()                    { *m = SCXueChiAutoBuy{} }
func (m *SCXueChiAutoBuy) String() string            { return proto.CompactTextString(m) }
func (*SCXueChiAutoBuy) ProtoMessage()               {}
func (*SCXueChiAutoBuy) Descriptor() ([]byte, []int) { return fileDescriptor138, []int{5} }

func (m *SCXueChiAutoBuy) GetItemId() int32 {
	if m != nil && m.ItemId != nil {
		return *m.ItemId
	}
	return 0
}

func (m *SCXueChiAutoBuy) GetItemNum() int32 {
	if m != nil && m.ItemNum != nil {
		return *m.ItemNum
	}
	return 0
}

func (m *SCXueChiAutoBuy) GetAddBlood() int64 {
	if m != nil && m.AddBlood != nil {
		return *m.AddBlood
	}
	return 0
}

type SCXueChiBlood struct {
	Blood            *int64 `protobuf:"varint,1,req,name=blood" json:"blood,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCXueChiBlood) Reset()                    { *m = SCXueChiBlood{} }
func (m *SCXueChiBlood) String() string            { return proto.CompactTextString(m) }
func (*SCXueChiBlood) ProtoMessage()               {}
func (*SCXueChiBlood) Descriptor() ([]byte, []int) { return fileDescriptor138, []int{6} }

func (m *SCXueChiBlood) GetBlood() int64 {
	if m != nil && m.Blood != nil {
		return *m.Blood
	}
	return 0
}

func init() {
	proto.RegisterType((*CSXueChiGet)(nil), "ui.CSXueChiGet")
	proto.RegisterType((*SCXueChiGet)(nil), "ui.SCXueChiGet")
	proto.RegisterType((*CSXueChiBloodLine)(nil), "ui.CSXueChiBloodLine")
	proto.RegisterType((*SCXueChiBloodLine)(nil), "ui.SCXueChiBloodLine")
	proto.RegisterType((*CSXueChiAutoBuy)(nil), "ui.CSXueChiAutoBuy")
	proto.RegisterType((*SCXueChiAutoBuy)(nil), "ui.SCXueChiAutoBuy")
	proto.RegisterType((*SCXueChiBlood)(nil), "ui.SCXueChiBlood")
}

func init() { proto.RegisterFile("xuechi.proto", fileDescriptor138) }

var fileDescriptor138 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xa9, 0x28, 0x4d, 0x4d,
	0xce, 0xc8, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0xcd, 0x54, 0xe2, 0xe5, 0xe2,
	0x76, 0x0e, 0x8e, 0x28, 0x4d, 0x75, 0xce, 0xc8, 0x74, 0x4f, 0x2d, 0x51, 0xd2, 0xe7, 0xe2, 0x0e,
	0x76, 0x86, 0x73, 0x85, 0x04, 0xb9, 0x38, 0x93, 0x72, 0xf2, 0xf3, 0x53, 0x7c, 0x32, 0xf3, 0x52,
	0x25, 0x18, 0x15, 0x98, 0x34, 0x58, 0x85, 0x78, 0xb9, 0x58, 0xc1, 0x42, 0x12, 0x4c, 0x0a, 0x4c,
	0x1a, 0xcc, 0x4a, 0x6a, 0x5c, 0x82, 0x30, 0xfd, 0x4e, 0x30, 0x95, 0x58, 0xb4, 0x81, 0xd4, 0xc1,
	0x0c, 0xc6, 0xab, 0xce, 0x88, 0x8b, 0x1f, 0x66, 0x9e, 0x63, 0x69, 0x49, 0xbe, 0x53, 0x69, 0xa5,
	0x10, 0x1f, 0x17, 0x5b, 0x66, 0x49, 0x6a, 0xae, 0x67, 0x0a, 0xd4, 0x05, 0xfc, 0x5c, 0xec, 0x20,
	0xbe, 0x5f, 0x69, 0x2e, 0xd8, 0x0d, 0xac, 0x4a, 0x2e, 0x5c, 0xfc, 0x30, 0xb3, 0x89, 0xd5, 0x23,
	0x24, 0xc0, 0xc5, 0x91, 0x98, 0x92, 0x02, 0x76, 0x8a, 0x04, 0x33, 0xd8, 0x27, 0x72, 0x5c, 0xbc,
	0x28, 0x2e, 0x44, 0xf8, 0x14, 0x64, 0x04, 0x33, 0x20, 0x00, 0x00, 0xff, 0xff, 0xdf, 0x97, 0x79,
	0x24, 0x3c, 0x01, 0x00, 0x00,
}
