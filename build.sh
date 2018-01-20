#!/usr/bin/env bash

GOOS=linux go build -o build/echo-server main.go

docker build -t echo .
