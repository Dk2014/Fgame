// Code generated by protoc-gen-go. DO NOT EDIT.
// source: body_shield.proto

package ui

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type BodyShieldInfo struct {
	AdvancedId       *int32 `protobuf:"varint,1,req,name=advancedId" json:"advancedId,omitempty"`
	JinjiadanLevel   *int32 `protobuf:"varint,2,req,name=jinjiadanLevel" json:"jinjiadanLevel,omitempty"`
	Progress         *int32 `protobuf:"varint,3,req,name=progress" json:"progress,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *BodyShieldInfo) Reset()                    { *m = BodyShieldInfo{} }
func (m *BodyShieldInfo) String() string            { return proto.CompactTextString(m) }
func (*BodyShieldInfo) ProtoMessage()               {}
func (*BodyShieldInfo) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

func (m *BodyShieldInfo) GetAdvancedId() int32 {
	if m != nil && m.AdvancedId != nil {
		return *m.AdvancedId
	}
	return 0
}

func (m *BodyShieldInfo) GetJinjiadanLevel() int32 {
	if m != nil && m.JinjiadanLevel != nil {
		return *m.JinjiadanLevel
	}
	return 0
}

func (m *BodyShieldInfo) GetProgress() int32 {
	if m != nil && m.Progress != nil {
		return *m.Progress
	}
	return 0
}

type ShieldInfo struct {
	ShieldId         *int32 `protobuf:"varint,1,req,name=shieldId" json:"shieldId,omitempty"`
	Progress         *int32 `protobuf:"varint,2,req,name=progress" json:"progress,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ShieldInfo) Reset()                    { *m = ShieldInfo{} }
func (m *ShieldInfo) String() string            { return proto.CompactTextString(m) }
func (*ShieldInfo) ProtoMessage()               {}
func (*ShieldInfo) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{1} }

func (m *ShieldInfo) GetShieldId() int32 {
	if m != nil && m.ShieldId != nil {
		return *m.ShieldId
	}
	return 0
}

func (m *ShieldInfo) GetProgress() int32 {
	if m != nil && m.Progress != nil {
		return *m.Progress
	}
	return 0
}

type CSBodyShieldGet struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSBodyShieldGet) Reset()                    { *m = CSBodyShieldGet{} }
func (m *CSBodyShieldGet) String() string            { return proto.CompactTextString(m) }
func (*CSBodyShieldGet) ProtoMessage()               {}
func (*CSBodyShieldGet) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{2} }

type SCBodyShieldGet struct {
	AdvancedId       *int32 `protobuf:"varint,1,req,name=advancedId" json:"advancedId,omitempty"`
	JinjiadanLevel   *int32 `protobuf:"varint,2,req,name=jinjiadanLevel" json:"jinjiadanLevel,omitempty"`
	Bless            *int32 `protobuf:"varint,3,opt,name=bless" json:"bless,omitempty"`
	BlessTime        *int64 `protobuf:"varint,4,opt,name=blessTime" json:"blessTime,omitempty"`
	Progress         *int32 `protobuf:"varint,5,req,name=progress" json:"progress,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCBodyShieldGet) Reset()                    { *m = SCBodyShieldGet{} }
func (m *SCBodyShieldGet) String() string            { return proto.CompactTextString(m) }
func (*SCBodyShieldGet) ProtoMessage()               {}
func (*SCBodyShieldGet) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{3} }

func (m *SCBodyShieldGet) GetAdvancedId() int32 {
	if m != nil && m.AdvancedId != nil {
		return *m.AdvancedId
	}
	return 0
}

func (m *SCBodyShieldGet) GetJinjiadanLevel() int32 {
	if m != nil && m.JinjiadanLevel != nil {
		return *m.JinjiadanLevel
	}
	return 0
}

func (m *SCBodyShieldGet) GetBless() int32 {
	if m != nil && m.Bless != nil {
		return *m.Bless
	}
	return 0
}

func (m *SCBodyShieldGet) GetBlessTime() int64 {
	if m != nil && m.BlessTime != nil {
		return *m.BlessTime
	}
	return 0
}

func (m *SCBodyShieldGet) GetProgress() int32 {
	if m != nil && m.Progress != nil {
		return *m.Progress
	}
	return 0
}

type CSBodyShieldJJDan struct {
	Num              *int32 `protobuf:"varint,1,req,name=num" json:"num,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSBodyShieldJJDan) Reset()                    { *m = CSBodyShieldJJDan{} }
func (m *CSBodyShieldJJDan) String() string            { return proto.CompactTextString(m) }
func (*CSBodyShieldJJDan) ProtoMessage()               {}
func (*CSBodyShieldJJDan) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{4} }

