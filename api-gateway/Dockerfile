# Для сборки образа API Gateway
FROM golang:1.22 AS builder


WORKDIR /app

# Копируем модули и устанавливаем их
COPY go.mod go.sum ./
RUN go mod tidy

# Копируем исходный код
COPY . .

# Сборка
RUN go build -o api-gateway .

EXPOSE 8080

CMD ["./api-gateway"]
