package service

import (
	"fmt"
	"github.com/iot-thermometer/server/internal/client"

	"github.com/iot-thermometer/server/internal/dto"
	"github.com/iot-thermometer/server/internal/model"
	"github.com/iot-thermometer/server/internal/repository"
	"github.com/iot-thermometer/server/internal/util"
)

type Device interface {
	List(userID uint) ([]model.Device, error)
	Create(userID uint, payload dto.CreateDeviceRequest) (model.Device, error)
	Update(userID uint, deviceID uint, payload dto.UpdateDeviceRequest) (model.Device, error)
	Delete(userID uint, deviceID uint) error
	Owns(userID uint, deviceID uint) (bool, error)

	Get(token string) (model.Device, error)
	Watch() error
}

type device struct {
	ownershipRepository repository.Ownership
	deviceRepository    repository.Device
	broker              client.Broker
}

func (d device) List(userID uint) ([]model.Device, error) {
	ownerships, err := d.ownershipRepository.FindByUserID(userID)
	if err != nil {
		return nil, err
	}
	devices := make([]model.Device, 0)
	for _, ownership := range ownerships {
		device, err := d.deviceRepository.GetByID(ownership.DeviceID)
		if err != nil {
			return nil, err
		}
		devices = append(devices, device)
	}
	return devices, nil
}

func (d device) Create(userID uint, payload dto.CreateDeviceRequest) (model.Device, error) {
	device := model.Device{
		Name:            payload.Name,
		Token:           util.RandStringRunes(16),
		ReadingInterval: payload.ReadingInterval,
		PushInterval:    payload.PushInterval,
	}
	result, err := d.deviceRepository.Create(device)
	if err != nil {
		return model.Device{}, err
	}
	_, err = d.ownershipRepository.Create(userID, result.ID, model.OwnershipRoleOwner)
	if err != nil {
		return model.Device{}, err
	}
	err = d.watch(result)
	if err != nil {
		return model.Device{}, err
	}

	return result, nil
}

func (d device) Update(userID uint, deviceID uint, payload dto.UpdateDeviceRequest) (model.Device, error) {
	owns, err := d.ownershipRepository.Exists(userID, deviceID)
	if err != nil {
		return model.Device{}, err
	}
	if !owns {
		return model.Device{}, fmt.Errorf("user %d does not own device %d", userID, deviceID)
	}
	device, err := d.deviceRepository.GetByID(deviceID)
	if err != nil {
		return model.Device{}, err
	}
	device.Name = payload.Name
	device.ReadingInterval = payload.ReadingInterval
	device.PushInterval = payload.PushInterval
	result, err := d.deviceRepository.Save(device)
	if err != nil {
		return model.Device{}, err
	}
	return result, nil
}

func (d device) Delete(userID uint, deviceID uint) error {
	owns, err := d.ownershipRepository.Exists(userID, deviceID)
	if err != nil {
		return err
	}
	if !owns {
		return fmt.Errorf("user %d does not own device %d", userID, deviceID)
	}
	err = d.ownershipRepository.Delete(userID, deviceID)
	if err != nil {
		return err
	}
	return d.deviceRepository.DeleteByID(deviceID)
}

func (d device) Owns(userID uint, deviceID uint) (bool, error) {
	return d.ownershipRepository.Exists(userID, deviceID)
}

func (d device) Get(token string) (model.Device, error) {
	return d.deviceRepository.GetByToken(token)
}

func (d device) watch(device model.Device) error {
	err := d.broker.Subscribe(fmt.Sprintf("sensors/%d/TEMP", device.ID))
	if err != nil {
		return nil
	}
	err = d.broker.Subscribe(fmt.Sprintf("sensors/%d/SOIL", device.ID))
	if err != nil {
		return nil
	}
	return nil
}

func (d device) Watch() error {
	devices, err := d.deviceRepository.GetAll()
	if err != nil {
		panic(err)
	}
	for _, device := range devices {
		err = d.watch(device)
		if err != nil {
			panic(err)
		}
	}
	return nil
}

func newDeviceService(ownershipRepository repository.Ownership, deviceRepository repository.Device, broker client.Broker) Device {
	return &device{
		ownershipRepository: ownershipRepository,
		deviceRepository:    deviceRepository,
		broker:              broker,
	}
}
