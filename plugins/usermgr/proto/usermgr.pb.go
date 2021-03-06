// Code generated by protoc-gen-go. DO NOT EDIT.
// source: usermgr.proto

package pluginusermgrpb

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

// UpdScriptCommand - updscript command
type UpdScriptCommand struct {
	// userID - userID
	UserID string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	// userName - userName
	UserName string `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	// scriptName - script name
	ScriptName string `protobuf:"bytes,3,opt,name=scriptName,proto3" json:"scriptName,omitempty"`
	// jarvisNodeName - jarvis node name
	JarvisNodeName       string   `protobuf:"bytes,4,opt,name=jarvisNodeName,proto3" json:"jarvisNodeName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdScriptCommand) Reset()         { *m = UpdScriptCommand{} }
func (m *UpdScriptCommand) String() string { return proto.CompactTextString(m) }
func (*UpdScriptCommand) ProtoMessage()    {}
func (*UpdScriptCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_6dfcc81643e16c43, []int{0}
}
func (m *UpdScriptCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdScriptCommand.Unmarshal(m, b)
}
func (m *UpdScriptCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdScriptCommand.Marshal(b, m, deterministic)
}
func (dst *UpdScriptCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdScriptCommand.Merge(dst, src)
}
func (m *UpdScriptCommand) XXX_Size() int {
	return xxx_messageInfo_UpdScriptCommand.Size(m)
}
func (m *UpdScriptCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdScriptCommand.DiscardUnknown(m)
}

var xxx_messageInfo_UpdScriptCommand proto.InternalMessageInfo

func (m *UpdScriptCommand) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *UpdScriptCommand) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UpdScriptCommand) GetScriptName() string {
	if m != nil {
		return m.ScriptName
	}
	return ""
}

func (m *UpdScriptCommand) GetJarvisNodeName() string {
	if m != nil {
		return m.JarvisNodeName
	}
	return ""
}

// UserScriptsCommand - userscripts command
type UserScriptsCommand struct {
	// userID - userID
	UserID string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	// userName - userName
	UserName string `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	// jarvisNodeName - jarvis node name
	JarvisNodeName       string   `protobuf:"bytes,3,opt,name=jarvisNodeName,proto3" json:"jarvisNodeName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserScriptsCommand) Reset()         { *m = UserScriptsCommand{} }
func (m *UserScriptsCommand) String() string { return proto.CompactTextString(m) }
func (*UserScriptsCommand) ProtoMessage()    {}
func (*UserScriptsCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_6dfcc81643e16c43, []int{1}
}
func (m *UserScriptsCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserScriptsCommand.Unmarshal(m, b)
}
func (m *UserScriptsCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserScriptsCommand.Marshal(b, m, deterministic)
}
func (dst *UserScriptsCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserScriptsCommand.Merge(dst, src)
}
func (m *UserScriptsCommand) XXX_Size() int {
	return xxx_messageInfo_UserScriptsCommand.Size(m)
}
func (m *UserScriptsCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_UserScriptsCommand.DiscardUnknown(m)
}

var xxx_messageInfo_UserScriptsCommand proto.InternalMessageInfo

func (m *UserScriptsCommand) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *UserScriptsCommand) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UserScriptsCommand) GetJarvisNodeName() string {
	if m != nil {
		return m.JarvisNodeName
	}
	return ""
}

// RemoveScriptCommand - rmscript command
type RemoveScriptCommand struct {
	// userID - userID
	UserID string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	// userName - userName
	UserName string `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	// scriptName - script name
	ScriptName           string   `protobuf:"bytes,3,opt,name=scriptName,proto3" json:"scriptName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveScriptCommand) Reset()         { *m = RemoveScriptCommand{} }
func (m *RemoveScriptCommand) String() string { return proto.CompactTextString(m) }
func (*RemoveScriptCommand) ProtoMessage()    {}
func (*RemoveScriptCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_6dfcc81643e16c43, []int{2}
}
func (m *RemoveScriptCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveScriptCommand.Unmarshal(m, b)
}
func (m *RemoveScriptCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveScriptCommand.Marshal(b, m, deterministic)
}
func (dst *RemoveScriptCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveScriptCommand.Merge(dst, src)
}
func (m *RemoveScriptCommand) XXX_Size() int {
	return xxx_messageInfo_RemoveScriptCommand.Size(m)
}
func (m *RemoveScriptCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveScriptCommand.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveScriptCommand proto.InternalMessageInfo

func (m *RemoveScriptCommand) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *RemoveScriptCommand) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *RemoveScriptCommand) GetScriptName() string {
	if m != nil {
		return m.ScriptName
	}
	return ""
}

