package service

import (
	"github.com/iot-thermometer/server/internal/dto"
	"github.com/iot-thermometer/server/internal/repository"
)

type Services interface {
	User() User
	Ownership() Ownership
	Device() Device
	Reading() Reading
	Alert() Alert
	Phone() Phone
}

type services struct {
	userService      User
	ownershipService Ownership
	deviceService    Device
	readingService   Reading
	alertService     Alert
	phoneService     Phone
}

func NewServices(repositories repository.Repositories, config dto.Config) Services {
	userService := newUserService(repositories.User(), config)
	ownershipService := newOwnershipService(repositories.Ownership(), repositories.User())
	deviceService := newDeviceService(repositories.Ownership(), repositories.Device())
	alertService := newAlertService(repositories.Alert(), repositories.Phone())
	readingService := newReadingService(userService, deviceService, alertService, repositories.Device(), repositories.Reading())
	phoneService := newPhoneService(repositories.Phone())
	return &services{
		userService:      userService,
		ownershipService: ownershipService,
		deviceService:    deviceService,
		readingService:   readingService,
		alertService:     alertService,
		phoneService:     phoneService,
	}
}

func (s services) User() User {
	return s.userService
}

func (s services) Ownership() Ownership {
	return s.ownershipService
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

func (s services) Phone() Phone {
	return s.phoneService
}
