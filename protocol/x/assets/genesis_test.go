package assets_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	keepertest "github.com/Bitoro-Network/chain/protocol/testutil/keeper"
	"github.com/Bitoro-Network/chain/protocol/x/assets"
	"github.com/Bitoro-Network/chain/protocol/x/assets/keeper"
	"github.com/Bitoro-Network/chain/protocol/x/assets/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	expected := types.DefaultGenesis()
	ctx, k, _, _, _, _ := keepertest.AssetsKeepers(t, true)
	assets.InitGenesis(ctx, *k, *expected)
	assertAssetCreateEventsInIndexerBlock(t, k, ctx, len(expected.Assets))
	actual := assets.ExportGenesis(ctx, *k)
	require.NotNil(t, actual)
	require.ElementsMatch(t, actual.Assets, expected.Assets)
}

// assertAssetCreateEventsInIndexerBlock checks that the number of asset create events
// included in the Indexer block kafka message.
func assertAssetCreateEventsInIndexerBlock(
	t *testing.T,
	k *keeper.Keeper,
	ctx sdk.Context,
	numAssets int,
) {
	assetEvents := keepertest.GetAssetCreateEventsFromIndexerBlock(ctx, k)
	require.Len(t, assetEvents, numAssets)
}
