#!/bin/bash
set -e

version=$(date +%m%d.%H%M)
echo $version

flags="-X go-hello/version.version=${version}"
echo $flags


mkdir -p output
go mod download && go mod verify
GOOS=linux go build -ldflags "$flags" -o output/go-hello main.go


docker build -t cr-cn-beijing.volces.com/matthew-demo/go-demo:latest .
docker push cr-cn-beijing.volces.com/matthew-demo/go-demo:latest
kubectl apply -f manifest/deployment.yml