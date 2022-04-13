#!/bin/bash

BINARY=interammd
CHAIN_DIR=./network/data
CHAINID_1=interamm
CHAINID_2=theta-testnet-001
VAL_MNEMONIC_1="clock post desk civil pottery foster expand merit dash seminar song memory figure uniform spice circle try happy obvious trash crime hybrid hood cushion"
DEMO_MNEMONIC_1="banner spread envelope side kite person disagree path silver will brother under couch edit food venture squirrel civil budget number acquire point work mass"
DEMO_MNEMONIC_2="veteran try aware erosion drink dance decade comic dawn museum release episode original list ability owner size tuition surface ceiling depth seminar capable only"
RLY_MNEMONIC_1="alley afraid soup fall idea toss can goose become valve initial strong forward bright dish figure check leopard decide warfare hub unusual join cart"
RLY_MNEMONIC_2="record gift you once hip style during joke field prize dust unique length more pencil transfer quit train device arrive energy sort steak upset"
P2PPORT_1=16656
RPCPORT_1=16657
RESTPORT_1=1316
ROSETTA_1=8080

# Stop interammd if it is already running 
if pgrep -x "$BINARY" >/dev/null; then
    echo "Terminating $BINARY..."
    killall $BINARY
fi

echo "Removing previous data..."
rm -rf $CHAIN_DIR/$CHAINID_1 &> /dev/null

# Add directories for all chains, exit if an error occurs
if ! mkdir -p $CHAIN_DIR/$CHAINID_1 2>/dev/null; then
    echo "Failed to create chain folder. Aborting..."
    exit 1
fi

cd $HOME/interamm
echo "Initializing $CHAINID_1..."
$BINARY init test --home $CHAIN_DIR/$CHAINID_1 --chain-id=$CHAINID_1

echo "Adding accounts..."
echo $VAL_MNEMONIC_1 | $BINARY keys add val --home $CHAIN_DIR/$CHAINID_1 --recover --keyring-backend=test
echo $DEMO_MNEMONIC_1 | $BINARY keys add demowallet --home $CHAIN_DIR/$CHAINID_1 --recover --keyring-backend=test
echo $DEMO_MNEMONIC_2 | gaiad keys add demowallet --home $CHAIN_DIR/$CHAINID_2 --recover --keyring-backend=test
echo $RLY_MNEMONIC_1 | $BINARY keys add rly --home $CHAIN_DIR/$CHAINID_1 --recover --keyring-backend=test 
echo $RLY_MNEMONIC_2 | gaiad keys add rly --home $CHAIN_DIR/$CHAINID_2 --recover --keyring-backend=test

$BINARY add-genesis-account $($BINARY --home $CHAIN_DIR/$CHAINID_1 keys show val --keyring-backend test -a) 100000000000stake  --home $CHAIN_DIR/$CHAINID_1
$BINARY add-genesis-account $($BINARY --home $CHAIN_DIR/$CHAINID_1 keys show demowallet --keyring-backend test -a) 100000000000stake  --home $CHAIN_DIR/$CHAINID_1
$BINARY add-genesis-account $($BINARY --home $CHAIN_DIR/$CHAINID_1 keys show rly --keyring-backend test -a) 100000000000stake  --home $CHAIN_DIR/$CHAINID_1

echo "Creating and collecting gentx..."
$BINARY gentx val 7000000000stake --home $CHAIN_DIR/$CHAINID_1 --chain-id $CHAINID_1 --keyring-backend test
$BINARY collect-gentxs --home $CHAIN_DIR/$CHAINID_1

echo "Changing defaults and ports in app.toml and config.toml files..."
sed -i -e 's#"tcp://0.0.0.0:26656"#"tcp://0.0.0.0:'"$P2PPORT_1"'"#g' $CHAIN_DIR/$CHAINID_1/config/config.toml
sed -i -e 's#"tcp://127.0.0.1:26657"#"tcp://0.0.0.0:'"$RPCPORT_1"'"#g' $CHAIN_DIR/$CHAINID_1/config/config.toml
sed -i -e 's/timeout_commit = "5s"/timeout_commit = "1s"/g' $CHAIN_DIR/$CHAINID_1/config/config.toml
sed -i -e 's/timeout_propose = "3s"/timeout_propose = "1s"/g' $CHAIN_DIR/$CHAINID_1/config/config.toml
sed -i -e 's/index_all_keys = false/index_all_keys = true/g' $CHAIN_DIR/$CHAINID_1/config/config.toml
sed -i -e 's/enable = false/enable = true/g' $CHAIN_DIR/$CHAINID_1/config/app.toml
sed -i -e 's/swagger = false/swagger = true/g' $CHAIN_DIR/$CHAINID_1/config/app.toml
sed -i -e 's#"tcp://0.0.0.0:1317"#"tcp://0.0.0.0:'"$RESTPORT_1"'"#g' $CHAIN_DIR/$CHAINID_1/config/app.toml
sed -i -e 's#":8080"#":'"$ROSETTA_1"'"#g' $CHAIN_DIR/$CHAINID_1/config/app.toml
