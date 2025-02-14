#!/usr/bin/env bash

cd "$(dirname "$0")"

CFG=docker-compose.yml

echo ">> compose environment"
docker compose -f "$CFG" up -d --remove-orphans

echo ">> using jconsole to connect JMX port"
# if jconsole not in $PATH print warning
if [[ ! -x "$(command -v jconsole)" ]]; then
  echo "WARN: jconsole not found in \$PATH"
else
  jconsole localhost:9999
fi

function cleanup {
  docker compose -f "$CFG" down --remove-orphans --volumes
  exit 0
}
trap 'cleanup' SIGINT
trap 'cleanup' SIGTERM
echo "press ctrl+c to stop"

# Wait for signals
while true; do
  sleep 1
done
