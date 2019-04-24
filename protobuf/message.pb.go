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
	BlockHeight          uint64   `protobuf:"varint,1,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
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

func (m *BlockHeightRequest) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

type BlockResponse struct {
	BlockHeight uint64 `protobuf:"varint,1,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	// Time of block intialization
	Time         *Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	Transactions uint64     `protobuf:"varint,3,opt,name=transactions,proto3" json:"transactions,omitempty"`
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

func (m *BlockResponse) GetBlockHeight() uint64 {
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

func (m *BlockResponse) GetTransactions() uint64 {
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
	Address              string            `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Nonce                uint64            `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	StorageRoot          string            `protobuf:"bytes,3,opt,name=storage_root,json=storageRoot,proto3" json:"storage_root,omitempty"`
	PublicKey            string            `protobuf:"bytes,4,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	Balance              uint64            `protobuf:"varint,5,opt,name=balance,proto3" json:"balance,omitempty"`
	Balances             map[string]uint64 `protobuf:"bytes,6,rep,name=balances,proto3" json:"balances,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
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

func (m *AccountResponse) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *AccountResponse) GetStorageRoot() string {
	if m != nil {
		return m.StorageRoot
	}
	return ""
}

func (m *AccountResponse) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

func (m *AccountResponse) GetBalance() uint64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *AccountResponse) GetBalances() map[string]uint64 {
	if m != nil {
		return m.Balances
	}
	return nil
}

type Tx struct {
	SenderAddress        string   `protobuf:"bytes,1,opt,name=sender_address,json=senderAddress,proto3" json:"sender_address,omitempty"`
	SenderPubkey         string   `protobuf:"bytes,2,opt,name=sender_pubkey,json=senderPubkey,proto3" json:"sender_pubkey,omitempty"`
	RecieverAddress      string   `protobuf:"bytes,3,opt,name=reciever_address,json=recieverAddress,proto3" json:"reciever_address,omitempty"`
	Asset                *Asset   `protobuf:"bytes,4,opt,name=asset,proto3" json:"asset,omitempty"`
	Message              string   `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	Sign                 string   `protobuf:"bytes,6,opt,name=sign,proto3" json:"sign,omitempty"`
	Type                 string   `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
	Status               string   `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tx) Reset()         { *m = Tx{} }
func (m *Tx) String() string { return proto.CompactTextString(m) }
func (*Tx) ProtoMessage()    {}
func (*Tx) Descriptor() ([]byte, []int) {
	return fileDescriptor_8368f5d77b0b9b7b, []int{5}
}

func (m *Tx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tx.Unmarshal(m, b)
}
func (m *Tx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tx.Marshal(b, m, deterministic)
}
func (m *Tx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tx.Merge(m, src)
}
func (m *Tx) XXX_Size() int {
	return xxx_messageInfo_Tx.Size(m)
}
func (m *Tx) XXX_DiscardUnknown() {
	xxx_messageInfo_Tx.DiscardUnknown(m)
}

var xxx_messageInfo_Tx proto.InternalMessageInfo

func (m *Tx) GetSenderAddress() string {
	if m != nil {
		return m.SenderAddress
	}
	return ""
}

func (m *Tx) GetSenderPubkey() string {
	if m != nil {
		return m.SenderPubkey
	}
	return ""
}

func (m *Tx) GetRecieverAddress() string {
	if m != nil {
		return m.RecieverAddress
	}
	return ""
}

func (m *Tx) GetAsset() *Asset {
	if m != nil {
		return m.Asset
	}
	return nil
}

func (m *Tx) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Tx) GetSign() string {
	if m != nil {
		return m.Sign
	}
	return ""
}

func (m *Tx) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Tx) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type Asset struct {
	Category             string   `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	Symbol               string   `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Network              string   `protobuf:"bytes,3,opt,name=network,proto3" json:"network,omitempty"`
	Value                uint64   `protobuf:"varint,4,opt,name=value,proto3" json:"value,omitempty"`
	Fee                  uint64   `protobuf:"varint,5,opt,name=fee,proto3" json:"fee,omitempty"`
	Nonce                uint64   `protobuf:"varint,6,opt,name=nonce,proto3" json:"nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Asset) Reset()         { *m = Asset{} }
func (m *Asset) String() string { return proto.CompactTextString(m) }
func (*Asset) ProtoMessage()    {}
func (*Asset) Descriptor() ([]byte, []int) {
	return fileDescriptor_8368f5d77b0b9b7b, []int{6}
}

func (m *Asset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Asset.Unmarshal(m, b)
}
func (m *Asset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Asset.Marshal(b, m, deterministic)
}
func (m *Asset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Asset.Merge(m, src)
}
func (m *Asset) XXX_Size() int {
	return xxx_messageInfo_Asset.Size(m)
}
func (m *Asset) XXX_DiscardUnknown() {
	xxx_messageInfo_Asset.DiscardUnknown(m)
}

var xxx_messageInfo_Asset proto.InternalMessageInfo

func (m *Asset) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *Asset) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *Asset) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *Asset) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Asset) GetFee() uint64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func (m *Asset) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

