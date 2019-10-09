package store

import (
	"strings"
	"time"

	"github.com/herdius/herdius-blockchain-api/protobuf"
)

// StatusPending indicates a TX is in pending state
const StatusPending = "pending"

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
	LockedAmount    uint64    `db:"locked_amount"`
	Nonce           uint64    `db:"nonce"`
	Message         string    `db:"message"`
	Sign            string    `db:"sign"`
	Status          string    `db:"status"`
	BlockID         uint64    `db:"block_id"`
	CreationDT      time.Time `db:"created_date"`
	Type            string    `db:"type"`
	Data            string    `db:"data"`
	EthAddress      string    `db:"eth_address"`
	BtcAddress      string    `db:"btc_address"`
	LtcAddress      string    `db:"ltc_address"`
	XtzAddress      string    `db:"xtz_address"`
	BnbAddress      string    `db:"bnb_address"`
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
			Category:     tx.Category,
			Symbol:       tx.Symbol,
			Network:      tx.Network,
			Value:        tx.Value,
			LockedAmount: tx.LockedAmount,
			Nonce:        tx.Nonce,
		},
		Type: tx.Type,
		Data: tx.Data,
	}

	if strings.EqualFold(tx.Type, "REGISTER") {
		externalAddresses := txDetail.Tx.ExternalAddress
		if externalAddresses == nil || len(externalAddresses) == 0 {
			externalAddresses = make(map[string]string)
		}
		externalAddresses["ETH"] = tx.EthAddress
		externalAddresses["BTC"] = tx.BtcAddress
		externalAddresses["LTC"] = tx.LtcAddress
		externalAddresses["XTZ"] = tx.XtzAddress
		externalAddresses["BNB"] = tx.BnbAddress
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
	tx.LockedAmount = txDetail.Tx.Asset.LockedAmount
	tx.Nonce = txDetail.Tx.Asset.Nonce
	tx.CreationDT = time.Now().UTC()
	tx.Sign = txDetail.Tx.Sign
	tx.BlockID = txDetail.BlockId
	tx.Type = txDetail.Tx.Type
	tx.Data = txDetail.Tx.Data
	if strings.EqualFold(txDetail.Tx.Type, "REGISTER") {
		tx.EthAddress = txDetail.Tx.ExternalAddress["ETH"]
		tx.BtcAddress = txDetail.Tx.ExternalAddress["BTC"]
		tx.LtcAddress = txDetail.Tx.ExternalAddress["LTC"]
		tx.XtzAddress = txDetail.Tx.ExternalAddress["XTZ"]
		tx.BnbAddress = txDetail.Tx.ExternalAddress["BNB"]
	}
	return tx
}
