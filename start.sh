#!/bin/bash

./scope.sh launch

ELK_CONFIG_DIR=log/docker-compose.yml

docker-compose -f $ELK_CONFIG_DIR up -d --build
#docker-compose -f $ELK_CONFIG_DIR up -d

./build.sh

LOGSTASH_HOST=$(docker inspect --format='{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' log_logstash) 
LOGSTASH_HOST=$LOGSTASH_HOST  docker-compose -f docker-compose.yml up -d --build