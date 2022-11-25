#!/usr/bin/env bash

docker compose up -d --remove-orphans

docker compose exec -it docker-client \
    docker info

docker compose down --remove-orphans
