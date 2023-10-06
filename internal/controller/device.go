package controller

import "github.com/iot-thermometer/server/internal/service"

type Device interface {
}

type device struct {
	deviceService service.Device
}

func newDeviceController(deviceService service.Device) Device {
	return &device{
		deviceService: deviceService,
	}
}
