#!/usr/bin/env bash

cd "$(dirname "$0")"

echo ">> compose environment"
docker compose up -d --remove-orphans

echo ">> attach CLI"
docker compose exec -it redis redis-cli

docker compose down --remove-orphans --volumes
