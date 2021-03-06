// Code generated by protoc-gen-go. DO NOT EDIT.
// source: anqi.proto

package ui

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type AnqiInfo struct {
	AdvancedId       *int32 `protobuf:"varint,1,req,name=advancedId" json:"advancedId,omitempty"`
	AnqiDanLevel     *int32 `protobuf:"varint,2,req,name=anqiDanLevel" json:"anqiDanLevel,omitempty"`
	Progress         *int32 `protobuf:"varint,3,req,name=progress" json:"progress,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *AnqiInfo) Reset()                    { *m = AnqiInfo{} }
func (m *AnqiInfo) String() string            { return proto.CompactTextString(m) }
func (*AnqiInfo) ProtoMessage()               {}
func (*AnqiInfo) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *AnqiInfo) GetAdvancedId() int32 {
	if m != nil && m.AdvancedId != nil {
		return *m.AdvancedId
	}
	return 0
}

func (m *AnqiInfo) GetAnqiDanLevel() int32 {
	if m != nil && m.AnqiDanLevel != nil {
		return *m.AnqiDanLevel
	}
	return 0
}

func (m *AnqiInfo) GetProgress() int32 {
	if m != nil && m.Progress != nil {
		return *m.Progress
	}
	return 0
}

type CSAnqiGet struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSAnqiGet) Reset()                    { *m = CSAnqiGet{} }
func (m *CSAnqiGet) String() string            { return proto.CompactTextString(m) }
func (*CSAnqiGet) ProtoMessage()               {}
func (*CSAnqiGet) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

type SCAnqiGet struct {
	AdvancedId       *int32 `protobuf:"varint,1,req,name=advancedId" json:"advancedId,omitempty"`
	AnqiDanLevel     *int32 `protobuf:"varint,2,req,name=anqiDanLevel" json:"anqiDanLevel,omitempty"`
	Bless            *int32 `protobuf:"varint,3,opt,name=bless" json:"bless,omitempty"`
	BlessTime        *int64 `protobuf:"varint,4,opt,name=blessTime" json:"blessTime,omitempty"`
	Progress         *int32 `protobuf:"varint,5,req,name=progress" json:"progress,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCAnqiGet) Reset()                    { *m = SCAnqiGet{} }
func (m *SCAnqiGet) String() string            { return proto.CompactTextString(m) }
func (*SCAnqiGet) ProtoMessage()               {}
func (*SCAnqiGet) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{2} }

func (m *SCAnqiGet) GetAdvancedId() int32 {
	if m != nil && m.AdvancedId != nil {
		return *m.AdvancedId
	}
	return 0
}

func (m *SCAnqiGet) GetAnqiDanLevel() int32 {
	if m != nil && m.AnqiDanLevel != nil {
		return *m.AnqiDanLevel
	}
	return 0
}

func (m *SCAnqiGet) GetBless() int32 {
	if m != nil && m.Bless != nil {
		return *m.Bless
	}
	return 0
}

func (m *SCAnqiGet) GetBlessTime() int64 {
	if m != nil && m.BlessTime != nil {
		return *m.BlessTime
	}
	return 0
}

func (m *SCAnqiGet) GetProgress() int32 {
	if m != nil && m.Progress != nil {
		return *m.Progress
	}
	return 0
}

type CSAnqiEatDan struct {
	Num              *int32 `protobuf:"varint,1,req,name=num" json:"num,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSAnqiEatDan) Reset()                    { *m = CSAnqiEatDan{} }
func (m *CSAnqiEatDan) String() string            { return proto.CompactTextString(m) }
func (*CSAnqiEatDan) ProtoMessage()               {}
func (*CSAnqiEatDan) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{3} }

func (m *CSAnqiEatDan) GetNum() int32 {
	if m != nil && m.Num != nil {
		return *m.Num
	}
	return 0
}

type SCAnqiEatDan struct {
	Level            *int32 `protobuf:"varint,1,req,name=level" json:"level,omitempty"`
	Progress         *int32 `protobuf:"varint,2,req,name=progress" json:"progress,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCAnqiEatDan) Reset()                    { *m = SCAnqiEatDan{} }
func (m *SCAnqiEatDan) String() string            { return proto.CompactTextString(m) }
func (*SCAnqiEatDan) ProtoMessage()               {}
func (*SCAnqiEatDan) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{4} }

