#!/usr/bin/env bash

echo ">> compose environment"
docker compose up -d --remove-orphans

echo ">> wait for some transmition"
sleep 10

echo ">> show some logs"
docker compose logs --timestamps --since=10s --tail=5 consumer producer

if [ -x "$(command -v kcat)" ]; then
    echo ">> produce 10 messages using kcat"
    for i in {00..10}; do printf "%02d\n" $i; done |
        kcat -P -b localhost:19092 -t test1

    echo ">> consume 10 messages using kcat"
    kcat -C -b localhost:19092 -t test1 -o -10 -e
fi

docker compose down --remove-orphans
