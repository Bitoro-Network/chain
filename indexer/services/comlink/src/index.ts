import {
  logger,
  wrapBackgroundTask,
  startBugsnag,
} from '@bitoroprotocol-indexer/base';
import { perpetualMarketRefresher, liquidityTierRefresher } from '@bitoroprotocol-indexer/postgres';

import { startVaultStartPnlCache } from './caches/vault-start-pnl';
import config from './config';
import IndexV4 from './controllers/api/index-v4';
import { connect as connectToRedis } from './helpers/redis/redis-controller';
import Server from './request-helpers/server';

process.on('SIGTERM', () => {
  logger.info({
    at: 'index#SIGTERM',
    message: 'Received SIGTERM, shutting down',
  });
  process.exit(0);
});

function startServer() {
  const app = Server(IndexV4);
  const port: number = config.PORT;
  const server = app.listen(port, () => {
    logger.info({
      at: 'index#startServer',
      message: `Api server is listening on ${port}`,
    });
  });

  server.keepAliveTimeout = config.KEEP_ALIVE_MS;
  server.headersTimeout = config.HEADERS_TIMEOUT_MS;
}

async function start() {
  startBugsnag();

  // Initialize PerpetualMarkets cache
  await Promise.all([
    perpetualMarketRefresher.updatePerpetualMarkets(),
    liquidityTierRefresher.updateLiquidityTiers(),
  ]);
  wrapBackgroundTask(perpetualMarketRefresher.start(), true, 'startUpdatePerpetualMarkets');
  wrapBackgroundTask(liquidityTierRefresher.start(), true, 'startUpdateLiquidityTiers');
  // Initialize cache for vault start PnL
  await startVaultStartPnlCache();

  await connectToRedis();
  logger.info({
    at: 'index#start',
    message: `Connected to redis at ${config.REDIS_URL}`,
  });

  startServer();
}

wrapBackgroundTask(start(), true, 'main');
