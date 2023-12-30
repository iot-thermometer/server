package service

import (
	"github.com/iot-thermometer/server/internal/dto"
	"github.com/iot-thermometer/server/internal/model"
	"github.com/iot-thermometer/server/internal/repository"
)

type Alert interface {
	List(userID uint) ([]model.Alert, error)
	Create(userID uint, payload dto.CreateAlertRequest) (model.Alert, error)
	Update(userID, alertID uint, payload dto.UpdateAlertRequest) (model.Alert, error)
	Delete(userID, alertID uint) error
	Owns(userID uint, alertID uint) (bool, error)

	Check(reading model.Reading) error
}

type alert struct {
	alertRepository repository.Alert
}

func newAlertService(alertRepository repository.Alert) Alert {
	return alert{
		alertRepository: alertRepository,
	}
}

func (a alert) List(userID uint) ([]model.Alert, error) {
	return a.alertRepository.GetByUserID(userID)
}

func (a alert) Create(userID uint, payload dto.CreateAlertRequest) (model.Alert, error) {
	return a.alertRepository.Create(model.Alert{
		AlertName:       payload.Name,
		UserID:          userID,
		DeviceID:        payload.DeviceID,
		SoilMoistureMin: payload.SoilMoistureMin,
		SoilMoistureMax: payload.SoilMoistureMax,
		TemperatureMin:  payload.TemperatureMin,
		TemperatureMax:  payload.TemperatureMax,
		PushToken:       payload.PushToken,
	})
}

func (a alert) Update(userID, alertID uint, payload dto.UpdateAlertRequest) (model.Alert, error) {
	alert, err := a.alertRepository.GetByUserIDAndID(userID, alertID)
	if err != nil {
		return model.Alert{}, err
	}
	alert.AlertName = payload.Name
	alert.SoilMoistureMin = payload.SoilMoistureMin
	alert.SoilMoistureMax = payload.SoilMoistureMax
	alert.TemperatureMin = payload.TemperatureMin
	alert.TemperatureMax = payload.TemperatureMax
	alert.PushToken = payload.PushToken
	alert, err = a.alertRepository.Save(alert)
	if err != nil {
		return model.Alert{}, err
	}
	return alert, nil
}

func (a alert) Delete(userID, alertID uint) error {
	_, err := a.alertRepository.GetByUserIDAndID(userID, alertID)
	if err != nil {
		return err
	}
	return a.alertRepository.DeleteByID(alertID)
}

func (a alert) Owns(userID uint, alertID uint) (bool, error) {
	_, err := a.alertRepository.GetByUserIDAndID(userID, alertID)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a alert) Check(reading model.Reading) error {
	alerts, err := a.alertRepository.GetAll()
	if err != nil {
		return err
	}

	for _, alert := range alerts {
		if alert.DeviceID == reading.DeviceID {
			if reading.SoilMoisture > alert.SoilMoistureMin && reading.SoilMoisture < alert.SoilMoistureMax &&
				reading.Temperature > alert.TemperatureMin && reading.Temperature < alert.TemperatureMax {
				// TODO send notification
			}
		}
	}
	return nil
}
