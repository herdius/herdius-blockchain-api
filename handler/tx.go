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
		json.NewEncoder(w).Encode("\nRequest invalid, Could not parse POST json data, invalid format, err:\n" + err.Error())
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
