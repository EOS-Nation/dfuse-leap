// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dfuse/eosio/statedb/v1/tablet.proto

package pbstatedb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type AuthLinkValue struct {
	Permission           uint64   `protobuf:"varint,1,opt,name=permission,proto3" json:"permission,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthLinkValue) Reset()         { *m = AuthLinkValue{} }
func (m *AuthLinkValue) String() string { return proto.CompactTextString(m) }
func (*AuthLinkValue) ProtoMessage()    {}
func (*AuthLinkValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cc566c0547764ba, []int{0}
}

func (m *AuthLinkValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthLinkValue.Unmarshal(m, b)
}
func (m *AuthLinkValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthLinkValue.Marshal(b, m, deterministic)
}
func (m *AuthLinkValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthLinkValue.Merge(m, src)
}
func (m *AuthLinkValue) XXX_Size() int {
	return xxx_messageInfo_AuthLinkValue.Size(m)
}
func (m *AuthLinkValue) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthLinkValue.DiscardUnknown(m)
}

var xxx_messageInfo_AuthLinkValue proto.InternalMessageInfo

func (m *AuthLinkValue) GetPermission() uint64 {
	if m != nil {
		return m.Permission
	}
	return 0
}

type ContractStateValue struct {
	Payer                uint64   `protobuf:"varint,1,opt,name=payer,proto3" json:"payer,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContractStateValue) Reset()         { *m = ContractStateValue{} }
func (m *ContractStateValue) String() string { return proto.CompactTextString(m) }
func (*ContractStateValue) ProtoMessage()    {}
func (*ContractStateValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cc566c0547764ba, []int{1}
}

func (m *ContractStateValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContractStateValue.Unmarshal(m, b)
}
func (m *ContractStateValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContractStateValue.Marshal(b, m, deterministic)
}
func (m *ContractStateValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractStateValue.Merge(m, src)
}
func (m *ContractStateValue) XXX_Size() int {
	return xxx_messageInfo_ContractStateValue.Size(m)
}
func (m *ContractStateValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractStateValue.DiscardUnknown(m)
}

var xxx_messageInfo_ContractStateValue proto.InternalMessageInfo

func (m *ContractStateValue) GetPayer() uint64 {
	if m != nil {
		return m.Payer
	}
	return 0
}

func (m *ContractStateValue) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ContractTableScopeValue struct {
	Payer                uint64   `protobuf:"varint,1,opt,name=payer,proto3" json:"payer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContractTableScopeValue) Reset()         { *m = ContractTableScopeValue{} }
func (m *ContractTableScopeValue) String() string { return proto.CompactTextString(m) }
func (*ContractTableScopeValue) ProtoMessage()    {}
func (*ContractTableScopeValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cc566c0547764ba, []int{2}
}

func (m *ContractTableScopeValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContractTableScopeValue.Unmarshal(m, b)
}
func (m *ContractTableScopeValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContractTableScopeValue.Marshal(b, m, deterministic)
}
func (m *ContractTableScopeValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractTableScopeValue.Merge(m, src)
}
func (m *ContractTableScopeValue) XXX_Size() int {
	return xxx_messageInfo_ContractTableScopeValue.Size(m)
}
func (m *ContractTableScopeValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractTableScopeValue.DiscardUnknown(m)
}

var xxx_messageInfo_ContractTableScopeValue proto.InternalMessageInfo

func (m *ContractTableScopeValue) GetPayer() uint64 {
	if m != nil {
		return m.Payer
	}
	return 0
}

// KeyAccountValue is actual empty and the bool field is there mainly to permit a future
// extension where we could add more fields in there.
type KeyAccountValue struct {
	Present              bool     `protobuf:"varint,1,opt,name=present,proto3" json:"present,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyAccountValue) Reset()         { *m = KeyAccountValue{} }
func (m *KeyAccountValue) String() string { return proto.CompactTextString(m) }
func (*KeyAccountValue) ProtoMessage()    {}
func (*KeyAccountValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_5cc566c0547764ba, []int{3}
}

func (m *KeyAccountValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyAccountValue.Unmarshal(m, b)
}
func (m *KeyAccountValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyAccountValue.Marshal(b, m, deterministic)
}
func (m *KeyAccountValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyAccountValue.Merge(m, src)
}
func (m *KeyAccountValue) XXX_Size() int {
	return xxx_messageInfo_KeyAccountValue.Size(m)
}
func (m *KeyAccountValue) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyAccountValue.DiscardUnknown(m)
}

var xxx_messageInfo_KeyAccountValue proto.InternalMessageInfo

func (m *KeyAccountValue) GetPresent() bool {
	if m != nil {
		return m.Present
	}
	return false
}

func init() {
	proto.RegisterType((*AuthLinkValue)(nil), "dfuse.eosio.statedb.v1.AuthLinkValue")
	proto.RegisterType((*ContractStateValue)(nil), "dfuse.eosio.statedb.v1.ContractStateValue")
	proto.RegisterType((*ContractTableScopeValue)(nil), "dfuse.eosio.statedb.v1.ContractTableScopeValue")
	proto.RegisterType((*KeyAccountValue)(nil), "dfuse.eosio.statedb.v1.KeyAccountValue")
}

func init() {
	proto.RegisterFile("dfuse/eosio/statedb/v1/tablet.proto", fileDescriptor_5cc566c0547764ba)
}

var fileDescriptor_5cc566c0547764ba = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0x59, 0xa9, 0x7f, 0x18, 0x14, 0x21, 0x88, 0xee, 0x49, 0xca, 0x7a, 0x29, 0x88, 0x09,
	0xc5, 0xa3, 0x20, 0xd4, 0xe2, 0x49, 0x4f, 0x5b, 0xf1, 0xe0, 0x2d, 0xc9, 0x8e, 0x36, 0xd8, 0x66,
	0x42, 0x32, 0x29, 0xf4, 0xdb, 0x4b, 0xb3, 0x5b, 0xf0, 0x20, 0xde, 0xde, 0x0b, 0xbf, 0xf7, 0xc8,
	0x3c, 0xb8, 0xe9, 0x3e, 0x73, 0x42, 0x85, 0x94, 0x1c, 0xa9, 0xc4, 0x9a, 0xb1, 0x33, 0x6a, 0x33,
	0x55, 0xac, 0xcd, 0x0a, 0x59, 0x86, 0x48, 0x4c, 0xe2, 0xb2, 0x40, 0xb2, 0x40, 0x72, 0x80, 0xe4,
	0x66, 0xda, 0x28, 0x38, 0x9b, 0x65, 0x5e, 0xbe, 0x3a, 0xff, 0xfd, 0xae, 0x57, 0x19, 0xc5, 0x35,
	0x40, 0xc0, 0xb8, 0x76, 0x29, 0x39, 0xf2, 0x75, 0x35, 0xae, 0x26, 0xa3, 0xf6, 0xd7, 0x4b, 0xf3,
	0x08, 0x62, 0x4e, 0x9e, 0xa3, 0xb6, 0xbc, 0xd8, 0xd5, 0xf4, 0xa9, 0x0b, 0x38, 0x0c, 0x7a, 0x8b,
	0x71, 0x08, 0xf4, 0x46, 0x08, 0x18, 0x75, 0x9a, 0x75, 0x7d, 0x30, 0xae, 0x26, 0xa7, 0x6d, 0xd1,
	0x8d, 0x82, 0xab, 0x7d, 0xfe, 0x6d, 0xf7, 0xc1, 0x85, 0xa5, 0xf0, 0x5f, 0x49, 0x73, 0x0b, 0xe7,
	0x2f, 0xb8, 0x9d, 0x59, 0x4b, 0xd9, 0x73, 0x0f, 0xd6, 0x70, 0x1c, 0x22, 0x26, 0xf4, 0x5c, 0xd0,
	0x93, 0x76, 0x6f, 0x9f, 0x9e, 0x3f, 0xe6, 0x5f, 0x8e, 0x97, 0xd9, 0x48, 0x4b, 0x6b, 0x55, 0x6e,
	0xbe, 0x73, 0x34, 0x88, 0x7e, 0xa1, 0x60, 0xd4, 0xdf, 0x83, 0x3d, 0x04, 0x33, 0x18, 0x73, 0x54,
	0x46, 0xbb, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x01, 0x25, 0x90, 0x11, 0x5b, 0x01, 0x00, 0x00,
}