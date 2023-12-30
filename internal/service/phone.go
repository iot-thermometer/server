package service

import (
	"fmt"
	"github.com/iot-thermometer/server/internal/model"
	"github.com/iot-thermometer/server/internal/repository"
)

type Phone interface {
	Add(userID uint, pushID, pushToken string) (model.Phone, error)
	Remove(userID uint, pushID string) error
}

type phone struct {
	phoneRepository repository.Phone
}

func newPhoneService(phoneRepository repository.Phone) Phone {
	return &phone{phoneRepository}
}

func (p phone) Add(userID uint, pushID, pushToken string) (model.Phone, error) {
	phone, err := p.phoneRepository.GetByUserIDAndPushID(userID, pushID)
	if err != nil {
		return model.Phone{}, err
	}

	if phone.ID == 0 {
		return p.phoneRepository.Create(model.Phone{
			UserID:            userID,
			FirebasePushID:    pushID,
			FirebasePushToken: pushToken,
		})
	} else {
		phone.FirebasePushToken = pushToken
		return p.phoneRepository.Update(phone)
	}
}

func (p phone) Remove(userID uint, pushID string) error {
	phone, err := p.phoneRepository.GetByUserIDAndPushID(userID, pushID)
	if err != nil {
		return err
	}
	if phone.ID > 0 {
		return p.phoneRepository.Delete(phone)
	}
	return fmt.Errorf("phone not found")
}
