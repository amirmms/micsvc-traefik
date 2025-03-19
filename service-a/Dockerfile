FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod .
COPY main.go .

RUN go build -o service-a main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/service-a .

EXPOSE 8080
CMD ["./service-a"]