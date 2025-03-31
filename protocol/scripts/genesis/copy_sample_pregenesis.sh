#!/bin/bash

./scripts/genesis/prod_pregenesis.sh bitoroprotocold
cp /tmp/prod-chain/.bitoroprotocol/config/sorted_genesis.json ./scripts/genesis/sample_pregenesis.json
