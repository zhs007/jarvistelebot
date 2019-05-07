// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dtdata2.proto

package plugindtdata2pb

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

// GameDayReportCommand - game day report
type GameDayReportCommand struct {
	// env - env
	Env string `protobuf:"bytes,1,opt,name=env,proto3" json:"env,omitempty"`
	// dayTime - is like 2019-05-01
	DayTime string `protobuf:"bytes,2,opt,name=dayTime,proto3" json:"dayTime,omitempty"`
	// currency - is like USD
	Currency string `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
	// scaleMoney - is like 10000
	ScaleMoney           int32    `protobuf:"varint,4,opt,name=scaleMoney,proto3" json:"scaleMoney,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GameDayReportCommand) Reset()         { *m = GameDayReportCommand{} }
func (m *GameDayReportCommand) String() string { return proto.CompactTextString(m) }
func (*GameDayReportCommand) ProtoMessage()    {}
func (*GameDayReportCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_dtdata2_1eed8d63cc5c6fd0, []int{0}
}
func (m *GameDayReportCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameDayReportCommand.Unmarshal(m, b)
}
func (m *GameDayReportCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameDayReportCommand.Marshal(b, m, deterministic)
}
func (dst *GameDayReportCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameDayReportCommand.Merge(dst, src)
}
func (m *GameDayReportCommand) XXX_Size() int {
	return xxx_messageInfo_GameDayReportCommand.Size(m)
}
func (m *GameDayReportCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_GameDayReportCommand.DiscardUnknown(m)
}

var xxx_messageInfo_GameDayReportCommand proto.InternalMessageInfo

func (m *GameDayReportCommand) GetEnv() string {
	if m != nil {
		return m.Env
	}
	return ""
}

func (m *GameDayReportCommand) GetDayTime() string {
	if m != nil {
		return m.DayTime
	}
	return ""
}

func (m *GameDayReportCommand) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *GameDayReportCommand) GetScaleMoney() int32 {
	if m != nil {
		return m.ScaleMoney
	}
	return 0
}

func init() {
	proto.RegisterType((*GameDayReportCommand)(nil), "plugindtdata2pb.GameDayReportCommand")
}

func init() { proto.RegisterFile("dtdata2.proto", fileDescriptor_dtdata2_1eed8d63cc5c6fd0) }

var fileDescriptor_dtdata2_1eed8d63cc5c6fd0 = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x29, 0x49, 0x49,
	0x2c, 0x49, 0x34, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2f, 0xc8, 0x29, 0x4d, 0xcf,
	0xcc, 0x83, 0x0a, 0x16, 0x24, 0x29, 0xd5, 0x71, 0x89, 0xb8, 0x27, 0xe6, 0xa6, 0xba, 0x24, 0x56,
	0x06, 0xa5, 0x16, 0xe4, 0x17, 0x95, 0x38, 0xe7, 0xe7, 0xe6, 0x26, 0xe6, 0xa5, 0x08, 0x09, 0x70,
	0x31, 0xa7, 0xe6, 0x95, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0x98, 0x42, 0x12, 0x5c,
	0xec, 0x29, 0x89, 0x95, 0x21, 0x99, 0xb9, 0xa9, 0x12, 0x4c, 0x60, 0x51, 0x18, 0x57, 0x48, 0x8a,
	0x8b, 0x23, 0xb9, 0xb4, 0xa8, 0x28, 0x35, 0x2f, 0xb9, 0x52, 0x82, 0x19, 0x2c, 0x05, 0xe7, 0x0b,
	0xc9, 0x71, 0x71, 0x15, 0x27, 0x27, 0xe6, 0xa4, 0xfa, 0xe6, 0xe7, 0xa5, 0x56, 0x4a, 0xb0, 0x28,
	0x30, 0x6a, 0xb0, 0x06, 0x21, 0x89, 0x24, 0xb1, 0x81, 0xdd, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0x4b, 0x02, 0xd3, 0xb7, 0xa8, 0x00, 0x00, 0x00,
}