func (m *SCAnqiEatDan) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *SCAnqiEatDan) GetProgress() int32 {
	if m != nil && m.Progress != nil {
		return *m.Progress
	}
	return 0
}

type CSAnqiAdvanced struct {
	AutoFlag         *bool  `protobuf:"varint,1,req,name=autoFlag" json:"autoFlag,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSAnqiAdvanced) Reset()                    { *m = CSAnqiAdvanced{} }
func (m *CSAnqiAdvanced) String() string            { return proto.CompactTextString(m) }
func (*CSAnqiAdvanced) ProtoMessage()               {}
func (*CSAnqiAdvanced) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{5} }

func (m *CSAnqiAdvanced) GetAutoFlag() bool {
	if m != nil && m.AutoFlag != nil {
		return *m.AutoFlag
	}
	return false
}

type SCAnqiAdvanced struct {
	AdvancedId       *int32 `protobuf:"varint,1,req,name=advancedId" json:"advancedId,omitempty"`
	Bless            *int32 `protobuf:"varint,2,opt,name=bless" json:"bless,omitempty"`
	BlessTime        *int64 `protobuf:"varint,3,opt,name=blessTime" json:"blessTime,omitempty"`
	AdvancedType     *int32 `protobuf:"varint,4,req,name=advancedType" json:"advancedType,omitempty"`
	IsDouble         *bool  `protobuf:"varint,5,opt,name=isDouble,def=0" json:"isDouble,omitempty"`
	TotalBless       *int32 `protobuf:"varint,6,opt,name=totalBless" json:"totalBless,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCAnqiAdvanced) Reset()                    { *m = SCAnqiAdvanced{} }
func (m *SCAnqiAdvanced) String() string            { return proto.CompactTextString(m) }
func (*SCAnqiAdvanced) ProtoMessage()               {}
func (*SCAnqiAdvanced) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{6} }

const Default_SCAnqiAdvanced_IsDouble bool = false

func (m *SCAnqiAdvanced) GetAdvancedId() int32 {
	if m != nil && m.AdvancedId != nil {
		return *m.AdvancedId
	}
	return 0
}

func (m *SCAnqiAdvanced) GetBless() int32 {
	if m != nil && m.Bless != nil {
		return *m.Bless
	}
	return 0
}

func (m *SCAnqiAdvanced) GetBlessTime() int64 {
	if m != nil && m.BlessTime != nil {
		return *m.BlessTime
	}
	return 0
}

func (m *SCAnqiAdvanced) GetAdvancedType() int32 {
	if m != nil && m.AdvancedType != nil {
		return *m.AdvancedType
	}
	return 0
}

func (m *SCAnqiAdvanced) GetIsDouble() bool {
	if m != nil && m.IsDouble != nil {
		return *m.IsDouble
	}
	return Default_SCAnqiAdvanced_IsDouble
}

func (m *SCAnqiAdvanced) GetTotalBless() int32 {
	if m != nil && m.TotalBless != nil {
		return *m.TotalBless
	}
	return 0
}

