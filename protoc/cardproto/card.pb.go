// Code generated by protoc-gen-go. DO NOT EDIT.
// source: card.proto

/*
Package cardproto is a generated protocol buffer package.

It is generated from these files:
	card.proto

It has these top-level messages:
	CardBlockSyncRequest
	CardBlockSyncResponse
	CardBlocksFetchRequest
	CardBlockFetchResponse
	CardBlockPushRequest
	CardBlock
*/
package cardproto

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

type CardBlockSyncRequest struct {
	Height int64  `protobuf:"varint,1,opt,name=Height" json:"Height,omitempty"`
	CardID string `protobuf:"bytes,2,opt,name=CardID" json:"CardID,omitempty"`
}

func (m *CardBlockSyncRequest) Reset()                    { *m = CardBlockSyncRequest{} }
func (m *CardBlockSyncRequest) String() string            { return proto.CompactTextString(m) }
func (*CardBlockSyncRequest) ProtoMessage()               {}
func (*CardBlockSyncRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CardBlockSyncRequest) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *CardBlockSyncRequest) GetCardID() string {
	if m != nil {
		return m.CardID
	}
	return ""
}

type CardBlockSyncResponse struct {
	Valid      bool   `protobuf:"varint,1,opt,name=Valid" json:"Valid,omitempty"`
	Height     int64  `protobuf:"varint,2,opt,name=Height" json:"Height,omitempty"`
	CardID     string `protobuf:"bytes,3,opt,name=CardID" json:"CardID,omitempty"`
	PrevCardID string `protobuf:"bytes,4,opt,name=PrevCardID" json:"PrevCardID,omitempty"`
}

func (m *CardBlockSyncResponse) Reset()                    { *m = CardBlockSyncResponse{} }
func (m *CardBlockSyncResponse) String() string            { return proto.CompactTextString(m) }
func (*CardBlockSyncResponse) ProtoMessage()               {}
func (*CardBlockSyncResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CardBlockSyncResponse) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *CardBlockSyncResponse) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *CardBlockSyncResponse) GetCardID() string {
	if m != nil {
		return m.CardID
	}
	return ""
}

func (m *CardBlockSyncResponse) GetPrevCardID() string {
	if m != nil {
		return m.PrevCardID
	}
	return ""
}

type CardBlocksFetchRequest struct {
	Height int64 `protobuf:"varint,1,opt,name=Height" json:"Height,omitempty"`
}

func (m *CardBlocksFetchRequest) Reset()                    { *m = CardBlocksFetchRequest{} }
func (m *CardBlocksFetchRequest) String() string            { return proto.CompactTextString(m) }
func (*CardBlocksFetchRequest) ProtoMessage()               {}
func (*CardBlocksFetchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CardBlocksFetchRequest) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type CardBlockFetchResponse struct {
	Valid      bool         `protobuf:"varint,1,opt,name=Valid" json:"Valid,omitempty"`
	Finish     bool         `protobuf:"varint,2,opt,name=Finish" json:"Finish,omitempty"`
	CardBlocks []*CardBlock `protobuf:"bytes,3,rep,name=CardBlocks" json:"CardBlocks,omitempty"`
}

func (m *CardBlockFetchResponse) Reset()                    { *m = CardBlockFetchResponse{} }
func (m *CardBlockFetchResponse) String() string            { return proto.CompactTextString(m) }
func (*CardBlockFetchResponse) ProtoMessage()               {}
func (*CardBlockFetchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CardBlockFetchResponse) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *CardBlockFetchResponse) GetFinish() bool {
	if m != nil {
		return m.Finish
	}
	return false
}

func (m *CardBlockFetchResponse) GetCardBlocks() []*CardBlock {
	if m != nil {
		return m.CardBlocks
	}
	return nil
}

type CardBlockPushRequest struct {
	CardBlock *CardBlock `protobuf:"bytes,1,opt,name=CardBlock" json:"CardBlock,omitempty"`
}

func (m *CardBlockPushRequest) Reset()                    { *m = CardBlockPushRequest{} }
func (m *CardBlockPushRequest) String() string            { return proto.CompactTextString(m) }
func (*CardBlockPushRequest) ProtoMessage()               {}
func (*CardBlockPushRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CardBlockPushRequest) GetCardBlock() *CardBlock {
	if m != nil {
		return m.CardBlock
	}
	return nil
}

type CardBlock struct {
	Version    string `protobuf:"bytes,1,opt,name=Version" json:"Version,omitempty"`
	Hard       int32  `protobuf:"varint,2,opt,name=Hard" json:"Hard,omitempty"`
	PubKey     string `protobuf:"bytes,3,opt,name=PubKey" json:"PubKey,omitempty"`
	Timestamp  int64  `protobuf:"varint,4,opt,name=Timestamp" json:"Timestamp,omitempty"`
	RandNumber int32  `protobuf:"varint,5,opt,name=RandNumber" json:"RandNumber,omitempty"`
	PrevCardID string `protobuf:"bytes,6,opt,name=PrevCardID" json:"PrevCardID,omitempty"`
	Height     int64  `protobuf:"varint,7,opt,name=Height" json:"Height,omitempty"`
}

func (m *CardBlock) Reset()                    { *m = CardBlock{} }
func (m *CardBlock) String() string            { return proto.CompactTextString(m) }
func (*CardBlock) ProtoMessage()               {}
func (*CardBlock) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CardBlock) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *CardBlock) GetHard() int32 {
	if m != nil {
		return m.Hard
	}
	return 0
}

