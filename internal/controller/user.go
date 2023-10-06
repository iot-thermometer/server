package controller

import "github.com/iot-thermometer/server/internal/service"

type User interface {
}

type user struct {
	userService service.User
}

func newUserController(userService service.User) User {
	return &user{
		userService: userService,
	}
}
