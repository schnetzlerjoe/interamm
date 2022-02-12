#!/bin/bash

cd $HOME
yes | rm -r cosmos
git clone https://github.com/cosmos/gaia
cd gaia
git checkout ica-acct-auth
make install
sudo cp $GOPATH/bin/gaiad /usr/local/bin