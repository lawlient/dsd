FROM golang:alpine

RUN apk add --update --no-cache graphviz ttf-dejavu 

COPY . /app

ENV GOPROXY=https://goproxy.cn

WORKDIR /app