package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/herdius/herdius-core/accounts/account"
	"github.com/herdius/herdius-core/supervisor/transaction"
)

func main() {
	log.Println("Opening API")
	LaunchServer()
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
	acct, err := validateGetAccount(w, r)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	log.Println("Account address:", acct.Address)
}

func GetAccountPortfolio(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "\nRequest Pending\n")
	acct, err := validateGetAccount(w, r)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	log.Println("Account address:", acct.Address)
}

// PostTransaction is an http.Handler POST implementation that accepts valid transactions
// and sends them to the memory pool for processing and batching
func PostTransaction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "\nTransaction Pending\n")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	tx := &transaction.Tx{}
	json.Unmarshal(body, &tx)
	log.Println("tx.Message:", tx.Message)

	if tx.Message == "" ||
		tx.Assetcategory == "" ||
		tx.Assetname == "" ||
		len(tx.Value) == 0 ||
		len(tx.Signature) == 0 ||
		len(tx.Senderpubkey) == 0 ||
		len(tx.Fee) == 0 {
		fmt.Fprint(w, "Transaction was missing necessary values")
		return
	}

	fmt.Fprint(w, "Transaction Posted\n")
}

func validateGetAccount(w http.ResponseWriter, r *http.Request) (*account.Account, error) {
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
	http.HandleFunc("/posttx", PostTransaction)
	log.Fatal(http.ListenAndServe(":80", nil))
}
