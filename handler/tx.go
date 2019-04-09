package handler

import (
	"context"
	"strings"

	"encoding/json"
	"net/http"
	"time"

	"github.com/herdius/herdius-blockchain-api/protobuf"

	"log"

	"github.com/herdius/herdius-core/p2p/network"
)

var (
	// newTxResponse variable holds the transaction response detail from Herdius blockchain.
	// This is an inappropriate approach to hold data coming from other services.
	// TODO: This needs to be switched to be getting handled in single context
	// of client request. Once the request is served the context also needs to be closed.
	newTxResponse = protobuf.TxResponse{}
)

// TXMessagePlugin will receive all Transaction specific messages.
type TXMessagePlugin struct{ *network.Plugin }

// Receive is Tx Receiver
func (state *TXMessagePlugin) Receive(ctx *network.PluginContext) error {
	switch msg := ctx.Message().(type) {
	case *protobuf.TxResponse:
		log.Printf("Tx ID: %v", msg.TxId)
		log.Printf("Tx ID: %v", msg.Status)

		newTxResponse = *msg
	}
	return nil
}

func (s *service) SendTxToBlockchain(txReq protobuf.TxRequest) (*protobuf.TxResponse, error) {
	net, err := NB.builder.Build()
	if err != nil {
		log.Print(err)
	}

	go net.Listen()
	defer net.Close()

	supervisorAddress := "tcp://localhost:3000"
	supervisorAdds := make([]string, 1)
	supervisorAdds = append(supervisorAdds, supervisorAddress)
	bootStrap(net, supervisorAdds)

	ctx := network.WithSignMessage(context.Background(), true)

	net.Broadcast(ctx, &txReq)

	// VERY IMPORTANT TODO: Having the process sleep is not an efficient way to handle requests.
	// This needs to be addressed by optimizing current P2P layer
	time.Sleep(1 * time.Second)

	return nil, nil
}

// SendTx ...
func SendTx(w http.ResponseWriter, r *http.Request) {

	var txRequest protobuf.TxRequest
	err := json.NewDecoder(r.Body).Decode(&txRequest)
	if err != nil {
		json.NewEncoder(w).Encode("\nRequest invalid, Could not parse POST json data, invalid format, err:\n" + err.Error())
		return
	}

	// Check if tx type is account registration
	if len(txRequest.Tx.Type) > 0 && strings.EqualFold(txRequest.Tx.Type, "register") {

		if len(txRequest.Tx.SenderAddress) == 0 ||
			len(txRequest.Tx.Sign) == 0 ||
			len(txRequest.Tx.SenderPubkey) == 0 {
			json.NewEncoder(w).Encode("\nRequest missing data, POST body did not include all required values\n")
			json.NewEncoder(w).Encode("\ntxreq:\n")
			json.NewEncoder(w).Encode(txRequest)
			return
		}
		srv := service{}
		srv.SendTxToBlockchain(txRequest)

		json.NewEncoder(w).Encode(newTxResponse)
		newTxResponse = protobuf.TxResponse{}
		return
	}

	if len(txRequest.Tx.SenderAddress) == 0 ||
		len(txRequest.Tx.Sign) == 0 ||
		len(txRequest.Tx.RecieverAddress) == 0 ||
		txRequest.Tx.Asset.Network == "" ||
		txRequest.Tx.Asset.Nonce <= 0 ||
		txRequest.Tx.Asset.Value < 0 {
		log.Println("POST body did not include all required values")

		json.NewEncoder(w).Encode("\nRequest missing data, POST body did not include all required values\n")
		json.NewEncoder(w).Encode("\ntxreq:\n")
		json.NewEncoder(w).Encode(txRequest)

		return
	}
	srv := service{}
	srv.SendTxToBlockchain(txRequest)

	json.NewEncoder(w).Encode(newTxResponse)
	newTxResponse = protobuf.TxResponse{}
}
