package services

import (
	"car-service/models"

	"gorm.io/gorm"
)

type CarService struct {
	DB *gorm.DB
}

func NewCarService(db *gorm.DB) *CarService {
	return &CarService{DB: db}
}

func (s *CarService) ListCars() ([]models.Car, error) {
	var cars []models.Car
	err := s.DB.Find(&cars).Error
	return cars, err
}

func (s *CarService) CreateCar(car *models.Car) error {
	return s.DB.Create(car).Error
}

func (s *CarService) GetCar(id uint) (models.Car, error) {
	var car models.Car
	err := s.DB.First(&car, id).Error
	return car, err
}
