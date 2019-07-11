package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/herdius/herdius-blockchain-api/config"
	"github.com/herdius/herdius-blockchain-api/protobuf"
	protoplugin "github.com/herdius/herdius-blockchain-api/protobuf"
	"github.com/herdius/herdius-core/p2p/network"
	"github.com/rs/zerolog/log"
)

// Account : Account Detail
type Account struct {
	Nonce        uint64 `json:"nonce"`
	Address      string `json:"address"`
	Balance      uint64 `json:"balance"`
	StorageRoot  string `json:"storageRoot"`
	PublickKey   string `json:"publicKey"`
	Erc20Address string `json:"erc20Address"`

	FirstExternalAddress map[string]string
	ExternalNonce        uint64

	LastBlockHeight uint64
	EBalances       map[string]map[string]EBalance
}

// EBalance is external balance model
type EBalance struct {
	Address         string
	Balance         uint64
	LastBlockHeight uint64
	Nonce           uint64
}

// GetAccountByAddress broadcasts a request to the supervisor to retrieve
// Account details for a given account address
func (s *service) GetAccountByAddress(accAddr string, net *network.Network, env string) (*Account, error) {
	configuration := config.GetConfiguration(env)
	supervisorAddress := configuration.GetSupervisorAddress()

	ctx := network.WithSignMessage(context.Background(), true)
	supervisorNode, err := net.Client(supervisorAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}
	if supervisorNode.Address == "" {
		fmt.Println("empty supervisornode:", supervisorNode)
	}
	res, err := supervisorNode.Request(ctx, &protoplugin.AccountRequest{Address: accAddr})
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Failed to find block due to: %v", err))
	}

	switch msg := res.(type) {
	case *protobuf.AccountResponse:
		acc := &Account{}
		acc.Address = msg.Address
		acc.Balance = msg.Balance
		acc.Nonce = msg.Nonce
		acc.PublickKey = msg.PublicKey
		acc.StorageRoot = msg.StorageRoot
		acc.Erc20Address = msg.Erc20Address
		acc.ExternalNonce = msg.ExternalNonce
		acc.LastBlockHeight = msg.LastBlockHeight
		acc.FirstExternalAddress = msg.FirstExternalAddress

		eBalances := make(map[string]map[string]EBalance)
		for asset, assetAccount := range msg.EBalances {
			eBalances[asset] = make(map[string]EBalance)
			for _, eb := range assetAccount.Asset {
				eBalance := EBalance{
					Address:         eb.Address,
					Balance:         eb.Balance,
					LastBlockHeight: eb.LastBlockHeight,
					Nonce:           eb.Nonce,
				}
				eBalances[asset][eb.Address] = eBalance
			}
		}
		acc.EBalances = eBalances
		return acc, nil
	}
	return nil, nil
}

// GetAccount handler called by http.HandleFunc
func GetAccount(w http.ResponseWriter, r *http.Request, net *network.Network, env string) {
	params := mux.Vars(r)
	if len(params["address"]) == 0 {
		json.NewEncoder(w).Encode("Request invalid, 'address' param missing\n")
		return
	}

	address := params["address"]

	srv := service{}
	account, err := srv.GetAccountByAddress(address, net, env)
	if err != nil {
		json.NewEncoder(w).Encode("Failed to retrieve acount detail due to: " + err.Error())
	} else {
		if account == nil || len(account.Address) == 0 {
			errMsg := "Please wait for block time of 15-30 seconds if account registration is requested otherwise accound details not found for address: "
			json.NewEncoder(w).Encode(errMsg + address)
			return
		}

		if len(account.Address) > 0 {
			log.Info().Msgf("Received Account detail for address: %s", account.Address)
			json.NewEncoder(w).Encode(account)
		}
	}
}
