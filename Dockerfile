FROM golang:1.24.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN apt-get update && apt-get install -y make && make build

FROM debian:bookworm-slim

LABEL org.opencontainers.image.source="https://github.com/halon176/hcpb-api"

WORKDIR /app

COPY --from=builder /app/bin/hcpb-api /app/hcpb-api
COPY --from=builder /app/migrations /app/migrations

RUN chmod +x /app/hcpb-api

CMD ["/app/hcpb-api"]
