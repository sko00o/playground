#!/usr/bin/env bash

cd "$(dirname "$0")"

echo ">> compose environment"
docker compose up -d --remove-orphans

echo ">> run demo"
docker compose exec -it local-host \
    ./run.sh

docker compose down --remove-orphans --volumes
