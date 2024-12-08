# Для сборки образа User-Service
# Dockerfile для сервисов
FROM golang:1.22 AS builder


WORKDIR /app

# Копируем модули и устанавливаем их
COPY go.mod go.sum ./
RUN go mod tidy

# Копируем исходный код
COPY . .

# Сборка
RUN go build -o user-service .

EXPOSE 8080

CMD ["./user-service"]
