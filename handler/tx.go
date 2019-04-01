package handler

import (
	"context"

	"encoding/json"
	"net/http"
	"time"

	"github.com/herdius/herdius-blockchain-api/protobuf"

	b64 "encoding/base64"

	"log"

	cryptoAmino "github.com/herdius/herdius-core-dev/crypto/encoding/amino"
	"github.com/herdius/herdius-core-dev/crypto/secp256k1"

	"github.com/herdius/herdius-core/p2p/network"
	amino "github.com/tendermint/go-amino"
)

var cdc = amino.NewCodec()

func init() {

	cryptoAmino.RegisterAmino(cdc)
}

var (
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
		json.NewEncoder(w).Encode(err)
		return
	}
	// Re-Create the TX to verify sign
	tx := txRequest.Tx
	asset := protobuf.Asset{
		Category: tx.Asset.Category,
		Symbol:   tx.Asset.Symbol,
		Network:  tx.Asset.Network,
		Value:    tx.Asset.Value,
		Fee:      tx.Asset.Fee,
		Nonce:    tx.Asset.Nonce,
	}
	txToVerify := protobuf.Tx{
		SenderAddress:   tx.SenderAddress,
		SenderPubkey:    tx.SenderPubkey,
		RecieverAddress: tx.RecieverAddress,
		Asset:           &asset,
		Message:         tx.Message,
	}

	txbBeforeSign, _ := json.Marshal(txToVerify)
	//verify the sign
	senderPubKeyBytes, _ := b64.StdEncoding.DecodeString(tx.SenderPubkey)

	//sendPubKey := senderPrivKey.PubKey()
	var sendPubkey secp256k1.PubKeySecp256k1

	cdc.UnmarshalBinaryBare(senderPubKeyBytes, &sendPubkey)

	decodedSig, _ := b64.StdEncoding.DecodeString(tx.Sign)

	log.Println(sendPubkey.VerifyBytes(txbBeforeSign, decodedSig))
	srv := service{}
	srv.SendTxToBlockchain(txRequest)

	json.NewEncoder(w).Encode(newTxResponse)
	newTxResponse = protobuf.TxResponse{}
}

/* package handler

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
	if len(txReq.Senderpubkey) == 0 ||
		len(txReq.Signature) == 0 ||
		txReq.Recaddress == "" ||
		txReq.Asset.Network == "" ||
		txReq.Asset.Nonce <= 0 ||
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
*/
