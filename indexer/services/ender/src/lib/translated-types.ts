import { Liquidity } from '@bitoroprotocol-indexer/postgres';
import {
  IndexerAssetPosition,
  IndexerOrder,
  IndexerPerpetualPosition,
  IndexerSubaccountId, LiquidationOrderV1,
} from '@bitoroprotocol-indexer/chain-protos';
import { Long } from '@bitoroprotocol-indexer/chain-protos/build/codegen/helpers';

/* Canonical object types for handling onchain messages from the protocol. */

export interface SubaccountUpdate {
  subaccountId?: IndexerSubaccountId,
  updatedPerpetualPositions: IndexerPerpetualPosition[],
  updatedAssetPositions: IndexerAssetPosition[],
}

export interface OrderFillWithLiquidity {
  makerOrder?: IndexerOrder,
  order?: IndexerOrder,
  liquidationOrder?: LiquidationOrderV1,
  /** Fill amount in base quantums. */
  fillAmount: Long,
  /** Maker fee in USDC quantums. */
  makerFee: Long,
  /**
   * Taker fee in USDC quantums. If the taker order is a liquidation, then this
   * represents the special liquidation fee, not the standard taker fee.
   */
  takerFee: Long,
  /** Total filled of the maker order in base quantums. */
  totalFilledMaker: Long,
  /** Total filled of the taker order in base quantums. */
  totalFilledTaker: Long,
  /** Liquidity of the order in the match to process in the handler. */
  liquidity: Liquidity,
  /** Affiliate rev share in USDC quantums. */
  affiliateRevShare: Long,
}
