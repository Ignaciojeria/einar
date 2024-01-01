package nats

import (
	"github.com/Ignaciojeria/einar/app/shared/archetype/container"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

var Conn *nats.Conn

var EinarTopic string = uuid.NewString()

func init() {
	container.InjectInstallation(func() error {
		//Customize your nats connection here :
		nc, err := nats.Connect("nats://37.16.23.225:4444")
		Conn = nc
		return err
	})
}
