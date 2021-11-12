// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dfuse/eosio/accounthist/v1/accounthist.proto

package pbaccounthist

import (
	context "context"
	fmt "fmt"
	v1 "github.com/dfuse-io/dfuse-eosio/pb/dfuse/eosio/codec/v1"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetActionsRequest struct {
	Account              uint64   `protobuf:"varint,1,opt,name=account,proto3" json:"account,omitempty"`
	Limit                uint32   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Cursor               *Cursor  `protobuf:"bytes,3,opt,name=cursor,proto3" json:"cursor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetActionsRequest) Reset()         { *m = GetActionsRequest{} }
func (m *GetActionsRequest) String() string { return proto.CompactTextString(m) }
func (*GetActionsRequest) ProtoMessage()    {}
func (*GetActionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c22ddb60199ece6, []int{0}
}

func (m *GetActionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetActionsRequest.Unmarshal(m, b)
}
func (m *GetActionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetActionsRequest.Marshal(b, m, deterministic)
}
func (m *GetActionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetActionsRequest.Merge(m, src)
}
func (m *GetActionsRequest) XXX_Size() int {
	return xxx_messageInfo_GetActionsRequest.Size(m)
}
func (m *GetActionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetActionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetActionsRequest proto.InternalMessageInfo

func (m *GetActionsRequest) GetAccount() uint64 {
	if m != nil {
		return m.Account
	}
	return 0
}

func (m *GetActionsRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetActionsRequest) GetCursor() *Cursor {
	if m != nil {
		return m.Cursor
	}
	return nil
}

type GetTokenActionsRequest struct {
	Account              uint64   `protobuf:"varint,1,opt,name=account,proto3" json:"account,omitempty"`
	Contract             uint64   `protobuf:"varint,2,opt,name=contract,proto3" json:"contract,omitempty"`
	Limit                uint32   `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Cursor               *Cursor  `protobuf:"bytes,4,opt,name=cursor,proto3" json:"cursor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTokenActionsRequest) Reset()         { *m = GetTokenActionsRequest{} }
func (m *GetTokenActionsRequest) String() string { return proto.CompactTextString(m) }
func (*GetTokenActionsRequest) ProtoMessage()    {}
func (*GetTokenActionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c22ddb60199ece6, []int{1}
}

func (m *GetTokenActionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTokenActionsRequest.Unmarshal(m, b)
}
func (m *GetTokenActionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTokenActionsRequest.Marshal(b, m, deterministic)
}
func (m *GetTokenActionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTokenActionsRequest.Merge(m, src)
}
func (m *GetTokenActionsRequest) XXX_Size() int {
	return xxx_messageInfo_GetTokenActionsRequest.Size(m)
}
func (m *GetTokenActionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTokenActionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTokenActionsRequest proto.InternalMessageInfo

func (m *GetTokenActionsRequest) GetAccount() uint64 {
	if m != nil {
		return m.Account
	}
	return 0
}

func (m *GetTokenActionsRequest) GetContract() uint64 {
	if m != nil {
		return m.Contract
	}
	return 0
}

func (m *GetTokenActionsRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetTokenActionsRequest) GetCursor() *Cursor {
	if m != nil {
		return m.Cursor
	}
	return nil
}

type ActionResponse struct {
	Cursor               *Cursor         `protobuf:"bytes,1,opt,name=cursor,proto3" json:"cursor,omitempty"`
	ActionTrace          *v1.ActionTrace `protobuf:"bytes,2,opt,name=action_trace,json=actionTrace,proto3" json:"action_trace,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ActionResponse) Reset()         { *m = ActionResponse{} }
func (m *ActionResponse) String() string { return proto.CompactTextString(m) }
func (*ActionResponse) ProtoMessage()    {}
func (*ActionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c22ddb60199ece6, []int{2}
}

func (m *ActionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActionResponse.Unmarshal(m, b)
}
func (m *ActionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActionResponse.Marshal(b, m, deterministic)
}
func (m *ActionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionResponse.Merge(m, src)
}
func (m *ActionResponse) XXX_Size() int {
	return xxx_messageInfo_ActionResponse.Size(m)
}
func (m *ActionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ActionResponse proto.InternalMessageInfo

func (m *ActionResponse) GetCursor() *Cursor {
	if m != nil {
		return m.Cursor
	}
	return nil
}

func (m *ActionResponse) GetActionTrace() *v1.ActionTrace {
	if m != nil {
		return m.ActionTrace
	}
	return nil
}

type ActionRow struct {
	Version              uint32          `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	ActionTrace          *v1.ActionTrace `protobuf:"bytes,2,opt,name=action_trace,json=actionTrace,proto3" json:"action_trace,omitempty"`
	LastDeletedSeq       uint64          `protobuf:"varint,3,opt,name=last_deleted_seq,json=lastDeletedSeq,proto3" json:"last_deleted_seq,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ActionRow) Reset()         { *m = ActionRow{} }
func (m *ActionRow) String() string { return proto.CompactTextString(m) }
func (*ActionRow) ProtoMessage()    {}
func (*ActionRow) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c22ddb60199ece6, []int{3}
}

func (m *ActionRow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActionRow.Unmarshal(m, b)
}
func (m *ActionRow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActionRow.Marshal(b, m, deterministic)
}
func (m *ActionRow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionRow.Merge(m, src)
}
func (m *ActionRow) XXX_Size() int {
	return xxx_messageInfo_ActionRow.Size(m)
}
func (m *ActionRow) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionRow.DiscardUnknown(m)
}

var xxx_messageInfo_ActionRow proto.InternalMessageInfo

func (m *ActionRow) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ActionRow) GetActionTrace() *v1.ActionTrace {
	if m != nil {
		return m.ActionTrace
	}
	return nil
}

func (m *ActionRow) GetLastDeletedSeq() uint64 {
	if m != nil {
		return m.LastDeletedSeq
	}
	return 0
}

type ShardCheckpoint struct {
	InitialStartBlock    uint64   `protobuf:"varint,1,opt,name=initial_start_block,json=initialStartBlock,proto3" json:"initial_start_block,omitempty"`
	TargetStopBlock      uint64   `protobuf:"varint,2,opt,name=target_stop_block,json=targetStopBlock,proto3" json:"target_stop_block,omitempty"`
	LastWrittenBlockNum  uint64   `protobuf:"varint,3,opt,name=last_written_block_num,json=lastWrittenBlockNum,proto3" json:"last_written_block_num,omitempty"`
	LastWrittenBlockId   string   `protobuf:"bytes,4,opt,name=last_written_block_id,json=lastWrittenBlockId,proto3" json:"last_written_block_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShardCheckpoint) Reset()         { *m = ShardCheckpoint{} }
func (m *ShardCheckpoint) String() string { return proto.CompactTextString(m) }
func (*ShardCheckpoint) ProtoMessage()    {}
func (*ShardCheckpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c22ddb60199ece6, []int{4}
}

func (m *ShardCheckpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShardCheckpoint.Unmarshal(m, b)
}
func (m *ShardCheckpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShardCheckpoint.Marshal(b, m, deterministic)
}
func (m *ShardCheckpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShardCheckpoint.Merge(m, src)
}
func (m *ShardCheckpoint) XXX_Size() int {
	return xxx_messageInfo_ShardCheckpoint.Size(m)
}
func (m *ShardCheckpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_ShardCheckpoint.DiscardUnknown(m)
}

var xxx_messageInfo_ShardCheckpoint proto.InternalMessageInfo

func (m *ShardCheckpoint) GetInitialStartBlock() uint64 {
	if m != nil {
		return m.InitialStartBlock
	}
	return 0
}

func (m *ShardCheckpoint) GetTargetStopBlock() uint64 {
	if m != nil {
		return m.TargetStopBlock
	}
	return 0
}

func (m *ShardCheckpoint) GetLastWrittenBlockNum() uint64 {
	if m != nil {
		return m.LastWrittenBlockNum
	}
	return 0
}

func (m *ShardCheckpoint) GetLastWrittenBlockId() string {
	if m != nil {
		return m.LastWrittenBlockId
	}
	return ""
}

type ActionRowAppend struct {
	LastDeletedSeq       uint64   `protobuf:"varint,3,opt,name=last_deleted_seq,json=lastDeletedSeq,proto3" json:"last_deleted_seq,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActionRowAppend) Reset()         { *m = ActionRowAppend{} }
func (m *ActionRowAppend) String() string { return proto.CompactTextString(m) }
func (*ActionRowAppend) ProtoMessage()    {}
func (*ActionRowAppend) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c22ddb60199ece6, []int{5}
}

func (m *ActionRowAppend) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActionRowAppend.Unmarshal(m, b)
}
func (m *ActionRowAppend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActionRowAppend.Marshal(b, m, deterministic)
}
func (m *ActionRowAppend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionRowAppend.Merge(m, src)
}
func (m *ActionRowAppend) XXX_Size() int {
	return xxx_messageInfo_ActionRowAppend.Size(m)
}
func (m *ActionRowAppend) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionRowAppend.DiscardUnknown(m)
}

