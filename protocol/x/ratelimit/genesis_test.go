package ratelimit_test

import (
	"testing"

	testapp "github.com/Bitoro-Network/chain/protocol/testutil/app"
	"github.com/Bitoro-Network/chain/protocol/x/ratelimit"
	"github.com/Bitoro-Network/chain/protocol/x/ratelimit/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	tApp := testapp.NewTestAppBuilder(t).Build()
	ctx := tApp.InitChain()
	got := ratelimit.ExportGenesis(ctx, tApp.App.RatelimitKeeper)
	require.NotNil(t, got)
	require.Equal(t, types.DefaultGenesis(), got)
}
