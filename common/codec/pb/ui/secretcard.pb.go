// Code generated by protoc-gen-go. DO NOT EDIT.
// source: secretcard.proto

package ui

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SecretCard struct {
	CardId           *int32 `protobuf:"varint,1,req,name=cardId" json:"cardId,omitempty"`
	Star             *int32 `protobuf:"varint,2,req,name=star" json:"star,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SecretCard) Reset()                    { *m = SecretCard{} }
func (m *SecretCard) String() string            { return proto.CompactTextString(m) }
func (*SecretCard) ProtoMessage()               {}
func (*SecretCard) Descriptor() ([]byte, []int) { return fileDescriptor92, []int{0} }

func (m *SecretCard) GetCardId() int32 {
	if m != nil && m.CardId != nil {
		return *m.CardId
	}
	return 0
}

func (m *SecretCard) GetStar() int32 {
	if m != nil && m.Star != nil {
		return *m.Star
	}
	return 0
}

type CSQuestSecretCardGet struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSQuestSecretCardGet) Reset()                    { *m = CSQuestSecretCardGet{} }
func (m *CSQuestSecretCardGet) String() string            { return proto.CompactTextString(m) }
func (*CSQuestSecretCardGet) ProtoMessage()               {}
func (*CSQuestSecretCardGet) Descriptor() ([]byte, []int) { return fileDescriptor92, []int{1} }

type SCQuestSecretCardGet struct {
	Num              *int32        `protobuf:"varint,1,req,name=num" json:"num,omitempty"`
	TotalStar        *int32        `protobuf:"varint,2,req,name=totalStar" json:"totalStar,omitempty"`
	OpenBoxList      []int32       `protobuf:"varint,3,rep,name=openBoxList" json:"openBoxList,omitempty"`
	CardList         []*SecretCard `protobuf:"bytes,4,rep,name=cardList" json:"cardList,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *SCQuestSecretCardGet) Reset()                    { *m = SCQuestSecretCardGet{} }
func (m *SCQuestSecretCardGet) String() string            { return proto.CompactTextString(m) }
func (*SCQuestSecretCardGet) ProtoMessage()               {}
func (*SCQuestSecretCardGet) Descriptor() ([]byte, []int) { return fileDescriptor92, []int{2} }

func (m *SCQuestSecretCardGet) GetNum() int32 {
	if m != nil && m.Num != nil {
		return *m.Num
	}
	return 0
}

func (m *SCQuestSecretCardGet) GetTotalStar() int32 {
	if m != nil && m.TotalStar != nil {
		return *m.TotalStar
	}
	return 0
}

func (m *SCQuestSecretCardGet) GetOpenBoxList() []int32 {
	if m != nil {
		return m.OpenBoxList
	}
	return nil
}

func (m *SCQuestSecretCardGet) GetCardList() []*SecretCard {
	if m != nil {
		return m.CardList
	}
	return nil
}

type CSQuestSecretSpy struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSQuestSecretSpy) Reset()                    { *m = CSQuestSecretSpy{} }
func (m *CSQuestSecretSpy) String() string            { return proto.CompactTextString(m) }
func (*CSQuestSecretSpy) ProtoMessage()               {}
func (*CSQuestSecretSpy) Descriptor() ([]byte, []int) { return fileDescriptor92, []int{3} }

