import { createKafkaMessage } from '@bitoroprotocol-indexer/kafka';
import { OrderSide } from '@bitoroprotocol-indexer/postgres';
import {
  redisTestConstants,
  OrderbookLevelsCache,
  CanceledOrdersCache,
  CanceledOrderStatus,
} from '@bitoroprotocol-indexer/redis';
import { OffChainUpdateV1 } from '@bitoroprotocol-indexer/chain-protos';
import { KafkaMessage } from 'kafkajs';

import { redisClient } from '../../src/helpers/redis/redis-controller';
import { onMessage } from '../../src/lib/on-message';
import { BitoroRecordHeaderKeys } from '../../src/lib/types';
import { defaultKafkaHeaders } from './constants';

export async function handleInitialOrderPlace(
  orderPlace: redisTestConstants.OffChainUpdateOrderPlaceUpdateMessage,
): Promise<void> {
  const update: OffChainUpdateV1 = {
    ...orderPlace,
  };
  const message: KafkaMessage = createKafkaMessage(
    Buffer.from(Uint8Array.from(OffChainUpdateV1.encode(update).finish())),
  );
  message.headers = defaultKafkaHeaders;

  await onMessage(message);
}

export async function handleOrderUpdate(
  orderUpdate: redisTestConstants.OffChainUpdateOrderUpdateUpdateMessage,
): Promise<void> {
  const update: OffChainUpdateV1 = {
    ...orderUpdate,
  };
  const message: KafkaMessage = createKafkaMessage(
    Buffer.from(Uint8Array.from(OffChainUpdateV1.encode(update).finish())),
  );
  message.headers = defaultKafkaHeaders;

  await onMessage(message);
}

export async function expectOrderbookLevelCache(
  ticker: string,
  orderSide: OrderSide,
  humanPrice: string,
  size: string,
): Promise<void> {
  const level: string = await OrderbookLevelsCache.getOrderbookLevel(
    ticker,
    orderSide,
    humanPrice,
    redisClient,
  );
  expect(level).toEqual(size);
}

export function setTransactionHash(
  kafkaMessage: KafkaMessage,
  txHash: Buffer,
): KafkaMessage {
  const messageWithTxhash: KafkaMessage = {
    ...kafkaMessage,
  };
  if (kafkaMessage.headers === undefined) {
    messageWithTxhash.headers = {};
  }

  messageWithTxhash.headers![BitoroRecordHeaderKeys.TRANSACTION_HASH_KEY] = txHash;
  return messageWithTxhash;
}

export async function expectCanceledOrderStatus(
  orderId: string,
  canceledOrderStatus: CanceledOrderStatus,
) {
  expect(await CanceledOrdersCache.getOrderCanceledStatus(orderId, redisClient)).toEqual(
    canceledOrderStatus,
  );
}
