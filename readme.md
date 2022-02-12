# Interamm Module
**interamm** is a Cosmos SDK module that allows for cross-chain interaction with IBC enabled DEX's and AMM's.

## Install
To install the latest version of this modules binary, a local Osmosis binary, a local Gaia (Cosmos Hub) binary and the Hermes relayer, execute the following command on your machine:

```
git clone https://github.com/schnetzlerjoe/interamm

make install
```

## Get Started

To get quick started, we have various commands that allow you to get running quickly. Running `make init` initializes a local Interamm , Osmosis and Gaia (Cosmos Hub) node. A Hermes relayer is also initialized.

The `make start-rly` command starts the hermes relayer. Run this command in another terminal.

`make start` starts up the local interamm, Osmosis, and Gaia nodes. You can then run commands cross-chain from the command line utilizing the interamm module.

```
make init

make start-rly

make start
```
