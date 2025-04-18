package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/bitoro-network/chain/protocol/lib"
	"github.com/bitoro-network/chain/protocol/x/clob/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// UpdateBlockRateLimitConfiguration updates the equity tier limit configuration returning an error
// if the configuration is invalid.
func (k msgServer) UpdateEquityTierLimitConfiguration(
	goCtx context.Context,
	msg *types.MsgUpdateEquityTierLimitConfiguration,
) (resp *types.MsgUpdateEquityTierLimitConfigurationResponse, err error) {
	ctx := lib.UnwrapSDKContext(goCtx, types.ModuleName)

	if !k.Keeper.HasAuthority(msg.Authority) {
		return nil, errorsmod.Wrapf(
			govtypes.ErrInvalidSigner,
			"invalid authority %s",
			msg.Authority,
		)
	}

	if err := k.Keeper.InitializeEquityTierLimit(ctx, msg.EquityTierLimitConfig); err != nil {
		return nil, err
	}
	return &types.MsgUpdateEquityTierLimitConfigurationResponse{}, nil
}
