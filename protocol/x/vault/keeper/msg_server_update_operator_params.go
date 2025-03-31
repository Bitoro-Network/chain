package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"

	"github.com/Bitoro-Network/chain/protocol/lib"
	"github.com/Bitoro-Network/chain/protocol/x/vault/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// UpdateOperatorParams updates the operator parameters of megavault.
func (k msgServer) UpdateOperatorParams(
	goCtx context.Context,
	msg *types.MsgUpdateOperatorParams,
) (*types.MsgUpdateOperatorParamsResponse, error) {
	if !k.HasAuthority(msg.Authority) {
		return nil, errorsmod.Wrapf(
			govtypes.ErrInvalidSigner,
			"invalid authority %s",
			msg.Authority,
		)
	}

	ctx := lib.UnwrapSDKContext(goCtx, types.ModuleName)
	if err := k.SetOperatorParams(ctx, msg.Params); err != nil {
		return nil, err
	}

	return &types.MsgUpdateOperatorParamsResponse{}, nil
}
