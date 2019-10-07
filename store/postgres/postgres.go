package postgres

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/herdius/herdius-blockchain-api/store"
)

const name = "postgres"

const txInsertStmt = `
INSERT INTO "transaction" (
	id,
	sender_address,
	sender_pubkey,
	receiver_address,
	category,
	symbol,
	network,
	value,
	locked_amount,
	nonce,
	message,
	sign,
	status,
	block_id,
	type,
	data,
	created_date
) VALUES (
	:id,
	:sender_address,
	:sender_pubkey,
	:receiver_address,
	:category,
	:symbol,
	:network,
	:value,
	:locked_amount,
	:nonce,
	:message,
	:sign,
	:status,
	:block_id,
	:type,
	:data,
	:created_date
)
`

const txSelectByIDStmt = `
SELECT
	id,
	sender_address,
	sender_pubkey,
	receiver_address,
	category,
	symbol,
	network,
	value,
	locked_amount,
	nonce,
	message,
	sign,
	status,
	block_id,
	type,
	data,
	created_date
FROM "transaction"
WHERE
	id = $1
`

const txSelectBySenderStmt = `
SELECT
	id,
	sender_address,
	sender_pubkey,
	receiver_address,
	category,
	symbol,
	network,
	value,
	locked_amount,
	nonce,
	message,
	sign,
	status,
	block_id,
	type,
	data,
	created_date
FROM "transaction"
WHERE
	sender_address = $1
`

const txSelectByStatusStmt = `
SELECT
	id,
	sender_address,
	sender_pubkey,
	receiver_address,
	category,
	symbol,
	network,
	value,
	nonce,
	message,
	sign,
	status,
	block_id,
	type,
	data,
	created_date
FROM "transaction"
WHERE
	status = $1
`

const txUpdateStmt = `
UPDATE "transaction" SET
	sender_address = :sender_address,
	sender_pubkey = :sender_pubkey,
	receiver_address = :receiver_address,
	category = :category,
	symbol = :symbol,
	network = :network,
	value = :value,
	locked_amount = :locked_amount,
	nonce = :nonce,
	message = :message,
	sign = :sign,
	status = :status,
	block_id = :block_id,
	type = :type,
	data = :data,
	created_date = :created_date
WHERE
	id = :id
`

const txSelectByAssetAndSenderStmt = `
SELECT
	id,
	sender_address,
	sender_pubkey,
	receiver_address,
	category,
	symbol,
	network,
	value,
	locked_amount,
	nonce,
	message,
	sign,
	status,
	block_id,
	type,
	data,
	created_date
FROM "transaction"
WHERE
	lower(symbol) = lower($1)
	AND
	sender_address = $2
`

const txSelectByLockBlockHeight = `
SELECT
	id,
	sender_address,
	sender_pubkey,
	receiver_address,
	category,
	symbol,
	network,
	value,
	locked_amount,
	nonce,
	message,
	sign,
	status,
	block_id,
	type,
	data,
	created_date
FROM "transaction"
WHERE
block_id = $2 AND lower("type") = lower($1)
`

// Store wraps *sqlx.DB
type Store struct {
	db *sqlx.DB
}

// NewStore returns a new Store from the provided databse string,
// Store implements store.Storer.
func NewStore(dbstring string) (*Store, error) {
	db, err := sqlx.Connect(name, dbstring)
	if err != nil {
		return nil, err
	}
	return &Store{db: db}, nil
}

// DB returns the underlying sqlx.DB object
func (s *Store) DB() *sql.DB {
	return s.db.DB
}

// Save stores store.Tx to database.
func (s *Store) Save(tx *store.Tx) error {
	if _, err := s.db.NamedExec(txInsertStmt, tx); err != nil {
		return err
	}
	return nil
}

// Update updatess a current store.Tx in database.
func (s *Store) Update(tx *store.Tx) error {
	if _, err := s.db.NamedExec(txUpdateStmt, tx); err != nil {
		return err
	}
	return nil
}

// Get returns transaction by given id.
func (s *Store) Get(id string) (*store.Tx, error) {
	var tx store.Tx
	if err := s.db.Get(&tx, txSelectByIDStmt, id); err != nil {
		return nil, err
	}
	return &tx, nil
}

// GetBySender returns list of transaction filter by given sender address.
func (s *Store) GetBySender(address string) ([]*store.Tx, error) {
	var txs []*store.Tx
	if err := s.db.Select(&txs, txSelectBySenderStmt, address); err != nil {
		return nil, err
	}
	return txs, nil
}

// GetByAssetAndSender returns list of transaction filter by given asset and sender address.
func (s *Store) GetByAssetAndSender(asset, address string) ([]*store.Tx, error) {
	var txs []*store.Tx
	if err := s.db.Select(&txs, txSelectByAssetAndSenderStmt, asset, address); err != nil {
		return nil, err
	}
	return txs, nil
}

// GetByStatus returns list of transaction filter by given status
func (s *Store) GetByStatus(status string) ([]*store.Tx, error) {
	var txs []*store.Tx
	if err := s.db.Select(&txs, txSelectByStatusStmt, status); err != nil {
		return nil, err
	}
	return txs, nil
}

// GetTxByTypeBlockHeight returns list of transaction filter by given block height.
func (s *Store) GetTxByTypeBlockHeight(txType string, height uint64) ([]*store.Tx, error) {
	var txs []*store.Tx
	if err := s.db.Select(&txs, txSelectByLockBlockHeight, txType, height); err != nil {
		return nil, err
	}
	return txs, nil
}

// GetLatestBlockID returns latest block id from saved transaction
func (s *Store) GetLatestBlockID() (uint64, error) {
	var blockID uint64
	if err := s.db.Get(&blockID, `SELECT block_id FROM transaction ORDER BY block_id DESC LIMIT 1`); err != nil {
		return 0, err
	}

	return blockID, nil
}

// GetLastSyncBlockID returns last synced block id.
func (s *Store) GetLastSyncBlockID() uint64 {
	var blockID uint64
	if err := s.db.Get(&blockID, `SELECT id FROM last_sync_block ORDER BY id DESC LIMIT 1`); err != nil {
		println(err.Error())
		return 0
	}

	return blockID
}

// SaveLastSyncBlockID save last synced block id to database.
func (s *Store) SaveLastSyncBlockID(blockID uint64) error {
	_, err := s.db.Exec(`UPDATE "last_sync_block" SET id = $1`, blockID)
	return err
}
