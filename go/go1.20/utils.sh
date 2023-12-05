#!/usr/bin/env bash

compose-up() {
    docker compose up -d --remove-orphans
}
compose-down() {
    docker compose down --remove-orphans --volumes
}
GO() {
    docker compose exec -it dev120 go $@
}
GO118() {
    docker compose exec -it dev118 go $@
}

run-comparable() {
    echo ">>> run in Go 1.18"
    GO118 run ./comparable

    echo ">>> run in Go 1.20"
    GO run ./comparable
}
