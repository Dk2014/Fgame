// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message_type.proto

package scene

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MessageType int32

const (
	MessageType_CS_PING_TYPE                  MessageType = 101
	MessageType_SC_PING_TYPE                  MessageType = 102
	MessageType_CS_ENTER_SCENE_TYPE           MessageType = 103
	MessageType_SC_ENTER_SCENE_TYPE           MessageType = 104
	MessageType_SC_OBJECT_ENTER_SCOPE_TYPE    MessageType = 105
	MessageType_SC_OBJECT_EXIT_SCOPE_TYPE     MessageType = 106
	MessageType_CS_OBJECT_MOVE_TYPE           MessageType = 107
	MessageType_SC_OBJECT_MOVE_TYPE           MessageType = 108
	MessageType_CS_OBJECT_ATTACK_TYPE         MessageType = 109
	MessageType_SC_OBJECT_ATTACK_TYPE         MessageType = 110
	MessageType_SC_OBJECT_DAMAGE_TYPE         MessageType = 111
	MessageType_SC_OBJECT_BUFF_TYPE           MessageType = 112
	MessageType_SC_OBJECT_BATTLE_TYPE         MessageType = 113
	MessageType_SC_PLAYER_DATA_CHANGED_TYPE   MessageType = 114
	MessageType_CS_PLAYER_RELIVE_TYPE         MessageType = 115
	MessageType_SC_PLAYER_RELIVE_TYPE         MessageType = 116
	MessageType_SC_OBJECT_BUFF_REMOVE_TYPE    MessageType = 117
	MessageType_CS_ITEM_GET_TYPE              MessageType = 118
	MessageType_SC_ITEM_GET_TYPE              MessageType = 119
	MessageType_SC_EXIT_SCENE_TYPE            MessageType = 120
	MessageType_SC_OBJECT_FIXED_POSITION_TYPE MessageType = 121
	MessageType_SC_MONSTER_CAMP_CHANGED_TYPE  MessageType = 122
	MessageType_SC_ITEM_OWNER_CHANGED_TYPE    MessageType = 123
	MessageType_CS_WORLD_ENTER_SCENE_TYPE     MessageType = 124
	MessageType_CS_ENTER_PORTAL_TYPE          MessageType = 125
	MessageType_SC_ENTER_PORTAL_TYPE          MessageType = 126
	MessageType_CS_REENTER_SCENE_TYPE         MessageType = 127
	MessageType_CS_SCENE_HEARTBEAT_TYPE       MessageType = 128
	MessageType_SC_SCENE_HEARTBEAT_TYPE       MessageType = 129
	MessageType_CS_PET_ATTACK_TYPE            MessageType = 130
	MessageType_SC_PET_ATTACK_TYPE            MessageType = 131
)

