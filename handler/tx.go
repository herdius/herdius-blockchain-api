package handler

import (
	"github.com/herdius/herdius-core/p2p/network"
)

// TXMessagePlugin will receive all Transaction specific messages.
type TXMessagePlugin struct{ *network.Plugin }
