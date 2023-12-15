package repository

import "gorm.io/gorm"

type Alert interface {
}

type alert struct {
	db *gorm.DB
}

func newAlertRepository(db *gorm.DB) Alert {
	return &alert{db}
}