type TxRequest struct {
	Tx                   *Tx      `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxRequest) Reset()         { *m = TxRequest{} }
func (m *TxRequest) String() string { return proto.CompactTextString(m) }
func (*TxRequest) ProtoMessage()    {}
func (*TxRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8368f5d77b0b9b7b, []int{7}
}

func (m *TxRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxRequest.Unmarshal(m, b)
}
func (m *TxRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxRequest.Marshal(b, m, deterministic)
}
func (m *TxRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxRequest.Merge(m, src)
}
func (m *TxRequest) XXX_Size() int {
	return xxx_messageInfo_TxRequest.Size(m)
}
func (m *TxRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TxRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TxRequest proto.InternalMessageInfo

func (m *TxRequest) GetTx() *Tx {
	if m != nil {
		return m.Tx
	}
	return nil
}

type TxResponse struct {
	TxId                 string   `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	Pending              int64    `protobuf:"varint,2,opt,name=pending,proto3" json:"pending,omitempty"`
	Queued               int64    `protobuf:"varint,3,opt,name=queued,proto3" json:"queued,omitempty"`
	Status               string   `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Message              string   `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxResponse) Reset()         { *m = TxResponse{} }
func (m *TxResponse) String() string { return proto.CompactTextString(m) }
func (*TxResponse) ProtoMessage()    {}
func (*TxResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8368f5d77b0b9b7b, []int{8}
}

