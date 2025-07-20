# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o log_dashboard

# Final stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/log_dashboard .
COPY app.env .

EXPOSE 8080

CMD ["./log_dashboard"]
