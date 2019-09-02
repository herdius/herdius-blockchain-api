package postgres

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/alecthomas/assert"
	"github.com/herdius/herdius-blockchain-api/store"
)

var s *Store

func TestMain(m *testing.M) {
	var err error
	s, err = NewStore("user=postgres host=127.0.0.1 port=5432 dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if _, err := s.DB().Exec("DELETE FROM transaction"); err != nil {
		log.Fatalf("Failed to cleanup db: %v", err)
	}

	os.Exit(m.Run())
}

func TestTx(t *testing.T) {
	now := time.Now().UTC()
	tx := &store.Tx{
		ID:              "id1",
		SenderAddress:   "sa",
		SenderPubKey:    "sp",
		ReceiverAddress: "ra",
		BlockID:         1,
		Category:        "category",
		Message:         "message",
		Nonce:           18,
		Sign:            "sign",
		Status:          "success",
		Symbol:          "ETH",
		Value:           1,
		LockedAmount:    0,
		Network:         "herdius",
		CreationDT:      now,
		Type:            "update",
	}

	assert.NoError(t, s.Save(tx))

	got, err := s.Get(tx.ID)
	assert.NoError(t, err)
	// Ignore time compare issue
	got.CreationDT = now
	assert.Equal(t, tx, got)

	tx.ID = "id2"
	tx.Symbol = "BTC"
	assert.NoError(t, s.Save(tx))

	txs, err := s.GetBySender(tx.SenderAddress)
	assert.NoError(t, err)
	assert.Len(t, txs, 2)

	tx.Status = "failed"
	assert.NoError(t, s.Update(tx))
	updated, err := s.Get(tx.ID)
	assert.NoError(t, err)
	assert.Equal(t, tx.Status, updated.Status)

	txs, err = s.GetByAssetAndSender("eth", tx.SenderAddress)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(txs))
	assert.Equal(t, "ETH", txs[0].Symbol)

	blockID, err := s.GetLatestBlockID()
	assert.NoError(t, err)
	assert.Equal(t, tx.BlockID, blockID)

	tx.ID = "id3"
	tx.BlockID = 10
	assert.NoError(t, s.Save(tx))

	blockID, err = s.GetLatestBlockID()
	assert.NoError(t, err)
	assert.Equal(t, tx.BlockID, blockID)
}

func TestLastSyncBlock(t *testing.T) {
	assert.NoError(t, s.SaveLastSyncBlockID(10))
	assert.Equal(t, uint64(10), s.GetLastSyncBlockID())
}
