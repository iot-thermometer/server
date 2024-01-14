package model

import (
	"gorm.io/gorm"
	"time"
)

type Alert struct {
	gorm.Model
	AlertName       string    `gorm:"type:varchar(100);not null" json:"alert_name"`
	UserID          uint      `gorm:"not null" json:"user_id"`
	DeviceID        uint      `gorm:"not null" json:"device_id"`
	SoilMoistureMin float64   `gorm:"not null" json:"soil_moisture_min"`
	SoilMoistureMax float64   `gorm:"not null" json:"soil_moisture_max"`
	TemperatureMin  float64   `gorm:"not null" json:"temperature_min"`
	TemperatureMax  float64   `gorm:"not null" json:"temperature_max"`
	LastSentAt      time.Time `json:"last_sent_at"`
}
