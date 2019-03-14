package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/herdius/herdius-core/p2p/network"
	"github.com/rs/zerolog/log"
)

// TXMessagePlugin will receive all Transaction specific messages.
type TXMessagePlugin struct{ *network.Plugin }
type TxReq struct {
	SenderAddress string `json:sender_address`
	SenderPubKey  string `json:sender_pubkey`
	RecAddress    string `json:rec_address`
	Signature     string `json:signature`
	Asset         struct {
		Network  string  `json:network`
		Category string  `json:category`
		Currency string  `json:currency`
		Value    float64 `json:value`
		Fee      float64 `json:fee`
	}
}

// PostTransaction handler called by http.HandleFunc
func PostTransaction(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var txReq TxReq
	err := decoder.Decode(&txReq)
	if err != nil {
		log.Info().Msg("Could not parse POST json data, invalid format")
		fmt.Fprint(w, "Request invalid, Could not parse POST json data, invalid format\n")
	}
	if txReq.SenderAddress == "" ||
		txReq.SenderPubKey == "" ||
		txReq.RecAddress == "" ||
		txReq.Signature == "" ||
		txReq.Asset.Network == "" ||
		txReq.Asset.Category == "" ||
		txReq.Asset.Currency == "" ||
		txReq.Asset.Value < 0 ||
		txReq.Asset.Fee < 0 {
		log.Info().Msg("POST body did not include all required values")
		fmt.Fprint(w, "Request invalid, POST body did not include all required values\n")
	}

	fmt.Fprint(w, txReq)
}
