package handler

import (
	"os/user"
	"strconv"

	apiProtobuf "github.com/herdius/herdius-blockchain-api/protobuf"
	blockProtobuf "github.com/herdius/herdius-core/blockchain/protobuf"
	"github.com/herdius/herdius-core/p2p/crypto"
	keystore "github.com/herdius/herdius-core/p2p/key"

	"github.com/herdius/herdius-core/p2p/log"
	"github.com/herdius/herdius-core/p2p/network"
	"github.com/herdius/herdius-core/p2p/network/discovery"
	"github.com/herdius/herdius-core/p2p/types/opcode"
)

// NetworkBuilder ...
type NetworkBuilder struct {
	builder *network.Builder
}

// CreateNetworkBuilder ...
func CreateNetworkBuilder() *NetworkBuilder {
	return &NetworkBuilder{
		builder: networkBuilder(),
	}
}

// GetNetworkBuilder ...
func (nb NetworkBuilder) GetNetworkBuilder() *network.Builder {
	return nb.builder
}

// NetworkBuilder ...
func networkBuilder() *network.Builder {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	port := 5555
	host := "localhost"

	peer := "tcp://localhost:3000"
	peers := make([]string, 1)
	peers = append(peers, peer)

	nodeAddress := host + ":" + strconv.Itoa(port)

	nodekey, err := keystore.LoadOrGenNodeKey(user.HomeDir + "/" + nodeAddress + "_peer_id.json")

	if err != nil {
		log.Error().Msgf("Failed to create or load node key: %v", err)
	}

	privKey := nodekey.PrivKey

	pubKey := privKey.PubKey()

	keys := &crypto.KeyPair{
		PublicKey:  pubKey.Bytes(),
		PrivateKey: privKey.Bytes(),
		PrivKey:    privKey,
		PubKey:     pubKey,
	}

	opcode.RegisterMessageType(opcode.Opcode(1112), &blockProtobuf.ConnectionMessage{})
	opcode.RegisterMessageType(opcode.Opcode(1113), &apiProtobuf.BlockHeightRequest{})
	opcode.RegisterMessageType(opcode.Opcode(1114), &apiProtobuf.BlockResponse{})
	opcode.RegisterMessageType(opcode.Opcode(1115), &apiProtobuf.AccountRequest{})
	opcode.RegisterMessageType(opcode.Opcode(1116), &apiProtobuf.AccountResponse{})
	opcode.RegisterMessageType(opcode.Opcode(1117), &apiProtobuf.TxRequest{})
	opcode.RegisterMessageType(opcode.Opcode(1118), &apiProtobuf.TxResponse{})

	builder := network.NewBuilder()
	builder.SetKeys(keys)
	builder.SetAddress(network.FormatAddress("tcp", host, uint16(port)))

	// // Register peer discovery plugin.
	builder.AddPlugin(new(discovery.Plugin))
	builder.AddPlugin(new(BlockMessagePlugin))
	builder.AddPlugin(new(AccountMessagePlugin))
	builder.AddPlugin(new(TXMessagePlugin))
	return builder

}