type SCQuestSecretSpy struct {
	Num              *int32        `protobuf:"varint,1,req,name=num" json:"num,omitempty"`
	CardList         []*SecretCard `protobuf:"bytes,2,rep,name=cardList" json:"cardList,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *SCQuestSecretSpy) Reset()                    { *m = SCQuestSecretSpy{} }
func (m *SCQuestSecretSpy) String() string            { return proto.CompactTextString(m) }
func (*SCQuestSecretSpy) ProtoMessage()               {}
func (*SCQuestSecretSpy) Descriptor() ([]byte, []int) { return fileDescriptor92, []int{4} }

func (m *SCQuestSecretSpy) GetNum() int32 {
	if m != nil && m.Num != nil {
		return *m.Num
	}
	return 0
}

func (m *SCQuestSecretSpy) GetCardList() []*SecretCard {
	if m != nil {
		return m.CardList
	}
	return nil
}

type CSQuestSecretPickUp struct {
	CardId           *int32 `protobuf:"varint,1,req,name=cardId" json:"cardId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSQuestSecretPickUp) Reset()                    { *m = CSQuestSecretPickUp{} }
func (m *CSQuestSecretPickUp) String() string            { return proto.CompactTextString(m) }
func (*CSQuestSecretPickUp) ProtoMessage()               {}
func (*CSQuestSecretPickUp) Descriptor() ([]byte, []int) { return fileDescriptor92, []int{5} }

func (m *CSQuestSecretPickUp) GetCardId() int32 {
	if m != nil && m.CardId != nil {
		return *m.CardId
	}
	return 0
}

type SCQuestSecretPickUp struct {
	CardId           *int32 `protobuf:"varint,1,req,name=cardId" json:"cardId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCQuestSecretPickUp) Reset()                    { *m = SCQuestSecretPickUp{} }
func (m *SCQuestSecretPickUp) String() string            { return proto.CompactTextString(m) }
func (*SCQuestSecretPickUp) ProtoMessage()               {}
func (*SCQuestSecretPickUp) Descriptor() ([]byte, []int) { return fileDescriptor92, []int{6} }

func (m *SCQuestSecretPickUp) GetCardId() int32 {
	if m != nil && m.CardId != nil {
		return *m.CardId
	}
	return 0
}

type CSQuestSecretDiscard struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSQuestSecretDiscard) Reset()                    { *m = CSQuestSecretDiscard{} }
func (m *CSQuestSecretDiscard) String() string            { return proto.CompactTextString(m) }
func (*CSQuestSecretDiscard) ProtoMessage()               {}
func (*CSQuestSecretDiscard) Descriptor() ([]byte, []int) { return fileDescriptor92, []int{7} }

type CSQuestSecretStarRew struct {
	OpenBox          *int32 `protobuf:"varint,1,req,name=openBox" json:"openBox,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSQuestSecretStarRew) Reset()                    { *m = CSQuestSecretStarRew{} }
func (m *CSQuestSecretStarRew) String() string            { return proto.CompactTextString(m) }
func (*CSQuestSecretStarRew) ProtoMessage()               {}
func (*CSQuestSecretStarRew) Descriptor() ([]byte, []int) { return fileDescriptor92, []int{8} }

func (m *CSQuestSecretStarRew) GetOpenBox() int32 {
	if m != nil && m.OpenBox != nil {
		return *m.OpenBox
	}
	return 0
}

type SCQuestSecretStarRew struct {
	OpenBox          *int32 `protobuf:"varint,1,req,name=openBox" json:"openBox,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCQuestSecretStarRew) Reset()                    { *m = SCQuestSecretStarRew{} }
func (m *SCQuestSecretStarRew) String() string            { return proto.CompactTextString(m) }
func (*SCQuestSecretStarRew) ProtoMessage()               {}
func (*SCQuestSecretStarRew) Descriptor() ([]byte, []int) { return fileDescriptor92, []int{9} }

func (m *SCQuestSecretStarRew) GetOpenBox() int32 {
	if m != nil && m.OpenBox != nil {
		return *m.OpenBox
	}
	return 0
}

type CSQuestSecretFinish struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSQuestSecretFinish) Reset()                    { *m = CSQuestSecretFinish{} }
func (m *CSQuestSecretFinish) String() string            { return proto.CompactTextString(m) }
func (*CSQuestSecretFinish) ProtoMessage()               {}
func (*CSQuestSecretFinish) Descriptor() ([]byte, []int) { return fileDescriptor92, []int{10} }

type SCQuestSecretFinish struct {
	ItemList         []*ItemInfo `protobuf:"bytes,1,rep,name=itemList" json:"itemList,omitempty"`
	Result           *bool       `protobuf:"varint,2,req,name=result" json:"result,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *SCQuestSecretFinish) Reset()                    { *m = SCQuestSecretFinish{} }
func (m *SCQuestSecretFinish) String() string            { return proto.CompactTextString(m) }
func (*SCQuestSecretFinish) ProtoMessage()               {}
func (*SCQuestSecretFinish) Descriptor() ([]byte, []int) { return fileDescriptor92, []int{11} }

func (m *SCQuestSecretFinish) GetItemList() []*ItemInfo {
	if m != nil {
		return m.ItemList
	}
	return nil
}

func (m *SCQuestSecretFinish) GetResult() bool {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return false
}

type CSQuestSecretImmediate struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSQuestSecretImmediate) Reset()                    { *m = CSQuestSecretImmediate{} }
func (m *CSQuestSecretImmediate) String() string            { return proto.CompactTextString(m) }
func (*CSQuestSecretImmediate) ProtoMessage()               {}
func (*CSQuestSecretImmediate) Descriptor() ([]byte, []int) { return fileDescriptor92, []int{12} }

type SCQuestSecretImmediate struct {
	ItemInfo         []*ItemInfo `protobuf:"bytes,1,rep,name=itemInfo" json:"itemInfo,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *SCQuestSecretImmediate) Reset()                    { *m = SCQuestSecretImmediate{} }
func (m *SCQuestSecretImmediate) String() string            { return proto.CompactTextString(m) }
func (*SCQuestSecretImmediate) ProtoMessage()               {}
func (*SCQuestSecretImmediate) Descriptor() ([]byte, []int) { return fileDescriptor92, []int{13} }

func (m *SCQuestSecretImmediate) GetItemInfo() []*ItemInfo {
	if m != nil {
		return m.ItemInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*SecretCard)(nil), "ui.SecretCard")
	proto.RegisterType((*CSQuestSecretCardGet)(nil), "ui.CSQuestSecretCardGet")
	proto.RegisterType((*SCQuestSecretCardGet)(nil), "ui.SCQuestSecretCardGet")
	proto.RegisterType((*CSQuestSecretSpy)(nil), "ui.CSQuestSecretSpy")
	proto.RegisterType((*SCQuestSecretSpy)(nil), "ui.SCQuestSecretSpy")
	proto.RegisterType((*CSQuestSecretPickUp)(nil), "ui.CSQuestSecretPickUp")
	proto.RegisterType((*SCQuestSecretPickUp)(nil), "ui.SCQuestSecretPickUp")
	proto.RegisterType((*CSQuestSecretDiscard)(nil), "ui.CSQuestSecretDiscard")
	proto.RegisterType((*CSQuestSecretStarRew)(nil), "ui.CSQuestSecretStarRew")
	proto.RegisterType((*SCQuestSecretStarRew)(nil), "ui.SCQuestSecretStarRew")
	proto.RegisterType((*CSQuestSecretFinish)(nil), "ui.CSQuestSecretFinish")
	proto.RegisterType((*SCQuestSecretFinish)(nil), "ui.SCQuestSecretFinish")
	proto.RegisterType((*CSQuestSecretImmediate)(nil), "ui.CSQuestSecretImmediate")
	proto.RegisterType((*SCQuestSecretImmediate)(nil), "ui.SCQuestSecretImmediate")
}

func init() { proto.RegisterFile("secretcard.proto", fileDescriptor92) }

var fileDescriptor92 = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x4f, 0xc2, 0x30,
	0x1c, 0xc5, 0xc3, 0x06, 0x8a, 0x5f, 0x08, 0x62, 0x41, 0xb2, 0x78, 0x30, 0x4b, 0x13, 0x03, 0xf1,
	0xc0, 0xc1, 0x93, 0x57, 0xc5, 0x1f, 0x59, 0xe2, 0x41, 0x6d, 0xfc, 0x03, 0x96, 0xad, 0xc6, 0x46,
	0xba, 0x2e, 0xed, 0x77, 0x51, 0xfe, 0x7b, 0xd3, 0x31, 0xc1, 0x56, 0x38, 0xf6, 0xf5, 0xed, 0xf3,
	0xde, 0x5e, 0x61, 0x68, 0x78, 0xa6, 0x39, 0x66, 0xa9, 0xce, 0xe7, 0xa5, 0x56, 0xa8, 0x48, 0x50,
	0x89, 0x33, 0x10, 0xc8, 0xe5, 0xfa, 0x4c, 0x2f, 0x01, 0x58, 0xed, 0x59, 0xa4, 0x3a, 0x27, 0x03,
	0x38, 0xb0, 0xde, 0x24, 0x8f, 0x5a, 0x71, 0x30, 0xeb, 0x90, 0x3e, 0xb4, 0x0d, 0xa6, 0x3a, 0x0a,
	0xec, 0x89, 0x4e, 0x60, 0xbc, 0x60, 0x2f, 0x15, 0x37, 0xb8, 0xfd, 0xe4, 0x91, 0x23, 0x95, 0x30,
	0x66, 0x8b, 0xff, 0x3a, 0xe9, 0x41, 0x58, 0x54, 0xb2, 0x41, 0x9d, 0xc0, 0x11, 0x2a, 0x4c, 0x97,
	0x6c, 0xc3, 0x23, 0x23, 0xe8, 0xa9, 0x92, 0x17, 0xb7, 0xea, 0xfb, 0x49, 0x18, 0x8c, 0xc2, 0x38,
	0x9c, 0x75, 0x48, 0x0c, 0x5d, 0x5b, 0xa1, 0x56, 0xda, 0x71, 0x38, 0xeb, 0x5d, 0x0d, 0xe6, 0x95,
	0x98, 0x6f, 0xc9, 0x94, 0xc0, 0xd0, 0xa9, 0xc1, 0xca, 0x15, 0xbd, 0x81, 0xa1, 0x53, 0x81, 0x95,
	0x2b, 0x37, 0xfe, 0x2f, 0x36, 0xd8, 0x89, 0xbd, 0x80, 0x91, 0x83, 0x7d, 0x16, 0xd9, 0xe7, 0x5b,
	0xe9, 0x4f, 0x62, 0x6d, 0x4e, 0xd2, 0x1e, 0x9b, 0xbf, 0xd5, 0x9d, 0x30, 0xf6, 0x9e, 0x4e, 0x3d,
	0xdd, 0xce, 0xf1, 0xca, 0xbf, 0xc8, 0x31, 0x1c, 0x36, 0x5b, 0x34, 0x80, 0xa9, 0x37, 0xea, 0x5e,
	0xe3, 0xa9, 0xd7, 0xfb, 0x41, 0x14, 0xc2, 0x7c, 0xd0, 0x7b, 0xaf, 0xe7, 0x5a, 0x26, 0xe7, 0xd0,
	0xb5, 0xaf, 0x5f, 0xef, 0xd0, 0xaa, 0x77, 0xe8, 0xdb, 0x1d, 0x12, 0xe4, 0x32, 0x29, 0xde, 0x95,
	0xfd, 0x0f, 0xcd, 0x4d, 0xb5, 0xc4, 0xfa, 0x8d, 0xba, 0x34, 0x82, 0x89, 0x43, 0x4f, 0xa4, 0xe4,
	0xb9, 0x48, 0x91, 0xd3, 0x6b, 0x98, 0x38, 0x01, 0x9b, 0x9b, 0xdf, 0x0c, 0xcb, 0xdb, 0x95, 0xf1,
	0x13, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x26, 0x14, 0x2a, 0x96, 0x02, 0x00, 0x00,
}