// ShowScriptCommand - showscript command
type ShowScriptCommand struct {
	// userID - userID
	UserID string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	// userName - userName
	UserName string `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	// scriptName - script name
	ScriptName           string   `protobuf:"bytes,3,opt,name=scriptName,proto3" json:"scriptName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShowScriptCommand) Reset()         { *m = ShowScriptCommand{} }
func (m *ShowScriptCommand) String() string { return proto.CompactTextString(m) }
func (*ShowScriptCommand) ProtoMessage()    {}
func (*ShowScriptCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_6dfcc81643e16c43, []int{3}
}
func (m *ShowScriptCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShowScriptCommand.Unmarshal(m, b)
}
func (m *ShowScriptCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShowScriptCommand.Marshal(b, m, deterministic)
}
func (dst *ShowScriptCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShowScriptCommand.Merge(dst, src)
}
func (m *ShowScriptCommand) XXX_Size() int {
	return xxx_messageInfo_ShowScriptCommand.Size(m)
}
func (m *ShowScriptCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_ShowScriptCommand.DiscardUnknown(m)
}

var xxx_messageInfo_ShowScriptCommand proto.InternalMessageInfo

func (m *ShowScriptCommand) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *ShowScriptCommand) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *ShowScriptCommand) GetScriptName() string {
	if m != nil {
		return m.ScriptName
	}
	return ""
}

// UpdFileTemplateCommand - updfiletemplate command
type UpdFileTemplateCommand struct {
	// userID - userID
	UserID string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	// userName - userName
	UserName string `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	// fileTemplateName - file template name
	FileTemplateName string `protobuf:"bytes,3,opt,name=fileTemplateName,proto3" json:"fileTemplateName,omitempty"`
	// jarvisNodeName - jarvis node name
	JarvisNodeName string `protobuf:"bytes,4,opt,name=jarvisNodeName,proto3" json:"jarvisNodeName,omitempty"`
	// fullPath - full path
	FullPath string `protobuf:"bytes,5,opt,name=fullPath,proto3" json:"fullPath,omitempty"`
	// subfilesPath - subfiles path
	SubfilesPath         string   `protobuf:"bytes,6,opt,name=subfilesPath,proto3" json:"subfilesPath,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdFileTemplateCommand) Reset()         { *m = UpdFileTemplateCommand{} }
func (m *UpdFileTemplateCommand) String() string { return proto.CompactTextString(m) }
func (*UpdFileTemplateCommand) ProtoMessage()    {}
func (*UpdFileTemplateCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_6dfcc81643e16c43, []int{4}
}
func (m *UpdFileTemplateCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdFileTemplateCommand.Unmarshal(m, b)
}
func (m *UpdFileTemplateCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdFileTemplateCommand.Marshal(b, m, deterministic)
}
func (dst *UpdFileTemplateCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdFileTemplateCommand.Merge(dst, src)
}
func (m *UpdFileTemplateCommand) XXX_Size() int {
	return xxx_messageInfo_UpdFileTemplateCommand.Size(m)
}
func (m *UpdFileTemplateCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdFileTemplateCommand.DiscardUnknown(m)
}

var xxx_messageInfo_UpdFileTemplateCommand proto.InternalMessageInfo

func (m *UpdFileTemplateCommand) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *UpdFileTemplateCommand) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UpdFileTemplateCommand) GetFileTemplateName() string {
	if m != nil {
		return m.FileTemplateName
	}
	return ""
}

func (m *UpdFileTemplateCommand) GetJarvisNodeName() string {
	if m != nil {
		return m.JarvisNodeName
	}
	return ""
}

func (m *UpdFileTemplateCommand) GetFullPath() string {
	if m != nil {
		return m.FullPath
	}
	return ""
}

func (m *UpdFileTemplateCommand) GetSubfilesPath() string {
	if m != nil {
		return m.SubfilesPath
	}
	return ""
}

