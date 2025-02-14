#!/usr/bin/env bash

compose-up() {
    docker compose up -d --remove-orphans
}
compose-down() {
    docker compose down --remove-orphans --volumes
}
GO() {
    docker compose exec -it dev118 go $@
}
