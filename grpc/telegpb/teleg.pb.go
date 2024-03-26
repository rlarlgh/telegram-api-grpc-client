// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: grpc/proto/teleg.proto

package telegpb

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_teleg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_teleg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_grpc_proto_teleg_proto_rawDescGZIP(), []int{0}
}

type ServerPushResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BotName          string `protobuf:"bytes,1,opt,name=bot_name,json=botName,proto3" json:"bot_name,omitempty"`
	Type             string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Text             string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	ChatId           int64  `protobuf:"varint,4,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	FromId           int64  `protobuf:"varint,5,opt,name=from_id,json=fromId,proto3" json:"from_id,omitempty"`
	MessageId        int32  `protobuf:"varint,6,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	ReplyToMessageId int32  `protobuf:"varint,7,opt,name=reply_to_message_id,json=replyToMessageId,proto3" json:"reply_to_message_id,omitempty"`
}

func (x *ServerPushResponse) Reset() {
	*x = ServerPushResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_teleg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerPushResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerPushResponse) ProtoMessage() {}

func (x *ServerPushResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_teleg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerPushResponse.ProtoReflect.Descriptor instead.
func (*ServerPushResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_teleg_proto_rawDescGZIP(), []int{1}
}

func (x *ServerPushResponse) GetBotName() string {
	if x != nil {
		return x.BotName
	}
	return ""
}

func (x *ServerPushResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ServerPushResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *ServerPushResponse) GetChatId() int64 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

func (x *ServerPushResponse) GetFromId() int64 {
	if x != nil {
		return x.FromId
	}
	return 0
}

func (x *ServerPushResponse) GetMessageId() int32 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

func (x *ServerPushResponse) GetReplyToMessageId() int32 {
	if x != nil {
		return x.ReplyToMessageId
	}
	return 0
}

type SendMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BotName          string `protobuf:"bytes,1,opt,name=bot_name,json=botName,proto3" json:"bot_name,omitempty"`
	Text             string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	ChatId           int64  `protobuf:"varint,3,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	ReplyToMessageId int32  `protobuf:"varint,4,opt,name=reply_to_message_id,json=replyToMessageId,proto3" json:"reply_to_message_id,omitempty"`
}

func (x *SendMessageRequest) Reset() {
	*x = SendMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_teleg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageRequest) ProtoMessage() {}

func (x *SendMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_teleg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageRequest.ProtoReflect.Descriptor instead.
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_teleg_proto_rawDescGZIP(), []int{2}
}

func (x *SendMessageRequest) GetBotName() string {
	if x != nil {
		return x.BotName
	}
	return ""
}

func (x *SendMessageRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *SendMessageRequest) GetChatId() int64 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

func (x *SendMessageRequest) GetReplyToMessageId() int32 {
	if x != nil {
		return x.ReplyToMessageId
	}
	return 0
}

type SendMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BotName string `protobuf:"bytes,1,opt,name=bot_name,json=botName,proto3" json:"bot_name,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SendMessageResponse) Reset() {
	*x = SendMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_teleg_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageResponse) ProtoMessage() {}

