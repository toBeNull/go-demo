FROM registry.cn-hangzhou.aliyuncs.com/matth/golang:1.18.3-alpine
WORKDIR /output
COPY ./output/go-hello /usr/local/bin/go-hello
CMD go-hello
