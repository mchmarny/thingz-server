#!/bin/bash

DIR="$(pwd)"

influxdb -config=$DIR/scripts/db.conf

