package msgs_test

import (
	"sort"
	"testing"

	"github.com/bitoro-network/chain/protocol/app/msgs"
	"github.com/bitoro-network/chain/protocol/lib"
	"github.com/stretchr/testify/require"
)

func TestInternalMsgSamples_All_Key(t *testing.T) {
	expectedAllInternalMsgs := lib.MergeAllMapsMustHaveDistinctKeys(msgs.InternalMsgSamplesGovAuth)
	require.Equal(t, expectedAllInternalMsgs, msgs.InternalMsgSamplesAll)
}

func TestInternalMsgSamples_All_Value(t *testing.T) {
	validateMsgValue(t, msgs.InternalMsgSamplesAll)
}

func TestInternalMsgSamples_Gov_Key(t *testing.T) {
	expectedMsgs := []string{
		// auth
		"/cosmos.auth.v1beta1.MsgUpdateParams",

		// bank
		"/cosmos.bank.v1beta1.MsgSetSendEnabled",
		"/cosmos.bank.v1beta1.MsgSetSendEnabledResponse",
		"/cosmos.bank.v1beta1.MsgUpdateParams",
		"/cosmos.bank.v1beta1.MsgUpdateParamsResponse",

		// consensus
		"/cosmos.consensus.v1.MsgUpdateParams",
		"/cosmos.consensus.v1.MsgUpdateParamsResponse",

		// crisis
		"/cosmos.crisis.v1beta1.MsgUpdateParams",
		"/cosmos.crisis.v1beta1.MsgUpdateParamsResponse",

		// distribution
		"/cosmos.distribution.v1beta1.MsgCommunityPoolSpend",
		"/cosmos.distribution.v1beta1.MsgCommunityPoolSpendResponse",
		"/cosmos.distribution.v1beta1.MsgUpdateParams",
		"/cosmos.distribution.v1beta1.MsgUpdateParamsResponse",

		// gov
		"/cosmos.gov.v1.MsgExecLegacyContent",
		"/cosmos.gov.v1.MsgExecLegacyContentResponse",
		"/cosmos.gov.v1.MsgUpdateParams",
		"/cosmos.gov.v1.MsgUpdateParamsResponse",

		// slashing
		"/cosmos.slashing.v1beta1.MsgUpdateParams",
		"/cosmos.slashing.v1beta1.MsgUpdateParamsResponse",

		// staking
		"/cosmos.staking.v1beta1.MsgUpdateParams",
		"/cosmos.staking.v1beta1.MsgUpdateParamsResponse",

		// upgrade
		"/cosmos.upgrade.v1beta1.MsgCancelUpgrade",
		"/cosmos.upgrade.v1beta1.MsgCancelUpgradeResponse",
		"/cosmos.upgrade.v1beta1.MsgSoftwareUpgrade",
		"/cosmos.upgrade.v1beta1.MsgSoftwareUpgradeResponse",

		// accountplus
		"/bitoroprotocol.accountplus.MsgSetActiveState",
		"/bitoroprotocol.accountplus.MsgSetActiveStateResponse",

		// affiliates
		"/bitoroprotocol.affiliates.MsgUpdateAffiliateTiers",
		"/bitoroprotocol.affiliates.MsgUpdateAffiliateTiersResponse",
		"/bitoroprotocol.affiliates.MsgUpdateAffiliateWhitelist",
		"/bitoroprotocol.affiliates.MsgUpdateAffiliateWhitelistResponse",

		// blocktime
		"/bitoroprotocol.blocktime.MsgUpdateDowntimeParams",
		"/bitoroprotocol.blocktime.MsgUpdateDowntimeParamsResponse",
		"/bitoroprotocol.blocktime.MsgUpdateSynchronyParams",
		"/bitoroprotocol.blocktime.MsgUpdateSynchronyParamsResponse",

		// bridge
		"/bitoroprotocol.bridge.MsgCompleteBridge",
		"/bitoroprotocol.bridge.MsgCompleteBridgeResponse",
		"/bitoroprotocol.bridge.MsgUpdateEventParams",
		"/bitoroprotocol.bridge.MsgUpdateEventParamsResponse",
		"/bitoroprotocol.bridge.MsgUpdateProposeParams",
		"/bitoroprotocol.bridge.MsgUpdateProposeParamsResponse",
		"/bitoroprotocol.bridge.MsgUpdateSafetyParams",
		"/bitoroprotocol.bridge.MsgUpdateSafetyParamsResponse",

		// clob
		"/bitoroprotocol.clob.MsgCreateClobPair",
		"/bitoroprotocol.clob.MsgCreateClobPairResponse",
		"/bitoroprotocol.clob.MsgUpdateBlockRateLimitConfiguration",
		"/bitoroprotocol.clob.MsgUpdateBlockRateLimitConfigurationResponse",
		"/bitoroprotocol.clob.MsgUpdateClobPair",
		"/bitoroprotocol.clob.MsgUpdateClobPairResponse",
		"/bitoroprotocol.clob.MsgUpdateEquityTierLimitConfiguration",
		"/bitoroprotocol.clob.MsgUpdateEquityTierLimitConfigurationResponse",
		"/bitoroprotocol.clob.MsgUpdateLiquidationsConfig",
		"/bitoroprotocol.clob.MsgUpdateLiquidationsConfigResponse",

		// delaymsg
		"/bitoroprotocol.delaymsg.MsgDelayMessage",
		"/bitoroprotocol.delaymsg.MsgDelayMessageResponse",

		// feetiers
		"/bitoroprotocol.feetiers.MsgUpdatePerpetualFeeParams",
		"/bitoroprotocol.feetiers.MsgUpdatePerpetualFeeParamsResponse",

		// govplus
		"/bitoroprotocol.govplus.MsgSlashValidator",
		"/bitoroprotocol.govplus.MsgSlashValidatorResponse",

		// listing
		"/bitoroprotocol.listing.MsgSetListingVaultDepositParams",
		"/bitoroprotocol.listing.MsgSetListingVaultDepositParamsResponse",
		"/bitoroprotocol.listing.MsgSetMarketsHardCap",
		"/bitoroprotocol.listing.MsgSetMarketsHardCapResponse",
		"/bitoroprotocol.listing.MsgUpgradeIsolatedPerpetualToCross",
		"/bitoroprotocol.listing.MsgUpgradeIsolatedPerpetualToCrossResponse",

		// perpeutals
		"/bitoroprotocol.perpetuals.MsgCreatePerpetual",
		"/bitoroprotocol.perpetuals.MsgCreatePerpetualResponse",
		"/bitoroprotocol.perpetuals.MsgSetLiquidityTier",
		"/bitoroprotocol.perpetuals.MsgSetLiquidityTierResponse",
		"/bitoroprotocol.perpetuals.MsgUpdateParams",
		"/bitoroprotocol.perpetuals.MsgUpdateParamsResponse",
		"/bitoroprotocol.perpetuals.MsgUpdatePerpetualParams",
		"/bitoroprotocol.perpetuals.MsgUpdatePerpetualParamsResponse",

		// prices
		"/bitoroprotocol.prices.MsgCreateOracleMarket",
		"/bitoroprotocol.prices.MsgCreateOracleMarketResponse",
		"/bitoroprotocol.prices.MsgUpdateMarketParam",
		"/bitoroprotocol.prices.MsgUpdateMarketParamResponse",

		// ratelimit
		"/bitoroprotocol.ratelimit.MsgSetLimitParams",
		"/bitoroprotocol.ratelimit.MsgSetLimitParamsResponse",

		// revshare
		"/bitoroprotocol.revshare.MsgSetMarketMapperRevShareDetailsForMarket",
		"/bitoroprotocol.revshare.MsgSetMarketMapperRevShareDetailsForMarketResponse",
		"/bitoroprotocol.revshare.MsgSetMarketMapperRevenueShare",
		"/bitoroprotocol.revshare.MsgSetMarketMapperRevenueShareResponse",
		"/bitoroprotocol.revshare.MsgUpdateUnconditionalRevShareConfig",
		"/bitoroprotocol.revshare.MsgUpdateUnconditionalRevShareConfigResponse",

		// rewards
		"/bitoroprotocol.rewards.MsgUpdateParams",
		"/bitoroprotocol.rewards.MsgUpdateParamsResponse",

		// sending
		"/bitoroprotocol.sending.MsgSendFromModuleToAccount",
		"/bitoroprotocol.sending.MsgSendFromModuleToAccountResponse",

		// stats
		"/bitoroprotocol.stats.MsgUpdateParams",
		"/bitoroprotocol.stats.MsgUpdateParamsResponse",

		// vault
		"/bitoroprotocol.vault.MsgUnlockShares",
		"/bitoroprotocol.vault.MsgUnlockSharesResponse",
		"/bitoroprotocol.vault.MsgUpdateOperatorParams",
		"/bitoroprotocol.vault.MsgUpdateOperatorParamsResponse",

		// vest
		"/bitoroprotocol.vest.MsgDeleteVestEntry",
		"/bitoroprotocol.vest.MsgDeleteVestEntryResponse",
		"/bitoroprotocol.vest.MsgSetVestEntry",
		"/bitoroprotocol.vest.MsgSetVestEntryResponse",

		// ibc
		"/ibc.applications.interchain_accounts.host.v1.MsgUpdateParams",
		"/ibc.applications.interchain_accounts.host.v1.MsgUpdateParamsResponse",
		"/ibc.applications.transfer.v1.MsgUpdateParams",
		"/ibc.applications.transfer.v1.MsgUpdateParamsResponse",
		"/ibc.core.client.v1.MsgUpdateParams",
		"/ibc.core.client.v1.MsgUpdateParamsResponse",
		"/ibc.core.connection.v1.MsgUpdateParams",
		"/ibc.core.connection.v1.MsgUpdateParamsResponse",
	}

	require.Equal(t, expectedMsgs, lib.GetSortedKeys[sort.StringSlice](msgs.InternalMsgSamplesGovAuth))
}

func TestInternalMsgSamples_Gov_Value(t *testing.T) {
	validateMsgValue(t, msgs.InternalMsgSamplesGovAuth)
}
