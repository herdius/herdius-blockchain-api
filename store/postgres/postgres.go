package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/herdius/herdius-blockchain-api/store"
)

const name = "postgres"

// Storage wraps *sqlx.DB
type Storage struct {
	db *sqlx.DB
}

// NewStorage returns a new Storage from the provided psql databse string
func NewStorage(dbstring string) (*Storage, error) {
	db, err := sqlx.Connect(name, dbstring)
	if err != nil {
		return nil, err
	}
	return &Storage{db: db}, nil
}

// DB returns the underlying sqlx.DB object
func (s *Storage) DB() *sqlx.DB {
	return s.db.DB
}

// Save stores store.Tx to database.
func (s *Storage) Save(*store.Tx) error {
	return nil
}

// Get returns transaction by given id.
func (s *Storage) Get(id string) *store.Tx {
	return nil
}

// GetBySender returns list of transaction filter by given sender address.
func (s *Storage) GetBySender(address string) []*store.Tx {
	return nil
}
