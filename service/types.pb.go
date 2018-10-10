// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types.proto

package service

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/golang/protobuf/ptypes/timestamp"

import time "time"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type CONNECTION int32

const (
	CONNECTION_UNKNOWN    CONNECTION = 0
	CONNECTION_BROADCAST  CONNECTION = 1
	CONNECTION_TRANSACTOR CONNECTION = 2
	CONNECTION_KV         CONNECTION = 3
)

var CONNECTION_name = map[int32]string{
	0: "UNKNOWN",
	1: "BROADCAST",
	2: "TRANSACTOR",
	3: "KV",
}
var CONNECTION_value = map[string]int32{
	"UNKNOWN":    0,
	"BROADCAST":  1,
	"TRANSACTOR": 2,
	"KV":         3,
}

func (x CONNECTION) String() string {
	return proto.EnumName(CONNECTION_name, int32(x))
}
func (CONNECTION) EnumDescriptor() ([]byte, []int) { return fileDescriptorTypes, []int{0} }

type Hello struct {
	Agent     string     `protobuf:"bytes,1,opt,name=agent,proto3" json:"agent,omitempty"`
	Payload   []byte     `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	Signature []byte     `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	Type      CONNECTION `protobuf:"varint,4,opt,name=type,proto3,enum=service.CONNECTION" json:"type,omitempty"`
}

func (m *Hello) Reset()                    { *m = Hello{} }
func (m *Hello) String() string            { return proto.CompactTextString(m) }
func (*Hello) ProtoMessage()               {}
func (*Hello) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{0} }

func (m *Hello) GetAgent() string {
	if m != nil {
		return m.Agent
	}
	return ""
}

func (m *Hello) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Hello) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Hello) GetType() CONNECTION {
	if m != nil {
		return m.Type
	}
	return CONNECTION_UNKNOWN
}

type HelloInfo struct {
	Client    uint64    `protobuf:"varint,1,opt,name=client,proto3" json:"client,omitempty"`
	Nonce     []byte    `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Server    uint64    `protobuf:"varint,3,opt,name=server,proto3" json:"server,omitempty"`
	Timestamp time.Time `protobuf:"bytes,4,opt,name=timestamp,stdtime" json:"timestamp"`
}

func (m *HelloInfo) Reset()                    { *m = HelloInfo{} }
func (m *HelloInfo) String() string            { return proto.CompactTextString(m) }
func (*HelloInfo) ProtoMessage()               {}
func (*HelloInfo) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{1} }

func (m *HelloInfo) GetClient() uint64 {
	if m != nil {
		return m.Client
	}
	return 0
}

func (m *HelloInfo) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *HelloInfo) GetServer() uint64 {
	if m != nil {
		return m.Server
	}
	return 0
}

func (m *HelloInfo) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

type Message struct {
	ID      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Opcode  int32  `protobuf:"varint,3,opt,name=opcode,proto3" json:"opcode,omitempty"`
	Payload []byte `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{2} }

func (m *Message) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Message) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *Message) GetOpcode() int32 {
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
	proto.RegisterType((*Hello)(nil), "service.Hello")
	proto.RegisterType((*HelloInfo)(nil), "service.HelloInfo")
	proto.RegisterType((*Message)(nil), "service.Message")
	proto.RegisterEnum("service.CONNECTION", CONNECTION_name, CONNECTION_value)
}

func init() { proto.RegisterFile("service/types.proto", fileDescriptorTypes) }

