package store

import (
	"time"

	"github.com/herdius/herdius-blockchain-api/protobuf"
)

// Tx ...
type Tx struct {
	ID              string    `db:"id"`
	SenderAddress   string    `db:"sender_address"`
	SenderPubKey    string    `db:"sender_pubkey"`
	ReceiverAddress string    `db:"receiver_address"`
	Category        string    `db:"category"`
	Symbol          string    `db:"symbol"`
	Network         string    `db:"network"`
	Value           uint64    `db:"value"`
	Nonce           uint64    `db:"nonce"`
	Message         string    `db:"message"`
	Sign            string    `db:"sign"`
	Status          string    `db:"status"`
	BlockID         uint64    `db:"block_id"`
	CreationDT      time.Time `db:"created_date"`
}

// ToTxDetailResponse converts a Tx to TxDetailResponse.
func (tx *Tx) ToTxDetailResponse() *protobuf.TxDetailResponse {
	txDetail := &protobuf.TxDetailResponse{}
	txDetail.TxId = tx.ID
	txDetail.BlockId = tx.BlockID
	txDetail.CreationDt = &protobuf.Timestamp{
		Seconds: tx.CreationDT.Unix(),
		Nanos:   tx.CreationDT.UnixNano(),
	}
	txDetail.Tx = &protobuf.Tx{
		SenderAddress:   tx.SenderAddress,
		SenderPubkey:    tx.SenderPubKey,
		RecieverAddress: tx.ReceiverAddress,
		Message:         tx.Message,
		Sign:            tx.Sign,
		Status:          tx.Status,
		Asset: &protobuf.Asset{
			Category: tx.Category,
			Symbol:   tx.Symbol,
			Network:  tx.Network,
			Value:    tx.Value,
			Nonce:    tx.Nonce,
		},
	}

	return txDetail
}

// FromTxDetailResponse creates a Tx from TxDetailResponse
func FromTxDetailResponse(txDetail *protobuf.TxDetailResponse) *Tx {
	tx := &Tx{}
	tx.ID = txDetail.TxId
	tx.Status = txDetail.Tx.Status
	tx.Message = txDetail.Tx.Message
	tx.SenderAddress = txDetail.Tx.SenderAddress
	tx.SenderPubKey = txDetail.Tx.SenderPubkey
	tx.ReceiverAddress = txDetail.Tx.RecieverAddress
	tx.Category = txDetail.Tx.Asset.Category
	tx.Symbol = txDetail.Tx.Asset.Symbol
	tx.Network = txDetail.Tx.Asset.Network
	tx.Value = txDetail.Tx.Asset.Value
	tx.Nonce = txDetail.Tx.Asset.Nonce
	tx.CreationDT = time.Now().UTC()
	tx.Sign = txDetail.Tx.Sign
	tx.BlockID = txDetail.BlockId

	return tx
}
