// Code generated by protoc-gen-go. DO NOT EDIT.
// source: assistant.proto

package assistantdbpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Message struct {
	MsgID                int64    `protobuf:"varint,1,opt,name=msgID,proto3" json:"msgID,omitempty"`
	Data                 string   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Keys                 []string `protobuf:"bytes,3,rep,name=keys,proto3" json:"keys,omitempty"`
	CreateTime           int64    `protobuf:"varint,4,opt,name=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime           int64    `protobuf:"varint,5,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_assistant_945acc512052abde, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (dst *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(dst, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetMsgID() int64 {
	if m != nil {
		return m.MsgID
	}
	return 0
}

func (m *Message) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *Message) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *Message) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *Message) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

type AssistantData struct {
	MaxMsgID             int64    `protobuf:"varint,1,opt,name=maxMsgID,proto3" json:"maxMsgID,omitempty"`
	Keys                 []string `protobuf:"bytes,2,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssistantData) Reset()         { *m = AssistantData{} }
func (m *AssistantData) String() string { return proto.CompactTextString(m) }
func (*AssistantData) ProtoMessage()    {}
func (*AssistantData) Descriptor() ([]byte, []int) {
	return fileDescriptor_assistant_945acc512052abde, []int{1}
}
func (m *AssistantData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssistantData.Unmarshal(m, b)
}
func (m *AssistantData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssistantData.Marshal(b, m, deterministic)
}
func (dst *AssistantData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssistantData.Merge(dst, src)
}
func (m *AssistantData) XXX_Size() int {
	return xxx_messageInfo_AssistantData.Size(m)
}
func (m *AssistantData) XXX_DiscardUnknown() {
	xxx_messageInfo_AssistantData.DiscardUnknown(m)
}

var xxx_messageInfo_AssistantData proto.InternalMessageInfo

func (m *AssistantData) GetMaxMsgID() int64 {
	if m != nil {
		return m.MaxMsgID
	}
	return 0
}

func (m *AssistantData) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "assistantdbpb.Message")
	proto.RegisterType((*AssistantData)(nil), "assistantdbpb.AssistantData")
}

func init() { proto.RegisterFile("assistant.proto", fileDescriptor_assistant_945acc512052abde) }

var fileDescriptor_assistant_945acc512052abde = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x2c, 0x2e, 0xce,
	0x2c, 0x2e, 0x49, 0xcc, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x85, 0x0b, 0xa4,
	0x24, 0x15, 0x24, 0x29, 0xb5, 0x33, 0x72, 0xb1, 0xfb, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0x0a,
	0x89, 0x70, 0xb1, 0xe6, 0x16, 0xa7, 0x7b, 0xba, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x30, 0x07, 0x41,
	0x38, 0x42, 0x42, 0x5c, 0x2c, 0x29, 0x89, 0x25, 0x89, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41,
	0x60, 0x36, 0x48, 0x2c, 0x3b, 0xb5, 0xb2, 0x58, 0x82, 0x59, 0x81, 0x19, 0x24, 0x06, 0x62, 0x0b,
	0xc9, 0x71, 0x71, 0x25, 0x17, 0xa5, 0x26, 0x96, 0xa4, 0x86, 0x64, 0xe6, 0xa6, 0x4a, 0xb0, 0x80,
	0x8d, 0x40, 0x12, 0x01, 0xc9, 0x97, 0x16, 0xa4, 0xc0, 0xe4, 0x59, 0x21, 0xf2, 0x08, 0x11, 0x25,
	0x7b, 0x2e, 0x5e, 0x47, 0x98, 0xd3, 0x5c, 0x40, 0x96, 0x48, 0x71, 0x71, 0xe4, 0x26, 0x56, 0xf8,
	0x22, 0xb9, 0x08, 0xce, 0x87, 0x3b, 0x80, 0x09, 0xe1, 0x80, 0x24, 0x36, 0xb0, 0x07, 0x8d, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x87, 0x92, 0x02, 0x85, 0xf3, 0x00, 0x00, 0x00,
}