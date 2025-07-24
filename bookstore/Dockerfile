# syntax=docker/dockerfile:1
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -o api_server ./api/bookstore.go && \
    CGO_ENABLED=0 GOOS=linux go build -o add_rpc ./rpc/add/add.go && \
    CGO_ENABLED=0 GOOS=linux go build -o check_rpc ./rpc/check/check.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/api_server ./api_server
COPY --from=builder /app/add_rpc ./add_rpc
COPY --from=builder /app/check_rpc ./check_rpc
COPY api/etc ./api/etc
COPY rpc/add/etc ./rpc/add/etc
COPY rpc/check/etc ./rpc/check/etc

# 默认启动api服务，如需启动rpc服务可覆盖CMD
CMD ["./api_server", "-f", "api/etc/bookstore-api.yaml"]