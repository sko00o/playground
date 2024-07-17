#!/usr/bin/env bash

cd "$(dirname "$0")"

echo ">> compose environment"
docker compose up -d --remove-orphans

defer() {
    docker compose down --remove-orphans --volumes
    exit 0
}

echo "Please open http://127.0.0.1:8081"

trap defer SIGINT

# Wait for Ctrl+C
while true; do
    sleep 1
done

