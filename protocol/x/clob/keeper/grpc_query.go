package keeper

import (
	"github.com/Bitoro-Network/chain/protocol/x/clob/types"
)

var _ types.QueryServer = Keeper{}
