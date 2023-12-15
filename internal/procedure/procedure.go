package procedure

import (
	"crypto/tls"
	"net"

	"github.com/iot-thermometer/server/gen"
	"github.com/iot-thermometer/server/internal/service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
	grpcCredentials, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		logrus.Panic(err)
	}
	grpcServer := grpc.NewServer(grpc.Creds(credentials.NewServerTLSFromCert(&grpcCredentials)))
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
