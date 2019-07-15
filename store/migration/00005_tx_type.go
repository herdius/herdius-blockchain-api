package migration

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up00005, Down00005)
}

func Up00005(tx *sql.Tx) error {
	if _, err := tx.Exec(`ALTER TABLE transaction ADD COLUMN tx_type TYPE VARCHAR(64)`); err != nil {
		return err
	}
	return nil
}

func Down00005(tx *sql.Tx) error {
	return nil
}
