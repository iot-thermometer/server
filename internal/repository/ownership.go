package repository

import (
	"errors"

	"github.com/iot-thermometer/server/internal/model"
	"gorm.io/gorm"
)

type Ownership interface {
	FindByUserID(userID uint) ([]model.Ownership, error)
	Create(userID uint, deviceID uint, role model.OwnershipRole) (model.Ownership, error)
	Exists(userID uint, deviceID uint) (bool, error)
	Delete(userID uint, deviceID uint) error
}

type ownership struct {
	db *gorm.DB
}

func newOwnershipRepository(db *gorm.DB) Ownership {
	return &ownership{
		db: db,
	}
}

func (o ownership) FindByUserID(userID uint) ([]model.Ownership, error) {
	var ownerships []model.Ownership
	result := o.db.Where("user_id = ?", userID).Find(&ownerships)
	if result.Error != nil {
		return nil, result.Error
	}
	return ownerships, nil
}

func (o ownership) Create(userID uint, deviceID uint, role model.OwnershipRole) (model.Ownership, error) {
	ownership := model.Ownership{
		UserID:   userID,
		DeviceID: deviceID,
		Role:     role,
	}
	result := o.db.Create(&ownership)
	if result.Error != nil {
		return model.Ownership{}, result.Error
	}
	return ownership, nil
}

func (o ownership) Exists(userID uint, deviceID uint) (bool, error) {
	var ownership model.Ownership
	result := o.db.Where("user_id = ? AND device_id = ?", userID, deviceID).First(&ownership)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}

func (o ownership) Delete(userID uint, deviceID uint) error {
	result := o.db.Where("user_id = ? AND device_id = ?", userID, deviceID).Delete(&model.Ownership{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
