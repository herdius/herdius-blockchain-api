package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/herdius/herdius-blockchain-api/config"
	"github.com/herdius/herdius-blockchain-api/protobuf"
	"github.com/herdius/herdius-core/p2p/network"
)

// TxServiceI is transaction service interface over blockchain
type TxServiceI interface {
	GetTx(id string, net *network.Network, env string) (*protobuf.TxDetailResponse, error)
}

// TxService ...
type TxService struct{}

var (
	_ TxServiceI = (*TxService)(nil)
)

func (s *service) PostTx(txReq protobuf.TxRequest, net *network.Network, env string) (*protobuf.TxResponse, error) {
	configuration := config.GetConfiguration(env)
	supervisorAddress := configuration.GetSupervisorAddress()

	ctx := network.WithSignMessage(context.Background(), true)

	supervisorNode, _ := net.Client(supervisorAddress)
	res, err := supervisorNode.Request(ctx, &txReq)

	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Failed to find block due to :%v", err))
	}

	switch msg := res.(type) {
	case *protobuf.TxResponse:
		log.Printf("Tx ID: %v", msg.TxId)
		log.Printf("Tx status: %v", msg.Status)
		return msg, nil
	}

	return nil, nil
}

// PostTx ...
func PostTx(w http.ResponseWriter, r *http.Request, net *network.Network, env string) {
	var txRequest protobuf.TxRequest
	err := json.NewDecoder(r.Body).Decode(&txRequest)
	if err != nil {
		json.NewEncoder(w).Encode("\nRequest invalid, Could not parse POST json data, invalid format, err:\n" + err.Error())
		return
	}

	// Check if tx type is account update
	if len(txRequest.Tx.Type) > 0 && strings.EqualFold(txRequest.Tx.Type, "update") {
		if len(txRequest.Tx.SenderAddress) == 0 ||
			len(txRequest.Tx.Sign) == 0 ||
			len(txRequest.Tx.SenderPubkey) == 0 {
			json.NewEncoder(w).Encode("\nRequest missing data, POST body did not include all required values\n")
			json.NewEncoder(w).Encode("\ntxreq:\n")
			json.NewEncoder(w).Encode(txRequest)
			return
		}
		srv := service{}

		newTxResponse, err := srv.PostTx(txRequest, net, env)

		if err != nil {
			json.NewEncoder(w).Encode(err.Error())
			return
		}
		json.NewEncoder(w).Encode(newTxResponse)
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
	srv := service{}
	newTxResponse, err := srv.PostTx(txRequest, net, env)

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode(newTxResponse)
	}
}

// GetTx will retrieve a transaction detail from blockchain
// for a give tx id
func GetTx(w http.ResponseWriter, r *http.Request, net *network.Network, env string) {
	params := mux.Vars(r)
	if len(params["id"]) == 0 {
		json.NewEncoder(w).Encode("Request invalid, 'id' param missing\n")
		return
	}

	id := params["id"]
	srv := TxService{}
	txDetailRes, err := srv.GetTx(id, net, env)
	if err != nil {
		json.NewEncoder(w).Encode("Failed to retrieve Tx detail for id: " + id)
	} else {
		json.NewEncoder(w).Encode(txDetailRes)
	}
}

// GetTx ...
func (t *TxService) GetTx(id string, net *network.Network, env string) (*protobuf.TxDetailResponse, error) {
	configuration := config.GetConfiguration(env)
	supervisorAddress := configuration.GetSupervisorAddress()
	ctx := network.WithSignMessage(context.Background(), true)
	supervisorNode, _ := net.Client(supervisorAddress)
	txDetailReq := protobuf.TxDetailRequest{
		TxId: id,
	}
	res, err := supervisorNode.Request(ctx, &txDetailReq)
	if err != nil {
		log.Println("Failed to get tx detail due to: " + err.Error())
		return nil, fmt.Errorf("Failed to get tx detail due to: %v", err)
	}

	switch msg := res.(type) {
	case *protobuf.TxDetailResponse:
		log.Printf("Tx Detail: %v", msg)
		return msg, nil
	}
	return nil, nil
}

// GetTxsByAddress ...
func GetTxsByAddress(w http.ResponseWriter, r *http.Request, net *network.Network, env string) {
	params := mux.Vars(r)
	if len(params["address"]) == 0 {
		json.NewEncoder(w).Encode("Request invalid, 'address' param missing\n")
		return
	}

	address := params["address"]
	srv := TxService{}
	txs, err := srv.GetTxsByAddress(address, net, env)

	if err != nil {
		log.Println(err.Error())
		json.NewEncoder(w).Encode(err.Error())
	} else {
		if len(txs.Txs) == 0 {
			json.NewEncoder(w).Encode("No transactions found")
		} else {
			json.NewEncoder(w).Encode(txs.Txs)
		}
	}
}

// GetTxsByAddress ...
func (t *TxService) GetTxsByAddress(address string, net *network.Network, env string) (*protobuf.TxsResponse, error) {
	configuration := config.GetConfiguration(env)
	supervisorAddress := configuration.GetSupervisorAddress()
	ctx := network.WithSignMessage(context.Background(), true)
	supervisorNode, err := net.Client(supervisorAddress)

	req := &protobuf.TxsByAddressRequest{
		Address: address,
	}

	res, err := supervisorNode.Request(ctx, req)
	if err != nil {
		log.Println("Failed to get all txs detail due to: " + err.Error())
		return nil, fmt.Errorf("Failed to get all txs detail due to: %v", err)
	}

	switch msg := res.(type) {
	case *protobuf.TxsResponse:
		log.Printf("Tx Detail: %v", msg)
		return msg, nil
	}
	return nil, nil
}

// GetTxsByAssetAndAddress ...
func GetTxsByAssetAndAddress(w http.ResponseWriter, r *http.Request, net *network.Network, env string) {
	params := mux.Vars(r)
	if len(params["asset"]) == 0 {
		json.NewEncoder(w).Encode("Request invalid, 'asset' param missing\n")
		return
	}
	if len(params["address"]) == 0 {
		json.NewEncoder(w).Encode("Request invalid, 'address' param missing\n")
		return
	}

	address := params["address"]
	asset := params["asset"]
	srv := TxService{}
	txs, err := srv.GetTxsByAssetAndAddress(asset, address, net, env)

	if err != nil {
		log.Println(err.Error())
		json.NewEncoder(w).Encode(err.Error())
	} else {
		if len(txs.Txs) == 0 {
			json.NewEncoder(w).Encode("No transactions found")
		} else {
			json.NewEncoder(w).Encode(txs.Txs)
		}

	}
}

// GetTxsByAssetAndAddress ...
func (t *TxService) GetTxsByAssetAndAddress(asset, address string, net *network.Network, env string) (*protobuf.TxsResponse, error) {
	configuration := config.GetConfiguration(env)
	supervisorAddress := configuration.GetSupervisorAddress()
	ctx := network.WithSignMessage(context.Background(), true)
	supervisorNode, err := net.Client(supervisorAddress)

	req := &protobuf.TxsByAssetAndAddressRequest{
		Address: address,
		Asset:   asset,
	}

	res, err := supervisorNode.Request(ctx, req)
	if err != nil {
		log.Println("Failed to get all txs detail due to: " + err.Error())
		return nil, fmt.Errorf("Failed to get all txs detail due to: %v", err)
	}

	switch msg := res.(type) {
	case *protobuf.TxsResponse:
		log.Printf("Tx Detail: %v", msg)
		return msg, nil
	}
	return nil, nil
}
