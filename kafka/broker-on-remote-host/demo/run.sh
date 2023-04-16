#!/usr/bin/env sh

cd "$(dirname "$0")"
set -e

echo ">> compose environment"
docker compose up -d --remove-orphans

broker="remote-host:19092"
cmd="docker run --rm -i --net=host confluentinc/cp-kafkacat:7.1.5 kafkacat"
echo ">> produce 10 messages using kcat"
for i in $(seq 1 10); do printf "%02d\n" $i; done |
    $cmd -P -b ${broker} -t test
echo ">> consume 10 messages using kcat"
$cmd -C -b ${broker} -t test -o -10 -e

docker compose down --remove-orphans --volumes