func (m *TxResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxResponse.Unmarshal(m, b)
}
func (m *TxResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxResponse.Marshal(b, m, deterministic)
}
func (m *TxResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxResponse.Merge(m, src)
}
func (m *TxResponse) XXX_Size() int {
	return xxx_messageInfo_TxResponse.Size(m)
}
func (m *TxResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TxResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TxResponse proto.InternalMessageInfo

func (m *TxResponse) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

func (m *TxResponse) GetPending() int64 {
	if m != nil {
		return m.Pending
	}
	return 0
}

func (m *TxResponse) GetQueued() int64 {
	if m != nil {
		return m.Queued
	}
	return 0
}

func (m *TxResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *TxResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type AccountRegisterRequest struct {
	SenderPubkey         string   `protobuf:"bytes,1,opt,name=sender_pubkey,json=senderPubkey,proto3" json:"sender_pubkey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountRegisterRequest) Reset()         { *m = AccountRegisterRequest{} }
func (m *AccountRegisterRequest) String() string { return proto.CompactTextString(m) }
func (*AccountRegisterRequest) ProtoMessage()    {}
func (*AccountRegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8368f5d77b0b9b7b, []int{9}
}

func (m *AccountRegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountRegisterRequest.Unmarshal(m, b)
}
func (m *AccountRegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountRegisterRequest.Marshal(b, m, deterministic)
}
func (m *AccountRegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountRegisterRequest.Merge(m, src)
}
func (m *AccountRegisterRequest) XXX_Size() int {
	return xxx_messageInfo_AccountRegisterRequest.Size(m)
}
func (m *AccountRegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountRegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountRegisterRequest proto.InternalMessageInfo

func (m *AccountRegisterRequest) GetSenderPubkey() string {
	if m != nil {
		return m.SenderPubkey
	}
	return ""
}

type AccountRegisterResponse struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountRegisterResponse) Reset()         { *m = AccountRegisterResponse{} }
func (m *AccountRegisterResponse) String() string { return proto.CompactTextString(m) }
func (*AccountRegisterResponse) ProtoMessage()    {}
func (*AccountRegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8368f5d77b0b9b7b, []int{10}
}

func (m *AccountRegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountRegisterResponse.Unmarshal(m, b)
}
func (m *AccountRegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountRegisterResponse.Marshal(b, m, deterministic)
}
func (m *AccountRegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountRegisterResponse.Merge(m, src)
}
func (m *AccountRegisterResponse) XXX_Size() int {
	return xxx_messageInfo_AccountRegisterResponse.Size(m)
}
func (m *AccountRegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountRegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccountRegisterResponse proto.InternalMessageInfo

func (m *AccountRegisterResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

// Send request to retrieve transaction committed in herdius blockchain
type TxDetailRequest struct {
	TxId                 string   `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxDetailRequest) Reset()         { *m = TxDetailRequest{} }
func (m *TxDetailRequest) String() string { return proto.CompactTextString(m) }
func (*TxDetailRequest) ProtoMessage()    {}
func (*TxDetailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8368f5d77b0b9b7b, []int{11}
}

func (m *TxDetailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxDetailRequest.Unmarshal(m, b)
}
func (m *TxDetailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxDetailRequest.Marshal(b, m, deterministic)
}
func (m *TxDetailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxDetailRequest.Merge(m, src)
}
func (m *TxDetailRequest) XXX_Size() int {
	return xxx_messageInfo_TxDetailRequest.Size(m)
}
func (m *TxDetailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TxDetailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TxDetailRequest proto.InternalMessageInfo

func (m *TxDetailRequest) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

// Transaction detail response from herdius blockchain
type TxDetailResponse struct {
	TxId                 string     `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	Tx                   *Tx        `protobuf:"bytes,2,opt,name=tx,proto3" json:"tx,omitempty"`
	CreationDt           *Timestamp `protobuf:"bytes,3,opt,name=creationDt,proto3" json:"creationDt,omitempty"`
	BlockId              uint64     `protobuf:"varint,4,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TxDetailResponse) Reset()         { *m = TxDetailResponse{} }
func (m *TxDetailResponse) String() string { return proto.CompactTextString(m) }
func (*TxDetailResponse) ProtoMessage()    {}
func (*TxDetailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8368f5d77b0b9b7b, []int{12}
}

func (m *TxDetailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxDetailResponse.Unmarshal(m, b)
}
func (m *TxDetailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxDetailResponse.Marshal(b, m, deterministic)
}
func (m *TxDetailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxDetailResponse.Merge(m, src)
}
func (m *TxDetailResponse) XXX_Size() int {
	return xxx_messageInfo_TxDetailResponse.Size(m)
}
func (m *TxDetailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TxDetailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TxDetailResponse proto.InternalMessageInfo

func (m *TxDetailResponse) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

func (m *TxDetailResponse) GetTx() *Tx {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *TxDetailResponse) GetCreationDt() *Timestamp {
	if m != nil {
		return m.CreationDt
	}
	return nil
}

func (m *TxDetailResponse) GetBlockId() uint64 {
	if m != nil {
		return m.BlockId
	}
	return 0
}

func init() {
	proto.RegisterType((*Timestamp)(nil), "protobuf.Timestamp")
	proto.RegisterType((*BlockHeightRequest)(nil), "protobuf.BlockHeightRequest")
	proto.RegisterType((*BlockResponse)(nil), "protobuf.BlockResponse")
	proto.RegisterType((*AccountRequest)(nil), "protobuf.AccountRequest")
	proto.RegisterType((*AccountResponse)(nil), "protobuf.AccountResponse")
	proto.RegisterMapType((map[string]uint64)(nil), "protobuf.AccountResponse.BalancesEntry")
	proto.RegisterType((*Tx)(nil), "protobuf.Tx")
	proto.RegisterType((*Asset)(nil), "protobuf.Asset")
	proto.RegisterType((*TxRequest)(nil), "protobuf.TxRequest")
	proto.RegisterType((*TxResponse)(nil), "protobuf.TxResponse")
	proto.RegisterType((*AccountRegisterRequest)(nil), "protobuf.AccountRegisterRequest")
	proto.RegisterType((*AccountRegisterResponse)(nil), "protobuf.AccountRegisterResponse")
	proto.RegisterType((*TxDetailRequest)(nil), "protobuf.TxDetailRequest")
	proto.RegisterType((*TxDetailResponse)(nil), "protobuf.TxDetailResponse")
}

func init() { proto.RegisterFile("protobuf/message.proto", fileDescriptor_8368f5d77b0b9b7b) }

var fileDescriptor_8368f5d77b0b9b7b = []byte{
	// 756 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xcd, 0x4e, 0xeb, 0x46,
	0x14, 0x96, 0x1d, 0xe7, 0xef, 0x24, 0x10, 0x18, 0x2a, 0xea, 0xa2, 0x56, 0x4a, 0x5d, 0x51, 0x42,
	0x25, 0x52, 0x35, 0x2c, 0x5a, 0x15, 0x75, 0x41, 0x4a, 0x25, 0x50, 0x37, 0xc8, 0xca, 0xaa, 0x9b,
	0xc8, 0x3f, 0x87, 0x60, 0x91, 0x78, 0x8c, 0x67, 0x4c, 0x9d, 0x07, 0xe8, 0x13, 0x54, 0x5d, 0xf7,
	0x89, 0xee, 0x1b, 0xdd, 0xc5, 0xd5, 0xfc, 0xd9, 0xe6, 0x02, 0x57, 0x77, 0x37, 0xdf, 0x37, 0xe7,
	0x9c, 0x39, 0xdf, 0x37, 0xc7, 0x63, 0x38, 0xcc, 0x72, 0xca, 0x69, 0x58, 0xdc, 0xfd, 0xb8, 0x41,
	0xc6, 0x82, 0x15, 0x4e, 0x25, 0x41, 0x7a, 0x86, 0xf7, 0x2e, 0xa0, 0xbf, 0x48, 0x36, 0xc8, 0x78,
	0xb0, 0xc9, 0x88, 0x0b, 0x5d, 0x86, 0x11, 0x4d, 0x63, 0xe6, 0x5a, 0x63, 0x6b, 0xd2, 0xf2, 0x0d,
	0x24, 0x5f, 0x40, 0x3b, 0x0d, 0x52, 0xca, 0x5c, 0x5b, 0xf2, 0x0a, 0x78, 0x3f, 0x03, 0x99, 0xaf,
	0x69, 0xf4, 0x70, 0x8d, 0xc9, 0xea, 0x9e, 0xfb, 0xf8, 0x58, 0x20, 0xe3, 0xe4, 0x5b, 0x18, 0x86,
	0x82, 0x5d, 0xde, 0x4b, 0x5a, 0x96, 0x72, 0xfc, 0x41, 0x58, 0x47, 0x7a, 0xef, 0x2c, 0xd8, 0x91,
	0x99, 0x3e, 0xb2, 0x8c, 0xa6, 0x0c, 0x3f, 0x23, 0x89, 0x9c, 0x80, 0xc3, 0x93, 0x0d, 0xca, 0x16,
	0x06, 0xb3, 0x83, 0xa9, 0xd1, 0x30, 0xad, 0x04, 0xf8, 0x32, 0x80, 0x78, 0x30, 0xe4, 0x79, 0x90,
	0xb2, 0x20, 0xe2, 0x09, 0x4d, 0x99, 0xdb, 0x92, 0xb5, 0x9e, 0x71, 0xe4, 0x0c, 0x08, 0x2b, 0x32,
	0xcc, 0x9f, 0x12, 0x46, 0xf3, 0x65, 0x10, 0xc7, 0x39, 0x32, 0xe6, 0x3a, 0x63, 0x6b, 0xd2, 0xf7,
	0xf7, 0xeb, 0x9d, 0x4b, 0xb5, 0x41, 0xbe, 0x01, 0x60, 0x3c, 0xe0, 0xb8, 0xcc, 0x29, 0xe5, 0x6e,
	0x7b, 0x6c, 0x4d, 0x86, 0x7e, 0x5f, 0x32, 0x3e, 0xa5, 0xdc, 0xfb, 0x01, 0x76, 0x2f, 0xa3, 0x88,
	0x16, 0x69, 0x65, 0x82, 0x0b, 0x5d, 0x53, 0xd4, 0x92, 0x45, 0x0d, 0xf4, 0xfe, 0xb7, 0x61, 0x54,
	0x05, 0x6b, 0xf5, 0x6f, 0x46, 0x4b, 0xe3, 0x69, 0x1a, 0x29, 0xd5, 0x8e, 0xaf, 0x80, 0x70, 0x8b,
	0x71, 0x9a, 0x07, 0x2b, 0xdd, 0x50, 0x4b, 0x26, 0x0d, 0x34, 0x27, 0x5a, 0x12, 0x1d, 0x67, 0x45,
	0xb8, 0x4e, 0xa2, 0xe5, 0x03, 0x6e, 0xb5, 0xb0, 0xbe, 0x62, 0xfe, 0xc4, 0xad, 0x38, 0x31, 0x0c,
	0xd6, 0x81, 0xa8, 0xdc, 0x96, 0x95, 0x0d, 0x24, 0xbf, 0x43, 0x4f, 0x2f, 0x99, 0xdb, 0x19, 0xb7,
	0x26, 0x83, 0xd9, 0x49, 0x6d, 0xf5, 0x47, 0x8d, 0x4f, 0xe7, 0x3a, 0xf2, 0x8f, 0x94, 0xe7, 0x5b,
	0xbf, 0x4a, 0x3c, 0xba, 0x80, 0x9d, 0x67, 0x5b, 0x64, 0x0f, 0x5a, 0xa2, 0x0f, 0xa5, 0x4e, 0x2c,
	0x85, 0xb2, 0xa7, 0x60, 0x5d, 0x54, 0xca, 0x24, 0xf8, 0xd5, 0xfe, 0xc5, 0xf2, 0xde, 0x5b, 0x60,
	0x2f, 0x4a, 0x72, 0x0c, 0xbb, 0x0c, 0xd3, 0x18, 0xeb, 0xeb, 0x51, 0xd9, 0x3b, 0x8a, 0x35, 0x57,
	0xf3, 0x1d, 0x68, 0x62, 0x99, 0x15, 0xa1, 0x38, 0xc3, 0x96, 0x51, 0x43, 0x45, 0xde, 0x4a, 0x8e,
	0x9c, 0xc2, 0x5e, 0x8e, 0x51, 0x82, 0x4f, 0x8d, 0x6a, 0xca, 0xb4, 0x91, 0xe1, 0x4d, 0xbd, 0x63,
	0x68, 0x07, 0x8c, 0x21, 0x97, 0x9e, 0x0d, 0x66, 0xa3, 0x86, 0x78, 0x41, 0xfb, 0x6a, 0x57, 0x18,
	0xa8, 0xbf, 0x29, 0x69, 0x60, 0xdf, 0x37, 0x90, 0x10, 0x70, 0x58, 0xb2, 0x4a, 0xdd, 0x8e, 0xa4,
	0xe5, 0x5a, 0x70, 0x7c, 0x9b, 0xa1, 0xdb, 0x55, 0x9c, 0x58, 0x93, 0x43, 0xe8, 0x88, 0x09, 0x2a,
	0x98, 0xdb, 0x93, 0xac, 0x46, 0xde, 0xbf, 0x16, 0xb4, 0xe5, 0x51, 0xe4, 0x08, 0x7a, 0x51, 0xc0,
	0x71, 0x45, 0x73, 0xe3, 0x5c, 0x85, 0x65, 0xf6, 0x76, 0x13, 0xd2, 0xb5, 0xd6, 0xab, 0x91, 0xe8,
	0x2b, 0x45, 0xfe, 0x37, 0xcd, 0x1f, 0xb4, 0x40, 0x03, 0x6b, 0xc3, 0x9d, 0x86, 0xe1, 0xe2, 0x62,
	0xee, 0xd0, 0x0c, 0x81, 0x58, 0xd6, 0x23, 0xd7, 0x69, 0x8c, 0x9c, 0x77, 0x0a, 0xfd, 0x45, 0x69,
	0xa6, 0xfb, 0x6b, 0xb0, 0x79, 0x29, 0x5b, 0x1a, 0xcc, 0x86, 0x8d, 0x0f, 0xb1, 0xf4, 0x6d, 0x5e,
	0x7a, 0xff, 0x58, 0x00, 0x22, 0x56, 0x0f, 0xf7, 0x01, 0xb4, 0x79, 0xb9, 0x4c, 0x62, 0x2d, 0xc1,
	0xe1, 0xe5, 0x4d, 0x2c, 0xda, 0xcc, 0x30, 0x8d, 0x93, 0x74, 0xa5, 0x9f, 0x14, 0x03, 0x85, 0xb0,
	0xc7, 0x02, 0x0b, 0x8c, 0x65, 0xff, 0x2d, 0x5f, 0xa3, 0x86, 0x5d, 0x4e, 0xd3, 0xae, 0xb7, 0x2f,
	0xc2, 0xfb, 0x0d, 0x0e, 0xab, 0x79, 0x5d, 0x25, 0x8c, 0x63, 0x6e, 0xfa, 0x7f, 0x31, 0x33, 0xd6,
	0xcb, 0x99, 0xf1, 0x7e, 0x82, 0x2f, 0x5f, 0xa4, 0x6b, 0x49, 0x75, 0x2f, 0x22, 0xb1, 0x57, 0x5d,
	0xdd, 0xf7, 0x30, 0x5a, 0x94, 0x57, 0xc8, 0x83, 0x64, 0x6d, 0x8e, 0x7a, 0x4d, 0xbd, 0xf7, 0x9f,
	0x05, 0x7b, 0x75, 0xe0, 0xa7, 0x7c, 0x52, 0x4e, 0xdb, 0xaf, 0x3b, 0x4d, 0xce, 0x01, 0xa2, 0x1c,
	0x03, 0xf1, 0xa4, 0x5d, 0xa9, 0x57, 0xe0, 0x8d, 0x87, 0xb1, 0x11, 0x46, 0xbe, 0x82, 0x9e, 0x7a,
	0x6a, 0x93, 0x58, 0x8f, 0x42, 0x57, 0xe2, 0x9b, 0x78, 0x7e, 0x06, 0xfb, 0x11, 0xdd, 0x4c, 0xef,
	0x31, 0x8f, 0x93, 0x82, 0xa9, 0x42, 0xf3, 0xe1, 0xb5, 0x82, 0xb7, 0x02, 0xdd, 0x5a, 0x7f, 0x55,
	0x3f, 0x8f, 0xb0, 0x23, 0x57, 0xe7, 0x1f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x7d, 0x74, 0x35, 0xa6,
	0x67, 0x06, 0x00, 0x00,
}
