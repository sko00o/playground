#!/usr/bin/env bash

cd "$(dirname "$0")"

echo ">> compose environment"
docker compose up -d --remove-orphans

echo ">> wait for some transmition"
sleep 10

echo ">> show some logs"
docker compose logs --timestamps --since=10s --tail=5 consumer producer

docker compose down --remove-orphans