func (x *SendMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_teleg_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageResponse.ProtoReflect.Descriptor instead.
func (*SendMessageResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_teleg_proto_rawDescGZIP(), []int{3}
}

func (x *SendMessageResponse) GetBotName() string {
	if x != nil {
		return x.BotName
	}
	return ""
}

func (x *SendMessageResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SendKeyboardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BotName          string   `protobuf:"bytes,1,opt,name=bot_name,json=botName,proto3" json:"bot_name,omitempty"`
	ChatId           int64    `protobuf:"varint,2,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	Text             string   `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	Buttons          []string `protobuf:"bytes,4,rep,name=buttons,proto3" json:"buttons,omitempty"`
	RowNum           int32    `protobuf:"varint,5,opt,name=row_num,json=rowNum,proto3" json:"row_num,omitempty"`
	ReplyToMessageId int32    `protobuf:"varint,6,opt,name=reply_to_message_id,json=replyToMessageId,proto3" json:"reply_to_message_id,omitempty"`
}

func (x *SendKeyboardRequest) Reset() {
	*x = SendKeyboardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_teleg_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendKeyboardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendKeyboardRequest) ProtoMessage() {}

func (x *SendKeyboardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_teleg_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendKeyboardRequest.ProtoReflect.Descriptor instead.
func (*SendKeyboardRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_teleg_proto_rawDescGZIP(), []int{4}
}

func (x *SendKeyboardRequest) GetBotName() string {
	if x != nil {
		return x.BotName
	}
	return ""
}

func (x *SendKeyboardRequest) GetChatId() int64 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

func (x *SendKeyboardRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *SendKeyboardRequest) GetButtons() []string {
	if x != nil {
		return x.Buttons
	}
	return nil
}

func (x *SendKeyboardRequest) GetRowNum() int32 {
	if x != nil {
		return x.RowNum
	}
	return 0
}

func (x *SendKeyboardRequest) GetReplyToMessageId() int32 {
	if x != nil {
		return x.ReplyToMessageId
	}
	return 0
}

type SendKeyboardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BotName string `protobuf:"bytes,1,opt,name=bot_name,json=botName,proto3" json:"bot_name,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SendKeyboardResponse) Reset() {
	*x = SendKeyboardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_teleg_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendKeyboardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendKeyboardResponse) ProtoMessage() {}

func (x *SendKeyboardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_teleg_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendKeyboardResponse.ProtoReflect.Descriptor instead.
func (*SendKeyboardResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_teleg_proto_rawDescGZIP(), []int{5}
}

func (x *SendKeyboardResponse) GetBotName() string {
	if x != nil {
		return x.BotName
	}
	return ""
}

func (x *SendKeyboardResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SendReplyInputRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BotName          string `protobuf:"bytes,1,opt,name=bot_name,json=botName,proto3" json:"bot_name,omitempty"`
	ChatId           int64  `protobuf:"varint,2,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	Text             string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	ReplyMarkup      string `protobuf:"bytes,4,opt,name=reply_markup,json=replyMarkup,proto3" json:"reply_markup,omitempty"`
	ReplyToMessageId int32  `protobuf:"varint,5,opt,name=reply_to_message_id,json=replyToMessageId,proto3" json:"reply_to_message_id,omitempty"`
}

func (x *SendReplyInputRequest) Reset() {
	*x = SendReplyInputRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_teleg_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendReplyInputRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendReplyInputRequest) ProtoMessage() {}

func (x *SendReplyInputRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_teleg_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendReplyInputRequest.ProtoReflect.Descriptor instead.
func (*SendReplyInputRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_teleg_proto_rawDescGZIP(), []int{6}
}

func (x *SendReplyInputRequest) GetBotName() string {
	if x != nil {
		return x.BotName
	}
	return ""
}

func (x *SendReplyInputRequest) GetChatId() int64 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

func (x *SendReplyInputRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *SendReplyInputRequest) GetReplyMarkup() string {
	if x != nil {
		return x.ReplyMarkup
	}
	return ""
}

func (x *SendReplyInputRequest) GetReplyToMessageId() int32 {
	if x != nil {
		return x.ReplyToMessageId
	}
	return 0
}

type SendReplyInputResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BotName string `protobuf:"bytes,1,opt,name=bot_name,json=botName,proto3" json:"bot_name,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SendReplyInputResponse) Reset() {
	*x = SendReplyInputResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_teleg_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendReplyInputResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendReplyInputResponse) ProtoMessage() {}

func (x *SendReplyInputResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_teleg_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendReplyInputResponse.ProtoReflect.Descriptor instead.
func (*SendReplyInputResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_teleg_proto_rawDescGZIP(), []int{7}
}

func (x *SendReplyInputResponse) GetBotName() string {
	if x != nil {
		return x.BotName
	}
	return ""
}

func (x *SendReplyInputResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_grpc_proto_teleg_proto protoreflect.FileDescriptor

var file_grpc_proto_teleg_proto_rawDesc = []byte{
	0x0a, 0x16, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x6c,
	0x65, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74,
	0x65, 0x6c, 0x65, 0x67, 0x70, 0x62, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0xd7, 0x01, 0x0a, 0x12, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6f, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x6f, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x66, 0x72, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x13, 0x72, 0x65,
	0x70, 0x6c, 0x79, 0x5f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x8b, 0x01, 0x0a, 0x12, 0x53, 0x65,
	0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x62, 0x6f, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x62, 0x6f, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x13, 0x72, 0x65, 0x70, 0x6c,
	0x79, 0x5f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x4a, 0x0a, 0x13, 0x53, 0x65, 0x6e, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x62, 0x6f, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x62, 0x6f, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0xbf, 0x01, 0x0a, 0x13, 0x53, 0x65, 0x6e, 0x64, 0x4b, 0x65, 0x79, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x62,
	0x6f, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62,
	0x6f, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x73, 0x12, 0x17, 0x0a,
	0x07, 0x72, 0x6f, 0x77, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x72, 0x6f, 0x77, 0x4e, 0x75, 0x6d, 0x12, 0x2d, 0x0a, 0x13, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f,
	0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x10, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x4b, 0x0a, 0x14, 0x53, 0x65, 0x6e, 0x64, 0x4b, 0x65, 0x79,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x62, 0x6f, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x62, 0x6f, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0xb1, 0x01, 0x0a, 0x15, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x62, 0x6f, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x62, 0x6f, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x5f, 0x6d, 0x61,
	0x72, 0x6b, 0x75, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x70, 0x6c,
	0x79, 0x4d, 0x61, 0x72, 0x6b, 0x75, 0x70, 0x12, 0x2d, 0x0a, 0x13, 0x72, 0x65, 0x70, 0x6c, 0x79,
	0x5f, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x54, 0x6f, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x4d, 0x0a, 0x16, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x62, 0x6f, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x62, 0x6f, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xe1, 0x02, 0x0a, 0x0b, 0x54, 0x65, 0x6c, 0x65, 0x67, 0x72,
	0x61, 0x6d, 0x41, 0x70, 0x69, 0x12, 0x4a, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x50, 0x75, 0x73, 0x68, 0x12, 0x13, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x74, 0x65, 0x6c, 0x65, 0x67, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x20, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30,
	0x01, 0x12, 0x52, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x20, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x70, 0x62, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x70,
	0x62, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64, 0x4b, 0x65, 0x79,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x21, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x6c,
	0x65, 0x67, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4b, 0x65, 0x79, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x74, 0x65, 0x6c, 0x65, 0x67, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4b, 0x65, 0x79, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x0e,
	0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x23,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x70, 0x62, 0x2e, 0x53, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x67,
	0x70, 0x62, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0e, 0x5a, 0x0c, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_grpc_proto_teleg_proto_rawDescOnce sync.Once
	file_grpc_proto_teleg_proto_rawDescData = file_grpc_proto_teleg_proto_rawDesc
)

func file_grpc_proto_teleg_proto_rawDescGZIP() []byte {
	file_grpc_proto_teleg_proto_rawDescOnce.Do(func() {
		file_grpc_proto_teleg_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_proto_teleg_proto_rawDescData)
	})
	return file_grpc_proto_teleg_proto_rawDescData
}

var file_grpc_proto_teleg_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_grpc_proto_teleg_proto_goTypes = []interface{}{
	(*Empty)(nil),                  // 0: grpc.telegpb.Empty
	(*ServerPushResponse)(nil),     // 1: grpc.telegpb.ServerPushResponse
	(*SendMessageRequest)(nil),     // 2: grpc.telegpb.SendMessageRequest
	(*SendMessageResponse)(nil),    // 3: grpc.telegpb.SendMessageResponse
	(*SendKeyboardRequest)(nil),    // 4: grpc.telegpb.SendKeyboardRequest
	(*SendKeyboardResponse)(nil),   // 5: grpc.telegpb.SendKeyboardResponse
	(*SendReplyInputRequest)(nil),  // 6: grpc.telegpb.SendReplyInputRequest
	(*SendReplyInputResponse)(nil), // 7: grpc.telegpb.SendReplyInputResponse
}
var file_grpc_proto_teleg_proto_depIdxs = []int32{
	0, // 0: grpc.telegpb.TelegramApi.StartServerPush:input_type -> grpc.telegpb.Empty
	2, // 1: grpc.telegpb.TelegramApi.SendMessage:input_type -> grpc.telegpb.SendMessageRequest
	4, // 2: grpc.telegpb.TelegramApi.SendKeyboard:input_type -> grpc.telegpb.SendKeyboardRequest
	6, // 3: grpc.telegpb.TelegramApi.SendReplyInput:input_type -> grpc.telegpb.SendReplyInputRequest
	1, // 4: grpc.telegpb.TelegramApi.StartServerPush:output_type -> grpc.telegpb.ServerPushResponse
	3, // 5: grpc.telegpb.TelegramApi.SendMessage:output_type -> grpc.telegpb.SendMessageResponse
	5, // 6: grpc.telegpb.TelegramApi.SendKeyboard:output_type -> grpc.telegpb.SendKeyboardResponse
	7, // 7: grpc.telegpb.TelegramApi.SendReplyInput:output_type -> grpc.telegpb.SendReplyInputResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_proto_teleg_proto_init() }
func file_grpc_proto_teleg_proto_init() {
	if File_grpc_proto_teleg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_proto_teleg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_grpc_proto_teleg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerPushResponse); i {
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
		file_grpc_proto_teleg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageRequest); i {
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
		file_grpc_proto_teleg_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageResponse); i {
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
		file_grpc_proto_teleg_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendKeyboardRequest); i {
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
		file_grpc_proto_teleg_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendKeyboardResponse); i {
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
		file_grpc_proto_teleg_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendReplyInputRequest); i {
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
		file_grpc_proto_teleg_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendReplyInputResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpc_proto_teleg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_proto_teleg_proto_goTypes,
		DependencyIndexes: file_grpc_proto_teleg_proto_depIdxs,
		MessageInfos:      file_grpc_proto_teleg_proto_msgTypes,
	}.Build()
	File_grpc_proto_teleg_proto = out.File
	file_grpc_proto_teleg_proto_rawDesc = nil
	file_grpc_proto_teleg_proto_goTypes = nil
	file_grpc_proto_teleg_proto_depIdxs = nil
}
