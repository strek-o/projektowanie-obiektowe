package main

import (
	"echo-weather/controllers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	weatherController := &controllers.WeatherController{}

	e.GET("/weather", weatherController.GetWeather)
	e.POST("/weather", weatherController.GetWeather)

	e.Logger.Fatal(e.Start(":8080"))
}
