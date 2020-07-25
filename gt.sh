#!/usr/bin/env bash
./devMode.sh dev

cd docker
./pushAll.sh
cd ..
#./devMode.sh dev