package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Общая структура Model для всех моделей
type Model struct {
	ID uint `gorm:"primaryKey"`
}

// Структура Car
type Car struct {
	Model
	Name     string
	Category string
	// Другие поля для машины
}

func ConnectDB() (*gorm.DB, error) {
	// Подключение к PostgreSQL для машин
	dsn := "host=postgres-car user=postgres password=example dbname=car_db port=5432 sslmode=disable"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
