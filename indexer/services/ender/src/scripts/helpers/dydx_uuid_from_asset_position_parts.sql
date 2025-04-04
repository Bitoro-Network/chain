CREATE OR REPLACE FUNCTION bitoro_uuid_from_asset_position_parts(subaccount_uuid uuid, asset_id text) RETURNS uuid AS $$
/**
  Returns a UUID using the parts of an asset position.

  (Note that no text should exist before the function declaration to ensure that exception line numbers are correct.)
*/
BEGIN
    return bitoro_uuid(concat(subaccount_uuid, '-', asset_id));
END;
$$ LANGUAGE plpgsql IMMUTABLE PARALLEL SAFE;
