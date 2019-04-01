package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	protoplugin "github.com/herdius/herdius-blockchain-api/protobuf"
	"github.com/herdius/herdius-core-dev/blockchain/protobuf"
	"github.com/herdius/herdius-core/p2p/network"
	"github.com/rs/zerolog/log"
)

var (
	//	accountResultTracker = 0
	account = Account{}
)

// Account : Account Detail
type Account struct {
	Nonce       uint64 `json:"nonce"`
	Address     string `json:"address"`
	Balance     int64  `json:"balance"`
	StorageRoot string `json:"storageRoot"`
}

// GetAccountByAddress broadcasts a request to the supervisor to retrieve
// Account details for a given account address
func (s *service) GetAccountByAddress(accAddr string) (*Account, error) {
	net, err := NB.builder.Build()
	if err != nil {
		log.Error().Msgf("Failed to build network:%v", err)
	}

	go net.Listen()
	defer net.Close()

	supervisorAddress := "tcp://localhost:3000"
	supervisorAdds := make([]string, 1)
	supervisorAdds = append(supervisorAdds, supervisorAddress)
	bootStrap(net, supervisorAdds)

	ctx := network.WithSignMessage(context.Background(), true)

	net.Broadcast(ctx, &protoplugin.AccountRequest{Address: accAddr})
	time.Sleep(1 * time.Second)

	acc := &Account{}
	// accountResultTracker will be 1 if request to get account detail using address is broadcasted to
	// blockchain
	// TODO: Need to remove global variable implmentation after the MVP
	//accountResultTracker = 1
	return acc, nil
}

// AccountMessagePlugin ...
type AccountMessagePlugin struct{ *network.Plugin }

// Receive handles block specific received messages
func (state *AccountMessagePlugin) Receive(ctx *network.PluginContext) error {
	switch msg := ctx.Message().(type) {
	case *protobuf.ConnectionMessage:
		log.Info().Msgf("Account detail not found: %v", msg)
	case *protoplugin.AccountResponse:
		account.Address = msg.Address
		account.Balance = msg.Balance
		account.Nonce = uint64(msg.Nonce)
		account.StorageRoot = msg.StorageRoot
		//accountResultTracker = 1
		log.Info().Msgf("Account Response: %v", msg)
	}
	return nil
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
	_, err := srv.GetAccountByAddress(address)
	if err != nil {
		fmt.Fprint(w, err)
	} else {
		if len(account.Address) == 0 {
			//fmt.Fprint(w, "Accound details not found for address: "+address)
			json.NewEncoder(w).Encode("Accound details not found for address: " + address)
			return
		}

		if len(account.Address) > 0 {
			log.Info().Msgf("Received Account detail for address: %s", account.Address)
			//fmt.Fprint(w, account)
			json.NewEncoder(w).Encode(account)
		}
		account = Account{}
	}
}
