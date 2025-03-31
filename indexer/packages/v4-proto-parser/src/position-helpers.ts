import {
  IndexerAssetPosition,
  IndexerPerpetualPosition,
} from '@bitoroprotocol-indexer/chain-protos';

import {
  bytesToBigInt,
} from './bytes-helpers';

export function getPositionIsLong(
  position: IndexerAssetPosition | IndexerPerpetualPosition,
): boolean {
  return bytesToBigInt(position.quantums) > 0;
}
