package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type WeatherController struct{}

func (wc *WeatherController) GetWeather(c echo.Context) error {
	mockedWeather := map[string]interface{}{
		"city":        "Krakow",
		"condition":   "Windy",
		"temperature": 12.5,
	}

	return c.JSON(http.StatusOK, mockedWeather)
}
