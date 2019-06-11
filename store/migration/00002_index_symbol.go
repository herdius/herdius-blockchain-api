package migration

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up00002, Down00002)
}

func Up00002(tx *sql.Tx) error {
	if _, err := tx.Exec(`CREATE INDEX idx_lower_symbol ON transaction(lower(symbol))`); err != nil {
		return err
	}
	return nil
}

func Down00002(tx *sql.Tx) error {
	return nil
}
