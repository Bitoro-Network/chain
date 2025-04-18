import { logger, ParseMessageError } from '@bitoroprotocol-indexer/base';
import {
  OrderSide,
  PerpetualMarketFromDatabase,
  PositionSide,
  protocolTranslations,
} from '@bitoroprotocol-indexer/postgres';
import {
  IndexerTendermintEvent,
  IndexerTendermintEvent_BlockEvent,
  OrderFillEventV1,
  MarketEventV1,
  SubaccountUpdateEventV1,
  TransferEventV1,
  IndexerOrder,
  StatefulOrderEventV1,
  FundingEventV1,
  AssetCreateEventV1,
  PerpetualMarketCreateEventV1,
  PerpetualMarketCreateEventV2,
  PerpetualMarketCreateEventV3,
  LiquidityTierUpsertEventV1,
  LiquidityTierUpsertEventV2,
  UpdatePerpetualEventV1,
  UpdatePerpetualEventV2,
  UpdatePerpetualEventV3,
  UpdateClobPairEventV1,
  SubaccountMessage,
  DeleveragingEventV1,
  TradingRewardsEventV1,
  RegisterAffiliateEventV1,
  UpsertVaultEventV1,
} from '@bitoroprotocol-indexer/chain-protos';
import Big from 'big.js';
import _ from 'lodash';
import { DateTime } from 'luxon';

import {
  AnnotatedSubaccountMessage,
  BitoroIndexerSubtypes,
  EventProtoWithTypeAndVersion,
} from './types';

export function indexerTendermintEventToTransactionIndex(
  event: IndexerTendermintEvent,
): number {
  if (event.transactionIndex !== undefined) {
    return event.transactionIndex;
  } else if (event.blockEvent !== undefined) {
    switch (event.blockEvent) {
      case IndexerTendermintEvent_BlockEvent.BLOCK_EVENT_BEGIN_BLOCK:
        return -2;
      case IndexerTendermintEvent_BlockEvent.BLOCK_EVENT_END_BLOCK:
        return -1;
      default:
        throw new ParseMessageError(`Received V4 event with invalid block event type: ${event.blockEvent}`);
    }
  }

  throw new ParseMessageError(
    'Either transactionIndex or blockEvent must be defined in IndexerTendermintEvent',
  );
}

export function convertToSubaccountMessage(
  annotatedMessage: AnnotatedSubaccountMessage,
): SubaccountMessage {
  const subaccountMessage: SubaccountMessage = _.omit(
    annotatedMessage,
    ['orderId', 'isFill', 'subaccountMessageContents'],
  );
  return subaccountMessage;
}

export function dateToDateTime(
  protoTime: Date,
): DateTime {
  return DateTime.fromJSDate(protoTime);
}

/**
 * Determines the event subtype and parses the IndexerTendermintEvent
 * to the correct EventProto and returns it all in an object.
 * @param blockEventIndex - the index of the event in the block.
 * @param event - the event.
 * @returns
 */
