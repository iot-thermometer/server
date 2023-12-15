package service

import (
	"github.com/iot-thermometer/server/internal/dto"
	"github.com/iot-thermometer/server/internal/repository"
)

type Services interface {
	User() User
	Device() Device
	Reading() Reading
	Alert() Alert
}

type services struct {
	userService    User
	deviceService  Device
	readingService Reading
	alertService   Alert
}

func NewServices(repositories repository.Repositories, config dto.Config) Services {
	userService := newUserService(repositories.User(), config)
	deviceService := newDeviceService(repositories.Ownership(), repositories.Device())
	readingService := newReadingService(userService, deviceService, repositories.Device(), repositories.Reading())
	alertService := newAlertService(repositories.Alert())
	return &services{
		userService:    userService,
		deviceService:  deviceService,
		readingService: readingService,
		alertService:   alertService,
	}
}

func (s services) User() User {
	return s.userService
}

func (s services) Device() Device {
	return s.deviceService
}

func (s services) Reading() Reading {
	return s.readingService
}

func (s services) Alert() Alert {
	return s.alertService
}
