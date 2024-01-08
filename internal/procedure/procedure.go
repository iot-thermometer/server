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
	userProcedure      User
	ownershipProcedure Ownership
	deviceProcedure    Device
	readingProcedure   Reading
	alertProcedure     Alert
	phoneProcedure     Phone

	grpcServer *grpc.Server
}

func NewProcedures(services service.Services) Procedures {
	userProcedure := newUserProcedure(services.User())
	ownershipProcedure := newOwnershipProcedure(services.Ownership(), services.User())
	deviceProcedure := newDeviceProcedure(services.Device(), services.User())
	readingProcedure := newReadingProcedure(services.Reading(), services.User())
	alertProcedure := newAlertProcedure(services.Alert(), services.User())
	phoneProcedure := newPhoneProcedure(services.Phone(), services.User())
	grpcCredentials, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		logrus.Panic(err)
	}
	grpcServer := grpc.NewServer(grpc.Creds(credentials.NewServerTLSFromCert(&grpcCredentials)))
	s := &server{
		userProcedure:      userProcedure,
		ownershipProcedure: ownershipProcedure,
		deviceProcedure:    deviceProcedure,
		readingProcedure:   readingProcedure,
		alertProcedure:     alertProcedure,
		phoneProcedure:     phoneProcedure,
	}
	gen.RegisterThermometerServiceServer(grpcServer, s)
	return &procedures{
		userProcedure:      userProcedure,
		ownershipProcedure: ownershipProcedure,
		deviceProcedure:    deviceProcedure,
		readingProcedure:   readingProcedure,
		alertProcedure:     alertProcedure,
		phoneProcedure:     phoneProcedure,
		grpcServer:         grpcServer,
	}
}

func (p *procedures) Serve(listener net.Listener) error {
	return p.grpcServer.Serve(listener)
}
