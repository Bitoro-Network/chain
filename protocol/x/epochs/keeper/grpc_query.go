package keeper

import (
	"github.com/bitoro-network/chain/protocol/x/epochs/types"
)

var _ types.QueryServer = Keeper{}
