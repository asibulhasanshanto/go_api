package model

import "gorm.io/gorm"

// Vehicle struct
type Vehicle struct {
	gorm.Model
	Brand        string `json:"brand"`
	VehicleModel string `json:"model"`
	Type         string `json:"type"`
}
