package repository

import (
	"github.com/iot-thermometer/server/internal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repositories interface {
	User() User
	Ownership() Ownership
	Device() Device
	Reading() Reading
	Alert() Alert
}

type repositories struct {
	userRepository      User
	ownershipRepository Ownership
	deviceRepository    Device
	readingRepository   Reading
	alertRepository     Alert
}

func NewRepositories(db *gorm.DB) Repositories {
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		logrus.Panic(err)
	}
	err = db.AutoMigrate(&model.Ownership{})
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
	err = db.AutoMigrate(&model.Alert{})
	if err != nil {
		logrus.Panic(err)
	}

	userRepository := newUserRepository(db)
	ownershipRepository := newOwnershipRepository(db)
	deviceRepository := newDeviceRepository(db)
	readingRepository := newReadingRepository(db)
	alertRepository := newAlertRepository(db)
	return &repositories{
		userRepository:      userRepository,
		ownershipRepository: ownershipRepository,
		deviceRepository:    deviceRepository,
		readingRepository:   readingRepository,
		alertRepository:     alertRepository,
	}
}

func (r repositories) User() User {
	return r.userRepository
}

func (r repositories) Ownership() Ownership {
	return r.ownershipRepository
}

func (r repositories) Device() Device {
	return r.deviceRepository
}

func (r repositories) Reading() Reading {
	return r.readingRepository
}

func (r repositories) Alert() Alert {
	return r.alertRepository
}
