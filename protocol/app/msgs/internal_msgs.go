package msgs

import (
	upgrade "cosmossdk.io/x/upgrade/types"
	"github.com/bitoro-network/chain/protocol/lib"
	accountplus "github.com/bitoro-network/chain/protocol/x/accountplus/types"
	affiliates "github.com/bitoro-network/chain/protocol/x/affiliates/types"
	blocktime "github.com/bitoro-network/chain/protocol/x/blocktime/types"
	bridge "github.com/bitoro-network/chain/protocol/x/bridge/types"
	clob "github.com/bitoro-network/chain/protocol/x/clob/types"
	delaymsg "github.com/bitoro-network/chain/protocol/x/delaymsg/types"
	feetiers "github.com/bitoro-network/chain/protocol/x/feetiers/types"
	govplus "github.com/bitoro-network/chain/protocol/x/govplus/types"
	listing "github.com/bitoro-network/chain/protocol/x/listing/types"
	perpetuals "github.com/bitoro-network/chain/protocol/x/perpetuals/types"
	prices "github.com/bitoro-network/chain/protocol/x/prices/types"
	ratelimit "github.com/bitoro-network/chain/protocol/x/ratelimit/types"
	revshare "github.com/bitoro-network/chain/protocol/x/revshare/types"
	rewards "github.com/bitoro-network/chain/protocol/x/rewards/types"
	sending "github.com/bitoro-network/chain/protocol/x/sending/types"
	stats "github.com/bitoro-network/chain/protocol/x/stats/types"
	vault "github.com/bitoro-network/chain/protocol/x/vault/types"
	vest "github.com/bitoro-network/chain/protocol/x/vest/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	auth "github.com/cosmos/cosmos-sdk/x/auth/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	consensus "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisis "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distribution "github.com/cosmos/cosmos-sdk/x/distribution/types"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	slashing "github.com/cosmos/cosmos-sdk/x/slashing/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
	icahosttypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/host/types"
	ibctransfer "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	ibcclient "github.com/cosmos/ibc-go/v8/modules/core/02-client/types" //nolint:staticcheck
	ibcconn "github.com/cosmos/ibc-go/v8/modules/core/03-connection/types"
)

