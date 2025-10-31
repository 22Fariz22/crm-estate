# syntax=docker/dockerfile:1
FROM golang:1.25.1-alpine AS builder

WORKDIR /app

# Кэшируем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем код
COPY . .

# Собираем из cmd/api/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

# Финальный образ
FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/migrations ./migrations
COPY .env .

EXPOSE 8080

CMD ["./main"]