export function indexerTendermintEventToEventProtoWithType(
  blockEventIndex: number,
  event: IndexerTendermintEvent,
): EventProtoWithTypeAndVersion | undefined {
  const eventDataBinary: Uint8Array = event.dataBytes;
  // set the default version to 1
  const version: number = event.version === 0 ? 1 : event.version;
  switch (event.subtype) {
    case (BitoroIndexerSubtypes.ORDER_FILL.toString()): {
      return {
        type: BitoroIndexerSubtypes.ORDER_FILL,
        eventProto: OrderFillEventV1.decode(eventDataBinary),
        indexerTendermintEvent: event,
        version,
        blockEventIndex,
      };
    }
    case (BitoroIndexerSubtypes.SUBACCOUNT_UPDATE.toString()): {
      return {
        type: BitoroIndexerSubtypes.SUBACCOUNT_UPDATE,
        eventProto: SubaccountUpdateEventV1.decode(eventDataBinary),
        indexerTendermintEvent: event,
        version,
        blockEventIndex,
      };
    }
    case (BitoroIndexerSubtypes.TRANSFER.toString()): {
      return {
        type: BitoroIndexerSubtypes.TRANSFER,
        eventProto: TransferEventV1.decode(eventDataBinary),
        indexerTendermintEvent: event,
        version,
        blockEventIndex,
      };
    }
    case (BitoroIndexerSubtypes.MARKET.toString()): {
      return {
        type: BitoroIndexerSubtypes.MARKET,
        eventProto: MarketEventV1.decode(eventDataBinary),
        indexerTendermintEvent: event,
        version,
        blockEventIndex,
      };
    }
    case (BitoroIndexerSubtypes.STATEFUL_ORDER.toString()): {
      return {
        type: BitoroIndexerSubtypes.STATEFUL_ORDER,
        eventProto: StatefulOrderEventV1.decode(eventDataBinary),
        indexerTendermintEvent: event,
        version,
        blockEventIndex,
      };
    }
    case (BitoroIndexerSubtypes.FUNDING.toString()): {
      return {
        type: BitoroIndexerSubtypes.FUNDING,
        eventProto: FundingEventV1.decode(eventDataBinary),
        indexerTendermintEvent: event,
        version,
        blockEventIndex,
      };
    }
    case (BitoroIndexerSubtypes.ASSET.toString()): {
      return {
        type: BitoroIndexerSubtypes.ASSET,
        eventProto: AssetCreateEventV1.decode(eventDataBinary),
        indexerTendermintEvent: event,
        version,
        blockEventIndex,
      };
    }
    case (BitoroIndexerSubtypes.PERPETUAL_MARKET.toString()): {
      if (version === 1) {
        return {
          type: BitoroIndexerSubtypes.PERPETUAL_MARKET,
          eventProto: PerpetualMarketCreateEventV1.decode(eventDataBinary),
          indexerTendermintEvent: event,
          version,
          blockEventIndex,
        };
      } else if (version === 2) {
        return {
          type: BitoroIndexerSubtypes.PERPETUAL_MARKET,
          eventProto: PerpetualMarketCreateEventV2.decode(eventDataBinary),
          indexerTendermintEvent: event,
          version,
          blockEventIndex,
        };
      } else if (version === 3) {
        return {
          type: BitoroIndexerSubtypes.PERPETUAL_MARKET,
          eventProto: PerpetualMarketCreateEventV3.decode(eventDataBinary),
          indexerTendermintEvent: event,
          version,
          blockEventIndex,
        };
      } else {
        const message: string = `Invalid version for perpetual market event: ${version}`;
        logger.error({
          at: 'helpers#indexerTendermintEventToEventWithType',
          message,
        });
        return undefined;
      }
    }
    case (BitoroIndexerSubtypes.LIQUIDITY_TIER.toString()): {
      if (version === 1) {
        return {
          type: BitoroIndexerSubtypes.LIQUIDITY_TIER,
          eventProto: LiquidityTierUpsertEventV1.decode(eventDataBinary),
          indexerTendermintEvent: event,
          version,
          blockEventIndex,
        };
      }
      return {
        type: BitoroIndexerSubtypes.LIQUIDITY_TIER,
        eventProto: LiquidityTierUpsertEventV2.decode(eventDataBinary),
        indexerTendermintEvent: event,
        version,
        blockEventIndex,
      };
    }
    case (BitoroIndexerSubtypes.UPDATE_PERPETUAL.toString()): {
      if (version === 1) {
        return {
          type: BitoroIndexerSubtypes.UPDATE_PERPETUAL,
          eventProto: UpdatePerpetualEventV1.decode(eventDataBinary),
          indexerTendermintEvent: event,
          version,
          blockEventIndex,
        };
      } else if (version === 2) {
        return {
          type: BitoroIndexerSubtypes.UPDATE_PERPETUAL,
          eventProto: UpdatePerpetualEventV2.decode(eventDataBinary),
          indexerTendermintEvent: event,
          version,
          blockEventIndex,
        };
      } else if (version === 3) {
        return {
          type: BitoroIndexerSubtypes.UPDATE_PERPETUAL,
          eventProto: UpdatePerpetualEventV3.decode(eventDataBinary),
          indexerTendermintEvent: event,
          version,
          blockEventIndex,
        };
      } else {
        const message: string = `Invalid version for update perpetual event: ${version}`;
        logger.error({
          at: 'helpers#indexerTendermintEventToEventWithType',
          message,
        });
        return undefined;
      }
    }
    case (BitoroIndexerSubtypes.UPDATE_CLOB_PAIR.toString()): {
      return {
        type: BitoroIndexerSubtypes.UPDATE_CLOB_PAIR,
        eventProto: UpdateClobPairEventV1.decode(eventDataBinary),
        indexerTendermintEvent: event,
        version,
        blockEventIndex,
      };
    }
    case (BitoroIndexerSubtypes.DELEVERAGING.toString()): {
      return {
        type: BitoroIndexerSubtypes.DELEVERAGING,
        eventProto: DeleveragingEventV1.decode(eventDataBinary),
        indexerTendermintEvent: event,
        version,
        blockEventIndex,
      };
    }
    case (BitoroIndexerSubtypes.TRADING_REWARD.toString()): {
      return {
        type: BitoroIndexerSubtypes.TRADING_REWARD,
        eventProto: TradingRewardsEventV1.decode(eventDataBinary),
        indexerTendermintEvent: event,
        version,
        blockEventIndex,
      };
    }
    case (BitoroIndexerSubtypes.REGISTER_AFFILIATE.toString()): {
      return {
        type: BitoroIndexerSubtypes.REGISTER_AFFILIATE,
        eventProto: RegisterAffiliateEventV1.decode(eventDataBinary),
        indexerTendermintEvent: event,
        version,
        blockEventIndex,
      };
    }
    case (BitoroIndexerSubtypes.UPSERT_VAULT.toString()): {
      return {
        type: BitoroIndexerSubtypes.UPSERT_VAULT,
        eventProto: UpsertVaultEventV1.decode(eventDataBinary),
        indexerTendermintEvent: event,
        version,
        blockEventIndex,
      };
    }
    default: {
      const message: string = `Unable to parse event subtype: ${event.subtype}`;
      logger.error({
        at: 'helpers#indexerTendermintEventToEventWithType',
        message,
      });
      return undefined;
    }
  }
}

