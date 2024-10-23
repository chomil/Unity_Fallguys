// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: messages.proto

package messages

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PlayerPosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X        float32 `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y        float32 `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
	Z        float32 `protobuf:"fixed32,3,opt,name=z,proto3" json:"z,omitempty"`
	Rx       float32 `protobuf:"fixed32,4,opt,name=rx,proto3" json:"rx,omitempty"` //각은 오일러로 받도록 하였음
	Ry       float32 `protobuf:"fixed32,5,opt,name=ry,proto3" json:"ry,omitempty"`
	Rz       float32 `protobuf:"fixed32,6,opt,name=rz,proto3" json:"rz,omitempty"`
	PlayerId string  `protobuf:"bytes,7,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
}

func (x *PlayerPosition) Reset() {
	*x = PlayerPosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerPosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerPosition) ProtoMessage() {}

func (x *PlayerPosition) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerPosition.ProtoReflect.Descriptor instead.
func (*PlayerPosition) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerPosition) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *PlayerPosition) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *PlayerPosition) GetZ() float32 {
	if x != nil {
		return x.Z
	}
	return 0
}

func (x *PlayerPosition) GetRx() float32 {
	if x != nil {
		return x.Rx
	}
	return 0
}

func (x *PlayerPosition) GetRy() float32 {
	if x != nil {
		return x.Ry
	}
	return 0
}

func (x *PlayerPosition) GetRz() float32 {
	if x != nil {
		return x.Rz
	}
	return 0
}

func (x *PlayerPosition) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

type ChatMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender  string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *ChatMessage) Reset() {
	*x = ChatMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMessage) ProtoMessage() {}

func (x *ChatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMessage.ProtoReflect.Descriptor instead.
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{1}
}

func (x *ChatMessage) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *ChatMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type LoginMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string `protobuf:"bytes,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
}

func (x *LoginMessage) Reset() {
	*x = LoginMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginMessage) ProtoMessage() {}

func (x *LoginMessage) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginMessage.ProtoReflect.Descriptor instead.
func (*LoginMessage) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{2}
}

func (x *LoginMessage) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

type PlayerAnimation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerAnimState string  `protobuf:"bytes,1,opt,name=player_anim_state,json=playerAnimState,proto3" json:"player_anim_state,omitempty"`
	PlayerId        string  `protobuf:"bytes,2,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	SpeedForward    float32 `protobuf:"fixed32,3,opt,name=SpeedForward,proto3" json:"SpeedForward,omitempty"`
	SpeedRight      float32 `protobuf:"fixed32,4,opt,name=SpeedRight,proto3" json:"SpeedRight,omitempty"`
}

func (x *PlayerAnimation) Reset() {
	*x = PlayerAnimation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerAnimation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerAnimation) ProtoMessage() {}

func (x *PlayerAnimation) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerAnimation.ProtoReflect.Descriptor instead.
func (*PlayerAnimation) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{3}
}

func (x *PlayerAnimation) GetPlayerAnimState() string {
	if x != nil {
		return x.PlayerAnimState
	}
	return ""
}

func (x *PlayerAnimation) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *PlayerAnimation) GetSpeedForward() float32 {
	if x != nil {
		return x.SpeedForward
	}
	return 0
}

func (x *PlayerAnimation) GetSpeedRight() float32 {
	if x != nil {
		return x.SpeedRight
	}
	return 0
}

type SpawnPlayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string  `protobuf:"bytes,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	X        float32 `protobuf:"fixed32,2,opt,name=x,proto3" json:"x,omitempty"`
	Y        float32 `protobuf:"fixed32,3,opt,name=y,proto3" json:"y,omitempty"`
	Z        float32 `protobuf:"fixed32,4,opt,name=z,proto3" json:"z,omitempty"`
	Rx       float32 `protobuf:"fixed32,5,opt,name=rx,proto3" json:"rx,omitempty"`
	Ry       float32 `protobuf:"fixed32,6,opt,name=ry,proto3" json:"ry,omitempty"`
	Rz       float32 `protobuf:"fixed32,7,opt,name=rz,proto3" json:"rz,omitempty"`
}

