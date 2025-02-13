#!/usr/bin/env bash

cd "$(dirname "$0")"

function cleanup {
  docker compose down --remove-orphans --volumes
  exit 0
}
trap 'cleanup' SIGINT
trap 'cleanup' SIGTERM

echo ">> compose environment"
docker compose up -d --remove-orphans

echo ">> show wg-client interfaces"
docker compose exec wg-client wg show

echo ">> show wg-client ip routes"
docker compose exec wg-client ip route show

echo ">> test ping wg-server from wg-client"
docker compose exec wg-client ping -c 4 10.13.13.1

echo "press ctrl+c to stop"

# Wait for signals
while true; do
  sleep 1
done
