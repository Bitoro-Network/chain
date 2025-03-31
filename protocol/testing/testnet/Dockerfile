FROM bitoroprotocol-base

RUN apk add --no-cache bash jq aws-cli
RUN go install cosmossdk.io/tools/cosmovisor/cmd/cosmovisor@v1.5.0

COPY ./testing/version/. /bitoroprotocol/
COPY ./testing/testnet/. /bitoroprotocol/

ENV HOME /bitoroprotocol
WORKDIR $HOME

RUN /bitoroprotocol/testnet.sh

ENTRYPOINT ["/bitoroprotocol/start.sh"]
