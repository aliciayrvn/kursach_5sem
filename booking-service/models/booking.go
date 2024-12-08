package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	UserID uint   `json:"user_id"`
	CarID  uint   `json:"car_id"`
	Status string `json:"status"`
}

func ConnectDB() (*gorm.DB, error) {
	// Подключение к PostgreSQL для бронирований
	dsn := "host=postgres-user user=postgres password=example dbname=user_db port=5432 sslmode=disable"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
