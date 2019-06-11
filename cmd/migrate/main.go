package main

import (
	"flag"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/pressly/goose"

	_ "github.com/herdius/herdius-blockchain-api/store/migration"
	"github.com/herdius/herdius-blockchain-api/store/postgres"
)

const driver = "postgres"

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
	dir   = flags.String("dir", "../../store/migration", "directory with migration files")
	dsn   = flags.String("dsn", "user=postgres host=127.0.0.1 port=5432 dbname=postgres password=postgres sslmode=disable", "database dsn")
)

func main() {
	flags.Parse(os.Args[1:])
	args := flags.Args()

	if len(args) < 1 {
		flags.Usage()
		return
	}

	if err := goose.SetDialect(driver); err != nil {
		log.Fatal(err)
	}

	pStore, err := postgres.NewStore(*dsn)
	if err != nil {
		log.Fatal(err)
	}

	command := args[0]
	arguments := []string{}
	if len(args) > 1 {
		arguments = append(arguments, args[1:]...)
	}

	if err := goose.Run(command, pStore.DB(), *dir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
