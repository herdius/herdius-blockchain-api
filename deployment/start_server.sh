#!/usr/bin/env bash
set -ex

source /etc/bashrc
cd /root/go/src/github.com/herdius/herdius-blockchain-api
go get ./...
go run /root/go/src/github.com/herdius/herdius-blockchain-api/server/server.go > /dev/null 2> /dev/null < /dev/null &
echo "server started in background"
