#!/bin/bash

DIR="$(pwd)"

db_target="local"

while getopts ":r" opt; do
  case $opt in
    r)
      db_target="remote"
      ;;
    \?)
      echo "Invalid option: -$OPTARG" >&2
      ;;
  esac
done


echo "DB target: ${db_target}"


if [ $db_target = "remote" ]; then

    $DIR/thingz-server --api-port=8080 \
                       --ui-port=8081 \
                       --db="http://root:${THINGZ_ROOT_SECRET}@${THINGZ_HOST}:8086/thingz"

else

    $DIR/thingz-server --api-port=8080 \
                       --ui-port=8081 \
                       --db="http://root:root@localhost:8086/thingz"

fi




