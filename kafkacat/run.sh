#!/usr/bin/env bash
set -e
cd "$(dirname "$0")"

echo ">> compose environment"
docker compose up -d --remove-orphans

echo ">> wait for some transmition"
sleep 5

echo ">> show consumer logs"
docker compose logs --tail=10 --no-log-prefix consumer

docker compose down --remove-orphans --volumes
