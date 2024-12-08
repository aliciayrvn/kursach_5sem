package controllers

import (
	"car-service/models"
	"car-service/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CarController struct {
	CarService *services.CarService
}

func NewCarController(carService *services.CarService) *CarController {
	return &CarController{CarService: carService}
}

func (ctrl *CarController) ListCars(c *gin.Context) {
	cars, err := ctrl.CarService.ListCars()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cars"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"cars": cars})
}

func (ctrl *CarController) CreateCar(c *gin.Context) {
	var input models.Car
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.CarService.CreateCar(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create car"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"car": input})
}

func (ctrl *CarController) GetCar(c *gin.Context) {
	idParam := c.Param("id")
	idInt, err := strconv.Atoi(idParam) // Преобразование строки в int
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid car ID"})
		return
	}

	id := uint(idInt) // Преобразование int в uint

	car, err := ctrl.CarService.GetCar(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"car": car})
}
