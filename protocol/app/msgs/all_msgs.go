package msgs

import (
	"github.com/bitoro-network/chain/protocol/lib"
)

var (
	// AllTypeMessages is a list of all messages and types that are used in the app.
	// This list comes from the app's `InterfaceRegistry`.
	AllTypeMessages = map[string]struct{}{
		// auth
		"/cosmos.auth.v1beta1.BaseAccount":      {},
		"/cosmos.auth.v1beta1.ModuleAccount":    {},
		"/cosmos.auth.v1beta1.ModuleCredential": {},
		"/cosmos.auth.v1beta1.MsgUpdateParams":  {},

		// authz
		"/cosmos.authz.v1beta1.GenericAuthorization": {},
		"/cosmos.authz.v1beta1.MsgExec":              {},
		"/cosmos.authz.v1beta1.MsgExecResponse":      {},
		"/cosmos.authz.v1beta1.MsgGrant":             {},
		"/cosmos.authz.v1beta1.MsgGrantResponse":     {},
		"/cosmos.authz.v1beta1.MsgRevoke":            {},
		"/cosmos.authz.v1beta1.MsgRevokeResponse":    {},

		// bank
		"/cosmos.bank.v1beta1.MsgMultiSend":              {},
		"/cosmos.bank.v1beta1.MsgMultiSendResponse":      {},
		"/cosmos.bank.v1beta1.MsgSend":                   {},
		"/cosmos.bank.v1beta1.MsgSendResponse":           {},
		"/cosmos.bank.v1beta1.MsgSetSendEnabled":         {},
		"/cosmos.bank.v1beta1.MsgSetSendEnabledResponse": {},
		"/cosmos.bank.v1beta1.MsgUpdateParams":           {},
		"/cosmos.bank.v1beta1.MsgUpdateParamsResponse":   {},
		"/cosmos.bank.v1beta1.SendAuthorization":         {},
		"/cosmos.bank.v1beta1.Supply":                    {},

		// consensus
		"/cosmos.consensus.v1.MsgUpdateParams":         {},
		"/cosmos.consensus.v1.MsgUpdateParamsResponse": {},

		// crisis
		"/cosmos.crisis.v1beta1.MsgUpdateParams":            {},
		"/cosmos.crisis.v1beta1.MsgUpdateParamsResponse":    {},
		"/cosmos.crisis.v1beta1.MsgVerifyInvariant":         {},
		"/cosmos.crisis.v1beta1.MsgVerifyInvariantResponse": {},

		// crypto
		"/cosmos.crypto.ed25519.PrivKey":            {},
		"/cosmos.crypto.ed25519.PubKey":             {},
		"/cosmos.crypto.multisig.LegacyAminoPubKey": {},
		"/cosmos.crypto.secp256k1.PrivKey":          {},
		"/cosmos.crypto.secp256k1.PubKey":           {},
		"/cosmos.crypto.secp256r1.PubKey":           {},

		// distribution
		"/cosmos.distribution.v1beta1.CommunityPoolSpendProposal":             {},
		"/cosmos.distribution.v1beta1.MsgCommunityPoolSpend":                  {},
		"/cosmos.distribution.v1beta1.MsgCommunityPoolSpendResponse":          {},
		"/cosmos.distribution.v1beta1.MsgDepositValidatorRewardsPool":         {},
		"/cosmos.distribution.v1beta1.MsgDepositValidatorRewardsPoolResponse": {},
		"/cosmos.distribution.v1beta1.MsgFundCommunityPool":                   {},
		"/cosmos.distribution.v1beta1.MsgFundCommunityPoolResponse":           {},
		"/cosmos.distribution.v1beta1.MsgSetWithdrawAddress":                  {},
		"/cosmos.distribution.v1beta1.MsgSetWithdrawAddressResponse":          {},
		"/cosmos.distribution.v1beta1.MsgUpdateParams":                        {},
		"/cosmos.distribution.v1beta1.MsgUpdateParamsResponse":                {},
		"/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward":             {},
		"/cosmos.distribution.v1beta1.MsgWithdrawDelegatorRewardResponse":     {},
		"/cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission":         {},
		"/cosmos.distribution.v1beta1.MsgWithdrawValidatorCommissionResponse": {},

		// evidence
		"/cosmos.evidence.v1beta1.Equivocation":              {},
		"/cosmos.evidence.v1beta1.MsgSubmitEvidence":         {},
		"/cosmos.evidence.v1beta1.MsgSubmitEvidenceResponse": {},

		// feegrant
		"/cosmos.feegrant.v1beta1.AllowedMsgAllowance":        {},
		"/cosmos.feegrant.v1beta1.BasicAllowance":             {},
		"/cosmos.feegrant.v1beta1.MsgGrantAllowance":          {},
		"/cosmos.feegrant.v1beta1.MsgGrantAllowanceResponse":  {},
		"/cosmos.feegrant.v1beta1.MsgPruneAllowances":         {},
		"/cosmos.feegrant.v1beta1.MsgPruneAllowancesResponse": {},
		"/cosmos.feegrant.v1beta1.MsgRevokeAllowance":         {},
		"/cosmos.feegrant.v1beta1.MsgRevokeAllowanceResponse": {},
		"/cosmos.feegrant.v1beta1.PeriodicAllowance":          {},

		// gov
		"/cosmos.gov.v1.MsgCancelProposal":              {},
		"/cosmos.gov.v1.MsgCancelProposalResponse":      {},
		"/cosmos.gov.v1.MsgDeposit":                     {},
		"/cosmos.gov.v1.MsgDepositResponse":             {},
		"/cosmos.gov.v1.MsgExecLegacyContent":           {},
		"/cosmos.gov.v1.MsgExecLegacyContentResponse":   {},
		"/cosmos.gov.v1.MsgSubmitProposal":              {},
		"/cosmos.gov.v1.MsgSubmitProposalResponse":      {},
		"/cosmos.gov.v1.MsgUpdateParams":                {},
		"/cosmos.gov.v1.MsgUpdateParamsResponse":        {},
		"/cosmos.gov.v1.MsgVote":                        {},
		"/cosmos.gov.v1.MsgVoteResponse":                {},
		"/cosmos.gov.v1.MsgVoteWeighted":                {},
		"/cosmos.gov.v1.MsgVoteWeightedResponse":        {},
		"/cosmos.gov.v1beta1.MsgDeposit":                {},
		"/cosmos.gov.v1beta1.MsgDepositResponse":        {},
		"/cosmos.gov.v1beta1.MsgSubmitProposal":         {},
		"/cosmos.gov.v1beta1.MsgSubmitProposalResponse": {},
		"/cosmos.gov.v1beta1.MsgVote":                   {},
		"/cosmos.gov.v1beta1.MsgVoteResponse":           {},
		"/cosmos.gov.v1beta1.MsgVoteWeighted":           {},
		"/cosmos.gov.v1beta1.MsgVoteWeightedResponse":   {},
		"/cosmos.gov.v1beta1.TextProposal":              {},

		// params
		"/cosmos.params.v1beta1.ParameterChangeProposal": {},

		// slashing
		"/cosmos.slashing.v1beta1.MsgUnjail":               {},
		"/cosmos.slashing.v1beta1.MsgUnjailResponse":       {},
		"/cosmos.slashing.v1beta1.MsgUpdateParams":         {},
		"/cosmos.slashing.v1beta1.MsgUpdateParamsResponse": {},

		// staking
		"/cosmos.staking.v1beta1.MsgBeginRedelegate":                   {},
		"/cosmos.staking.v1beta1.MsgBeginRedelegateResponse":           {},
		"/cosmos.staking.v1beta1.MsgCancelUnbondingDelegation":         {},
		"/cosmos.staking.v1beta1.MsgCancelUnbondingDelegationResponse": {},
		"/cosmos.staking.v1beta1.MsgCreateValidator":                   {},
		"/cosmos.staking.v1beta1.MsgCreateValidatorResponse":           {},
		"/cosmos.staking.v1beta1.MsgDelegate":                          {},
		"/cosmos.staking.v1beta1.MsgDelegateResponse":                  {},
		"/cosmos.staking.v1beta1.MsgEditValidator":                     {},
		"/cosmos.staking.v1beta1.MsgEditValidatorResponse":             {},
		"/cosmos.staking.v1beta1.MsgUndelegate":                        {},
		"/cosmos.staking.v1beta1.MsgUndelegateResponse":                {},
		"/cosmos.staking.v1beta1.MsgUpdateParams":                      {},
		"/cosmos.staking.v1beta1.MsgUpdateParamsResponse":              {},
		"/cosmos.staking.v1beta1.StakeAuthorization":                   {},

		// tx
		"/cosmos.tx.v1beta1.Tx": {},

		// upgrade
		"/cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal": {},
		"/cosmos.upgrade.v1beta1.MsgCancelUpgrade":              {},
		"/cosmos.upgrade.v1beta1.MsgCancelUpgradeResponse":      {},
		"/cosmos.upgrade.v1beta1.MsgSoftwareUpgrade":            {},
		"/cosmos.upgrade.v1beta1.MsgSoftwareUpgradeResponse":    {},
		"/cosmos.upgrade.v1beta1.SoftwareUpgradeProposal":       {},

		// affiliates
		"/bitoroprotocol.affiliates.MsgRegisterAffiliate":                {},
		"/bitoroprotocol.affiliates.MsgRegisterAffiliateResponse":        {},
		"/bitoroprotocol.affiliates.MsgUpdateAffiliateTiers":             {},
		"/bitoroprotocol.affiliates.MsgUpdateAffiliateTiersResponse":     {},
		"/bitoroprotocol.affiliates.MsgUpdateAffiliateWhitelist":         {},
		"/bitoroprotocol.affiliates.MsgUpdateAffiliateWhitelistResponse": {},

		// accountplus
		"/bitoroprotocol.accountplus.MsgAddAuthenticator":            {},
		"/bitoroprotocol.accountplus.MsgAddAuthenticatorResponse":    {},
		"/bitoroprotocol.accountplus.MsgRemoveAuthenticator":         {},
		"/bitoroprotocol.accountplus.MsgRemoveAuthenticatorResponse": {},
		"/bitoroprotocol.accountplus.MsgSetActiveState":              {},
		"/bitoroprotocol.accountplus.MsgSetActiveStateResponse":      {},
		"/bitoroprotocol.accountplus.TxExtension":                    {},

		// blocktime
		"/bitoroprotocol.blocktime.MsgUpdateDowntimeParams":          {},
		"/bitoroprotocol.blocktime.MsgUpdateDowntimeParamsResponse":  {},
		"/bitoroprotocol.blocktime.MsgUpdateSynchronyParams":         {},
		"/bitoroprotocol.blocktime.MsgUpdateSynchronyParamsResponse": {},

		// bridge
		"/bitoroprotocol.bridge.MsgAcknowledgeBridges":          {},
		"/bitoroprotocol.bridge.MsgAcknowledgeBridgesResponse":  {},
		"/bitoroprotocol.bridge.MsgCompleteBridge":              {},
		"/bitoroprotocol.bridge.MsgCompleteBridgeResponse":      {},
		"/bitoroprotocol.bridge.MsgUpdateEventParams":           {},
		"/bitoroprotocol.bridge.MsgUpdateEventParamsResponse":   {},
		"/bitoroprotocol.bridge.MsgUpdateProposeParams":         {},
		"/bitoroprotocol.bridge.MsgUpdateProposeParamsResponse": {},
		"/bitoroprotocol.bridge.MsgUpdateSafetyParams":          {},
		"/bitoroprotocol.bridge.MsgUpdateSafetyParamsResponse":  {},

		// clob
		"/bitoroprotocol.clob.MsgBatchCancel":                                {},
		"/bitoroprotocol.clob.MsgBatchCancelResponse":                        {},
		"/bitoroprotocol.clob.MsgCancelOrder":                                {},
		"/bitoroprotocol.clob.MsgCancelOrderResponse":                        {},
		"/bitoroprotocol.clob.MsgCreateClobPair":                             {},
		"/bitoroprotocol.clob.MsgCreateClobPairResponse":                     {},
		"/bitoroprotocol.clob.MsgPlaceOrder":                                 {},
		"/bitoroprotocol.clob.MsgPlaceOrderResponse":                         {},
		"/bitoroprotocol.clob.MsgProposedOperations":                         {},
		"/bitoroprotocol.clob.MsgProposedOperationsResponse":                 {},
		"/bitoroprotocol.clob.MsgUpdateBlockRateLimitConfiguration":          {},
		"/bitoroprotocol.clob.MsgUpdateBlockRateLimitConfigurationResponse":  {},
		"/bitoroprotocol.clob.MsgUpdateClobPair":                             {},
		"/bitoroprotocol.clob.MsgUpdateClobPairResponse":                     {},
		"/bitoroprotocol.clob.MsgUpdateEquityTierLimitConfiguration":         {},
		"/bitoroprotocol.clob.MsgUpdateEquityTierLimitConfigurationResponse": {},
		"/bitoroprotocol.clob.MsgUpdateLiquidationsConfig":                   {},
		"/bitoroprotocol.clob.MsgUpdateLiquidationsConfigResponse":           {},

		// delaymsg
		"/bitoroprotocol.delaymsg.MsgDelayMessage":         {},
		"/bitoroprotocol.delaymsg.MsgDelayMessageResponse": {},

		// feetiers
		"/bitoroprotocol.feetiers.MsgUpdatePerpetualFeeParams":         {},
		"/bitoroprotocol.feetiers.MsgUpdatePerpetualFeeParamsResponse": {},

		// govplus
		"/bitoroprotocol.govplus.MsgSlashValidator":         {},
		"/bitoroprotocol.govplus.MsgSlashValidatorResponse": {},

		// listing
		"/bitoroprotocol.listing.MsgSetMarketsHardCap":                       {},
		"/bitoroprotocol.listing.MsgSetMarketsHardCapResponse":               {},
		"/bitoroprotocol.listing.MsgCreateMarketPermissionless":              {},
		"/bitoroprotocol.listing.MsgCreateMarketPermissionlessResponse":      {},
		"/bitoroprotocol.listing.MsgSetListingVaultDepositParams":            {},
		"/bitoroprotocol.listing.MsgSetListingVaultDepositParamsResponse":    {},
		"/bitoroprotocol.listing.MsgUpgradeIsolatedPerpetualToCross":         {},
		"/bitoroprotocol.listing.MsgUpgradeIsolatedPerpetualToCrossResponse": {},

		// perpetuals
		"/bitoroprotocol.perpetuals.MsgAddPremiumVotes":               {},
		"/bitoroprotocol.perpetuals.MsgAddPremiumVotesResponse":       {},
		"/bitoroprotocol.perpetuals.MsgCreatePerpetual":               {},
		"/bitoroprotocol.perpetuals.MsgCreatePerpetualResponse":       {},
		"/bitoroprotocol.perpetuals.MsgSetLiquidityTier":              {},
		"/bitoroprotocol.perpetuals.MsgSetLiquidityTierResponse":      {},
		"/bitoroprotocol.perpetuals.MsgUpdateParams":                  {},
		"/bitoroprotocol.perpetuals.MsgUpdateParamsResponse":          {},
		"/bitoroprotocol.perpetuals.MsgUpdatePerpetualParams":         {},
		"/bitoroprotocol.perpetuals.MsgUpdatePerpetualParamsResponse": {},

		// prices
		"/bitoroprotocol.prices.MsgCreateOracleMarket":         {},
		"/bitoroprotocol.prices.MsgCreateOracleMarketResponse": {},
		"/bitoroprotocol.prices.MsgUpdateMarketPrices":         {},
		"/bitoroprotocol.prices.MsgUpdateMarketPricesResponse": {},
		"/bitoroprotocol.prices.MsgUpdateMarketParam":          {},
		"/bitoroprotocol.prices.MsgUpdateMarketParamResponse":  {},

		// ratelimit
		"/bitoroprotocol.ratelimit.MsgSetLimitParams":         {},
		"/bitoroprotocol.ratelimit.MsgSetLimitParamsResponse": {},

		// sending
		"/bitoroprotocol.sending.MsgCreateTransfer":                  {},
		"/bitoroprotocol.sending.MsgCreateTransferResponse":          {},
		"/bitoroprotocol.sending.MsgDepositToSubaccount":             {},
		"/bitoroprotocol.sending.MsgDepositToSubaccountResponse":     {},
		"/bitoroprotocol.sending.MsgWithdrawFromSubaccount":          {},
		"/bitoroprotocol.sending.MsgWithdrawFromSubaccountResponse":  {},
		"/bitoroprotocol.sending.MsgSendFromModuleToAccount":         {},
		"/bitoroprotocol.sending.MsgSendFromModuleToAccountResponse": {},

		// stats
		"/bitoroprotocol.stats.MsgUpdateParams":         {},
		"/bitoroprotocol.stats.MsgUpdateParamsResponse": {},

		// vault
		"/bitoroprotocol.vault.MsgAllocateToVault":                    {},
		"/bitoroprotocol.vault.MsgAllocateToVaultResponse":            {},
		"/bitoroprotocol.vault.MsgDepositToMegavault":                 {},
		"/bitoroprotocol.vault.MsgDepositToMegavaultResponse":         {},
		"/bitoroprotocol.vault.MsgRetrieveFromVault":                  {},
		"/bitoroprotocol.vault.MsgRetrieveFromVaultResponse":          {},
		"/bitoroprotocol.vault.MsgSetVaultParams":                     {},
		"/bitoroprotocol.vault.MsgSetVaultParamsResponse":             {},
		"/bitoroprotocol.vault.MsgSetVaultQuotingParams":              {}, // deprecated
		"/bitoroprotocol.vault.MsgUnlockShares":                       {},
		"/bitoroprotocol.vault.MsgUnlockSharesResponse":               {},
		"/bitoroprotocol.vault.MsgUpdateDefaultQuotingParams":         {},
		"/bitoroprotocol.vault.MsgUpdateDefaultQuotingParamsResponse": {},
		"/bitoroprotocol.vault.MsgUpdateOperatorParams":               {},
		"/bitoroprotocol.vault.MsgUpdateOperatorParamsResponse":       {},
		"/bitoroprotocol.vault.MsgUpdateParams":                       {}, // deprecated
		"/bitoroprotocol.vault.MsgWithdrawFromMegavault":              {},
		"/bitoroprotocol.vault.MsgWithdrawFromMegavaultResponse":      {},

		// vest
		"/bitoroprotocol.vest.MsgSetVestEntry":            {},
		"/bitoroprotocol.vest.MsgSetVestEntryResponse":    {},
		"/bitoroprotocol.vest.MsgDeleteVestEntry":         {},
		"/bitoroprotocol.vest.MsgDeleteVestEntryResponse": {},

		// revshare
		"/bitoroprotocol.revshare.MsgSetMarketMapperRevenueShare":                     {},
		"/bitoroprotocol.revshare.MsgSetMarketMapperRevenueShareResponse":             {},
		"/bitoroprotocol.revshare.MsgSetMarketMapperRevShareDetailsForMarket":         {},
		"/bitoroprotocol.revshare.MsgSetMarketMapperRevShareDetailsForMarketResponse": {},
		"/bitoroprotocol.revshare.MsgUpdateUnconditionalRevShareConfig":               {},
		"/bitoroprotocol.revshare.MsgUpdateUnconditionalRevShareConfigResponse":       {},

		// rewards
		"/bitoroprotocol.rewards.MsgUpdateParams":         {},
		"/bitoroprotocol.rewards.MsgUpdateParamsResponse": {},

		// ibc.applications
		"/ibc.applications.transfer.v1.MsgTransfer":             {},
		"/ibc.applications.transfer.v1.MsgTransferResponse":     {},
		"/ibc.applications.transfer.v1.MsgUpdateParams":         {},
		"/ibc.applications.transfer.v1.MsgUpdateParamsResponse": {},
		"/ibc.applications.transfer.v1.TransferAuthorization":   {},

		// ibc.core.channel
		"/ibc.core.channel.v1.Channel":                        {},
		"/ibc.core.channel.v1.Counterparty":                   {},
		"/ibc.core.channel.v1.MsgAcknowledgement":             {},
		"/ibc.core.channel.v1.MsgAcknowledgementResponse":     {},
		"/ibc.core.channel.v1.MsgChannelCloseConfirm":         {},
		"/ibc.core.channel.v1.MsgChannelCloseConfirmResponse": {},
		"/ibc.core.channel.v1.MsgChannelCloseInit":            {},
		"/ibc.core.channel.v1.MsgChannelCloseInitResponse":    {},
		"/ibc.core.channel.v1.MsgChannelOpenAck":              {},
		"/ibc.core.channel.v1.MsgChannelOpenAckResponse":      {},
		"/ibc.core.channel.v1.MsgChannelOpenConfirm":          {},
		"/ibc.core.channel.v1.MsgChannelOpenConfirmResponse":  {},
		"/ibc.core.channel.v1.MsgChannelOpenInit":             {},
		"/ibc.core.channel.v1.MsgChannelOpenInitResponse":     {},
		"/ibc.core.channel.v1.MsgChannelOpenTry":              {},
		"/ibc.core.channel.v1.MsgChannelOpenTryResponse":      {},
		"/ibc.core.channel.v1.MsgRecvPacket":                  {},
		"/ibc.core.channel.v1.MsgRecvPacketResponse":          {},
		"/ibc.core.channel.v1.MsgTimeout":                     {},
		"/ibc.core.channel.v1.MsgTimeoutOnClose":              {},
		"/ibc.core.channel.v1.MsgTimeoutOnCloseResponse":      {},
		"/ibc.core.channel.v1.MsgTimeoutResponse":             {},
		"/ibc.core.channel.v1.Packet":                         {},

		// ibc.core.client
		"/ibc.core.client.v1.ClientUpdateProposal":          {},
		"/ibc.core.client.v1.Height":                        {},
		"/ibc.core.client.v1.MsgCreateClient":               {},
		"/ibc.core.client.v1.MsgCreateClientResponse":       {},
		"/ibc.core.client.v1.MsgIBCSoftwareUpgrade":         {},
		"/ibc.core.client.v1.MsgIBCSoftwareUpgradeResponse": {},
		"/ibc.core.client.v1.MsgRecoverClient":              {},
		"/ibc.core.client.v1.MsgRecoverClientResponse":      {},
		"/ibc.core.client.v1.MsgSubmitMisbehaviour":         {},
		"/ibc.core.client.v1.MsgSubmitMisbehaviourResponse": {},
		"/ibc.core.client.v1.MsgUpdateClient":               {},
		"/ibc.core.client.v1.MsgUpdateClientResponse":       {},
		"/ibc.core.client.v1.MsgUpgradeClient":              {},
		"/ibc.core.client.v1.MsgUpgradeClientResponse":      {},
		"/ibc.core.client.v1.MsgUpdateParams":               {},
		"/ibc.core.client.v1.MsgUpdateParamsResponse":       {},
		"/ibc.core.client.v1.UpgradeProposal":               {},

		// ibc.core.commitment
		"/ibc.core.commitment.v1.MerklePath":   {},
		"/ibc.core.commitment.v1.MerklePrefix": {},
		"/ibc.core.commitment.v1.MerkleProof":  {},
		"/ibc.core.commitment.v1.MerkleRoot":   {},

		// ibc.core.connection
		"/ibc.core.connection.v1.ConnectionEnd":                    {},
		"/ibc.core.connection.v1.Counterparty":                     {},
		"/ibc.core.connection.v1.MsgConnectionOpenAck":             {},
		"/ibc.core.connection.v1.MsgConnectionOpenAckResponse":     {},
		"/ibc.core.connection.v1.MsgConnectionOpenConfirm":         {},
		"/ibc.core.connection.v1.MsgConnectionOpenConfirmResponse": {},
		"/ibc.core.connection.v1.MsgConnectionOpenInit":            {},
		"/ibc.core.connection.v1.MsgConnectionOpenInitResponse":    {},
		"/ibc.core.connection.v1.MsgConnectionOpenTry":             {},
		"/ibc.core.connection.v1.MsgConnectionOpenTryResponse":     {},
		"/ibc.core.connection.v1.MsgUpdateParams":                  {},
		"/ibc.core.connection.v1.MsgUpdateParamsResponse":          {},

		// ibc.lightclients
		"/ibc.lightclients.localhost.v2.ClientState":     {},
		"/ibc.lightclients.tendermint.v1.ClientState":    {},
		"/ibc.lightclients.tendermint.v1.ConsensusState": {},
		"/ibc.lightclients.tendermint.v1.Header":         {},
		"/ibc.lightclients.tendermint.v1.Misbehaviour":   {},

		// ica messages
		// Note: the `interchain_accounts.controller` messages are not actually used by the app,
		// since ICA Controller Keeper is initialized as nil.
		// However, since the ica.AppModuleBasic{} needs to be passed to basic_mananger as a whole, these messages
		// registered in the interface registry.
		"/ibc.applications.interchain_accounts.v1.InterchainAccount":                               {},
		"/ibc.applications.interchain_accounts.controller.v1.MsgSendTx":                            {},
		"/ibc.applications.interchain_accounts.controller.v1.MsgSendTxResponse":                    {},
		"/ibc.applications.interchain_accounts.controller.v1.MsgRegisterInterchainAccount":         {},
		"/ibc.applications.interchain_accounts.controller.v1.MsgRegisterInterchainAccountResponse": {},
		"/ibc.applications.interchain_accounts.controller.v1.MsgUpdateParams":                      {},
		"/ibc.applications.interchain_accounts.controller.v1.MsgUpdateParamsResponse":              {},
		"/ibc.applications.interchain_accounts.host.v1.MsgUpdateParams":                            {},
		"/ibc.applications.interchain_accounts.host.v1.MsgUpdateParamsResponse":                    {},

		// slinky marketmap messages
		"/slinky.marketmap.v1.MsgCreateMarkets":                   {},
		"/slinky.marketmap.v1.MsgCreateMarketsResponse":           {},
		"/slinky.marketmap.v1.MsgParams":                          {},
		"/slinky.marketmap.v1.MsgParamsResponse":                  {},
		"/slinky.marketmap.v1.MsgRemoveMarkets":                   {},
		"/slinky.marketmap.v1.MsgRemoveMarketsResponse":           {},
		"/slinky.marketmap.v1.MsgRemoveMarketAuthorities":         {},
		"/slinky.marketmap.v1.MsgRemoveMarketAuthoritiesResponse": {},
		"/slinky.marketmap.v1.MsgUpdateMarkets":                   {},
		"/slinky.marketmap.v1.MsgUpdateMarketsResponse":           {},
		"/slinky.marketmap.v1.MsgUpsertMarkets":                   {},
		"/slinky.marketmap.v1.MsgUpsertMarketsResponse":           {},
	}

	// DisallowMsgs are messages that cannot be externally submitted.
	DisallowMsgs = lib.MergeAllMapsMustHaveDistinctKeys(
		AppInjectedMsgSamples,
		InternalMsgSamplesAll,
		NestedMsgSamples,
		UnsupportedMsgSamples,
	)

	// AllowMsgs are messages that can be externally submitted.
	AllowMsgs = NormalMsgs
)
