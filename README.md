<p align="center"><img src="https://bitoro.foundation/bitoro-token.svg?" width="256" /></p>

<h1 align="center">Bitoro Chain</h1>

<div align="center">
  <a href='https://github.com/bitoro-network/chain/blob/main/LICENSE'>
    <img src='https://img.shields.io/badge/License-AGPL_v3-blue.svg' alt='License' />
  </a>
</div>

The Bitoro software (the "Bitoro Chain") is a sovereign blockchain software built using Cosmos SDK and CometBFT. It powers a robust decentralized perpetual futures exchange with a high-performance orderbook and matching engine for a feature-rich, self-custodial perpetual trading experience on any market.

This repository contains the source code for the Cosmos SDK application responsible for running the chain and the associated indexer services. The indexer services are a read-only layer that indexes on-chain and off-chain data from a node and serves it to users in a more performant and user-friendly way.

# Getting Started

[Architecture Overview](https://bitoro.foundation/blog/chain-technical-architecture-overview)

[Indexer Deep Dive](https://bitoro.foundation/blog/chain-deep-dive-indexer)

[Front End Deep Dive](https://bitoro.foundation/blog/chain-deep-dive-front-end)

# Resources and Documentation

[Official Documentation](https://docs.bitoro.foundation/)

[Bitoro Academy](https://bitoro.foundation/crypto-learning#)

[Bitoro Blog](https://bitoro.foundation/blog#)

# Clients

[chain-clients](https://github.com/bitoro-network/chain-clients)

# Third-party Clients

[C++ Client](https://github.com/asnefedovv/bitoro-client-cpp)

[Python Client](https://github.com/kaloureyes3/chain-clients/tree/main/chain-client-py)

By clicking the above links to third-party clients, you will leave the Bitoro Trading Inc. ("Bitoro") GitHub repository and join repositories made available by third parties, which are independent from and unaffiliated with Bitoro. Bitoro is not responsible for any action taken or content on third-party repositories.

# Directory Structure

`audits` — Audit reports live here.

`docs` — Home for documentation pertaining to the entire repo.

`indexer` — Contains source code for indexer services. See its [README](https://github.com/bitoro-network/chain/blob/main/indexer/README.md) for developer documentation.

`proto` — Contains all protocol buffers used by protocol and indexer.

`protocol` — Contains source code for the Cosmos SDK app that runs the protocol. See its [README](https://github.com/bitoro-network/chain/blob/main/protocol/README.md) for developer documentation.

`chain-proto-js` — Scripts for publishing proto package to [npm](https://www.npmjs.com/package/@bitoroprotocol/chain-proto).

`chain-proto-py` — Scripts for publishing proto package to [PyPI](https://pypi.org/project/chain-proto/).

# Contributing

If you encounter a bug, please file an [issue](https://github.com/bitoro-network/chain/issues) to let us know. Alternatively, please feel free to contribute a bug fix directly. If you are planning to contribute changes that involve significant design choices, please open an issue for discussion instead.

# License and Terms

The Bitoro Chain is licensed under AGPLv3 and subject to the [v4 Terms of Use](https://bitoro.foundation/chain-terms). The full license can be found [here](https://github.com/bitoro-network/chain/blob/main/LICENSE). Bitoro does not deploy or run v4 software for public use, or operate or control any Bitoro Chain infrastructure. Bitoro products and services are not available to persons or entities who reside in, are located in, are incorporated in, or have registered offices in the United States or Canada, or Restricted Persons (as defined in the Bitoro [Terms of Use](https://bitoro.foundation/terms)).
