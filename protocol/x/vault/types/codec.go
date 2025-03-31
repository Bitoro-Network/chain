package types

import (
	"github.com/bitoro-network/chain/protocol/app/module"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	"github.com/cosmos/gogoproto/proto"
)

func RegisterCodec(cdc *codec.LegacyAmino) {}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
	// Register deprecated MsgUpdateParams as it's not part of msg service.
	registry.RegisterInterface(
		"/"+proto.MessageName(&MsgUpdateParams{}),
		(*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	// Register deprecated MsgSetVaultQuotingParams as it's not part of msg service.
	registry.RegisterInterface(
		"/"+proto.MessageName(&MsgSetVaultQuotingParams{}),
		(*sdk.Msg)(nil),
		&MsgSetVaultQuotingParams{},
	)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(module.InterfaceRegistry)
)
