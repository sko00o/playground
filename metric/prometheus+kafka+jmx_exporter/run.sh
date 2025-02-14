#!/usr/bin/env bash
set -e

cd "$(dirname "$0")"

function cleanup {
    docker compose down --remove-orphans --volumes
    exit 0
}

echo ">> compose environment"
docker compose up -d --remove-orphans

trap 'cleanup' SIGINT
trap 'cleanup' SIGTERM

echo ">> please visit http://127.0.0.1:19090"
echo "press ctrl+c to stop"

tail -f /dev/null
