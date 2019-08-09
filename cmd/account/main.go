package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"flag"
	"fmt"
	"log"
	"math/big"

	"github.com/btcd/btcec"
	"github.com/herdius/herdius-core/crypto/secp256k1"
)

/*
	To run it execute the below command
	go run cmd/account/main.go -secret=My-Secret
*/
func main() {
	secret := flag.String("secret", "default", "secret to create private key")
	flag.Parse()
	fmt.Println(GetHerdiusAddress(*secret))
}

// GetHerdiusAddress creates a Herdius Address based
// on the provided secret to private key

func GetHerdiusAddress(secret string) string {
	privateKey := createPrivateKey([]byte(secret))
	var senderPrivKey btcec.PrivateKey
	senderPrivKey = btcec.PrivateKey(*privateKey)

	var sendPubKey secp256k1.PubKeySecp256k1
	copy(sendPubKey[:], senderPrivKey.PubKey().SerializeCompressed())
	senderAddress := sendPubKey.GetAddress()
	return senderAddress
}

// GetPrivateKey ...
func createPrivateKey(ch []byte) *ecdsa.PrivateKey {
	bigInt, err := createBigInt(btcec.S256(), ch)
	if err != nil {
		log.Printf("Failed to create big number : %v", err)
		return nil
	}

	//Create Private Key using Big Int
	privKey, _ := generateKey(btcec.S256(), bigInt)
	return privKey
}

// generateKey generates a public and private key pair.
func generateKey(c elliptic.Curve, k *big.Int) (*ecdsa.PrivateKey, error) {

	priv := new(ecdsa.PrivateKey)
	priv.PublicKey.Curve = c
	priv.D = k
	priv.PublicKey.X, priv.PublicKey.Y = c.ScalarBaseMult(k.Bytes())
	return priv, nil
}

// Create Big Int underlying the given
// curve using the procedure given in [NSA] A.2.1.
func createBigInt(c elliptic.Curve, ch []byte) (k *big.Int, err error) {
	params := c.Params()
	one := new(big.Int).SetInt64(1)

	k = new(big.Int).SetBytes(ch)
	n := new(big.Int).Sub(params.N, one)
	k.Mod(k, n)
	k.Add(k, one)
	return
}
