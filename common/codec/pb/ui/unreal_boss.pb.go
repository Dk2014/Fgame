// Code generated by protoc-gen-go. DO NOT EDIT.
// source: unreal_boss.proto

package ui

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CSUnrealBossList struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSUnrealBossList) Reset()                    { *m = CSUnrealBossList{} }
func (m *CSUnrealBossList) String() string            { return proto.CompactTextString(m) }
func (*CSUnrealBossList) ProtoMessage()               {}
func (*CSUnrealBossList) Descriptor() ([]byte, []int) { return fileDescriptor123, []int{0} }

type SCUnrealBossList struct {
	BossInfoList     []*BossInfo `protobuf:"bytes,1,rep,name=bossInfoList" json:"bossInfoList,omitempty"`
	CurPilao         *int32      `protobuf:"varint,2,req,name=curPilao" json:"curPilao,omitempty"`
	CurBuyTimes      *int32      `protobuf:"varint,3,req,name=curBuyTimes" json:"curBuyTimes,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *SCUnrealBossList) Reset()                    { *m = SCUnrealBossList{} }
func (m *SCUnrealBossList) String() string            { return proto.CompactTextString(m) }
func (*SCUnrealBossList) ProtoMessage()               {}
func (*SCUnrealBossList) Descriptor() ([]byte, []int) { return fileDescriptor123, []int{1} }

func (m *SCUnrealBossList) GetBossInfoList() []*BossInfo {
	if m != nil {
		return m.BossInfoList
	}
	return nil
}

func (m *SCUnrealBossList) GetCurPilao() int32 {
	if m != nil && m.CurPilao != nil {
		return *m.CurPilao
	}
	return 0
}

func (m *SCUnrealBossList) GetCurBuyTimes() int32 {
	if m != nil && m.CurBuyTimes != nil {
		return *m.CurBuyTimes
	}
	return 0
}

type CSUnrealBossChallenge struct {
	BiologyId        *int32 `protobuf:"varint,1,req,name=biologyId" json:"biologyId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSUnrealBossChallenge) Reset()                    { *m = CSUnrealBossChallenge{} }
func (m *CSUnrealBossChallenge) String() string            { return proto.CompactTextString(m) }
func (*CSUnrealBossChallenge) ProtoMessage()               {}
func (*CSUnrealBossChallenge) Descriptor() ([]byte, []int) { return fileDescriptor123, []int{2} }

func (m *CSUnrealBossChallenge) GetBiologyId() int32 {
	if m != nil && m.BiologyId != nil {
		return *m.BiologyId
	}
	return 0
}

type SCUnrealBossChallenge struct {
	Pos              *Position `protobuf:"bytes,1,req,name=pos" json:"pos,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *SCUnrealBossChallenge) Reset()                    { *m = SCUnrealBossChallenge{} }
func (m *SCUnrealBossChallenge) String() string            { return proto.CompactTextString(m) }
func (*SCUnrealBossChallenge) ProtoMessage()               {}
func (*SCUnrealBossChallenge) Descriptor() ([]byte, []int) { return fileDescriptor123, []int{3} }

func (m *SCUnrealBossChallenge) GetPos() *Position {
	if m != nil {
		return m.Pos
	}
	return nil
}

type SCUnrealBossInfoBroadcast struct {
	BossInfo         *BossInfo `protobuf:"bytes,1,req,name=bossInfo" json:"bossInfo,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *SCUnrealBossInfoBroadcast) Reset()                    { *m = SCUnrealBossInfoBroadcast{} }
func (m *SCUnrealBossInfoBroadcast) String() string            { return proto.CompactTextString(m) }
func (*SCUnrealBossInfoBroadcast) ProtoMessage()               {}
func (*SCUnrealBossInfoBroadcast) Descriptor() ([]byte, []int) { return fileDescriptor123, []int{4} }

func (m *SCUnrealBossInfoBroadcast) GetBossInfo() *BossInfo {
	if m != nil {
		return m.BossInfo
	}
	return nil
}

type SCUnrealBossListInfoNotice struct {
	BossInfoList     []*BossInfo `protobuf:"bytes,1,rep,name=bossInfoList" json:"bossInfoList,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *SCUnrealBossListInfoNotice) Reset()                    { *m = SCUnrealBossListInfoNotice{} }
func (m *SCUnrealBossListInfoNotice) String() string            { return proto.CompactTextString(m) }
func (*SCUnrealBossListInfoNotice) ProtoMessage()               {}
func (*SCUnrealBossListInfoNotice) Descriptor() ([]byte, []int) { return fileDescriptor123, []int{5} }

func (m *SCUnrealBossListInfoNotice) GetBossInfoList() []*BossInfo {
	if m != nil {
		return m.BossInfoList
	}
	return nil
}

type CSUnrealBossBuyPilaoNum struct {
	BuyNum           *int32 `protobuf:"varint,1,req,name=buyNum" json:"buyNum,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSUnrealBossBuyPilaoNum) Reset()                    { *m = CSUnrealBossBuyPilaoNum{} }
func (m *CSUnrealBossBuyPilaoNum) String() string            { return proto.CompactTextString(m) }
func (*CSUnrealBossBuyPilaoNum) ProtoMessage()               {}
func (*CSUnrealBossBuyPilaoNum) Descriptor() ([]byte, []int) { return fileDescriptor123, []int{6} }

func (m *CSUnrealBossBuyPilaoNum) GetBuyNum() int32 {
	if m != nil && m.BuyNum != nil {
		return *m.BuyNum
	}
	return 0
}

type SCUnrealBossBuyPilaoNum struct {
	CurPilao         *int32 `protobuf:"varint,1,req,name=curPilao" json:"curPilao,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCUnrealBossBuyPilaoNum) Reset()                    { *m = SCUnrealBossBuyPilaoNum{} }
func (m *SCUnrealBossBuyPilaoNum) String() string            { return proto.CompactTextString(m) }
func (*SCUnrealBossBuyPilaoNum) ProtoMessage()               {}
func (*SCUnrealBossBuyPilaoNum) Descriptor() ([]byte, []int) { return fileDescriptor123, []int{7} }

func (m *SCUnrealBossBuyPilaoNum) GetCurPilao() int32 {
	if m != nil && m.CurPilao != nil {
		return *m.CurPilao
	}
	return 0
}

type SCUnrealBossEnemiesNotice struct {
	AttackId         *int64      `protobuf:"varint,1,req,name=attackId" json:"attackId,omitempty"`
	AttackPlayerName *string     `protobuf:"bytes,2,req,name=attackPlayerName" json:"attackPlayerName,omitempty"`
	BossName         *string     `protobuf:"bytes,3,req,name=bossName" json:"bossName,omitempty"`
	DropInfoList     []*DropInfo `protobuf:"bytes,4,rep,name=dropInfoList" json:"dropInfoList,omitempty"`
	BiologyId        *int32      `protobuf:"varint,5,req,name=biologyId" json:"biologyId,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *SCUnrealBossEnemiesNotice) Reset()                    { *m = SCUnrealBossEnemiesNotice{} }
func (m *SCUnrealBossEnemiesNotice) String() string            { return proto.CompactTextString(m) }
func (*SCUnrealBossEnemiesNotice) ProtoMessage()               {}
func (*SCUnrealBossEnemiesNotice) Descriptor() ([]byte, []int) { return fileDescriptor123, []int{8} }

func (m *SCUnrealBossEnemiesNotice) GetAttackId() int64 {
	if m != nil && m.AttackId != nil {
		return *m.AttackId
	}
	return 0
}

func (m *SCUnrealBossEnemiesNotice) GetAttackPlayerName() string {
	if m != nil && m.AttackPlayerName != nil {
		return *m.AttackPlayerName
	}
	return ""
}

func (m *SCUnrealBossEnemiesNotice) GetBossName() string {
	if m != nil && m.BossName != nil {
		return *m.BossName
	}
	return ""
}

func (m *SCUnrealBossEnemiesNotice) GetDropInfoList() []*DropInfo {
	if m != nil {
		return m.DropInfoList
	}
	return nil
}

func (m *SCUnrealBossEnemiesNotice) GetBiologyId() int32 {
	if m != nil && m.BiologyId != nil {
		return *m.BiologyId
	}
	return 0
}

type SCUnrealBossPilaoInfo struct {
	CurPilao         *int32 `protobuf:"varint,1,req,name=curPilao" json:"curPilao,omitempty"`
	CurBuyTimes      *int32 `protobuf:"varint,2,req,name=curBuyTimes" json:"curBuyTimes,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCUnrealBossPilaoInfo) Reset()                    { *m = SCUnrealBossPilaoInfo{} }
func (m *SCUnrealBossPilaoInfo) String() string            { return proto.CompactTextString(m) }
func (*SCUnrealBossPilaoInfo) ProtoMessage()               {}
func (*SCUnrealBossPilaoInfo) Descriptor() ([]byte, []int) { return fileDescriptor123, []int{9} }

func (m *SCUnrealBossPilaoInfo) GetCurPilao() int32 {
	if m != nil && m.CurPilao != nil {
		return *m.CurPilao
	}
	return 0
}

func (m *SCUnrealBossPilaoInfo) GetCurBuyTimes() int32 {
	if m != nil && m.CurBuyTimes != nil {
		return *m.CurBuyTimes
	}
	return 0
}

func init() {
	proto.RegisterType((*CSUnrealBossList)(nil), "ui.CSUnrealBossList")
	proto.RegisterType((*SCUnrealBossList)(nil), "ui.SCUnrealBossList")
	proto.RegisterType((*CSUnrealBossChallenge)(nil), "ui.CSUnrealBossChallenge")
	proto.RegisterType((*SCUnrealBossChallenge)(nil), "ui.SCUnrealBossChallenge")
	proto.RegisterType((*SCUnrealBossInfoBroadcast)(nil), "ui.SCUnrealBossInfoBroadcast")
	proto.RegisterType((*SCUnrealBossListInfoNotice)(nil), "ui.SCUnrealBossListInfoNotice")
	proto.RegisterType((*CSUnrealBossBuyPilaoNum)(nil), "ui.CSUnrealBossBuyPilaoNum")
	proto.RegisterType((*SCUnrealBossBuyPilaoNum)(nil), "ui.SCUnrealBossBuyPilaoNum")
	proto.RegisterType((*SCUnrealBossEnemiesNotice)(nil), "ui.SCUnrealBossEnemiesNotice")
	proto.RegisterType((*SCUnrealBossPilaoInfo)(nil), "ui.SCUnrealBossPilaoInfo")
}

func init() { proto.RegisterFile("unreal_boss.proto", fileDescriptor123) }

var fileDescriptor123 = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x8b, 0xdb, 0x30,
	0x10, 0x85, 0x71, 0xdc, 0x94, 0x64, 0x62, 0x8a, 0xa3, 0x12, 0xe2, 0xf8, 0x50, 0x82, 0x4e, 0x69,
	0x0b, 0x39, 0xe4, 0x5a, 0x28, 0xc5, 0x69, 0x0f, 0x81, 0x62, 0x02, 0x69, 0x8f, 0xa5, 0x28, 0xb6,
	0x36, 0x2b, 0xd6, 0xf6, 0x04, 0xc9, 0x62, 0xf1, 0xbf, 0xd8, 0x9f, 0xbc, 0x48, 0x76, 0x36, 0xf2,
	0xee, 0x1e, 0xf6, 0x38, 0x8f, 0x27, 0xcd, 0xfb, 0x66, 0x06, 0xa6, 0xba, 0x92, 0x9c, 0x15, 0xff,
	0x8f, 0xa8, 0xd4, 0xfa, 0x2c, 0xb1, 0x46, 0x32, 0xd0, 0x22, 0x86, 0x5c, 0xe2, 0xb9, 0xad, 0xe3,
	0xf0, 0x1e, 0x65, 0x91, 0x3b, 0x8e, 0x38, 0xc8, 0xb0, 0x2c, 0xb1, 0x6a, 0x2b, 0x4a, 0x20, 0xdc,
	0x1e, 0xfe, 0xda, 0x6f, 0x12, 0x54, 0xea, 0xb7, 0x50, 0x35, 0xfd, 0x07, 0xe1, 0x61, 0xdb, 0xd7,
	0x08, 0x85, 0xc0, 0xfc, 0xb1, 0xab, 0x6e, 0xd0, 0xd4, 0x91, 0xb7, 0xf4, 0x57, 0x93, 0x4d, 0xb0,
	0xd6, 0x62, 0x9d, 0x74, 0x3a, 0x09, 0x61, 0x94, 0x69, 0xb9, 0x17, 0x05, 0xc3, 0x68, 0xb0, 0x1c,
	0xac, 0x86, 0xe4, 0x23, 0x4c, 0x32, 0x2d, 0x13, 0xdd, 0xfc, 0x11, 0x25, 0x57, 0x91, 0x6f, 0x44,
	0xfa, 0x05, 0x66, 0x6e, 0xcb, 0xed, 0x2d, 0x2b, 0x0a, 0x5e, 0x9d, 0x38, 0x99, 0xc2, 0xf8, 0x28,
	0xb0, 0xc0, 0x53, 0xb3, 0xcb, 0x23, 0xcf, 0x7a, 0x37, 0x30, 0x73, 0xa3, 0x5c, 0xbd, 0x0b, 0xf0,
	0xcf, 0xa8, 0xac, 0xab, 0x8b, 0xb1, 0x47, 0x25, 0x6a, 0x81, 0x15, 0xfd, 0x06, 0x0b, 0xf7, 0x8d,
	0x89, 0x96, 0x48, 0x64, 0x79, 0xc6, 0x54, 0x4d, 0x3e, 0xc1, 0xe8, 0xc2, 0xe1, 0x3e, 0xbe, 0x18,
	0xe9, 0x0f, 0x88, 0x9f, 0xb3, 0x1b, 0x3d, 0xc5, 0x5a, 0x64, 0xfc, 0x2d, 0x53, 0xa0, 0x9f, 0x61,
	0xee, 0xe2, 0x25, 0xba, 0xb1, 0x13, 0x49, 0x75, 0x49, 0x3e, 0xc0, 0xfb, 0xa3, 0x6e, 0x52, 0x5d,
	0x76, 0x74, 0x5f, 0x61, 0xee, 0x36, 0x73, 0xad, 0xee, 0x2c, 0x5b, 0xf3, 0x83, 0xd7, 0xe7, 0xfa,
	0x55, 0xf1, 0x52, 0x70, 0xd5, 0x25, 0x0b, 0x61, 0xc4, 0xea, 0x9a, 0x65, 0x77, 0xdd, 0xe8, 0x7c,
	0x12, 0x41, 0xd8, 0x2a, 0xfb, 0x82, 0x35, 0x5c, 0xa6, 0xac, 0xe4, 0x76, 0x2b, 0x63, 0xe3, 0x35,
	0x14, 0x56, 0xf1, 0xad, 0x42, 0x21, 0x30, 0x37, 0xf3, 0xc4, 0xf5, 0xee, 0xca, 0xf5, 0xb3, 0xd3,
	0xfb, 0xdb, 0x19, 0xda, 0x48, 0xdf, 0xfb, 0xdb, 0xb1, 0x69, 0x5f, 0x5c, 0x82, 0xf7, 0xda, 0x25,
	0xd8, 0xf3, 0x78, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x8b, 0xc1, 0xe7, 0xbe, 0xc0, 0x02, 0x00, 0x00,
}
