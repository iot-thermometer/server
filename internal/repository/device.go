package repository

import "gorm.io/gorm"

type Device interface {
}
type device struct {
	db *gorm.DB
}

func newDeviceRepository(db *gorm.DB) Device {
	return &device{
		db: db,
	}
}
