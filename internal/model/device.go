package model

import (
	"time"

	"gorm.io/gorm"
)

type Device struct {
	gorm.Model
	Name            string `gorm:"not null"`
	RecentlySeenAt  time.Time
	Token           string `gorm:"not null"`
	ReadingInterval int    `gorm:"not null"`
	PushInterval    int    `gorm:"not null"`
}
