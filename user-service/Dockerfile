FROM golang:alpine AS builder

RUN apt update && apt add --no-cache git

WORKDIR /go/src/go-docker/user-service

COPY . .

