package types

import (
	"math/big"

	stattypes "github.com/Bitoro-Network/chain/protocol/x/stats/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type StatsKeeper interface {
	GetStakedAmount(ctx sdk.Context, delegatorAddr string) *big.Int
	GetBlockStats(ctx sdk.Context) *stattypes.BlockStats
}

type FeetiersKeeper interface {
	GetAffiliateRefereeLowestTakerFee(ctx sdk.Context) int32
	GetLowestMakerFee(ctx sdk.Context) int32
}
