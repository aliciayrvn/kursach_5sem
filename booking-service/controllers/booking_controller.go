package controllers

import (
	"booking-service/models"
	"booking-service/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookingController struct {
	BookingService *services.BookingService
}

func NewBookingController(bookingService *services.BookingService) *BookingController {
	return &BookingController{BookingService: bookingService}
}

// ListBookings возвращает список всех бронирований
func (ctrl *BookingController) ListBookings(c *gin.Context) {
	bookings, err := ctrl.BookingService.ListBookings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bookings"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"bookings": bookings})
}

// CreateBooking создает новое бронирование
func (ctrl *BookingController) CreateBooking(c *gin.Context) {
	var input models.Booking
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.BookingService.CreateBooking(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"booking": input})
}

// GetBooking возвращает информацию о бронировании по ID
func (ctrl *BookingController) GetBooking(c *gin.Context) {
	idParam := c.Param("id")
	idInt, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
		return
	}

	id := uint(idInt)

	booking, err := ctrl.BookingService.GetBooking(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"booking": booking})
}
