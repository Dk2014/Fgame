// Code generated by protoc-gen-go. DO NOT EDIT.
// source: myboss.proto

package ui

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CSMyBossChallenge struct {
	BiologyId        *int32 `protobuf:"varint,1,req,name=biologyId" json:"biologyId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSMyBossChallenge) Reset()                    { *m = CSMyBossChallenge{} }
func (m *CSMyBossChallenge) String() string            { return proto.CompactTextString(m) }
func (*CSMyBossChallenge) ProtoMessage()               {}
func (*CSMyBossChallenge) Descriptor() ([]byte, []int) { return fileDescriptor76, []int{0} }

func (m *CSMyBossChallenge) GetBiologyId() int32 {
	if m != nil && m.BiologyId != nil {
		return *m.BiologyId
	}
	return 0
}

type SCMyBossChallenge struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCMyBossChallenge) Reset()                    { *m = SCMyBossChallenge{} }
func (m *SCMyBossChallenge) String() string            { return proto.CompactTextString(m) }
func (*SCMyBossChallenge) ProtoMessage()               {}
func (*SCMyBossChallenge) Descriptor() ([]byte, []int) { return fileDescriptor76, []int{1} }

type SCMyBossChallengeResult struct {
	IsSuccess        *bool       `protobuf:"varint,1,req,name=isSuccess" json:"isSuccess,omitempty"`
	DropList         []*DropInfo `protobuf:"bytes,2,rep,name=dropList" json:"dropList,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *SCMyBossChallengeResult) Reset()                    { *m = SCMyBossChallengeResult{} }
func (m *SCMyBossChallengeResult) String() string            { return proto.CompactTextString(m) }
func (*SCMyBossChallengeResult) ProtoMessage()               {}
func (*SCMyBossChallengeResult) Descriptor() ([]byte, []int) { return fileDescriptor76, []int{2} }

func (m *SCMyBossChallengeResult) GetIsSuccess() bool {
	if m != nil && m.IsSuccess != nil {
		return *m.IsSuccess
	}
	return false
}

func (m *SCMyBossChallengeResult) GetDropList() []*DropInfo {
	if m != nil {
		return m.DropList
	}
	return nil
}

type SCMyBossSceneInfo struct {
	CreateTime       *int64 `protobuf:"varint,1,req,name=createTime" json:"createTime,omitempty"`
	BossId           *int32 `protobuf:"varint,2,req,name=bossId" json:"bossId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCMyBossSceneInfo) Reset()                    { *m = SCMyBossSceneInfo{} }
func (m *SCMyBossSceneInfo) String() string            { return proto.CompactTextString(m) }
func (*SCMyBossSceneInfo) ProtoMessage()               {}
func (*SCMyBossSceneInfo) Descriptor() ([]byte, []int) { return fileDescriptor76, []int{3} }

func (m *SCMyBossSceneInfo) GetCreateTime() int64 {
	if m != nil && m.CreateTime != nil {
		return *m.CreateTime
	}
	return 0
}

func (m *SCMyBossSceneInfo) GetBossId() int32 {
	if m != nil && m.BossId != nil {
		return *m.BossId
	}
	return 0
}

type MyBossInfo struct {
	BiologyId        *int32 `protobuf:"varint,1,req,name=biologyId" json:"biologyId,omitempty"`
	AttendTimes      *int32 `protobuf:"varint,2,req,name=attendTimes" json:"attendTimes,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MyBossInfo) Reset()                    { *m = MyBossInfo{} }
func (m *MyBossInfo) String() string            { return proto.CompactTextString(m) }
func (*MyBossInfo) ProtoMessage()               {}
func (*MyBossInfo) Descriptor() ([]byte, []int) { return fileDescriptor76, []int{4} }

func (m *MyBossInfo) GetBiologyId() int32 {
	if m != nil && m.BiologyId != nil {
		return *m.BiologyId
	}
	return 0
}

func (m *MyBossInfo) GetAttendTimes() int32 {
	if m != nil && m.AttendTimes != nil {
		return *m.AttendTimes
	}
	return 0
}

type SCMyBossInfoNotice struct {
	InfoList         []*MyBossInfo `protobuf:"bytes,1,rep,name=infoList" json:"infoList,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *SCMyBossInfoNotice) Reset()                    { *m = SCMyBossInfoNotice{} }
func (m *SCMyBossInfoNotice) String() string            { return proto.CompactTextString(m) }
func (*SCMyBossInfoNotice) ProtoMessage()               {}
func (*SCMyBossInfoNotice) Descriptor() ([]byte, []int) { return fileDescriptor76, []int{5} }

func (m *SCMyBossInfoNotice) GetInfoList() []*MyBossInfo {
	if m != nil {
		return m.InfoList
	}
	return nil
}

func init() {
	proto.RegisterType((*CSMyBossChallenge)(nil), "ui.CSMyBossChallenge")
	proto.RegisterType((*SCMyBossChallenge)(nil), "ui.SCMyBossChallenge")
	proto.RegisterType((*SCMyBossChallengeResult)(nil), "ui.SCMyBossChallengeResult")
	proto.RegisterType((*SCMyBossSceneInfo)(nil), "ui.SCMyBossSceneInfo")
	proto.RegisterType((*MyBossInfo)(nil), "ui.MyBossInfo")
	proto.RegisterType((*SCMyBossInfoNotice)(nil), "ui.SCMyBossInfoNotice")
}

func init() { proto.RegisterFile("myboss.proto", fileDescriptor76) }

var fileDescriptor76 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x69, 0x8a, 0x12, 0xa7, 0xa5, 0x90, 0xe9, 0xc1, 0xe0, 0x41, 0x42, 0x0e, 0xd2, 0x53,
	0x0e, 0x22, 0x7a, 0x37, 0x5e, 0x02, 0xd5, 0x83, 0xf1, 0x0f, 0xa4, 0x9b, 0x69, 0x1d, 0x48, 0x77,
	0x42, 0x66, 0x73, 0xc8, 0xbf, 0x97, 0xdd, 0x2a, 0x05, 0x7b, 0xdc, 0xe1, 0xbd, 0xef, 0x7d, 0x0b,
	0xcb, 0xe3, 0xb4, 0x13, 0xd5, 0xa2, 0x1f, 0xc4, 0x09, 0x46, 0x23, 0xdf, 0x41, 0x3b, 0x48, 0x7f,
	0x7a, 0xe7, 0x0f, 0x90, 0x94, 0xf5, 0xfb, 0xf4, 0x2a, 0xaa, 0xe5, 0x77, 0xd3, 0x75, 0x64, 0x0f,
	0x84, 0x09, 0xdc, 0xec, 0x58, 0x3a, 0x39, 0x4c, 0x55, 0x9b, 0xce, 0xb2, 0x68, 0x73, 0x95, 0xaf,
	0x21, 0xa9, 0xcb, 0x7f, 0xb9, 0x7c, 0x0b, 0xb7, 0x17, 0xc7, 0x4f, 0xd2, 0xb1, 0x73, 0x1e, 0xc1,
	0x5a, 0x8f, 0xc6, 0x90, 0x6a, 0x40, 0xc4, 0x78, 0x0f, 0xb1, 0x1f, 0xde, 0xb2, 0xba, 0x34, 0xca,
	0xe6, 0x9b, 0xc5, 0xe3, 0xb2, 0x18, 0xb9, 0x78, 0x1b, 0xa4, 0xaf, 0xec, 0x5e, 0xf2, 0x97, 0xf3,
	0x44, 0x6d, 0xc8, 0x92, 0x3f, 0x22, 0x02, 0x98, 0x81, 0x1a, 0x47, 0x5f, 0x7c, 0xa4, 0x00, 0x9a,
	0xe3, 0x0a, 0xae, 0xfd, 0x8f, 0xaa, 0x36, 0x8d, 0x82, 0xdb, 0x13, 0xc0, 0xa9, 0x16, 0x1a, 0x97,
	0xf2, 0xb8, 0x86, 0x45, 0xe3, 0x1c, 0xd9, 0xd6, 0x43, 0xf4, 0xb7, 0xf5, 0x0c, 0xf8, 0x37, 0xe7,
	0x7b, 0x1f, 0xe2, 0xd8, 0x10, 0x66, 0x10, 0xb3, 0xdd, 0x4b, 0x90, 0x9c, 0x05, 0xc9, 0x95, 0x97,
	0x3c, 0xe7, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x3e, 0x2e, 0xf3, 0x50, 0x01, 0x00, 0x00,
}
