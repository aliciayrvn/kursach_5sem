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

func (cc *CarController) GetCar(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	car, err := cc.CarService.GetCar(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
		return
	}

	c.JSON(http.StatusOK, car)
}

func (cc *CarController) ListCars(c *gin.Context) {
	cars, err := cc.CarService.ListCars()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cars"})
		return
	}

	c.JSON(http.StatusOK, cars)
}

func (cc *CarController) CreateCar(c *gin.Context) {
	var car models.Car
	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cc.CarService.CreateCar(&car); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create car"})
		return
	}

	c.JSON(http.StatusCreated, car)
}
