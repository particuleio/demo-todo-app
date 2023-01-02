#!/bin/bash

docker pull docker.io/library/redis:7
docker tag docker.io/library/redis:7 redis

# SERVER_ADDR should be set
# example:
# SERVER_ADDR=34.155.67.216:1323

docker build -t docker.io/sametma/server:1 -t docker.io/sametma/server:latest ./server
docker build -t docker.io/sametma/front:1 -t docker.io/sametma/front:latest --build-arg=SERVER_ADDR=localhost:1323 ./front


docker run -d --rm --net=host --name=redis redis
docker run -d --rm --net=host --name=server server
docker run -d --rm --net=host --name=front front