/**
 * Returns the size of an order in human readable form.
 * @param order
 * @param perpetualMarket
 */
export function getSize(
  order: IndexerOrder,
  perpetualMarket: PerpetualMarketFromDatabase,
): string {
  return protocolTranslations.quantumsToHumanFixedString(
    order.quantums.toString(10),
    perpetualMarket.atomicResolution,
  );
}

/**
 * Returns the price of an order in human readable form.
 *
 * @param order
 * @param perpetualMarket
 */
export function getPrice(
  order: IndexerOrder,
  perpetualMarket: PerpetualMarketFromDatabase,
): string {
  return protocolTranslations.subticksToPrice(
    order.subticks.toString(10),
    perpetualMarket,
  );
}

/**
 * Returns the trigger price of an order in human readable form.
 *
 * @param order
 * @param perpetualMarket
 * @returns
 */
export function getTriggerPrice(
  order: IndexerOrder,
  perpetualMarket: PerpetualMarketFromDatabase,
): string {
  return protocolTranslations.subticksToPrice(
    order.conditionalOrderTriggerSubticks.toString(10),
    perpetualMarket,
  );
}

/**
 * Returns the weighted average between two prices
 * @param firstPrice
 * @param firstWeight
 * @param secondPrice
 * @param secondWeight
 * @returns
 */
export function getWeightedAverage(
  firstPrice: string,
  firstWeight: string,
  secondPrice: string,
  secondWeight: string,
): string {
  return Big(firstPrice).times(firstWeight).plus(
    Big(secondPrice).times(secondWeight),
  )
    .div(Big(firstWeight).plus(secondWeight))
    .toFixed();
}

/**
 * Returns true if perpetualPostionSide is LONG and orderSide is BUY or
 * if perpetualPostionSide is SHORT and orderSide is SELL
 * @param perpetualPositionSide
 * @param orderSide
 */
export function perpetualPositionAndOrderSideMatching(
  perpetualPositionSide: PositionSide,
  orderSide: OrderSide,
): boolean {
  return (perpetualPositionSide === PositionSide.LONG && orderSide === OrderSide.BUY) ||
    (perpetualPositionSide === PositionSide.SHORT && orderSide === OrderSide.SELL);
}
