package bridge_test

import (
	"testing"

	testapp "github.com/Bitoro-Network/chain/protocol/testutil/app"
	"github.com/Bitoro-Network/chain/protocol/x/bridge"
	"github.com/Bitoro-Network/chain/protocol/x/bridge/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	tApp := testapp.NewTestAppBuilder(t).Build()
	ctx := tApp.InitChain()
	got := bridge.ExportGenesis(ctx, tApp.App.BridgeKeeper)
	require.NotNil(t, got)
	require.Equal(t, types.DefaultGenesis(), got)
}
