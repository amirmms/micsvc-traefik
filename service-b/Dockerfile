FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod .
COPY main.go .

RUN go build -o service-b main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/service-b .

EXPOSE 8080
CMD ["./service-b"]