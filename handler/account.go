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
func (s *service) ReceiveRequest(r *http.Request, reqChan chan string, errChan chan error) {
	params := mux.Vars(r)
	if len(params["address"]) == 0 {
		errChan <- errors.New("Request invalid, 'address' param missing\n")
		return
	}
	reqChan <- params["address"]
}

func (s *service) ProcessRequest(reqChan chan string, resChan chan interface{}, errChan chan error) {
	select {
	case req := <-reqChan:
		res, err := supervisorNode.Request(ctx, &protoplugin.AccountRequest{Address: accAddr})
		if err != nil {
			errChan <- fmt.Errorf("error with request to Supervisor: %v", err)
			return
		}
		resChan <- res
		return
	default:
		errChan <- fmt.Errorf("no request received for Supervisor")
		return
	}
}

func (s *service) ProcessResponse(resChan chan interface{}, errChan chan error) *Account {
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
			errChan <- fmt.Errorf("response from Supervisor was not of AccountResponseType. Instead: %T", msg)
			return
		}
	default:
		errChan <- fmt.Errorf("no response received from Supervisor")
		return
	}
}

// GetAccount handler called by http.HandleFunc
func GetAccount(w http.ResponseWriter, r *http.Request, reqChan chan string, resChan chan interface{}) {
	errChan := make(chan error, 0)
	s := service{}

	go func() {
		s.ReceiveRequest(r, reqChan, errChan)
	}()
	go func() {
		s.ProcessRequest(reqChan, resChan, errChan)
	}()
	go func() {
		account := s.ProcessResponse(resChan, errChan)
		if len(account.Address) <= 0 {
			json.NewEncoder(w).Encode("Accound details not found for address: " + address)
			return
		}
		log.Info().Msgf("Received Account detail for address: %s", account.Address)
		json.NewEncoder(w).Encode(account)
		return
	}()

	select {
	case err := <-errors:
		json.NewEncoder(w).Encode("Failed to retrieve account details: " + err)
		return
	}
	return
}
