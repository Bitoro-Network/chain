FROM bitoroprotocol-base

COPY ./testing/testnet-dev/dev.sh /bitoroprotocol/dev.sh
COPY ./testing/genesis.sh /bitoroprotocol/genesis.sh
COPY ./testing/start.sh /bitoroprotocol/start.sh
COPY ./daemons/pricefeed/client/constants/testdata /bitoroprotocol/exchange_config
COPY ./testing/delaymsg_config /bitoroprotocol/delaymsg_config

RUN go install cosmossdk.io/tools/cosmovisor/cmd/cosmovisor@v1.5.0

RUN /bitoroprotocol/dev.sh

ENV DAEMON_NAME=bitoroprotocold
ENTRYPOINT ["cosmovisor", "run"]
