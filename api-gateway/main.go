package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

func main() {
	// Подключаемся к RabbitMQ
	rmqConn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}
	defer rmqConn.Close()

	ch, err := rmqConn.Channel()
	if err != nil {
		log.Fatal("Failed to open a channel:", err)
	}
	defer ch.Close()

	// Настройка очереди для уведомлений о бронированиях
	q, err := ch.QueueDeclare(
		"booking", // имя очереди
		false,     // не персистентная
		false,     // не удаляется при закрытии соединения
		true,      // эксклюзивная
		false,     // не ожидать ответа
		nil,       // дополнительные аттрибуты
	)
	if err != nil {
		log.Fatal("Failed to declare a queue:", err)
	}

	// Инициализация Gin
	r := gin.Default()

	// Прокси для User-Service
	r.GET("/users/:id", func(c *gin.Context) {
		// Прокси-запрос к user-service
		c.Redirect(http.StatusMovedPermanently, "http://user-service:8081/users/"+c.Param("id"))
	})

	// Прокси для Car-Service
	r.GET("/cars/:id", func(c *gin.Context) {
		// Прокси-запрос к car-service
		c.Redirect(http.StatusMovedPermanently, "http://car-service:8081/cars/"+c.Param("id"))
	})

	// Прокси для Booking-Service
	r.POST("/bookings", func(c *gin.Context) {
		// Прокси-запрос к booking-service
		c.Redirect(http.StatusMovedPermanently, "http://booking-service:8082/bookings")
	})

	// Обработка уведомлений от RabbitMQ (например, уведомление о новом бронировании)
	go func() {
		msgs, err := ch.Consume(
			q.Name, // имя очереди
			"",     // consumer
			true,   // автоматически подтверждать сообщения
			false,  // не эксклюзивная
			false,  // не ожидать ответа
			false,  // не временная
			nil,    // дополнительные аттрибуты
		)
		if err != nil {
			log.Fatal("Failed to register a consumer:", err)
		}

		for msg := range msgs {
			log.Printf("Received a message: %s", msg.Body)
			// Здесь можно обработать уведомление (например, отправить email или логировать)
		}
	}()

	// Запуск сервера
	r.Run(":8080")
}
