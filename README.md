# synapsecommon-go

[![Tests](https://img.shields.io/github/workflow/status/synapsecns/synapsecommon-go/Test/master?event=push&label=tests)](https://github.com/synapsecns/synapsecommon-go/actions/workflows/test.yml)
[![Coverags](https://img.shields.io/coveralls/github/synapsecns/synapsecommon-go/master?label=coverage)](https://coveralls.io/github/synapsecns/synapsecommon-go?branch=master)


Utility functions, ABI wrappers, and more, most or all of which are used in various synapsecns Go codebases.

Made public for the greater good of all Go developers.

## Installation

Supported Go versions:

- 1.16.x
- 1.17.x
- Likely 1.18, once it hits RC

Install (Go 1.16.x): `go get -u github.com/synapsecns/synapsecommon-go`

Install (Go 1.17.x and higher): `go get -d -u github.com/synapsecns/synapsecommon-go`

## Contents

- [abis/erc20](./abis/erc20) - Output of abigen for the IERC20 interface. Should be wrapped.
- [addrutils](./addrutils) - Utilities for dealing with go-ethereum's common.Address type.
- [bigfloat](./bigfloat) - Wrapper functions around common tasks when dealing with `big.Float` objects
- [bigint](./bigint) - Wrapper functions around common tasks when dealing with `big.Int` objects
- [env](./env) - Easily fetch values set in the local environment, with a fallback. 
- [unitconv](./uniconv) - Convert arbitrary units of Ethereum currency units of Wei or Ether. 