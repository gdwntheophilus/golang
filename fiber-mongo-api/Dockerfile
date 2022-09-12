#!/bin/bash

FROM golang:1.15-alpine

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app/fiber-mongo-api

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .



# Build the Go app
RUN go build -o /app/fiber-mongo-api/out/fiber-mongo-api .
RUN chmod +x /app/fiber-mongo-api/out/fiber-mongo-api

# This container exposes port 8080 to the outside world
EXPOSE 6000

# Run the binary program produced by `go install`
CMD ["/app/fiber-mongo-api/out/fiber-mongo-api"]
