package testutils

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/herdius/herdius-blockchain-api/protobuf"
	cryptoAmino "github.com/herdius/herdius-core/crypto/encoding/amino"
	"github.com/herdius/herdius-core/crypto/secp256k1"
	"github.com/herdius/herdius-core/p2p/key"
	amino "github.com/tendermint/go-amino"
)

var cdc = amino.NewCodec()

func init() {
	cryptoAmino.RegisterAmino(cdc)
}

// 1) master account have balance
// 2) master will send her

var endpoint string

type CMDUtils struct {
	endpoint string
}

//New Create Instance of CMDUtils
func New(ep string) *CMDUtils {
	return &CMDUtils{endpoint: ep}

}

func (cmd *CMDUtils) RegisterAccount(senderPrivKey key.NodeKey, accountType string, nonce uint64, externalSenderAddressstring string) {
	senderAddress := senderPrivKey.PubKey().GetAddress()
	var pubkeyBytes secp256k1.PubKeySecp256k1
	sendPubKey := senderPrivKey.PubKey()
	pubkeyBytes = sendPubKey.(secp256k1.PubKeySecp256k1)
	senderB64 := b64.StdEncoding.EncodeToString(pubkeyBytes[:])
	log.Println("Account Address : " + senderAddress)
	msg := "Register account"
	asset := &protobuf.Asset{
		Category:              "crypto",
		Symbol:                accountType,
		Network:               "Herdius",
		Value:                 15,
		Fee:                   0,
		Nonce:                 nonce,
		ExternalSenderAddress: externalSenderAddressstring,
	}
	tx := protobuf.Tx{
		SenderAddress: senderAddress,
		SenderPubkey:  senderB64,
		Message:       msg,
		Type:          "update",
		Asset:         asset,
	}
	signedTX, err := signTX(senderPrivKey, tx)
	log.Println(signedTX)
	res, err := postTx(cmd.endpoint, signedTX)
	log.Println(res, err)
}

//ToTx send balance to from master account
func (cmd *CMDUtils) ToTx(senderPrivKey key.NodeKey, accountType string, nonce uint64, value uint64, recAddress string) (txResponse protobuf.TxResponse, err error) {
	senderAddress := senderPrivKey.PubKey().GetAddress()
	var pubkeyBytes secp256k1.PubKeySecp256k1
	sendPubKey := senderPrivKey.PubKey()
	pubkeyBytes = sendPubKey.(secp256k1.PubKeySecp256k1)
	senderB64 := b64.StdEncoding.EncodeToString(pubkeyBytes[:])
	log.Println("Account Address : " + senderAddress)
	msg := "Send asset"
	asset := &protobuf.Asset{
		Category: "crypto",
		Symbol:   accountType,
		Network:  "Herdius",
		Value:    value,
		Fee:      1,
		Nonce:    nonce,
	}
	tx := protobuf.Tx{
		SenderAddress:   senderAddress,
		SenderPubkey:    senderB64,
		RecieverAddress: recAddress,
		Message:         msg,
		Asset:           asset,
	}
	signedTX, err := signTX(senderPrivKey, tx)
	return postTx(cmd.endpoint, signedTX)

}

func signTX(pk key.NodeKey, tx protobuf.Tx) (protobuf.Tx, error) {
	txbBeforeSign, err := json.Marshal(tx)
	if err != nil {
		return tx, err
	}
	signature, err := pk.PrivKey.Sign(txbBeforeSign)
	if err != nil {
		return tx, err
	}
	tx.Sign = b64.StdEncoding.EncodeToString(signature)
	return tx, nil
}

func postTx(endpoint string, tx protobuf.Tx) (txResponse protobuf.TxResponse, err error) {
	txReq := protobuf.TxRequest{
		Tx: &tx,
	}
	txJSON, err := json.Marshal(txReq)
	if err != nil {
		log.Println("Error marshalling json")
		return
	}
	response, err := http.Post(endpoint+"/tx", "application/json", bytes.NewBuffer(txJSON))
	if err != nil {
		log.Println("Error http post to endpoint :", endpoint)
		return
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Error reading ioutil")

		return
	}
	log.Println(string(body))
	err = json.Unmarshal(body, &txResponse)
	log.Println("---------")
	return
}

func (cmd *CMDUtils) GetAccount(address string) (accRes protobuf.AccountResponse, err error) {
	getURL := cmd.endpoint + "/account/" + address
	response, err := http.Get(getURL)
	if err != nil {
		log.Println("Failed to get response: ", err)
		return
	}
	defer response.Body.Close()
	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		log.Println(fmt.Sprintf("Failed to read http response %s.", readErr))
	}
	if strings.Contains(string(body), "details not found") {
		return
	}
	log.Println("Account response: ", string(body))
	json.Unmarshal(body, &accRes)
	return

}

type HerKey struct {
	PrivateKey Secret `json:"priv_key"`
}

type Secret struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func GetHERKeyPair(secret string) (nodeKey key.NodeKey, err error) {
	sec := Secret{Type: "herdius/PrivKeySecp256k1", Value: secret}
	key1 := HerKey{PrivateKey: sec}
	jsonBytes, err := json.Marshal(key1)
	err = cdc.UnmarshalJSON(jsonBytes, &nodeKey)
	return

}