var fileDescriptorTypes = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xd1, 0xca, 0x9b, 0x30,
	0x18, 0x86, 0xab, 0xb5, 0x3a, 0xbf, 0x6e, 0x45, 0xd2, 0x31, 0xa4, 0x0c, 0x2c, 0x3d, 0x59, 0x19,
	0xcc, 0x42, 0x77, 0x03, 0x53, 0x3b, 0x58, 0x29, 0x53, 0x48, 0xdd, 0x76, 0x6c, 0x35, 0xcd, 0x02,
	0xd6, 0x88, 0xda, 0x41, 0x4f, 0x76, 0x0d, 0x3b, 0xd9, 0x3d, 0xed, 0x2a, 0x3a, 0xd8, 0x95, 0x0c,
	0x13, 0xfd, 0xfd, 0xff, 0x33, 0x9f, 0xe4, 0x8d, 0xdf, 0xf3, 0xbd, 0x30, 0xaf, 0x49, 0xf5, 0x83,
	0xa5, 0x64, 0xd3, 0xdc, 0x4a, 0x52, 0xbb, 0x65, 0xc5, 0x1b, 0x8e, 0x8c, 0xee, 0x70, 0xf1, 0x8e,
	0xb2, 0xe6, 0xfb, 0xf5, 0xe4, 0xa6, 0xfc, 0xb2, 0xa1, 0x9c, 0xf2, 0x8d, 0xb8, 0x3f, 0x5d, 0xcf,
	0x82, 0x04, 0x88, 0x2f, 0xf9, 0x6e, 0xe1, 0x50, 0xce, 0x69, 0x4e, 0x86, 0x54, 0xc3, 0x2e, 0xa4,
	0x6e, 0x92, 0x4b, 0x29, 0x03, 0xab, 0x9f, 0x30, 0xf9, 0x44, 0xf2, 0x9c, 0xa3, 0x97, 0x30, 0x49,
	0x28, 0x29, 0x1a, 0x5b, 0x59, 0x2a, 0x6b, 0x13, 0x4b, 0x40, 0x36, 0x18, 0x65, 0x72, 0xcb, 0x79,
	0x92, 0xd9, 0xea, 0x52, 0x59, 0x3f, 0xc7, 0x3d, 0xa2, 0xd7, 0x60, 0xd6, 0x8c, 0x16, 0x49, 0x73,
	0xad, 0x88, 0x3d, 0x16, 0x77, 0xc3, 0x01, 0x7a, 0x03, 0x5a, 0xab, 0x6f, 0x6b, 0x4b, 0x65, 0x3d,
	0xdb, 0xce, 0xdd, 0x4e, 0xdf, 0x0d, 0xa2, 0x30, 0xfc, 0x18, 0xc4, 0xfb, 0x28, 0xc4, 0x22, 0xb0,
	0xfa, 0xad, 0x80, 0x29, 0x04, 0xf6, 0xc5, 0x99, 0xa3, 0x57, 0xa0, 0xa7, 0x39, 0xeb, 0x2d, 0x34,
	0xdc, 0x51, 0x2b, 0x57, 0xf0, 0x22, 0x25, 0x9d, 0x84, 0x84, 0x36, 0xdd, 0xfe, 0x97, 0x54, 0x62,
	0xbe, 0x86, 0x3b, 0x42, 0x3e, 0x98, 0x0f, 0x6b, 0x0a, 0x83, 0xe9, 0x76, 0xe1, 0xca, 0x22, 0xdc,
	0xbe, 0x08, 0x37, 0xee, 0x13, 0xfe, 0xb3, 0x3f, 0x77, 0x67, 0xf4, 0xeb, 0xaf, 0xa3, 0xe0, 0xe1,
	0xd9, 0x8a, 0x81, 0xf1, 0x99, 0xd4, 0x75, 0x42, 0xdb, 0x31, 0x2a, 0xcb, 0xa4, 0x90, 0xaf, 0xff,
	0xbb, 0x3b, 0xea, 0x7e, 0x87, 0x55, 0x96, 0xb5, 0x52, 0xa4, 0xaa, 0x78, 0x25, 0xa4, 0x4c, 0x2c,
	0xa1, 0x95, 0xe2, 0x65, 0xca, 0x33, 0x59, 0xca, 0x04, 0x77, 0xf4, 0xb8, 0x49, 0xed, 0x49, 0x93,
	0x6f, 0x3f, 0x00, 0x0c, 0xb5, 0xa0, 0x29, 0x18, 0x5f, 0xc2, 0x43, 0x18, 0x7d, 0x0b, 0xad, 0x11,
	0x7a, 0x01, 0xa6, 0x8f, 0x23, 0x6f, 0x17, 0x78, 0xc7, 0xd8, 0x52, 0xd0, 0x0c, 0x20, 0xc6, 0x5e,
	0x78, 0xf4, 0x82, 0x38, 0xc2, 0x96, 0x8a, 0x74, 0x50, 0x0f, 0x5f, 0xad, 0xf1, 0x49, 0x17, 0x5b,
	0xbd, 0xff, 0x1f, 0x00, 0x00, 0xff, 0xff, 0xff, 0xac, 0x6b, 0x63, 0x3b, 0x02, 0x00, 0x00,
}
