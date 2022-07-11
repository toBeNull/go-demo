#!/bin/bash
set -e

version=$(date +%m%d.%H%M)
echo $version

flags="-X go-hello/main/version.version=${version}"
echo $flags


mkdir -p output
GOOS=linux go build -ldflags "$flags" -o output/go-hello main.go


