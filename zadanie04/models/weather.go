package models

import "gorm.io/gorm"

type Weather struct {
	gorm.Model
	City        string  `json:"city"`
	Condition   string  `json:"condition"`
	Temperature float64 `json:"temperature"`
}
