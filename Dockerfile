FROM golang:1.16 as build-env

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

FROM debian:10.8-slim
RUN apt-get update && \
    apt-get -y install curl git openssl && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /tmp/etisbew
COPY --from=build-env /go/bin/etisbew /usr/bin/etisbew
