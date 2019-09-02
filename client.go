package main

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
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

	endpoint = "http://" + endpoint + ":80"
	log.Println("endpoint:", endpoint)
	switch *txType {
	case "all":
		performAllTxs(endpoint)
	case "register":
		sendAccountRegisterTx(endpoint)
	case "external":
		postExternalTx(endpoint)
	case "internal":
		postTx(endpoint)
	case "lock":
		sendLockedTx(endpoint)
	case "redeem":
		sendRedeemTx(endpoint)
	}

}

// This function always registers a new HER Account.
// It randomly creates a private key and uses the same key to register HER account
// Once the new HER Account is added, it will add a ETH address to new HER Account
func performAllTxs(endpoint string) {
	senderPrivKey := secp256k1.GenPrivKey()
	pk64 := b64.StdEncoding.EncodeToString(senderPrivKey[:])
	log.Println("Pk in B64: ", pk64)
	senderAddress := senderPrivKey.PubKey().GetAddress()
	var pubkeyBytes secp256k1.PubKeySecp256k1
	sendPubKey := senderPrivKey.PubKey()
	pubkeyBytes = sendPubKey.(secp256k1.PubKeySecp256k1)

	senderB64 := b64.StdEncoding.EncodeToString(pubkeyBytes[:])
	log.Println("Account Address : " + senderAddress)

	msg := "Register New HER Account"

	asset := &protobuf.Asset{
		Category:              "crypto",
		Symbol:                "HER",
		Network:               "Herdius",
		Value:                 15,
		Fee:                   0,
		Nonce:                 0,
		ExternalSenderAddress: "",
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
	fmt.Println("txbBeforeSign: ", string(txbBeforeSign))
	signature, err := senderPrivKey.Sign(txbBeforeSign)
	if err != nil {
		log.Panic("Failed to sign: ", err)
	}

	tx.Sign = b64.StdEncoding.EncodeToString(signature)

	// Post tx to blockchain.
	txReq := protobuf.TxRequest{
		Tx: &tx,
	}
	txJSON, err := json.Marshal(txReq)
	if err != nil {
		log.Fatalf("Failed to Marshal %v", err)
	}
	response, err := http.Post(endpoint+"/tx", "application/json", bytes.NewBuffer(txJSON))
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

	log.Println("Retrieving new HER Account detail : ", senderAddress)
	getURL := endpoint + "/account/" + senderAddress
	log.Println("Get account url: ", getURL)
	//time.Sleep(16 * time.Second)
	for {
		response, err = http.Get(getURL)
		if err != nil {
			log.Println("Failed to get response: ", err)
		} else {
			defer response.Body.Close()
			body, readErr := ioutil.ReadAll(response.Body)
			if readErr != nil {
				log.Println(fmt.Sprintf("Failed to read http response %s.", readErr))
			}
			if strings.Contains(string(body), "details not found") {
				continue
			}
			log.Println("Account response: ", string(body))
			break
		}
	}

	log.Println("Register an ethereum address to created HER account: ", senderAddress)

	msg = "Add New Tezos Address"
	asset = &protobuf.Asset{
		Category:              "crypto",
		Symbol:                "XTZ",
		Network:               "Herdius",
		Value:                 0,
		Fee:                   0,
		Nonce:                 1,
		ExternalSenderAddress: "tz1WEKDqQrjP49KUBp27AKuzVC2fBv4veirm",
	}

	tx = protobuf.Tx{
		SenderAddress: senderAddress,
		SenderPubkey:  senderB64,
		Message:       msg,
		Type:          "update",
		Asset:         asset,
	}

	// Sign the transaction detail
	txbBeforeSign, _ = json.Marshal(tx)
	fmt.Println("txbBeforeSign: ", string(txbBeforeSign))
	signature, err = senderPrivKey.Sign(txbBeforeSign)
	if err != nil {
		log.Panic("Failed to sign: ", err)
	}

	tx.Sign = b64.StdEncoding.EncodeToString(signature)

	// Post tx to blockchain.
	txReq = protobuf.TxRequest{
		Tx: &tx,
	}
	txJSON, err = json.Marshal(txReq)
	if err != nil {
		log.Fatalf("Failed to Marshal %v", err)
	}
	response, err = http.Post(endpoint+"/tx", "application/json", bytes.NewBuffer(txJSON))
	if err != nil {
		log.Fatalf("Failed to Marshal %v", err)
	}

	defer response.Body.Close()

	if readErr != nil {
		log.Fatalf("Failed to read http response %s.", readErr)
	}

	jsonErr = json.Unmarshal(body, &txResponse)
	if jsonErr != nil {
		log.Fatalf("Failed to Unmarshal %s.", jsonErr)
	}

	log.Println(txResponse.TxId)
	log.Println(txResponse.Status)

	log.Println("Retrieving new HER Account detail : ", senderAddress)
	getURL = endpoint + "/account/" + senderAddress
	//	time.Sleep(16 * time.Second)
	log.Println("Get account url: ", getURL)
	for {
		response, err = http.Get(getURL)
		if err != nil {
			log.Println("Failed to get response: ", err)
		} else {
			defer response.Body.Close()
			body, readErr := ioutil.ReadAll(response.Body)
			if readErr != nil {
				log.Println(fmt.Sprintf("Failed to read http response %s.", readErr))
			}
			if strings.Contains(string(body), "details not found") {
				continue
			}
			var accRes protobuf.AccountResponse
			jsonErr := json.Unmarshal(body, &accRes)
			if jsonErr != nil {
				log.Fatalf("Failed to Unmarshal %s.", jsonErr)
			}
			if len(accRes.EBalances) == 0 {
				continue
			}
			log.Println("Account response: ", accRes)
			break
		}
	}

}

func sendLockedTx(endpoint string) {
	pkBz, _ := b64.StdEncoding.DecodeString("galGfukw3C2vLKLeAJU5bO4gb7KbNQM6qMyiVoPKuow=")
	var senderPrivKey secp256k1.PrivKeySecp256k1
	copy(senderPrivKey[:], pkBz)
	senderAddress := senderPrivKey.PubKey().GetAddress()
	var pubkeyBytes secp256k1.PubKeySecp256k1
	sendPubKey := senderPrivKey.PubKey()
	pubkeyBytes = sendPubKey.(secp256k1.PubKeySecp256k1)

	senderB64 := b64.StdEncoding.EncodeToString(pubkeyBytes[:])
	log.Println("Account Address : " + senderAddress)
	// Lock Ethereum Transaction to Mint HBTC
	log.Println("\nLock Ethereum Transaction to Mint HBTC: ", senderAddress)
	msg := "Locked my ether"
	lockedValue := uint64(0.1 * math.Pow10(18))
	rcvrAddrs := "Hx00000000000000000000000000000000"
	lockedAsset := &protobuf.Asset{
		Category:              "crypto",
		Symbol:                "ETH",
		Network:               "Herdius",
		Value:                 0,
		Fee:                   0,
		Nonce:                 2,
		LockedAmount:          lockedValue,
		ExternalSenderAddress: "0xD8f647855876549d2623f52126CE40D053a2ef6A",
	}
	lockedTx := protobuf.Tx{
		SenderAddress:   senderAddress,
		SenderPubkey:    senderB64,
		RecieverAddress: rcvrAddrs,
		Message:         msg,
		Type:            "lock",
		Asset:           lockedAsset,
	}
	// Sign the transaction detail
	lockedTxbBeforeSign, _ := json.Marshal(lockedTx)
	fmt.Println("lockedTxbBeforeSign: ", string(lockedTxbBeforeSign))
	lockedSignature, err := senderPrivKey.Sign(lockedTxbBeforeSign)
	if err != nil {
		log.Panic("Failed to sign: ", err)
	}
	lockedTx.Sign = b64.StdEncoding.EncodeToString(lockedSignature)
	// Post tx to blockchain.
	lockedTxReq := protobuf.TxRequest{
		Tx: &lockedTx,
	}
	lockedTxJSON, err := json.Marshal(lockedTxReq)
	if err != nil {
		log.Fatalf("Failed to Marshal %v", err)
	}
	response, err := http.Post(endpoint+"/tx", "application/json", bytes.NewBuffer(lockedTxJSON))
	if err != nil {
		log.Fatalf("Failed to Marshal %v", err)
	}
	defer response.Body.Close()
	if err != nil {
		log.Fatalf("Failed to read http response %s.", err)
	}

	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		log.Println(fmt.Sprintf("Failed to read http response %s.", readErr))
	}
	var txResponse protobuf.TxResponse

	jsonErr := json.Unmarshal(body, &txResponse)
	if jsonErr != nil {
		log.Fatalf("Failed to Unmarshal %s.", jsonErr)
	}
	log.Println(txResponse.TxId)
	log.Println(txResponse.Status)
}

