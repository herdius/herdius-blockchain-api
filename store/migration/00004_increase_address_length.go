package migration

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up00004, Down00004)
}

func Up00004(tx *sql.Tx) error {
	if _, err := tx.Exec(`ALTER TABLE transaction ALTER COLUMN sender_address TYPE VARCHAR(64)`); err != nil {
		return err
	}
	if _, err := tx.Exec(`ALTER TABLE transaction ALTER COLUMN receiver_address TYPE VARCHAR(64)`); err != nil {
		return err
	}
	return nil
}

func Down00004(tx *sql.Tx) error {
	return nil
}
