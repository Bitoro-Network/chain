package cmd_test

import (
	"fmt"
	"testing"

	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/bitoro-network/chain/protocol/app/config"
	"github.com/bitoro-network/chain/protocol/app/constants"
	"github.com/bitoro-network/chain/protocol/cmd/bitoroprotocold/cmd"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/stretchr/testify/require"
)

func TestNewRootCmd_UsesClientConfig(t *testing.T) {
	tempDir := t.TempDir()

	config.SetupConfig()

	// Set the client config to point to a fake chain id since this is a required option
	{
		option := cmd.GetOptionWithCustomStartCmd()
		rootCmd := cmd.NewRootCmd(option, tempDir)

		cmd.AddTendermintSubcommands(rootCmd)
		cmd.AddInitCmdPostRunE(rootCmd)
		rootCmd.SetArgs([]string{"config", "set", "client", "chain-id", "fakeChainId"})
		require.NoError(t, svrcmd.Execute(rootCmd, constants.AppDaemonName, tempDir))
	}

	// Set the client config to point to a fake address
	{
		option := cmd.GetOptionWithCustomStartCmd()
		rootCmd := cmd.NewRootCmd(option, tempDir)

		cmd.AddTendermintSubcommands(rootCmd)
		cmd.AddInitCmdPostRunE(rootCmd)
		rootCmd.SetArgs([]string{"config", "set", "client", "node", "fakeTestAddress"})
		require.NoError(t, svrcmd.Execute(rootCmd, constants.AppDaemonName, tempDir))
	}

	// Run a query command (that will fail) to ensure that we are reading the client config
	option := cmd.GetOptionWithCustomStartCmd()
	rootCmd := cmd.NewRootCmd(option, tempDir)

	cmd.AddTendermintSubcommands(rootCmd)
	cmd.AddInitCmdPostRunE(rootCmd)
	rootCmd.SetArgs([]string{"query", "auth", "params"})
	require.ErrorContains(t, svrcmd.Execute(rootCmd, constants.AppDaemonName, tempDir), "fakeTestAddress")
}

func TestCmdModuleNameToAddress(t *testing.T) {
	expectedModuleNameAddress := map[string]string{
		"subaccounts":       "btoro1v88c3xv9xyv3eetdx0tvcmq7ung3dywp5upwc6",
		"subaccounts:37":    "btoro16lwrx54mh9aru9ulzpknd429wldkhdwekhlswf",
		"insurance_fund":    "btoro1c7ptc87hkd54e3r7zjy92q29xkq7t79w64slrq",
		"insurance_fund:37": "btoro10mlrxmaquwjwsj59ywp8xttc8rfxn9jfvzswtn",
	}
	for moduleName, expectedAddress := range expectedModuleNameAddress {
		t.Run(
			fmt.Sprintf("ModuleNameToAddress %s", moduleName), func(t *testing.T) {
				ctx := client.Context{}
				out, err := clitestutil.ExecTestCLICmd(ctx, cmd.CmdModuleNameToAddress(), []string{moduleName})
				require.NoError(t, err)
				require.Equal(t, expectedAddress, out.String())
			},
		)
	}
}
