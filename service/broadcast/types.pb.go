// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types.proto

package broadcast

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

type OP int32

const (
	OP_UNKNOWN         OP = 0
	OP_BLOCK           OP = 1
	OP_BLOCKS_REQUEST  OP = 2
	OP_BLOCKS_RESPONSE OP = 3
)

var OP_name = map[int32]string{
	0: "UNKNOWN",
	1: "BLOCK",
	2: "BLOCKS_REQUEST",
	3: "BLOCKS_RESPONSE",
}
var OP_value = map[string]int32{
	"UNKNOWN":         0,
	"BLOCK":           1,
	"BLOCKS_REQUEST":  2,
	"BLOCKS_RESPONSE": 3,
}

func (x OP) String() string {
	return proto.EnumName(OP_name, int32(x))
}
func (OP) EnumDescriptor() ([]byte, []int) { return fileDescriptorTypes, []int{0} }

type Block struct {
	Node     uint64   `protobuf:"varint,1,opt,name=node,proto3" json:"node,omitempty"`
	Number   uint64   `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	Entries  []*Entry `protobuf:"bytes,3,rep,name=entries" json:"entries,omitempty"`
	Previous []byte   `protobuf:"bytes,4,opt,name=previous,proto3" json:"previous,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{0} }

func (m *Block) GetNode() uint64 {
	if m != nil {
		return m.Node
	}
	return 0
}

func (m *Block) GetNumber() uint64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Block) GetEntries() []*Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *Block) GetPrevious() []byte {
	if m != nil {
		return m.Previous
	}
	return nil
}

type BlockReference struct {
	Node   uint64 `protobuf:"varint,1,opt,name=node,proto3" json:"node,omitempty"`
	Number uint64 `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	Hash   []byte `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *BlockReference) Reset()                    { *m = BlockReference{} }
func (m *BlockReference) String() string            { return proto.CompactTextString(m) }
func (*BlockReference) ProtoMessage()               {}
func (*BlockReference) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{1} }

func (m *BlockReference) GetNode() uint64 {
	if m != nil {
		return m.Node
	}
	return 0
}

func (m *BlockReference) GetNumber() uint64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *BlockReference) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type BlocksRequest struct {
	Blocks []*BlockReference `protobuf:"bytes,1,rep,name=blocks" json:"blocks,omitempty"`
}

func (m *BlocksRequest) Reset()                    { *m = BlocksRequest{} }
func (m *BlocksRequest) String() string            { return proto.CompactTextString(m) }
func (*BlocksRequest) ProtoMessage()               {}
func (*BlocksRequest) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{2} }

func (m *BlocksRequest) GetBlocks() []*BlockReference {
	if m != nil {
		return m.Blocks
	}
	return nil
}

type BlocksResponse struct {
	Signed []*SignedBlock `protobuf:"bytes,1,rep,name=signed" json:"signed,omitempty"`
}

func (m *BlocksResponse) Reset()                    { *m = BlocksResponse{} }
func (m *BlocksResponse) String() string            { return proto.CompactTextString(m) }
func (*BlocksResponse) ProtoMessage()               {}
func (*BlocksResponse) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{3} }

func (m *BlocksResponse) GetSigned() []*SignedBlock {
	if m != nil {
		return m.Signed
	}
	return nil
}

