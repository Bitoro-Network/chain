import {
  IndexerTendermintBlock,
  UpsertVaultEventV1,
  VaultStatus,
} from '@bitoroprotocol-indexer/chain-protos';
import {
  dbHelpers,
  testMocks,
  testConstants,
  VaultFromDatabase,
  VaultTable,
  VaultStatus as IndexerVaultStatus,
} from '@bitoroprotocol-indexer/postgres';
import { KafkaMessage } from 'kafkajs';
import { createKafkaMessage } from '@bitoroprotocol-indexer/kafka';
import { onMessage } from '../../src/lib/on-message';
import { BitoroIndexerSubtypes } from '../../src/lib/types';
import { createIndexerTendermintBlock, createIndexerTendermintEvent } from '../helpers/indexer-proto-helpers';
import {
  defaultHeight,
  defaultPreviousHeight,
  defaultTime,
  defaultTxHash,
} from '../helpers/constants';
import { updateBlockCache } from '../../src/caches/block-cache';
import { createPostgresFunctions } from '../../src/helpers/postgres/postgres-functions';

describe('upsertVaultHandler', () => {
  beforeAll(async () => {
    await dbHelpers.migrate();
    await dbHelpers.clearData();
    await createPostgresFunctions();
  });

  beforeEach(async () => {
    await testMocks.seedData();
    updateBlockCache(defaultPreviousHeight);
  });

  afterEach(async () => {
    await dbHelpers.clearData();
    jest.clearAllMocks();
  });

  afterAll(async () => {
    await dbHelpers.teardown();
    jest.resetAllMocks();
  });

  it('should upsert new vaults in single block', async () => {
    const events: UpsertVaultEventV1[] = [
      {
        address: testConstants.defaultVaultAddress,
        clobPairId: 0,
        status: VaultStatus.VAULT_STATUS_QUOTING,
      }, {
        address: testConstants.defaultAddress,
        clobPairId: 1,
        status: VaultStatus.VAULT_STATUS_STAND_BY,
      },
    ];
    const block: IndexerTendermintBlock = createBlockFromEvents(
      defaultHeight,
      ...events,
    );
    const binaryBlock: Uint8Array = IndexerTendermintBlock.encode(block).finish();
    const kafkaMessage: KafkaMessage = createKafkaMessage(Buffer.from(binaryBlock));

    await onMessage(kafkaMessage);

    const vaults: VaultFromDatabase[] = await VaultTable.findAll(
      {},
      [],
      {},
    );
    expect(vaults).toHaveLength(2);
    expect(vaults[0]).toEqual({
      address: testConstants.defaultVaultAddress,
      clobPairId: '0',
      status: IndexerVaultStatus.QUOTING,
      createdAt: block.time?.toISOString(),
      updatedAt: block.time?.toISOString(),
    });
    expect(vaults[1]).toEqual({
      address: testConstants.defaultAddress,
      clobPairId: '1',
      status: IndexerVaultStatus.STAND_BY,
      createdAt: block.time?.toISOString(),
      updatedAt: block.time?.toISOString(),
    });
  });

  it('should upsert an existing vault', async () => {
    await VaultTable.create(testConstants.defaultVault);

    const events: UpsertVaultEventV1[] = [
      {
        address: testConstants.defaultVaultAddress,
        clobPairId: 0,
        status: VaultStatus.VAULT_STATUS_CLOSE_ONLY,
      },
    ];
    const block: IndexerTendermintBlock = createBlockFromEvents(
      defaultHeight,
      ...events,
    );
    const binaryBlock: Uint8Array = IndexerTendermintBlock.encode(block).finish();
    const kafkaMessage: KafkaMessage = createKafkaMessage(Buffer.from(binaryBlock));

    await onMessage(kafkaMessage);

    const vaults: VaultFromDatabase[] = await VaultTable.findAll(
      {},
      [],
      {},
    );
    expect(vaults).toHaveLength(1);
    expect(vaults[0]).toEqual({
      address: testConstants.defaultVault.address,
      clobPairId: testConstants.defaultVault.clobPairId,
      status: IndexerVaultStatus.CLOSE_ONLY,
      createdAt: testConstants.defaultVault.createdAt,
      updatedAt: block.time?.toISOString(),
    });
  });
});

function createBlockFromEvents(
  height: number,
  ...events: UpsertVaultEventV1[]
): IndexerTendermintBlock {
  const transactionIndex = 0;
  let eventIndex = 0;

  return createIndexerTendermintBlock(
    height,
    defaultTime,
    events.map((event) => {
      const indexerEvent = createIndexerTendermintEvent(
        BitoroIndexerSubtypes.UPSERT_VAULT,
        UpsertVaultEventV1.encode(event).finish(),
        transactionIndex,
        eventIndex,
      );
      eventIndex += 1;
      return indexerEvent;
    }),
    [defaultTxHash],
  );
}
