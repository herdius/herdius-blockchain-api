package migration

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up00006, Down00006)
}

func Up00006(tx *sql.Tx) error {
	if _, err := tx.Exec(`CREATE INDEX idx_lower_type ON transaction(lower(type))`); err != nil {
		return err
	}
	return nil
}

func Down00006(tx *sql.Tx) error {
	return nil
}
