package repository

import (
	"github.com/iot-thermometer/server/internal/model"
	"gorm.io/gorm"
)

type Alert interface {
	GetByUserID(userID uint) ([]model.Alert, error)
	GetByUserIDAndID(userID, id uint) (model.Alert, error)
	Create(alert model.Alert) (model.Alert, error)
	Save(alert model.Alert) (model.Alert, error)
	DeleteByID(id uint) error
}

type alert struct {
	db *gorm.DB
}

func (a alert) GetByUserID(userID uint) ([]model.Alert, error) {
	alerts := make([]model.Alert, 0)
	if err := a.db.Where("user_id = ?", userID).Find(&alerts).Error; err != nil {
		return nil, err
	}
	return alerts, nil
}

func (a alert) GetByUserIDAndID(userID, id uint) (model.Alert, error) {
	var alert model.Alert
	if err := a.db.Where("user_id = ? and id = ?", userID, id).Find(&alert).Error; err != nil {
		return model.Alert{}, err
	}
	return alert, nil
}

func (a alert) Create(alert model.Alert) (model.Alert, error) {
	result := a.db.Create(&alert)
	if result.Error != nil {
		return model.Alert{}, result.Error
	}
	return alert, nil
}

func (a alert) Save(alert model.Alert) (model.Alert, error) {
	result := a.db.Save(&alert)
	if result.Error != nil {
		return model.Alert{}, result.Error
	}
	return alert, nil
}

func (a alert) DeleteByID(id uint) error {
	result := a.db.Delete(&model.Alert{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func newAlertRepository(db *gorm.DB) Alert {
	return &alert{db}
}
