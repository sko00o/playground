#!/usr/bin/env bash
set -e

mkdir -p models output

# if models not found, prompt for download
if [ ! -f models/config.json ] || [ ! -f models/model.pth ]; then
    echo "Models not found, downloading..."

    wget --continue https://huggingface.co/sko00o/vits/resolve/main/paimon6k_390k.pth?download=true -O models/model.pth
    wget --continue https://huggingface.co/sko00o/vits/resolve/main/paimon6k.json?download=true -O models/config.json
fi

docker compose up vits --build

echo "Open ./output/test.wav to listen to the result."
