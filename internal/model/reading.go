package model

import (
	"time"
)

type Reading struct {
	ID         uint64 `gorm:"primarykey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeviceID   uint
	Value      float64 `gorm:"not null"`
	Type       string  `gorm:"not null"`
	MeasuredAt time.Time
	UploadedAt time.Time
}
