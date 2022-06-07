FROM golang:1.18.3-alpine3.16

WORKDIR /usr/src/app

COPY ./go.mod ./go.sum ./
RUN go mod download
