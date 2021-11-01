FROM golang:1.16-alpine AS build-env

WORKDIR /src
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go build cmd/server -o server

FROM alpine

WORKDIR /app
COPY --from=build-env /src/server .
ENTRYPOINT ./server
