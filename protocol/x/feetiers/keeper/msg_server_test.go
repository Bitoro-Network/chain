package keeper_test

import (
	"context"
	"testing"

	"github.com/bitoro-network/chain/protocol/lib"

	testapp "github.com/bitoro-network/chain/protocol/testutil/app"
	"github.com/bitoro-network/chain/protocol/x/feetiers/keeper"
	"github.com/bitoro-network/chain/protocol/x/feetiers/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t *testing.T) (keeper.Keeper, types.MsgServer, context.Context) {
	tApp := testapp.NewTestAppBuilder(t).Build()
	ctx := tApp.InitChain()
	k := tApp.App.FeeTiersKeeper

	return *k, keeper.NewMsgServerImpl(k), ctx
}

func TestMsgServer(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	require.NotNil(t, k)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
}

func TestMsgUpdateParams(t *testing.T) {
	_, ms, ctx := setupMsgServer(t)

	testCases := []struct {
		name      string
		input     *types.MsgUpdatePerpetualFeeParams
		expErr    bool
		expErrMsg string
	}{
		{
			name: "valid params",
			input: &types.MsgUpdatePerpetualFeeParams{
				Authority: lib.GovModuleAddress.String(),
				Params:    types.DefaultGenesis().Params,
			},
			expErr: false,
		},
		{
			name: "invalid authority",
			input: &types.MsgUpdatePerpetualFeeParams{
				Authority: "invalid",
				Params:    types.DefaultGenesis().Params,
			},
			expErr:    true,
			expErrMsg: "invalid authority",
		},
		{
			name: "invalid params: negative duration",
			input: &types.MsgUpdatePerpetualFeeParams{
				Authority: lib.GovModuleAddress.String(),
				Params: types.PerpetualFeeParams{
					Tiers: []*types.PerpetualFeeTier{
						{TotalVolumeShareRequirementPpm: 1},
					},
				},
			},
			expErr:    true,
			expErrMsg: "First fee tier must not have volume requirements",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := ms.UpdatePerpetualFeeParams(ctx, tc.input)
			if tc.expErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expErrMsg)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
