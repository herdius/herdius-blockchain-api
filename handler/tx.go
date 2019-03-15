package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	apiProtobuf "github.com/herdius/herdius-blockchain-api/protobuf"
	"github.com/herdius/herdius-core/p2p/network"
	"github.com/rs/zerolog/log"
)

// TXMessagePlugin will receive all Transaction specific messages.
type TXMessagePlugin struct{ *network.Plugin }

// PostTransaction handler called by http.HandleFunc
func PostTransaction(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	var txReq *apiProtobuf.Transaction
	err = json.Unmarshal(body, &txReq)
	if err != nil {
		log.Info().Msg("Could not parse POST json data, invalid format")
		fmt.Fprint(w, "\nRequest invalid, Could not parse POST json data, invalid format, err:\n", err)
		return
	}
	if txReq.Senderpubkey == "" ||
		txReq.Signature == "" ||
		txReq.Recaddress == "" ||
		txReq.Asset.Network == "" ||
		txReq.Asset.Value < 0 ||
		txReq.Asset.Fee < 0 {
		log.Info().Msg("POST body did not include all required values")
		fmt.Fprint(w, "\nRequest missing data, POST body did not include all required values\n")
		fmt.Fprint(w, "\ntxreq:\n", txReq)
		return
	}

	fmt.Fprint(w, "\nRequest syntax validated")
	net, err := NB.builder.Build()
	if err != nil {
		log.Error().Msgf("Failed to build network:", err)
		fmt.Fprint(w, "\nFailed to build network, request aborted")
		return
	}

	go net.Listen()
	defer net.Close()

	supervisorAddress := "tcp://localhost:3000"
	supervisorAdds := make([]string, 1)
	supervisorAdds = append(supervisorAdds, supervisorAddress)
	bootStrap(net, supervisorAdds)

	ctx := network.WithSignMessage(context.Background(), true)
	net.Broadcast(ctx, &apiProtobuf.TransactionRequest{Tx: txReq})
	time.Sleep(2 * time.Second)
}

type TxMessagePlugin struct{ *network.Plugin }

// Receive handles block specific received messages
func (p *TxMessagePlugin) Receive(ctx *network.PluginContext) error {
	switch msg := ctx.Message().(type) {
	case *apiProtobuf.TransactionResponse:
		log.Info().Msgf("Tx Response: %v", msg)
	}
	return nil
}
