package model

import (
	"gorm.io/gorm"
	"time"
)

type Device struct {
	gorm.Model
	Name           string `gorm:"not null"`
	RecentlySeenAt time.Time
	UserID         uint
	Token          string `gorm:"not null"`
}
