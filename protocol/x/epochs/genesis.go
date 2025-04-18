package epochs

import (
	"github.com/bitoro-network/chain/protocol/x/epochs/keeper"
	"github.com/bitoro-network/chain/protocol/x/epochs/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the epochs module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.InitializeForGenesis(ctx)

	// Set all the epochInfo
	for _, elem := range genState.EpochInfoList {
		if err := k.CreateEpochInfo(ctx, elem); err != nil {
			panic(err)
		}
	}
}

// ExportGenesis returns the epochs module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		EpochInfoList: k.GetAllEpochInfo(ctx),
	}
}
