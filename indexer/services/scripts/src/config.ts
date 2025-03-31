import { baseConfigSchema, parseSchema } from '@bitoroprotocol-indexer/base';
import { kafkaConfigSchema } from '@bitoroprotocol-indexer/kafka';
import { postgresConfigSchema } from '@bitoroprotocol-indexer/postgres';

export const configSchema = {
  ...baseConfigSchema,
  ...postgresConfigSchema,
  ...kafkaConfigSchema,
};

export default parseSchema(configSchema);
