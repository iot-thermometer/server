package service

import (
	"fmt"
	"github.com/iot-thermometer/server/internal/model"
	"github.com/iot-thermometer/server/internal/repository"
)

type Reading interface {
	List(userID, deviceID uint) ([]model.Reading, error)
}

type reading struct {
	userService       User
	deviceService     Device
	readingRepository repository.Reading
}

func (r reading) List(userID, deviceID uint) ([]model.Reading, error) {
	owns, err := r.deviceService.Owns(userID, deviceID)
	if err != nil {
		return nil, err
	}
	if !owns {
		return nil, fmt.Errorf("user %d does not own device %d", userID, deviceID)
	}
	return r.readingRepository.List(deviceID)
}

func newReadingService(userService User, deviceService Device, readingRepository repository.Reading) Reading {
	return &reading{
		userService:       userService,
		deviceService:     deviceService,
		readingRepository: readingRepository,
	}
}
