#!/bin/bash

# this is only for demo
# obviously you'll change the root password, right?

HOST="localhost"
PSWD="root"

echo "Using"
echo "   http://${HOST}:8083/"

# add db
echo "Creating DB..."
curl -X POST "http://${HOST}:8086/db?u=root&p=${PSWD}" \
     -d '{"name": "thingz"}'

# add db user
echo "Creating DB user..."
curl -X POST "http://${HOST}:8086/db/thingz/users?u=root&p=${PSWD}" \
  -d '{"name": "thingz", "password": "thingz"}'

echo "Created"
echo "   Database: thingz"
echo "   User: thingz"
echo "   Password: thingz"