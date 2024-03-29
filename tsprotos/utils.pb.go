// Code generated by protoc-gen-go. DO NOT EDIT.
// source: utils.proto

package tsprotos

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PBError struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PBError) Reset()         { *m = PBError{} }
func (m *PBError) String() string { return proto.CompactTextString(m) }
func (*PBError) ProtoMessage()    {}
func (*PBError) Descriptor() ([]byte, []int) {
	return fileDescriptor_c91c651f4717a5f2, []int{0}
}

func (m *PBError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PBError.Unmarshal(m, b)
}
func (m *PBError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PBError.Marshal(b, m, deterministic)
}
func (m *PBError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBError.Merge(m, src)
}
func (m *PBError) XXX_Size() int {
	return xxx_messageInfo_PBError.Size(m)
}
func (m *PBError) XXX_DiscardUnknown() {
	xxx_messageInfo_PBError.DiscardUnknown(m)
}

var xxx_messageInfo_PBError proto.InternalMessageInfo

func (m *PBError) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *PBError) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*PBError)(nil), "tsprotos.PBError")
}

func init() { proto.RegisterFile("utils.proto", fileDescriptor_c91c651f4717a5f2) }

var fileDescriptor_c91c651f4717a5f2 = []byte{
	// 93 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x2d, 0xc9, 0xcc,
	0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x28, 0x29, 0x06, 0x33, 0x8a, 0x95, 0xcc,
	0xb9, 0xd8, 0x03, 0x9c, 0x5c, 0x8b, 0x8a, 0xf2, 0x8b, 0x84, 0x84, 0xb8, 0x58, 0x92, 0xf3, 0x53,
	0x52, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x21, 0x09, 0x2e, 0xf6, 0xdc, 0xd4,
	0xe2, 0xe2, 0xc4, 0xf4, 0x54, 0x09, 0x26, 0xb0, 0x30, 0x8c, 0x9b, 0xc4, 0x06, 0x36, 0xc0, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0x39, 0xba, 0x94, 0xd8, 0x58, 0x00, 0x00, 0x00,
}
