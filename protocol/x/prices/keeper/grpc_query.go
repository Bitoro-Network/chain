package keeper

import (
	"github.com/Bitoro-Network/chain/protocol/x/prices/types"
)

var _ types.QueryServer = Keeper{}
