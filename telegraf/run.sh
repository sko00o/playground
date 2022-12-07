#!/usr/bin/env bash

cd "$(dirname "$0")"

echo ">> compose environment"
docker compose up -d --remove-orphans

echo ">> show es index"
docker compose exec -it tool bash -c ./run.sh

docker compose down --remove-orphans
