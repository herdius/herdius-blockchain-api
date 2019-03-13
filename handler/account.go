package handler

import (
	"context"
	"fmt"

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
