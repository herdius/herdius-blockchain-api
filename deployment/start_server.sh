#!/usr/bin/env bash
set -ex

source /etc/bashrc
export ENV=staging
export GOPATH=/root/go
cd /root/go/src/github.com/herdius/herdius-blockchain-api
mkdir -p /var/log/herdius/herdius-blockchain-api/log/
go get ./...
go run /root/go/src/github.com/herdius/herdius-blockchain-api/server/server.go -env=staging > /var/log/herdius/herdius-blockchain-api/log/server.log 2> /var/log/herdius/herdius-blockchain-api/log/server.log < /dev/null &
echo "server started in background"