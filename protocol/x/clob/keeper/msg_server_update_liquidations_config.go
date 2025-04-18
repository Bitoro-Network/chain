package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"

	"github.com/bitoro-network/chain/protocol/lib"
	"github.com/bitoro-network/chain/protocol/x/clob/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// UpdateLiquidationsConfig updates the liquidation config in state.
func (k msgServer) UpdateLiquidationsConfig(
	goCtx context.Context,
	msg *types.MsgUpdateLiquidationsConfig,
) (resp *types.MsgUpdateLiquidationsConfigResponse, err error) {
	ctx := lib.UnwrapSDKContext(goCtx, types.ModuleName)

	if !k.Keeper.HasAuthority(msg.Authority) {
		return nil, errorsmod.Wrapf(
			govtypes.ErrInvalidSigner,
			"invalid authority %s",
			msg.Authority,
		)
	}

	if err := k.Keeper.UpdateLiquidationsConfig(ctx, msg.LiquidationsConfig); err != nil {
		return nil, err
	}
	return &types.MsgUpdateLiquidationsConfigResponse{}, nil
}