func sendRedeemTx(endpoint string) {
	pkBz, _ := b64.StdEncoding.DecodeString("galGfukw3C2vLKLeAJU5bO4gb7KbNQM6qMyiVoPKuow=")
	var senderPrivKey secp256k1.PrivKeySecp256k1
	copy(senderPrivKey[:], pkBz)
	senderAddress := senderPrivKey.PubKey().GetAddress()
	var pubkeyBytes secp256k1.PubKeySecp256k1
	sendPubKey := senderPrivKey.PubKey()
	pubkeyBytes = sendPubKey.(secp256k1.PubKeySecp256k1)

	senderB64 := b64.StdEncoding.EncodeToString(pubkeyBytes[:])
	log.Println("Account Address : " + senderAddress)
	// Lock Ethereum Transaction to Mint HBTC
	log.Println("\nRedeem Ethereum Transaction to Burn HBTC: ", senderAddress)

	msg := "Redeem my ether"
	value := uint64(0.01 * math.Pow10(18))
	rcvrAddrs := "Hx00000000000000000000000000000000"
	redeemAsset := &protobuf.Asset{
		Category:              "crypto",
		Symbol:                "ETH",
		Network:               "Herdius",
		Value:                 0,
		Fee:                   0,
		Nonce:                 3,
		RedeemedAmount:        value,
		ExternalSenderAddress: "0xD8f647855876549d2623f52126CE40D053a2ef6A",
	}
	lockedTx := protobuf.Tx{
		SenderAddress:   senderAddress,
		SenderPubkey:    senderB64,
		RecieverAddress: rcvrAddrs,
		Message:         msg,
		Type:            "redeem",
		Asset:           redeemAsset,
	}
	// Sign the transaction detail
	lockedTxbBeforeSign, _ := json.Marshal(lockedTx)
	fmt.Println("lockedTxbBeforeSign: ", string(lockedTxbBeforeSign))
	lockedSignature, err := senderPrivKey.Sign(lockedTxbBeforeSign)
	if err != nil {
		log.Panic("Failed to sign: ", err)
	}
	lockedTx.Sign = b64.StdEncoding.EncodeToString(lockedSignature)
	// Post tx to blockchain.
	lockedTxReq := protobuf.TxRequest{
		Tx: &lockedTx,
	}
	lockedTxJSON, err := json.Marshal(lockedTxReq)
	if err != nil {
		log.Fatalf("Failed to Marshal %v", err)
	}
	response, err := http.Post(endpoint+"/tx", "application/json", bytes.NewBuffer(lockedTxJSON))
	if err != nil {
		log.Fatalf("Failed to Marshal %v", err)
	}
	defer response.Body.Close()
	if err != nil {
		log.Fatalf("Failed to read http response %s.", err)
	}

	body, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		log.Println(fmt.Sprintf("Failed to read http response %s.", readErr))
	}
	var txResponse protobuf.TxResponse

	jsonErr := json.Unmarshal(body, &txResponse)
	if jsonErr != nil {
		log.Fatalf("Failed to Unmarshal %s.", jsonErr)
	}
	log.Println(txResponse.TxId)
	log.Println(txResponse.Status)
}

