import {
  ComplianceClient,
  getComplianceClient,
} from '@bitoroprotocol-indexer/compliance';
import { ComplianceProvider } from '@bitoroprotocol-indexer/postgres';

export interface ClientAndProvider {
  client: ComplianceClient,
  provider: ComplianceProvider,
}

export const complianceProvider: ClientAndProvider = {
  client: getComplianceClient(),
  provider: ComplianceProvider.ELLIPTIC,
};
