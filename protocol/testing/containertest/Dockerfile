FROM bitoroprotocol-base

COPY ./testing/containertest/bitoroprotocold_preupgrade* /bin/
COPY ./testing/containertest/containertest.sh /bitoroprotocol/containertest.sh
COPY ./testing/containertest/preupgrade_entrypoint.sh /bitoroprotocol/preupgrade_entrypoint.sh
COPY ./testing/containertest/preupgrade_genesis.json /bitoroprotocol/preupgrade_genesis.json
COPY ./testing/genesis.sh /bitoroprotocol/genesis.sh
COPY ./daemons/pricefeed/client/constants/testdata /bitoroprotocol/exchange_config
COPY ./testing/delaymsg_config /bitoroprotocol/delaymsg_config
COPY ./testing/version/. /bitoroprotocol/

RUN go install cosmossdk.io/tools/cosmovisor/cmd/cosmovisor@v1.5.0

RUN /bitoroprotocol/containertest.sh

COPY ./testing/containertest/config/pricefeed_exchange_config.toml /bitoroprotocol/chain/.alice/config/pricefeed_exchange_config.toml
COPY ./testing/containertest/config/pricefeed_exchange_config.toml /bitoroprotocol/chain/.bob/config/pricefeed_exchange_config.toml
COPY ./testing/containertest/config/pricefeed_exchange_config.toml /bitoroprotocol/chain/.carl/config/pricefeed_exchange_config.toml
COPY ./testing/containertest/config/pricefeed_exchange_config.toml /bitoroprotocol/chain/.dave/config/pricefeed_exchange_config.toml

ENTRYPOINT ["bitoroprotocold"]
