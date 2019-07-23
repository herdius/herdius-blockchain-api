package store

import (
	"context"
	"fmt"

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

	ctx := network.WithSignMessage(context.Background(), true)
	req := protobuf.TxsByBlockHeightRequest{BlockHeight: int64(blockID)}
	res, err := supervisorNode.Request(ctx, &req)
	if err != nil {
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
	}
	return nil
}
