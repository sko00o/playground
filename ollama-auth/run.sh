#!/usr/bin/env bash
set -e

cd "$(dirname "$0")"

COMPOSE_FILE="docker-compose.yml:nginx.yml:curl.yml"

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

defer() {
    docker compose down --remove-orphans --volumes
    exit 0
}
trap defer EXIT

echo ">> send some requests"
CURL="docker compose exec curl curl"

echo ">> 1. without api key"
$CURL -i http://localhost:11434 || true # ignore error
echo

echo ">> 2. with correct api key"
$CURL -i http://localhost:11434 -H "Authorization: Bearer YOUR_SECRET_API_KEY"
echo
