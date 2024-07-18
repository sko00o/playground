#!/usr/bin/env bash
set -e
cd "$(dirname "$0")"

echo ">> compose environment"
docker compose up -d --remove-orphans

echo ">> please visit http://127.0.0.1:18090"

function cleanup {
    docker compose down --remove-orphans --volumes
    exit 0
}
trap 'cleanup' SIGINT
trap 'cleanup' SIGTERM
echo "press ctrl+c to stop"

# Wait for signals
while true; do
  sleep 1
done
