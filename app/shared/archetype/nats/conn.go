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
		nc, err := nats.Connect("nats://127.0.0.1:4444")
		Conn = nc
		return err
	})
}
