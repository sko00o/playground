#!/usr/bin/env bash
set -e

cd "$(dirname "$0")"

echo ">> compose environment"
docker compose up -d --build --remove-orphans

defer() {
    docker compose down --remove-orphans --volumes
    exit 0
}
trap defer TERM INT EXIT

echo "Please open http://127.0.0.1:8080 in browser."
tail -f /dev/null # stop here
