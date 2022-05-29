#! /usr/bin/bash

echo "Starting docker-dev.sh"

docker build -f Dockerfile.dev -t firstneverrest/flexdict:1.0.0 .
docker run -d -p 3000:3000 firstneverrest/flexdict:1.0.0

echo "Finished docker-dev.sh"

