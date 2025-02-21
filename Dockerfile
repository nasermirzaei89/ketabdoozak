FROM alpine:3.21.3 AS runner

RUN apk update

###

FROM golang:1.24.0-alpine3.21 AS builder

RUN apk update
RUN apk add make build-base

WORKDIR /src

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN IN_DOCKER=true make build

###

FROM runner

COPY --from=builder /src/bin/ketabdoozak /

ENTRYPOINT ["/ketabdoozak"]
