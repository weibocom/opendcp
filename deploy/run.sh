#!/bin/sh

export VER=latest
export REPO=weibo.com
docker-compose down $1
docker-compose up $1
