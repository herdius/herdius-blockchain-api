package migration

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up00003, Down00003)
}

func Up00003(tx *sql.Tx) error {
	if _, err := tx.Exec(`ALTER TABLE transaction ADD COLUMN locked_amount numeric`); err != nil {
		return err
	}
	return nil
}

func Down00003(tx *sql.Tx) error {
	return nil
}
