package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Bitoro-Network/chain/protocol/x/revshare/types"
	"github.com/cosmos/cosmos-sdk/client"
)

// GetQueryCmd returns the cli query commands for this module.
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group x/revshare queries under a subcommand.
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryRevShareParams())
	cmd.AddCommand(CmdQueryRevShareDetailsForMarket())
	cmd.AddCommand(CmdQueryUnconditionalRevShareConfig())

	return cmd
}
