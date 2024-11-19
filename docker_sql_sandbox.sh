#!/bin/bash
# setup container
docker run --name sandbox -e MYSQL_ROOT_PASSWORD=root -d mysql:latest

docker start sandbox
docker exec -it sandbox bash
