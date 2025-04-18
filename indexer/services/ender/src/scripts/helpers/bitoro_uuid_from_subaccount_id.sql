CREATE OR REPLACE FUNCTION bitoro_uuid_from_subaccount_id(subaccount_id jsonb) RETURNS uuid AS $$
/**
  Returns a UUID using the JSON.stringify format of an IndexerSubaccountId (https://github.com/bitoro-network/chain/blob/9ed26bd/proto/bitoroprotocol/indexer/protocol/v1/subaccount.proto#L15).

  (Note that no text should exist before the function declaration to ensure that exception line numbers are correct.)
*/
BEGIN
    RETURN bitoro_uuid_from_subaccount_id_parts(subaccount_id->>'owner', subaccount_id->>'number');
END;
$$ LANGUAGE plpgsql IMMUTABLE PARALLEL SAFE;
