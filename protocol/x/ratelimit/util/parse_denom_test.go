package util_test

import (
	"fmt"
	"testing"

	"github.com/bitoro-network/chain/protocol/lib"
	assettypes "github.com/bitoro-network/chain/protocol/x/assets/types"
	"github.com/bitoro-network/chain/protocol/x/ratelimit/util"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	"github.com/stretchr/testify/require"
)

func TestParseDenomFromRecvPacket(t *testing.T) {
	nobleChannelOnBitoro := "channel-0"
	nobleChannelOnOsmo := "channel-200"
	osmoChannelOnBitoro := "channel-5"
	bitoroChannelOnNoble := "channel-100"
	bitoroChannelOnOsmo := "channel-101"
	originalUsdcDenom := "uusdc"

	testCases := []struct {
		name               string
		packetDenomTrace   string
		sourceChannel      string
		destinationChannel string
		expectedDenom      string
	}{
		// Sink asset one hop away:
		//   uusdc sent from Noble to Bitoro
		//   -> tack on prefix (transfer/channel-0/uusdc) and hash
		{
			name:               "sink_one_hop",
			packetDenomTrace:   assettypes.UusdcDenom,
			sourceChannel:      bitoroChannelOnNoble,
			destinationChannel: nobleChannelOnBitoro,
			expectedDenom: hashDenomTrace(fmt.Sprintf(
				"%s/%s/%s",
				transferPort,
				nobleChannelOnBitoro,
				assettypes.UusdcDenom,
			)),
		},
		// Native source assets
		//    lib.DefaultBaseDenom sent from Bitoro to Noble and then back to Bitoro (transfer/channel-0/adv4tnt)
		//    -> remove prefix and leave as is (adv4tnt)
		{
			name:               lib.DefaultBaseDenom,
			packetDenomTrace:   fmt.Sprintf("%s/%s/%s", transferPort, bitoroChannelOnNoble, lib.DefaultBaseDenom),
			sourceChannel:      bitoroChannelOnNoble,
			destinationChannel: nobleChannelOnBitoro,
			expectedDenom:      lib.DefaultBaseDenom,
		},
		// Sink asset two hops away:
		//   uusdc sent from Noble to Osmosis to Bitoro (transfer/channel-200/uusdc)
		//   -> tack on prefix (transfer/channel-0/transfer/channel-200/uusdc) and hash
		{
			name:               "sink_two_hops",
			packetDenomTrace:   fmt.Sprintf("%s/%s/%s", transferPort, nobleChannelOnOsmo, originalUsdcDenom),
			sourceChannel:      bitoroChannelOnOsmo,
			destinationChannel: osmoChannelOnBitoro,
			expectedDenom: hashDenomTrace(
				fmt.Sprintf(
					"%s/%s/%s/%s/%s",
					transferPort,
					osmoChannelOnBitoro,
					transferPort,
					nobleChannelOnOsmo,
					originalUsdcDenom,
				),
			),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			packet := channeltypes.Packet{
				SourcePort:         transferPort,
				DestinationPort:    transferPort,
				SourceChannel:      tc.sourceChannel,
				DestinationChannel: tc.destinationChannel,
			}
			packetData := ibctransfertypes.FungibleTokenPacketData{
				Denom: tc.packetDenomTrace,
			}

			parsedDenom := util.ParseDenomFromRecvPacket(packet, packetData)
			require.Equal(t, tc.expectedDenom, parsedDenom, tc.name)
		})
	}
}

func TestParseDenomFromSendPacket(t *testing.T) {
	testCases := []struct {
		name             string
		packetDenomTrace string
		expectedDenom    string
	}{
		// Native assets stay as is
		{
			name:             "base denom",
			packetDenomTrace: lib.DefaultBaseDenom,
			expectedDenom:    lib.DefaultBaseDenom,
		},
		// Non-native assets are hashed
		{
			name:             "uusdc on Bitoro",
			packetDenomTrace: "transfer/channel-0/uusdc",
			expectedDenom:    assettypes.UusdcDenom,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			packet := ibctransfertypes.FungibleTokenPacketData{
				Denom: tc.packetDenomTrace,
			}

			parsedDenom := util.ParseDenomFromSendPacket(packet)
			require.Equal(t, tc.expectedDenom, parsedDenom, tc.name)
		})
	}
}