// FileTemplatesCommand - filetemplates command
type FileTemplatesCommand struct {
	// userID - userID
	UserID string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	// userName - userName
	UserName string `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	// jarvisNodeName - jarvis node name
	JarvisNodeName       string   `protobuf:"bytes,3,opt,name=jarvisNodeName,proto3" json:"jarvisNodeName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileTemplatesCommand) Reset()         { *m = FileTemplatesCommand{} }
func (m *FileTemplatesCommand) String() string { return proto.CompactTextString(m) }
func (*FileTemplatesCommand) ProtoMessage()    {}
func (*FileTemplatesCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_6dfcc81643e16c43, []int{5}
}
func (m *FileTemplatesCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileTemplatesCommand.Unmarshal(m, b)
}
func (m *FileTemplatesCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileTemplatesCommand.Marshal(b, m, deterministic)
}
func (dst *FileTemplatesCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileTemplatesCommand.Merge(dst, src)
}
func (m *FileTemplatesCommand) XXX_Size() int {
	return xxx_messageInfo_FileTemplatesCommand.Size(m)
}
func (m *FileTemplatesCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_FileTemplatesCommand.DiscardUnknown(m)
}

var xxx_messageInfo_FileTemplatesCommand proto.InternalMessageInfo

func (m *FileTemplatesCommand) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *FileTemplatesCommand) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *FileTemplatesCommand) GetJarvisNodeName() string {
	if m != nil {
		return m.JarvisNodeName
	}
	return ""
}

// RemoveFileTemplateCommand - rmfiletemplate command
type RemoveFileTemplateCommand struct {
	// userID - userID
	UserID string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	// userName - userName
	UserName string `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	// fileTemplateName - filetemplate name
	FileTemplateName     string   `protobuf:"bytes,3,opt,name=fileTemplateName,proto3" json:"fileTemplateName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveFileTemplateCommand) Reset()         { *m = RemoveFileTemplateCommand{} }
func (m *RemoveFileTemplateCommand) String() string { return proto.CompactTextString(m) }
func (*RemoveFileTemplateCommand) ProtoMessage()    {}
func (*RemoveFileTemplateCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_6dfcc81643e16c43, []int{6}
}
func (m *RemoveFileTemplateCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveFileTemplateCommand.Unmarshal(m, b)
}
func (m *RemoveFileTemplateCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveFileTemplateCommand.Marshal(b, m, deterministic)
}
func (dst *RemoveFileTemplateCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveFileTemplateCommand.Merge(dst, src)
}
func (m *RemoveFileTemplateCommand) XXX_Size() int {
	return xxx_messageInfo_RemoveFileTemplateCommand.Size(m)
}
func (m *RemoveFileTemplateCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveFileTemplateCommand.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveFileTemplateCommand proto.InternalMessageInfo

func (m *RemoveFileTemplateCommand) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *RemoveFileTemplateCommand) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *RemoveFileTemplateCommand) GetFileTemplateName() string {
	if m != nil {
		return m.FileTemplateName
	}
	return ""
}

// ShowFileTemplateCommand - showfiletemplate command
type ShowFileTemplateCommand struct {
	// userID - userID
	UserID string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	// userName - userName
	UserName string `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	// scriptName - script name
	FileTemplateName     string   `protobuf:"bytes,3,opt,name=fileTemplateName,proto3" json:"fileTemplateName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShowFileTemplateCommand) Reset()         { *m = ShowFileTemplateCommand{} }
func (m *ShowFileTemplateCommand) String() string { return proto.CompactTextString(m) }
func (*ShowFileTemplateCommand) ProtoMessage()    {}
func (*ShowFileTemplateCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_6dfcc81643e16c43, []int{7}
}
func (m *ShowFileTemplateCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShowFileTemplateCommand.Unmarshal(m, b)
}
func (m *ShowFileTemplateCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShowFileTemplateCommand.Marshal(b, m, deterministic)
}
func (dst *ShowFileTemplateCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShowFileTemplateCommand.Merge(dst, src)
}
func (m *ShowFileTemplateCommand) XXX_Size() int {
	return xxx_messageInfo_ShowFileTemplateCommand.Size(m)
}
func (m *ShowFileTemplateCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_ShowFileTemplateCommand.DiscardUnknown(m)
}

var xxx_messageInfo_ShowFileTemplateCommand proto.InternalMessageInfo

func (m *ShowFileTemplateCommand) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *ShowFileTemplateCommand) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *ShowFileTemplateCommand) GetFileTemplateName() string {
	if m != nil {
		return m.FileTemplateName
	}
	return ""
}

