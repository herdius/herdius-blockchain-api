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
	nonce,
	message,
	sign,
	status,
	block_id,
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
	:nonce,
	:message,
	:sign,
	:status,
	:block_id,
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
	nonce,
	message,
	sign,
	status,
	block_id,
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
	nonce,
	message,
	sign,
	status,
	block_id,
	created_date
FROM "transaction"
WHERE
	sender_address = $1
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
	nonce = :nonce,
	message = :message,
	sign = :sign,
	status = :status,
	block_id = :block_id,
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
	nonce,
	message,
	sign,
	status,
	block_id,
	created_date
FROM "transaction"
WHERE
	lower(symbol) = lower($1)
	AND
	sender_address = $2
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