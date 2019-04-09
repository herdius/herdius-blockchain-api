package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/herdius/herdius-blockchain-api/config"
	nb "github.com/herdius/herdius-blockchain-api/network"
	"github.com/herdius/herdius-blockchain-api/protobuf"
	protoplugin "github.com/herdius/herdius-blockchain-api/protobuf"
	"github.com/herdius/herdius-core/p2p/log"
	"github.com/herdius/herdius-core/p2p/network"
)

// Block will hold block detail retrieved from blockchain
type Block struct {
	BlockHeight       uint64
	Timestamp         int64 // Timestamp of block intialization
	Transactions      uint64
	SupervisorAddress string // Herdius Address of Supervisor node
	StateRoot         []byte
}

// BlockI is an interface to provide block specific services
type BlockI interface {
	GetBlockByHeight(height uint64) (*protobuf.BlockResponse, error)
}

// Service ...
type service struct{}

var (
	_ BlockI = (*service)(nil)
)

func (s *service) GetBlockByHeight(height uint64) (*protobuf.BlockResponse, error) {
	net, err := nb.GetNetworkBuilder().Build()
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Failed to build network:%v", err))
	}

	go net.Listen()
	defer net.Close()

	configuration := config.GetConfiguration()

	supervisorAddress := configuration.TCP + "://" + configuration.SupervisorHost + ":" + strconv.Itoa(configuration.SupervisorPort)

	ctx := network.WithSignMessage(context.Background(), true)

	supervisorNode, _ := net.Client(supervisorAddress)
	res, err := supervisorNode.Request(ctx, &protoplugin.BlockHeightRequest{BlockHeight: height})

	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Failed to find block due to :%v", err))
	}

	switch msg := res.(type) {
	case *protobuf.BlockResponse:
		return msg, nil
	}

	return nil, nil
}

func bootStrap(net *network.Network, peers []string) {
	if len(peers) > 0 {
		net.Bootstrap(peers...)
	}
}

// GetBlockByHeight handler
func GetBlockByHeight(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if len(params["height"]) == 0 {
		json.NewEncoder(w).Encode("Request invalid, 'height' param missing\n")
		return
	}

	heightJSON := params["height"]

	height, err := strconv.ParseInt(heightJSON, 10, 64)

	if err != nil {
		log.Error().Msgf("Failed to Parse %v", err)
	}

	srv := service{}
	block, err := srv.GetBlockByHeight(uint64(height))
	if err != nil {
		fmt.Fprint(w, err)
	}
	if block.Time != nil {
		b := Block{
			BlockHeight:       block.BlockHeight,
			Timestamp:         block.Time.Seconds,
			Transactions:      block.Transactions,
			SupervisorAddress: block.SupervisorAddress,
			StateRoot:         block.StateRoot,
		}
		log.Info().Msgf("Processed for Block Height: %d", block.BlockHeight)
		fmt.Fprint(w, b)

	} else {
		fmt.Fprint(w, "Block not found for block height: "+heightJSON)
	}

}