// MyScriptsCommand - myscripts command
type MyScriptsCommand struct {
	// jarvisNodeName - jarvis node name
	JarvisNodeName       string   `protobuf:"bytes,1,opt,name=jarvisNodeName,proto3" json:"jarvisNodeName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MyScriptsCommand) Reset()         { *m = MyScriptsCommand{} }
func (m *MyScriptsCommand) String() string { return proto.CompactTextString(m) }
func (*MyScriptsCommand) ProtoMessage()    {}
func (*MyScriptsCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_6dfcc81643e16c43, []int{8}
}
func (m *MyScriptsCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MyScriptsCommand.Unmarshal(m, b)
}
func (m *MyScriptsCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MyScriptsCommand.Marshal(b, m, deterministic)
}
func (dst *MyScriptsCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MyScriptsCommand.Merge(dst, src)
}
func (m *MyScriptsCommand) XXX_Size() int {
	return xxx_messageInfo_MyScriptsCommand.Size(m)
}
func (m *MyScriptsCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_MyScriptsCommand.DiscardUnknown(m)
}

var xxx_messageInfo_MyScriptsCommand proto.InternalMessageInfo

func (m *MyScriptsCommand) GetJarvisNodeName() string {
	if m != nil {
		return m.JarvisNodeName
	}
	return ""
}

// MyFileTemplatesCommand - myfiletemplates command
type MyFileTemplatesCommand struct {
	// jarvisNodeName - jarvis node name
	JarvisNodeName       string   `protobuf:"bytes,1,opt,name=jarvisNodeName,proto3" json:"jarvisNodeName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MyFileTemplatesCommand) Reset()         { *m = MyFileTemplatesCommand{} }
func (m *MyFileTemplatesCommand) String() string { return proto.CompactTextString(m) }
func (*MyFileTemplatesCommand) ProtoMessage()    {}
func (*MyFileTemplatesCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_6dfcc81643e16c43, []int{9}
}
func (m *MyFileTemplatesCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MyFileTemplatesCommand.Unmarshal(m, b)
}
func (m *MyFileTemplatesCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MyFileTemplatesCommand.Marshal(b, m, deterministic)
}
func (dst *MyFileTemplatesCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MyFileTemplatesCommand.Merge(dst, src)
}
func (m *MyFileTemplatesCommand) XXX_Size() int {
	return xxx_messageInfo_MyFileTemplatesCommand.Size(m)
}
func (m *MyFileTemplatesCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_MyFileTemplatesCommand.DiscardUnknown(m)
}

var xxx_messageInfo_MyFileTemplatesCommand proto.InternalMessageInfo

func (m *MyFileTemplatesCommand) GetJarvisNodeName() string {
	if m != nil {
		return m.JarvisNodeName
	}
	return ""
}

// ExpScriptsCommand - expscripts command
type ExpScriptsCommand struct {
	// userID - userID
	UserID string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	// userName - userName
	UserName             string   `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExpScriptsCommand) Reset()         { *m = ExpScriptsCommand{} }
func (m *ExpScriptsCommand) String() string { return proto.CompactTextString(m) }
func (*ExpScriptsCommand) ProtoMessage()    {}
func (*ExpScriptsCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_6dfcc81643e16c43, []int{10}
}
func (m *ExpScriptsCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExpScriptsCommand.Unmarshal(m, b)
}
func (m *ExpScriptsCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExpScriptsCommand.Marshal(b, m, deterministic)
}
func (dst *ExpScriptsCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExpScriptsCommand.Merge(dst, src)
}
func (m *ExpScriptsCommand) XXX_Size() int {
	return xxx_messageInfo_ExpScriptsCommand.Size(m)
}
func (m *ExpScriptsCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_ExpScriptsCommand.DiscardUnknown(m)
}

var xxx_messageInfo_ExpScriptsCommand proto.InternalMessageInfo

func (m *ExpScriptsCommand) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *ExpScriptsCommand) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

// ExpFileTemplatesCommand - expfiletemplates command
type ExpFileTemplatesCommand struct {
	// userID - userID
	UserID string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	// userName - userName
	UserName             string   `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExpFileTemplatesCommand) Reset()         { *m = ExpFileTemplatesCommand{} }
func (m *ExpFileTemplatesCommand) String() string { return proto.CompactTextString(m) }
func (*ExpFileTemplatesCommand) ProtoMessage()    {}
func (*ExpFileTemplatesCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_6dfcc81643e16c43, []int{11}
}
func (m *ExpFileTemplatesCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExpFileTemplatesCommand.Unmarshal(m, b)
}
func (m *ExpFileTemplatesCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExpFileTemplatesCommand.Marshal(b, m, deterministic)
}
func (dst *ExpFileTemplatesCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExpFileTemplatesCommand.Merge(dst, src)
}
func (m *ExpFileTemplatesCommand) XXX_Size() int {
	return xxx_messageInfo_ExpFileTemplatesCommand.Size(m)
}
func (m *ExpFileTemplatesCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_ExpFileTemplatesCommand.DiscardUnknown(m)
}

var xxx_messageInfo_ExpFileTemplatesCommand proto.InternalMessageInfo

func (m *ExpFileTemplatesCommand) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *ExpFileTemplatesCommand) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

