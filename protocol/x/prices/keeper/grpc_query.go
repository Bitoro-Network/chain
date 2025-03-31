package keeper

import (
	"github.com/bitoro-network/chain/protocol/x/prices/types"
)

var _ types.QueryServer = Keeper{}
