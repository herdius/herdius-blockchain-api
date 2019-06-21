package store

import (
	"context"

	"github.com/herdius/herdius-blockchain-api/config"
	"github.com/herdius/herdius-blockchain-api/protobuf"
	"github.com/herdius/herdius-core/p2p/network"
)

// SyncPendingTxs syncs pending status with core
func SyncPendingTxs(s Storer, net *network.Network, env string) error {
	configuration := config.GetConfiguration(env)
	supervisorAddress := configuration.GetSupervisorAddress()
	supervisorNode, _ := net.Client(supervisorAddress)
	txs, err := s.GetByStatus(StatusPending)
	if err != nil {
		return err
	}

	for _, tx := range txs {
		txDetailReq := protobuf.TxDetailRequest{TxId: tx.ID}
		ctx := network.WithSignMessage(context.Background(), true)
		res, err := supervisorNode.Request(ctx, &txDetailReq)
		if err != nil {
			return err
		}
		if msg, ok := res.(*protobuf.TxDetailResponse); ok {
			if msg.Tx != nil {
				if err := s.Update(FromTxDetailResponse(msg)); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