func sendAccountRegisterTx(endpoint string) {
	senderPrivKey := secp256k1.GenPrivKey()
	senderAddress := senderPrivKey.PubKey().GetAddress()
	var pubkeyBytes secp256k1.PubKeySecp256k1
	sendPubKey := senderPrivKey.PubKey()
	pubkeyBytes = sendPubKey.(secp256k1.PubKeySecp256k1)

	senderB64 := b64.StdEncoding.EncodeToString(pubkeyBytes[:])
	log.Println("Account Address : " + senderAddress)

	pubkeyBytes = sendPubKey.(secp256k1.PubKeySecp256k1)

	msg := "Register HER account"
	asset := &protobuf.Asset{
		Category:              "crypto",
		Symbol:                "HER",
		Network:               "Herdius",
		Value:                 15,
		Fee:                   0,
		Nonce:                 0,
		ExternalSenderAddress: "0x44c46Ed496B94fafE8A81b9Ab93B27935fcA1603",
	}

	// In case ETH or external asset address is required to be registered
	// use the below Asset Object

	asset = &protobuf.Asset{
		Category:              "crypto",
		Symbol:                "XTZ",
		Network:               "Herdius",
		Value:                 0,
		Fee:                   0,
		Nonce:                 1,
		ExternalSenderAddress: "tz1WEKDqQrjP49KUBp27AKuzVC2fBv4veirm",
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

	sig, err := senderPrivKey.Sign(txbBeforeSign)

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

func postExternalTx(endpoint string) {
	// Create key pairs and store in a local file
	// User 1
	// Address: HHy1CuT3UxCGJ3BHydLEvR5ut6HRy2qUvm
	senderPrivKey, err := key.LoadOrGenNodeKey("./tempKeySign.json")
	if err != nil {
		panic(err)
	}
	// User 2
	// Address: HKTXmdsHyZn1B2ErRKiG4iN34YixCgdQgx
	rcvrPrivKey, err := key.LoadOrGenNodeKey("./tempKeySign.json")

	sendPubKey := senderPrivKey.PubKey()
	var pubkeyBytes secp256k1.PubKeySecp256k1
	pubkeyBytes = sendPubKey.(secp256k1.PubKeySecp256k1)

	senderB64 := b64.StdEncoding.EncodeToString(pubkeyBytes[:])

	senderAddress := sendPubKey.GetAddress()

	recPubKey := rcvrPrivKey.PubKey()
	recAddress := recPubKey.GetAddress()
	log.Println("Sender Address: " + senderAddress)
	log.Println("Receiver Address: " + recAddress)

	msg := "Send ETH Tokens"
	asset := &protobuf.Asset{
		Category: "crypto",
		Symbol:   "ETH",
		Network:  "Herdius",
		Value:    1,
		Fee:      1,
		Nonce:    uint64(2),
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
	fmt.Printf("txJSON: %v\n", string(txJSON))
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
	recAddress = "HFw2Qo2ZCaqxDqv8CsYmiczhuCz5UeGdXb"
	if err != nil {
		panic(err)
	}
	for i := 1; i <= 1; i++ {

		asset := &protobuf.Asset{
			Category: "crypto",
			Symbol:   "HER",
			Network:  "Herdius",
			//Value:    100,
			Value: 500,
			Fee:   1,
			Nonce: uint64(5),
		}

		//sig = b64.StdEncoding.EncodeToString(sig)
		tx := protobuf.Tx{
			SenderAddress:   senderAddress,
			SenderPubkey:    senderB64,
			RecieverAddress: recAddress,
			Asset:           asset,
			Message:         msg,
			//Type:            "update",
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

		response, err := http.Post(endpoint+"/tx", "application/json", bytes.NewBuffer(txJSON))
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
