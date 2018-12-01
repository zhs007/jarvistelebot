// Code generated by protoc-gen-go. DO NOT EDIT.
// source: userscript.proto

package pluginuserscriptpb

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

// RunScriptCommand - run command
type RunScriptCommand struct {
	// scriptName - script name
	ScriptName           string   `protobuf:"bytes,1,opt,name=scriptName,proto3" json:"scriptName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunScriptCommand) Reset()         { *m = RunScriptCommand{} }
func (m *RunScriptCommand) String() string { return proto.CompactTextString(m) }
func (*RunScriptCommand) ProtoMessage()    {}
func (*RunScriptCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_userscript_476aba5f53152a6b, []int{0}
}
func (m *RunScriptCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunScriptCommand.Unmarshal(m, b)
}
func (m *RunScriptCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunScriptCommand.Marshal(b, m, deterministic)
}
func (dst *RunScriptCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunScriptCommand.Merge(dst, src)
}
func (m *RunScriptCommand) XXX_Size() int {
	return xxx_messageInfo_RunScriptCommand.Size(m)
}
func (m *RunScriptCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_RunScriptCommand.DiscardUnknown(m)
}

var xxx_messageInfo_RunScriptCommand proto.InternalMessageInfo

func (m *RunScriptCommand) GetScriptName() string {
	if m != nil {
		return m.ScriptName
	}
	return ""
}

func init() {
	proto.RegisterType((*RunScriptCommand)(nil), "pluginuserscriptpb.RunScriptCommand")
}

func init() { proto.RegisterFile("userscript.proto", fileDescriptor_userscript_476aba5f53152a6b) }

var fileDescriptor_userscript_476aba5f53152a6b = []byte{
	// 95 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x2d, 0x4e, 0x2d,
	0x2a, 0x4e, 0x2e, 0xca, 0x2c, 0x28, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x2a, 0xc8,
	0x29, 0x4d, 0xcf, 0xcc, 0x43, 0x88, 0x17, 0x24, 0x29, 0x19, 0x71, 0x09, 0x04, 0x95, 0xe6, 0x05,
	0x83, 0xb9, 0xce, 0xf9, 0xb9, 0xb9, 0x89, 0x79, 0x29, 0x42, 0x72, 0x5c, 0x5c, 0x10, 0x79, 0xbf,
	0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x24, 0x91, 0x24, 0x36, 0xb0, 0x71,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe9, 0x2f, 0xe1, 0xde, 0x62, 0x00, 0x00, 0x00,
}
