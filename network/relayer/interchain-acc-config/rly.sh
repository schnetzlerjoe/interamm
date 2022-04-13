#!/bin/bash

# Configure predefined mnemonic pharses
BINARY=rly
CHAIN_DIR=./network/data
RELAYER_DIR=./relayer
MNEMONIC_1="alley afraid soup fall idea toss can goose become valve initial strong forward bright dish figure check leopard decide warfare hub unusual join cart"
MNEMONIC_2="record gift you once hip style during joke field prize dust unique length more pencil transfer quit train device arrive energy sort steak upset"

# Ensure rly is installed
if ! [ -x "$(command -v $BINARY)" ]; then
    echo "$BINARY is required to run this script..."
    echo "You can download at https://github.com/cosmos/relayer"
    exit 1
fi

echo "Initializing $BINARY..."
$BINARY config init --home $CHAIN_DIR/$RELAYER_DIR

echo "Adding configurations for all chains..."
$BINARY config add-chains $PWD/network/relayer/interchain-acc-config/chains --home $CHAIN_DIR/$RELAYER_DIR
$BINARY config add-paths $PWD/network/relayer/interchain-acc-config/paths --home $CHAIN_DIR/$RELAYER_DIR

echo "Restoring accounts..."
$BINARY keys restore interamm interamm "$MNEMONIC_1" --home $CHAIN_DIR/$RELAYER_DIR
$BINARY keys restore theta-testnet-001 theta-testnet-001 "$MNEMONIC_2" --home $CHAIN_DIR/$RELAYER_DIR

echo "Initializing light clients for all chains..."
$BINARY light init interamm -f --home $CHAIN_DIR/$RELAYER_DIR
$BINARY light init theta-testnet-001 -f --home $CHAIN_DIR/$RELAYER_DIR

echo "Linking all chains..."
$BINARY tx link interamm-account-gaia --home $CHAIN_DIR/$RELAYER_DIR

echo "Starting to listen relayer..."
$BINARY start interamm-account-gaia --home $CHAIN_DIR/$RELAYER_DIR
