// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat.proto

package ui

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CSChatSend struct {
	Channel          *int32  `protobuf:"varint,1,req,name=channel" json:"channel,omitempty"`
	RecvId           *int64  `protobuf:"varint,2,req,name=recvId" json:"recvId,omitempty"`
	MsgType          *int32  `protobuf:"varint,3,req,name=msgType" json:"msgType,omitempty"`
	Content          []byte  `protobuf:"bytes,4,req,name=content" json:"content,omitempty"`
	Args             *string `protobuf:"bytes,5,opt,name=args" json:"args,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CSChatSend) Reset()                    { *m = CSChatSend{} }
func (m *CSChatSend) String() string            { return proto.CompactTextString(m) }
func (*CSChatSend) ProtoMessage()               {}
func (*CSChatSend) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{0} }

func (m *CSChatSend) GetChannel() int32 {
	if m != nil && m.Channel != nil {
		return *m.Channel
	}
	return 0
}

func (m *CSChatSend) GetRecvId() int64 {
	if m != nil && m.RecvId != nil {
		return *m.RecvId
	}
	return 0
}

func (m *CSChatSend) GetMsgType() int32 {
	if m != nil && m.MsgType != nil {
		return *m.MsgType
	}
	return 0
}

func (m *CSChatSend) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *CSChatSend) GetArgs() string {
	if m != nil && m.Args != nil {
		return *m.Args
	}
	return ""
}

type SCChatSend struct {
	Channel          *int32  `protobuf:"varint,1,req,name=channel" json:"channel,omitempty"`
	RecvId           *int64  `protobuf:"varint,2,req,name=recvId" json:"recvId,omitempty"`
	MsgType          *int32  `protobuf:"varint,3,req,name=msgType" json:"msgType,omitempty"`
	Content          []byte  `protobuf:"bytes,4,req,name=content" json:"content,omitempty"`
	Args             *string `protobuf:"bytes,5,opt,name=args" json:"args,omitempty"`
	ChatCount        *int32  `protobuf:"varint,6,opt,name=chatCount" json:"chatCount,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SCChatSend) Reset()                    { *m = SCChatSend{} }
func (m *SCChatSend) String() string            { return proto.CompactTextString(m) }
func (*SCChatSend) ProtoMessage()               {}
func (*SCChatSend) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{1} }

func (m *SCChatSend) GetChannel() int32 {
	if m != nil && m.Channel != nil {
		return *m.Channel
	}
	return 0
}

func (m *SCChatSend) GetRecvId() int64 {
	if m != nil && m.RecvId != nil {
		return *m.RecvId
	}
	return 0
}

func (m *SCChatSend) GetMsgType() int32 {
	if m != nil && m.MsgType != nil {
		return *m.MsgType
	}
	return 0
}

func (m *SCChatSend) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *SCChatSend) GetArgs() string {
	if m != nil && m.Args != nil {
		return *m.Args
	}
	return ""
}

func (m *SCChatSend) GetChatCount() int32 {
	if m != nil && m.ChatCount != nil {
		return *m.ChatCount
	}
	return 0
}

type SCChatRecv struct {
	SendId           *int64  `protobuf:"varint,1,req,name=sendId" json:"sendId,omitempty"`
	Channel          *int32  `protobuf:"varint,2,req,name=channel" json:"channel,omitempty"`
	RecvId           *int64  `protobuf:"varint,3,req,name=recvId" json:"recvId,omitempty"`
	MsgType          *int32  `protobuf:"varint,4,req,name=msgType" json:"msgType,omitempty"`
	Content          []byte  `protobuf:"bytes,5,req,name=content" json:"content,omitempty"`
	Args             *string `protobuf:"bytes,6,opt,name=args" json:"args,omitempty"`
	SendName         *string `protobuf:"bytes,7,opt,name=sendName" json:"sendName,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SCChatRecv) Reset()                    { *m = SCChatRecv{} }
func (m *SCChatRecv) String() string            { return proto.CompactTextString(m) }
func (*SCChatRecv) ProtoMessage()               {}
func (*SCChatRecv) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{2} }

func (m *SCChatRecv) GetSendId() int64 {
	if m != nil && m.SendId != nil {
		return *m.SendId
	}
	return 0
}

func (m *SCChatRecv) GetChannel() int32 {
	if m != nil && m.Channel != nil {
		return *m.Channel
	}
	return 0
}

func (m *SCChatRecv) GetRecvId() int64 {
	if m != nil && m.RecvId != nil {
		return *m.RecvId
	}
	return 0
}

func (m *SCChatRecv) GetMsgType() int32 {
	if m != nil && m.MsgType != nil {
		return *m.MsgType
	}
	return 0
}

func (m *SCChatRecv) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *SCChatRecv) GetArgs() string {
	if m != nil && m.Args != nil {
		return *m.Args
	}
	return ""
}

func (m *SCChatRecv) GetSendName() string {
	if m != nil && m.SendName != nil {
		return *m.SendName
	}
	return ""
}

type ChatData struct {
	SendId           *int64  `protobuf:"varint,1,req,name=sendId" json:"sendId,omitempty"`
	MsgType          *int32  `protobuf:"varint,2,req,name=msgType" json:"msgType,omitempty"`
	Content          []byte  `protobuf:"bytes,3,req,name=content" json:"content,omitempty"`
	SendTime         *int64  `protobuf:"varint,4,req,name=sendTime" json:"sendTime,omitempty"`
	Args             *string `protobuf:"bytes,5,opt,name=args" json:"args,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ChatData) Reset()                    { *m = ChatData{} }
func (m *ChatData) String() string            { return proto.CompactTextString(m) }
func (*ChatData) ProtoMessage()               {}
func (*ChatData) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{3} }

