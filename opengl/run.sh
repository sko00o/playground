#!/usr/bin/env bash

cd "$(dirname "$0")"

docker build -t glfw-app .
docker run --gpus all glfw-app python3 main.py
