#!/bin/bash

DIR="$(pwd)"

db_target="http://root:root@localhost:8086/thingz"
kafka="false"

while getopts ":r:k" opt; do
  case $opt in
    r)
      db_target="http://root:${THINGZ_ROOT_SECRET}@${THINGZ_HOST}:8086/thingz"
      ;;
    k)
      kafka="true"
      ;;
    \?)
      echo "Invalid option: -$OPTARG" >&2
      ;;
  esac
done


$DIR/thingz-server --api-port=8080 \
                   --ui-port=8081 \
                   --db="${db_target}" \
                   --kafka=$kafka \
                   --kafka-topic="thingz" \
                   --kafka-brokers="localhost:9092"





