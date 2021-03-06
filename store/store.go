// Package store provides interface to interact with herdius blockchain
// transaction without connect to herdius blockchain server.
package store

// Storer is the interface to manipulate blockchain transaction.
type Storer interface {
	Save(*Tx) error
	Update(*Tx) error
	Get(id string) (*Tx, error)
	GetBySender(address string) ([]*Tx, error)
	GetByAssetAndSender(asset, address string) ([]*Tx, error)
	GetByStatus(status string) ([]*Tx, error)
	GetTxByTypeBlockHeight(txType string, height uint64) ([]*Tx, error)
	GetLatestBlockID() (uint64, error)
	GetLastSyncBlockID() uint64
	SaveLastSyncBlockID(blockID uint64) error
}
