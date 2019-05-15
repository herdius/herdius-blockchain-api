package main

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/herdius/herdius-blockchain-api/protobuf"
	"github.com/herdius/herdius-core/crypto/secp256k1"
	"github.com/herdius/herdius-core/p2p/key"
)

/**
This file is just for testing purpose.
It uses the keys that are loaded during blockchain startup
*/

func main() {

	var host = flag.String("host", "localhost", "host to listen to")
	var txType = flag.String("txtype", "value", "transaction type")
	flag.Parse()
	var endpoint string
	endpoint = *host

	endpoint = "http://" + endpoint + ":80/tx"
	log.Println("endpoint:", endpoint)

	if strings.EqualFold(*txType, "update") {
		sendAccountRegisterTx(endpoint)
	} else {
		postTx(endpoint)
	}

}
func sendAccountRegisterTx(endpoint string) {
	// Create key pairs and store in a local file
	senderPrivKey, err := key.LoadOrGenNodeKey("./tempKey.json")
	if err != nil {
		panic(err)
	}

	sendPubKey := senderPrivKey.PubKey()
	var pubkeyBytes secp256k1.PubKeySecp256k1
	pubkeyBytes = sendPubKey.(secp256k1.PubKeySecp256k1)

	senderB64 := b64.StdEncoding.EncodeToString(pubkeyBytes[:])
	senderAddress := sendPubKey.GetAddress()
	log.Println("Account update request : " + senderAddress)
	msg := "Update my account"
	asset := &protobuf.Asset{
		Category: "crypto",
		Symbol:   "HER",
		Network:  "Herdius",
		Value:    0,
		Fee:      0,
		Nonce:    2,
	}
	tx := protobuf.Tx{
		SenderAddress: senderAddress,
		SenderPubkey:  senderB64,
		Message:       msg,
		Type:          "update",
		Asset:         asset,
	}

	// Sign the transaction detail
	txbBeforeSign, _ := json.Marshal(tx)

	sig, err := senderPrivKey.PrivKey.Sign(txbBeforeSign)

	tx.Sign = b64.StdEncoding.EncodeToString(sig)

	// Post tx to blockchain.
	txReq := protobuf.TxRequest{
		Tx: &tx,
	}
	txJSON, err := json.Marshal(txReq)
	if err != nil {
		log.Fatalf("Failed to Marshal %v", err)
	}
	response, err := http.Post(endpoint, "application/json", bytes.NewBuffer(txJSON))
	if err != nil {
		log.Fatalf("Failed to Marshal %v", err)
	}

	defer response.Body.Close()
	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		log.Fatalf("Failed to read http response %s.", readErr)
	}
	var txResponse protobuf.TxResponse
	jsonErr := json.Unmarshal(body, &txResponse)
	if jsonErr != nil {
		log.Fatalf("Failed to Unmarshal %s.", jsonErr)
	}

	log.Println(txResponse.TxId)
	log.Println(txResponse.Status)
}

func postTx(endpoint string) {
	// Copied the user id's from herdius-core package
	dataPath := "./testdata/secp205k1Accts/"
	senderPrivKey, err := key.LoadOrGenNodeKey(dataPath + "1_peer_id.json")
	if err != nil {
		panic(err)
	}
	recPrivKey, err := key.LoadOrGenNodeKey(dataPath + "2_peer_id.json")
	if err != nil {
		panic(err)
	}

	sendPubKey := senderPrivKey.PubKey()
	var pubkeyBytes secp256k1.PubKeySecp256k1
	pubkeyBytes = sendPubKey.(secp256k1.PubKeySecp256k1)

	senderB64 := b64.StdEncoding.EncodeToString(pubkeyBytes[:])

	senderAddress := sendPubKey.GetAddress()

	recPubKey := recPrivKey.PubKey()
	recAddress := recPubKey.GetAddress()
	log.Println("Sender Address: " + senderAddress)
	log.Println("Receiver Address: " + recAddress)

	msg := "Send Her Token"

	if err != nil {
		panic(err)
	}
	for i := 45; i <= 45; i++ {

		asset := &protobuf.Asset{
			Category: "crypto",
			Symbol:   "HER",
			Network:  "Herdius",
			Value:    100,
			Fee:      1,
			Nonce:    uint64(i),
		}

		//sig = b64.StdEncoding.EncodeToString(sig)
		tx := protobuf.Tx{
			SenderAddress:   senderAddress,
			SenderPubkey:    senderB64,
			RecieverAddress: recAddress,
			Asset:           asset,
			Message:         msg,
		}

		// Sign the transaction detail
		txbBeforeSign, _ := json.Marshal(tx)

		sig, err := senderPrivKey.PrivKey.Sign(txbBeforeSign)

		tx.Sign = b64.StdEncoding.EncodeToString(sig)

		// Post tx to blockchain.
		txReq := protobuf.TxRequest{
			Tx: &tx,
		}
		txJSON, err := json.Marshal(txReq)
		if err != nil {
			log.Fatalf("Failed to Marshal %v", err)
		}

		response, err := http.Post(endpoint, "application/json", bytes.NewBuffer(txJSON))
		if err != nil {
			log.Fatalf("Failed to Marshal %v", err)
		}

		defer response.Body.Close()
		body, readErr := ioutil.ReadAll(response.Body)
		if readErr != nil {
			log.Fatalf("Failed to read http response %s.", readErr)
		}
		var txResponse protobuf.TxResponse
		jsonErr := json.Unmarshal(body, &txResponse)
		if jsonErr != nil {
			log.Fatalf("Failed to Unmarshal %s.", jsonErr)
		}

		log.Println(txResponse.TxId)
		log.Println(txResponse.Status)
	}
}
