package controller

import (
	"github.com/iot-thermometer/server/internal/service"
	"github.com/labstack/echo/v4"
)

type Reading interface {
	List(c echo.Context) error
}

type reading struct {
	readingService service.Reading
}

func (r reading) List(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func newReadingController(readingService service.Reading) Reading {
	return &reading{
		readingService: readingService,
	}
}
