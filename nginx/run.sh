#!/usr/bin/env bash

cd "$(dirname "$0")"

echo ">> compose environment"
docker compose up -d --remove-orphans

echo ">> send some requests"
CURL="docker compose exec curl curl"

$CURL -s http://nginx:80/
$CURL -s http://nginx:80/__1
$CURL -s 'http://nginx:80/"/__1'

echo ">> show nginx logs"
docker compose logs --tail 5 nginx

docker compose down --remove-orphans --volumes
