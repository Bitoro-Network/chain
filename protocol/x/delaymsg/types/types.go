package types

import (
	"cosmossdk.io/log"
	"github.com/bitoro-network/chain/protocol/lib"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type DelayMsgKeeper interface {
	// Delayed messages
	SetNextDelayedMessageId(
		ctx sdk.Context,
		nextDelayedMessageId uint32,
	)

	SetDelayedMessage(
		ctx sdk.Context,
		msg *DelayedMessage,
	) (
		err error,
	)
	DelayMessageByBlocks(
		ctx sdk.Context,
		msg sdk.Msg,
		blockDelay uint32,
	) (
		id uint32,
		err error,
	)

	GetMessage(
		ctx sdk.Context,
		id uint32,
	) (
		delayedMessage DelayedMessage,
		found bool,
	)

	GetAllDelayedMessages(ctx sdk.Context) []*DelayedMessage

	DeleteMessage(
		ctx sdk.Context,
		id uint32,
	) (
		err error,
	)

	// Block message ids
	GetBlockMessageIds(
		ctx sdk.Context,
		blockHeight uint32,
	) (
		blockMessageIds BlockMessageIds,
		found bool,
	)

	// HasAuthority returns whether the authority is permitted to send delayed messages.
	HasAuthority(authority string) bool

	Logger(ctx sdk.Context) log.Logger

	Router() lib.MsgRouter
}
