package grpc

import pricetypes "github.com/Bitoro-Network/chain/protocol/x/prices/types"

type QueryServer interface {
	pricetypes.QueryServer
}