var xxx_messageInfo_ActionRowAppend proto.InternalMessageInfo

func (m *ActionRowAppend) GetLastDeletedSeq() uint64 {
	if m != nil {
		return m.LastDeletedSeq
	}
	return 0
}

type Cursor struct {
	Version              uint32   `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Magic                uint32   `protobuf:"varint,2,opt,name=magic,proto3" json:"magic,omitempty"`
	Key                  []byte   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	ShardNum             uint32   `protobuf:"varint,5,opt,name=shard_num,json=shardNum,proto3" json:"shard_num,omitempty"`
	SequenceNumber       uint64   `protobuf:"varint,6,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cursor) Reset()         { *m = Cursor{} }
func (m *Cursor) String() string { return proto.CompactTextString(m) }
func (*Cursor) ProtoMessage()    {}
func (*Cursor) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c22ddb60199ece6, []int{6}
}

func (m *Cursor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cursor.Unmarshal(m, b)
}
func (m *Cursor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cursor.Marshal(b, m, deterministic)
}
func (m *Cursor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cursor.Merge(m, src)
}
func (m *Cursor) XXX_Size() int {
	return xxx_messageInfo_Cursor.Size(m)
}
func (m *Cursor) XXX_DiscardUnknown() {
	xxx_messageInfo_Cursor.DiscardUnknown(m)
}

