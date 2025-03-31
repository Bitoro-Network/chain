package keeper

import (
	"context"

	"github.com/bitoro-network/chain/protocol/x/listing/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) ListingVaultDepositParams(
	ctx context.Context,
	req *types.QueryListingVaultDepositParams,
) (*types.QueryListingVaultDepositParamsResponse, error) {
	params := k.GetListingVaultDepositParams(sdk.UnwrapSDKContext(ctx))
	return &types.QueryListingVaultDepositParamsResponse{
		Params: params,
	}, nil
}
