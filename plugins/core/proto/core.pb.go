// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core.proto

package plugincorepb

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

// UserCommand - user command
type UserCommand struct {
	// userID - User's unique identifier
	UserID string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	// userName - Username
	UserName             string   `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserCommand) Reset()         { *m = UserCommand{} }
func (m *UserCommand) String() string { return proto.CompactTextString(m) }
func (*UserCommand) ProtoMessage()    {}
func (*UserCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_core_a41b49cb79ee2a6d, []int{0}
}
func (m *UserCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserCommand.Unmarshal(m, b)
}
func (m *UserCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserCommand.Marshal(b, m, deterministic)
}
func (dst *UserCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserCommand.Merge(dst, src)
}
func (m *UserCommand) XXX_Size() int {
	return xxx_messageInfo_UserCommand.Size(m)
}
func (m *UserCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_UserCommand.DiscardUnknown(m)
}

var xxx_messageInfo_UserCommand proto.InternalMessageInfo

func (m *UserCommand) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *UserCommand) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

// UsersCommand - users command
type UsersCommand struct {
	// nums - user nums
	Nums                 int32    `protobuf:"varint,1,opt,name=nums,proto3" json:"nums,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UsersCommand) Reset()         { *m = UsersCommand{} }
func (m *UsersCommand) String() string { return proto.CompactTextString(m) }
func (*UsersCommand) ProtoMessage()    {}
func (*UsersCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_core_a41b49cb79ee2a6d, []int{1}
}
func (m *UsersCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsersCommand.Unmarshal(m, b)
}
func (m *UsersCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsersCommand.Marshal(b, m, deterministic)
}
func (dst *UsersCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsersCommand.Merge(dst, src)
}
func (m *UsersCommand) XXX_Size() int {
	return xxx_messageInfo_UsersCommand.Size(m)
}
func (m *UsersCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_UsersCommand.DiscardUnknown(m)
}

var xxx_messageInfo_UsersCommand proto.InternalMessageInfo

func (m *UsersCommand) GetNums() int32 {
	if m != nil {
		return m.Nums
	}
	return 0
}

func init() {
	proto.RegisterType((*UserCommand)(nil), "plugincorepb.UserCommand")
	proto.RegisterType((*UsersCommand)(nil), "plugincorepb.UsersCommand")
}

func init() { proto.RegisterFile("core.proto", fileDescriptor_core_a41b49cb79ee2a6d) }

var fileDescriptor_core_a41b49cb79ee2a6d = []byte{
	// 124 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xce, 0x2f, 0x4a,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x29, 0xc8, 0x29, 0x4d, 0xcf, 0xcc, 0x03, 0x89,
	0x14, 0x24, 0x29, 0x39, 0x72, 0x71, 0x87, 0x16, 0xa7, 0x16, 0x39, 0xe7, 0xe7, 0xe6, 0x26, 0xe6,
	0xa5, 0x08, 0x89, 0x71, 0xb1, 0x95, 0x16, 0xa7, 0x16, 0x79, 0xba, 0x48, 0x30, 0x2a, 0x30, 0x6a,
	0x70, 0x06, 0x41, 0x79, 0x42, 0x52, 0x5c, 0x1c, 0x20, 0x96, 0x5f, 0x62, 0x6e, 0xaa, 0x04, 0x13,
	0x58, 0x06, 0xce, 0x57, 0x52, 0xe2, 0xe2, 0x01, 0x19, 0x51, 0x0c, 0x33, 0x43, 0x88, 0x8b, 0x25,
	0xaf, 0x34, 0xb7, 0x18, 0x6c, 0x02, 0x6b, 0x10, 0x98, 0x9d, 0xc4, 0x06, 0xb6, 0xdb, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0x0f, 0x1a, 0xa2, 0xe6, 0x89, 0x00, 0x00, 0x00,
}
