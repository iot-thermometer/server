package controller

import (
	"github.com/iot-thermometer/server/internal/service"
	"github.com/labstack/echo/v4"
)

type Controllers interface {
	User() User
	Device() Device
	Reading() Reading

	Route(e *echo.Echo)
}

type controllers struct {
	userController    User
	deviceController  Device
	readingController Reading
}

func NewControllers(services service.Services) Controllers {
	userController := newUserController(services.User())
	deviceController := newDeviceController(services.Device())
	readingController := newReadingController(services.Reading())
	return &controllers{
		userController:    userController,
		deviceController:  deviceController,
		readingController: readingController,
	}
}

func (c controllers) User() User {
	return c.userController
}

func (c controllers) Device() Device {
	return c.deviceController
}

func (c controllers) Reading() Reading {
	return c.readingController
}

func (c controllers) Route(e *echo.Echo) {
	e.POST("/api/auth/login", c.userController.Login)
	e.POST("/api/auth/register", c.userController.Register)

	e.GET("/api/devices", c.deviceController.List)
	e.POST("/api/devices", c.deviceController.Create)
	e.PUT("/api/devices/:id", c.deviceController.Update)
	e.DELETE("/api/devices/:id", c.deviceController.Delete)

	e.GET("/api/devices/:id/readings", c.readingController.List)
}
