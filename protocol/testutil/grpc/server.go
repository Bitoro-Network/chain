package grpc

import pricetypes "github.com/bitoro-network/chain/protocol/x/prices/types"

type QueryServer interface {
	pricetypes.QueryServer
}
