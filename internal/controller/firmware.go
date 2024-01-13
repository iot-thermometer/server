package controller

import (
	"github.com/labstack/echo/v4"
)

type Firmware interface {
	Index(c echo.Context) error
	Download(c echo.Context) error
}

type firmware struct {
}

func newFirmware() Firmware {
	return &firmware{}
}

func (f firmware) Index(c echo.Context) error {
	return c.File("/firmware/version.json")
}

func (f firmware) Download(c echo.Context) error {
	return c.File("/firmware/iot.bin")
}
