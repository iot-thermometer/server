package client

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/iot-thermometer/server/internal/dto"
	"github.com/sirupsen/logrus"
	"time"
)

type MessageCallback func(message mqtt.Message) error

type Broker interface {
	Connect(topic string, callback MessageCallback) error
}

type broker struct {
	config dto.Config
	client mqtt.Client
}

func newBroker(config dto.Config) Broker {
	return &broker{config: config}
}

func (b broker) Connect(topic string, callback MessageCallback) error {
	opts := mqtt.NewClientOptions().AddBroker(b.config.Broker).SetClientID("server")
	opts.SetKeepAlive(60 * time.Second)
	opts.SetDefaultPublishHandler(func(client mqtt.Client, message mqtt.Message) {
		logrus.Infof("Received message on topic %s: %s", message.Topic(), message.Payload())
		err := callback(message)
		if err != nil {
			logrus.Error(err)
		}
	})
	opts.SetPingTimeout(1 * time.Second)
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	b.client = client

	if token := client.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}