func (m *CSBodyShieldJJDan) GetNum() int32 {
	if m != nil && m.Num != nil {
		return *m.Num
	}
	return 0
}

type SCBodyShieldJJDan struct {
	Level            *int32 `protobuf:"varint,1,req,name=level" json:"level,omitempty"`
	Progress         *int32 `protobuf:"varint,2,req,name=progress" json:"progress,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCBodyShieldJJDan) Reset()                    { *m = SCBodyShieldJJDan{} }
func (m *SCBodyShieldJJDan) String() string            { return proto.CompactTextString(m) }
func (*SCBodyShieldJJDan) ProtoMessage()               {}
func (*SCBodyShieldJJDan) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{5} }

func (m *SCBodyShieldJJDan) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *SCBodyShieldJJDan) GetProgress() int32 {
	if m != nil && m.Progress != nil {
		return *m.Progress
	}
	return 0
}

type CSBodyShieldAdvanced struct {
	AutoFlag         *bool  `protobuf:"varint,1,req,name=autoFlag" json:"autoFlag,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSBodyShieldAdvanced) Reset()                    { *m = CSBodyShieldAdvanced{} }
func (m *CSBodyShieldAdvanced) String() string            { return proto.CompactTextString(m) }
func (*CSBodyShieldAdvanced) ProtoMessage()               {}
func (*CSBodyShieldAdvanced) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{6} }

func (m *CSBodyShieldAdvanced) GetAutoFlag() bool {
	if m != nil && m.AutoFlag != nil {
		return *m.AutoFlag
	}
	return false
}

type SCBodyShieldAdvanced struct {
	AdvancedId       *int32 `protobuf:"varint,1,req,name=advancedId" json:"advancedId,omitempty"`
	Bless            *int32 `protobuf:"varint,2,opt,name=bless" json:"bless,omitempty"`
	BlessTime        *int64 `protobuf:"varint,3,opt,name=blessTime" json:"blessTime,omitempty"`
	AdvancedType     *int32 `protobuf:"varint,4,req,name=advancedType" json:"advancedType,omitempty"`
	IsDouble         *bool  `protobuf:"varint,5,opt,name=isDouble,def=0" json:"isDouble,omitempty"`
	TotalBless       *int32 `protobuf:"varint,6,opt,name=totalBless" json:"totalBless,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCBodyShieldAdvanced) Reset()                    { *m = SCBodyShieldAdvanced{} }
func (m *SCBodyShieldAdvanced) String() string            { return proto.CompactTextString(m) }
func (*SCBodyShieldAdvanced) ProtoMessage()               {}
func (*SCBodyShieldAdvanced) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{7} }

const Default_SCBodyShieldAdvanced_IsDouble bool = false

func (m *SCBodyShieldAdvanced) GetAdvancedId() int32 {
	if m != nil && m.AdvancedId != nil {
		return *m.AdvancedId
	}
	return 0
}

func (m *SCBodyShieldAdvanced) GetBless() int32 {
	if m != nil && m.Bless != nil {
		return *m.Bless
	}
	return 0
}

func (m *SCBodyShieldAdvanced) GetBlessTime() int64 {
	if m != nil && m.BlessTime != nil {
		return *m.BlessTime
	}
	return 0
}

func (m *SCBodyShieldAdvanced) GetAdvancedType() int32 {
	if m != nil && m.AdvancedType != nil {
		return *m.AdvancedType
	}
	return 0
}

func (m *SCBodyShieldAdvanced) GetIsDouble() bool {
	if m != nil && m.IsDouble != nil {
		return *m.IsDouble
	}
	return Default_SCBodyShieldAdvanced_IsDouble
}

func (m *SCBodyShieldAdvanced) GetTotalBless() int32 {
	if m != nil && m.TotalBless != nil {
		return *m.TotalBless
	}
	return 0
}

type CSShieldGet struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSShieldGet) Reset()                    { *m = CSShieldGet{} }
func (m *CSShieldGet) String() string            { return proto.CompactTextString(m) }
func (*CSShieldGet) ProtoMessage()               {}
func (*CSShieldGet) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{8} }

type SCShieldGet struct {
	ShieldId         *int32 `protobuf:"varint,1,req,name=shieldId" json:"shieldId,omitempty"`
	Progress         *int32 `protobuf:"varint,2,req,name=progress" json:"progress,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCShieldGet) Reset()                    { *m = SCShieldGet{} }
func (m *SCShieldGet) String() string            { return proto.CompactTextString(m) }
func (*SCShieldGet) ProtoMessage()               {}
func (*SCShieldGet) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{9} }

