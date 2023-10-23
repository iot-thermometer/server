package procedure

import (
	"context"
	"github.com/iot-thermometer/server/gen"
)

type server struct {
	gen.UnimplementedThermometerServiceServer
}

func (s server) Login(ctx context.Context, request *gen.LoginRequest) (*gen.LoginResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s server) Register(ctx context.Context, request *gen.RegisterRequest) (*gen.RegisterResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s server) GetDevices(ctx context.Context, request *gen.GetDevicesRequest) (*gen.GetDevicesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s server) AddDevice(ctx context.Context, request *gen.AddDeviceRequest) (*gen.AddDeviceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s server) UpdateDevice(ctx context.Context, request *gen.UpdateDeviceRequest) (*gen.UpdateDeviceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s server) GetReadings(request *gen.GetReadingsRequest, readingsServer gen.ThermometerService_GetReadingsServer) error {
	//TODO implement me
	panic("implement me")
}

func (s server) GetConfig(ctx context.Context, request *gen.GetConfigRequest) (*gen.GetConfigResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s server) mustEmbedUnimplementedThermometerServiceServer() {
	//TODO implement me
	panic("implement me")
}
