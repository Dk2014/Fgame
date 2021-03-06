// Code generated by protoc-gen-go. DO NOT EDIT.
// source: arenapvp.proto

package cross

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SIArenapvpAttend struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *SIArenapvpAttend) Reset()                    { *m = SIArenapvpAttend{} }
func (m *SIArenapvpAttend) String() string            { return proto.CompactTextString(m) }
func (*SIArenapvpAttend) ProtoMessage()               {}
func (*SIArenapvpAttend) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

type ISArenapvpAttend struct {
	IsLineUp         *bool  `protobuf:"varint,1,req,name=isLineUp" json:"isLineUp,omitempty"`
	SceneId          *int64 `protobuf:"varint,2,req,name=sceneId" json:"sceneId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ISArenapvpAttend) Reset()                    { *m = ISArenapvpAttend{} }
func (m *ISArenapvpAttend) String() string            { return proto.CompactTextString(m) }
func (*ISArenapvpAttend) ProtoMessage()               {}
func (*ISArenapvpAttend) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *ISArenapvpAttend) GetIsLineUp() bool {
	if m != nil && m.IsLineUp != nil {
		return *m.IsLineUp
	}
	return false
}

func (m *ISArenapvpAttend) GetSceneId() int64 {
	if m != nil && m.SceneId != nil {
		return *m.SceneId
	}
	return 0
}

type ISArenapvpRelive struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ISArenapvpRelive) Reset()                    { *m = ISArenapvpRelive{} }
func (m *ISArenapvpRelive) String() string            { return proto.CompactTextString(m) }
func (*ISArenapvpRelive) ProtoMessage()               {}
func (*ISArenapvpRelive) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

type SIArenapvpRelive struct {
	Success          *bool  `protobuf:"varint,1,req,name=success" json:"success,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SIArenapvpRelive) Reset()                    { *m = SIArenapvpRelive{} }
func (m *SIArenapvpRelive) String() string            { return proto.CompactTextString(m) }
func (*SIArenapvpRelive) ProtoMessage()               {}
func (*SIArenapvpRelive) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *SIArenapvpRelive) GetSuccess() bool {
	if m != nil && m.Success != nil {
		return *m.Success
	}
	return false
}

type PlayerArenapvpData struct {
	ReliveTimes      *int32 `protobuf:"varint,1,opt,name=reliveTimes" json:"reliveTimes,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PlayerArenapvpData) Reset()                    { *m = PlayerArenapvpData{} }
func (m *PlayerArenapvpData) String() string            { return proto.CompactTextString(m) }
func (*PlayerArenapvpData) ProtoMessage()               {}
func (*PlayerArenapvpData) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *PlayerArenapvpData) GetReliveTimes() int32 {
	if m != nil && m.ReliveTimes != nil {
		return *m.ReliveTimes
	}
	return 0
}

type SIPlayerArenapvpDataChanged struct {
	PlayerArenapvpData *PlayerArenapvpData `protobuf:"bytes,1,req,name=playerArenapvpData" json:"playerArenapvpData,omitempty"`
	XXX_unrecognized   []byte              `json:"-"`
}

func (m *SIPlayerArenapvpDataChanged) Reset()                    { *m = SIPlayerArenapvpDataChanged{} }
func (m *SIPlayerArenapvpDataChanged) String() string            { return proto.CompactTextString(m) }
func (*SIPlayerArenapvpDataChanged) ProtoMessage()               {}
func (*SIPlayerArenapvpDataChanged) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func (m *SIPlayerArenapvpDataChanged) GetPlayerArenapvpData() *PlayerArenapvpData {
	if m != nil {
		return m.PlayerArenapvpData
	}
	return nil
}

type ISArenapvpResetReliveTimes struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ISArenapvpResetReliveTimes) Reset()                    { *m = ISArenapvpResetReliveTimes{} }
func (m *ISArenapvpResetReliveTimes) String() string            { return proto.CompactTextString(m) }
func (*ISArenapvpResetReliveTimes) ProtoMessage()               {}
func (*ISArenapvpResetReliveTimes) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

type ISArenapvpAttendSuccess struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ISArenapvpAttendSuccess) Reset()                    { *m = ISArenapvpAttendSuccess{} }
func (m *ISArenapvpAttendSuccess) String() string            { return proto.CompactTextString(m) }
func (*ISArenapvpAttendSuccess) ProtoMessage()               {}
func (*ISArenapvpAttendSuccess) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{7} }

type ISArenapvpResultBattle struct {
	Win              *bool  `protobuf:"varint,1,req,name=win" json:"win,omitempty"`
	PvpType          *int32 `protobuf:"varint,2,req,name=pvpType" json:"pvpType,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ISArenapvpResultBattle) Reset()                    { *m = ISArenapvpResultBattle{} }
func (m *ISArenapvpResultBattle) String() string            { return proto.CompactTextString(m) }
func (*ISArenapvpResultBattle) ProtoMessage()               {}
func (*ISArenapvpResultBattle) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{8} }

