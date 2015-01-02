#!/bin/bash

TMP="./temp"
DIR="$(pwd)"

cd $DIR

if [ -d "$TMP" ]; then
    echo "deleting $TMP directory..."
    rm -r $TMP
fi

influxdb -config=./scripts/db.conf

