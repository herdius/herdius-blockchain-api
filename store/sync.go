package store

import (
	"context"
	"fmt"
	"log"

	"github.com/herdius/herdius-blockchain-api/config"
	"github.com/herdius/herdius-blockchain-api/protobuf"
	"github.com/herdius/herdius-core/p2p/network"
)

// SyncPendingTxs syncs pending status with core
func SyncPendingTxs(s Storer, net *network.Network, env string) error {
	configuration := config.GetConfiguration(env)
	supervisorAddress := configuration.GetSupervisorAddress()
	supervisorNode, _ := net.Client(supervisorAddress)
	blockID, err := s.GetLatestBlockID()
	if err != nil {
		return fmt.Errorf("s.GetLatestBlockID: %v", err)
	}

	for {
		blockID++
		ctx := network.WithSignMessage(context.Background(), true)
		req := protobuf.BlockHeightRequest{BlockHeight: blockID}
		msg, err := supervisorNode.Request(ctx, &req)
		if err != nil {
			log.Printf("failed to query block info: %d\n", blockID)
			break
		}

		block, ok := msg.(*protobuf.BlockResponse)
		if !ok || block.Time == nil {
			log.Printf("block %d does not exist\n", blockID)
			break
		}

		ctx = network.WithSignMessage(context.Background(), true)
		txsReq := protobuf.TxsByBlockHeightRequest{BlockHeight: blockID}
		res, err := supervisorNode.Request(ctx, &txsReq)
		if err != nil {
			log.Printf("failed to get txs in block: %d\n", blockID)
			return fmt.Errorf("supervisorNode.Request: %v", err)
		}

		if msg, ok := res.(*protobuf.TxsResponse); ok {
			for _, tx := range msg.GetTxs() {
				if tx.Tx == nil {
					continue
				}
				if err := s.Update(FromTxDetailResponse(tx)); err != nil {
					return err
				}
			}
			return nil
		}
	}
	return nil
}
