#!/usr/bin/env bash

set -ex

source /etc/bashrc
export ENV=staging
export GOPATH=/root/go
service_name="herdius-api"
cd /root/go/src/github.com/herdius/herdius-blockchain-api

supervisorctl stop "$service_name"

export GO111MODULE=on
go build -o ./api-server ./server/server.go

supervisorctl start "$service_name"
