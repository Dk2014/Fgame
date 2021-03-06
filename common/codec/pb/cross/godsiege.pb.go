// Code generated by protoc-gen-go. DO NOT EDIT.
// source: godsiege.proto

package cross

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SIGodSiegeAttend struct {
	GodType          *int32 `protobuf:"varint,1,req,name=godType" json:"godType,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SIGodSiegeAttend) Reset()                    { *m = SIGodSiegeAttend{} }
func (m *SIGodSiegeAttend) String() string            { return proto.CompactTextString(m) }
func (*SIGodSiegeAttend) ProtoMessage()               {}
func (*SIGodSiegeAttend) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

func (m *SIGodSiegeAttend) GetGodType() int32 {
	if m != nil && m.GodType != nil {
		return *m.GodType
	}
	return 0
}

type ISGodSiegeAttend struct {
	GodType          *int32 `protobuf:"varint,1,req,name=godType" json:"godType,omitempty"`
	IsLineUp         *bool  `protobuf:"varint,2,req,name=isLineUp" json:"isLineUp,omitempty"`
	BeforeNum        *int32 `protobuf:"varint,3,req,name=beforeNum" json:"beforeNum,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ISGodSiegeAttend) Reset()                    { *m = ISGodSiegeAttend{} }
func (m *ISGodSiegeAttend) String() string            { return proto.CompactTextString(m) }
func (*ISGodSiegeAttend) ProtoMessage()               {}
func (*ISGodSiegeAttend) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{1} }

func (m *ISGodSiegeAttend) GetGodType() int32 {
	if m != nil && m.GodType != nil {
		return *m.GodType
	}
	return 0
}

func (m *ISGodSiegeAttend) GetIsLineUp() bool {
	if m != nil && m.IsLineUp != nil {
		return *m.IsLineUp
	}
	return false
}

func (m *ISGodSiegeAttend) GetBeforeNum() int32 {
	if m != nil && m.BeforeNum != nil {
		return *m.BeforeNum
	}
	return 0
}

type ISGodSiegeLineUpSuccess struct {
	GodType          *int32 `protobuf:"varint,1,req,name=godType" json:"godType,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ISGodSiegeLineUpSuccess) Reset()                    { *m = ISGodSiegeLineUpSuccess{} }
func (m *ISGodSiegeLineUpSuccess) String() string            { return proto.CompactTextString(m) }
func (*ISGodSiegeLineUpSuccess) ProtoMessage()               {}
func (*ISGodSiegeLineUpSuccess) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{2} }

func (m *ISGodSiegeLineUpSuccess) GetGodType() int32 {
	if m != nil && m.GodType != nil {
		return *m.GodType
	}
	return 0
}

type SIGodSiegeLineUpSuccess struct {
	GodType          *int32 `protobuf:"varint,1,req,name=godType" json:"godType,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SIGodSiegeLineUpSuccess) Reset()                    { *m = SIGodSiegeLineUpSuccess{} }
func (m *SIGodSiegeLineUpSuccess) String() string            { return proto.CompactTextString(m) }
func (*SIGodSiegeLineUpSuccess) ProtoMessage()               {}
func (*SIGodSiegeLineUpSuccess) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{3} }

func (m *SIGodSiegeLineUpSuccess) GetGodType() int32 {
	if m != nil && m.GodType != nil {
		return *m.GodType
	}
	return 0
}

type SIGodSiegeCancleLineUp struct {
	GodType          *int32 `protobuf:"varint,1,req,name=godType" json:"godType,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SIGodSiegeCancleLineUp) Reset()                    { *m = SIGodSiegeCancleLineUp{} }
func (m *SIGodSiegeCancleLineUp) String() string            { return proto.CompactTextString(m) }
func (*SIGodSiegeCancleLineUp) ProtoMessage()               {}
func (*SIGodSiegeCancleLineUp) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{4} }

func (m *SIGodSiegeCancleLineUp) GetGodType() int32 {
	if m != nil && m.GodType != nil {
		return *m.GodType
	}
	return 0
}

type ISGodSiegeCancleLineUp struct {
	GodType          *int32 `protobuf:"varint,1,req,name=godType" json:"godType,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ISGodSiegeCancleLineUp) Reset()                    { *m = ISGodSiegeCancleLineUp{} }
func (m *ISGodSiegeCancleLineUp) String() string            { return proto.CompactTextString(m) }
func (*ISGodSiegeCancleLineUp) ProtoMessage()               {}
func (*ISGodSiegeCancleLineUp) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{5} }

func (m *ISGodSiegeCancleLineUp) GetGodType() int32 {
	if m != nil && m.GodType != nil {
		return *m.GodType
	}
	return 0
}

