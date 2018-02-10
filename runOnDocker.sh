#!/usr/bin/env bash
go build
docker build -t goserver .
docker run --rm -it goserver