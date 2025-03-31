import { KafkaTopics } from '@bitoroprotocol-indexer/kafka';
import {
  Liquidity,
  PerpetualPositionColumns,
  PerpetualPositionFromDatabase,
  SubaccountMessageContents,
} from '@bitoroprotocol-indexer/postgres';
import {
  PerpetualMarketCreateEventV2,
  PerpetualMarketCreateEventV3,
  StatefulOrderEventV1,
  IndexerTendermintEvent,
  CandleMessage,
  LiquidationOrderV1,
  MarketCreateEventV1,
  MarketEventV1,
  MarketMessage,
  MarketModifyEventV1,
  MarketPriceUpdateEventV1,
  IndexerOrder,
  OrderFillEventV1,
  SubaccountMessage,
  SubaccountUpdateEventV1,
  TradeMessage,
  TransferEventV1,
  OffChainUpdateV1,
  FundingEventV1_Type,
  FundingEventV1,
  FundingUpdateV1,
  AssetCreateEventV1,
  PerpetualMarketCreateEventV1,
  LiquidityTierUpsertEventV1,
  LiquidityTierUpsertEventV2,
  UpdatePerpetualEventV1,
  UpdatePerpetualEventV2,
  UpdatePerpetualEventV3,
  UpdateClobPairEventV1,
  DeleveragingEventV1,
  TradingRewardsEventV1,
  BlockHeightMessage,
  RegisterAffiliateEventV1,
  UpsertVaultEventV1,
} from '@bitoroprotocol-indexer/chain-protos';
import { IHeaders } from 'kafkajs';
import Long from 'long';

// Type sourced from protocol:
// https://github.com/bitoro-network/chain/blob/main/protocol/indexer/events/constants.go
export enum BitoroIndexerSubtypes {
  ORDER_FILL = 'order_fill',
  SUBACCOUNT_UPDATE = 'subaccount_update',
  TRANSFER = 'transfer',
  MARKET = 'market',
  STATEFUL_ORDER = 'stateful_order',
  FUNDING = 'funding_values',
  ASSET = 'asset',
  PERPETUAL_MARKET = 'perpetual_market',
  LIQUIDITY_TIER = 'liquidity_tier',
  UPDATE_PERPETUAL = 'update_perpetual',
  UPDATE_CLOB_PAIR = 'update_clob_pair',
  DELEVERAGING = 'deleveraging',
  TRADING_REWARD = 'trading_reward',
  REGISTER_AFFILIATE = 'register_affiliate',
  UPSERT_VAULT = 'upsert_vault',
}

export const SKIPPED_EVENT_SUBTYPE = 'skipped_event';

// Generic interface used for creating the Handler objects
// eslint-disable-next-line @typescript-eslint/no-explicit-any
export type EventMessage = any;

export type EventProtoWithTypeAndVersion = {
  type: string,
  eventProto: EventMessage,
  indexerTendermintEvent: IndexerTendermintEvent,
  version: number,
  blockEventIndex: number,
} & ({
  type: BitoroIndexerSubtypes.ORDER_FILL,
  eventProto: OrderFillEventV1,
  indexerTendermintEvent: IndexerTendermintEvent,
  version: number,
  blockEventIndex: number,
} | {
  type: BitoroIndexerSubtypes.SUBACCOUNT_UPDATE,
  eventProto: SubaccountUpdateEventV1,
  indexerTendermintEvent: IndexerTendermintEvent,
  version: number,
  blockEventIndex: number,
} | {
  type: BitoroIndexerSubtypes.TRANSFER,
  eventProto: TransferEventV1,
  indexerTendermintEvent: IndexerTendermintEvent,
  version: number,
  blockEventIndex: number,
} | {
  type: BitoroIndexerSubtypes.MARKET,
  eventProto: MarketEventV1,
  indexerTendermintEvent: IndexerTendermintEvent,
  version: number,
  blockEventIndex: number,
} | {
  type: BitoroIndexerSubtypes.STATEFUL_ORDER,
  eventProto: StatefulOrderEventV1,
  indexerTendermintEvent: IndexerTendermintEvent,
  version: number,
  blockEventIndex: number,
} | {
  type: BitoroIndexerSubtypes.FUNDING,
  eventProto: FundingEventV1,
  indexerTendermintEvent: IndexerTendermintEvent,
  version: number,
  blockEventIndex: number,
} | {
  type: BitoroIndexerSubtypes.ASSET,
  eventProto: AssetCreateEventV1,
  indexerTendermintEvent: IndexerTendermintEvent,
  version: number,
  blockEventIndex: number,
} | {
  type: BitoroIndexerSubtypes.PERPETUAL_MARKET,
  eventProto: PerpetualMarketCreateEventV1,
  indexerTendermintEvent: IndexerTendermintEvent,
  version: number,
  blockEventIndex: number,
} | {
  type: BitoroIndexerSubtypes.PERPETUAL_MARKET,
  eventProto: PerpetualMarketCreateEventV2,
  indexerTendermintEvent: IndexerTendermintEvent,
  version: number,
  blockEventIndex: number,
} | {
  type: BitoroIndexerSubtypes.PERPETUAL_MARKET,
  eventProto: PerpetualMarketCreateEventV3,
  indexerTendermintEvent: IndexerTendermintEvent,
  version: number,
  blockEventIndex: number,
} | {
  type: BitoroIndexerSubtypes.LIQUIDITY_TIER,
  eventProto: LiquidityTierUpsertEventV1,
  indexerTendermintEvent: IndexerTendermintEvent,
  version: number,
  blockEventIndex: number,
} | {
  type: BitoroIndexerSubtypes.LIQUIDITY_TIER,
  eventProto: LiquidityTierUpsertEventV2,
  indexerTendermintEvent: IndexerTendermintEvent,
  version: number,
  blockEventIndex: number,
} | {
  type: BitoroIndexerSubtypes.UPDATE_PERPETUAL,
  eventProto: UpdatePerpetualEventV1,
  indexerTendermintEvent: IndexerTendermintEvent,
  version: number,
  blockEventIndex: number,
} | {
  type: BitoroIndexerSubtypes.UPDATE_PERPETUAL,
  eventProto: UpdatePerpetualEventV2,
  indexerTendermintEvent: IndexerTendermintEvent,
  version: number,
  blockEventIndex: number,
} | {
  type: BitoroIndexerSubtypes.UPDATE_PERPETUAL,
  eventProto: UpdatePerpetualEventV3,
  indexerTendermintEvent: IndexerTendermintEvent,
  version: number,
  blockEventIndex: number,
} | {
  type: BitoroIndexerSubtypes.UPDATE_CLOB_PAIR,
  eventProto: UpdateClobPairEventV1,
  indexerTendermintEvent: IndexerTendermintEvent,
  version: number,
  blockEventIndex: number,
} | {
  type: BitoroIndexerSubtypes.DELEVERAGING,
  eventProto: DeleveragingEventV1,
  indexerTendermintEvent: IndexerTendermintEvent,
  version: number,
  blockEventIndex: number,
} | {
  type: BitoroIndexerSubtypes.TRADING_REWARD,
  eventProto: TradingRewardsEventV1,
  indexerTendermintEvent: IndexerTendermintEvent,
  version: number,
  blockEventIndex: number,
} | {
  type: BitoroIndexerSubtypes.REGISTER_AFFILIATE,
  eventProto: RegisterAffiliateEventV1,
  indexerTendermintEvent: IndexerTendermintEvent,
  version: number,
  blockEventIndex: number,
} | {
  type: BitoroIndexerSubtypes.UPSERT_VAULT,
  eventProto: UpsertVaultEventV1,
  indexerTendermintEvent: IndexerTendermintEvent,
  version: number,
  blockEventIndex: number,
});

