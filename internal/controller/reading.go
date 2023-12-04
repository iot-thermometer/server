package controller

import (
	"time"

	"github.com/iot-thermometer/server/internal/service"
	"github.com/iot-thermometer/server/internal/util"
	"github.com/labstack/echo/v4"
)

type Reading interface {
	List(c echo.Context) error
}

type reading struct {
	userService    service.User
	readingService service.Reading
}

func (r reading) List(c echo.Context) error {
	userID, err := r.userService.DecodeToken(util.GetTokenFromContext(c))
	if err != nil {
		return err
	}
	readings, err := r.readingService.List(userID, util.ParseParamID(c.Param("id")), time.Unix(int64(util.ParseParamID(c.QueryParam("start"))), 0), time.Unix(int64(util.ParseParamID("end")), 0))
	if err != nil {
		return err
	}
	return c.JSON(200, readings)
}

func newReadingController(userService service.User, readingService service.Reading) Reading {
	return &reading{
		userService:    userService,
		readingService: readingService,
	}
}
