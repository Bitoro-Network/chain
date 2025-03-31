package keeper

import (
	"github.com/Bitoro-Network/chain/protocol/x/perpetuals/types"
)

var _ types.QueryServer = Keeper{}
