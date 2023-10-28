package procedure

import (
	"github.com/iot-thermometer/server/gen"
	"github.com/iot-thermometer/server/internal/service"
	"google.golang.org/grpc"
	"net"
)

type Procedures interface {
	Serve(listener net.Listener) error
}

type procedures struct {
	userProcedure    User
	deviceProcedure  Device
	readingProcedure Reading

	grpcServer *grpc.Server
}

func NewProcedures(services service.Services) Procedures {
	userProcedure := newUserProcedure(services.User())
	deviceProcedure := newDeviceProcedure(services.Device(), services.User())
	readingProcedure := newReadingProcedure(services.Reading(), services.User())
	grpcServer := grpc.NewServer()
	s := &server{
		userProcedure:    userProcedure,
		deviceProcedure:  deviceProcedure,
		readingProcedure: readingProcedure,
	}
	gen.RegisterThermometerServiceServer(grpcServer, s)
	return &procedures{
		userProcedure:    userProcedure,
		deviceProcedure:  deviceProcedure,
		readingProcedure: readingProcedure,
		grpcServer:       grpcServer,
	}
}

func (p *procedures) Serve(listener net.Listener) error {
	return p.grpcServer.Serve(listener)
}
