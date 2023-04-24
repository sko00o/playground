#!/usr/bin/env bash

cd "$(dirname "$0")"

echo ">> compose environment"
docker compose up -d --remove-orphans

echo ">> list all clusters"
docker compose exec curl sh -c '
    echo "select * from system.clusters
    format PrettyCompact" |
    curl http://ch00:8123/?database=default \
    --data-binary @-
    '

docker compose down --remove-orphans --volumes