// ImpScriptsCommand - impscripts command
type ImpScriptsCommand struct {
	// userID - userID
	UserID string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	// userName - userName
	UserName             string   `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImpScriptsCommand) Reset()         { *m = ImpScriptsCommand{} }
func (m *ImpScriptsCommand) String() string { return proto.CompactTextString(m) }
func (*ImpScriptsCommand) ProtoMessage()    {}
func (*ImpScriptsCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_6dfcc81643e16c43, []int{12}
}
func (m *ImpScriptsCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImpScriptsCommand.Unmarshal(m, b)
}
func (m *ImpScriptsCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImpScriptsCommand.Marshal(b, m, deterministic)
}
func (dst *ImpScriptsCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImpScriptsCommand.Merge(dst, src)
}
func (m *ImpScriptsCommand) XXX_Size() int {
	return xxx_messageInfo_ImpScriptsCommand.Size(m)
}
func (m *ImpScriptsCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_ImpScriptsCommand.DiscardUnknown(m)
}

var xxx_messageInfo_ImpScriptsCommand proto.InternalMessageInfo

func (m *ImpScriptsCommand) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *ImpScriptsCommand) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

// ImpFileTemplatesCommand - impfiletemplates command
type ImpFileTemplatesCommand struct {
	// userID - userID
	UserID string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	// userName - userName
	UserName             string   `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImpFileTemplatesCommand) Reset()         { *m = ImpFileTemplatesCommand{} }
func (m *ImpFileTemplatesCommand) String() string { return proto.CompactTextString(m) }
func (*ImpFileTemplatesCommand) ProtoMessage()    {}
func (*ImpFileTemplatesCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_usermgr_6dfcc81643e16c43, []int{13}
}
func (m *ImpFileTemplatesCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImpFileTemplatesCommand.Unmarshal(m, b)
}
func (m *ImpFileTemplatesCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImpFileTemplatesCommand.Marshal(b, m, deterministic)
}
func (dst *ImpFileTemplatesCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImpFileTemplatesCommand.Merge(dst, src)
}
func (m *ImpFileTemplatesCommand) XXX_Size() int {
	return xxx_messageInfo_ImpFileTemplatesCommand.Size(m)
}
func (m *ImpFileTemplatesCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_ImpFileTemplatesCommand.DiscardUnknown(m)
}

var xxx_messageInfo_ImpFileTemplatesCommand proto.InternalMessageInfo

func (m *ImpFileTemplatesCommand) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *ImpFileTemplatesCommand) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func init() {
	proto.RegisterType((*UpdScriptCommand)(nil), "pluginusermgrpb.UpdScriptCommand")
	proto.RegisterType((*UserScriptsCommand)(nil), "pluginusermgrpb.UserScriptsCommand")
	proto.RegisterType((*RemoveScriptCommand)(nil), "pluginusermgrpb.RemoveScriptCommand")
	proto.RegisterType((*ShowScriptCommand)(nil), "pluginusermgrpb.ShowScriptCommand")
	proto.RegisterType((*UpdFileTemplateCommand)(nil), "pluginusermgrpb.UpdFileTemplateCommand")
	proto.RegisterType((*FileTemplatesCommand)(nil), "pluginusermgrpb.FileTemplatesCommand")
	proto.RegisterType((*RemoveFileTemplateCommand)(nil), "pluginusermgrpb.RemoveFileTemplateCommand")
	proto.RegisterType((*ShowFileTemplateCommand)(nil), "pluginusermgrpb.ShowFileTemplateCommand")
	proto.RegisterType((*MyScriptsCommand)(nil), "pluginusermgrpb.MyScriptsCommand")
	proto.RegisterType((*MyFileTemplatesCommand)(nil), "pluginusermgrpb.MyFileTemplatesCommand")
	proto.RegisterType((*ExpScriptsCommand)(nil), "pluginusermgrpb.ExpScriptsCommand")
	proto.RegisterType((*ExpFileTemplatesCommand)(nil), "pluginusermgrpb.ExpFileTemplatesCommand")
	proto.RegisterType((*ImpScriptsCommand)(nil), "pluginusermgrpb.ImpScriptsCommand")
	proto.RegisterType((*ImpFileTemplatesCommand)(nil), "pluginusermgrpb.ImpFileTemplatesCommand")
}

