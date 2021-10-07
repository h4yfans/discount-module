FROM golang:1-alpine AS builder


LABEL maintainer="Kaan Karaca, kaan94karaca@gmail.com"

RUN apk add --no-cache git


# Set the Current Working Directory inside the container
WORKDIR /go/src/github.com/h4yfans/discount-module/


ENV GO111MODULE=on

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o .

# Second stage
FROM alpine

COPY --from=builder /go/src/github.com/h4yfans/discount-module/ .

ENV PROVIDERS_PATH="./resource/input/providers.txt"
ENV INPUT_PATH="./resource/input/input.txt"
ENV CONFIG_PATH="./resource/config/config.json"

CMD ./discount-module