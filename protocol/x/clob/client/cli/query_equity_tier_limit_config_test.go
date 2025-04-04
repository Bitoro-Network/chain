//go:build all || integration_test

package cli_test

import (
	"fmt"
	"testing"

	"github.com/bitoro-network/chain/protocol/x/clob/client/cli"
	"github.com/bitoro-network/chain/protocol/x/clob/types"
	tmcli "github.com/cometbft/cometbft/libs/cli"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
)

var (
	emptyEquityTierLimitConfig = types.EquityTierLimitConfiguration{
		ShortTermOrderEquityTiers: []types.EquityTierLimit{},
		StatefulOrderEquityTiers:  []types.EquityTierLimit{},
	}
)

func TestCmdGetEquityTierLimitConfig(t *testing.T) {
	net, _ := networkWithClobPairObjects(t, 2)
	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}

	out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdGetEquityTierLimitConfig(), common)
	require.NoError(t, err)
	var resp types.QueryEquityTierLimitConfigurationResponse
	require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
	require.NotNil(t, resp.EquityTierLimitConfig)
	require.Equal(
		t,
		emptyEquityTierLimitConfig,
		resp.EquityTierLimitConfig,
	)
}
