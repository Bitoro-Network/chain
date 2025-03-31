package rate_limit_test

import (
	"testing"

	testapp "github.com/Bitoro-Network/chain/protocol/testutil/app"
	"github.com/Bitoro-Network/chain/protocol/x/clob/rate_limit"
	"github.com/stretchr/testify/require"
)

func TestNoOpRateLimiter(t *testing.T) {
	tApp := testapp.NewTestAppBuilder(t).Build()
	ctx := tApp.InitChain()
	rl := rate_limit.NewNoOpRateLimiter[int]()
	for i := 0; i < 100_000; i += 1 {
		require.NoError(t, rl.RateLimit(ctx, 42))
	}
	rl.PruneRateLimits(ctx)
}
