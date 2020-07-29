#!/usr/bin/env bash

# -tags netgo apline构建golang编译问题

# go mod 中的静态资源引入问题
#GOOS=linux GOARCH=amd64 go build -v -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main
export CGO_ENABLED=0
GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -tags netgo -o main
docker build -f ./Dockerfile -t registry.cn-hangzhou.aliyuncs.com/dreamlu/common:w2socks-client .

# remove build
rm -rf main