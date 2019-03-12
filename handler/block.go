package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

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
	BlockHeight       int64
	Timestamp         int64 // Timestamp of block intialization
	Transactions      int32
	SupervisorAddress string // Herdius Address of Supervisor node
	StateRoot         []byte
}

// BlockI is an interface to provide block specific services
type BlockI interface {
	GetBlockByHeight(height int64) (*Block, error)
}

// Service ...
type service struct{}

var (
	_ BlockI = (*service)(nil)
)

func (s *service) GetBlockByHeight(height int64) (*Block, error) {
	net, err := NB.builder.Build()
	if err != nil {
		log.Error().Msgf("Failed to build network:%v", err)
	}
	go net.Listen()
	defer net.Close()

	peer := "tcp://localhost:3000"
	peers := make([]string, 1)
	peers = append(peers, peer)
	bootStrap(net, peers)

	ctx := network.WithSignMessage(context.Background(), true)

	net.Broadcast(ctx, &protoplugin.BlockHeightRequest{BlockHeight: 2})
	time.Sleep(2 * time.Second)
	block := Block{
		BlockHeight:       1,
		Timestamp:         79273,
		Transactions:      500,
		SupervisorAddress: "xafaljowx",
	}
	blockByHeight = 1
	return &block, nil
}

// BlockMessagePlugin will receive all Block specific messages.
type BlockMessagePlugin struct{ *network.Plugin }

// Receive handles block specific received messages
func (state *BlockMessagePlugin) Receive(ctx *network.PluginContext) error {
	switch msg := ctx.Message().(type) {
	case *protobuf.BlockResponse:
		log.Info().Msgf("Block Response: %v", msg)
		log.Info().Msgf("Block blockByHeight: %v", blockByHeight)
		//if blockByHeight == 1 {
		//blockResponse.mux.Lock()
		blockResponse.BlockHeight = msg.BlockHeight
		blockResponse.Transactions = msg.Transactions
		blockResponse.Timestamp = msg.Time.Seconds
		blockResponse.SupervisorAddress = msg.SupervisorAddress
		blockByHeight = 0
		//blockResponse.mux.Unlock()
		log.Info().Msgf("Block Response: %v", blockResponse)
		//}
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
	params, ok := r.URL.Query()["height"]
	if !ok || len(params[0]) < 1 {
		log.Info().Msg("Url Param 'height' is missing")
		fmt.Fprint(w, "Request invalid, 'height' param missing\n")
	}

	heightJSON := params[0]

	height, err := strconv.ParseInt(heightJSON, 10, 64)

	if err != nil {
		log.Error().Msgf("Failed to Parse %v", err)
	}

	srv := service{}
	block, _ := srv.GetBlockByHeight(height)
	log.Info().Msgf("Processed for Block Height: %d", block.BlockHeight)
	fmt.Fprint(w, blockResponse)
}
