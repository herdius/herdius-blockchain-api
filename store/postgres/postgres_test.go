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
		Network:         "herdius",
		CreationDT:      now,
	}

	assert.NoError(t, s.Save(tx))

	got, err := s.Get(tx.ID)
	assert.NoError(t, err)
	// Ignore time compare issue
	got.CreationDT = now
	assert.Equal(t, tx, got)

	tx.ID = "id2"
	assert.NoError(t, s.Save(tx))

	txs, err := s.GetBySender(tx.SenderAddress)
	assert.NoError(t, err)
	assert.Len(t, txs, 2)
}
