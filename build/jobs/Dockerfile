FROM golang:1.16-alpine3.13 as builder
WORKDIR /go/src/vsantosalmeida/go-app-engine-demo

COPY . .
RUN apk add --no-cache \
   build-base \
   gcc \
   git \
   pkgconf \
   musl-dev
RUN export GOPRIVATE=github.com/vsantosalmeida/*
RUN make build-job

FROM alpine
WORKDIR /root/
COPY --from=builder /go/src/vsantosalmeida/go-app-engine-demo/bin/job.bin .