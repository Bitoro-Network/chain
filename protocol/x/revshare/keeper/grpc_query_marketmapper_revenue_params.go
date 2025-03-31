package keeper

import (
	"context"

	"github.com/Bitoro-Network/chain/protocol/x/revshare/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) MarketMapperRevenueShareParams(
	ctx context.Context,
	req *types.QueryMarketMapperRevenueShareParams,
) (*types.QueryMarketMapperRevenueShareParamsResponse, error) {
	params := k.GetMarketMapperRevenueShareParams(sdk.UnwrapSDKContext(ctx))
	return &types.QueryMarketMapperRevenueShareParamsResponse{
		Params: params,
	}, nil
}
