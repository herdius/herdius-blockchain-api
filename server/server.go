package main

import (
	"log"
	"net/http"

	"github.com/herdius/herdius-blockchain-api/handler"
)

func main() {
	log.Println("Opening API")
	LaunchServer()
}

// LaunchServer opens a standard REST server, through which chain- and transaction-based
// CRUD requests are performed
func LaunchServer() {
	http.HandleFunc("/block", handler.GetBlockByHeight)
	http.HandleFunc("/account", handler.GetAccount)
	//http.HandleFunc("/accountportfolio", handler.GetAccountPortfolio)
	log.Fatal(http.ListenAndServe(":80", nil))
}
