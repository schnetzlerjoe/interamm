#!/bin/bash

cd $HOME
yes | rm -r osmosis
git clone https://github.com/osmosis-labs/osmosis
cd osmosis
git checkout v6.2.0
make install
sudo cp $GOPATH/bin/osmosisd /usr/local/bin