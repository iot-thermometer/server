package procedure

import (
	"context"
	"fmt"

	"github.com/iot-thermometer/server/gen"
	"github.com/iot-thermometer/server/internal/dto"
	"github.com/iot-thermometer/server/internal/service"
	"github.com/iot-thermometer/server/internal/util"
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
		userService:   userService,
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
		fmt.Println(device.ID, device.Name)
		pbDevices = append(pbDevices, &gen.Device{
			Id:              util.Uint64(uint64(device.ID)),
			Name:            &device.Name,
			RecentlySeenAt:  util.Int64(device.RecentlySeenAt.Unix()),
			Token:           &device.Token,
			ReadingInterval: util.Int64(int64(device.ReadingInterval)),
			PushInterval:    util.Int64(int64(device.PushInterval)),
		})
	}
	return &gen.ListDevicesResponse{
		Devices: pbDevices,
	}, nil
}

func (d deviceProcedure) AddDevice(ctx context.Context, request *gen.CreateDeviceRequest) (*gen.CreateDeviceResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing context metadata")
	}
	token := md.Get("Authorization")[0]
	userID, err := d.userService.DecodeToken(token)
	if err != nil {
		return nil, err
	}
	device, err := d.deviceService.Create(userID, dto.CreateDeviceRequest{
		Name:            request.Name,
		ReadingInterval: int(request.ReadingInterval),
		PushInterval:    int(request.PushInterval),
	})
	if err != nil {
		return nil, err
	}
	return &gen.CreateDeviceResponse{
		Device: &gen.Device{
			Id:              util.Uint64(uint64(device.ID)),
			Name:            &device.Name,
			RecentlySeenAt:  util.Int64(device.RecentlySeenAt.Unix()),
			Token:           &device.Token,
			ReadingInterval: util.Int64(int64(device.ReadingInterval)),
			PushInterval:    util.Int64(int64(device.PushInterval)),
		},
	}, nil
}

func (d deviceProcedure) UpdateDevice(ctx context.Context, request *gen.UpdateDeviceRequest) (*gen.UpdateDeviceResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing context metadata")
	}
	token := md.Get("Authorization")[0]
	userID, err := d.userService.DecodeToken(token)
	if err != nil {
		return nil, err
	}
	device, err := d.deviceService.Update(userID, uint(request.Id), dto.UpdateDeviceRequest{
		Name:            request.Name,
		ReadingInterval: int(request.ReadingInterval),
		PushInterval:    int(request.PushInterval),
	})
	if err != nil {
		return nil, err
	}

	return &gen.UpdateDeviceResponse{
		Device: &gen.Device{
			Id:              util.Uint64(uint64(device.ID)),
			Name:            &device.Name,
			RecentlySeenAt:  util.Int64(device.RecentlySeenAt.Unix()),
			Token:           &device.Token,
			ReadingInterval: util.Int64(int64(device.ReadingInterval)),
			PushInterval:    util.Int64(int64(device.PushInterval)),
		},
	}, nil
}

func (d deviceProcedure) DeleteDevice(ctx context.Context, request *gen.DeleteDeviceRequest) (*gen.DeleteDeviceResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing context metadata")
	}
	token := md.Get("Authorization")[0]
	userID, err := d.userService.DecodeToken(token)
	if err != nil {
		return nil, err
	}
	err = d.deviceService.Delete(userID, uint(request.Id))
	if err != nil {
		return nil, err
	}
	return &gen.DeleteDeviceResponse{}, nil
}
