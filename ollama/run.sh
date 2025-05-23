#!/usr/bin/env bash
set -e

cd "$(dirname "$0")"

COMPOSE_FILE="docker-compose.yml:ollama.ports.yml"

# check if ollama model dir exists
default_model_dir="/usr/share/ollama/.ollama"
if [ ! -d "$default_model_dir" ]; then
    echo ">> ollama model dir not found"
else
    echo ">> using existing ollama model dir"
    COMPOSE_FILE="${COMPOSE_FILE}:volumes.yml"
fi
export COMPOSE_FILE

echo ">> compose environment"
docker compose up -d --build --remove-orphans

cleanup() {
    docker compose down --remove-orphans --volumes
    exit 0
}
trap cleanup SIGINT
trap cleanup SIGTERM

echo ">> please visit http://127.0.0.1:11434"
echo "press ctrl+c to stop"
tail -f /dev/null
