package repository

import (
	"github.com/iot-thermometer/server/internal/model"
	"gorm.io/gorm"
)

type Reading interface {
	List(deviceID uint) ([]model.Reading, error)
	Create(reading model.Reading) (model.Reading, error)
}

type reading struct {
	db *gorm.DB
}

func newReadingRepository(db *gorm.DB) Reading {
	return &reading{
		db: db,
	}
}

func (r reading) List(deviceID uint) ([]model.Reading, error) {
	readings := make([]model.Reading, 0)
	if err := r.db.Where("device_id = ?", deviceID).Find(&readings).Error; err != nil {
		return nil, err
	}
	return readings, nil
}

func (r reading) Create(reading model.Reading) (model.Reading, error) {
	if err := r.db.Create(&reading).Error; err != nil {
		return model.Reading{}, err
	}
	return reading, nil
}
