// Code generated by protoc-gen-go. DO NOT EDIT.
// source: login.proto

package cross

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SILogin struct {
	PlayerId         *int64 `protobuf:"varint,1,req,name=playerId" json:"playerId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SILogin) Reset()                    { *m = SILogin{} }
func (m *SILogin) String() string            { return proto.CompactTextString(m) }
func (*SILogin) ProtoMessage()               {}
func (*SILogin) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{0} }

func (m *SILogin) GetPlayerId() int64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

type ISLogin struct {
	PlayerId         *int64 `protobuf:"varint,1,req,name=playerId" json:"playerId,omitempty"`
	Match            *bool  `protobuf:"varint,2,req,name=match" json:"match,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ISLogin) Reset()                    { *m = ISLogin{} }
func (m *ISLogin) String() string            { return proto.CompactTextString(m) }
func (*ISLogin) ProtoMessage()               {}
func (*ISLogin) Descriptor() ([]byte, []int) { return fileDescriptor17, []int{1} }

func (m *ISLogin) GetPlayerId() int64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

func (m *ISLogin) GetMatch() bool {
	if m != nil && m.Match != nil {
		return *m.Match
	}
	return false
}

func init() {
	proto.RegisterType((*SILogin)(nil), "cross.SILogin")
	proto.RegisterType((*ISLogin)(nil), "cross.ISLogin")
}

func init() { proto.RegisterFile("login.proto", fileDescriptor17) }

var fileDescriptor17 = []byte{
	// 94 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0xc9, 0x4f, 0xcf,
	0xcc, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0x2e, 0xca, 0x2f, 0x2e, 0x56, 0x92,
	0xe6, 0x62, 0x0f, 0xf6, 0xf4, 0x01, 0x89, 0x0b, 0x09, 0x70, 0x71, 0x14, 0xe4, 0x24, 0x56, 0xa6,
	0x16, 0x79, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x69, 0x30, 0x2b, 0x69, 0x71, 0xb1, 0x7b, 0x06, 0xe3,
	0x90, 0x14, 0xe2, 0xe5, 0x62, 0xcd, 0x4d, 0x2c, 0x49, 0xce, 0x90, 0x60, 0x52, 0x60, 0xd2, 0xe0,
	0x00, 0x04, 0x00, 0x00, 0xff, 0xff, 0x90, 0x35, 0x74, 0x5a, 0x5d, 0x00, 0x00, 0x00,
}