func (m *SCShieldGet) GetShieldId() int32 {
	if m != nil && m.ShieldId != nil {
		return *m.ShieldId
	}
	return 0
}

func (m *SCShieldGet) GetProgress() int32 {
	if m != nil && m.Progress != nil {
		return *m.Progress
	}
	return 0
}

type CSShieldAdvanced struct {
	AutoFlag         *bool  `protobuf:"varint,1,req,name=autoFlag" json:"autoFlag,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSShieldAdvanced) Reset()                    { *m = CSShieldAdvanced{} }
func (m *CSShieldAdvanced) String() string            { return proto.CompactTextString(m) }
func (*CSShieldAdvanced) ProtoMessage()               {}
func (*CSShieldAdvanced) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{10} }

func (m *CSShieldAdvanced) GetAutoFlag() bool {
	if m != nil && m.AutoFlag != nil {
		return *m.AutoFlag
	}
	return false
}

type SCShieldAdvanced struct {
	ShieldId         *int32 `protobuf:"varint,1,req,name=shieldId" json:"shieldId,omitempty"`
	Progress         *int32 `protobuf:"varint,2,req,name=progress" json:"progress,omitempty"`
	AdvancedType     *int32 `protobuf:"varint,3,req,name=advancedType" json:"advancedType,omitempty"`
	IsDouble         *bool  `protobuf:"varint,4,opt,name=isDouble,def=0" json:"isDouble,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCShieldAdvanced) Reset()                    { *m = SCShieldAdvanced{} }
func (m *SCShieldAdvanced) String() string            { return proto.CompactTextString(m) }
func (*SCShieldAdvanced) ProtoMessage()               {}
func (*SCShieldAdvanced) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{11} }

const Default_SCShieldAdvanced_IsDouble bool = false

func (m *SCShieldAdvanced) GetShieldId() int32 {
	if m != nil && m.ShieldId != nil {
		return *m.ShieldId
	}
	return 0
}

func (m *SCShieldAdvanced) GetProgress() int32 {
	if m != nil && m.Progress != nil {
		return *m.Progress
	}
	return 0
}

func (m *SCShieldAdvanced) GetAdvancedType() int32 {
	if m != nil && m.AdvancedType != nil {
		return *m.AdvancedType
	}
	return 0
}

func (m *SCShieldAdvanced) GetIsDouble() bool {
	if m != nil && m.IsDouble != nil {
		return *m.IsDouble
	}
	return Default_SCShieldAdvanced_IsDouble
}

func init() {
	proto.RegisterType((*BodyShieldInfo)(nil), "ui.BodyShieldInfo")
	proto.RegisterType((*ShieldInfo)(nil), "ui.ShieldInfo")
	proto.RegisterType((*CSBodyShieldGet)(nil), "ui.CSBodyShieldGet")
	proto.RegisterType((*SCBodyShieldGet)(nil), "ui.SCBodyShieldGet")
	proto.RegisterType((*CSBodyShieldJJDan)(nil), "ui.CSBodyShieldJJDan")
	proto.RegisterType((*SCBodyShieldJJDan)(nil), "ui.SCBodyShieldJJDan")
	proto.RegisterType((*CSBodyShieldAdvanced)(nil), "ui.CSBodyShieldAdvanced")
	proto.RegisterType((*SCBodyShieldAdvanced)(nil), "ui.SCBodyShieldAdvanced")
	proto.RegisterType((*CSShieldGet)(nil), "ui.CSShieldGet")
	proto.RegisterType((*SCShieldGet)(nil), "ui.SCShieldGet")
	proto.RegisterType((*CSShieldAdvanced)(nil), "ui.CSShieldAdvanced")
	proto.RegisterType((*SCShieldAdvanced)(nil), "ui.SCShieldAdvanced")
}

func init() { proto.RegisterFile("body_shield.proto", fileDescriptor11) }

