#!/usr/bin/env bash
set -ex

source /etc/bashrc
export ENV=staging
export GOPATH=/root/go
cd /root/go/src/github.com/herdius/herdius-blockchain-api
mkdir -p /var/log/herdius/herdius-blockchain-api/log/
RUNDIR=/var/run/herdius
if [ ! -d "$RUNDIR" ]; then
  mkdir -p "$RUNDIR"
fi

pidfile="$RUNDIR/api-server.pid"

# Kill old process if existed
if [[ -f "$pidfile" ]]; then
  kill "$(cat "$pidfile")" || :
fi

# Wait api server quit
sleep 3

go build -o ./api-server ./server/server.go
./api-server -env=staging > /var/log/herdius/herdius-blockchain-api/log/server.log 2>&1 </dev/null &

# Save the pid to kill later
printf '%s\n' "$!" >"$pidfile"

echo "server started in background"
