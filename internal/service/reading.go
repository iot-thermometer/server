package service

import (
	"encoding/json"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/iot-thermometer/server/internal/dto"
	"github.com/iot-thermometer/server/internal/model"
	"github.com/iot-thermometer/server/internal/repository"
	"github.com/sirupsen/logrus"
)

type Reading interface {
	List(userID, deviceID uint) ([]model.Reading, error)
	Handle(message mqtt.Message) error
	Channel() chan model.Reading
}

type reading struct {
	userService       User
	deviceService     Device
	deviceRepository  repository.Device
	readingRepository repository.Reading

	readings chan model.Reading
}

func newReadingService(userService User, deviceService Device, deviceRepository repository.Device, readingRepository repository.Reading) Reading {
	return &reading{
		userService:       userService,
		deviceService:     deviceService,
		deviceRepository:  deviceRepository,
		readingRepository: readingRepository,
		readings:          make(chan model.Reading),
	}
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

func (r reading) Handle(message mqtt.Message) error {
	var msg dto.TopicMessage

	err := json.Unmarshal(message.Payload(), &msg)
	if err != nil {
		return err
	}

	reading := model.Reading{
		DeviceID:   msg.DeviceID,
		Value:      msg.Value,
		Type:       msg.Type,
		MeasuredAt: time.Unix(msg.Time, 0),
		UploadedAt: time.Now(),
	}
	_, err = r.readingRepository.Create(reading)
	if err != nil {
		return err
	}

	device, err := r.deviceRepository.GetByID(msg.DeviceID)
	if err != nil {
		return err
	}
	device.RecentlySeenAt = time.Now()
	_, err = r.deviceRepository.Save(device)
	if err != nil {
		return err
	}

	r.readings <- reading

	logrus.Info("Received reading: ", reading)

	return nil
}

func (r *reading) Channel() chan model.Reading {
	return r.readings
}
