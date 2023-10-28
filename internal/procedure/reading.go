package procedure

import (
	"fmt"

	"github.com/iot-thermometer/server/gen"
	"github.com/iot-thermometer/server/internal/service"
	"google.golang.org/grpc/metadata"
)

type Reading interface {
	List(request *gen.ListReadingsRequest, stream gen.ThermometerService_ListReadingsServer) error
}

type readingProcedure struct {
	readingService service.Reading
	userService    service.User
}

func newReadingProcedure(readingService service.Reading, userService service.User) Reading {
	return &readingProcedure{
		readingService: readingService,
		userService:    userService,
	}
}

func (r readingProcedure) List(request *gen.ListReadingsRequest, stream gen.ThermometerService_ListReadingsServer) error {
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return fmt.Errorf("missing context metadata")
	}
	token := md.Get("Authorization")[0]
	userID, err := r.userService.DecodeToken(token)
	if err != nil {
		return err
	}
	readings, err := r.readingService.List(userID, uint(request.Id))
	if err != nil {
		return err
	}

	var readingsProtobuf []*gen.Reading
	for _, reading := range readings {
		readingsProtobuf = append(readingsProtobuf, reading.Protobuf())
	}
	err = stream.Send(&gen.ListReadingsResponse{
		Readings: readingsProtobuf,
	})
	if err != nil {
		return err
	}

	for {
		select {
		case <-stream.Context().Done():
			return nil
		case reading := <-r.readingService.Channel():
			stream.Send(&gen.ListReadingsResponse{
				Readings: []*gen.Reading{reading.Protobuf()},
			})
		}
	}
}