// Events grouped into events block events and events for each transactionIndex
export interface GroupedEvents {
  transactionEvents: EventProtoWithTypeAndVersion[][],
  blockEvents: EventProtoWithTypeAndVersion[],
}

export type MarketPriceUpdateEventMessage = {
  marketId: number,
  priceUpdate: MarketPriceUpdateEventV1,
};

export type MarketCreateEventMessage = {
  marketId: number,
  marketCreate: MarketCreateEventV1,
};

export type MarketModifyEventMessage = {
  marketId: number,
  marketModify: MarketModifyEventV1,
};

export type OrderFillEventWithOrder = {
  makerOrder: IndexerOrder,
  order: IndexerOrder,
  fillAmount: Long,
  totalFilledMaker: Long,
  totalFilledTaker: Long,
  makerFee: Long,
  takerFee: Long,
  affiliateRevShare: Long,
};

export type OrderFillEventWithLiquidation = {
  makerOrder: IndexerOrder,
  liquidationOrder: LiquidationOrderV1,
  fillAmount: Long,
  totalFilledMaker: Long,
  totalFilledTaker: Long,
  makerFee: Long,
  takerFee: Long,
  affiliateRevShare: Long,
};

export type FundingEventMessage = {
  type: FundingEventV1_Type.TYPE_FUNDING_RATE_AND_INDEX | FundingEventV1_Type.TYPE_PREMIUM_SAMPLE,
  updates: FundingUpdateV1[],
};

export type SumFields = PerpetualPositionColumns.sumOpen | PerpetualPositionColumns.sumClose;
export type PriceFields = PerpetualPositionColumns.entryPrice |
PerpetualPositionColumns.exitPrice;

export type OrderFillEventWithLiquidity = {
  event: OrderFillEventV1,
  liquidity: Liquidity,
};

export interface PositionWithPnl extends PerpetualPositionFromDatabase {
  realizedPnl?: string,
  unrealizedPnl?: string,
}

export interface SingleTradeMessage extends TradeMessage {
  transactionIndex: number,
  eventIndex: number,
}

export interface AnnotatedSubaccountMessage extends SubaccountMessage {
  orderId?: string,
  isFill?: boolean,
  subaccountMessageContents?: SubaccountMessageContents,
}

export interface VulcanMessage {
  key: Buffer,
  value: OffChainUpdateV1,
  headers?: IHeaders,
}

export type ConsolidatedKafkaEvent = {
  topic: KafkaTopics.TO_WEBSOCKETS_SUBACCOUNTS,
  message: AnnotatedSubaccountMessage,
} | {
  topic: KafkaTopics.TO_WEBSOCKETS_TRADES,
  message: SingleTradeMessage,
} | {
  topic: KafkaTopics.TO_WEBSOCKETS_MARKETS,
  message: MarketMessage,
} | {
  topic: KafkaTopics.TO_WEBSOCKETS_CANDLES,
  message: CandleMessage,
} | {
  topic: KafkaTopics.TO_VULCAN,
  message: VulcanMessage,
} | {
  topic: KafkaTopics.TO_WEBSOCKETS_BLOCK_HEIGHT,
  message: BlockHeightMessage,
};

export enum TransferEventType {
  DEPOSIT = 'deposit',
  WITHDRAWAL = 'withdrawal',
  TRANSFER = 'transfer',
}
