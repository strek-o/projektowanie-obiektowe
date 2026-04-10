package controllers

import (
	"echo-weather/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type WeatherController struct {
	DB *gorm.DB
}

func (wc *WeatherController) GetWeather(c echo.Context) error {
	var weathers []models.Weather

	result := wc.DB.Find(&weathers)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, "NotFound")
	}

	return c.JSON(http.StatusOK, weathers)
}
