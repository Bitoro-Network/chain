CREATE OR REPLACE FUNCTION bitoro_update_perpetual_v3_handler(event_data jsonb) RETURNS jsonb AS $$
/**
  Parameters:
    - event_data: The 'data' field of the IndexerTendermintEvent (https://github.com/bitoro-network/chain/blob/9ed26bd/proto/bitoroprotocol/indexer/indexer_manager/event.proto#L25)
        converted to JSON format. Conversion to JSON is expected to be done by JSON.stringify.
  Returns: JSON object containing fields:
    - perpetual_market: The updated perpetual market in perpetual-market-model format (https://github.com/bitoro-network/chain/blob/9ed26bd/indexer/packages/postgres/src/models/perpetual-market-model.ts).

  (Note that no text should exist before the function declaration to ensure that exception line numbers are correct.)
*/
DECLARE
    PPM_EXPONENT constant numeric = -6;
    FUNDING_RATE_FROM_PROTOCOL_IN_HOURS constant numeric = 8;
    perpetual_market_id bigint;
    perpetual_market_record perpetual_markets%ROWTYPE;
BEGIN
    perpetual_market_id = (event_data->'id')::bigint;
    perpetual_market_record."ticker" = event_data->>'ticker';
    perpetual_market_record."marketId" = (event_data->'marketId')::integer;
    perpetual_market_record."atomicResolution" = (event_data->'atomicResolution')::integer;
    perpetual_market_record."liquidityTierId" = (event_data->'liquidityTier')::integer;
    perpetual_market_record."marketType" = bitoro_protocol_market_type_to_perpetual_market_type(event_data->'marketType');
    perpetual_market_record."defaultFundingRate1H" = bitoro_trim_scale(
      power(10, PPM_EXPONENT) /
      FUNDING_RATE_FROM_PROTOCOL_IN_HOURS *
      (event_data->'defaultFunding8hrPpm')::numeric);
    UPDATE perpetual_markets
    SET
        "ticker" = perpetual_market_record."ticker",
        "marketId" = perpetual_market_record."marketId",
        "atomicResolution" = perpetual_market_record."atomicResolution",
        "liquidityTierId" = perpetual_market_record."liquidityTierId",
        "marketType" = perpetual_market_record."marketType",
        "defaultFundingRate1H" = perpetual_market_record."defaultFundingRate1H"
    WHERE "id" = perpetual_market_id
    RETURNING * INTO perpetual_market_record;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'Could not find perpetual market with corresponding id %', perpetual_market_id;
    END IF;

    RETURN jsonb_build_object(
            'perpetual_market',
            bitoro_to_jsonb(perpetual_market_record)
        );
END;
$$ LANGUAGE plpgsql;