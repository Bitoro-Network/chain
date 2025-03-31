package keeper

import (
	"github.com/Bitoro-Network/chain/protocol/x/delaymsg/types"
)

var _ types.QueryServer = Keeper{}
