#!/bin/bash

docker pull docker.io/library/redis:7
docker tag docker.io/library/redis:7 redis

# SERVER_ADDR should be set
# example:
# SERVER_ADDR=34.155.67.216:1323

docker build -t server ./server
docker build -t front --build-arg=SERVER_ADDR ./front


docker run -d --rm --net=host --name=redis redis
docker run -d --rm --net=host --name=server server
docker run -d --rm --net=host --name=front front
