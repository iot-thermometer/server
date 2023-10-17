package controller

import (
	"github.com/iot-thermometer/server/internal/dto"
	"github.com/labstack/echo/v4"
	"net/http"
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
	return c.JSON(http.StatusOK, dto.FirmwareIndexResponse{
		Version: 0,
		Source:  "",
	})
}

func (f firmware) Download(c echo.Context) error {
	return c.Blob(http.StatusOK, "application/octet-stream", []byte{})
}