func (m *CardBlock) GetPubKey() string {
	if m != nil {
		return m.PubKey
	}
	return ""
}

func (m *CardBlock) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *CardBlock) GetRandNumber() int32 {
	if m != nil {
		return m.RandNumber
	}
	return 0
}

func (m *CardBlock) GetPrevCardID() string {
	if m != nil {
		return m.PrevCardID
	}
	return ""
}

func (m *CardBlock) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func init() {
	proto.RegisterType((*CardBlockSyncRequest)(nil), "cardproto.CardBlockSyncRequest")
	proto.RegisterType((*CardBlockSyncResponse)(nil), "cardproto.CardBlockSyncResponse")
	proto.RegisterType((*CardBlocksFetchRequest)(nil), "cardproto.CardBlocksFetchRequest")
	proto.RegisterType((*CardBlockFetchResponse)(nil), "cardproto.CardBlockFetchResponse")
	proto.RegisterType((*CardBlockPushRequest)(nil), "cardproto.CardBlockPushRequest")
	proto.RegisterType((*CardBlock)(nil), "cardproto.CardBlock")
}

func init() { proto.RegisterFile("card.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x25, 0x4d, 0x93, 0x36, 0xe3, 0x6d, 0x89, 0x25, 0x07, 0x91, 0xb0, 0xa7, 0x9c, 0x82, 0x54,
	0x7f, 0x81, 0x4a, 0xa8, 0x0a, 0x12, 0x56, 0xe9, 0x7d, 0x93, 0x2c, 0x66, 0xb1, 0xf9, 0x70, 0x37,
	0x11, 0x0a, 0xfa, 0xfb, 0xfc, 0x5b, 0x92, 0x6d, 0x9a, 0x2c, 0x81, 0xd2, 0xdb, 0xbc, 0x37, 0x33,
	0x6f, 0x3e, 0x1e, 0x40, 0x4a, 0x45, 0x16, 0xd6, 0xa2, 0x6a, 0x2a, 0xe4, 0x74, 0xb1, 0x0a, 0x71,
	0x04, 0xee, 0x03, 0x15, 0xd9, 0xfd, 0xae, 0x4a, 0x3f, 0xdf, 0xf6, 0x65, 0x4a, 0xd8, 0x57, 0xcb,
	0x64, 0x83, 0x56, 0x60, 0x6f, 0x18, 0xff, 0xc8, 0x1b, 0xcf, 0xf0, 0x8d, 0xc0, 0x24, 0x3d, 0xea,
	0xf8, 0xae, 0xfe, 0xe9, 0xd1, 0x9b, 0xf9, 0x46, 0xe0, 0x90, 0x1e, 0xe1, 0x5f, 0xb8, 0x9c, 0xe8,
	0xc8, 0xba, 0x2a, 0x25, 0x43, 0x2e, 0x58, 0x5b, 0xba, 0xe3, 0x99, 0xd2, 0x59, 0x92, 0x03, 0xd0,
	0xe4, 0x67, 0x27, 0xe4, 0x4d, 0x5d, 0x1e, 0x5d, 0x03, 0xc4, 0x82, 0x7d, 0xf7, 0xb9, 0xb9, 0xca,
	0x69, 0x0c, 0xbe, 0x81, 0xd5, 0x30, 0x5e, 0x46, 0xac, 0x49, 0xf3, 0x33, 0x87, 0xe0, 0x1f, 0xad,
	0xa3, 0x6f, 0x38, 0xb7, 0x71, 0xc4, 0x4b, 0x2e, 0x73, 0xb5, 0xf1, 0x92, 0xf4, 0x08, 0xdd, 0x01,
	0x8c, 0x93, 0x3d, 0xd3, 0x37, 0x83, 0x8b, 0xb5, 0x1b, 0x0e, 0x0f, 0x0e, 0x87, 0x24, 0xd1, 0xea,
	0xf0, 0xb3, 0xf6, 0xf6, 0xb8, 0x95, 0xc3, 0xb6, 0x6b, 0x70, 0x06, 0x5e, 0xcd, 0x3f, 0x25, 0x36,
	0x96, 0xe1, 0x3f, 0x43, 0x6b, 0x42, 0x1e, 0x2c, 0xb6, 0x4c, 0x48, 0x5e, 0x95, 0xaa, 0xdf, 0x21,
	0x47, 0x88, 0x10, 0xcc, 0x37, 0x54, 0x64, 0x6a, 0x7f, 0x8b, 0xa8, 0xb8, 0xbb, 0x2a, 0x6e, 0x93,
	0x17, 0xb6, 0x3f, 0xfe, 0xfb, 0x80, 0xd0, 0x15, 0x38, 0xef, 0xbc, 0x60, 0xb2, 0xa1, 0x45, 0xad,
	0xde, 0x6d, 0x92, 0x91, 0xe8, 0xdc, 0x20, 0xb4, 0xcc, 0x5e, 0xdb, 0x22, 0x61, 0xc2, 0xb3, 0x94,
	0x9e, 0xc6, 0x4c, 0xdc, 0xb2, 0xa7, 0x6e, 0x69, 0x9e, 0x2c, 0x74, 0x4f, 0x12, 0x5b, 0x5d, 0x79,
	0xfb, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x25, 0x09, 0xb7, 0xac, 0x02, 0x00, 0x00,
}
