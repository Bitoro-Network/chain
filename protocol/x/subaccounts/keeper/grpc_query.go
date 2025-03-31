package keeper

import (
	"github.com/Bitoro-Network/chain/protocol/x/subaccounts/types"
)

var _ types.QueryServer = Keeper{}
