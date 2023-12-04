package procedure

import (
	"context"
	"fmt"
	"time"

	"github.com/iot-thermometer/server/gen"
	"github.com/iot-thermometer/server/internal/service"
	"google.golang.org/grpc/metadata"
)

type Reading interface {
	List(ctx context.Context, request *gen.ListReadingsRequest) (*gen.ListReadingsResponse, error)
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

func (r readingProcedure) List(ctx context.Context, request *gen.ListReadingsRequest) (*gen.ListReadingsResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing context metadata")
	}
	token := md.Get("Authorization")[0]
	userID, err := r.userService.DecodeToken(token)
	if err != nil {
		return nil, err
	}
	readings, err := r.readingService.List(userID, uint(request.Id), time.Unix(request.Start, 0), time.Unix(request.End, 0))
	if err != nil {
		return nil, err
	}
	var readingsProtobuf []*gen.Reading
	for _, reading := range readings {
		readingsProtobuf = append(readingsProtobuf, reading.Protobuf())
	}
	return &gen.ListReadingsResponse{
		Readings: readingsProtobuf,
	}, nil
}
