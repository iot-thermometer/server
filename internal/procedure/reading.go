package procedure

import (
	"context"
	"github.com/iot-thermometer/server/gen"
	"github.com/iot-thermometer/server/internal/service"
)

type Reading interface {
	GetReadings(ctx context.Context, request *gen.ListReadingsRequest) (*gen.ListReadingsResponse, error)
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

func (r readingProcedure) GetReadings(ctx context.Context, request *gen.ListReadingsRequest) (*gen.ListReadingsResponse, error) {
	//TODO implement me
	panic("implement me")
}
