package main

import (
	"os"

	"github.com/bitoro-network/chain/protocol/app"
	"github.com/bitoro-network/chain/protocol/app/config"
	"github.com/bitoro-network/chain/protocol/app/constants"
	"github.com/bitoro-network/chain/protocol/cmd/bitoroprotocold/cmd"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	config.SetupConfig()

	option := cmd.GetOptionWithCustomStartCmd()
	rootCmd := cmd.NewRootCmd(option, app.DefaultNodeHome)

	cmd.AddTendermintSubcommands(rootCmd)
	cmd.AddInitCmdPostRunE(rootCmd)

	if err := svrcmd.Execute(rootCmd, constants.AppDaemonName, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
