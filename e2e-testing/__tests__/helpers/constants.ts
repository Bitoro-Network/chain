import {
  IPlaceOrder, Order_Side, Order_TimeInForce, OrderFlags,
} from '@bitoroprotocol/chain-client-js';
import Long from 'long';

import { OrderDetails } from './types';

export const BITORO_LOCAL_ADDRESS = 'btoro199tqg4wdlnu4qjlxchpd7seg454937hjrknju4';
export const BITORO_LOCAL_MNEMONIC = 'merge panther lobster crazy road hollow amused security before critic about cliff exhibit cause coyote talent happy where lion river tobacco option coconut small';
export const BITORO_LOCAL_ADDRESS_2 = 'btoro10fx7sy6ywd5senxae9dwytf8jxek3t2gcen2vs';
export const BITORO_LOCAL_MNEMONIC_2 = 'color habit donor nurse dinosaur stable wonder process post perfect raven gold census inside worth inquiry mammal panic olive toss shadow strong name drum';

export const MNEMONIC_TO_ADDRESS: Record<string, string> = {
  [BITORO_LOCAL_MNEMONIC]: BITORO_LOCAL_ADDRESS,
  [BITORO_LOCAL_MNEMONIC_2]: BITORO_LOCAL_ADDRESS_2,
};

export const ADDRESS_TO_MNEMONIC: Record<string, string> = {
  [BITORO_LOCAL_ADDRESS]: BITORO_LOCAL_MNEMONIC,
  [BITORO_LOCAL_ADDRESS_2]: BITORO_LOCAL_MNEMONIC_2,
};

export const PERPETUAL_PAIR_BTC_USD: number = 0;
export const quantums: Long = new Long(1_000_000_000);
export const subticks: Long = new Long(1_000_000_000);

export const defaultOrder: IPlaceOrder = {
  clientId: 0,
  orderFlags: OrderFlags.SHORT_TERM,
  clobPairId: PERPETUAL_PAIR_BTC_USD,
  side: Order_Side.SIDE_BUY,
  quantums,
  subticks,
  timeInForce: Order_TimeInForce.TIME_IN_FORCE_UNSPECIFIED,
  reduceOnly: false,
  clientMetadata: 0,
};

export const orderDetails: OrderDetails[] = [
  {
    mnemonic: BITORO_LOCAL_MNEMONIC,
    timeInForce: 0,
    orderFlags: 64,
    side: 1,
    clobPairId: PERPETUAL_PAIR_BTC_USD,
    quantums: 10000000,
    subticks: 5000000000,
  },
  {
    mnemonic: BITORO_LOCAL_MNEMONIC_2,
    timeInForce: 0,
    orderFlags: 64,
    side: 2,
    clobPairId: PERPETUAL_PAIR_BTC_USD,
    quantums: 5000000,
    subticks: 5000000000,
  },
];
