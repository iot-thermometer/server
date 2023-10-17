package controller

import (
	"github.com/iot-thermometer/server/internal/dto"
	"github.com/iot-thermometer/server/internal/service"
	"github.com/iot-thermometer/server/internal/util"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Device interface {
	List(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error

	Config(c echo.Context) error
}

type device struct {
	userService   service.User
	deviceService service.Device
}

func (d device) List(c echo.Context) error {
	userID, err := d.userService.DecodeToken(util.GetTokenFromContext(c))
	if err != nil {
		return err
	}

	devices, err := d.deviceService.List(userID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, devices)
}

func (d device) Create(c echo.Context) error {
	userID, err := d.userService.DecodeToken(util.GetTokenFromContext(c))
	if err != nil {
		return err
	}

	var payload dto.CreateDeviceRequest
	if err := c.Bind(&payload); err != nil {
		return err
	}

	device, err := d.deviceService.Create(userID, payload)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, device)
}

func (d device) Update(c echo.Context) error {
	userID, err := d.userService.DecodeToken(util.GetTokenFromContext(c))
	if err != nil {
		return err
	}

	var payload dto.UpdateDeviceRequest
	if err := c.Bind(&payload); err != nil {
		return err
	}

	device, err := d.deviceService.Update(userID, util.ParseParamID(c.Param("id")), payload)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, device)
}

func (d device) Delete(c echo.Context) error {
	userID, err := d.userService.DecodeToken(util.GetTokenFromContext(c))
	if err != nil {
		return err
	}

	if err := d.deviceService.Delete(userID, util.ParseParamID(c.Param("id"))); err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}

func (d device) Config(c echo.Context) error {
	device, err := d.deviceService.Get(c.Param("token"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, device)
}

func newDeviceController(userService service.User, deviceService service.Device) Device {
	return &device{
		userService:   userService,
		deviceService: deviceService,
	}
}
