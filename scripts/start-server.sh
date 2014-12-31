#!/bin/bash

DIR="$(pwd)"

db_target="http://root:root@localhost:8086/thingz"
kafka=false
verbose=false

while getopts ":vrk" opt; do
  case $opt in
    r)
      db_target="http://root:${THINGZ_ROOT_SECRET}@${THINGZ_HOST}:8086/thingz"
      ;;
    k)
      kafka=true
      ;;
    v)
      verbose=true
      ;;
    \?)
      echo "Invalid option: -$OPTARG" >&2
      ;;
  esac
done


$DIR/thingz-server --verbose=$verbose \
                   --api-port=8080 \
                   --ui-port=8081 \
                   --db="${db_target}" \
                   --laod=$kafka \
                   --topic="thingz" \
                   --brokers="localhost:9092" \
                   --pub-db="udp://thingz:thingz@localhost:4444/thingz"





