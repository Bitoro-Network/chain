CREATE OR REPLACE FUNCTION bitoro_uuid_from_funding_index_update_parts(block_height int, event_id bytea, perpetual_id bigint) RETURNS uuid AS $$
/**
  Returns a UUID using the parts of a funding index update.

  (Note that no text should exist before the function declaration to ensure that exception line numbers are correct.)
*/
BEGIN
    return bitoro_uuid(concat(block_height, '-', encode(event_id, 'hex'), '-', perpetual_id));
END;
$$ LANGUAGE plpgsql IMMUTABLE PARALLEL SAFE;
