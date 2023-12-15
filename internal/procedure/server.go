package procedure

import (
	"context"
	"github.com/iot-thermometer/server/gen"
)

type server struct {
	gen.UnimplementedThermometerServiceServer

	userProcedure    User
	deviceProcedure  Device
	readingProcedure Reading
	alertProcedure   Alert
}

func (s server) Login(ctx context.Context, request *gen.LoginRequest) (*gen.LoginResponse, error) {
	return s.userProcedure.Login(request)
}

func (s server) Register(ctx context.Context, request *gen.RegisterRequest) (*gen.RegisterResponse, error) {
	return s.userProcedure.Register(request)
}

func (s server) ListDevices(ctx context.Context, request *gen.ListDevicesRequest) (*gen.ListDevicesResponse, error) {
	return s.deviceProcedure.GetDevices(ctx, request)
}

func (s server) CreateDevice(ctx context.Context, request *gen.CreateDeviceRequest) (*gen.CreateDeviceResponse, error) {
	return s.deviceProcedure.AddDevice(ctx, request)
}

func (s server) UpdateDevice(ctx context.Context, request *gen.UpdateDeviceRequest) (*gen.UpdateDeviceResponse, error) {
	return s.deviceProcedure.UpdateDevice(ctx, request)
}

func (s server) DeleteDevice(ctx context.Context, request *gen.DeleteDeviceRequest) (*gen.DeleteDeviceResponse, error) {
	return s.deviceProcedure.DeleteDevice(ctx, request)
}

func (s server) ListReadings(ctx context.Context, request *gen.ListReadingsRequest) (*gen.ListReadingsResponse, error) {
	return s.readingProcedure.List(ctx, request)
}

func (s server) ListAlerts(ctx context.Context, request *gen.ListAlertsRequest) (*gen.ListAlertsResponse, error) {
	return s.alertProcedure.List(ctx, request)
}

func (s server) CreateAlert(ctx context.Context, request *gen.CreateAlertRequest) (*gen.CreateAlertResponse, error) {
	return s.alertProcedure.Create(ctx, request)
}

func (s server) UpdateAlert(ctx context.Context, request *gen.UpdateAlertRequest) (*gen.UpdateAlertResponse, error) {
	return s.alertProcedure.Update(ctx, request)
}

func (s server) DeleteAlert(ctx context.Context, request *gen.DeleteAlertRequest) (*gen.DeleteAlertResponse, error) {
	return s.alertProcedure.Delete(ctx, request)
}
