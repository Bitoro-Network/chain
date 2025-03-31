package affiliates

import (
	"github.com/bitoro-network/chain/protocol/lib/log"
	"github.com/bitoro-network/chain/protocol/x/affiliates/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func EndBlocker(
	ctx sdk.Context,
	keeper *keeper.Keeper,
) {
	if err := keeper.AggregateAffiliateReferredVolumeForFills(ctx); err != nil {
		log.ErrorLogWithError(ctx, "error aggregating affiliate volume for fills", err)
	}
}
