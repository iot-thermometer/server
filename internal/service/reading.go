package service

import (
	"crypto/aes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/andreburgaud/crypt2go/ecb"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/iot-thermometer/server/internal/dto"
	"github.com/iot-thermometer/server/internal/model"
	"github.com/iot-thermometer/server/internal/repository"
	"github.com/sirupsen/logrus"
)

type Reading interface {
	List(userID, deviceID uint, start, end time.Time) ([]model.Reading, error)
	Handle(message mqtt.Message) error
}

type reading struct {
	userService       User
	deviceService     Device
	alertService      Alert
	deviceRepository  repository.Device
	readingRepository repository.Reading

	readings chan model.Reading
}

func newReadingService(userService User, deviceService Device, alertService Alert, deviceRepository repository.Device, readingRepository repository.Reading) Reading {
	return &reading{
		userService:       userService,
		deviceService:     deviceService,
		alertService:      alertService,
		deviceRepository:  deviceRepository,
		readingRepository: readingRepository,
		readings:          make(chan model.Reading),
	}
}

func (r reading) List(userID, deviceID uint, start, end time.Time) ([]model.Reading, error) {
	owns, err := r.deviceService.Owns(userID, deviceID)
	if err != nil {
		return nil, err
	}
	if !owns {
		return nil, fmt.Errorf("user %d does not own device %d", userID, deviceID)
	}
	return r.readingRepository.List(deviceID, start, end)
}

func (r reading) Handle(message mqtt.Message) error {
	payloadString := message.Payload()
	if strings.HasPrefix(string(payloadString), "XDEBUG") {
		return nil
	}
	parts := strings.Split(string(payloadString), ";")

	deviceID, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}
	device, err := r.deviceRepository.GetByID(uint(deviceID))
	if err != nil {
		return err
	}
	now := time.Now()
	device.RecentlySeenAt = &now
	_, err = r.deviceRepository.Save(device)
	if err != nil {
		return err
	}

	cipher, err := aes.NewCipher([]byte(device.Token))
	if err != nil {
		return err
	}

	encrypted, err := hex.DecodeString(parts[1])
	if err != nil {
		return err
	}
	decrypted := make([]byte, 512)
	e := ecb.NewECBDecrypter(cipher)
	e.CryptBlocks(decrypted, encrypted)

	index := strings.Index(string(decrypted), "}")
	j := string(decrypted)[:index+1]

	var msg dto.TopicMessage
	err = json.Unmarshal([]byte(j), &msg)
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

	logrus.Info("Received reading: ", reading)
	go r.alertService.Check(reading)

	return nil
}
