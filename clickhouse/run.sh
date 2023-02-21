#!/usr/bin/env bash

cd "$(dirname "$0")"

echo ">> compose environment"
docker compose up -d --remove-orphans

echo ">> run migrations"
docker compose exec -it golang-migrate sh -c ./run.sh

echo ">> run queries"
docker compose exec -it curl sh -c ./run.sh

docker compose down --remove-orphans --volumes
