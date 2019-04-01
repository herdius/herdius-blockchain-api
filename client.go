package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	b64 "encoding/base64"
	"encoding/hex"
	"encoding/json"

	"github.com/herdius/herdius-blockchain-api/protobuf"
	"github.com/herdius/herdius-core-dev/crypto/herhash"
	"github.com/herdius/herdius-core-dev/p2p/key"
)

/**
This file is just for testing purpose. It will be removed.
*/

func main() {

	dataPath := "/Users/lingrajmahanand/go/src/github.com/herdius/herdius-core-dev/cmd/testdata/secp205k1Accts/"
	senderPrivKey, err := key.LoadOrGenNodeKey(dataPath + "1_peer_id.json")
	if err != nil {
		panic(err)
	}
	recPrivKey, err := key.LoadOrGenNodeKey(dataPath + "2_peer_id.json")
	if err != nil {
		panic(err)
	}

	sendPubKey := senderPrivKey.PubKey()
	senderB64 := b64.StdEncoding.EncodeToString(sendPubKey.Bytes())
	senderAddress := sendPubKey.GetAddress()

	recPubKey := recPrivKey.PubKey()
	recAddress := recPubKey.GetAddress()

	//	receiverB64 := b64.StdEncoding.EncodeToString(recPubKey.Bytes())

	msg := "Send Her Token"

	if err != nil {
		panic(err)
	}

	asset := &protobuf.Asset{
		Category: "crypto",
		Symbol:   "HER",

		Network: "Herdius",
		Value:   10,
		Fee:     0,
		Nonce:   2,
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

	// decodedSig, _ := b64.StdEncoding.DecodeString(tx.Sign)

	// fmt.Println(sendPubKey.VerifyBytes(txbBeforeSign, decodedSig))

	// Post tx to blockchain.
	txReq := protobuf.TxRequest{
		Tx: &tx,
	}
	txJSON, err := json.Marshal(txReq)
	if err != nil {
		log.Fatalf("Failed to Marshal %v", err)
	}

	response, err := http.Post("http://localhost:80/tx", "application/json", bytes.NewBuffer(txJSON))
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

// CreateTxID ...
func CreateTxID(txbBeforeSign []byte) string {
	txhash := herhash.Sum(txbBeforeSign)
	flen := 20

	if len(txhash) == flen {
		return "H" + hex.EncodeToString(txhash)
	}
	if len(txhash) > flen {

		return "H" + hex.EncodeToString(txhash[len(txhash)-flen:])
	}
	hh := make([]byte, flen)
	copy(hh[flen-len(txhash):flen], txhash)

	return "H" + hex.EncodeToString(hh)
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