func (x *SpawnPlayer) Reset() {
	*x = SpawnPlayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpawnPlayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpawnPlayer) ProtoMessage() {}

func (x *SpawnPlayer) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpawnPlayer.ProtoReflect.Descriptor instead.
func (*SpawnPlayer) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{4}
}

func (x *SpawnPlayer) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *SpawnPlayer) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *SpawnPlayer) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *SpawnPlayer) GetZ() float32 {
	if x != nil {
		return x.Z
	}
	return 0
}

func (x *SpawnPlayer) GetRx() float32 {
	if x != nil {
		return x.Rx
	}
	return 0
}

func (x *SpawnPlayer) GetRy() float32 {
	if x != nil {
		return x.Ry
	}
	return 0
}

func (x *SpawnPlayer) GetRz() float32 {
	if x != nil {
		return x.Rz
	}
	return 0
}

type LogoutMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string `protobuf:"bytes,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
}

func (x *LogoutMessage) Reset() {
	*x = LogoutMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutMessage) ProtoMessage() {}

func (x *LogoutMessage) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutMessage.ProtoReflect.Descriptor instead.
func (*LogoutMessage) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{5}
}

func (x *LogoutMessage) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

type RaceFinishMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId   string `protobuf:"bytes,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	FinishTime int64  `protobuf:"varint,2,opt,name=finishTime,proto3" json:"finishTime,omitempty"`
}

func (x *RaceFinishMessage) Reset() {
	*x = RaceFinishMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaceFinishMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaceFinishMessage) ProtoMessage() {}

func (x *RaceFinishMessage) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaceFinishMessage.ProtoReflect.Descriptor instead.
func (*RaceFinishMessage) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{6}
}

func (x *RaceFinishMessage) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *RaceFinishMessage) GetFinishTime() int64 {
	if x != nil {
		return x.FinishTime
	}
	return 0
}

type CostumeMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId          string `protobuf:"bytes,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	PlayerCostumeType int32  `protobuf:"varint,2,opt,name=playerCostumeType,proto3" json:"playerCostumeType,omitempty"`
	PlayerCostumeName string `protobuf:"bytes,3,opt,name=playerCostumeName,proto3" json:"playerCostumeName,omitempty"`
}

func (x *CostumeMessage) Reset() {
	*x = CostumeMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CostumeMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CostumeMessage) ProtoMessage() {}

func (x *CostumeMessage) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CostumeMessage.ProtoReflect.Descriptor instead.
func (*CostumeMessage) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{7}
}

func (x *CostumeMessage) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *CostumeMessage) GetPlayerCostumeType() int32 {
	if x != nil {
		return x.PlayerCostumeType
	}
	return 0
}

func (x *CostumeMessage) GetPlayerCostumeName() string {
	if x != nil {
		return x.PlayerCostumeName
	}
	return ""
}

type RaceEndMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string `protobuf:"bytes,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
}

func (x *RaceEndMessage) Reset() {
	*x = RaceEndMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RaceEndMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RaceEndMessage) ProtoMessage() {}

func (x *RaceEndMessage) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RaceEndMessage.ProtoReflect.Descriptor instead.
func (*RaceEndMessage) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{8}
}

func (x *RaceEndMessage) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

type GameMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//
	//	*GameMessage_PlayerPosition
	//	*GameMessage_Chat
	//	*GameMessage_Login
	//	*GameMessage_PlayerAnimState
	//	*GameMessage_SpawnPlayer
	//	*GameMessage_SpawnExistingPlayer
	//	*GameMessage_Logout
	//	*GameMessage_RaceFinish
	//	*GameMessage_PlayerCostume
	//	*GameMessage_RaceEnd
	Message isGameMessage_Message `protobuf_oneof:"message"`
}

func (x *GameMessage) Reset() {
	*x = GameMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameMessage) ProtoMessage() {}

