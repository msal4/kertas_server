FROM golang:1.16-alpine AS build-env

WORKDIR /src
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN ["go", "build", "-o", "build/migrate", "cmd/migrate-schema/main.go"]
RUN ["go", "build", "-o", "build/server", "cmd/server/main.go"]

FROM alpine

WORKDIR /app
COPY --from=build-env /src/build/* ./
COPY --from=build-env /src/config.yml .
CMD ["./server"]
