package main

import (
	"car-service/controllers"
	"car-service/models"
	"car-service/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Подключение к БД PostgreSQL для машин
	db, err := models.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	// Инициализация сервисов
	carService := services.NewCarService(db)
	carController := controllers.NewCarController(carService)

	// Инициализация Gin
	r := gin.Default()

	// Роуты
	r.GET("/cars", carController.ListCars)
	r.POST("/cars", carController.CreateCar)
	r.GET("/cars/:id", carController.GetCar)

	// Запуск сервера
	r.Run(":8081")
}
