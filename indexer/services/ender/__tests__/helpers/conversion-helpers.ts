import { bigIntToBytes } from '@bitoroprotocol-indexer/chain-proto-parser';

export function intToSerializedInt(int: number): Uint8Array {
  return bigIntToBytes(BigInt(int));
}