func (m *ChatData) GetSendId() int64 {
	if m != nil && m.SendId != nil {
		return *m.SendId
	}
	return 0
}

func (m *ChatData) GetMsgType() int32 {
	if m != nil && m.MsgType != nil {
		return *m.MsgType
	}
	return 0
}

func (m *ChatData) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *ChatData) GetSendTime() int64 {
	if m != nil && m.SendTime != nil {
		return *m.SendTime
	}
	return 0
}

func (m *ChatData) GetArgs() string {
	if m != nil && m.Args != nil {
		return *m.Args
	}
	return ""
}

type SCChatListNotice struct {
	WorldList        []*ChatData `protobuf:"bytes,1,rep,name=worldList" json:"worldList,omitempty"`
	SystemList       []*ChatData `protobuf:"bytes,2,rep,name=systemList" json:"systemList,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *SCChatListNotice) Reset()                    { *m = SCChatListNotice{} }
func (m *SCChatListNotice) String() string            { return proto.CompactTextString(m) }
func (*SCChatListNotice) ProtoMessage()               {}
func (*SCChatListNotice) Descriptor() ([]byte, []int) { return fileDescriptor15, []int{4} }

func (m *SCChatListNotice) GetWorldList() []*ChatData {
	if m != nil {
		return m.WorldList
	}
	return nil
}

func (m *SCChatListNotice) GetSystemList() []*ChatData {
	if m != nil {
		return m.SystemList
	}
	return nil
}

func init() {
	proto.RegisterType((*CSChatSend)(nil), "ui.CSChatSend")
	proto.RegisterType((*SCChatSend)(nil), "ui.SCChatSend")
	proto.RegisterType((*SCChatRecv)(nil), "ui.SCChatRecv")
	proto.RegisterType((*ChatData)(nil), "ui.ChatData")
	proto.RegisterType((*SCChatListNotice)(nil), "ui.SCChatListNotice")
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor15) }

var fileDescriptor15 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x14, 0x45, 0x15, 0xbb, 0x49, 0x9b, 0x47, 0x04, 0xc5, 0x93, 0x37, 0xac, 0x4c, 0x9e, 0x32, 0xf0,
	0x0b, 0x61, 0xa9, 0x84, 0x3a, 0x90, 0xb2, 0xc1, 0x60, 0x25, 0x56, 0x13, 0xa9, 0xb1, 0xa3, 0xf8,
	0xa5, 0xa8, 0x13, 0xbf, 0x8e, 0xec, 0x52, 0x04, 0x81, 0x95, 0xd5, 0xba, 0xbe, 0xe7, 0xdc, 0x07,
	0x50, 0xb7, 0x0a, 0x8b, 0x61, 0xb4, 0x68, 0x19, 0x99, 0xba, 0xfc, 0x15, 0xa0, 0xac, 0xca, 0x56,
	0x61, 0xa5, 0x4d, 0xc3, 0x6e, 0x60, 0x59, 0xb7, 0xca, 0x18, 0x7d, 0xe0, 0x91, 0x20, 0x32, 0x66,
	0xd7, 0x90, 0x8c, 0xba, 0x3e, 0x6e, 0x1a, 0x4e, 0x04, 0x91, 0xd4, 0x07, 0x7a, 0xb7, 0xdf, 0x9d,
	0x06, 0xcd, 0x69, 0x08, 0xf8, 0x1f, 0xd6, 0xa0, 0x36, 0xc8, 0x17, 0x82, 0xc8, 0x8c, 0x65, 0xb0,
	0x50, 0xe3, 0xde, 0xf1, 0x58, 0x44, 0x32, 0xcd, 0x07, 0x80, 0xaa, 0xfc, 0xb7, 0x7a, 0x76, 0x0b,
	0xa9, 0xdf, 0x53, 0xda, 0xc9, 0x20, 0x4f, 0x44, 0x24, 0xe3, 0xfc, 0xfd, 0x42, 0x7c, 0xd2, 0xf5,
	0xd1, 0x03, 0x9c, 0x36, 0xcd, 0xa6, 0x09, 0x40, 0xfa, 0xdd, 0x80, 0xcc, 0x0c, 0xe8, 0xdc, 0x60,
	0x31, 0x37, 0x88, 0x7f, 0x18, 0x24, 0xc1, 0x60, 0x0d, 0x2b, 0x0f, 0xd8, 0xaa, 0x5e, 0xf3, 0x65,
	0x98, 0xfc, 0x02, 0x2b, 0x8f, 0x7f, 0x50, 0xa8, 0xfe, 0xc2, 0x5f, 0xda, 0xc9, 0xbc, 0x9d, 0x86,
	0xf6, 0xcf, 0xbe, 0x5d, 0xd7, 0x9f, 0x05, 0xe8, 0xec, 0xa0, 0xcf, 0xb0, 0x3e, 0xcf, 0x7b, 0xec,
	0x1c, 0x6e, 0x2d, 0x76, 0xb5, 0x66, 0x77, 0x90, 0xbe, 0xd9, 0xf1, 0xd0, 0xf8, 0x27, 0x1e, 0x09,
	0x2a, 0xaf, 0xee, 0xb3, 0x62, 0xea, 0x8a, 0x2f, 0x0d, 0x01, 0xe0, 0x4e, 0x0e, 0x75, 0x1f, 0x12,
	0xe4, 0x77, 0xe2, 0x23, 0x00, 0x00, 0xff, 0xff, 0x78, 0xa7, 0x94, 0x0b, 0x17, 0x02, 0x00, 0x00,
}
