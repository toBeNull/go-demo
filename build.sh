#!/bin/bash
set -e

version=$(date +%m%d.%H%M)
echo $version

flags="-X go-hello/version.version=${version}"
echo $flags


mkdir -p output
go mod download && go mod verify
GOOS=linux go build -ldflags "$flags" -o output/go-hello main.go