func init() { proto.RegisterFile("usermgr.proto", fileDescriptor_usermgr_6dfcc81643e16c43) }

var fileDescriptor_usermgr_6dfcc81643e16c43 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x2d, 0x4e, 0x2d,
	0xca, 0x4d, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2f, 0xc8, 0x29, 0x4d, 0xcf,
	0xcc, 0x83, 0x0a, 0x16, 0x24, 0x29, 0xf5, 0x31, 0x72, 0x09, 0x84, 0x16, 0xa4, 0x04, 0x27, 0x17,
	0x65, 0x16, 0x94, 0x38, 0xe7, 0xe7, 0xe6, 0x26, 0xe6, 0xa5, 0x08, 0x89, 0x71, 0xb1, 0x81, 0x54,
	0x78, 0xba, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x79, 0x42, 0x52, 0x5c, 0x1c, 0x20,
	0x96, 0x5f, 0x62, 0x6e, 0xaa, 0x04, 0x13, 0x58, 0x06, 0xce, 0x17, 0x92, 0xe3, 0xe2, 0x2a, 0x06,
	0x1b, 0x02, 0x96, 0x65, 0x06, 0xcb, 0x22, 0x89, 0x08, 0xa9, 0x71, 0xf1, 0x65, 0x25, 0x16, 0x95,
	0x65, 0x16, 0xfb, 0xe5, 0xa7, 0xa4, 0x82, 0xd5, 0xb0, 0x80, 0xd5, 0xa0, 0x89, 0x2a, 0x15, 0x70,
	0x09, 0x85, 0x16, 0xa7, 0x16, 0x41, 0x1c, 0x54, 0x4c, 0x89, 0x8b, 0x30, 0x6d, 0x64, 0xc6, 0x6a,
	0x63, 0x26, 0x97, 0x70, 0x50, 0x6a, 0x6e, 0x7e, 0x59, 0x2a, 0xcd, 0x03, 0x41, 0x29, 0x9d, 0x4b,
	0x30, 0x38, 0x23, 0xbf, 0x9c, 0xf6, 0x16, 0x3d, 0x60, 0xe4, 0x12, 0x0b, 0x2d, 0x48, 0x71, 0xcb,
	0xcc, 0x49, 0x0d, 0x49, 0xcd, 0x2d, 0xc8, 0x49, 0x2c, 0x49, 0xa5, 0xc4, 0x3a, 0x2d, 0x2e, 0x81,
	0x34, 0x24, 0xa3, 0x90, 0x2c, 0xc5, 0x10, 0x27, 0x36, 0xa2, 0x41, 0xf6, 0xa5, 0x95, 0xe6, 0xe4,
	0x04, 0x24, 0x96, 0x64, 0x48, 0xb0, 0x42, 0xec, 0x83, 0xf1, 0x85, 0x94, 0xb8, 0x78, 0x8a, 0x4b,
	0x93, 0x40, 0x46, 0x17, 0x83, 0xe5, 0xd9, 0xc0, 0xf2, 0x28, 0x62, 0x4a, 0x45, 0x5c, 0x22, 0xc8,
	0xde, 0xa3, 0x4b, 0x52, 0xa9, 0xe6, 0x92, 0x84, 0x24, 0x95, 0x01, 0x08, 0x58, 0xa5, 0x4a, 0x2e,
	0x71, 0x50, 0xe2, 0x19, 0x08, 0xab, 0xad, 0xb8, 0x04, 0x7c, 0x2b, 0xd1, 0xb2, 0x24, 0x66, 0x98,
	0x31, 0x62, 0x0d, 0x33, 0x07, 0x2e, 0x31, 0xdf, 0x4a, 0xac, 0x31, 0x45, 0xac, 0x09, 0xee, 0x5c,
	0x82, 0xae, 0x15, 0x05, 0x94, 0x97, 0x08, 0x4a, 0xbe, 0x5c, 0xe2, 0xae, 0x15, 0x05, 0xd4, 0x4a,
	0x35, 0x20, 0x77, 0x79, 0xe6, 0x52, 0xc9, 0x5d, 0x9e, 0xb9, 0x54, 0x73, 0x57, 0x12, 0x1b, 0xb8,
	0xac, 0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x75, 0x0d, 0xb6, 0xfc, 0x05, 0x00, 0x00,
}
