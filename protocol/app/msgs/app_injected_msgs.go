package msgs

import (
	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bitoro-network/chain/protocol/testutil/constants"
	bridgetypes "github.com/bitoro-network/chain/protocol/x/bridge/types"
	clobtypes "github.com/bitoro-network/chain/protocol/x/clob/types"
	perptypes "github.com/bitoro-network/chain/protocol/x/perpetuals/types"
	pricestypes "github.com/bitoro-network/chain/protocol/x/prices/types"
)

var (
	// AppInjectedMsgSamples are msgs that are injected into the block by the proposing validator.
	// These messages are reserved for proposing validator's use only.
	AppInjectedMsgSamples = map[string]sdk.Msg{
		// bridge
		"/bitoroprotocol.bridge.MsgAcknowledgeBridges": &bridgetypes.MsgAcknowledgeBridges{
			Events: []bridgetypes.BridgeEvent{
				{
					Id: 0,
					Coin: sdk.NewCoin(
						"bridge-token",
						sdkmath.NewIntFromUint64(1234),
					),
					Address: constants.Alice_Num0.Owner,
				},
			},
		},
		"/bitoroprotocol.bridge.MsgAcknowledgeBridgesResponse": nil,

		// clob
		"/bitoroprotocol.clob.MsgProposedOperations": &clobtypes.MsgProposedOperations{
			OperationsQueue: make([]clobtypes.OperationRaw, 0),
		},
		"/bitoroprotocol.clob.MsgProposedOperationsResponse": nil,

		// perpetuals
		"/bitoroprotocol.perpetuals.MsgAddPremiumVotes": &perptypes.MsgAddPremiumVotes{
			Votes: []perptypes.FundingPremium{
				{PerpetualId: 0, PremiumPpm: 1_000},
			},
		},
		"/bitoroprotocol.perpetuals.MsgAddPremiumVotesResponse": nil,

		// prices
		"/bitoroprotocol.prices.MsgUpdateMarketPrices": &pricestypes.MsgUpdateMarketPrices{
			MarketPriceUpdates: []*pricestypes.MsgUpdateMarketPrices_MarketPrice{
				pricestypes.NewMarketPriceUpdate(constants.MarketId0, 123_000),
			},
		},
		"/bitoroprotocol.prices.MsgUpdateMarketPricesResponse": nil,
	}
)
