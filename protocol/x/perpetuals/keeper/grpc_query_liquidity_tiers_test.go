package keeper_test

import (
	"context"
	"testing"

	testapp "github.com/bitoro-network/chain/protocol/testutil/app"
	"github.com/bitoro-network/chain/protocol/testutil/constants"
	perptypes "github.com/bitoro-network/chain/protocol/x/perpetuals/types"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
)

func TestAllLiquidityTiers(
	t *testing.T,
) {
	tApp := testapp.NewTestAppBuilder(t).
		WithGenesisDocFn(func() types.GenesisDoc {
			genesis := testapp.DefaultGenesis()
			testapp.UpdateGenesisDocWithAppStateForModule(&genesis, func(state *perptypes.GenesisState) {
				state.LiquidityTiers = constants.LiquidityTiers
			})
			return genesis
		}).Build()

	tApp.InitChain()

	request := perptypes.QueryAllLiquidityTiersRequest{}
	abciResponse, err := tApp.App.Query(
		context.Background(),
		&abci.RequestQuery{
			Path: "/bitoroprotocol.perpetuals.Query/AllLiquidityTiers",
			Data: tApp.App.AppCodec().MustMarshal(&request),
		})
	require.NoError(t, err)
	require.True(t, abciResponse.IsOK())

	var actual perptypes.QueryAllLiquidityTiersResponse
	tApp.App.AppCodec().MustUnmarshal(abciResponse.Value, &actual)

	expected := perptypes.QueryAllLiquidityTiersResponse{
		LiquidityTiers: constants.LiquidityTiers,
		Pagination: &query.PageResponse{
			NextKey: nil,
			Total:   uint64(len(constants.LiquidityTiers)),
		},
	}
	require.Equal(t, expected, actual)
}
