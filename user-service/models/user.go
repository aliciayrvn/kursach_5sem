package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ConnectDB() (*gorm.DB, error) {
	// Подключение к PostgreSQL
	dsn := "host=postgres-user user=postgres password=example dbname=user_db port=5432 sslmode=disable"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
