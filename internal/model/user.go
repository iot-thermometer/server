package model

import (
	"time"
)

type User struct {
	ID        uint64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string `gorm:"uniqueIndex;not null"`
	Password  string `gorm:"not null"`
}
