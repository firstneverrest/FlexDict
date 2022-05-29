#! /usr/bin/bash

echo "Starting docker-dev.sh"

docker-compose down
docker-compose build
docker-compose up -d

echo "Finished docker-dev.sh"

