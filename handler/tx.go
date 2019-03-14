package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/herdius/herdius-core/p2p/network"
	"github.com/rs/zerolog/log"
)

// TXMessagePlugin will receive all Transaction specific messages.
type TXMessagePlugin struct{ *network.Plugin }
type TxReq struct {
	SenderAddress string `json:senderaddress`
	SenderPubKey  string `json:senderpubkey`
	RecAddress    string `json:recaddress`
	Signature     string `json:signature`
	Asset         struct {
		Message  string  `json:message`
		Network  string  `json:network`
		Category string  `json:category`
		Currency string  `json:currency`
		Value    float64 `json:value`
		Fee      float64 `json:fee`
	} `json:asset`
}

// PostTransaction handler called by http.HandleFunc
func PostTransaction(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	var txReq TxReq
	err = json.Unmarshal(body, &txReq)
	if err != nil {
		log.Info().Msg("Could not parse POST json data, invalid format")
		fmt.Fprint(w, "Request invalid, Could not parse POST json data, invalid format, err: %v\n", err)
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
