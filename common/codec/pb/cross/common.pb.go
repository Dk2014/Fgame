// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package cross

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SIHeartBeat struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *SIHeartBeat) Reset()                    { *m = SIHeartBeat{} }
func (m *SIHeartBeat) String() string            { return proto.CompactTextString(m) }
func (*SIHeartBeat) ProtoMessage()               {}
func (*SIHeartBeat) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

type ISHeartBeat struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ISHeartBeat) Reset()                    { *m = ISHeartBeat{} }
func (m *ISHeartBeat) String() string            { return proto.CompactTextString(m) }
func (*ISHeartBeat) ProtoMessage()               {}
func (*ISHeartBeat) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

func init() {
	proto.RegisterType((*SIHeartBeat)(nil), "cross.SIHeartBeat")
	proto.RegisterType((*ISHeartBeat)(nil), "cross.ISHeartBeat")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 66 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0x2e, 0xca, 0x2f, 0x2e, 0x56,
	0xe2, 0xe5, 0xe2, 0x0e, 0xf6, 0xf4, 0x48, 0x4d, 0x2c, 0x2a, 0x71, 0x4a, 0x4d, 0x2c, 0x01, 0x71,
	0x3d, 0x83, 0xe1, 0x5c, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x39, 0xbf, 0x3c, 0xc7, 0x33, 0x00,
	0x00, 0x00,
}
