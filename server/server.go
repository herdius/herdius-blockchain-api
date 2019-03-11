package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/herdius/herdius-core/accounts/account"
)

func main() {
	log.Println("Opening API")
	// Launch HTTP server
	LaunchServer()

	// Instantiate AccountDetail()
}

// LaunchServer opens a standard REST server, through which chain- and transaction-based
// CRUD requests are performed
func LaunchServer() {
	http.HandleFunc("/getaccount", GetAccount)
	log.Fatal(http.ListenAndServe(":80", nil))
}

// GetAccount is an http.Handler GET implementation that accepts valid Herdius
// tokens and returns an account Detail of the associated Herdius token address
func GetAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "\nRequest Pending\n")
	params, ok := r.URL.Query()["address"]
	if !ok || len(params[0]) < 1 {
		log.Println("Url Param 'address' is missing")
		fmt.Fprint(w, "Request invalid, 'address' param missing\n")
		return
	}
	addressJson := params[0]
	account := &account.Account{}
	account.Address = addressJson

	fmt.Fprint(w, "Request Valid\n")
	log.Println("Account address:", account.Address)
}
