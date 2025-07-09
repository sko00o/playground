#!/bin/bash

# Build the Docker image
docker build -t png-chunk-reader .

# Run the Docker container
pushd ../../assets
docker run --rm -v $(pwd):/app/test/ png-chunk-reader
popd
