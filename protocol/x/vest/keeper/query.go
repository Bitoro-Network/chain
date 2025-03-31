package keeper

import (
	"github.com/bitoro-network/chain/protocol/x/vest/types"
)

var _ types.QueryServer = Keeper{}
