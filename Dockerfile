FROM golang:1.18.3

WORKDIR /usr/src/app

COPY go.mod go.sum ./
COPY . .
RUN go mod tidy
