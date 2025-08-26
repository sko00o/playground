#!/usr/bin/env bash
set -e

cd "$(dirname "$0")"

COMPOSE_FILE="docker-compose.yml:ports.ollama.yml:ports.openwebui.yml"

# check if ollama model dir exists
default_model_dir="/usr/share/ollama/.ollama"
if [ -d "$default_model_dir" ]; then
    echo ">> using existing ollama model dir"
    COMPOSE_FILE="${COMPOSE_FILE}:volumes.ollama.local.yml"
elif docker volume ls | grep -q "ollama"; then
    echo ">> using existing ollama volume"
    COMPOSE_FILE="${COMPOSE_FILE}:volumes.ollama.yml"
else
    echo ">> no ollama volume found"
fi

if docker volume ls | grep -q "open-webui"; then
    echo ">> using existing open-webui volume"
    COMPOSE_FILE="${COMPOSE_FILE}:volumes.openwebui.yml"
else
    echo ">> no open-webui volume found"
fi

export COMPOSE_FILE

echo ">> compose environment"
docker compose up -d --build --remove-orphans

echo ">> please visit http://127.0.0.1:18080"
