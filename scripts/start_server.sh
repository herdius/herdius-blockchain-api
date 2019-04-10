set -ex

echo "this yaaa is a working cicd pipeline test"
whoami
pwd
ls -lah
source /etc/bashrc
echo "GOPATH = $GOPATH"
cd /root/go/src/github.com/herdius/herdius-blockchain-api
ls -lah
go get ./...
go run /root/go/src/github.com/herdius/herdius-blockchain-api/server/server.go &
echo $! > /var/run/api-server.pid
