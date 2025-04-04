package ratelimit

import (
	"github.com/bitoro-network/chain/protocol/x/ratelimit/keeper"
	"github.com/bitoro-network/chain/protocol/x/ratelimit/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	for _, limitParams := range genState.LimitParamsList {
		if err := k.SetLimitParams(ctx, limitParams); err != nil {
			panic(err)
		}
	}
	k.InitializeForGenesis(ctx)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		LimitParamsList: k.GetAllLimitParams(ctx),
	}
}
