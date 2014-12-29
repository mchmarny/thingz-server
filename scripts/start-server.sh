#!/bin/bash

DIR="$(pwd)"

db_target="http://root:root@localhost:8086/thingz"
http_proxy=""

while getopts ":r" opt; do
  case $opt in
    r)
      db_target="http://root:${THINGZ_ROOT_SECRET}@${THINGZ_HOST}:8086/thingz"
      ;;
    p)
      http_proxy="http://proxy-us.intel.com:911"
      ;;
    \?)
      echo "Invalid option: -$OPTARG" >&2
      ;;
  esac
done


echo "DB target: ${db_target}"
echo "HTTP proxy: ${http_proxy}"


$DIR/thingz-server --api-port=8080 \
                   --ui-port=8081 \
                   --proxy="${http_proxy}" \
                   --db="${db_target}"




