#!/bin/bash

(cd src && go build .)
mkdir -p build
mv src/safe-server ./build/safe-server