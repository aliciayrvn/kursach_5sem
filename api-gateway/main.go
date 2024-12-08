package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Пример маршрутов
	r.GET("/api/users", func(c *gin.Context) {
		// gRPC вызов в user-service
		c.JSON(http.StatusOK, gin.H{"message": "Users route"})
	})

	r.GET("/api/cars", func(c *gin.Context) {
		// gRPC вызов в car-service
		c.JSON(http.StatusOK, gin.H{"message": "Cars route"})
	})

	log.Println("API Gateway запущен на :8080")
	r.Run(":8080")
}
