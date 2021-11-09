FROM golang:1.16-alpine AS build-env

WORKDIR /src
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN ["go", "build", "-o", "build/migrate", "cmd/migrate-schema/main.go"]
RUN ["go", "build", "-o", "build/server", "cmd/server/main.go"]

FROM node:14.18 as node-env

ARG REACT_APP_GRAPHQL_URL
ARG REACT_APP_CDN_URL

WORKDIR /src/dashboard

COPY dashboard/package.json .
COPY dashboard/yarn.lock .

RUN yarn

COPY --from=build-env /src/dashboard .

RUN PUBLIC_URL=/dashboard REACT_APP_GRAPHQL_URL=${REACT_APP_GRAPHQL_URL} REACT_APP_GRAPHQL_URL=${REACT_APP_GRAPHQL_URL} yarn build

FROM alpine

WORKDIR /app
COPY --from=build-env /src/build/* ./
COPY --from=build-env /src/config.yml .
COPY --from=node-env /src/dashboard/build ./dashboard/build

EXPOSE 3000

CMD ["./server"]
