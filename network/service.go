package network

import (
	"os/user"
	"strconv"
	"sync"

	"github.com/herdius/herdius-blockchain-api/config"
	apiProtobuf "github.com/herdius/herdius-blockchain-api/protobuf"
	"github.com/herdius/herdius-core/p2p/crypto"
	keystore "github.com/herdius/herdius-core/p2p/key"

	"github.com/herdius/herdius-core/p2p/log"
	"github.com/herdius/herdius-core/p2p/network"
	"github.com/herdius/herdius-core/p2p/types/opcode"
)

var builder *network.Builder
var once sync.Once

// GetNetworkBuilder will instantiate network builder only once
func GetNetworkBuilder(env string) *network.Builder {
	once.Do(func() {
		builder = networkBuilder(env)
	})
	return builder
}

func networkBuilder(env string) *network.Builder {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	configuration := config.GetConfiguration(env)
	nodeAddress := configuration.SelfIP + ":" + strconv.Itoa(configuration.ConnectionPort)

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

	opcode.RegisterMessageType(opcode.Opcode(1113), &apiProtobuf.BlockHeightRequest{})
	opcode.RegisterMessageType(opcode.Opcode(1114), &apiProtobuf.BlockResponse{})
	opcode.RegisterMessageType(opcode.Opcode(1115), &apiProtobuf.AccountRequest{})
	opcode.RegisterMessageType(opcode.Opcode(1116), &apiProtobuf.AccountResponse{})
	opcode.RegisterMessageType(opcode.Opcode(1117), &apiProtobuf.TxRequest{})
	opcode.RegisterMessageType(opcode.Opcode(1118), &apiProtobuf.TxResponse{})
	opcode.RegisterMessageType(opcode.Opcode(1119), &apiProtobuf.TxDetailRequest{})
	opcode.RegisterMessageType(opcode.Opcode(1120), &apiProtobuf.TxDetailResponse{})
	opcode.RegisterMessageType(opcode.Opcode(1121), &apiProtobuf.TxsByAddressRequest{})
	opcode.RegisterMessageType(opcode.Opcode(1122), &apiProtobuf.TxsResponse{})
	opcode.RegisterMessageType(opcode.Opcode(1123), &apiProtobuf.TxsByAssetAndAddressRequest{})
	opcode.RegisterMessageType(opcode.Opcode(1124), &apiProtobuf.TxCancelRequest{})
	opcode.RegisterMessageType(opcode.Opcode(1125), &apiProtobuf.TxCancelResponse{})

	builder := network.NewBuilder(env)
	builder.SetKeys(keys)
	builder.SetAddress(network.FormatAddress(configuration.TCP, configuration.SelfIP, uint16(configuration.ConnectionPort)))

	return builder

}