func (m *ISArenapvpResultBattle) GetWin() bool {
	if m != nil && m.Win != nil {
		return *m.Win
	}
	return false
}

func (m *ISArenapvpResultBattle) GetPvpType() int32 {
	if m != nil && m.PvpType != nil {
		return *m.PvpType
	}
	return 0
}

type ISArenapvpResultElection struct {
	Win              *bool  `protobuf:"varint,1,req,name=win" json:"win,omitempty"`
	Ranking          *int32 `protobuf:"varint,2,req,name=ranking" json:"ranking,omitempty"`
	PvpType          *int32 `protobuf:"varint,3,req,name=pvpType" json:"pvpType,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ISArenapvpResultElection) Reset()                    { *m = ISArenapvpResultElection{} }
func (m *ISArenapvpResultElection) String() string            { return proto.CompactTextString(m) }
func (*ISArenapvpResultElection) ProtoMessage()               {}
func (*ISArenapvpResultElection) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{9} }

func (m *ISArenapvpResultElection) GetWin() bool {
	if m != nil && m.Win != nil {
		return *m.Win
	}
	return false
}

func (m *ISArenapvpResultElection) GetRanking() int32 {
	if m != nil && m.Ranking != nil {
		return *m.Ranking
	}
	return 0
}

func (m *ISArenapvpResultElection) GetPvpType() int32 {
	if m != nil && m.PvpType != nil {
		return *m.PvpType
	}
	return 0
}

type ISAreanapvpElectionLuckyRew struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ISAreanapvpElectionLuckyRew) Reset()                    { *m = ISAreanapvpElectionLuckyRew{} }
func (m *ISAreanapvpElectionLuckyRew) String() string            { return proto.CompactTextString(m) }
func (*ISAreanapvpElectionLuckyRew) ProtoMessage()               {}
func (*ISAreanapvpElectionLuckyRew) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{10} }

func init() {
	proto.RegisterType((*SIArenapvpAttend)(nil), "cross.SIArenapvpAttend")
	proto.RegisterType((*ISArenapvpAttend)(nil), "cross.ISArenapvpAttend")
	proto.RegisterType((*ISArenapvpRelive)(nil), "cross.ISArenapvpRelive")
	proto.RegisterType((*SIArenapvpRelive)(nil), "cross.SIArenapvpRelive")
	proto.RegisterType((*PlayerArenapvpData)(nil), "cross.PlayerArenapvpData")
	proto.RegisterType((*SIPlayerArenapvpDataChanged)(nil), "cross.SIPlayerArenapvpDataChanged")
	proto.RegisterType((*ISArenapvpResetReliveTimes)(nil), "cross.ISArenapvpResetReliveTimes")
	proto.RegisterType((*ISArenapvpAttendSuccess)(nil), "cross.ISArenapvpAttendSuccess")
	proto.RegisterType((*ISArenapvpResultBattle)(nil), "cross.ISArenapvpResultBattle")
	proto.RegisterType((*ISArenapvpResultElection)(nil), "cross.ISArenapvpResultElection")
	proto.RegisterType((*ISAreanapvpElectionLuckyRew)(nil), "cross.ISAreanapvpElectionLuckyRew")
}

func init() { proto.RegisterFile("arenapvp.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xcd, 0x6a, 0xeb, 0x30,
	0x10, 0x85, 0x49, 0x82, 0xb9, 0x61, 0x0c, 0xb7, 0x46, 0x85, 0xd6, 0x69, 0x5a, 0x08, 0xea, 0xc6,
	0xdd, 0x78, 0x51, 0x48, 0xf7, 0xe9, 0xcf, 0xc2, 0x90, 0x45, 0xb1, 0xdd, 0x07, 0x10, 0xf2, 0x90,
	0x8a, 0xb8, 0xb2, 0x90, 0x64, 0x07, 0xbf, 0x7d, 0x89, 0xdc, 0x60, 0x27, 0xde, 0x6a, 0xce, 0xf7,
	0xcd, 0xe8, 0xc0, 0x7f, 0xa6, 0x51, 0x32, 0xd5, 0xa8, 0x58, 0xe9, 0xca, 0x56, 0xc4, 0xe3, 0xba,
	0x32, 0x86, 0x12, 0x08, 0xb2, 0x64, 0xf3, 0x37, 0xda, 0x58, 0x8b, 0xb2, 0xa0, 0x6b, 0x08, 0x92,
	0xec, 0xfc, 0x8d, 0x04, 0x30, 0x17, 0x66, 0x2b, 0x24, 0x7e, 0xa9, 0x70, 0xb2, 0x9a, 0x46, 0x73,
	0x72, 0x05, 0xff, 0x0c, 0x47, 0x89, 0x49, 0x11, 0x4e, 0x57, 0xd3, 0x68, 0x76, 0x54, 0xf5, 0x58,
	0x8a, 0xa5, 0x68, 0x90, 0x3e, 0x0e, 0xf5, 0xdd, 0x9b, 0x03, 0x6b, 0xce, 0xd1, 0x98, 0xce, 0x44,
	0x9f, 0x80, 0x7c, 0x96, 0xac, 0x45, 0x7d, 0x0a, 0xbe, 0x33, 0xcb, 0xc8, 0x35, 0xf8, 0xda, 0x01,
	0xb9, 0xf8, 0xc1, 0x63, 0x74, 0x12, 0x79, 0x34, 0x87, 0x65, 0x96, 0x8c, 0xc3, 0x6f, 0xdf, 0x4c,
	0xee, 0xb0, 0x20, 0x6b, 0x20, 0x6a, 0x34, 0x74, 0x5b, 0xfc, 0xe7, 0x45, 0xec, 0x7e, 0x1c, 0x8f,
	0x69, 0x7a, 0x0f, 0x77, 0xc3, 0xcb, 0x0d, 0xda, 0xb4, 0xdf, 0x4c, 0x17, 0x70, 0x7b, 0x59, 0x47,
	0xd6, 0xdd, 0x4f, 0x5f, 0xe0, 0xe6, 0x0c, 0xac, 0x4b, 0xfb, 0xca, 0xac, 0x2d, 0x91, 0xf8, 0x30,
	0x3b, 0x08, 0xd9, 0x57, 0xa5, 0x1a, 0x95, 0xb7, 0x0a, 0x5d, 0x55, 0x1e, 0x4d, 0x20, 0xbc, 0xe4,
	0x3e, 0x4a, 0xe4, 0x56, 0x54, 0x72, 0x44, 0x6a, 0x26, 0xf7, 0x42, 0xee, 0x3a, 0x72, 0xa8, 0x9a,
	0x39, 0xd5, 0x03, 0x2c, 0x9d, 0x8a, 0x39, 0xd7, 0xc9, 0xb2, 0xad, 0xf9, 0xbe, 0x4d, 0xf1, 0xf0,
	0x1b, 0x00, 0x00, 0xff, 0xff, 0x9c, 0xd4, 0x2d, 0x93, 0xf7, 0x01, 0x00, 0x00,
}
