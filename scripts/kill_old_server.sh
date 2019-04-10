set -ex

echo "attempting to kill old server proc"

whoami
pwd
ls -lah
cd /var/run
ls -lah
ls -lah | grep api-server.pid
! pkill -f /tmp/go-build

if [ $? -gt 0 ]
then
	echo "no pid found for api server, api server not running"
else
	echo "api server running, killing server"
	kill `cat api-server.pid`
	echo "api server destroyed"
fi
