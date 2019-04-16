package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/herdius/herdius-blockchain-api/protobuf"
	protoplugin "github.com/herdius/herdius-blockchain-api/protobuf"
	"github.com/rs/zerolog/log"
)

// Account : Account Detail
type Account struct {
	Nonce       uint64            `json:"nonce"`
	Address     string            `json:"address"`
	Balance     uint64            `json:"balance"`
	StorageRoot string            `json:"storageRoot"`
	PublickKey  string            `json:"publicKey"`
	Balances    map[string]uint64 // Balances will store balances of assets e.g. [BTC]=10 or [HER]=1000
}

// GetAccountByAddress broadcasts a request to the supervisor to retrieve
// Account details for a given account address
func (s *service) ReceiveRequest(r *http.Request, reqChan chan string) error {
	params := mux.Vars(r)
	if len(params["address"]) == 0 {
		return errors.New("Request invalid, 'address' param missing\n")
	}
	reqChan <- params["address"]
}

func (s *service) ProcessRequest(reqChan chan string, resChan chan interface{}) error {
	select {
	case req := <-reqChan:
		res, err := supervisorNode.Request(ctx, &protoplugin.AccountRequest{Address: accAddr})
		if err != nil {
			return nil, fmt.Errorf("error with request to Supervisor: %v", err)
		}
		resChan <- res
	default:
		return fmt.Errorf("no request received for Supervisor")
	}
}

func (s *service) ProcessResponse(resChan chan interface{}) (*Account, error) {
	select {
	case res := <-resChan:
		switch msg := res.(type) {
		case *protobuf.AccountResponse:
			acc := &Account{}
			acc.Address = msg.Address
			acc.Balance = msg.Balance
			acc.Nonce = msg.Nonce
			acc.PublickKey = msg.PublicKey
			acc.StorageRoot = msg.StorageRoot
			acc.Balances = msg.Balances
			return acc, nil
		default:
			return fmt.Errorf("response from Supervisor was not of AccountResponseType. Instead: %T", msg)
		}
	default:
		return fmt.Errorf("no response received from Supervisor")
	}
}

// GetAccount handler called by http.HandleFunc
func GetAccount(w http.ResponseWriter, r *http.Request, reqChan chan string, resChan chan interface{}) {
	go ReceiveRequest()
	go ProcessRequest()
	go ProcessResponse()
	if err != nil {
		json.NewEncoder(w).Encode("Failed to retrieve acount detail due to: " + err)
		return
	}
	if len(account.Address) <= 0 {
		json.NewEncoder(w).Encode("Accound details not found for address: " + address)
		return
	}
	log.Info().Msgf("Received Account detail for address: %s", account.Address)
	json.NewEncoder(w).Encode(account)
	return
}
