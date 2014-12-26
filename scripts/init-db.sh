#!/bin/bash

# this is only for demo
# obviously you'll change the root password, right?

HOST="localhost"
PSWD="root"
NAME="thingz"


echo "Using"
echo "   http://${HOST}:8083/"

# add db
echo "Creating DB..."
curl -X POST "http://${HOST}:8086/db?u=root&p=${PSWD}" \
     -d '{"name": "'"${NAME}"'"}'

# add db user
echo "Creating DB user..."
curl -X POST "http://${HOST}:8086/db/${NAME}/users?u=root&p=${PSWD}" \
  -d '{"name": "'"${NAME}"'", "password": "'"${NAME}"'"}'

echo "Created"
echo "   Database: ${NAME}"
echo "   User: ${NAME}"
echo "   Password: ${NAME}"