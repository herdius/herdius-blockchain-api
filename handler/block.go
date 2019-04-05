package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/herdius/herdius-blockchain-api/protobuf"
	protoplugin "github.com/herdius/herdius-blockchain-api/protobuf"
	"github.com/herdius/herdius-core/p2p/log"
	"github.com/herdius/herdius-core/p2p/network"
)

// NB is a Network Builder
var NB *NetworkBuilder

// Response Tracker from Blockchain
var (
	blockByHeight = 0
	blockResponse = Block{}
)

func init() {
	// Setup network specific configuration to call blockchain services
	NB = CreateNetworkBuilder()
}

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
	GetBlockByHeight(height uint64) error
}

// Service ...
type service struct{}

var (
	_ BlockI = (*service)(nil)
)

func (s *service) GetBlockByHeight(height uint64) error {
	net, err := NB.builder.Build()
	NB.GetNetworkBuilder()
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("Failed to build network:%v", err))
	}
	go net.Listen()
	defer net.Close()

	peer := "tcp://localhost:3000"
	peers := make([]string, 1)
	peers = append(peers, peer)
	bootStrap(net, peers)

	ctx := network.WithSignMessage(context.Background(), true)

	net.Broadcast(ctx, &protoplugin.BlockHeightRequest{BlockHeight: height})
	time.Sleep(2 * time.Second)

	// TODO: Need to have a better implementation rather than using
	// the global variable blockByHeight
	blockByHeight = 1
	return nil
}

// BlockMessagePlugin will receive all Block specific messages.
type BlockMessagePlugin struct{ *network.Plugin }

// Receive handles block specific received messages
func (state *BlockMessagePlugin) Receive(ctx *network.PluginContext) error {
	switch msg := ctx.Message().(type) {
	case *protobuf.BlockResponse:
		log.Info().Msgf("Block Response: %v", msg)
		log.Info().Msgf("Block blockByHeight: %v", blockByHeight)

		blockResponse.BlockHeight = msg.BlockHeight
		blockResponse.Transactions = msg.Transactions
		blockResponse.Timestamp = msg.Time.Seconds
		blockResponse.SupervisorAddress = msg.SupervisorAddress
		blockByHeight = 0

		log.Info().Msgf("Block Response: %v", blockResponse)
	}
	return nil
}

func bootStrap(net *network.Network, peers []string) {
	if len(peers) > 0 {
		net.Bootstrap(peers...)
	}
}

////////////////////////////////////////////////
// Block Handlers
///////////////////////////////////////////////

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
	err = srv.GetBlockByHeight(uint64(height))
	if err != nil {
		fmt.Fprint(w, err)
	} else {
		log.Info().Msgf("Processed for Block Height: %d", blockResponse.BlockHeight)
		fmt.Fprint(w, blockResponse)
	}

}
