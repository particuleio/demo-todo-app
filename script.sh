#!/bin/bash

docker pull docker.io/library/redis:7
docker tag docker.io/library/redis:7 redis


docker build -t servertest ./server
docker build -t fronttest ./front


docker run -d --rm --net=host --name=redis redis
docker run -d --rm --net=host --name=server servertest
docker run -d --rm --net=host --name=front fronttest
