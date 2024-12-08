package services

import (
	"booking-service/models"

	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

type BookingService struct {
	DB        *gorm.DB
	RMQClient *amqp.Connection
}

func NewBookingService(db *gorm.DB, rmqClient *amqp.Connection) *BookingService {
	return &BookingService{
		DB:        db,
		RMQClient: rmqClient,
	}
}

// GetBooking получает информацию о бронировании по ID
func (s *BookingService) GetBooking(id uint) (*models.Booking, error) {
	var booking models.Booking
	if err := s.DB.First(&booking, id).Error; err != nil {
		return nil, err
	}
	return &booking, nil
}

// ListBookings возвращает список всех бронирований
func (s *BookingService) ListBookings() ([]models.Booking, error) {
	var bookings []models.Booking
	if err := s.DB.Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}

// CreateBooking создает новое бронирование
func (s *BookingService) CreateBooking(booking *models.Booking) error {
	return s.DB.Create(booking).Error
}
