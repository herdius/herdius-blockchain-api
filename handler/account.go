package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/herdius/herdius-blockchain-api/protobuf"
	protoplugin "github.com/herdius/herdius-blockchain-api/protobuf"
	"github.com/herdius/herdius-core/p2p/network"
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

func (s *service) ProcessRequest(ctx context.Context, reqChan chan string, supervisorNode *network.PeerClient, resChan chan interface{}, errChan chan error) {
	i := 0
	select {
	case req := <-reqChan:
		fmt.Println("i:", i)
		fmt.Println("req:", req)
		res, err := supervisorNode.Request(ctx, &protoplugin.AccountRequest{Address: req})
		if err != nil {
			fmt.Println("req error:", err)
			errChan <- fmt.Errorf("error with request to Supervisor: %v", err)
			return
		}
		fmt.Println("res:", res)
		resChan <- res
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
			return acc
		default:
			errChan <- fmt.Errorf("response from Supervisor was not of AccountResponseType. Instead: %T", msg)
			return nil
		}
	}
}

// GetAccount handler called by http.HandleFunc
// TODO make these all pointers
func GetAccount(w http.ResponseWriter, r *http.Request, ctx context.Context, supervisorNode *network.PeerClient, reqChan chan string, resChan chan interface{}) {
	errChan := make(chan error, 0)
	s := service{}

	go func() {
		s.ReceiveRequest(r, reqChan, errChan)
	}()
	go func() {
		s.ProcessRequest(ctx, reqChan, supervisorNode, resChan, errChan)
	}()
	go func() {
		account := s.ProcessResponse(resChan, errChan)
		fmt.Println("account:", account)
		if len(account.Address) <= 0 {
			json.NewEncoder(w).Encode("Accound details not found for address: " + account.Address)
			return
		}
		log.Info().Msgf("Received Account detail for address: %s", account.Address)
		json.NewEncoder(w).Encode(account)
		return
	}()

	select {
	case err := <-errChan:
		json.NewEncoder(w).Encode("Failed to retrieve account details: " + fmt.Sprint(err))
		return
	}
	return
}
