import { IndexerTendermintBlock, IndexerTendermintEvent } from '@bitoroprotocol-indexer/chain-protos';

export interface AnnotatedIndexerTendermintEvent extends IndexerTendermintEvent {
  data: string,
}

export interface AnnotatedIndexerTendermintBlock extends IndexerTendermintBlock {
  annotatedEvents: AnnotatedIndexerTendermintEvent[],
}

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
}
