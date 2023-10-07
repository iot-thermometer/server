package model

import "gorm.io/gorm"

type Ownership struct {
	gorm.Model
	UserID   uint
	DeviceID uint
}
