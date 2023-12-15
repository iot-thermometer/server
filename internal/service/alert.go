package service

import "github.com/iot-thermometer/server/internal/repository"

type Alert interface {
}

type alert struct {
	alertRepository repository.Alert
}

func newAlertService(alertRepository repository.Alert) Alert {
	return alert{
		alertRepository: alertRepository,
	}
}
