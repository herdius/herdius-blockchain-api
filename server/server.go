package main

import (
	"errors"
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
	http.HandleFunc("/getaccountportfolio", GetAccountPortfolio)
	log.Fatal(http.ListenAndServe(":80", nil))
}

// GetAccount is an http.Handler GET implementation that accepts valid Herdius
// tokens and returns an account Detail of the associated Herdius token address
func GetAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "\nRequest Pending\n")
	acct, err := validateInput(w, r)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	log.Println("Account address:", acct.Address)
}

func GetAccountPortfolio(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "\nRequest Pending\n")
	acct, err := validateInput(w, r)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	log.Println("Account address:", acct.Address)
}

func validateInput(w http.ResponseWriter, r *http.Request) (*account.Account, error) {
	params, ok := r.URL.Query()["address"]
	if !ok || len(params[0]) < 1 {
		log.Println("Url Param 'address' is missing")
		fmt.Fprint(w, "Request invalid, 'address' param missing\n")
		return nil, errors.New("Invalid address parameter provided")
	}
	addressJson := params[0]
	account := &account.Account{}
	account.Address = addressJson

	fmt.Fprint(w, "Request Valid\n")
	return account, nil
}
