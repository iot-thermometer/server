package procedure

import "github.com/iot-thermometer/server/internal/service"

type Reading interface {
}

type readingProcedure struct {
	readingService service.Reading
}

func newReadingProcedure(readingService service.Reading) Reading {
	return &readingProcedure{
		readingService: readingService,
	}
}
