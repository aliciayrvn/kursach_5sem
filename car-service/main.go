package main

import (
	"car-service/controllers"
	"car-service/models"
	"car-service/services"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Настройка строки подключения
	dsn := "host=postgres-car user=postgres password=example dbname=car_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	// Миграция
	db.AutoMigrate(&models.Car{})

	// Инициализация сервисов и контроллеров
	carService := services.NewCarService(db)
	carController := controllers.NewCarController(carService)

	// Настройка маршрутов Gin
	r := gin.Default()

	r.GET("/cars/:id", carController.GetCar)
	r.GET("/cars", carController.ListCars)
	r.POST("/cars", carController.CreateCar)

	// Запуск сервера
	r.Run(":8082")
}
