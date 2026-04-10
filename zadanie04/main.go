package main

import (
	"echo-weather/controllers"
	"echo-weather/models"
	"log"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("weather.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("ConnectionError", err)
	}

	db.AutoMigrate(&models.Weather{})

	var count int64
	db.Model(&models.Weather{}).Count(&count)

	if count == 0 {
		initialData := []models.Weather{
			{City: "Cracow", Condition: "Windy", Temperature: 12.5},
			{City: "Warsaw", Condition: "Sunny", Temperature: 15.0},
		}

		for _, w := range initialData {
			db.Create(&w)
		}
	}

	e := echo.New()

	weatherController := &controllers.WeatherController{
		DB: db,
	}

	e.GET("/weather", weatherController.GetWeather)
	e.POST("/weather", weatherController.GetWeather)

	e.Logger.Fatal(e.Start(":8080"))
}
