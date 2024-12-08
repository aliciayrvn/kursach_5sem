package services

import (
	"user-service/models"

	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

// GetUser получает пользователя по ID
func (s *UserService) GetUser(id uint) (*models.User, error) {
	var user models.User
	if err := s.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// ListUsers возвращает список всех пользователей
func (s *UserService) ListUsers() ([]models.User, error) {
	var users []models.User
	if err := s.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// CreateUser создает нового пользователя
func (s *UserService) CreateUser(user *models.User) error {
	return s.DB.Create(user).Error
}
