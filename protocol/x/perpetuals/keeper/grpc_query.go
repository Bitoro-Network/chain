package keeper

import (
	"github.com/bitoro-network/chain/protocol/x/perpetuals/types"
)

var _ types.QueryServer = Keeper{}
