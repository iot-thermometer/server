package repository

import (
	"github.com/iot-thermometer/server/internal/model"
	"gorm.io/gorm"
)

type Phone interface {
	GetByUserIDAndPushID(userID uint, pushID string) (model.Phone, error)
	Update(phone model.Phone) (model.Phone, error)
	Create(phone model.Phone) (model.Phone, error)
	Delete(phone model.Phone) error
}

type phone struct {
	db *gorm.DB
}

func newPhoneRepository(db *gorm.DB) Phone {
	return phone{db}
}

func (p phone) GetByUserIDAndPushID(userID uint, pushID string) (model.Phone, error) {
	var exists bool
	err := p.db.Model(model.Phone{}).
		Select("count(*) > 0").
		Where("user_id = ? and push_id = ?", userID, pushID).
		Find(&exists).
		Error
	if err != nil {
		return model.Phone{}, err
	}
	if !exists {
		return model.Phone{}, nil
	}

	var phone model.Phone
	tx := p.db.Where("user_id = ? and push_id = ?", userID, pushID).First(&phone)
	if tx.Error != nil {
		return model.Phone{}, tx.Error
	}
	return phone, nil
}

func (p phone) Update(phone model.Phone) (model.Phone, error) {
	tx := p.db.Save(&phone)
	return phone, tx.Error
}

func (p phone) Create(phone model.Phone) (model.Phone, error) {
	tx := p.db.Create(&phone)
	return phone, tx.Error
}

func (p phone) Delete(phone model.Phone) error {
	return p.db.Delete(&phone).Error
}
