package keeper

import (
	"github.com/bitoro-network/chain/protocol/x/sending/types"
)

var _ types.QueryServer = Keeper{}
