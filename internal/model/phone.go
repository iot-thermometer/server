package model

import "gorm.io/gorm"

type Phone struct {
	gorm.Model
	UserID            uint
	FirebasePushID    string
	FirebasePushToken string
}
