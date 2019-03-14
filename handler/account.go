package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	protoplugin "github.com/herdius/herdius-blockchain-api/protobuf"
	"github.com/herdius/herdius-core/accounts/account"
	"github.com/herdius/herdius-core/p2p/network"
	"github.com/rs/zerolog/log"
)

// GetAccountByAddress broadcasts a request to the supervisor to retrieve
// Account details for a given account address
func (s *service) GetAccountByAddress(accAddr string) (*account.Account, error) {
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
	time.Sleep(2 * time.Second)

	acc := &account.Account{}
	return acc, nil
}

type AccountMessagePlugin struct{ *network.Plugin }

// Receive handles block specific received messages
func (state *AccountMessagePlugin) Receive(ctx *network.PluginContext) error {
	switch msg := ctx.Message().(type) {
	case *protoplugin.AccountResponse:
		log.Info().Msgf("Account Response: %v", msg)
	}
	return nil
}

// GetAccount handler called by http.HandleFunc
func GetAccount(w http.ResponseWriter, r *http.Request) {
	params, ok := r.URL.Query()["address"]
	if !ok || len(params[0]) < 1 {
		log.Info().Msg("Url Param 'address' is missing")
		fmt.Fprint(w, "Request invalid, 'address' param missing\n")
	}

	address := params[0]

	srv := service{}
	account, _ := srv.GetAccountByAddress(address)
	fmt.Fprint(w, account)
}
