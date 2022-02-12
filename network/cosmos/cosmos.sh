#!/bin/bash

# Install the specialized, ica-27 enabled gaiad binary
cd $HOME
yes | rm -r cosmos
git clone https://github.com/cosmos/gaia
cd gaia
git checkout ica-acct-auth
make install
sudo cp $GOPATH/bin/gaiad /usr/local/bin