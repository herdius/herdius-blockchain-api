package migration

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up00009, Down00009)
}

func Up00009(tx *sql.Tx) error {
	if _, err := tx.Exec(`ALTER TABLE transaction ADD COLUMN "eth_address" VARCHAR(64)`); err != nil {
		return err
	}
	if _, err := tx.Exec(`ALTER TABLE transaction ADD COLUMN "btc_address" VARCHAR(64)`); err != nil {
		return err
	}
	if _, err := tx.Exec(`ALTER TABLE transaction ADD COLUMN "ltc_address" VARCHAR(64)`); err != nil {
		return err
	}
	if _, err := tx.Exec(`ALTER TABLE transaction ADD COLUMN "xtz_address" VARCHAR(64)`); err != nil {
		return err
	}
	if _, err := tx.Exec(`ALTER TABLE transaction ADD COLUMN "bnb_address" VARCHAR(64)`); err != nil {
		return err
	}
	return nil
}

func Down00009(tx *sql.Tx) error {
	return nil
}
