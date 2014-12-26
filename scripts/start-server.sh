#!/bin/bash

DIR="$(pwd)"

$DIR/thingz-server --api-port=8080 \
                   --ui-port=8081 \
                   --db="http://root:root@localhost:8086/thingz"