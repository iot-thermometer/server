package controller

import (
	"github.com/iot-thermometer/server/internal/dto"
	"github.com/iot-thermometer/server/internal/service"
	"github.com/labstack/echo/v4"
)

type User interface {
	Login(c echo.Context) error
	Register(c echo.Context) error
}

type userController struct {
	userService service.User
}

func newUserController(userService service.User) User {
	return &userController{
		userService: userService,
	}
}

func (u userController) Login(c echo.Context) error {
	var loginRequest dto.LoginRequest
	if err := c.Bind(&loginRequest); err != nil {
		return err
	}
	token, err := u.userService.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return err
	}
	return c.JSON(200, dto.LoginResponse{Token: token})
}

func (u userController) Register(c echo.Context) error {
	var registerRequest dto.RegisterRequest
	if err := c.Bind(&registerRequest); err != nil {
		return err
	}
	token, err := u.userService.Register(registerRequest.Email, registerRequest.Password)
	if err != nil {
		return err
	}
	return c.JSON(200, dto.RegisterResponse{Token: token})
}