func (x *GameMessage) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameMessage.ProtoReflect.Descriptor instead.
func (*GameMessage) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{9}
}

func (m *GameMessage) GetMessage() isGameMessage_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *GameMessage) GetPlayerPosition() *PlayerPosition {
	if x, ok := x.GetMessage().(*GameMessage_PlayerPosition); ok {
		return x.PlayerPosition
	}
	return nil
}

func (x *GameMessage) GetChat() *ChatMessage {
	if x, ok := x.GetMessage().(*GameMessage_Chat); ok {
		return x.Chat
	}
	return nil
}

func (x *GameMessage) GetLogin() *LoginMessage {
	if x, ok := x.GetMessage().(*GameMessage_Login); ok {
		return x.Login
	}
	return nil
}

func (x *GameMessage) GetPlayerAnimState() *PlayerAnimation {
	if x, ok := x.GetMessage().(*GameMessage_PlayerAnimState); ok {
		return x.PlayerAnimState
	}
	return nil
}

func (x *GameMessage) GetSpawnPlayer() *SpawnPlayer {
	if x, ok := x.GetMessage().(*GameMessage_SpawnPlayer); ok {
		return x.SpawnPlayer
	}
	return nil
}

func (x *GameMessage) GetSpawnExistingPlayer() *SpawnPlayer {
	if x, ok := x.GetMessage().(*GameMessage_SpawnExistingPlayer); ok {
		return x.SpawnExistingPlayer
	}
	return nil
}

func (x *GameMessage) GetLogout() *LogoutMessage {
	if x, ok := x.GetMessage().(*GameMessage_Logout); ok {
		return x.Logout
	}
	return nil
}

func (x *GameMessage) GetRaceFinish() *RaceFinishMessage {
	if x, ok := x.GetMessage().(*GameMessage_RaceFinish); ok {
		return x.RaceFinish
	}
	return nil
}

func (x *GameMessage) GetPlayerCostume() *CostumeMessage {
	if x, ok := x.GetMessage().(*GameMessage_PlayerCostume); ok {
		return x.PlayerCostume
	}
	return nil
}

func (x *GameMessage) GetRaceEnd() *RaceEndMessage {
	if x, ok := x.GetMessage().(*GameMessage_RaceEnd); ok {
		return x.RaceEnd
	}
	return nil
}

type isGameMessage_Message interface {
	isGameMessage_Message()
}

type GameMessage_PlayerPosition struct {
	PlayerPosition *PlayerPosition `protobuf:"bytes,1,opt,name=player_position,json=playerPosition,proto3,oneof"`
}

type GameMessage_Chat struct {
	Chat *ChatMessage `protobuf:"bytes,2,opt,name=chat,proto3,oneof"`
}

type GameMessage_Login struct {
	Login *LoginMessage `protobuf:"bytes,3,opt,name=login,proto3,oneof"`
}

type GameMessage_PlayerAnimState struct {
	PlayerAnimState *PlayerAnimation `protobuf:"bytes,4,opt,name=player_anim_state,json=playerAnimState,proto3,oneof"`
}

type GameMessage_SpawnPlayer struct {
	SpawnPlayer *SpawnPlayer `protobuf:"bytes,5,opt,name=spawnPlayer,proto3,oneof"`
}

type GameMessage_SpawnExistingPlayer struct {
	SpawnExistingPlayer *SpawnPlayer `protobuf:"bytes,6,opt,name=spawnExistingPlayer,proto3,oneof"`
}

type GameMessage_Logout struct {
	Logout *LogoutMessage `protobuf:"bytes,7,opt,name=logout,proto3,oneof"`
}

type GameMessage_RaceFinish struct {
	RaceFinish *RaceFinishMessage `protobuf:"bytes,8,opt,name=race_finish,json=raceFinish,proto3,oneof"`
}

type GameMessage_PlayerCostume struct {
	PlayerCostume *CostumeMessage `protobuf:"bytes,9,opt,name=player_costume,json=playerCostume,proto3,oneof"`
}

