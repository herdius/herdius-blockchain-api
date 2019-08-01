package migration

import (
	"database/sql"

	"github.com/pressly/goose"
)

const createLastSyncBlockTableStatement = `
CREATE TABLE "last_sync_block" (
  id INTEGER
);

INSERT INTO "last_sync_block" VALUES (0);
`

func init() {
	goose.AddMigration(Up00007, Down00007)
}

func Up00007(tx *sql.Tx) error {
	if _, err := tx.Exec(createLastSyncBlockTableStatement); err != nil {
		return err
	}
	return nil
}

func Down00007(tx *sql.Tx) error {
	return nil
}
