#!/usr/bin/env bash
set -ex

source /etc/bashrc
export ENV=staging
export GOPATH=/root/go
cd /root/go/src/github.com/herdius/herdius-blockchain-api
go get ./...
go run /root/go/src/github.com/herdius/herdius-blockchain-api/server/server.go -env=staging > /dev/null 2> /dev/null < /dev/null &
echo "server started in background"
