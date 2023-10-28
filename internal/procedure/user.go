package procedure

import (
	"errors"

	"github.com/iot-thermometer/server/gen"
	"github.com/iot-thermometer/server/internal/dto"
	"github.com/iot-thermometer/server/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type User interface {
	Login(request *gen.LoginRequest) (*gen.LoginResponse, error)
	Register(request *gen.RegisterRequest) (*gen.RegisterResponse, error)
}

type userProcedure struct {
	userService service.User
}

func newUserProcedure(userService service.User) User {
	return &userProcedure{userService: userService}
}

func (u userProcedure) Login(request *gen.LoginRequest) (*gen.LoginResponse, error) {
	token, err := u.userService.Login(request.GetEmail(), request.GetPassword())
	var appError dto.AppError
	if err != nil {
		if errors.As(err, &appError) {
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}
		return nil, err
	}
	return &gen.LoginResponse{Token: token}, nil
}

func (u userProcedure) Register(request *gen.RegisterRequest) (*gen.RegisterResponse, error) {
	token, err := u.userService.Register(request.GetEmail(), request.GetPassword())
	var appError dto.AppError
	if err != nil {
		if errors.As(err, &appError) {
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}
		return nil, err
	}
	return &gen.RegisterResponse{Token: token}, nil
}
