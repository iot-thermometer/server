package model

import (
	"time"

	"gorm.io/gorm"
)

type Reading struct {
	gorm.Model
	DeviceID   uint
	Value      float64 `gorm:"not null"`
	Type       string  `gorm:"not null"`
	MeasuredAt time.Time
	UploadedAt time.Time
}