var (
	// InternalMsgSamplesAll are msgs that are used only used internally.
	InternalMsgSamplesAll = lib.MergeAllMapsMustHaveDistinctKeys(InternalMsgSamplesGovAuth)

	// InternalMsgSamplesGovAuth are msgs that are used only used internally.
	// GovAuth means that these messages must originate from the gov module and
	// signed by gov module account.
	// InternalMsgSamplesAll are msgs that are used only used internally.
	InternalMsgSamplesGovAuth = lib.MergeAllMapsMustHaveDistinctKeys(
		InternalMsgSamplesDefault,
		InternalMsgSamplesBitoroCustom,
	)

	// CosmosSDK default modules
	InternalMsgSamplesDefault = map[string]sdk.Msg{
		// auth
		"/cosmos.auth.v1beta1.MsgUpdateParams": &auth.MsgUpdateParams{},

		// bank
		"/cosmos.bank.v1beta1.MsgSetSendEnabled":         &bank.MsgSetSendEnabled{},
		"/cosmos.bank.v1beta1.MsgSetSendEnabledResponse": nil,
		"/cosmos.bank.v1beta1.MsgUpdateParams":           &bank.MsgUpdateParams{},
		"/cosmos.bank.v1beta1.MsgUpdateParamsResponse":   nil,

		// consensus
		"/cosmos.consensus.v1.MsgUpdateParams":         &consensus.MsgUpdateParams{},
		"/cosmos.consensus.v1.MsgUpdateParamsResponse": nil,

		// crisis
		"/cosmos.crisis.v1beta1.MsgUpdateParams":         &crisis.MsgUpdateParams{},
		"/cosmos.crisis.v1beta1.MsgUpdateParamsResponse": nil,

		// distribution
		"/cosmos.distribution.v1beta1.MsgCommunityPoolSpend":         &distribution.MsgCommunityPoolSpend{},
		"/cosmos.distribution.v1beta1.MsgCommunityPoolSpendResponse": nil,
		"/cosmos.distribution.v1beta1.MsgUpdateParams":               &distribution.MsgUpdateParams{},
		"/cosmos.distribution.v1beta1.MsgUpdateParamsResponse":       nil,

		// gov
		"/cosmos.gov.v1.MsgExecLegacyContent":         &gov.MsgExecLegacyContent{},
		"/cosmos.gov.v1.MsgExecLegacyContentResponse": nil,
		"/cosmos.gov.v1.MsgUpdateParams":              &gov.MsgUpdateParams{},
		"/cosmos.gov.v1.MsgUpdateParamsResponse":      nil,

		// slashing
		"/cosmos.slashing.v1beta1.MsgUpdateParams":         &slashing.MsgUpdateParams{},
		"/cosmos.slashing.v1beta1.MsgUpdateParamsResponse": nil,

		// staking
		"/cosmos.staking.v1beta1.MsgUpdateParams":         &staking.MsgUpdateParams{},
		"/cosmos.staking.v1beta1.MsgUpdateParamsResponse": nil,

		// upgrade
		"/cosmos.upgrade.v1beta1.MsgCancelUpgrade":           &upgrade.MsgCancelUpgrade{},
		"/cosmos.upgrade.v1beta1.MsgCancelUpgradeResponse":   nil,
		"/cosmos.upgrade.v1beta1.MsgSoftwareUpgrade":         &upgrade.MsgSoftwareUpgrade{},
		"/cosmos.upgrade.v1beta1.MsgSoftwareUpgradeResponse": nil,

		// ibc
		"/ibc.applications.interchain_accounts.host.v1.MsgUpdateParams":         &icahosttypes.MsgUpdateParams{},
		"/ibc.applications.interchain_accounts.host.v1.MsgUpdateParamsResponse": nil,
		"/ibc.applications.transfer.v1.MsgUpdateParams":                         &ibctransfer.MsgUpdateParams{},
		"/ibc.applications.transfer.v1.MsgUpdateParamsResponse":                 nil,
		"/ibc.core.client.v1.MsgUpdateParams":                                   &ibcclient.MsgUpdateParams{},
		"/ibc.core.client.v1.MsgUpdateParamsResponse":                           nil,
		"/ibc.core.connection.v1.MsgUpdateParams":                               &ibcconn.MsgUpdateParams{},
		"/ibc.core.connection.v1.MsgUpdateParamsResponse":                       nil,
	}

	// Custom modules
	InternalMsgSamplesBitoroCustom = map[string]sdk.Msg{
		// affiliates
		"/bitoroprotocol.affiliates.MsgUpdateAffiliateTiers":             &affiliates.MsgUpdateAffiliateTiers{},
		"/bitoroprotocol.affiliates.MsgUpdateAffiliateTiersResponse":     nil,
		"/bitoroprotocol.affiliates.MsgUpdateAffiliateWhitelist":         &affiliates.MsgUpdateAffiliateWhitelist{},
		"/bitoroprotocol.affiliates.MsgUpdateAffiliateWhitelistResponse": nil,

		// accountplus
		"/bitoroprotocol.accountplus.MsgSetActiveState":         &accountplus.MsgSetActiveState{},
		"/bitoroprotocol.accountplus.MsgSetActiveStateResponse": nil,

		// blocktime
		"/bitoroprotocol.blocktime.MsgUpdateDowntimeParams":          &blocktime.MsgUpdateDowntimeParams{},
		"/bitoroprotocol.blocktime.MsgUpdateDowntimeParamsResponse":  nil,
		"/bitoroprotocol.blocktime.MsgUpdateSynchronyParams":         &blocktime.MsgUpdateSynchronyParams{},
		"/bitoroprotocol.blocktime.MsgUpdateSynchronyParamsResponse": nil,

		// bridge
		"/bitoroprotocol.bridge.MsgCompleteBridge":              &bridge.MsgCompleteBridge{},
		"/bitoroprotocol.bridge.MsgCompleteBridgeResponse":      nil,
		"/bitoroprotocol.bridge.MsgUpdateEventParams":           &bridge.MsgUpdateEventParams{},
		"/bitoroprotocol.bridge.MsgUpdateEventParamsResponse":   nil,
		"/bitoroprotocol.bridge.MsgUpdateProposeParams":         &bridge.MsgUpdateProposeParams{},
		"/bitoroprotocol.bridge.MsgUpdateProposeParamsResponse": nil,
		"/bitoroprotocol.bridge.MsgUpdateSafetyParams":          &bridge.MsgUpdateSafetyParams{},
		"/bitoroprotocol.bridge.MsgUpdateSafetyParamsResponse":  nil,

		// clob
		"/bitoroprotocol.clob.MsgCreateClobPair":                             &clob.MsgCreateClobPair{},
		"/bitoroprotocol.clob.MsgCreateClobPairResponse":                     nil,
		"/bitoroprotocol.clob.MsgUpdateBlockRateLimitConfiguration":          &clob.MsgUpdateBlockRateLimitConfiguration{},
		"/bitoroprotocol.clob.MsgUpdateBlockRateLimitConfigurationResponse":  nil,
		"/bitoroprotocol.clob.MsgUpdateClobPair":                             &clob.MsgUpdateClobPair{},
		"/bitoroprotocol.clob.MsgUpdateClobPairResponse":                     nil,
		"/bitoroprotocol.clob.MsgUpdateEquityTierLimitConfiguration":         &clob.MsgUpdateEquityTierLimitConfiguration{},
		"/bitoroprotocol.clob.MsgUpdateEquityTierLimitConfigurationResponse": nil,
		"/bitoroprotocol.clob.MsgUpdateLiquidationsConfig":                   &clob.MsgUpdateLiquidationsConfig{},
		"/bitoroprotocol.clob.MsgUpdateLiquidationsConfigResponse":           nil,

		// delaymsg
		"/bitoroprotocol.delaymsg.MsgDelayMessage":         &delaymsg.MsgDelayMessage{},
		"/bitoroprotocol.delaymsg.MsgDelayMessageResponse": nil,

		// feetiers
		"/bitoroprotocol.feetiers.MsgUpdatePerpetualFeeParams":         &feetiers.MsgUpdatePerpetualFeeParams{},
		"/bitoroprotocol.feetiers.MsgUpdatePerpetualFeeParamsResponse": nil,

		// govplus
		"/bitoroprotocol.govplus.MsgSlashValidator":         &govplus.MsgSlashValidator{},
		"/bitoroprotocol.govplus.MsgSlashValidatorResponse": nil,

		// listing
		"/bitoroprotocol.listing.MsgSetMarketsHardCap":                       &listing.MsgSetMarketsHardCap{},
		"/bitoroprotocol.listing.MsgSetMarketsHardCapResponse":               nil,
		"/bitoroprotocol.listing.MsgSetListingVaultDepositParams":            &listing.MsgSetListingVaultDepositParams{},
		"/bitoroprotocol.listing.MsgSetListingVaultDepositParamsResponse":    nil,
		"/bitoroprotocol.listing.MsgUpgradeIsolatedPerpetualToCross":         &listing.MsgUpgradeIsolatedPerpetualToCross{},
		"/bitoroprotocol.listing.MsgUpgradeIsolatedPerpetualToCrossResponse": nil,

		// perpetuals
		"/bitoroprotocol.perpetuals.MsgCreatePerpetual":               &perpetuals.MsgCreatePerpetual{},
		"/bitoroprotocol.perpetuals.MsgCreatePerpetualResponse":       nil,
		"/bitoroprotocol.perpetuals.MsgSetLiquidityTier":              &perpetuals.MsgSetLiquidityTier{},
		"/bitoroprotocol.perpetuals.MsgSetLiquidityTierResponse":      nil,
		"/bitoroprotocol.perpetuals.MsgUpdateParams":                  &perpetuals.MsgUpdateParams{},
		"/bitoroprotocol.perpetuals.MsgUpdateParamsResponse":          nil,
		"/bitoroprotocol.perpetuals.MsgUpdatePerpetualParams":         &perpetuals.MsgUpdatePerpetualParams{},
		"/bitoroprotocol.perpetuals.MsgUpdatePerpetualParamsResponse": nil,

		// prices
		"/bitoroprotocol.prices.MsgCreateOracleMarket":         &prices.MsgCreateOracleMarket{},
		"/bitoroprotocol.prices.MsgCreateOracleMarketResponse": nil,
		"/bitoroprotocol.prices.MsgUpdateMarketParam":          &prices.MsgUpdateMarketParam{},
		"/bitoroprotocol.prices.MsgUpdateMarketParamResponse":  nil,

		// ratelimit
		"/bitoroprotocol.ratelimit.MsgSetLimitParams":         &ratelimit.MsgSetLimitParams{},
		"/bitoroprotocol.ratelimit.MsgSetLimitParamsResponse": nil,

		// revshare
		"/bitoroprotocol.revshare.MsgSetMarketMapperRevShareDetailsForMarket":         &revshare.MsgSetMarketMapperRevShareDetailsForMarket{}, //nolint:lll
		"/bitoroprotocol.revshare.MsgSetMarketMapperRevShareDetailsForMarketResponse": nil,
		"/bitoroprotocol.revshare.MsgSetMarketMapperRevenueShare":                     &revshare.MsgSetMarketMapperRevenueShare{}, //nolint:lll
		"/bitoroprotocol.revshare.MsgSetMarketMapperRevenueShareResponse":             nil,
		"/bitoroprotocol.revshare.MsgUpdateUnconditionalRevShareConfig":               &revshare.MsgUpdateUnconditionalRevShareConfig{}, //nolint:lll
		"/bitoroprotocol.revshare.MsgUpdateUnconditionalRevShareConfigResponse":       nil,

		// rewards
		"/bitoroprotocol.rewards.MsgUpdateParams":         &rewards.MsgUpdateParams{},
		"/bitoroprotocol.rewards.MsgUpdateParamsResponse": nil,

		// sending
		"/bitoroprotocol.sending.MsgSendFromModuleToAccount":         &sending.MsgSendFromModuleToAccount{},
		"/bitoroprotocol.sending.MsgSendFromModuleToAccountResponse": nil,

		// stats
		"/bitoroprotocol.stats.MsgUpdateParams":         &stats.MsgUpdateParams{},
		"/bitoroprotocol.stats.MsgUpdateParamsResponse": nil,

		// vault
		"/bitoroprotocol.vault.MsgUnlockShares":                 &vault.MsgUnlockShares{},
		"/bitoroprotocol.vault.MsgUnlockSharesResponse":         nil,
		"/bitoroprotocol.vault.MsgUpdateOperatorParams":         &vault.MsgUpdateOperatorParams{},
		"/bitoroprotocol.vault.MsgUpdateOperatorParamsResponse": nil,

		// vest
		"/bitoroprotocol.vest.MsgSetVestEntry":            &vest.MsgSetVestEntry{},
		"/bitoroprotocol.vest.MsgSetVestEntryResponse":    nil,
		"/bitoroprotocol.vest.MsgDeleteVestEntry":         &vest.MsgDeleteVestEntry{},
		"/bitoroprotocol.vest.MsgDeleteVestEntryResponse": nil,
	}
)
