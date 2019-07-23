package network

import (
	"os/user"
	"strconv"
	"sync"

	"github.com/herdius/herdius-core/blockchain/protobuf"
	"github.com/herdius/herdius-core/p2p/crypto"
	keystore "github.com/herdius/herdius-core/p2p/key"
	"github.com/herdius/herdius-core/p2p/log"
	"github.com/herdius/herdius-core/p2p/network"
	"github.com/herdius/herdius-core/p2p/types/opcode"
	"github.com/herdius/herdius-core/types"

	"github.com/herdius/herdius-blockchain-api/config"
	apiProtobuf "github.com/herdius/herdius-blockchain-api/protobuf"
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

	opcode.RegisterMessageType(types.OpcodeBlockHeightRequest, &apiProtobuf.BlockHeightRequest{})
	opcode.RegisterMessageType(types.OpcodeBlockResponse, &apiProtobuf.BlockResponse{})
	opcode.RegisterMessageType(types.OpcodeAccountRequest, &apiProtobuf.AccountRequest{})
	opcode.RegisterMessageType(types.OpcodeAccountResponse, &apiProtobuf.AccountResponse{})
	opcode.RegisterMessageType(types.OpcodeTxRequest, &apiProtobuf.TxRequest{})
	opcode.RegisterMessageType(types.OpcodeTxResponse, &apiProtobuf.TxResponse{})
	opcode.RegisterMessageType(types.OpcodeTxDetailRequest, &apiProtobuf.TxDetailRequest{})
	opcode.RegisterMessageType(types.OpcodeTxDetailResponse, &apiProtobuf.TxDetailResponse{})
	opcode.RegisterMessageType(types.OpcodeTxsByAddressRequest, &apiProtobuf.TxsByAddressRequest{})
	opcode.RegisterMessageType(types.OpcodeTxsResponse, &apiProtobuf.TxsResponse{})
	opcode.RegisterMessageType(types.OpcodeTxsByAssetAndAddressRequest, &apiProtobuf.TxsByAssetAndAddressRequest{})
	opcode.RegisterMessageType(types.OpcodeTxUpdateRequest, &apiProtobuf.TxUpdateRequest{})
	opcode.RegisterMessageType(types.OpcodeTxUpdateResponse, &apiProtobuf.TxUpdateResponse{})
	opcode.RegisterMessageType(types.OpcodeTxDeleteRequest, &apiProtobuf.TxDeleteRequest{})
	opcode.RegisterMessageType(types.OpcodeTxLockedRequest, &apiProtobuf.TxLockedRequest{})
	opcode.RegisterMessageType(types.OpcodeTxLockedResponse, &apiProtobuf.TxLockedResponse{})
	opcode.RegisterMessageType(types.OpcodePing, &protobuf.Ping{})
	opcode.RegisterMessageType(types.OpcodePong, &protobuf.Pong{})
	opcode.RegisterMessageType(types.OpcodeTxRedeemRequest, &apiProtobuf.TxRedeemRequest{})
	opcode.RegisterMessageType(types.OpcodeTxRedeemResponse, &apiProtobuf.TxRedeemResponse{})
	opcode.RegisterMessageType(types.OpcodeTxsByBlockHeightRequest, &apiProtobuf.TxsByBlockHeightRequest{})
	opcode.RegisterMessageType(types.OpcodeLastBlockRequest, &apiProtobuf.LastBlockRequest{})

	builder := network.NewBuilder(env)
	builder.SetKeys(keys)
	builder.SetAddress(network.FormatAddress(configuration.TCP, configuration.SelfIP, uint16(configuration.ConnectionPort)))

	return builder

}