type Entry struct {
	// Types that are valid to be assigned to Value:
	//	*Entry_Block
	//	*Entry_Transaction
	Value isEntry_Value `protobuf_oneof:"value"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{4} }

type isEntry_Value interface {
	isEntry_Value()
}

type Entry_Block struct {
	Block *BlockReference `protobuf:"bytes,1,opt,name=block,oneof"`
}
type Entry_Transaction struct {
	Transaction *TransactionData `protobuf:"bytes,2,opt,name=transaction,oneof"`
}

func (*Entry_Block) isEntry_Value()       {}
func (*Entry_Transaction) isEntry_Value() {}

func (m *Entry) GetValue() isEntry_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Entry) GetBlock() *BlockReference {
	if x, ok := m.GetValue().(*Entry_Block); ok {
		return x.Block
	}
	return nil
}

func (m *Entry) GetTransaction() *TransactionData {
	if x, ok := m.GetValue().(*Entry_Transaction); ok {
		return x.Transaction
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Entry) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Entry_OneofMarshaler, _Entry_OneofUnmarshaler, _Entry_OneofSizer, []interface{}{
		(*Entry_Block)(nil),
		(*Entry_Transaction)(nil),
	}
}

func _Entry_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Entry)
	// value
	switch x := m.Value.(type) {
	case *Entry_Block:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Block); err != nil {
			return err
		}
	case *Entry_Transaction:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Transaction); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Entry.Value has unexpected type %T", x)
	}
	return nil
}

func _Entry_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Entry)
	switch tag {
	case 1: // value.block
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BlockReference)
		err := b.DecodeMessage(msg)
		m.Value = &Entry_Block{msg}
		return true, err
	case 2: // value.transaction
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TransactionData)
		err := b.DecodeMessage(msg)
		m.Value = &Entry_Transaction{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Entry_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Entry)
	// value
	switch x := m.Value.(type) {
	case *Entry_Block:
		s := proto.Size(x.Block)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Entry_Transaction:
		s := proto.Size(x.Transaction)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type SignedBlock struct {
	Data      []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *SignedBlock) Reset()                    { *m = SignedBlock{} }
func (m *SignedBlock) String() string            { return proto.CompactTextString(m) }
func (*SignedBlock) ProtoMessage()               {}
func (*SignedBlock) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{5} }

func (m *SignedBlock) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *SignedBlock) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type TransactionData struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Fee  uint64 `protobuf:"varint,2,opt,name=fee,proto3" json:"fee,omitempty"`
}

func (m *TransactionData) Reset()                    { *m = TransactionData{} }
func (m *TransactionData) String() string            { return proto.CompactTextString(m) }
func (*TransactionData) ProtoMessage()               {}
func (*TransactionData) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{6} }

func (m *TransactionData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *TransactionData) GetFee() uint64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func init() {
	proto.RegisterType((*Block)(nil), "broadcast.Block")
	proto.RegisterType((*BlockReference)(nil), "broadcast.BlockReference")
	proto.RegisterType((*BlocksRequest)(nil), "broadcast.BlocksRequest")
	proto.RegisterType((*BlocksResponse)(nil), "broadcast.BlocksResponse")
	proto.RegisterType((*Entry)(nil), "broadcast.Entry")
	proto.RegisterType((*SignedBlock)(nil), "broadcast.SignedBlock")
	proto.RegisterType((*TransactionData)(nil), "broadcast.TransactionData")
	proto.RegisterEnum("broadcast.OP", OP_name, OP_value)
}

func init() { proto.RegisterFile("types.proto", fileDescriptorTypes) }

var fileDescriptorTypes = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0x93, 0x38, 0x21, 0xe3, 0xd0, 0x5a, 0x83, 0x54, 0x99, 0x88, 0x43, 0xe4, 0x53, 0x54,
	0x09, 0x57, 0x2d, 0x07, 0x6e, 0x80, 0x02, 0x96, 0x2a, 0x15, 0x25, 0x61, 0xdd, 0x8a, 0x23, 0x5a,
	0x3b, 0x93, 0xc4, 0xa2, 0xdd, 0x0d, 0xbb, 0xeb, 0x4a, 0x15, 0x47, 0x7e, 0x1c, 0x65, 0x62, 0x52,
	0x83, 0x10, 0x12, 0xb7, 0x37, 0x6f, 0xdf, 0xbc, 0xf7, 0x76, 0x6d, 0x08, 0xdc, 0xc3, 0x96, 0x6c,
	0xb2, 0x35, 0xda, 0x69, 0x1c, 0xe4, 0x46, 0xcb, 0x65, 0x21, 0xad, 0x1b, 0xbd, 0x5c, 0x97, 0x6e,
	0x53, 0xe5, 0x49, 0xa1, 0xef, 0xce, 0xd6, 0x7a, 0xad, 0xcf, 0x58, 0x91, 0x57, 0x2b, 0x9e, 0x78,
	0x60, 0xb4, 0xdf, 0x8c, 0xbf, 0x83, 0x3f, 0xbd, 0xd5, 0xc5, 0x57, 0x44, 0xe8, 0x2a, 0xbd, 0xa4,
	0xc8, 0x1b, 0x7b, 0x93, 0xae, 0x60, 0x8c, 0x27, 0xd0, 0x53, 0xd5, 0x5d, 0x4e, 0x26, 0x6a, 0x33,
	0x5b, 0x4f, 0x78, 0x0a, 0x7d, 0x52, 0xce, 0x94, 0x64, 0xa3, 0xce, 0xb8, 0x33, 0x09, 0x2e, 0xc2,
	0xe4, 0x50, 0x20, 0x49, 0x95, 0x33, 0x0f, 0xe2, 0x97, 0x00, 0x47, 0xf0, 0x64, 0x6b, 0xe8, 0xbe,
	0xd4, 0x95, 0x8d, 0xba, 0x63, 0x6f, 0x32, 0x14, 0x87, 0x39, 0x5e, 0xc0, 0x11, 0x87, 0x0b, 0x5a,
	0x91, 0x21, 0x55, 0xd0, 0x7f, 0xb5, 0x40, 0xe8, 0x6e, 0xa4, 0xdd, 0x44, 0x1d, 0x76, 0x65, 0x1c,
	0x4f, 0xe1, 0x29, 0x3b, 0x5a, 0x41, 0xdf, 0x2a, 0xb2, 0x0e, 0xcf, 0xa1, 0x97, 0x33, 0x11, 0x79,
	0xdc, 0xf4, 0x79, 0xa3, 0xe9, 0xef, 0xd9, 0xa2, 0x16, 0xc6, 0xef, 0xea, 0x56, 0x56, 0x90, 0xdd,
	0x6a, 0x65, 0x09, 0x13, 0xe8, 0xd9, 0x72, 0xad, 0x68, 0x59, 0x9b, 0x9c, 0x34, 0x4c, 0x32, 0x3e,
	0xd8, 0x5b, 0xd5, 0xaa, 0xf8, 0x87, 0x07, 0x3e, 0x3f, 0x03, 0x9e, 0x83, 0xcf, 0xae, 0x7c, 0xa1,
	0x7f, 0xa5, 0x5f, 0xb6, 0xc4, 0x5e, 0x89, 0x6f, 0x20, 0x70, 0x46, 0x2a, 0x2b, 0x0b, 0x57, 0x6a,
	0xc5, 0x77, 0x0e, 0x2e, 0x46, 0x8d, 0xc5, 0xeb, 0xc7, 0xd3, 0x0f, 0xd2, 0xc9, 0xcb, 0x96, 0x68,
	0x2e, 0x4c, 0xfb, 0xe0, 0xdf, 0xcb, 0xdb, 0x8a, 0xe2, 0xb7, 0x10, 0x34, 0xca, 0xed, 0x9e, 0x6b,
	0x29, 0x9d, 0xe4, 0x26, 0x43, 0xc1, 0x18, 0x5f, 0xc0, 0x60, 0x57, 0x59, 0xba, 0xca, 0x10, 0x27,
	0x0d, 0xc5, 0x23, 0x11, 0xbf, 0x86, 0xe3, 0x3f, 0xb2, 0xfe, 0x6a, 0x12, 0x42, 0x67, 0x45, 0x54,
	0x7f, 0x9c, 0x1d, 0x3c, 0x4d, 0xa1, 0x3d, 0x5f, 0x60, 0x00, 0xfd, 0x9b, 0xd9, 0xd5, 0x6c, 0xfe,
	0x79, 0x16, 0xb6, 0x70, 0x00, 0xfe, 0xf4, 0xe3, 0xfc, 0xfd, 0x55, 0xe8, 0x21, 0xc2, 0x11, 0xc3,
	0xec, 0x8b, 0x48, 0x3f, 0xdd, 0xa4, 0xd9, 0x75, 0xd8, 0xc6, 0x67, 0x70, 0x7c, 0xe0, 0xb2, 0xc5,
	0x7c, 0x96, 0xa5, 0x61, 0x27, 0xef, 0xf1, 0x2f, 0xfa, 0xea, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xaf, 0x5e, 0x8c, 0x6e, 0xeb, 0x02, 0x00, 0x00,
}