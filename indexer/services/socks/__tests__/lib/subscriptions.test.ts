import WebSocket from 'ws';
import { Channel, OutgoingMessageType } from '../../src/types';
import { Subscriptions } from '../../src/lib/subscription';
import { sendMessage, sendMessageString } from '../../src/helpers/wss';
import { RateLimiter } from '../../src/lib/rate-limit';
import {
  dbHelpers,
  testMocks,
  perpetualMarketRefresher,
  CandleResolution,
  MAX_PARENT_SUBACCOUNTS,
  blockHeightRefresher,
} from '@bitoroprotocol-indexer/postgres';
import { btcTicker, invalidChannel, invalidTicker } from '../constants';
import { axiosRequest } from '../../src/lib/axios';
import { AxiosSafeServerError, makeAxiosSafeServerError } from '@bitoroprotocol-indexer/base';
import { BlockedError } from '../../src/lib/errors';

jest.mock('ws');
jest.mock('../../src/helpers/wss');
jest.mock('../../src/lib/axios');
jest.mock('@bitoroprotocol-indexer/compliance');

describe('Subscriptions', () => {
  let subscriptions: Subscriptions;
  let mockWs: WebSocket;
  let sendMessageMock: jest.Mock;
  let sendMessageStringMock: jest.Mock;
  let rateLimiterSpy: jest.SpyInstance;
  let axiosRequestMock: jest.Mock;

  const connectionId: string = 'connectionId';
  const initialMsgId: number = 1;
  const defaultId: string = 'id';
  const mockSubaccountId: string = 'address/0';
  const invalidCandleResolution: string = 'candleResolution';
  const validIds: Record<Channel, string> = {
    [Channel.CHAIN_ACCOUNTS]: mockSubaccountId,
    [Channel.CHAIN_CANDLES]: `${btcTicker}/${CandleResolution.ONE_DAY}`,
    [Channel.CHAIN_MARKETS]: defaultId,
    [Channel.CHAIN_ORDERBOOK]: btcTicker,
    [Channel.CHAIN_TRADES]: btcTicker,
    [Channel.CHAIN_PARENT_ACCOUNTS]: mockSubaccountId,
    [Channel.CHAIN_BLOCK_HEIGHT]: defaultId,
  };
  const invalidIdsMap:
  Record<Exclude<Channel, Channel.CHAIN_MARKETS | Channel.CHAIN_BLOCK_HEIGHT>, string[]> = {
    [Channel.CHAIN_ACCOUNTS]: [invalidTicker],
    [Channel.CHAIN_CANDLES]: [
      `${invalidTicker}/${CandleResolution.ONE_DAY}`,
      `${btcTicker}/${invalidCandleResolution}`,
      btcTicker,
    ],
    [Channel.CHAIN_ORDERBOOK]: [invalidTicker],
    [Channel.CHAIN_TRADES]: [invalidTicker],
    [Channel.CHAIN_PARENT_ACCOUNTS]: [`address/${MAX_PARENT_SUBACCOUNTS}`],
  };
  const initialResponseUrlPatterns: Record<Channel, string[] | undefined> = {
    [Channel.CHAIN_ACCOUNTS]: [
      '/v4/addresses/.+/subaccountNumber/.+',
      '/v4/orders?.+subaccountNumber.+OPEN,UNTRIGGERED,BEST_EFFORT_OPENED',
      '/v4/orders?.+subaccountNumber.+BEST_EFFORT_CANCELED.+goodTilBlockAfter=[0-9]+',
    ],
    [Channel.CHAIN_CANDLES]: ['/v4/candles/perpetualMarkets/.+?resolution=.+'],
    [Channel.CHAIN_MARKETS]: ['/v4/perpetualMarkets'],
    [Channel.CHAIN_ORDERBOOK]: ['/v4/orderbooks/perpetualMarket/.+'],
    [Channel.CHAIN_TRADES]: ['/v4/trades/perpetualMarket/.+'],
    [Channel.CHAIN_PARENT_ACCOUNTS]: [
      '/v4/addresses/.+/parentSubaccountNumber/.+',
      '/v4/orders/parentSubaccountNumber?.+parentSubaccountNumber.+OPEN,UNTRIGGERED,BEST_EFFORT_OPENED',
      '/v4/orders/parentSubaccountNumber?.+parentSubaccountNumber.+BEST_EFFORT_CANCELED.+goodTilBlockAfter=[0-9]+',
    ],
    [Channel.CHAIN_BLOCK_HEIGHT]: ['v4/height'],
  };
  const initialMessage: Object = ['a', 'b'];
  const country: string = 'AR';

  beforeAll(async () => {
    await dbHelpers.migrate();
    await testMocks.seedData();
    await Promise.all([
      perpetualMarketRefresher.updatePerpetualMarkets(),
      blockHeightRefresher.updateBlockHeight(),
    ]);
  });

  afterAll(async () => {
    await dbHelpers.clearData();
    await dbHelpers.teardown();
  });

  beforeEach(() => {
    (WebSocket as unknown as jest.Mock).mockClear();
    subscriptions = new Subscriptions();
    subscriptions.start(jest.fn());
    mockWs = new WebSocket(null);
    sendMessageMock = (sendMessage as jest.Mock);
    sendMessageStringMock = (sendMessageString as jest.Mock);
    rateLimiterSpy = jest.spyOn(RateLimiter.prototype, 'rateLimit');
    axiosRequestMock = (axiosRequest as jest.Mock);
    axiosRequestMock.mockClear();
    axiosRequestMock.mockImplementation(() => (JSON.stringify(initialMessage)));
  });

  describe('subscribe', () => {
    it.each([
      [Channel.CHAIN_ACCOUNTS, validIds[Channel.CHAIN_ACCOUNTS]],
      [Channel.CHAIN_CANDLES, validIds[Channel.CHAIN_CANDLES]],
      [Channel.CHAIN_MARKETS, validIds[Channel.CHAIN_MARKETS]],
      [Channel.CHAIN_ORDERBOOK, validIds[Channel.CHAIN_ORDERBOOK]],
      [Channel.CHAIN_TRADES, validIds[Channel.CHAIN_TRADES]],
      [Channel.CHAIN_PARENT_ACCOUNTS, validIds[Channel.CHAIN_PARENT_ACCOUNTS]],
      [Channel.CHAIN_BLOCK_HEIGHT, validIds[Channel.CHAIN_BLOCK_HEIGHT]],
    ])('handles valid subscription request to channel %s', async (
      channel: Channel,
      id: string,
    ) => {
      await subscriptions.subscribe(
        mockWs,
        channel,
        connectionId,
        initialMsgId,
        id,
        false,
        country,
      );

      expect(sendMessageStringMock).toHaveBeenCalledTimes(1);
      expect(sendMessageStringMock).toHaveBeenCalledWith(
        mockWs,
        connectionId,
        expect.stringContaining(OutgoingMessageType.SUBSCRIBED),
      );
      expect(subscriptions.subscriptions[channel][id])
        .toContainEqual(expect.objectContaining({ connectionId }));
      expect(subscriptions.subscriptionLists[connectionId]).toHaveLength(1);
      expect(subscriptions.subscriptionLists[connectionId])
        .toContainEqual(expect.objectContaining({ channel, id }));

      const urlPatterns: string[] | undefined = initialResponseUrlPatterns[channel];
      if (urlPatterns !== undefined) {
        for (const urlPattern of urlPatterns) {
          expect(axiosRequestMock).toHaveBeenCalledWith(expect.objectContaining({
            url: expect.stringMatching(RegExp(urlPattern)),
            headers: {
              'cf-ipcountry': country,
            },
          }));
        }
      } else {
        expect(axiosRequestMock).not.toHaveBeenCalled();
      }
    });

    it.each([
      [Channel.CHAIN_ACCOUNTS, invalidIdsMap[Channel.CHAIN_ACCOUNTS]],
      [Channel.CHAIN_CANDLES, invalidIdsMap[Channel.CHAIN_CANDLES]],
      [Channel.CHAIN_ORDERBOOK, invalidIdsMap[Channel.CHAIN_ORDERBOOK]],
      [Channel.CHAIN_TRADES, invalidIdsMap[Channel.CHAIN_TRADES]],
      [Channel.CHAIN_PARENT_ACCOUNTS, invalidIdsMap[Channel.CHAIN_PARENT_ACCOUNTS]],
    ])('sends error message if invalid subscription request to channel %s', async (
      channel: Channel,
      invalidIds: string[],
    ) => {
      for (const id of invalidIds) {
        await subscriptions.subscribe(
          mockWs,
          channel,
          connectionId,
          initialMsgId,
          id,
          false,
        );

        expect(sendMessageMock).toHaveBeenCalledTimes(1);
        expect(sendMessageMock).toHaveBeenCalledWith(
          mockWs,
          connectionId,
          expect.objectContaining({
            connection_id: connectionId,
            type: 'error',
            message: `Invalid subscription id for channel: (${channel}-${id})`,
          }),
        );
        expect(subscriptions.subscriptions[channel]).toBeUndefined();

        sendMessageMock.mockClear();
      }
    });

    it('throws error if channel is invalid', async () => {
      await expect(
        async () => {
          await subscriptions.subscribe(
            mockWs,
            (invalidChannel as Channel),
            connectionId,
            initialMsgId,
            defaultId,
            false,
          );
        },
      ).rejects.toEqual(new Error(`Invalid channel: ${invalidChannel}`));
    });

    it('sends error message if rate limit exceeded', async () => {
      rateLimiterSpy.mockImplementation(() => 1);
      await subscriptions.subscribe(
        mockWs,
        Channel.CHAIN_ACCOUNTS,
        connectionId,
        initialMsgId,
        mockSubaccountId,
        false,
      );

      expect(sendMessageMock).toHaveBeenCalledTimes(1);
      expect(sendMessageMock).toHaveBeenCalledWith(
        mockWs,
        connectionId,
        expect.objectContaining({
          message: expect.stringContaining('Too many subscribe attempts'),
        }));
      expect(subscriptions.subscriptions[Channel.CHAIN_ACCOUNTS]).toBeUndefined();
      expect(subscriptions.subscriptionLists[connectionId]).toBeUndefined();
    });

    it('sends error message if initial message request fails', async () => {
      axiosRequestMock.mockImplementation(() => { throw Error(); });
      await subscriptions.subscribe(
        mockWs,
        Channel.CHAIN_ACCOUNTS,
        connectionId,
        initialMsgId,
        mockSubaccountId,
        false,
      );

      expect(sendMessageMock).toHaveBeenCalledTimes(1);
      expect(sendMessageMock).toHaveBeenCalledWith(
        mockWs,
        connectionId,
        expect.objectContaining({
          connection_id: connectionId,
          type: 'error',
          message: expect.stringContaining(
            `Internal error, could not fetch data for subscription: ${Channel.CHAIN_ACCOUNTS}`,
          ),
          channel: Channel.CHAIN_ACCOUNTS,
          id: mockSubaccountId,
        }));
      expect(subscriptions.subscriptions[Channel.CHAIN_ACCOUNTS]).toBeUndefined();
      expect(subscriptions.subscriptionLists[connectionId]).toBeUndefined();
    });

    it('sends blocked error message if initial message request fails with 403', async () => {
      const expectedError: BlockedError = new BlockedError();
      axiosRequestMock.mockImplementation(
        () => {
          throw new AxiosSafeServerError({
            data: {},
            status: 403,
            statusText: '',
          }, {});
        },
      );
      await subscriptions.subscribe(
        mockWs,
        Channel.CHAIN_ACCOUNTS,
        connectionId,
        initialMsgId,
        mockSubaccountId,
        false,
        country,
      );

      expect(sendMessageMock).toHaveBeenCalledTimes(1);
      expect(sendMessageMock).toHaveBeenCalledWith(
        mockWs,
        connectionId,
        expect.objectContaining({
          connection_id: connectionId,
          type: 'error',
          message: expectedError.message,
          channel: Channel.CHAIN_ACCOUNTS,
          id: mockSubaccountId,
        }));
      expect(subscriptions.subscriptions[Channel.CHAIN_ACCOUNTS]).toBeUndefined();
      expect(subscriptions.subscriptionLists[connectionId]).toBeUndefined();
    });

    it('sends empty contents if initial message request fails with 404 for accounts', async () => {
      axiosRequestMock.mockImplementation(() => {
        return Promise.reject(makeAxiosSafeServerError(404, '', ''));
      });
      await subscriptions.subscribe(
        mockWs,
        Channel.CHAIN_ACCOUNTS,
        connectionId,
        initialMsgId,
        mockSubaccountId,
        false,
        country,
      );

      expect(sendMessageStringMock).toHaveBeenCalledTimes(1);
      expect(sendMessageStringMock).toHaveBeenCalledWith(
        mockWs,
        connectionId,
        expect.stringContaining(OutgoingMessageType.SUBSCRIBED),
      );
      expect(subscriptions.subscriptions[Channel.CHAIN_ACCOUNTS][mockSubaccountId])
        .toContainEqual(expect.objectContaining({ connectionId }));
      expect(subscriptions.subscriptionLists[connectionId]).toHaveLength(1);
      expect(subscriptions.subscriptionLists[connectionId])
        .toContainEqual(
          expect.objectContaining({ channel: Channel.CHAIN_ACCOUNTS, id: mockSubaccountId }),
        );
    });
  });

  describe('unsubscribe', () => {
    it.each([
      [Channel.CHAIN_ACCOUNTS, validIds[Channel.CHAIN_ACCOUNTS]],
      [Channel.CHAIN_CANDLES, validIds[Channel.CHAIN_CANDLES]],
      [Channel.CHAIN_MARKETS, validIds[Channel.CHAIN_MARKETS]],
      [Channel.CHAIN_ORDERBOOK, validIds[Channel.CHAIN_ORDERBOOK]],
      [Channel.CHAIN_TRADES, validIds[Channel.CHAIN_TRADES]],
      [Channel.CHAIN_BLOCK_HEIGHT, validIds[Channel.CHAIN_BLOCK_HEIGHT]],
    ])('handles valid unsubscription request to channel %s', async (
      channel: Channel,
      id: string,
    ) => {
      await subscriptions.subscribe(
        mockWs,
        channel,
        connectionId,
        initialMsgId,
        id,
        false,
        country,
      );
      subscriptions.unsubscribe(
        connectionId,
        channel,
        id,
      );

      expect(subscriptions.subscriptions[channel][id]).toHaveLength(0);
      expect(subscriptions.subscriptionLists[connectionId]).toHaveLength(0);
    });

    it('is no-op if connection is not subscribed to channel and id', async () => {
      await subscriptions.subscribe(
        mockWs,
        Channel.CHAIN_ACCOUNTS,
        connectionId,
        initialMsgId,
        mockSubaccountId,
        false,
        country,
      );
      subscriptions.unsubscribe(
        connectionId,
        Channel.CHAIN_CANDLES,
        defaultId,
      );

      expect(subscriptions.subscriptions[Channel.CHAIN_ACCOUNTS][mockSubaccountId]).toHaveLength(1);
      expect(subscriptions.subscriptions[Channel.CHAIN_CANDLES]).toBeUndefined();
      expect(subscriptions.subscriptionLists[connectionId]).toHaveLength(1);
    });
  });

  describe('remove', () => {
    it('removes connection id from all subscriptions', async () => {
      await Promise.all(Object.values(Channel).map((channel: Channel): Promise<void> => {
        return subscriptions.subscribe(
          mockWs,
          channel,
          connectionId,
          initialMsgId,
          validIds[channel],
          false,
        );
      }));

      for (const channel of Object.values(Channel)) {
        expect(subscriptions.subscriptions[channel][validIds[channel]]).toHaveLength(1);
        expect(subscriptions.subscriptions[channel][validIds[channel]]).toContainEqual(
          expect.objectContaining({ connectionId }),
        );
      }
      expect(subscriptions.subscriptionLists[connectionId]).toHaveLength(
        Object.values(Channel).length,
      );

      subscriptions.remove(connectionId);

      for (const channel of Object.values(Channel)) {
        expect(subscriptions.subscriptions[channel][validIds[channel]]).toHaveLength(0);
      }
      expect(subscriptions.subscriptionLists[connectionId]).toBeUndefined();
    });
  });
});
