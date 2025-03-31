package keeper

import (
	"github.com/Bitoro-Network/chain/protocol/x/epochs/types"
)

var _ types.QueryServer = Keeper{}
