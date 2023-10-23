package client

import (
	"github.com/iot-thermometer/server/internal/dto"
)

type Clients interface {
	Broker() Broker
}

type clients struct {
	broker Broker
}

func NewClients(config dto.Config) Clients {
	broker := newBroker(config)
	return &clients{
		broker: broker,
	}
}

func (c clients) Broker() Broker {
	return c.broker
}
