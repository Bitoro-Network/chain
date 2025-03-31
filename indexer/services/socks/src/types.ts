import { IncomingMessage as IncomingMessageHttp } from 'http';

import express from 'express';
import WebSocket from 'ws';

export enum IncomingMessageType {
  SUBSCRIBE = 'subscribe',
  PING = 'ping',
  UNSUBSCRIBE = 'unsubscribe',
}

export enum RequestMethod {
  GET = 'get',
  POST = 'post',
  DELETE = 'delete',
}

export enum Channel {
  CHAIN_ORDERBOOK = 'chain_orderbook',
  CHAIN_ACCOUNTS = 'chain_subaccounts',
  CHAIN_TRADES = 'chain_trades',
  CHAIN_MARKETS = 'chain_markets',
  CHAIN_CANDLES = 'chain_candles',
  CHAIN_PARENT_ACCOUNTS = 'chain_parent_subaccounts',
  CHAIN_BLOCK_HEIGHT = 'chain_block_height',
}

export const ALL_CHANNELS = Object.values(Channel);

export interface IncomingMessage extends IncomingMessageHttp {
  type: IncomingMessageType,
}

export interface SubscribeMessage extends IncomingMessage {
  channel: Channel,
  id?: string,
  batched?: boolean,
  timestamp?: string,
  includeOffsets?: boolean,
}

export interface UnsubscribeMessage extends IncomingMessage {
  channel: Channel,
  id?: string,
  timestamp?: string,
}

export interface PingMessage extends IncomingMessage {
  id?: number,
}

export enum OutgoingMessageType {
  ERROR = 'error',
  CONNECTED = 'connected',
  SUBSCRIBED = 'subscribed',
  UNSUBSCRIBED = 'unsubscribed',
  CHANNEL_DATA = 'channel_data',
  CHANNEL_BATCH_DATA = 'channel_batch_data',
  PONG = 'pong',
}

export interface OutgoingMessage {
  type: OutgoingMessageType,
  connection_id: string,
  message_id: number,
}

export interface ErrorMessage extends OutgoingMessage {
  message: string,
  channel?: string,
  id?: string,
}

export interface SubscribedMessage extends OutgoingMessage {
  channel: Channel,
  id?: string,
  // eslint-disable-next-line  @typescript-eslint/no-explicit-any
  contents: any,
}

export interface UnsubscribedMessage extends OutgoingMessage {
  channel: Channel,
  id?: string,
}

export interface ChannelDataMessage extends OutgoingMessage {
  // eslint-disable-next-line  @typescript-eslint/no-explicit-any
  contents: any,
  channel: Channel,
  id?: string,
  version: string,
  subaccountNumber?: number,
}

export interface ChannelBatchDataMessage extends OutgoingMessage {
  // eslint-disable-next-line  @typescript-eslint/no-explicit-any
  contents: any[],
  channel: Channel,
  id?: string,
  version: string,
  subaccountNumber?: number,
}

export interface ConnectedMessage extends OutgoingMessage {}

export interface PongMessage extends OutgoingMessage {
  id?: number,
}

export interface Subscription {
  channel: Channel,
  id: string,
  batched?: boolean,
}

export interface SubscriptionInfo {
  connectionId: string,
  pending: boolean,
  pendingMessages: MessageToForward[],
  batched?: boolean,
}

export interface Connection {
  ws: WebSocket,
  messageId: number,
  heartbeat?: NodeJS.Timeout,
  disconnect?: NodeJS.Timeout,
  countryCode?: string,
}

export interface MessageToForward {
  channel: Channel,
  id: string,
  // eslint-disable-next-line  @typescript-eslint/no-explicit-any
  contents: any,
  version: string,
  subaccountNumber?: number,
}

export interface ResponseWithBody extends express.Response {
  body: unknown,
}

export enum WebsocketTopics {
  TO_WEBSOCKETS_ORDERBOOKS = 'to-websockets-orderbooks',
  TO_WEBSOCKETS_SUBACCOUNTS = 'to-websockets-subaccounts',
  TO_WEBSOCKETS_TRADES = 'to-websockets-trades',
  TO_WEBSOCKETS_MARKETS = 'to-websockets-markets',
  TO_WEBSOCKETS_CANDLES = 'to-websockets-candles',
  TO_WEBSOCKETS_BLOCK_HEIGHT = 'to-websockets-block-height',
}

export enum WebsocketEvents {
  CLOSE = 'close',
  CONNECTION = 'connection',
  ERROR = 'error',
  LISTENING = 'listening',
  MESSAGE = 'message',
  PONG = 'pong',
  PING = 'ping',
}
