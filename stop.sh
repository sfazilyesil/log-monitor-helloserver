#!/bin/bash

docker-compose -f docker-compose.yml down
docker-compose -f log/docker-compose.yml down

docker ps |grep scope |awk '{print $1}' |xargs docker rm -f