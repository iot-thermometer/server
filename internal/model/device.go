package model

import (
	"time"
)

type Device struct {
	ID              uint64 `gorm:"primarykey"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Name            string `gorm:"not null"`
	RecentlySeenAt  time.Time
	Token           string `gorm:"not null"`
	ReadingInterval int    `gorm:"not null"`
	PushInterval    int    `gorm:"not null"`
}
