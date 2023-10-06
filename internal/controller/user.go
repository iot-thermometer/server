package controller

import (
	"github.com/iot-thermometer/server/internal/service"
	"github.com/labstack/echo/v4"
)

type User interface {
	Login(c echo.Context) error
	Register(c echo.Context) error
}

type user struct {
	userService service.User
}

func (u user) Login(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (u user) Register(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func newUserController(userService service.User) User {
	return &user{
		userService: userService,
	}
}