var xxx_messageInfo_Cursor proto.InternalMessageInfo

func (m *Cursor) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Cursor) GetMagic() uint32 {
	if m != nil {
		return m.Magic
	}
	return 0
}

func (m *Cursor) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Cursor) GetShardNum() uint32 {
	if m != nil {
		return m.ShardNum
	}
	return 0
}

func (m *Cursor) GetSequenceNumber() uint64 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

func init() {
	proto.RegisterType((*GetActionsRequest)(nil), "dfuse.eosio.accounthist.v1.GetActionsRequest")
	proto.RegisterType((*GetTokenActionsRequest)(nil), "dfuse.eosio.accounthist.v1.GetTokenActionsRequest")
	proto.RegisterType((*ActionResponse)(nil), "dfuse.eosio.accounthist.v1.ActionResponse")
	proto.RegisterType((*ActionRow)(nil), "dfuse.eosio.accounthist.v1.ActionRow")
	proto.RegisterType((*ShardCheckpoint)(nil), "dfuse.eosio.accounthist.v1.ShardCheckpoint")
	proto.RegisterType((*ActionRowAppend)(nil), "dfuse.eosio.accounthist.v1.ActionRowAppend")
	proto.RegisterType((*Cursor)(nil), "dfuse.eosio.accounthist.v1.Cursor")
}

func init() {
	proto.RegisterFile("dfuse/eosio/accounthist/v1/accounthist.proto", fileDescriptor_4c22ddb60199ece6)
}