type ISGodSiegeFinishLineUpCancle struct {
	GodType          *int32 `protobuf:"varint,1,req,name=godType" json:"godType,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ISGodSiegeFinishLineUpCancle) Reset()                    { *m = ISGodSiegeFinishLineUpCancle{} }
func (m *ISGodSiegeFinishLineUpCancle) String() string            { return proto.CompactTextString(m) }
func (*ISGodSiegeFinishLineUpCancle) ProtoMessage()               {}
func (*ISGodSiegeFinishLineUpCancle) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{6} }

func (m *ISGodSiegeFinishLineUpCancle) GetGodType() int32 {
	if m != nil && m.GodType != nil {
		return *m.GodType
	}
	return 0
}

type SIGodSiegeFinishLineUpCancle struct {
	GodType          *int32 `protobuf:"varint,1,req,name=godType" json:"godType,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SIGodSiegeFinishLineUpCancle) Reset()                    { *m = SIGodSiegeFinishLineUpCancle{} }
func (m *SIGodSiegeFinishLineUpCancle) String() string            { return proto.CompactTextString(m) }
func (*SIGodSiegeFinishLineUpCancle) ProtoMessage()               {}
func (*SIGodSiegeFinishLineUpCancle) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{7} }

func (m *SIGodSiegeFinishLineUpCancle) GetGodType() int32 {
	if m != nil && m.GodType != nil {
		return *m.GodType
	}
	return 0
}

func init() {
	proto.RegisterType((*SIGodSiegeAttend)(nil), "cross.SIGodSiegeAttend")
	proto.RegisterType((*ISGodSiegeAttend)(nil), "cross.ISGodSiegeAttend")
	proto.RegisterType((*ISGodSiegeLineUpSuccess)(nil), "cross.ISGodSiegeLineUpSuccess")
	proto.RegisterType((*SIGodSiegeLineUpSuccess)(nil), "cross.SIGodSiegeLineUpSuccess")
	proto.RegisterType((*SIGodSiegeCancleLineUp)(nil), "cross.SIGodSiegeCancleLineUp")
	proto.RegisterType((*ISGodSiegeCancleLineUp)(nil), "cross.ISGodSiegeCancleLineUp")
	proto.RegisterType((*ISGodSiegeFinishLineUpCancle)(nil), "cross.ISGodSiegeFinishLineUpCancle")
	proto.RegisterType((*SIGodSiegeFinishLineUpCancle)(nil), "cross.SIGodSiegeFinishLineUpCancle")
}

func init() { proto.RegisterFile("godsiege.proto", fileDescriptor11) }

var fileDescriptor11 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xcf, 0x4f, 0x29,
	0xce, 0x4c, 0x4d, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0x2e, 0xca, 0x2f,
	0x2e, 0x56, 0x52, 0xe6, 0x12, 0x08, 0xf6, 0x74, 0xcf, 0x4f, 0x09, 0x06, 0x49, 0x39, 0x96, 0x94,
	0xa4, 0xe6, 0xa5, 0x08, 0xf1, 0x73, 0xb1, 0xa7, 0xe7, 0xa7, 0x84, 0x54, 0x16, 0xa4, 0x4a, 0x30,
	0x2a, 0x30, 0x69, 0xb0, 0x2a, 0x79, 0x70, 0x09, 0x78, 0x06, 0x13, 0x50, 0x24, 0x24, 0xc0, 0xc5,
	0x91, 0x59, 0xec, 0x93, 0x99, 0x97, 0x1a, 0x5a, 0x20, 0xc1, 0xa4, 0xc0, 0xa4, 0xc1, 0x21, 0x24,
	0xc8, 0xc5, 0x99, 0x94, 0x9a, 0x96, 0x5f, 0x94, 0xea, 0x57, 0x9a, 0x2b, 0xc1, 0x0c, 0x36, 0x49,
	0x8b, 0x4b, 0x1c, 0x61, 0x12, 0x44, 0x71, 0x70, 0x69, 0x72, 0x72, 0x6a, 0x71, 0x31, 0xa6, 0xad,
	0x5a, 0x5c, 0xe2, 0x08, 0xa7, 0x11, 0x50, 0xab, 0xc9, 0x25, 0x86, 0x50, 0xeb, 0x9c, 0x98, 0x97,
	0x9c, 0x03, 0xd5, 0x81, 0x55, 0x29, 0xc2, 0x09, 0xf8, 0x95, 0xea, 0x73, 0xc9, 0x20, 0x94, 0xba,
	0x65, 0xe6, 0x65, 0x16, 0x67, 0x40, 0x94, 0x42, 0xb4, 0x61, 0xd5, 0x80, 0x70, 0x06, 0x11, 0x1a,
	0x00, 0x01, 0x00, 0x00, 0xff, 0xff, 0xda, 0x73, 0xf5, 0x42, 0x96, 0x01, 0x00, 0x00,
}
