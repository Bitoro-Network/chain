package prices

import (
	indexerevents "github.com/bitoro-network/chain/protocol/indexer/events"
	"github.com/bitoro-network/chain/protocol/indexer/indexer_manager"
	"github.com/bitoro-network/chain/protocol/x/prices/keeper"
	"github.com/bitoro-network/chain/protocol/x/prices/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the x/prices module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.InitializeForGenesis(ctx)

	if len(genState.MarketPrices) != len(genState.MarketParams) {
		panic("Expected the same number of market prices and market params")
	}

	// Set all the market params and prices.
	for i, param := range genState.MarketParams {
		if _, err := k.CreateMarket(ctx, param, genState.MarketPrices[i]); err != nil {
			panic(err)
		}
	}

	// Generate indexer events.
	priceUpdateIndexerEvents := keeper.GenerateMarketPriceUpdateIndexerEvents(genState.MarketPrices)
	for _, update := range priceUpdateIndexerEvents {
		k.GetIndexerEventManager().AddTxnEvent(
			ctx,
			indexerevents.SubtypeMarket,
			indexerevents.MarketEventVersion,
			indexer_manager.GetBytes(
				update,
			),
		)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.MarketParams = k.GetAllMarketParams(ctx)
	genesis.MarketPrices = k.GetAllMarketPrices(ctx)

	return genesis
}
