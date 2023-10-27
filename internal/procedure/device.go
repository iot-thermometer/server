package procedure

import (
	"context"
	"fmt"
	"github.com/iot-thermometer/server/gen"
	"github.com/iot-thermometer/server/internal/service"
	"google.golang.org/grpc/metadata"
)

type Device interface {
	GetDevices(ctx context.Context, request *gen.ListDevicesRequest) (*gen.ListDevicesResponse, error)
	AddDevice(ctx context.Context, request *gen.CreateDeviceRequest) (*gen.CreateDeviceResponse, error)
	UpdateDevice(ctx context.Context, request *gen.UpdateDeviceRequest) (*gen.UpdateDeviceResponse, error)
	DeleteDevice(ctx context.Context, request *gen.DeleteDeviceRequest) (*gen.DeleteDeviceResponse, error)
}

type deviceProcedure struct {
	deviceService service.Device
	userService   service.User
}

func newDeviceProcedure(deviceService service.Device, userService service.User) Device {
	return &deviceProcedure{
		deviceService: deviceService,
	}
}

func (d deviceProcedure) GetDevices(ctx context.Context, request *gen.ListDevicesRequest) (*gen.ListDevicesResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing context metadata")
	}
	token := md.Get("Authorization")[0]
	userID, err := d.userService.DecodeToken(token)
	if err != nil {
		return nil, err
	}
	devices, err := d.deviceService.List(userID)
	if err != nil {
		return nil, err
	}

	var pbDevices []*gen.Device
	for _, device := range devices {
		pbDevices = append(pbDevices, &gen.Device{
			Id:   &device.ID,
			Name: &device.Name,
		})
	}
	return &gen.ListDevicesResponse{
		Devices: pbDevices,
	}, nil
}

func (d deviceProcedure) AddDevice(ctx context.Context, request *gen.CreateDeviceRequest) (*gen.CreateDeviceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (d deviceProcedure) UpdateDevice(ctx context.Context, request *gen.UpdateDeviceRequest) (*gen.UpdateDeviceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (d deviceProcedure) DeleteDevice(ctx context.Context, request *gen.DeleteDeviceRequest) (*gen.DeleteDeviceResponse, error) {
	//TODO implement me
	panic("implement me")
}
