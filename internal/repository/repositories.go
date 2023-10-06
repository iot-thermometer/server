package repository

import (
	"github.com/iot-thermometer/server/internal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repositories interface {
	User() User
	Device() Device
	Reading() Reading
}

type repositories struct {
	userRepository    User
	deviceRepository  Device
	readingRepository Reading
}

func NewRepositories(db *gorm.DB) Repositories {
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		logrus.Panic(err)
	}
	err = db.AutoMigrate(&model.Device{})
	if err != nil {
		logrus.Panic(err)
	}
	err = db.AutoMigrate(&model.Reading{})
	if err != nil {
		logrus.Panic(err)
	}

	userRepository := newUserRepository(db)
	deviceRepository := newDeviceRepository(db)
	readingRepository := newReadingRepository(db)
	return &repositories{
		userRepository:    userRepository,
		deviceRepository:  deviceRepository,
		readingRepository: readingRepository,
	}
}

func (r repositories) User() User {
	return r.userRepository
}

func (r repositories) Device() Device {
	return r.deviceRepository
}

func (r repositories) Reading() Reading {
	return r.readingRepository
}
