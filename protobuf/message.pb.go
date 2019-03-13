// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/message.proto

package protobuf

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

type Timestamp struct {
	Seconds              int64    `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	Nanos                int64    `protobuf:"varint,2,opt,name=nanos,proto3" json:"nanos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Timestamp) Reset()         { *m = Timestamp{} }
func (m *Timestamp) String() string { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()    {}
func (*Timestamp) Descriptor() ([]byte, []int) {
	return fileDescriptor_8368f5d77b0b9b7b, []int{0}
}

func (m *Timestamp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Timestamp.Unmarshal(m, b)
}
func (m *Timestamp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Timestamp.Marshal(b, m, deterministic)
}
func (m *Timestamp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Timestamp.Merge(m, src)
}
func (m *Timestamp) XXX_Size() int {
	return xxx_messageInfo_Timestamp.Size(m)
}
func (m *Timestamp) XXX_DiscardUnknown() {
	xxx_messageInfo_Timestamp.DiscardUnknown(m)
}

var xxx_messageInfo_Timestamp proto.InternalMessageInfo

func (m *Timestamp) GetSeconds() int64 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *Timestamp) GetNanos() int64 {
	if m != nil {
		return m.Nanos
	}
	return 0
}

type BlockHeightRequest struct {
	BlockHeight          int64    `protobuf:"varint,1,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockHeightRequest) Reset()         { *m = BlockHeightRequest{} }
func (m *BlockHeightRequest) String() string { return proto.CompactTextString(m) }
func (*BlockHeightRequest) ProtoMessage()    {}
func (*BlockHeightRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8368f5d77b0b9b7b, []int{1}
}

func (m *BlockHeightRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockHeightRequest.Unmarshal(m, b)
}
func (m *BlockHeightRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockHeightRequest.Marshal(b, m, deterministic)
}
func (m *BlockHeightRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeightRequest.Merge(m, src)
}
func (m *BlockHeightRequest) XXX_Size() int {
	return xxx_messageInfo_BlockHeightRequest.Size(m)
}
func (m *BlockHeightRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeightRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeightRequest proto.InternalMessageInfo

func (m *BlockHeightRequest) GetBlockHeight() int64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

type BlockResponse struct {
	BlockHeight int64 `protobuf:"varint,1,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	// Time of block intialization
	Time         *Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	Transactions int32      `protobuf:"varint,3,opt,name=transactions,proto3" json:"transactions,omitempty"`
	// Supervisor herdius token address who created the block
	SupervisorAddress    string   `protobuf:"bytes,4,opt,name=supervisor_address,json=supervisorAddress,proto3" json:"supervisor_address,omitempty"`
	StateRoot            []byte   `protobuf:"bytes,5,opt,name=state_root,json=stateRoot,proto3" json:"state_root,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockResponse) Reset()         { *m = BlockResponse{} }
func (m *BlockResponse) String() string { return proto.CompactTextString(m) }
func (*BlockResponse) ProtoMessage()    {}
func (*BlockResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8368f5d77b0b9b7b, []int{2}
}

func (m *BlockResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockResponse.Unmarshal(m, b)
}
func (m *BlockResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockResponse.Marshal(b, m, deterministic)
}
func (m *BlockResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockResponse.Merge(m, src)
}
func (m *BlockResponse) XXX_Size() int {
	return xxx_messageInfo_BlockResponse.Size(m)
}
func (m *BlockResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BlockResponse proto.InternalMessageInfo

func (m *BlockResponse) GetBlockHeight() int64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *BlockResponse) GetTime() *Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *BlockResponse) GetTransactions() int32 {
	if m != nil {
		return m.Transactions
	}
	return 0
}

func (m *BlockResponse) GetSupervisorAddress() string {
	if m != nil {
		return m.SupervisorAddress
	}
	return ""
}

func (m *BlockResponse) GetStateRoot() []byte {
	if m != nil {
		return m.StateRoot
	}
	return nil
}

type AccountRequest struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountRequest) Reset()         { *m = AccountRequest{} }
func (m *AccountRequest) String() string { return proto.CompactTextString(m) }
func (*AccountRequest) ProtoMessage()    {}
func (*AccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8368f5d77b0b9b7b, []int{3}
}

func (m *AccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountRequest.Unmarshal(m, b)
}
func (m *AccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountRequest.Marshal(b, m, deterministic)
}
func (m *AccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountRequest.Merge(m, src)
}
func (m *AccountRequest) XXX_Size() int {
	return xxx_messageInfo_AccountRequest.Size(m)
}
func (m *AccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountRequest proto.InternalMessageInfo

func (m *AccountRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type AccountResponse struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Nonce                int64    `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Balance              int64    `protobuf:"varint,3,opt,name=balance,proto3" json:"balance,omitempty"`
	StorageRoot          string   `protobuf:"bytes,4,opt,name=storage_root,json=storageRoot,proto3" json:"storage_root,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountResponse) Reset()         { *m = AccountResponse{} }
func (m *AccountResponse) String() string { return proto.CompactTextString(m) }
func (*AccountResponse) ProtoMessage()    {}
func (*AccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8368f5d77b0b9b7b, []int{4}
}

func (m *AccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountResponse.Unmarshal(m, b)
}
func (m *AccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountResponse.Marshal(b, m, deterministic)
}
func (m *AccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountResponse.Merge(m, src)
}
func (m *AccountResponse) XXX_Size() int {
	return xxx_messageInfo_AccountResponse.Size(m)
}
func (m *AccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccountResponse proto.InternalMessageInfo

func (m *AccountResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AccountResponse) GetNonce() int64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *AccountResponse) GetBalance() int64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *AccountResponse) GetStorageRoot() string {
	if m != nil {
		return m.StorageRoot
	}
	return ""
}

func init() {
	proto.RegisterType((*Timestamp)(nil), "protobuf.Timestamp")
	proto.RegisterType((*BlockHeightRequest)(nil), "protobuf.BlockHeightRequest")
	proto.RegisterType((*BlockResponse)(nil), "protobuf.BlockResponse")
	proto.RegisterType((*AccountRequest)(nil), "protobuf.AccountRequest")
	proto.RegisterType((*AccountResponse)(nil), "protobuf.AccountResponse")
}

func init() { proto.RegisterFile("protobuf/message.proto", fileDescriptor_8368f5d77b0b9b7b) }

var fileDescriptor_8368f5d77b0b9b7b = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0x65, 0xda, 0x52, 0x7a, 0x0d, 0xa0, 0x1a, 0x84, 0xb2, 0x20, 0x85, 0x2c, 0x44, 0x48,
	0x0d, 0x12, 0x0c, 0x0c, 0x4c, 0xed, 0xd4, 0xb1, 0xb2, 0x98, 0x58, 0x2a, 0x27, 0x31, 0x6d, 0x44,
	0xe3, 0x0b, 0x3e, 0x87, 0x91, 0x6f, 0xc8, 0x77, 0x42, 0x71, 0xfe, 0x54, 0x0c, 0x48, 0x6c, 0x79,
	0x3f, 0x3f, 0x9f, 0xee, 0xbd, 0x18, 0xae, 0x4a, 0x83, 0x16, 0x93, 0xea, 0xed, 0xbe, 0x50, 0x44,
	0x72, 0xab, 0x62, 0x07, 0xf8, 0x49, 0xc7, 0xc3, 0x67, 0x98, 0xbc, 0xe4, 0x85, 0x22, 0x2b, 0x8b,
	0x92, 0xfb, 0x30, 0x26, 0x95, 0xa2, 0xce, 0xc8, 0x67, 0x01, 0x8b, 0x06, 0xa2, 0x93, 0xfc, 0x12,
	0x46, 0x5a, 0x6a, 0x24, 0xff, 0xc8, 0xf1, 0x46, 0x84, 0x4f, 0xc0, 0x97, 0x7b, 0x4c, 0xdf, 0x57,
	0x2a, 0xdf, 0xee, 0xac, 0x50, 0x1f, 0x95, 0x22, 0xcb, 0x6f, 0xc0, 0x4b, 0x6a, 0xba, 0xd9, 0x39,
	0xdc, 0x8e, 0x9a, 0x26, 0x07, 0x67, 0xf8, 0xcd, 0xe0, 0xd4, 0xdd, 0x14, 0x8a, 0x4a, 0xd4, 0xa4,
	0xfe, 0x71, 0x89, 0xdf, 0xc2, 0xd0, 0xe6, 0x85, 0x72, 0x2b, 0x4c, 0x1f, 0x2e, 0xe2, 0x2e, 0x43,
	0xdc, 0x07, 0x10, 0xce, 0xc0, 0x43, 0xf0, 0xac, 0x91, 0x9a, 0x64, 0x6a, 0x73, 0xd4, 0xe4, 0x0f,
	0x02, 0x16, 0x8d, 0xc4, 0x2f, 0xc6, 0xe7, 0xc0, 0xa9, 0x2a, 0x95, 0xf9, 0xcc, 0x09, 0xcd, 0x46,
	0x66, 0x99, 0x51, 0x44, 0xfe, 0x30, 0x60, 0xd1, 0x44, 0xcc, 0x0e, 0x27, 0x8b, 0xe6, 0x80, 0x5f,
	0x03, 0x90, 0x95, 0x56, 0x6d, 0x0c, 0xa2, 0xf5, 0x47, 0x01, 0x8b, 0x3c, 0x31, 0x71, 0x44, 0x20,
	0xda, 0xf0, 0x0e, 0xce, 0x16, 0x69, 0x8a, 0x95, 0xee, 0x4b, 0xf0, 0x61, 0xdc, 0x0d, 0x65, 0x6e,
	0x68, 0x27, 0xc3, 0x2f, 0x38, 0xef, 0xbd, 0x6d, 0xf8, 0x3f, 0xcd, 0xae, 0x77, 0xd4, 0xa9, 0xea,
	0x7b, 0xaf, 0x45, 0xed, 0x4f, 0xe4, 0x5e, 0xd6, 0x7c, 0xd0, 0xfc, 0xa7, 0x56, 0xd6, 0x35, 0x92,
	0x45, 0x23, 0xb7, 0xed, 0xa6, 0x4d, 0xa0, 0x69, 0xcb, 0xea, 0x5d, 0x97, 0x73, 0x98, 0xa5, 0x58,
	0xc4, 0x3b, 0x65, 0xb2, 0xbc, 0xa2, 0xa6, 0xc5, 0xa5, 0xb7, 0x6a, 0xe4, 0xba, 0x56, 0x6b, 0xf6,
	0xda, 0x3f, 0x90, 0xe4, 0xd8, 0x7d, 0x3d, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xde, 0xe5, 0xf1,
	0xc7, 0x4b, 0x02, 0x00, 0x00,
}
