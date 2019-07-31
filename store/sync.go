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

	lastSyncBlockID := s.GetLastSyncBlockID()
	if blockID < lastSyncBlockID {
		blockID = lastSyncBlockID
	}

	ctx := network.WithSignMessage(context.Background(), true)
	req := protobuf.LastBlockRequest{}
	msg, err := supervisorNode.Request(ctx, &req)
	if err != nil {
		log.Printf("failed to query last block info: %d\n", blockID)
		return err
	}

	block, ok := msg.(*protobuf.BlockResponse)
	if !ok {
		log.Printf("block %d does not exist\n", blockID)
		return nil
	}

	for blockID < block.GetBlockHeight() {
		blockID++
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
					log.Printf("failed to save tx to database: %v - %v\n", tx, err)
				}
			}
		}
	}

	if err := s.SaveLastSyncBlockID(block.GetBlockHeight()); err != nil {
		log.Printf("failed to save last sync block id: %v", err)
	}

	return nil
}
