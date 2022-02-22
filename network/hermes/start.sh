#!/bin/bash

# Load shell variables
. ./network/hermes/variables.sh

echo "Hermes Relayer Version Check......"
$HERMES_BINARY version

# Start the hermes relayer in multi-paths mode
echo "Starting hermes relayer..."
$HERMES_BINARY -c $CONFIG_DIR start
