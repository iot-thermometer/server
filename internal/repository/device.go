package repository

import (
	"github.com/iot-thermometer/server/internal/model"
	"gorm.io/gorm"
)

type Device interface {
	GetAll() ([]model.Device, error)
	GetByID(id uint) (model.Device, error)
	GetByToken(token string) (model.Device, error)
	Create(device model.Device) (model.Device, error)
	Save(device model.Device) (model.Device, error)
	DeleteByID(id uint) error
}

type device struct {
	db *gorm.DB
}

func newDeviceRepository(db *gorm.DB) Device {
	return &device{
		db: db,
	}
}

func (d device) GetAll() ([]model.Device, error) {
	var devices []model.Device
	result := d.db.Find(&devices)
	if result.Error != nil {
		return nil, result.Error
	}
	return devices, nil
}
func (d device) GetByID(id uint) (model.Device, error) {
	var device model.Device
	result := d.db.First(&device, id)
	if result.Error != nil {
		return model.Device{}, result.Error
	}
	return device, nil
}

func (d device) GetByToken(token string) (model.Device, error) {
	var device model.Device
	result := d.db.Where("token = ?", token).First(&device)
	if result.Error != nil {
		return model.Device{}, result.Error
	}
	return device, nil
}

func (d device) Create(device model.Device) (model.Device, error) {
	result := d.db.Create(&device)
	if result.Error != nil {
		return model.Device{}, result.Error
	}
	return device, nil
}

func (d device) Save(device model.Device) (model.Device, error) {
	result := d.db.Save(&device)
	if result.Error != nil {
		return model.Device{}, result.Error
	}
	return device, nil
}

func (d device) DeleteByID(id uint) error {
	result := d.db.Delete(&model.Device{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
