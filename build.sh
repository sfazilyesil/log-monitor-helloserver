#!/usr/bin/env bash

# go kütüphanelerini statik olarak (derleme anında) bağlayarak derle
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o helloserver .

# docker imajını oluştur
#docker build -t sfy/helloserver .