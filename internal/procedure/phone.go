package procedure

import (
	"context"
	"fmt"
	"github.com/iot-thermometer/server/gen"
	"github.com/iot-thermometer/server/internal/service"
	"google.golang.org/grpc/metadata"
)

type Phone interface {
	Add(ctx context.Context, request *gen.AddPhoneRequest) (*gen.AddPhoneResponse, error)
	Remove(ctx context.Context, request *gen.RemovePhoneRequest) (*gen.RemovePhoneResponse, error)
}

type phone struct {
	phoneService service.Phone
	userService  service.User
}

func newPhoneProcedure(phoneService service.Phone, userService service.User) Phone {
	return &phone{phoneService: phoneService, userService: userService}
}

func (p phone) Add(ctx context.Context, request *gen.AddPhoneRequest) (*gen.AddPhoneResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing context metadata")
	}
	token := md.Get("Authorization")[0]
	userID, err := p.userService.DecodeToken(token)
	if err != nil {
		return nil, err
	}
	_, err = p.phoneService.Add(userID, request.PushID, request.PushToken)
	if err != nil {
		return nil, err
	}
	return &gen.AddPhoneResponse{}, nil
}

func (p phone) Remove(ctx context.Context, request *gen.RemovePhoneRequest) (*gen.RemovePhoneResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing context metadata")
	}
	token := md.Get("Authorization")[0]
	userID, err := p.userService.DecodeToken(token)
	if err != nil {
		return nil, err
	}
	err = p.phoneService.Remove(userID, request.PushID)
	if err != nil {
		return nil, err
	}
	return &gen.RemovePhoneResponse{}, nil
}
