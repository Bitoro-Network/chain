import { PerpetualMarketFromDatabase } from '@bitoroprotocol-indexer/postgres';
import { PerpetualMarketCreateEventV1 } from '@bitoroprotocol-indexer/chain-protos';

export function expectPerpetualMarketMatchesEvent(
  perpetual: PerpetualMarketCreateEventV1,
  perpetualMarket: PerpetualMarketFromDatabase,
) {
  expect(perpetualMarket).toEqual(expect.objectContaining({
    id: perpetual.id.toString(),
    clobPairId: perpetual.clobPairId.toString(),
    ticker: perpetual.ticker,
    marketId: perpetual.marketId,
    quantumConversionExponent: perpetual.quantumConversionExponent,
    atomicResolution: perpetual.atomicResolution,
    subticksPerTick: perpetual.subticksPerTick,
    stepBaseQuantums: Number(perpetual.stepBaseQuantums),
    liquidityTierId: perpetual.liquidityTier,
  }));
}
