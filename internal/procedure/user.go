package procedure

import "github.com/iot-thermometer/server/internal/service"

type User interface {
}

type userProcedure struct {
	userService service.User
}

func newUserProcedure(userService service.User) User {
	return &userProcedure{userService: userService}
}
