package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/herdius/herdius-core/accounts/account"
	"github.com/herdius/herdius-core/p2p/network"
	"github.com/rs/zerolog/log"
)

// NB is a Network Builder
var NB *NetworkBuilder

type service struct{}

// GetAccountByAddress broadcasts a request to the supervisor to retrieve
// Account details for a given account address
func (s *service) GetAccountByAddress(accAddr string) (*account.Account, error) {
	fmt.Println("entering service.GetAccountByAddress")
	net, err := NB.builder.Build()
	if err != nil {
		log.Error().Msgf("Failed to build network:%v", err)
	}

	go net.Listen()
	defer net.Close()

	fmt.Println("checkpoint1")
	supervisorAddress := "tcp://localhost:3000"
	supervisorAdds := make([]string, 1)
	supervisorAdds = append(supervisorAdds, supervisorAddress)
	bootStrap(net, supervisorAdds)

	fmt.Println("checkpoint2")
	ctx := network.WithSignMessage(context.Background(), true)

	net.Broadcast(ctx, &protoplugin.AccountRequest{AccountAddress: accountAddr})

	acc := account.Account{}
	fmt.Println("checkpoint3")
	return acc
}

// TODO add details about wtf bootStrap() does
func bootStrap(net *network.Network, peers []string) {
	if len(peers) > 0 {
		net.Bootstrap(peers...)
	}
}

// GetAccount handler called by http.HandleFunc
func GetAccount(w http.ResponseWriter, r *http.Request) {
	params, ok := r.URL.Query()["address"]
	if !ok || len(params[0]) < 1 {
		log.Info().Msg("Url Param 'address' is missing")
		fmt.Fprint(w, "Request invalid, 'address' param missing\n")
	}

	addressJSON := params[0]

	address, err := strconv.ParseInt(addressJSON, 10, 64)

	if err != nil {
		log.Error().Msgf("Failed to Parse %v", err)
	}

	srv := service{}
	account, _ := srv.GetAccount(height)
	log.Info().Msgf("Processed for Account: %d", account.Address)
	fmt.Fprint(w, accountResponse)
}
