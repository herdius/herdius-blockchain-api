package handler

import (
	"context"
	"fmt"
	"strings"

	"encoding/json"
	"net/http"

	"github.com/herdius/herdius-blockchain-api/config"
	apiNet "github.com/herdius/herdius-blockchain-api/network"
	"github.com/herdius/herdius-blockchain-api/protobuf"

	"log"

	"github.com/herdius/herdius-core/p2p/network"
)

func (s *service) SendTxToBlockchain(txReq protobuf.TxRequest) (*protobuf.TxResponse, error) {
	net, err := apiNet.GetNetworkBuilder().Build()
	if err != nil {
		log.Print(err)
	}

	go net.Listen()
	defer net.Close()

	configuration := config.GetConfiguration()

	supervisorAddress := configuration.GetSupervisorAddress()

	supervisorAdds := make([]string, 1)
	supervisorAdds = append(supervisorAdds, supervisorAddress)
	bootStrap(net, supervisorAdds)

	ctx := network.WithSignMessage(context.Background(), true)

	supervisorNode, _ := net.Client(supervisorAddress)
	res, err := supervisorNode.Request(ctx, &txReq)

	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Failed to find block due to :%v", err))
	}

	switch msg := res.(type) {
	case *protobuf.TxResponse:
		log.Printf("Tx ID: %v", msg.TxId)
		log.Printf("Tx ID: %v", msg.Status)
		return msg, nil
	}

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

	// Check if tx type is account update
	if len(txRequest.Tx.Type) > 0 && strings.EqualFold(txRequest.Tx.Type, "update") {

		if len(txRequest.Tx.SenderAddress) == 0 ||
			len(txRequest.Tx.Sign) == 0 ||
			len(txRequest.Tx.SenderPubkey) == 0 {
			json.NewEncoder(w).Encode("\nRequest missing data, POST body did not include all required values\n")
			json.NewEncoder(w).Encode("\ntxreq:\n")
			json.NewEncoder(w).Encode(txRequest)
			return
		}
		srv := service{}
		newTxResponse, _ := srv.SendTxToBlockchain(txRequest)

		json.NewEncoder(w).Encode(newTxResponse)
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
	newTxResponse, err := srv.SendTxToBlockchain(txRequest)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode(newTxResponse)
	}

}
