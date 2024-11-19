#!/bin/bash

addr=$(docker inspect sandbox | jq -r '.[0].NetworkSettings.Networks[].IPAddress')

go run server.go $addr
