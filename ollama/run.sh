#!/usr/bin/env bash
set -e

cd "$(dirname "$0")"

# Parse command line arguments
SKIP_TEARDOWN=false
for arg in "$@"; do
    case $arg in
    --skip-teardown | -s)
        SKIP_TEARDOWN=true
        shift
        ;;
    --help | -h)
        echo "Usage: $0 [--skip-teardown|-s]"
        exit 0
        ;;
    *)
        # Unknown option, ignore
        shift
        ;;
    esac
done

COMPOSE_FILE="docker-compose.yml:ollama.ports.yml"

# check if ollama model dir exists
default_model_dir="/usr/share/ollama/.ollama"
if [ ! -d "$default_model_dir" ]; then
    echo ">> ollama model dir not found"
else
    echo ">> using existing ollama model dir"
    COMPOSE_FILE="${COMPOSE_FILE}:volumes.yml"
fi
export COMPOSE_FILE

echo ">> compose environment"
docker compose up -d --build --remove-orphans

echo ">> please visit http://127.0.0.1:11434"

if [ "$SKIP_TEARDOWN" = true ]; then
    echo ">> skipping teardown (--skip-teardown flag used)"
    exit 0
fi

cleanup() {
    docker compose down --remove-orphans --volumes
    exit 0
}
trap cleanup SIGINT
trap cleanup SIGTERM

echo "press ctrl+c to stop"
tail -f /dev/null
