package migration

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up00008, Down00008)
}

func Up00008(tx *sql.Tx) error {
	if _, err := tx.Exec(`ALTER TABLE transaction ADD COLUMN "data" VARCHAR(64)`); err != nil {
		return err
	}
	return nil
}

func Down00008(tx *sql.Tx) error {
	return nil
}