type GameMessage_RaceEnd struct {
	RaceEnd *RaceEndMessage `protobuf:"bytes,10,opt,name=race_end,json=raceEnd,proto3,oneof"`
}

func (*GameMessage_PlayerPosition) isGameMessage_Message() {}

func (*GameMessage_Chat) isGameMessage_Message() {}

func (*GameMessage_Login) isGameMessage_Message() {}

func (*GameMessage_PlayerAnimState) isGameMessage_Message() {}

func (*GameMessage_SpawnPlayer) isGameMessage_Message() {}

func (*GameMessage_SpawnExistingPlayer) isGameMessage_Message() {}

func (*GameMessage_Logout) isGameMessage_Message() {}

func (*GameMessage_RaceFinish) isGameMessage_Message() {}

func (*GameMessage_PlayerCostume) isGameMessage_Message() {}

func (*GameMessage_RaceEnd) isGameMessage_Message() {}

var File_messages_proto protoreflect.FileDescriptor

var file_messages_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x22, 0x87, 0x01, 0x0a, 0x0e, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x01, 0x7a, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x02, 0x72, 0x78, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x02, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x7a, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x02, 0x72, 0x7a, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x3f, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x22, 0x2a, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x22, 0x9e, 0x01,
	0x0a, 0x0f, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x61, 0x6e, 0x69, 0x6d,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x41, 0x6e, 0x69, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x70,
	0x65, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0c, 0x53, 0x70, 0x65, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x53, 0x70, 0x65, 0x65, 0x64, 0x52, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0a, 0x53, 0x70, 0x65, 0x65, 0x64, 0x52, 0x69, 0x67, 0x68, 0x74, 0x22, 0x83,
	0x01, 0x0a, 0x0b, 0x53, 0x70, 0x61, 0x77, 0x6e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x01, 0x7a, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x02, 0x72, 0x78, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x02, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x7a, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x02, 0x72, 0x7a, 0x22, 0x2b, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x4f, 0x0a, 0x11, 0x52, 0x61, 0x63, 0x65, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x54, 0x69,
	0x6d, 0x65, 0x22, 0x88, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x73, 0x74, 0x75, 0x6d, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x2c, 0x0a, 0x11, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6f, 0x73, 0x74, 0x75,
	0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x43, 0x6f, 0x73, 0x74, 0x75, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x2c, 0x0a, 0x11, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6f, 0x73, 0x74, 0x75, 0x6d, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x43, 0x6f, 0x73, 0x74, 0x75, 0x6d, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2c, 0x0a,
	0x0e, 0x52, 0x61, 0x63, 0x65, 0x45, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x22, 0xce, 0x04, 0x0a, 0x0b,
	0x47, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3f, 0x0a, 0x0f, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0e, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x04,
	0x63, 0x68, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x61, 0x6d,
	0x65, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52,
	0x04, 0x63, 0x68, 0x61, 0x74, 0x12, 0x2a, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x12, 0x43, 0x0a, 0x11, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x61, 0x6e, 0x69, 0x6d,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x6e, 0x69,
	0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x35, 0x0a, 0x0b, 0x73, 0x70, 0x61, 0x77, 0x6e, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x61,
	0x6d, 0x65, 0x2e, 0x53, 0x70, 0x61, 0x77, 0x6e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x48, 0x00,
	0x52, 0x0b, 0x73, 0x70, 0x61, 0x77, 0x6e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x45, 0x0a,
	0x13, 0x73, 0x70, 0x61, 0x77, 0x6e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x61, 0x6d,
	0x65, 0x2e, 0x53, 0x70, 0x61, 0x77, 0x6e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x48, 0x00, 0x52,
	0x13, 0x73, 0x70, 0x61, 0x77, 0x6e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x06, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x6f,
	0x75, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x06, 0x6c, 0x6f, 0x67,
	0x6f, 0x75, 0x74, 0x12, 0x3a, 0x0a, 0x0b, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x66, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e,
	0x52, 0x61, 0x63, 0x65, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x48, 0x00, 0x52, 0x0a, 0x72, 0x61, 0x63, 0x65, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x12,
	0x3d, 0x0a, 0x0e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x75, 0x6d,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x43,
	0x6f, 0x73, 0x74, 0x75, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52,
	0x0d, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6f, 0x73, 0x74, 0x75, 0x6d, 0x65, 0x12, 0x31,
	0x0a, 0x08, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x52, 0x61, 0x63, 0x65, 0x45, 0x6e, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x07, 0x72, 0x61, 0x63, 0x65, 0x45, 0x6e,
	0x64, 0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x0d, 0x5a, 0x0b,
	0x47, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_messages_proto_rawDescOnce sync.Once
	file_messages_proto_rawDescData = file_messages_proto_rawDesc
)

