package service

type Reading interface {
}

type reading struct {
}

func newReadingService() Reading {
	return &reading{}
}
