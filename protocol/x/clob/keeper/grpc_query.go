package keeper

import (
	"github.com/bitoro-network/chain/protocol/x/clob/types"
)

var _ types.QueryServer = Keeper{}
