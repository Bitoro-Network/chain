package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"

	"github.com/bitoro-network/chain/protocol/lib"
	"github.com/bitoro-network/chain/protocol/x/blocktime/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) UpdateDowntimeParams(
	goCtx context.Context,
	msg *types.MsgUpdateDowntimeParams,
) (*types.MsgUpdateDowntimeParamsResponse, error) {
	if !k.HasAuthority(msg.Authority) {
		return nil, errorsmod.Wrapf(
			govtypes.ErrInvalidSigner,
			"invalid authority %s",
			msg.Authority,
		)
	}

	ctx := lib.UnwrapSDKContext(goCtx, types.ModuleName)
	if err := k.SetDowntimeParams(ctx, msg.Params); err != nil {
		return nil, err
	}

	return &types.MsgUpdateDowntimeParamsResponse{}, nil
}

func (k msgServer) UpdateSynchronyParams(
	goCtx context.Context,
	msg *types.MsgUpdateSynchronyParams,
) (*types.MsgUpdateSynchronyParamsResponse, error) {
	if !k.HasAuthority(msg.Authority) {
		return nil, errorsmod.Wrapf(
			govtypes.ErrInvalidSigner,
			"invalid authority %s",
			msg.Authority,
		)
	}

	ctx := lib.UnwrapSDKContext(goCtx, types.ModuleName)
	k.SetSynchronyParams(ctx, msg.Params)

	return &types.MsgUpdateSynchronyParamsResponse{}, nil
}
