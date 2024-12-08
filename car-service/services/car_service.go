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

func (s *CarService) GetCar(id uint) (*models.Car, error) {
	var car models.Car
	if err := s.DB.First(&car, id).Error; err != nil {
		return nil, err
	}
	return &car, nil
}

func (s *CarService) ListCars() ([]models.Car, error) {
	var cars []models.Car
	if err := s.DB.Find(&cars).Error; err != nil {
		return nil, err
	}
	return cars, nil
}

func (s *CarService) CreateCar(car *models.Car) error {
	return s.DB.Create(car).Error
}
