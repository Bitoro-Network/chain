package types_test

import (
	"testing"

	"cosmossdk.io/math"
	"github.com/bitoro-network/chain/protocol/dtypes"
	"github.com/bitoro-network/chain/protocol/lib"
	types "github.com/bitoro-network/chain/protocol/x/govplus/types"
	"github.com/stretchr/testify/require"
)

var (
	validConsensusAddress = "cosmosvalcons1ntk8eualewuprz0gamh8hnvcem2nrcdsgz563h"
)

func TestValidateBasic(t *testing.T) {
	tests := map[string]struct {
		msg         types.MsgSlashValidator
		expectedErr error
	}{
		"invalid authority": {
			msg: types.MsgSlashValidator{
				Authority: "invalid",
			},
			expectedErr: types.ErrInvalidAuthority,
		},
		"bad address": {
			msg: types.MsgSlashValidator{
				Authority:        lib.GovModuleAddress.String(),
				ValidatorAddress: "asdfasdfasdfaf",
			},
			expectedErr: types.ErrValidatorAddress,
		},
		"bad tokens at infraction height": {
			msg: types.MsgSlashValidator{
				Authority:                lib.GovModuleAddress.String(),
				ValidatorAddress:         validConsensusAddress,
				TokensAtInfractionHeight: dtypes.NewInt(-10),
				SlashFactor:              math.LegacyMustNewDecFromStr("0.5"),
			},
			expectedErr: types.ErrInvalidTokensAtInfractionHeight,
		},
		"bad slash factor": {
			msg: types.MsgSlashValidator{
				Authority:                lib.GovModuleAddress.String(),
				ValidatorAddress:         validConsensusAddress,
				TokensAtInfractionHeight: dtypes.NewInt(100),
				SlashFactor:              math.LegacyMustNewDecFromStr("1.1"),
			},
			expectedErr: types.ErrInvalidSlashFactor,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.expectedErr == nil {
				require.NoError(t, err)
			} else {
				require.ErrorIs(t, err, tc.expectedErr)
			}
		})
	}
}
