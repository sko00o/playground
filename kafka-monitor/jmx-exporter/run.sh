#!/usr/bin/env bash

cd "$(dirname "$0")"

CFG=docker-compose.yml

echo ">> compose environment"
docker compose -f "$CFG" up -d --remove-orphans

echo "please visit http://localhost:9090"

function cleanup {
  docker compose -f "$CFG" down --remove-orphans --volumes
  exit 0
}
trap 'cleanup' SIGINT
trap 'cleanup' SIGTERM
echo "press ctrl+c to stop"

# Wait for signals
while true; do
  sleep 1
done
