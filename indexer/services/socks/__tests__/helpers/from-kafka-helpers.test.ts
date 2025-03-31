import { getChannels, getMessageToForward, getMessagesToForward } from '../../src/helpers/from-kafka-helpers';
import { InvalidForwardMessageError, InvalidTopicError } from '../../src/lib/errors';
import {
  Channel,
  MessageToForward,
  WebsocketTopics,
} from '../../src/types';
import {
  btcTicker,
  candlesMessage,
  defaultAccNumber,
  defaultContents,
  defaultOwner,
  invalidChannel,
  invalidClobPairId,
  invalidTopic,
  marketsMessage,
  orderbookMessage,
  subaccountMessage,
  childSubaccountMessage,
  tradesMessage,
  defaultChildAccNumber,
  defaultTransferContents,
  defaultBlockHeightMessage,
} from '../constants';
import { KafkaMessage } from 'kafkajs';
import { createKafkaMessage } from './kafka';
import {
  BlockHeightMessage,
  CandleMessage,
  MarketMessage,
  OrderbookMessage,
  SubaccountMessage,
  TradeMessage,
} from '@bitoroprotocol-indexer/chain-protos';
import { CHAIN_MARKETS_ID } from '../../src/lib/constants';
import {
  dbHelpers,
  testMocks,
  perpetualMarketRefresher,
  CandleResolution,
  TransferSubaccountMessageContents,
  SubaccountMessageContents, TransferType,
} from '@bitoroprotocol-indexer/postgres';

