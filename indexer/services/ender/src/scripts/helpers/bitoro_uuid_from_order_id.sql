CREATE OR REPLACE FUNCTION bitoro_uuid_from_order_id(order_id jsonb) RETURNS uuid AS $$
/**
  Returns a UUID using the JSON.stringify format of an IndexerOrderId (https://github.com/bitoro-network/chain/blob/9ed26bd/proto/bitoroprotocol/indexer/protocol/v1/clob.proto#L15).

  (Note that no text should exist before the function declaration to ensure that exception line numbers are correct.)
*/
BEGIN
    return bitoro_uuid_from_order_id_parts(
        bitoro_uuid_from_subaccount_id(order_id->'subaccountId'),
        order_id->>'clientId',
        order_id->>'clobPairId',
        order_id->>'orderFlags');
END;
$$ LANGUAGE plpgsql IMMUTABLE PARALLEL SAFE;
