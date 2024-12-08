package main

import (
	"log"
	"user-service/controllers"
	"user-service/models"
	"user-service/services"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Настройка строки подключения к базе данных
	dsn := "host=postgres-user user=postgres password=example dbname=user_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	// Миграция базы данных
	db.AutoMigrate(&models.User{})

	// Инициализация сервисов и контроллеров
	userService := services.NewUserService(db)
	userController := controllers.NewUserController(userService)

	// Настройка маршрутов с помощью Gin
	r := gin.Default()

	// Роуты для управления пользователями
	r.GET("/users/:id", userController.GetUser)
	r.GET("/users", userController.ListUsers)
	r.POST("/users", userController.CreateUser)

	// Запуск сервера на порту 8081
	r.Run(":8081")
}
