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
type TxReq struct {
	SenderAddress string `json:senderaddress`
	SenderPubKey  string `json:senderpubkey`
	RecAddress    string `json:recaddress`
	Signature     string `json:signature`
	Asset         struct {
		Message  string  `json:message`
		Network  string  `json:network`
		Category string  `json:category`
		Currency string  `json:currency`
		Value    float64 `json:value`
		Fee      float64 `json:fee`
	} `json:asset`
}

// PostTransaction handler called by http.HandleFunc
func PostTransaction(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	var txReq TxReq
	err = json.Unmarshal(body, &txReq)
	if err != nil {
		log.Info().Msg("Could not parse POST json data, invalid format")
		fmt.Fprint(w, "Request invalid, Could not parse POST json data, invalid format, err: %v\n", err)
	}
	if txReq.SenderAddress == "" ||
		txReq.SenderPubKey == "" ||
		txReq.RecAddress == "" ||
		txReq.Signature == "" ||
		txReq.Asset.Network == "" ||
		txReq.Asset.Category == "" ||
		txReq.Asset.Currency == "" ||
		txReq.Asset.Value < 0 ||
		txReq.Asset.Fee < 0 {
		log.Info().Msg("POST body did not include all required values")
		fmt.Fprint(w, "Request invalid, POST body did not include all required values\n")
	}

	fmt.Fprint(w, "Request syntax validated")
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
