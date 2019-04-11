set -ex

whoami
pwd
ls -lah
source /etc/bashrc
echo "GOPATH = $GOPATH"
cd /root/go/src/github.com/herdius/herdius-blockchain-api
go get ./...
go run /root/go/src/github.com/herdius/herdius-blockchain-api/server/server.go &
echo "server started in background"