func file_messages_proto_rawDescGZIP() []byte {
	file_messages_proto_rawDescOnce.Do(func() {
		file_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_messages_proto_rawDescData)
	})
	return file_messages_proto_rawDescData
}

var file_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_messages_proto_goTypes = []any{
	(*PlayerPosition)(nil),    // 0: game.PlayerPosition
	(*ChatMessage)(nil),       // 1: game.ChatMessage
	(*LoginMessage)(nil),      // 2: game.LoginMessage
	(*PlayerAnimation)(nil),   // 3: game.PlayerAnimation
	(*SpawnPlayer)(nil),       // 4: game.SpawnPlayer
	(*LogoutMessage)(nil),     // 5: game.LogoutMessage
	(*RaceFinishMessage)(nil), // 6: game.RaceFinishMessage
	(*CostumeMessage)(nil),    // 7: game.CostumeMessage
	(*RaceEndMessage)(nil),    // 8: game.RaceEndMessage
	(*GameMessage)(nil),       // 9: game.GameMessage
}
var file_messages_proto_depIdxs = []int32{
	0,  // 0: game.GameMessage.player_position:type_name -> game.PlayerPosition
	1,  // 1: game.GameMessage.chat:type_name -> game.ChatMessage
	2,  // 2: game.GameMessage.login:type_name -> game.LoginMessage
	3,  // 3: game.GameMessage.player_anim_state:type_name -> game.PlayerAnimation
	4,  // 4: game.GameMessage.spawnPlayer:type_name -> game.SpawnPlayer
	4,  // 5: game.GameMessage.spawnExistingPlayer:type_name -> game.SpawnPlayer
	5,  // 6: game.GameMessage.logout:type_name -> game.LogoutMessage
	6,  // 7: game.GameMessage.race_finish:type_name -> game.RaceFinishMessage
	7,  // 8: game.GameMessage.player_costume:type_name -> game.CostumeMessage
	8,  // 9: game.GameMessage.race_end:type_name -> game.RaceEndMessage
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_messages_proto_init() }
func file_messages_proto_init() {
	if File_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_messages_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PlayerPosition); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_messages_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ChatMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_messages_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*LoginMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_messages_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*PlayerAnimation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_messages_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*SpawnPlayer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_messages_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*LogoutMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_messages_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*RaceFinishMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_messages_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*CostumeMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_messages_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*RaceEndMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_messages_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*GameMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_messages_proto_msgTypes[9].OneofWrappers = []any{
		(*GameMessage_PlayerPosition)(nil),
		(*GameMessage_Chat)(nil),
		(*GameMessage_Login)(nil),
		(*GameMessage_PlayerAnimState)(nil),
		(*GameMessage_SpawnPlayer)(nil),
		(*GameMessage_SpawnExistingPlayer)(nil),
		(*GameMessage_Logout)(nil),
		(*GameMessage_RaceFinish)(nil),
		(*GameMessage_PlayerCostume)(nil),
		(*GameMessage_RaceEnd)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messages_proto_goTypes,
		DependencyIndexes: file_messages_proto_depIdxs,
		MessageInfos:      file_messages_proto_msgTypes,
	}.Build()
	File_messages_proto = out.File
	file_messages_proto_rawDesc = nil
	file_messages_proto_goTypes = nil
	file_messages_proto_depIdxs = nil
}
