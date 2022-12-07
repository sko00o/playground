#!/usr/bin/env bash

cd "$(dirname "$0")"

echo ">> compose environment"
docker compose up -d --remove-orphans

echo ">> show docker info by Docker-CLI"
docker compose exec -it docker-client \
    docker info

echo ">> show docker info by API"
# API Reference: https://docs.docker.com/engine/api/v1.41/
docker compose exec -it curl \
    curl --unix-socket /var/run/docker.sock http://localhost/v1.41/info

docker compose down --remove-orphans