var fileDescriptor_4c22ddb60199ece6 = []byte{
	// 604 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0x95, 0x69, 0x1a, 0xda, 0xe9, 0x47, 0x5a, 0xb7, 0x54, 0x21, 0x5c, 0xc2, 0x5e, 0x88, 0x2a,
	0xba, 0x21, 0xe9, 0x8d, 0x9e, 0xfa, 0x21, 0x15, 0x84, 0xe8, 0x61, 0x53, 0x09, 0x89, 0xcb, 0x6a,
	0xd7, 0x3b, 0x34, 0x56, 0x12, 0x7b, 0xbb, 0xf6, 0xa6, 0xaa, 0x10, 0xe2, 0x2f, 0x20, 0x24, 0xc4,
	0x95, 0x3f, 0xc5, 0xff, 0x41, 0xb6, 0x77, 0xdb, 0x14, 0xd2, 0x40, 0x24, 0x6e, 0x9e, 0x79, 0x6f,
	0x9e, 0x9f, 0xed, 0xf1, 0xc0, 0xf3, 0xe4, 0x43, 0xae, 0xb0, 0x8d, 0x52, 0x71, 0xd9, 0x8e, 0x18,
	0x93, 0xb9, 0xd0, 0x7d, 0xae, 0x74, 0x7b, 0xdc, 0x99, 0x0c, 0xfd, 0x34, 0x93, 0x5a, 0xd2, 0x86,
	0x65, 0xfb, 0x96, 0xed, 0x4f, 0xc2, 0xe3, 0x4e, 0xa3, 0x39, 0xa9, 0xc4, 0x64, 0x82, 0xcc, 0x68,
	0xd8, 0x85, 0xab, 0xf6, 0x3e, 0xc3, 0xe6, 0x29, 0xea, 0x43, 0xa6, 0xb9, 0x14, 0x2a, 0xc0, 0xcb,
	0x1c, 0x95, 0xa6, 0x75, 0x78, 0x58, 0x08, 0xd5, 0x49, 0x93, 0xb4, 0x2a, 0x41, 0x19, 0xd2, 0x6d,
	0x58, 0x1c, 0xf2, 0x11, 0xd7, 0xf5, 0x07, 0x4d, 0xd2, 0x5a, 0x0b, 0x5c, 0x40, 0x5f, 0x42, 0x95,
	0xe5, 0x99, 0x92, 0x59, 0x7d, 0xa1, 0x49, 0x5a, 0x2b, 0x5d, 0xcf, 0xbf, 0xdf, 0x93, 0x7f, 0x6c,
	0x99, 0x41, 0x51, 0xe1, 0xfd, 0x20, 0xb0, 0x73, 0x8a, 0xfa, 0x5c, 0x0e, 0x50, 0xfc, 0xb3, 0x8d,
	0x06, 0x2c, 0x31, 0x29, 0x74, 0x16, 0x31, 0xe7, 0xa4, 0x12, 0xdc, 0xc4, 0xb7, 0x16, 0x17, 0xa6,
	0x5b, 0xac, 0xcc, 0x6d, 0xf1, 0x2b, 0x81, 0x75, 0x67, 0x2d, 0x40, 0x95, 0x4a, 0xa1, 0x70, 0x42,
	0x8e, 0xcc, 0x2b, 0x47, 0x4f, 0x60, 0x35, 0xb2, 0x6a, 0xa1, 0x31, 0x8c, 0xf6, 0x00, 0x2b, 0xdd,
	0xa7, 0x77, 0x14, 0xdc, 0x13, 0x8d, 0x3b, 0xbe, 0xdb, 0xf7, 0xdc, 0x10, 0x83, 0x95, 0xe8, 0x36,
	0xf0, 0xbe, 0x11, 0x58, 0x2e, 0x4c, 0xc9, 0x2b, 0x73, 0x55, 0x63, 0xcc, 0x14, 0x97, 0xc2, 0x1a,
	0x5a, 0x0b, 0xca, 0xf0, 0xff, 0xec, 0x46, 0x5b, 0xb0, 0x31, 0x8c, 0x94, 0x0e, 0x13, 0x1c, 0xa2,
	0xc6, 0x24, 0x54, 0x78, 0x69, 0xef, 0xb7, 0x12, 0xac, 0x9b, 0xfc, 0x89, 0x4b, 0xf7, 0xf0, 0xd2,
	0xfb, 0x49, 0xa0, 0xd6, 0xeb, 0x47, 0x59, 0x72, 0xdc, 0x47, 0x36, 0x48, 0x25, 0x17, 0x9a, 0xfa,
	0xb0, 0xc5, 0x05, 0xd7, 0x3c, 0x1a, 0x86, 0x4a, 0x47, 0x99, 0x0e, 0xe3, 0xa1, 0x64, 0x83, 0xe2,
	0x51, 0x37, 0x0b, 0xa8, 0x67, 0x90, 0x23, 0x03, 0xd0, 0x5d, 0xd8, 0xd4, 0x51, 0x76, 0x81, 0x3a,
	0x54, 0x5a, 0xa6, 0x05, 0xdb, 0xbd, 0x73, 0xcd, 0x01, 0x3d, 0x2d, 0x53, 0xc7, 0xdd, 0x87, 0x1d,
	0xeb, 0xec, 0x2a, 0xe3, 0x5a, 0xa3, 0x70, 0xe4, 0x50, 0xe4, 0xa3, 0xc2, 0xdf, 0x96, 0x41, 0xdf,
	0x39, 0xd0, 0x56, 0x9c, 0xe5, 0x23, 0xda, 0x81, 0x47, 0x53, 0x8a, 0x78, 0x62, 0x9b, 0x63, 0x39,
	0xa0, 0xbf, 0xd7, 0xbc, 0x4e, 0xbc, 0x03, 0xa8, 0xdd, 0x5c, 0xf7, 0x61, 0x9a, 0xa2, 0x48, 0xe6,
	0xb8, 0x94, 0x2f, 0x04, 0xaa, 0xae, 0x0b, 0x66, 0xbc, 0xd4, 0x36, 0x2c, 0x8e, 0xa2, 0x0b, 0xce,
	0xca, 0xbf, 0x65, 0x03, 0xba, 0x01, 0x0b, 0x03, 0xbc, 0xb6, 0xba, 0xab, 0x81, 0x59, 0xd2, 0x27,
	0xb0, 0xac, 0xcc, 0x05, 0xdb, 0x43, 0x2e, 0x5a, 0xee, 0x92, 0x4d, 0x98, 0x93, 0x3d, 0x83, 0x9a,
	0x32, 0xdf, 0x47, 0x30, 0x34, 0x78, 0x8c, 0x59, 0xbd, 0xea, 0x2c, 0x95, 0xe9, 0x33, 0x9b, 0xed,
	0x7e, 0x34, 0x3d, 0x6d, 0xdb, 0xf4, 0x15, 0x57, 0x5a, 0x66, 0xd7, 0x94, 0x03, 0xdc, 0x8e, 0x02,
	0xba, 0x37, 0xab, 0xa3, 0xff, 0x18, 0x19, 0x8d, 0xdd, 0x59, 0xf4, 0xbb, 0x9f, 0xe7, 0x05, 0xe9,
	0x7e, 0x27, 0xb0, 0x53, 0xec, 0x7e, 0x5c, 0xfc, 0xdb, 0xd2, 0xc5, 0x27, 0x78, 0x6c, 0xd5, 0xef,
	0x80, 0xa5, 0xa9, 0xee, 0x5f, 0x4c, 0x4d, 0x99, 0x22, 0xf3, 0x39, 0x3b, 0x7a, 0xfb, 0xfe, 0xcd,
	0x05, 0xd7, 0xfd, 0x3c, 0xf6, 0x99, 0x1c, 0xb5, 0x6d, 0xe5, 0x1e, 0x97, 0xc5, 0xc2, 0xcd, 0xd1,
	0x34, 0x6e, 0xdf, 0x3f, 0xa0, 0x0f, 0xd2, 0x78, 0x22, 0x11, 0x57, 0xed, 0x94, 0xdd, 0xff, 0x15,
	0x00, 0x00, 0xff, 0xff, 0x43, 0xa1, 0x53, 0xf9, 0xd3, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AccountHistoryClient is the client API for AccountHistory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountHistoryClient interface {
	GetActions(ctx context.Context, in *GetActionsRequest, opts ...grpc.CallOption) (AccountHistory_GetActionsClient, error)
}

type accountHistoryClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountHistoryClient(cc grpc.ClientConnInterface) AccountHistoryClient {
	return &accountHistoryClient{cc}
}

func (c *accountHistoryClient) GetActions(ctx context.Context, in *GetActionsRequest, opts ...grpc.CallOption) (AccountHistory_GetActionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AccountHistory_serviceDesc.Streams[0], "/dfuse.eosio.accounthist.v1.AccountHistory/GetActions", opts...)
	if err != nil {
		return nil, err
	}
	x := &accountHistoryGetActionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AccountHistory_GetActionsClient interface {
	Recv() (*ActionResponse, error)
	grpc.ClientStream
}

type accountHistoryGetActionsClient struct {
	grpc.ClientStream
}

func (x *accountHistoryGetActionsClient) Recv() (*ActionResponse, error) {
	m := new(ActionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AccountHistoryServer is the server API for AccountHistory service.
type AccountHistoryServer interface {
	GetActions(*GetActionsRequest, AccountHistory_GetActionsServer) error
}

// UnimplementedAccountHistoryServer can be embedded to have forward compatible implementations.
type UnimplementedAccountHistoryServer struct {
}

func (*UnimplementedAccountHistoryServer) GetActions(req *GetActionsRequest, srv AccountHistory_GetActionsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetActions not implemented")
}

func RegisterAccountHistoryServer(s *grpc.Server, srv AccountHistoryServer) {
	s.RegisterService(&_AccountHistory_serviceDesc, srv)
}

func _AccountHistory_GetActions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetActionsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AccountHistoryServer).GetActions(m, &accountHistoryGetActionsServer{stream})
}

type AccountHistory_GetActionsServer interface {
	Send(*ActionResponse) error
	grpc.ServerStream
}

type accountHistoryGetActionsServer struct {
	grpc.ServerStream
}

func (x *accountHistoryGetActionsServer) Send(m *ActionResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _AccountHistory_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dfuse.eosio.accounthist.v1.AccountHistory",
	HandlerType: (*AccountHistoryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetActions",
			Handler:       _AccountHistory_GetActions_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "dfuse/eosio/accounthist/v1/accounthist.proto",
}

// AccountContractHistoryClient is the client API for AccountContractHistory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountContractHistoryClient interface {
	GetAccountContractActions(ctx context.Context, in *GetTokenActionsRequest, opts ...grpc.CallOption) (AccountContractHistory_GetAccountContractActionsClient, error)
}

type accountContractHistoryClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountContractHistoryClient(cc grpc.ClientConnInterface) AccountContractHistoryClient {
	return &accountContractHistoryClient{cc}
}

func (c *accountContractHistoryClient) GetAccountContractActions(ctx context.Context, in *GetTokenActionsRequest, opts ...grpc.CallOption) (AccountContractHistory_GetAccountContractActionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AccountContractHistory_serviceDesc.Streams[0], "/dfuse.eosio.accounthist.v1.AccountContractHistory/GetAccountContractActions", opts...)
	if err != nil {
		return nil, err
	}
	x := &accountContractHistoryGetAccountContractActionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AccountContractHistory_GetAccountContractActionsClient interface {
	Recv() (*ActionResponse, error)
	grpc.ClientStream
}

type accountContractHistoryGetAccountContractActionsClient struct {
	grpc.ClientStream
}

func (x *accountContractHistoryGetAccountContractActionsClient) Recv() (*ActionResponse, error) {
	m := new(ActionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AccountContractHistoryServer is the server API for AccountContractHistory service.
type AccountContractHistoryServer interface {
	GetAccountContractActions(*GetTokenActionsRequest, AccountContractHistory_GetAccountContractActionsServer) error
}

// UnimplementedAccountContractHistoryServer can be embedded to have forward compatible implementations.
type UnimplementedAccountContractHistoryServer struct {
}

func (*UnimplementedAccountContractHistoryServer) GetAccountContractActions(req *GetTokenActionsRequest, srv AccountContractHistory_GetAccountContractActionsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAccountContractActions not implemented")
}

func RegisterAccountContractHistoryServer(s *grpc.Server, srv AccountContractHistoryServer) {
	s.RegisterService(&_AccountContractHistory_serviceDesc, srv)
}

func _AccountContractHistory_GetAccountContractActions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetTokenActionsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AccountContractHistoryServer).GetAccountContractActions(m, &accountContractHistoryGetAccountContractActionsServer{stream})
}

type AccountContractHistory_GetAccountContractActionsServer interface {
	Send(*ActionResponse) error
	grpc.ServerStream
}

type accountContractHistoryGetAccountContractActionsServer struct {
	grpc.ServerStream
}

func (x *accountContractHistoryGetAccountContractActionsServer) Send(m *ActionResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _AccountContractHistory_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dfuse.eosio.accounthist.v1.AccountContractHistory",
	HandlerType: (*AccountContractHistoryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAccountContractActions",
			Handler:       _AccountContractHistory_GetAccountContractActions_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "dfuse/eosio/accounthist/v1/accounthist.proto",
}
