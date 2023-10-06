package model

import (
	"gorm.io/gorm"
	"time"
)

type Reading struct {
	gorm.Model
	DeviceID   uint
	Value      float64 `gorm:"not null"`
	MeasuredAt time.Time
	UploadedAt time.Time
}
