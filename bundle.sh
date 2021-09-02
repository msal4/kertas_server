#!/bin/bash

rm -rf ./build
if [ $? -ne 0 ]; then
    exit
fi
echo removed ./build directory

go build -o build/migrate cmd/migrate-schema/main.go
if [ $? -ne 0 ]; then
    echo failed to build migrate tool
    exit
fi
echo built migrate cmd in build/migrate

go build -o build/seed cmd/seed/main.go
if [ $? -ne 0 ]; then
    echo failed to build seed tool
    exit
fi
echo built seed cmd in build/seed

go build -o build/server cmd/server/main.go
if [ $? -ne 0 ]; then
    echo failed to build server
    exit
fi
echo built server cmd in build/server

cp config.yml build/
cp -r testfiles build/testfiles
echo copied config.yml and testfiles to ./build

zip -r kertas.zip ./build

echo bundled in ./kertas.zip
