#!/bin/bash

if rm -r ./build; then
  echo removed ./build directory
fi

if ! go build -o build/migrate cmd/migrate-schema/main.go; then
  echo failed to build migrate tool
  exit
fi
echo built migrate cmd in build/migrate

if ! go build -o build/seed cmd/seed/main.go; then
  echo failed to build seed tool
  exit
fi
echo built seed cmd in build/seed

if ! go build -o build/server cmd/server/main.go; then
  echo failed to build server
  exit
fi
echo built server cmd in build/server

cp config.yml build/
cp -r testfiles build/testfiles
echo copied config.yml and testfiles to ./build

zip -r kertas.zip ./build

KERTAS_PATH=$(pwd)/kertas.zip
echo bundled in $KERTAS_PATH

if [[ $OSTYPE == 'darwin'* ]]; then
  open -R $KERTAS_PATH
fi
