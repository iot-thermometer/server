package service

import "github.com/iot-thermometer/server/internal/repository"

type Services interface {
	User() User
	Device() Device
	Reading() Reading
}

type services struct {
	userService    User
	deviceService  Device
	readingService Reading
}

func NewServices(repositories repository.Repositories) Services {
	userService := newUserService()
	deviceService := newDeviceService()
	readingService := newReadingService()
	return &services{
		userService:    userService,
		deviceService:  deviceService,
		readingService: readingService,
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
