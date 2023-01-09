#!/bin/bash

docker pull docker.io/library/redis:7
docker tag docker.io/library/redis:7 redis

# SERVER_ADDR should be set
# example:
export SERVER_ADDR=http://localhost:1323/api

docker build -t front -t docker.io/sametma/front:1 -t docker.io/sametma/front:latest --build-arg=SERVER_ADDR=$SERVER_ADDR ./front
docker build -t server -t docker.io/sametma/server:1 -t docker.io/sametma/server:latest ./server


docker run -d --rm --net=host --name=redis redis
docker run -d --rm --net=host --name=server server
docker run -d --rm --net=host --name=front front



docker run -d --rm --net=test --name=redis redis
docker run -d --rm --net=test -p 1323:1323 --name=server server
docker run -d --rm --net=test -p 8888:8080 --name=front front
