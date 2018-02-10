#!/usr/bin/env bash
go build
docker build -t goserver .
docker run --rm -it -p 8080:8080 goserver