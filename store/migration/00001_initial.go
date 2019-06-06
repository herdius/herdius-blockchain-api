package main

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose"
)

const createTxTableStatement = `
CREATE TABLE "transaction" (
  id varchar(64) PRIMARY KEY,
  sender_address varchar(35),
  sender_pubkey varchar(256),
  receiver_address varchar(35),
  category varchar(32),
  symbol varchar(16),
  network varchar(32),
  value numeric,
  nonce numeric,
  message text,
  sign varchar(128),
  status varchar(8),
  block_id numeric,
  created_date timestamp
) INHERITS existing_table_name;

CREATE INDEX idx_sender_address ON transaction(sender_address);
CREATE INDEX idx_receiver_address ON transaction(receiver_address);
CREATE INDEX idx_nonce ON transaction(nonce);
`

func init() {
	goose.AddMigration(Up00001, Down00001)
}

func Up00001(tx *sqlx.Tx) error {
	if _, err := tx.Exec(createTxTableStatement); err != nil {
		return err
	}
	return nil
}

func Down00001(tx *sql.Tx) error {
	return nil
}