var MessageType_name = map[int32]string{
	101: "CS_PING_TYPE",
	102: "SC_PING_TYPE",
	103: "CS_ENTER_SCENE_TYPE",
	104: "SC_ENTER_SCENE_TYPE",
	105: "SC_OBJECT_ENTER_SCOPE_TYPE",
	106: "SC_OBJECT_EXIT_SCOPE_TYPE",
	107: "CS_OBJECT_MOVE_TYPE",
	108: "SC_OBJECT_MOVE_TYPE",
	109: "CS_OBJECT_ATTACK_TYPE",
	110: "SC_OBJECT_ATTACK_TYPE",
	111: "SC_OBJECT_DAMAGE_TYPE",
	112: "SC_OBJECT_BUFF_TYPE",
	113: "SC_OBJECT_BATTLE_TYPE",
	114: "SC_PLAYER_DATA_CHANGED_TYPE",
	115: "CS_PLAYER_RELIVE_TYPE",
	116: "SC_PLAYER_RELIVE_TYPE",
	117: "SC_OBJECT_BUFF_REMOVE_TYPE",
	118: "CS_ITEM_GET_TYPE",
	119: "SC_ITEM_GET_TYPE",
	120: "SC_EXIT_SCENE_TYPE",
	121: "SC_OBJECT_FIXED_POSITION_TYPE",
	122: "SC_MONSTER_CAMP_CHANGED_TYPE",
	123: "SC_ITEM_OWNER_CHANGED_TYPE",
	124: "CS_WORLD_ENTER_SCENE_TYPE",
	125: "CS_ENTER_PORTAL_TYPE",
	126: "SC_ENTER_PORTAL_TYPE",
	127: "CS_REENTER_SCENE_TYPE",
	128: "CS_SCENE_HEARTBEAT_TYPE",
	129: "SC_SCENE_HEARTBEAT_TYPE",
	130: "CS_PET_ATTACK_TYPE",
	131: "SC_PET_ATTACK_TYPE",
}
var MessageType_value = map[string]int32{
	"CS_PING_TYPE":                  101,
	"SC_PING_TYPE":                  102,
	"CS_ENTER_SCENE_TYPE":           103,
	"SC_ENTER_SCENE_TYPE":           104,
	"SC_OBJECT_ENTER_SCOPE_TYPE":    105,
	"SC_OBJECT_EXIT_SCOPE_TYPE":     106,
	"CS_OBJECT_MOVE_TYPE":           107,
	"SC_OBJECT_MOVE_TYPE":           108,
	"CS_OBJECT_ATTACK_TYPE":         109,
	"SC_OBJECT_ATTACK_TYPE":         110,
	"SC_OBJECT_DAMAGE_TYPE":         111,
	"SC_OBJECT_BUFF_TYPE":           112,
	"SC_OBJECT_BATTLE_TYPE":         113,
	"SC_PLAYER_DATA_CHANGED_TYPE":   114,
	"CS_PLAYER_RELIVE_TYPE":         115,
	"SC_PLAYER_RELIVE_TYPE":         116,
	"SC_OBJECT_BUFF_REMOVE_TYPE":    117,
	"CS_ITEM_GET_TYPE":              118,
	"SC_ITEM_GET_TYPE":              119,
	"SC_EXIT_SCENE_TYPE":            120,
	"SC_OBJECT_FIXED_POSITION_TYPE": 121,
	"SC_MONSTER_CAMP_CHANGED_TYPE":  122,
	"SC_ITEM_OWNER_CHANGED_TYPE":    123,
	"CS_WORLD_ENTER_SCENE_TYPE":     124,
	"CS_ENTER_PORTAL_TYPE":          125,
	"SC_ENTER_PORTAL_TYPE":          126,
	"CS_REENTER_SCENE_TYPE":         127,
	"CS_SCENE_HEARTBEAT_TYPE":       128,
	"SC_SCENE_HEARTBEAT_TYPE":       129,
	"CS_PET_ATTACK_TYPE":            130,
	"SC_PET_ATTACK_TYPE":            131,
}

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}
func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}
func (x *MessageType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MessageType_value, data, "MessageType")
	if err != nil {
		return err
	}
	*x = MessageType(value)
	return nil
}
func (MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func init() {
	proto.RegisterEnum("scene.MessageType", MessageType_name, MessageType_value)
}

func init() { proto.RegisterFile("message_type.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xd2, 0x49, 0xb3, 0xd3, 0x30,
	0x0c, 0x00, 0xe0, 0xe9, 0x01, 0x0e, 0x86, 0x43, 0x47, 0x3c, 0x28, 0x0f, 0xde, 0x63, 0xb9, 0x72,
	0xe0, 0x3f, 0x38, 0x8a, 0x9a, 0x17, 0x48, 0xe2, 0x8c, 0x6d, 0x68, 0x7b, 0xf2, 0x30, 0x4c, 0x28,
	0x5b, 0xdb, 0xd0, 0x94, 0xa5, 0xec, 0xcb, 0x6f, 0xe4, 0xff, 0x30, 0x8e, 0x4d, 0x96, 0xa6, 0x57,
	0x7d, 0x96, 0x65, 0x4b, 0x62, 0xb0, 0x2a, 0xaa, 0xea, 0xd9, 0xb2, 0x30, 0xbb, 0x7d, 0x59, 0x3c,
	0x2c, 0xb7, 0x9b, 0xdd, 0x06, 0x2e, 0x55, 0xcf, 0x8b, 0x75, 0xf1, 0xe0, 0xef, 0x65, 0x76, 0x25,
	0x75, 0xaa, 0xf7, 0x65, 0x01, 0x63, 0x76, 0x15, 0x95, 0xc9, 0xe3, 0x2c, 0x32, 0x7a, 0x91, 0xd3,
	0xb8, 0x8e, 0x28, 0xec, 0x44, 0x5e, 0xc0, 0x84, 0x5d, 0x43, 0x65, 0x28, 0xd3, 0x24, 0x8d, 0x42,
	0xca, 0xc8, 0xc1, 0xd2, 0x82, 0xc2, 0x21, 0xbc, 0x84, 0x3b, 0xec, 0x96, 0x42, 0x23, 0x82, 0x47,
	0x84, 0xba, 0x71, 0x91, 0x7b, 0x7f, 0x05, 0xe7, 0xec, 0xb4, 0xe3, 0xf3, 0x58, 0x77, 0xf9, 0xb5,
	0x2f, 0xe8, 0x39, 0x15, 0x4f, 0x3d, 0xbc, 0xf1, 0x05, 0x07, 0xf0, 0x16, 0x4e, 0xd9, 0xf5, 0x36,
	0x83, 0x6b, 0xcd, 0xf1, 0xb1, 0xa3, 0x95, 0xa5, 0x36, 0xa7, 0x4b, 0xeb, 0x3e, 0x85, 0x3c, 0xe5,
	0x91, 0xbf, 0x70, 0xd3, 0xaf, 0x14, 0x3c, 0x99, 0x4e, 0x1d, 0x94, 0xfd, 0x9c, 0x80, 0x6b, 0x9d,
	0xf8, 0x9c, 0x77, 0x70, 0x97, 0xdd, 0xb6, 0x9d, 0x4b, 0xf8, 0x82, 0xa4, 0x09, 0xb9, 0xe6, 0x06,
	0x2f, 0x78, 0x16, 0x51, 0xe8, 0x0e, 0x6c, 0xfd, 0x2b, 0xfd, 0x01, 0x49, 0x49, 0xfc, 0xff, 0x03,
	0x95, 0xbf, 0xf6, 0x08, 0xed, 0xfa, 0xcd, 0xac, 0x9f, 0x22, 0xa9, 0xfd, 0xfb, 0x7b, 0x38, 0x61,
	0x63, 0x54, 0x26, 0xd6, 0x94, 0x9a, 0x88, 0xb4, 0x8b, 0x7e, 0xb0, 0x51, 0x85, 0x07, 0xd1, 0x8f,
	0x70, 0x83, 0x81, 0x9d, 0x98, 0xeb, 0x78, 0x33, 0xb0, 0x4f, 0x70, 0x9f, 0x9d, 0xb7, 0x35, 0xa6,
	0xf1, 0x9c, 0x42, 0x93, 0x0b, 0x15, 0xeb, 0x58, 0x64, 0xee, 0xc8, 0x1e, 0xee, 0xb1, 0x33, 0x85,
	0x26, 0x15, 0x99, 0xb2, 0xe3, 0x44, 0x9e, 0xe6, 0xfd, 0xef, 0x7d, 0xf6, 0x0f, 0xad, 0x4b, 0x8a,
	0x59, 0x66, 0x0f, 0x75, 0xfd, 0x8b, 0x9d, 0x3a, 0x2a, 0x33, 0x13, 0x32, 0x09, 0x87, 0x4b, 0xf3,
	0x15, 0x6e, 0xb2, 0x93, 0x66, 0xcd, 0x72, 0x21, 0x35, 0x4f, 0x9c, 0x7c, 0xb3, 0xd2, 0xec, 0x59,
	0x57, 0xbe, 0xfb, 0x8e, 0x4a, 0x1a, 0x5c, 0xf7, 0x03, 0xce, 0xd8, 0x04, 0x95, 0x0f, 0x5d, 0x10,
	0x97, 0x3a, 0x20, 0xee, 0xfb, 0xf0, 0x73, 0x64, 0x55, 0xe1, 0x71, 0xfd, 0x35, 0x82, 0x09, 0x03,
	0x3b, 0x28, 0xea, 0x2f, 0xcc, 0xef, 0x1a, 0xec, 0x98, 0x0e, 0xe0, 0xcf, 0xe8, 0x5f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xa3, 0xde, 0xac, 0x9d, 0x73, 0x03, 0x00, 0x00,
}
