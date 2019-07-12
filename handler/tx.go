package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/herdius/herdius-blockchain-api/config"
	"github.com/herdius/herdius-blockchain-api/protobuf"
	"github.com/herdius/herdius-blockchain-api/store"
	"github.com/herdius/herdius-core/p2p/network"
)

// TxServiceI is transaction service interface over blockchain
type TxServiceI interface {
	GetTx(string, *network.Network, string) (*protobuf.TxDetailResponse, error)
	GetTxsByAddress(string, *network.Network, string) (*protobuf.TxsResponse, error)
	GetTxsByAssetAndAddress(string, string, *network.Network, string) (*protobuf.TxsResponse, error)
	GetLockedTxsByBlockNumber(int64, *network.Network, string) (*protobuf.TxLockedResponse, error)
	PutUpdateTxByTxID(*protobuf.TxUpdateRequest, *network.Network, string) (*protobuf.TxUpdateResponse, error)
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
		if s := getStore(configuration.DBConnString()); s != nil && msg.TxId != "" {
			if err := s.Save(&store.Tx{ID: msg.TxId, Status: store.StatusPending}); err != nil {
				log.Printf("Failed to save Tx to database: %v", err)
			}
			log.Printf("Tx saved to database")
		}
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
	configuration := config.GetConfiguration(env)
	s := getStore(configuration.DBConnString())
	if s != nil {
		if tx, err := s.Get(id); err == nil {
			json.NewEncoder(w).Encode(tx.ToTxDetailResponse())
			return
		}
	}

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
	configuration := config.GetConfiguration(env)
	s := getStore(configuration.DBConnString())
	if s != nil {
		if txs, err := s.GetBySender(address); err == nil && len(txs) > 0 {
			res := &protobuf.TxsResponse{}
			res.Txs = make([]*protobuf.TxDetailResponse, len(txs))
			for i, tx := range txs {
				res.Txs[i] = tx.ToTxDetailResponse()
			}
			json.NewEncoder(w).Encode(res)
			return
		}
	}
	srv := TxService{}
	txs, err := srv.GetTxsByAddress(address, net, env)

	if err != nil {
		log.Println(err.Error())
		json.NewEncoder(w).Encode(err.Error())
	} else {
		if len(txs.Txs) == 0 {
			json.NewEncoder(w).Encode("No transactions found")
		} else {
			if s != nil {
				for _, txDetail := range txs.Txs {
					s.Save(store.FromTxDetailResponse(txDetail))
				}
			}
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
	configuration := config.GetConfiguration(env)
	s := getStore(configuration.DBConnString())
	if s != nil {
		if txs, err := s.GetByAssetAndSender(asset, address); err == nil && len(txs) > 0 {
			res := &protobuf.TxsResponse{}
			res.Txs = make([]*protobuf.TxDetailResponse, len(txs))
			for i, tx := range txs {
				res.Txs[i] = tx.ToTxDetailResponse()
			}
			json.NewEncoder(w).Encode(res)
			return
		}
	}
	srv := TxService{}
	txs, err := srv.GetTxsByAssetAndAddress(asset, address, net, env)

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		if len(txs.Txs) == 0 {
			json.NewEncoder(w).Encode("No transactions found")
		} else {
			if s != nil {
				for _, txDetail := range txs.Txs {
					s.Save(store.FromTxDetailResponse(txDetail))
				}
			}
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

// PutUpdateTxByTxID forwards a request to update a TX that is currently queued in the Supervisor memory pool
func PutUpdateTxByTxID(w http.ResponseWriter, r *http.Request, net *network.Network, env string) {
	params := mux.Vars(r)
	if len(params["id"]) == 0 {
		json.NewEncoder(w).Encode("Request invalid, 'id' param missing\n")
		return
	}
	id := params["id"]
	var txRequest protobuf.TxUpdateRequest
	err := json.NewDecoder(r.Body).Decode(&txRequest.Tx)
	if err != nil {
		json.NewEncoder(w).Encode("\nRequest invalid, Could not parse PUT json data, invalid format, err:\n" + err.Error())
		return
	}
	txRequest.TxId = id
	log.Println("Update request received for tx:", id)
	srv := TxService{}
	res, err := srv.PutUpdateTxByTxID(&txRequest, net, env)

	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(res)
	}
}

// PutUpdateTxByTxID forwards a request to update a TX that is currently queued in the Supervisor memory pool
func (t *TxService) PutUpdateTxByTxID(txRequest *protobuf.TxUpdateRequest, net *network.Network, env string) (*protobuf.TxUpdateResponse, error) {
	configuration := config.GetConfiguration(env)
	supervisorAddress := configuration.GetSupervisorAddress()
	ctx := network.WithSignMessage(context.Background(), true)
	supervisorNode, err := net.Client(supervisorAddress)

	res, err := supervisorNode.Request(ctx, txRequest)
	if err != nil {
		log.Println("Failed to get all txs detail due to: " + err.Error())
		return nil, fmt.Errorf("Failed to get all txs detail due to: %v", err)
	}

	switch msg := res.(type) {
	case *protobuf.TxUpdateResponse:
		log.Printf("Tx Detail: %v", msg)
		if s := getStore(configuration.DBConnString()); s != nil && msg.TxId != "" {
			txDetailReq := protobuf.TxDetailRequest{TxId: msg.TxId}
			res, err := supervisorNode.Request(ctx, &txDetailReq)
			if err != nil {
				log.Printf("Failed to get TX after updating")
			} else {
				if txDetail, ok := res.(*protobuf.TxDetailResponse); ok {
					if err := s.Update(store.FromTxDetailResponse(txDetail)); err != nil {
						log.Printf("Failed to update Tx in database: %v", err)
					}
					log.Printf("Tx updated in database")
				}
			}
		}
		return msg, nil
	}
	return nil, nil
}

// CancelRequest forwards a request to cancel a TX that is currently queued in the Supervisor memory pool
func DeleteTx(w http.ResponseWriter, r *http.Request, net *network.Network, env string) {
	params := mux.Vars(r)
	if len(params["id"]) == 0 {
		json.NewEncoder(w).Encode("Request invalid, 'id' param missing\n")
		return
	}

	id := params["id"]
	log.Println("tx id:", id)

	configuration := config.GetConfiguration(env)
	supervisorAddress := configuration.GetSupervisorAddress()
	ctx := network.WithSignMessage(context.Background(), true)
	supervisorNode, err := net.Client(supervisorAddress)

	req := &protobuf.TxDeleteRequest{
		TxId: id,
	}

	res, err := supervisorNode.Request(ctx, req)
	if err != nil {
		log.Println("Failed to cancel tx: " + err.Error())
		json.NewEncoder(w).Encode(fmt.Sprintf("Failed to cancel tx: %v", err))
	}

	switch msg := res.(type) {
	case *protobuf.TxUpdateResponse:
		log.Printf("Tx Detail: %v", msg)
		json.NewEncoder(w).Encode(msg)
	}
}

// GetLockedTxsByBlockNumber returns all locked txs in a block by its number
func GetLockedTxsByBlockNumber(w http.ResponseWriter, r *http.Request, net *network.Network, env string) {
	params := mux.Vars(r)
	if len(params["block_number"]) == 0 {
		json.NewEncoder(w).Encode("Request invalid, 'block_number' param missing\n")
		return
	}

	blockNumber, err := strconv.ParseInt(params["block_number"], 10, 64)
	if err != nil {
		json.NewEncoder(w).Encode("Request invalid, invalid 'block_number'\n")
		return
	}

	srv := TxService{}
	txResp, err := srv.GetLockedTxsByBlockNumber(blockNumber, net, env)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	if txResp == nil {
		json.NewEncoder(w).Encode("No locked transactions found in block")
	} else {
		json.NewEncoder(w).Encode(txResp.Txs)
	}
}

// GetLockedTxsByBlockNumber ...
func (t *TxService) GetLockedTxsByBlockNumber(blockNumber int64, net *network.Network, env string) (*protobuf.TxLockedResponse, error) {
	configuration := config.GetConfiguration(env)
	supervisorAddress := configuration.GetSupervisorAddress()
	ctx := network.WithSignMessage(context.Background(), true)
	supervisorNode, err := net.Client(supervisorAddress)

	req := &protobuf.TxLockedRequest{BlockNumber: blockNumber}

	res, err := supervisorNode.Request(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get locked txs due to: %v", err)
	}

	switch msg := res.(type) {
	case *protobuf.TxLockedResponse:
		return msg, nil
	}
	return nil, nil
}
