#!/usr/bin/env bash

echo ">> compose environment"
docker compose up -d --remove-orphans

echo ">> wait for some transmition"
sleep 10

echo ">> show some logs"
docker compose logs --timestamps --since=10s --tail=5 consumer producer

for cmd in kcat kafkacat; do
if [ -x "$(command -v $cmd)" ]; then
    echo ">> found $cmd"
    echo ">> produce 10 messages using kcat"
    for i in {00..10}; do printf "%02d\n" $i; done |
        $cmd -P -b localhost:19092 -t test1

    echo ">> consume 10 messages using kcat"
    $cmd -C -b localhost:19092 -t test1 -o -10 -e

    break
fi
done

docker compose down --remove-orphans
