# syntax=docker/dockerfile:1
FROM golang:1.15-buster as builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download -x

COPY . ./


RUN go build -v -o /bin/server cmd/app/main.go

FROM debian:buster-slim
RUN set -x && apt-get update && \
  DEBIAN_FRONTEND=noninteractive apt-get install -y ca-certificates && \
  rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY config.yaml ./
COPY --from=builder /bin/server ./

ENV PORT=8080
ENV DATABASE_URL=""

EXPOSE $PORT

CMD ["./server", "-config", "config.yaml"]