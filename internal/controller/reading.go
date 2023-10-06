package controller

import "github.com/iot-thermometer/server/internal/service"

type Reading interface {
}

type reading struct {
	readingService service.Reading
}

func newReadingController(readingService service.Reading) Reading {
	return &reading{
		readingService: readingService,
	}
}