describe('from-kafka-helpers', () => {
  describe('getChannel', () => {
    it.each([
      [WebsocketTopics.TO_WEBSOCKETS_CANDLES, [Channel.CHAIN_CANDLES]],
      [WebsocketTopics.TO_WEBSOCKETS_MARKETS, [Channel.CHAIN_MARKETS]],
      [WebsocketTopics.TO_WEBSOCKETS_ORDERBOOKS, [Channel.CHAIN_ORDERBOOK]],
      [
        WebsocketTopics.TO_WEBSOCKETS_SUBACCOUNTS,
        [Channel.CHAIN_ACCOUNTS, Channel.CHAIN_PARENT_ACCOUNTS],
      ],
      [WebsocketTopics.TO_WEBSOCKETS_TRADES, [Channel.CHAIN_TRADES]],
      [WebsocketTopics.TO_WEBSOCKETS_BLOCK_HEIGHT, [Channel.CHAIN_BLOCK_HEIGHT]],
    ])('gets correct channel for topic %s', (topic: WebsocketTopics, channels: Channel[]) => {
      expect(getChannels(topic)).toEqual(channels);
    });

    it('throws InvalidTopicError for invalid topic', () => {
      expect(() => { getChannels(invalidTopic); }).toThrow(new InvalidTopicError(invalidTopic));
    });
  });

  describe('getMessageToForward', () => {
    beforeAll(async () => {
      await dbHelpers.migrate();
      await testMocks.seedData();
      await perpetualMarketRefresher.updatePerpetualMarkets();
    });

    afterAll(async () => {
      await dbHelpers.clearData();
      await dbHelpers.teardown();
    });

    it('gets correct MessageToForward for subaccount message', () => {
      const message: KafkaMessage = createKafkaMessage(
        Buffer.from(Uint8Array.from(SubaccountMessage.encode(subaccountMessage).finish())),
      );
      const messageToForward: MessageToForward = getMessageToForward(
        Channel.CHAIN_ACCOUNTS,
        message,
      );

      expect(messageToForward.channel).toEqual(Channel.CHAIN_ACCOUNTS);
      expect(messageToForward.id).toEqual(`${defaultOwner}/${defaultAccNumber}`);
      expect(messageToForward.contents).toEqual(defaultContents);
    });

    it('gets correct MessageToForward for candles message', () => {
      const message: KafkaMessage = createKafkaMessage(
        Buffer.from(Uint8Array.from(CandleMessage.encode(candlesMessage).finish())),
      );
      const messageToForward: MessageToForward = getMessageToForward(
        Channel.CHAIN_CANDLES,
        message,
      );

      expect(messageToForward.channel).toEqual(Channel.CHAIN_CANDLES);
      expect(messageToForward.id).toEqual(`${btcTicker}/${CandleResolution.ONE_MINUTE}`);
      expect(messageToForward.contents).toEqual(defaultContents);
    });

    it('gets correct MessageToForward for market message', () => {
      const message: KafkaMessage = createKafkaMessage(
        Buffer.from(Uint8Array.from(MarketMessage.encode(marketsMessage).finish())),
      );
      const messageToForward: MessageToForward = getMessageToForward(Channel.CHAIN_MARKETS, message);

      expect(messageToForward.channel).toEqual(Channel.CHAIN_MARKETS);
      expect(messageToForward.id).toEqual(CHAIN_MARKETS_ID);
      expect(messageToForward.contents).toEqual(defaultContents);
    });

    it('gets correct MessageToForward for orderbook message', () => {
      const message: KafkaMessage = createKafkaMessage(
        Buffer.from(Uint8Array.from(OrderbookMessage.encode(orderbookMessage).finish())),
      );
      const messageToForward: MessageToForward = getMessageToForward(
        Channel.CHAIN_ORDERBOOK,
        message,
      );

      expect(messageToForward.channel).toEqual(Channel.CHAIN_ORDERBOOK);
      expect(messageToForward.id).toEqual(btcTicker);
      expect(messageToForward.contents).toEqual(defaultContents);
    });

    it('gets correct MessageToForward for trade message', () => {
      const message: KafkaMessage = createKafkaMessage(
        Buffer.from(Uint8Array.from(TradeMessage.encode(tradesMessage).finish())),
      );
      const messageToForward: MessageToForward = getMessageToForward(
        Channel.CHAIN_TRADES,
        message,
      );

      expect(messageToForward.channel).toEqual(Channel.CHAIN_TRADES);
      expect(messageToForward.id).toEqual(btcTicker);
      expect(messageToForward.contents).toEqual(defaultContents);
    });

    it('gets correct MessageToForward for subaccount message for parent subaccount channel', () => {
      const message: KafkaMessage = createKafkaMessage(
        Buffer.from(Uint8Array.from(SubaccountMessage.encode(childSubaccountMessage).finish())),
      );
      const messageToForward: MessageToForward = getMessageToForward(
        Channel.CHAIN_PARENT_ACCOUNTS,
        message,
      );

      expect(messageToForward.channel).toEqual(Channel.CHAIN_PARENT_ACCOUNTS);
      expect(messageToForward.id).toEqual(`${defaultOwner}/${defaultAccNumber}`);
      expect(messageToForward.contents).toEqual(defaultContents);
      expect(messageToForward.subaccountNumber).toBeDefined();
      expect(messageToForward.subaccountNumber).toEqual(defaultChildAccNumber);
    });

    it('gets correct MessageToForward for BlockHeight message', () => {
      const message: KafkaMessage = createKafkaMessage(
        Buffer.from(Uint8Array.from(BlockHeightMessage.encode(defaultBlockHeightMessage).finish())),
      );
      const messageToForward: MessageToForward = getMessagesToForward(
        WebsocketTopics.TO_WEBSOCKETS_BLOCK_HEIGHT,
        message,
      ).pop()!;
      expect(messageToForward.channel).toEqual(Channel.CHAIN_BLOCK_HEIGHT);
      expect(messageToForward.version).toEqual(defaultBlockHeightMessage.version);
      expect(messageToForward.contents).toEqual(
        {
          blockHeight: defaultBlockHeightMessage.blockHeight,
          time: defaultBlockHeightMessage.time,
        },
      );
    });

    it('filters out transfers between child subaccounts for parent subaccount channel', () => {
      const transferContents: SubaccountMessageContents = {
        transfers: {
          ...defaultTransferContents,
          sender: {
            address: defaultOwner,
            subaccountNumber: defaultAccNumber,
          },
          recipient: {
            address: defaultOwner,
            subaccountNumber: defaultChildAccNumber,
          },
          type: TransferType.TRANSFER_IN,
        },
      };
      const message: KafkaMessage = createKafkaMessage(
        Buffer.from(Uint8Array.from(SubaccountMessage.encode(
          {
            ...childSubaccountMessage,
            contents: JSON.stringify(transferContents),
          },
        ).finish())),
      );
      const messageToForward: MessageToForward = getMessageToForward(
        Channel.CHAIN_PARENT_ACCOUNTS,
        message,
      );

      expect(messageToForward.channel).toEqual(Channel.CHAIN_PARENT_ACCOUNTS);
      expect(messageToForward.id).toEqual(`${defaultOwner}/${defaultAccNumber}`);
      expect(messageToForward.contents).toEqual({});
      expect(messageToForward.subaccountNumber).toBeDefined();
      expect(messageToForward.subaccountNumber).toEqual(defaultChildAccNumber);
    });

    it.each([
      [
        'transfer between other parent/child subaccount',
        {
          ...defaultTransferContents,
          sender: {
            address: defaultOwner,
            subaccountNumber: defaultAccNumber + 1,
          },
          recipient: {
            address: defaultOwner,
            subaccountNumber: defaultChildAccNumber,
          },
        },
      ],
      [
        'deposit',
        {
          ...defaultTransferContents,
          sender: {
            address: defaultOwner,
            subaccountNumber: undefined,
          },
          recipient: {
            address: defaultOwner,
            subaccountNumber: defaultChildAccNumber,
          },
          type: TransferType.DEPOSIT,
        },
      ],
      [
        'withdraw',
        {
          ...defaultTransferContents,
          sender: {
            address: defaultOwner,
            subaccountNumber: defaultChildAccNumber,
          },
          recipient: {
            address: defaultOwner,
            subaccountNumber: undefined,
          },
          type: TransferType.WITHDRAWAL,
        },
      ],
    ])('does not filter out transfer message for (%s)', (
      _name: string,
      transfer: TransferSubaccountMessageContents,
    ) => {
      const transferContents: SubaccountMessageContents = {
        transfers: transfer,
      };
      const message: KafkaMessage = createKafkaMessage(
        Buffer.from(Uint8Array.from(SubaccountMessage.encode(
          {
            ...childSubaccountMessage,
            contents: JSON.stringify(transferContents),
          },
        ).finish())),
      );
      const messageToForward: MessageToForward = getMessageToForward(
        Channel.CHAIN_PARENT_ACCOUNTS,
        message,
      );

      expect(messageToForward.channel).toEqual(Channel.CHAIN_PARENT_ACCOUNTS);
      expect(messageToForward.id).toEqual(`${defaultOwner}/${defaultAccNumber}`);
      expect(messageToForward.contents).toEqual(transferContents);
      expect(messageToForward.subaccountNumber).toBeDefined();
      expect(messageToForward.subaccountNumber).toEqual(defaultChildAccNumber);
    });

    it('throws InvalidForwardMessageError for empty message', () => {
      const message: KafkaMessage = createKafkaMessage(null);

      expect(() => { getMessageToForward(Channel.CHAIN_ACCOUNTS, message); }).toThrow(
        new InvalidForwardMessageError('Got empty kafka message'),
      );
    });

    it('throws InvalidForwardMessageError for invalid channel', () => {
      const message: KafkaMessage = createKafkaMessage(Buffer.from(''));

      expect(() => { getMessageToForward((invalidChannel as Channel), message); }).toThrow(
        new InvalidForwardMessageError(`Unknown channel: ${invalidChannel}`),
      );
    });

    it('throw InvalidForwardMessageError for invalid clobPairId on candles message', () => {
      const invalidCandlesMessage: CandleMessage = {
        ...candlesMessage,
        clobPairId: invalidClobPairId,
      };
      const message: KafkaMessage = createKafkaMessage(
        Buffer.from(Uint8Array.from(CandleMessage.encode(invalidCandlesMessage).finish())),
      );

      expect(() => { getMessageToForward(Channel.CHAIN_CANDLES, message); }).toThrow(
        new InvalidForwardMessageError(`Invalid clob pair id: ${invalidClobPairId}`),
      );
    });

    it('throw InvalidForwardMessageError for invalid clobPairId on orderbook message', () => {
      const invalidOrderbookMessage: OrderbookMessage = {
        ...orderbookMessage,
        clobPairId: invalidClobPairId,
      };
      const message: KafkaMessage = createKafkaMessage(
        Buffer.from(Uint8Array.from(OrderbookMessage.encode(invalidOrderbookMessage).finish())),
      );

      expect(() => { getMessageToForward(Channel.CHAIN_ORDERBOOK, message); }).toThrow(
        new InvalidForwardMessageError(`Invalid clob pair id: ${invalidClobPairId}`),
      );
    });

    it('throw InvalidForwardMessageError for invalid clobPairId on trade message', () => {
      const invalidTradeMessage: TradeMessage = {
        ...tradesMessage,
        clobPairId: invalidClobPairId,
      };
      const message: KafkaMessage = createKafkaMessage(
        Buffer.from(Uint8Array.from(TradeMessage.encode(invalidTradeMessage).finish())),
      );

      expect(() => { getMessageToForward(Channel.CHAIN_TRADES, message); }).toThrow(
        new InvalidForwardMessageError(`Invalid clob pair id: ${invalidClobPairId}`),
      );
    });
  });
});