var fileDescriptor11 = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xcd, 0x6e, 0xe2, 0x30,
	0x14, 0x85, 0x95, 0x04, 0x8f, 0x32, 0x87, 0x01, 0x92, 0x08, 0xcd, 0x64, 0x19, 0x59, 0xb3, 0xc8,
	0xaa, 0x6a, 0xa5, 0xae, 0xba, 0x2b, 0x41, 0xad, 0x40, 0x55, 0x37, 0x61, 0x5f, 0x99, 0xda, 0xd0,
	0x20, 0x13, 0x53, 0x92, 0x20, 0xf1, 0x14, 0x7d, 0xe5, 0x2a, 0x4e, 0x01, 0x23, 0xfa, 0xbb, 0xf4,
	0xd5, 0xbd, 0xdf, 0x39, 0x9f, 0x0c, 0x7f, 0xaa, 0xf8, 0xf6, 0xa1, 0x78, 0xca, 0x84, 0xe4, 0x67,
	0xab, 0xb5, 0x2a, 0x55, 0x60, 0x57, 0x19, 0xbd, 0x47, 0x77, 0xa0, 0xf8, 0x36, 0xd5, 0xf3, 0x51,
	0x3e, 0x53, 0x41, 0x00, 0x30, 0xbe, 0x61, 0xf9, 0xa3, 0xe0, 0x23, 0x1e, 0x5a, 0x91, 0x1d, 0x93,
	0xe0, 0x2f, 0xba, 0x8b, 0x2c, 0x5f, 0x64, 0x8c, 0xb3, 0xfc, 0x4e, 0x6c, 0x84, 0x0c, 0x6d, 0x3d,
	0xf7, 0xe0, 0xae, 0xd6, 0x6a, 0xbe, 0x16, 0x45, 0x11, 0x3a, 0xf5, 0x84, 0x9e, 0x03, 0x06, 0xcb,
	0x83, 0xdb, 0x24, 0xee, 0x49, 0xe6, 0x85, 0x66, 0x50, 0x1f, 0xbd, 0x24, 0x3d, 0x74, 0xb8, 0x15,
	0x25, 0x7d, 0x46, 0x2f, 0x4d, 0x8e, 0x46, 0x3f, 0x6a, 0xd5, 0x01, 0x99, 0xca, 0xa6, 0x92, 0x15,
	0x93, 0xc0, 0xc7, 0x6f, 0xfd, 0x9c, 0x64, 0x4b, 0x11, 0xb6, 0x22, 0x2b, 0x76, 0x8e, 0x5a, 0x10,
	0xdd, 0x22, 0x82, 0x6f, 0xb6, 0x18, 0x8f, 0x87, 0x2c, 0x0f, 0xda, 0x70, 0xf2, 0x6a, 0xd9, 0xa4,
	0xd1, 0x4b, 0xf8, 0x66, 0xa9, 0x66, 0xa3, 0x03, 0x22, 0x75, 0xf2, 0x47, 0x76, 0x31, 0xfa, 0x26,
	0xf7, 0xfa, 0xcd, 0xa1, 0xde, 0x64, 0x55, 0xa9, 0x6e, 0x24, 0x9b, 0xeb, 0x5b, 0x97, 0xbe, 0x58,
	0xe8, 0x9b, 0x01, 0xfb, 0xd5, 0xf7, 0xd4, 0xf7, 0x8a, 0xf6, 0xa9, 0xa2, 0xa3, 0x15, 0xfb, 0xf8,
	0xb3, 0xbb, 0x9a, 0x6c, 0x57, 0xb5, 0x78, 0x7d, 0xf7, 0x0f, 0x6e, 0x56, 0x0c, 0x55, 0x35, 0x95,
	0x22, 0x24, 0x91, 0x15, 0xbb, 0x57, 0x64, 0xc6, 0x64, 0x21, 0xea, 0x90, 0x52, 0x95, 0x4c, 0x0e,
	0x34, 0xf5, 0x57, 0x4d, 0xa5, 0x1d, 0xb4, 0x93, 0xf4, 0xf0, 0x2b, 0x17, 0x68, 0xa7, 0xc9, 0xe1,
	0x47, 0xbe, 0xf3, 0xb7, 0xff, 0xe1, 0xed, 0x08, 0x9f, 0x98, 0x0b, 0x78, 0x3b, 0xb0, 0xb9, 0xf5,
	0x15, 0xfd, 0x44, 0xd1, 0x39, 0x51, 0x6c, 0x19, 0x8a, 0xaf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6b,
	0x1e, 0x37, 0x39, 0x02, 0x03, 0x00, 0x00,
}
