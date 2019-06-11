package handler

import (
	"log"
	"sync"

	"github.com/herdius/herdius-blockchain-api/store"
	"github.com/herdius/herdius-blockchain-api/store/postgres"
)

var (
	db   store.Storer
	once sync.Once
)

func getStore(dsn string) store.Storer {
	once.Do(func() {
		var err error
		db, err = postgres.NewStore(dsn)
		if err != nil {
			log.Printf("Failed to init store: %v", err)
		}
	})
	return db
}
