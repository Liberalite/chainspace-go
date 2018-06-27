// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types.proto

package service

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Message struct {
	Opcode  uint32 `protobuf:"varint,1,opt,name=opcode,proto3" json:"opcode,omitempty"`
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{0} }

func (m *Message) GetOpcode() uint32 {
	if m != nil {
		return m.Opcode
	}
	return 0
}

func (m *Message) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "service.Message")
}

func init() { proto.RegisterFile("types.proto", fileDescriptorTypes) }

var fileDescriptorTypes = []byte{
	// 131 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0xa9, 0x2c, 0x48,
	0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e,
	0x95, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f,
	0xcf, 0xd7, 0x07, 0xcb, 0x27, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0xa7, 0x64,
	0xcd, 0xc5, 0xee, 0x9b, 0x5a, 0x5c, 0x9c, 0x98, 0x9e, 0x2a, 0x24, 0xc6, 0xc5, 0x96, 0x5f, 0x90,
	0x9c, 0x9f, 0x92, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x1b, 0x04, 0xe5, 0x09, 0x49, 0x70, 0xb1,
	0x17, 0x24, 0x56, 0xe6, 0xe4, 0x27, 0xa6, 0x48, 0x30, 0x29, 0x30, 0x6a, 0xf0, 0x04, 0xc1, 0xb8,
	0x49, 0x6c, 0x60, 0x33, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x88, 0x57, 0xeb, 0x8a,
	0x00, 0x00, 0x00,
}
