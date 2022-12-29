#!/bin/bash

docker pull docker.io/library/redis:7
docker tag docker.io/library/redis:7 redis


docker build -t server ./server
docker build -t front ./front


docker run -d --rm --net=host --name=redis redis
docker run -d --rm --net=host --name=server server
docker run -d --rm --net=host --name=front front