type SCAnqiInfo struct {
	AnqiInfo         *AnqiInfo `protobuf:"bytes,1,req,name=anqiInfo" json:"anqiInfo,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *SCAnqiInfo) Reset()                    { *m = SCAnqiInfo{} }
func (m *SCAnqiInfo) String() string            { return proto.CompactTextString(m) }
func (*SCAnqiInfo) ProtoMessage()               {}
func (*SCAnqiInfo) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{7} }

func (m *SCAnqiInfo) GetAnqiInfo() *AnqiInfo {
	if m != nil {
		return m.AnqiInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*AnqiInfo)(nil), "ui.AnqiInfo")
	proto.RegisterType((*CSAnqiGet)(nil), "ui.CSAnqiGet")
	proto.RegisterType((*SCAnqiGet)(nil), "ui.SCAnqiGet")
	proto.RegisterType((*CSAnqiEatDan)(nil), "ui.CSAnqiEatDan")
	proto.RegisterType((*SCAnqiEatDan)(nil), "ui.SCAnqiEatDan")
	proto.RegisterType((*CSAnqiAdvanced)(nil), "ui.CSAnqiAdvanced")
	proto.RegisterType((*SCAnqiAdvanced)(nil), "ui.SCAnqiAdvanced")
	proto.RegisterType((*SCAnqiInfo)(nil), "ui.SCAnqiInfo")
}

func init() { proto.RegisterFile("anqi.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x1c, 0xc5, 0x49, 0x6b, 0x24, 0xfd, 0xb7, 0x1b, 0x1a, 0x04, 0x0b, 0x82, 0x94, 0x9c, 0x7a, 0x90,
	0x0a, 0x1e, 0xbd, 0xcd, 0xd5, 0xc9, 0xc0, 0xdb, 0xf6, 0x05, 0x32, 0x9b, 0x8d, 0x40, 0x96, 0x74,
	0x6d, 0x3a, 0xf0, 0x23, 0xf8, 0xad, 0x25, 0x49, 0x1d, 0xc8, 0xbc, 0x78, 0x6b, 0xf2, 0xff, 0xbf,
	0xdf, 0x7b, 0xaf, 0x01, 0xe0, 0xfa, 0x20, 0xab, 0xb6, 0x33, 0xd6, 0xd0, 0x68, 0x90, 0x6c, 0x01,
	0x64, 0xa6, 0x0f, 0x72, 0xa9, 0xb7, 0x86, 0x52, 0x00, 0xde, 0x1c, 0xb9, 0xfe, 0x10, 0xcd, 0xb2,
	0xc9, 0x51, 0x11, 0x95, 0x98, 0xde, 0x40, 0xe6, 0x14, 0x35, 0xd7, 0xef, 0xe2, 0x28, 0x54, 0x1e,
	0xf9, 0xdb, 0x2b, 0x20, 0x6d, 0x67, 0x76, 0x9d, 0xe8, 0xfb, 0x3c, 0x76, 0x37, 0x2c, 0x85, 0x64,
	0xbe, 0x72, 0xa4, 0x37, 0x61, 0x99, 0x84, 0x64, 0x35, 0x1f, 0x0f, 0xff, 0xa0, 0x4e, 0x00, 0x6f,
	0x54, 0x40, 0xa2, 0x12, 0xd3, 0x6b, 0x48, 0xfc, 0x71, 0x2d, 0xf7, 0x22, 0xbf, 0x28, 0x50, 0x19,
	0xff, 0xf2, 0xc5, 0xde, 0xf7, 0x0e, 0xb2, 0xe0, 0xfb, 0xca, 0x6d, 0xcd, 0x35, 0x4d, 0x21, 0xd6,
	0xc3, 0x3e, 0xd8, 0xb0, 0x47, 0xc8, 0x42, 0x8e, 0x71, 0x38, 0x01, 0xac, 0xbc, 0x1f, 0x3a, 0x6b,
	0xe1, 0x13, 0x30, 0x06, 0xd3, 0x40, 0x9b, 0x8d, 0x89, 0xdd, 0x0e, 0x1f, 0xac, 0x59, 0x28, 0xbe,
	0xf3, 0x2a, 0xc2, 0xbe, 0x10, 0x4c, 0x03, 0xf5, 0xb4, 0xf4, 0x57, 0xc5, 0x53, 0x99, 0xe8, 0xbc,
	0x4c, 0xec, 0xcb, 0xb8, 0x9f, 0x30, 0xaa, 0xd6, 0x9f, 0xad, 0xab, 0xe8, 0x74, 0xb7, 0x40, 0x64,
	0x5f, 0x9b, 0x61, 0xa3, 0x44, 0x8e, 0x0b, 0x54, 0x92, 0x67, 0xbc, 0xe5, 0xaa, 0x17, 0xce, 0xc4,
	0x1a, 0xcb, 0xd5, 0x8b, 0xa7, 0x5e, 0x3a, 0x2a, 0x7b, 0x00, 0x08, 0x51, 0xfc, 0xfb, 0xdd, 0x03,
	0xe1, 0xe3, 0xb7, 0x0f, 0x91, 0x3e, 0x65, 0xd5, 0x20, 0xab, 0x9f, 0xf9, 0x77, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x79, 0x8a, 0x8c, 0xde, 0xfc, 0x01, 0x00, 0x00,
}
