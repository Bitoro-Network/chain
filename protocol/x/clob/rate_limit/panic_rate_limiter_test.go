package rate_limit_test

import (
	"testing"

	testapp "github.com/Bitoro-Network/chain/protocol/testutil/app"
	"github.com/Bitoro-Network/chain/protocol/x/clob/rate_limit"
	"github.com/stretchr/testify/require"
)

func TestPanicRateLimiter(t *testing.T) {
	tApp := testapp.NewTestAppBuilder(t).Build()
	ctx := tApp.InitChain()
	rl := rate_limit.NewPanicRateLimiter[int]()
	require.Panics(t, func() {
		//nolint:errcheck
		rl.RateLimit(ctx, 5)
	})
	require.Panics(t, func() {
		//nolint:errcheck
		rl.PruneRateLimits(ctx)
	})
}
