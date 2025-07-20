# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o log_collector

# Final stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/log_collector .
COPY app.env .

EXPOSE 50051

CMD ["./log_collector"]
