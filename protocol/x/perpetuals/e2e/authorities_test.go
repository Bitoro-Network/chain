package perpetuals_e2e_test

import (
	"testing"

	testapp "github.com/bitoro-network/chain/protocol/testutil/app"
	delaymsgtypes "github.com/bitoro-network/chain/protocol/x/delaymsg/types"
	"github.com/cometbft/cometbft/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/stretchr/testify/require"
)

func TestHasAuthority(t *testing.T) {
	tests := map[string]struct {
		authorityAddress string
		hasAuthority     bool
	}{
		"gov is an authority": {
			authorityAddress: authtypes.NewModuleAddress(govtypes.ModuleName).String(),
			hasAuthority:     true,
		},
		"delaymsg is an authority": {
			authorityAddress: authtypes.NewModuleAddress(delaymsgtypes.ModuleName).String(),
			hasAuthority:     true,
		},
		"random invalid authority": {
			authorityAddress: "random",
			hasAuthority:     false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tApp := testapp.NewTestAppBuilder(t).WithGenesisDocFn(func() (genesis types.GenesisDoc) {
				return testapp.DefaultGenesis()
			}).Build()
			tApp.InitChain()
			require.Equal(
				t,
				tc.hasAuthority,
				tApp.App.PerpetualsKeeper.HasAuthority(tc.authorityAddress),
			)
		})
	}
}
