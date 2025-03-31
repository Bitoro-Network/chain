package main

import (
	"os"

	"github.com/Bitoro-Network/chain/protocol/app"
	"github.com/Bitoro-Network/chain/protocol/app/config"
	"github.com/Bitoro-Network/chain/protocol/app/constants"
	"github.com/Bitoro-Network/chain/protocol/cmd/bitoroprotocold/cmd"
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
