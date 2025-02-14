#!/usr/bin/env bash

cd "$(dirname "$0")"

SET_CFG="-f docker-compose-kafka.yml -f docker-compose-monitor.yml"

echo ">> compose environment"
docker compose $SET_CFG up -d --remove-orphans

echo "please visit http://localhost:3000"

function cleanup {
  docker compose $SET_CFG down --remove-orphans --volumes
  exit 0
}
trap 'cleanup' SIGINT
trap 'cleanup' SIGTERM
echo "press ctrl+c to stop"

# Wait for signals
while true; do
  sleep 1
done
