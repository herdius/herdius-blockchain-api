package main

import (
	"flag"
	"fmt"

	"github.com/btcd/btcec"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/herdius/herdius-core/crypto/secp256k1"
)

/*
	To get herdius address against a hexified private key, execute the below command:
	go run cmd/account/main.go -hexprivkey=00000000000000000000000000000000000000000000004d792d536563726575
*/
func main() {
	hexPrivateKey := flag.String("hexprivkey", "default", "Hexified Private Key")
	flag.Parse()
	fmt.Println(GetHerdiusAddress(*hexPrivateKey))
}

// GetHerdiusAddress creates a Herdius Address based
// on the provided hexPrivateKey to private key
func GetHerdiusAddress(hexPrivateKey string) string {
	var senderPrivKey btcec.PrivateKey
	ecdsaPK, _ := crypto.HexToECDSA(hexPrivateKey)
	senderPrivKey = btcec.PrivateKey(*ecdsaPK)

	var sendPubKey secp256k1.PubKeySecp256k1
	copy(sendPubKey[:], senderPrivKey.PubKey().SerializeCompressed())
	senderAddress := sendPubKey.GetAddress()
	return senderAddress
}
