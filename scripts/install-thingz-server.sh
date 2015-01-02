#!/bin/bash

set -x
cd

# Go
if [ -d go ]; then
  echo "Go already installed"
else
  wget https://storage.googleapis.com/golang/go1.4.linux-amd64.tar.gz
  sudo tar -C /usr/local -xzf go1.4.linux-amd64.tar.gz
fi

# GOPATH

echo "
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin" >> /etc/profile

source /etc/profile


# InfluxDB
wget http://s3.amazonaws.com/influxdb/influxdb_latest_amd64.deb
sudo dpkg -i influxdb_latest_amd64.deb

sudo mv ./db-ubuntu.conf /opt/influxdb/shared/config.toml

sudo /etc/init.d/influxdb restart
