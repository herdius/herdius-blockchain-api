package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/herdius/herdius-blockchain-api/config"
	apiNet "github.com/herdius/herdius-blockchain-api/network"
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
func (s *service) GetAccountByAddress(accAddr string) (*Account, error) {

	net, err := apiNet.GetNetworkBuilder().Build()
	if err != nil {
		log.Error().Msgf("Failed to build network:%v", err)
	}

	go net.Listen()
	defer net.Close()

	configuration := config.GetConfiguration()
	supervisorAddress := configuration.GetSupervisorAddress()

	ctx := network.WithSignMessage(context.Background(), true)
	supervisorNode, _ := net.Client(supervisorAddress)
	res, err := supervisorNode.Request(ctx, &protoplugin.AccountRequest{Address: accAddr})

	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Failed to find block due to :%v", err))
	}

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
	}
	return nil, nil
}

// GetAccount handler called by http.HandleFunc
func GetAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if len(params["address"]) == 0 {
		json.NewEncoder(w).Encode("Request invalid, 'address' param missing\n")
		return
	}

	address := params["address"]

	srv := service{}
	account, err := srv.GetAccountByAddress(address)
	if err != nil {
		json.NewEncoder(w).Encode("Failed to retrieve acount detail due to: " + err.Error())
	} else {
		if len(account.Address) == 0 {
			json.NewEncoder(w).Encode("Accound details not found for address: " + address)
			return
		}

		if len(account.Address) > 0 {
			log.Info().Msgf("Received Account detail for address: %s", account.Address)
			json.NewEncoder(w).Encode(account)
		}
	}
}
