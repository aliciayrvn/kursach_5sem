package main

import (
	"log"
	"user-service/controllers"
	"user-service/models"
	"user-service/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Подключение к БД PostgreSQL
	db, err := models.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	// Инициализация сервисов
	userService := services.NewUserService(db)
	userController := controllers.NewUserController(userService)

	// Инициализация Gin
	r := gin.Default()

	// Роуты
	r.GET("/users", userController.ListUsers)
	r.POST("/users", userController.CreateUser)
	r.GET("/users/:id", userController.GetUser)

	// Запуск сервера
	r.Run(":8080")
}
