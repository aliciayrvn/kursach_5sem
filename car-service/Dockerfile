# Для сборки образа Car-Service
# Dockerfile для сервисов
FROM golang:1.22 AS builder


WORKDIR /app

# Копируем модули и устанавливаем их
COPY go.mod go.sum ./
RUN go mod tidy

# Копируем исходный код
COPY . .

# Сборка
RUN go build -o car-service .

EXPOSE 8081

CMD ["./car-service"]
