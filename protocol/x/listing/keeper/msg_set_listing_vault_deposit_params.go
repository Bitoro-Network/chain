package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/bitoro-network/chain/protocol/lib"
	"github.com/bitoro-network/chain/protocol/x/listing/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func (k msgServer) SetListingVaultDepositParams(
	goCtx context.Context,
	msg *types.MsgSetListingVaultDepositParams,
) (*types.MsgSetListingVaultDepositParamsResponse, error) {
	ctx := lib.UnwrapSDKContext(goCtx, types.ModuleName)

	// Check if the sender has the authority to set the vault deposit params
	if !k.HasAuthority(msg.Authority) {
		return nil, errorsmod.Wrapf(
			govtypes.ErrInvalidSigner,
			"invalid authority %s",
			msg.Authority,
		)
	}

	// Set the vault deposit params for the listing
	err := k.Keeper.SetListingVaultDepositParams(ctx, msg.Params)
	if err != nil {
		return nil, err
	}

	return &types.MsgSetListingVaultDepositParamsResponse{}, nil
}
