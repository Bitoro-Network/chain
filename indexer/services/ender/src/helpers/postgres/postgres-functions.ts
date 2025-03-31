import { readFileSync } from 'fs';
import path from 'path';

import { logger } from '@bitoroprotocol-indexer/base';
import { dbHelpers, storeHelpers } from '@bitoroprotocol-indexer/postgres';

export type PostgresFunction = {
  // The name of the script
  readonly name: string,
  // The contents of the script
  readonly script: string,
};

/**
 * Loads a named script from the specified path.
 *
 * @param name The name of the script.
 * @param scriptPath The path to the script.
 * @returns The created script object
 */
function newScript(name: string, scriptPath: string): PostgresFunction {
  const script: string = readFileSync(path.resolve(__dirname, scriptPath)).toString();
  return {
    name,
    script,
  };
}

const HANDLER_SCRIPTS: string[] = [
  'bitoro_asset_create_handler.sql',
  'bitoro_block_processor_ordered_handlers.sql',
  'bitoro_block_processor_unordered_handlers.sql',
  'bitoro_deleveraging_handler.sql',
  'bitoro_funding_handler.sql',
  'bitoro_liquidity_tier_handler.sql',
  'bitoro_market_create_handler.sql',
  'bitoro_market_modify_handler.sql',
  'bitoro_market_price_update_handler.sql',
  'bitoro_perpetual_market_v1_handler.sql',
  'bitoro_perpetual_market_v2_handler.sql',
  'bitoro_perpetual_market_v3_handler.sql',
  'bitoro_register_affiliate_handler.sql',
  'bitoro_stateful_order_handler.sql',
  'bitoro_subaccount_update_handler.sql',
  'bitoro_trading_rewards_handler.sql',
  'bitoro_transfer_handler.sql',
  'bitoro_update_clob_pair_handler.sql',
  'bitoro_update_perpetual_v1_handler.sql',
  'bitoro_update_perpetual_v2_handler.sql',
  'bitoro_update_perpetual_v3_handler.sql',
  'bitoro_vault_upsert_handler.sql',
];

const DB_SETUP_SCRIPTS: string[] = [
  'create_extension_pg_stat_statements.sql',
  'create_extension_uuid_ossp.sql',
];

const HELPER_SCRIPTS: string[] = [
  'bitoro_clob_pair_status_to_market_status.sql',
  'bitoro_create_initial_rows_for_tendermint_block.sql',
  'bitoro_create_tendermint_event.sql',
  'bitoro_create_transaction.sql',
  'bitoro_event_id_from_parts.sql',
  'bitoro_from_jsonlib_long.sql',
  'bitoro_from_protocol_order_side.sql',
  'bitoro_from_protocol_time_in_force.sql',
  'bitoro_from_serializable_int.sql',
  'bitoro_get_fee_from_liquidity.sql',
  'bitoro_get_order_status.sql',
  'bitoro_get_perpetual_market_for_clob_pair.sql',
  'bitoro_get_market_for_id.sql',
  'bitoro_get_total_filled_from_liquidity.sql',
  'bitoro_get_weighted_average.sql',
  'bitoro_liquidation_fill_handler_per_order.sql',
  'bitoro_order_fill_handler_per_order.sql',
  'bitoro_perpetual_position_and_order_side_matching.sql',
  'bitoro_process_trading_reward_event.sql',
  'bitoro_protocol_condition_type_to_order_type.sql',
  'bitoro_tendermint_event_to_transaction_index.sql',
  'bitoro_trim_scale.sql',
  'bitoro_update_perpetual_position_aggregate_fields.sql',
  'bitoro_uuid.sql',
  'bitoro_uuid_from_asset_position_parts.sql',
  'bitoro_uuid_from_fill_event_parts.sql',
  'bitoro_uuid_from_funding_index_update_parts.sql',
  'bitoro_uuid_from_oracle_price_parts.sql',
  'bitoro_uuid_from_order_id.sql',
  'bitoro_uuid_from_order_id_parts.sql',
  'bitoro_uuid_from_perpetual_position_parts.sql',
  'bitoro_uuid_from_subaccount_id.sql',
  'bitoro_uuid_from_subaccount_id_parts.sql',
  'bitoro_uuid_from_trading_rewards_parts.sql',
  'bitoro_uuid_from_transaction_parts.sql',
  'bitoro_uuid_from_transfer_parts.sql',
  'bitoro_protocol_market_type_to_perpetual_market_type.sql',
  'bitoro_protocol_vault_status_to_vault_status.sql',
];

const MAIN_SCRIPTS: string[] = [
  'bitoro_block_processor.sql',
];

const SCRIPTS: string[] = [
  ...HANDLER_SCRIPTS.map((script: string) => `handlers/${script}`),
  ...HELPER_SCRIPTS.map((script: string) => `helpers/${script}`),
  ...DB_SETUP_SCRIPTS.map((script: string) => `setup/${script}`),
  ...MAIN_SCRIPTS,
];

export async function createPostgresFunctions(): Promise<void> {
  await Promise.all([
    dbHelpers.createModelToJsonFunctions(),
    ...SCRIPTS.map((script: string) => storeHelpers.rawQuery(newScript(script, `../../scripts/${script}`).script, {})
      .catch((error: Error) => {
        logger.error({
          at: 'postgres-functions#createPostgresFunctions',
          message: `Failed to create or replace function contained in ${script}`,
          error,
        });
        throw error;
      }),
    ),
  ]);
}
