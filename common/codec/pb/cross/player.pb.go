// Code generated by protoc-gen-go. DO NOT EDIT.
// source: player.proto

package cross

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PlayerBasicData struct {
	UserId           *int64  `protobuf:"varint,1,req,name=userId" json:"userId,omitempty"`
	PlayerId         *int64  `protobuf:"varint,2,req,name=playerId" json:"playerId,omitempty"`
	Name             *string `protobuf:"bytes,3,req,name=name" json:"name,omitempty"`
	Role             *int32  `protobuf:"varint,4,req,name=role" json:"role,omitempty"`
	Sex              *int32  `protobuf:"varint,5,req,name=sex" json:"sex,omitempty"`
	ServerId         *int32  `protobuf:"varint,6,req,name=serverId" json:"serverId,omitempty"`
	GuaJi            *bool   `protobuf:"varint,7,req,name=guaJi" json:"guaJi,omitempty"`
	Platform         *int32  `protobuf:"varint,8,req,name=platform" json:"platform,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PlayerBasicData) Reset()                    { *m = PlayerBasicData{} }
func (m *PlayerBasicData) String() string            { return proto.CompactTextString(m) }
func (*PlayerBasicData) ProtoMessage()               {}
func (*PlayerBasicData) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{0} }

func (m *PlayerBasicData) GetUserId() int64 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *PlayerBasicData) GetPlayerId() int64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

func (m *PlayerBasicData) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *PlayerBasicData) GetRole() int32 {
	if m != nil && m.Role != nil {
		return *m.Role
	}
	return 0
}

func (m *PlayerBasicData) GetSex() int32 {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return 0
}

func (m *PlayerBasicData) GetServerId() int32 {
	if m != nil && m.ServerId != nil {
		return *m.ServerId
	}
	return 0
}

func (m *PlayerBasicData) GetGuaJi() bool {
	if m != nil && m.GuaJi != nil {
		return *m.GuaJi
	}
	return false
}

func (m *PlayerBasicData) GetPlatform() int32 {
	if m != nil && m.Platform != nil {
		return *m.Platform
	}
	return 0
}

func init() {
	proto.RegisterType((*PlayerBasicData)(nil), "cross.PlayerBasicData")
}

func init() { proto.RegisterFile("player.proto", fileDescriptor21) }

var fileDescriptor21 = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x2c, 0x8c, 0x41, 0x0e, 0x82, 0x30,
	0x10, 0x45, 0x43, 0xa1, 0x58, 0x47, 0x54, 0xd2, 0xd5, 0x2c, 0x1b, 0x57, 0x5d, 0x79, 0x08, 0xe3,
	0x06, 0x57, 0x5e, 0xa1, 0x81, 0x6a, 0x48, 0xc0, 0x92, 0x99, 0x62, 0xf4, 0x14, 0x5e, 0xd9, 0x50,
	0x5d, 0xbe, 0x97, 0xff, 0x3e, 0x54, 0xd3, 0xe0, 0xde, 0x9e, 0x8e, 0x13, 0x85, 0x18, 0xb4, 0x6c,
	0x29, 0x30, 0x1f, 0x3e, 0x19, 0xec, 0xaf, 0xc9, 0x9f, 0x1c, 0xf7, 0xed, 0xd9, 0x45, 0xa7, 0x77,
	0x50, 0xce, 0xec, 0xa9, 0xe9, 0x30, 0x33, 0xc2, 0xe6, 0xba, 0x06, 0xf5, 0x4b, 0x9b, 0x0e, 0x45,
	0x32, 0x15, 0x14, 0x0f, 0x37, 0x7a, 0xcc, 0x8d, 0xb0, 0xeb, 0x85, 0x28, 0x0c, 0x1e, 0x0b, 0x23,
	0xac, 0xd4, 0x1b, 0xc8, 0xd9, 0xbf, 0x50, 0x26, 0xa8, 0x41, 0xb1, 0xa7, 0x67, 0x4a, 0xcb, 0x64,
	0xb6, 0x20, 0xef, 0xb3, 0xbb, 0xf4, 0xb8, 0x32, 0xc2, 0xaa, 0xff, 0x77, 0xbc, 0x05, 0x1a, 0x51,
	0x2d, 0x83, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x35, 0x30, 0x33, 0xa7, 0x00, 0x00, 0x00,
}
