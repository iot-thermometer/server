package controller

import (
	"github.com/iot-thermometer/server/internal/service"
	"github.com/labstack/echo/v4"
)

type Device interface {
	List(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type device struct {
	deviceService service.Device
}

func (d device) List(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (d device) Create(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (d device) Update(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (d device) Delete(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func newDeviceController(deviceService service.Device) Device {
	return &device{
		deviceService: deviceService,
	}
}
