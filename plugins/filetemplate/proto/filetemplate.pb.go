// Code generated by protoc-gen-go. DO NOT EDIT.
// source: filetemplate.proto

package pluginfiletemplatepb

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

// FileTemplateCommand - file template command
type FileTemplateCommand struct {
	// fileTemplateName - file template name
	FileTemplateName     string   `protobuf:"bytes,1,opt,name=fileTemplateName,proto3" json:"fileTemplateName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileTemplateCommand) Reset()         { *m = FileTemplateCommand{} }
func (m *FileTemplateCommand) String() string { return proto.CompactTextString(m) }
func (*FileTemplateCommand) ProtoMessage()    {}
func (*FileTemplateCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_filetemplate_cbbacf01edb0b519, []int{0}
}
func (m *FileTemplateCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileTemplateCommand.Unmarshal(m, b)
}
func (m *FileTemplateCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileTemplateCommand.Marshal(b, m, deterministic)
}
func (dst *FileTemplateCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileTemplateCommand.Merge(dst, src)
}
func (m *FileTemplateCommand) XXX_Size() int {
	return xxx_messageInfo_FileTemplateCommand.Size(m)
}
func (m *FileTemplateCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_FileTemplateCommand.DiscardUnknown(m)
}

var xxx_messageInfo_FileTemplateCommand proto.InternalMessageInfo

func (m *FileTemplateCommand) GetFileTemplateName() string {
	if m != nil {
		return m.FileTemplateName
	}
	return ""
}

func init() {
	proto.RegisterType((*FileTemplateCommand)(nil), "pluginfiletemplatepb.FileTemplateCommand")
}

func init() { proto.RegisterFile("filetemplate.proto", fileDescriptor_filetemplate_cbbacf01edb0b519) }

var fileDescriptor_filetemplate_cbbacf01edb0b519 = []byte{
	// 99 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0xcb, 0xcc, 0x49,
	0x2d, 0x49, 0xcd, 0x2d, 0xc8, 0x49, 0x2c, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12,
	0x29, 0xc8, 0x29, 0x4d, 0xcf, 0xcc, 0x43, 0x96, 0x29, 0x48, 0x52, 0x72, 0xe4, 0x12, 0x76, 0xcb,
	0xcc, 0x49, 0x0d, 0x81, 0x8a, 0x38, 0xe7, 0xe7, 0xe6, 0x26, 0xe6, 0xa5, 0x08, 0x69, 0x71, 0x09,
	0xa4, 0x21, 0x09, 0xfb, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x61, 0x88,
	0x27, 0xb1, 0x81, 0xcd, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x41, 0xbc, 0x79, 0x75,
	0x00, 0x00, 0x00,
}
