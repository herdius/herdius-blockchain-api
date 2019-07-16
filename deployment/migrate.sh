#!/usr/bin/env bash

set -ex

source /etc/bashrc
export GOPATH=/root/go
cd /root/go/src/github.com/herdius/herdius-blockchain-api/cmd/migrate
go run main.go up
