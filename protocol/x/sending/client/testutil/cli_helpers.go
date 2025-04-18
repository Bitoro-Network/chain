package testutil

import (
	"fmt"

	sendingcli "github.com/bitoro-network/chain/protocol/x/sending/client/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/testutil"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
)

// MsgCreateTransferExec broadcasts a transfer message.
func MsgCreateTransferExec(
	clientCtx client.Context,
	senderOwner sdk.AccAddress,
	senderNumber uint32,
	recipientOwner sdk.AccAddress,
	recipientNumber uint32,
	amount uint64,
) (testutil.BufferWriter, error) {
	args := []string{
		senderOwner.String(),
		fmt.Sprint(senderNumber),
		recipientOwner.String(),
		fmt.Sprint(recipientNumber),
		fmt.Sprint(amount),
	}

	args = append(args,
		fmt.Sprintf("--%s=%s", flags.FlagFrom, "node0"),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
	)

	return clitestutil.ExecTestCLICmd(clientCtx, sendingcli.CmdCreateTransfer(), args)
}
