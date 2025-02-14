#!/usr/bin/env bash
set -e

cd "$(dirname "$0")"

echo ">> compose environment"
docker compose up -d --build --remove-orphans

defer() {
    docker compose down --remove-orphans --volumes
    exit 0
}
trap defer EXIT

echo "Please open http://127.0.0.1:8188 in browser."
echo "Now exec any commands in container: (Enter Ctrl+D to quit.)"
docker compose exec -it comfyui bash
