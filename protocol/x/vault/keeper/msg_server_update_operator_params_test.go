package keeper_test

import (
	"testing"

	"github.com/bitoro-network/chain/protocol/lib"

	testapp "github.com/bitoro-network/chain/protocol/testutil/app"
	"github.com/bitoro-network/chain/protocol/testutil/constants"

	"github.com/bitoro-network/chain/protocol/x/vault/keeper"
	"github.com/bitoro-network/chain/protocol/x/vault/types"
	"github.com/stretchr/testify/require"
)

func TestMsgUpdateOperatorParams(t *testing.T) {
	tests := map[string]struct {
		// Msg.
		msg *types.MsgUpdateOperatorParams
		// Expected error
		expectedErr string
	}{
		"Success. Update to gov module account": {
			msg: &types.MsgUpdateOperatorParams{
				Authority: lib.GovModuleAddress.String(),
				Params: types.OperatorParams{
					Operator: constants.GovAuthority,
				},
			},
		},
		"Success. Update to Alice": {
			msg: &types.MsgUpdateOperatorParams{
				Authority: lib.GovModuleAddress.String(),
				Params: types.OperatorParams{
					Operator: constants.AliceAccAddress.String(),
				},
			},
		},
		"Failure - Invalid Authority": {
			msg: &types.MsgUpdateOperatorParams{
				Authority: constants.AliceAccAddress.String(),
				Params: types.OperatorParams{
					Operator: constants.GovAuthority,
				},
			},
			expectedErr: "invalid authority",
		},
		"Failure - Invalid Params": {
			msg: &types.MsgUpdateOperatorParams{
				Authority: lib.GovModuleAddress.String(),
				Params: types.OperatorParams{
					Operator: "",
				},
			},
			expectedErr: types.ErrEmptyOperator.Error(),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tApp := testapp.NewTestAppBuilder(t).Build()
			ctx := tApp.InitChain()
			k := tApp.App.VaultKeeper
			ms := keeper.NewMsgServerImpl(k)

			_, err := ms.UpdateOperatorParams(ctx, tc.msg)
			if tc.expectedErr != "" {
				require.ErrorContains(t, err, tc.expectedErr)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.msg.Params, k.GetOperatorParams(ctx))
			}
		})
	}
}
