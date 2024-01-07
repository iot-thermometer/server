package procedure

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/iot-thermometer/server/gen"
	"github.com/iot-thermometer/server/internal/dto"
	"github.com/iot-thermometer/server/internal/service"
	"github.com/iot-thermometer/server/internal/util"
	"google.golang.org/grpc/metadata"
)

type Alert interface {
	List(ctx context.Context, request *gen.ListAlertsRequest) (*gen.ListAlertsResponse, error)
	Create(ctx context.Context, request *gen.CreateAlertRequest) (*gen.CreateAlertResponse, error)
	Update(ctx context.Context, request *gen.UpdateAlertRequest) (*gen.UpdateAlertResponse, error)
	Delete(ctx context.Context, request *gen.DeleteAlertRequest) (*gen.DeleteAlertResponse, error)
}

type alert struct {
	alertService service.Alert
	userService  service.User
}

func (a alert) List(ctx context.Context, request *gen.ListAlertsRequest) (*gen.ListAlertsResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing context metadata")
	}
	token := md.Get("Authorization")[0]
	userID, err := a.userService.DecodeToken(token)
	if err != nil {
		return nil, err
	}

	alerts, err := a.alertService.List(userID)
	if err != nil {
		return nil, err
	}

	var pbAlerts []*gen.Alert
	for _, alert := range alerts {
		pbAlerts = append(pbAlerts, &gen.Alert{
			Id:              util.Uint64(uint64(alert.ID)),
			Name:            aws.String(alert.AlertName),
			UserID:          util.Uint64(uint64(alert.UserID)),
			DeviceID:        util.Uint64(uint64(alert.DeviceID)),
			SoilMoistureMin: util.Float32(float32(alert.SoilMoistureMin)),
			SoilMoistureMax: util.Float32(float32(alert.SoilMoistureMax)),
			TemperatureMin:  util.Float32(float32(alert.TemperatureMin)),
			TemperatureMax:  util.Float32(float32(alert.TemperatureMax)),
		})
	}
	return &gen.ListAlertsResponse{Alerts: pbAlerts}, nil
}

func (a alert) Create(ctx context.Context, request *gen.CreateAlertRequest) (*gen.CreateAlertResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing context metadata")
	}
	token := md.Get("Authorization")[0]
	userID, err := a.userService.DecodeToken(token)
	if err != nil {
		return nil, err
	}

	alert, err := a.alertService.Create(userID, dto.CreateAlertRequest{
		DeviceID:        uint(request.DeviceID),
		Name:            request.Name,
		SoilMoistureMin: float64(request.SoilMoistureMin),
		SoilMoistureMax: float64(request.SoilMoistureMax),
		TemperatureMin:  float64(request.TemperatureMin),
		TemperatureMax:  float64(request.TemperatureMax),
	})
	if err != nil {
		return nil, err
	}
	return &gen.CreateAlertResponse{Alert: &gen.Alert{
		Id:              util.Uint64(uint64(alert.ID)),
		Name:            aws.String(alert.AlertName),
		UserID:          util.Uint64(uint64(alert.UserID)),
		DeviceID:        util.Uint64(uint64(alert.DeviceID)),
		SoilMoistureMin: util.Float32(float32(alert.SoilMoistureMin)),
		SoilMoistureMax: util.Float32(float32(alert.SoilMoistureMax)),
		TemperatureMin:  util.Float32(float32(alert.TemperatureMin)),
		TemperatureMax:  util.Float32(float32(alert.TemperatureMax)),
	}}, nil
}

func (a alert) Update(ctx context.Context, request *gen.UpdateAlertRequest) (*gen.UpdateAlertResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing context metadata")
	}
	token := md.Get("Authorization")[0]
	userID, err := a.userService.DecodeToken(token)
	if err != nil {
		return nil, err
	}

	alert, err := a.alertService.Update(userID, uint(request.DeviceID), dto.UpdateAlertRequest{
		Name:            request.Name,
		SoilMoistureMin: float64(request.SoilMoistureMin),
		SoilMoistureMax: float64(request.SoilMoistureMax),
		TemperatureMin:  float64(request.TemperatureMin),
		TemperatureMax:  float64(request.TemperatureMax),
	})
	if err != nil {
		return nil, err
	}
	return &gen.UpdateAlertResponse{Alert: &gen.Alert{
		Id:              util.Uint64(uint64(alert.ID)),
		Name:            aws.String(alert.AlertName),
		UserID:          util.Uint64(uint64(alert.UserID)),
		DeviceID:        util.Uint64(uint64(alert.DeviceID)),
		SoilMoistureMin: util.Float32(float32(alert.SoilMoistureMin)),
		SoilMoistureMax: util.Float32(float32(alert.SoilMoistureMax)),
		TemperatureMin:  util.Float32(float32(alert.TemperatureMin)),
		TemperatureMax:  util.Float32(float32(alert.TemperatureMax)),
	}}, nil
}

func (a alert) Delete(ctx context.Context, request *gen.DeleteAlertRequest) (*gen.DeleteAlertResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing context metadata")
	}
	token := md.Get("Authorization")[0]
	userID, err := a.userService.DecodeToken(token)
	if err != nil {
		return nil, err
	}
	err = a.alertService.Delete(userID, uint(request.Id))
	if err != nil {
		return nil, err
	}
	return &gen.DeleteAlertResponse{}, nil
}

func newAlertProcedure(alertService service.Alert, userService service.User) Alert {
	return &alert{
		alertService: alertService,
		userService:  userService,
	}
}
