import { IndexerTendermintEvent, UpsertVaultEventV1 } from '@bitoroprotocol-indexer/chain-protos';

import { Handler } from '../handlers/handler';
import { UpsertVaultHandler } from '../handlers/upsert-vault-handler';
import { Validator } from './validator';

export class UpsertVaultValidator extends Validator<UpsertVaultEventV1> {
  public validate(): void {
    if (this.event.address === '') {
      return this.logAndThrowParseMessageError(
        'UpsertVaultEvent address is not populated',
      );
    }
  }

  public createHandlers(
    indexerTendermintEvent: IndexerTendermintEvent,
    txId: number,
    _: string,
  ): Handler<UpsertVaultEventV1>[] {
    return [
      new UpsertVaultHandler(
        this.block,
        this.blockEventIndex,
        indexerTendermintEvent,
        txId,
        this.event,
      ),
    ];
  }
}
