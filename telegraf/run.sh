#!/usr/bin/env bash

echo ">> compose environment"
docker compose up -d --remove-orphans

echo ">> show es index"
docker compose exec -it tool bash -c ./run.sh
