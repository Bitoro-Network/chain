import { logger } from '@bitoroprotocol-indexer/base';
import { AssetCreateEventV1, IndexerTendermintBlock, IndexerTendermintEvent } from '@bitoroprotocol-indexer/chain-protos';
import { dbHelpers, testMocks } from '@bitoroprotocol-indexer/postgres';
import { BitoroIndexerSubtypes } from '../../src/lib/types';
import {
  defaultAssetCreateEvent, defaultHeight, defaultTime, defaultTxHash,
} from '../helpers/constants';
import {
  createIndexerTendermintBlock,
  createIndexerTendermintEvent,
} from '../helpers/indexer-proto-helpers';
import { expectDidntLogError } from '../helpers/validator-helpers';
import { AssetValidator } from '../../src/validators/asset-validator';
import { createPostgresFunctions } from '../../src/helpers/postgres/postgres-functions';

describe('asset-validator', () => {
  beforeAll(async () => {
    await dbHelpers.migrate();
    await createPostgresFunctions();
  });

  beforeEach(async () => {
    await testMocks.seedData();
    jest.spyOn(logger, 'error');
  });

  afterEach(async () => {
    await dbHelpers.clearData();
    jest.clearAllMocks();
  });

  afterAll(async () => {
    await dbHelpers.teardown();
    jest.resetAllMocks();
  });

  describe('validate', () => {
    it('does not throw error on valid asset create event', () => {
      const validator: AssetValidator = new AssetValidator(
        defaultAssetCreateEvent,
        createBlock(defaultAssetCreateEvent),
        0,
      );

      validator.validate();
      expectDidntLogError();
    });
  });
});

function createBlock(
  assetCreateEvent: AssetCreateEventV1,
): IndexerTendermintBlock {
  const event: IndexerTendermintEvent = createIndexerTendermintEvent(
    BitoroIndexerSubtypes.ASSET,
    AssetCreateEventV1.encode(assetCreateEvent).finish(),
    0,
    0,
  );

  return createIndexerTendermintBlock(
    defaultHeight,
    defaultTime,
    [event],
    [defaultTxHash],
  );
}
