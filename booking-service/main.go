package main

import (
	"booking-service/controllers"
	"booking-service/models"
	"booking-service/services"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

func main() {
	// Подключение к БД PostgreSQL для бронирований
	db, err := models.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	// Подключение к RabbitMQ
	rmqConn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer rmqConn.Close()

	// Инициализация сервисов
	bookingService := services.NewBookingService(db, rmqConn)
	bookingController := controllers.NewBookingController(bookingService)

	// Инициализация Gin
	r := gin.Default()

	// Роуты
	r.GET("/bookings", bookingController.ListBookings)
	r.POST("/bookings", bookingController.CreateBooking)
	r.GET("/bookings/:id", bookingController.GetBooking)

	// Запуск сервера
	r.Run(":8082")
}
