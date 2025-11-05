FROM golang:1.25-alpine AS builder

WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod* go.sum* ./

# Если go.mod не существует, инициализируем проект
RUN if [ ! -f go.mod ]; then \
      go mod init github.com/VovanOst/integration-app; \
    fi

# Копируем весь исходный код
COPY . .

# Скачиваем зависимости
RUN go mod download || true

# Тidy модули
RUN go mod tidy

# ✅ ИСПРАВЛЕНО: Компилируем всю папку cmd/api вместо только main.go
# Это гарантирует что Go видит все файлы пакета (main.go, server.go, migrate.go, health.go)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s" \
    -o api \
    ./cmd/api

# Runtime stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/api .

# Копируем .env файл если существует
COPY .env* ./

EXPOSE 8080

# Аргументы по умолчанию для сервера
CMD ["./api", "server"]
