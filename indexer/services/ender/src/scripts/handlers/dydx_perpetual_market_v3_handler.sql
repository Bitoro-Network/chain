CREATE OR REPLACE FUNCTION bitoro_perpetual_market_v3_handler(event_data jsonb) RETURNS jsonb AS $$
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
    perpetual_market_record perpetual_markets%ROWTYPE;
BEGIN
    perpetual_market_record."id" = (event_data->'id')::bigint;
    perpetual_market_record."clobPairId" = (event_data->'clobPairId')::bigint;
    perpetual_market_record."ticker" = event_data->>'ticker';
    perpetual_market_record."marketId" = (event_data->'marketId')::integer;
    perpetual_market_record."status" = bitoro_clob_pair_status_to_market_status(event_data->'status');
    perpetual_market_record."priceChange24H" = 0;
    perpetual_market_record."trades24H" = 0;
    perpetual_market_record."volume24H" = 0;
    perpetual_market_record."nextFundingRate" = 0;
    perpetual_market_record."openInterest"= 0;
    perpetual_market_record."quantumConversionExponent" = (event_data->'quantumConversionExponent')::integer;
    perpetual_market_record."atomicResolution" = (event_data->'atomicResolution')::integer;
    perpetual_market_record."subticksPerTick" = (event_data->'subticksPerTick')::integer;
    perpetual_market_record."stepBaseQuantums" = bitoro_from_jsonlib_long(event_data->'stepBaseQuantums');
    perpetual_market_record."liquidityTierId" = (event_data->'liquidityTier')::integer;
    perpetual_market_record."marketType" = bitoro_protocol_market_type_to_perpetual_market_type(event_data->'marketType');
    perpetual_market_record."baseOpenInterest" = 0;
    perpetual_market_record."defaultFundingRate1H" = bitoro_trim_scale(
      power(10, PPM_EXPONENT) /
      FUNDING_RATE_FROM_PROTOCOL_IN_HOURS *
      (event_data->'defaultFunding8hrPpm')::numeric);

    INSERT INTO perpetual_markets VALUES (perpetual_market_record.*) RETURNING * INTO perpetual_market_record;

    RETURN jsonb_build_object(
            'perpetual_market',
            bitoro_to_jsonb(perpetual_market_record)
        );
END;
$$ LANGUAGE plpgsql;
