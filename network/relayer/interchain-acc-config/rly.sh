#!/bin/bash

# Configure predefined mnemonic pharses
BINARY=rly
CHAIN_DIR=./network/data
RELAYER_DIR=./relayer
MNEMONIC_1="alley afraid soup fall idea toss can goose become valve initial strong forward bright dish figure check leopard decide warfare hub unusual join cart"
MNEMONIC_2="record gift you once hip style during joke field prize dust unique length more pencil transfer quit train device arrive energy sort steak upset"
MNEMONIC_3="license bulk van praise awake clown annual material direct angry unfair hollow spoil wage peasant tissue gold supreme boost slab parade artwork spawn couch"

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
$BINARY keys restore test-1 test-1 "$MNEMONIC_1" --home $CHAIN_DIR/$RELAYER_DIR
$BINARY keys restore test-2 test-2 "$MNEMONIC_2" --home $CHAIN_DIR/$RELAYER_DIR
$BINARY keys restore test-3 test-3 "$MNEMONIC_3" --home $CHAIN_DIR/$RELAYER_DIR

echo "Initializing light clients for all chains..."
$BINARY light init test-1 -f --home $CHAIN_DIR/$RELAYER_DIR
$BINARY light init test-2 -f --home $CHAIN_DIR/$RELAYER_DIR
$BINARY light init test-3 -f --home $CHAIN_DIR/$RELAYER_DIR

echo "Linking all chains..."
$BINARY tx link test1-account-test2 --home $CHAIN_DIR/$RELAYER_DIR
$BINARY tx link test1-account-test3 --home $CHAIN_DIR/$RELAYER_DIR

echo "Starting to listen relayer..."
$BINARY start test1-account-test2 --home $CHAIN_DIR/$RELAYER_DIR
$BINARY start test1-account-test3 --home $CHAIN_DIR/$RELAYER_DIR
