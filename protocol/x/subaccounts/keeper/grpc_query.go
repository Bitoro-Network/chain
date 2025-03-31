package keeper

import (
	"github.com/bitoro-network/chain/protocol/x/subaccounts/types"
)

var _ types.QueryServer = Keeper{}
