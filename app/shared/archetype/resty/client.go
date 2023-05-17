package resty

import (
	"archetype/app/shared/archetype/container"
	"archetype/app/shared/config"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

var Client *resty.Client

func init() {
	config.Installations.EnableRestyClient = true
	container.InjectInstallation(func() error {
		Client = resty.New()
		return nil
	}, container.InjectionProps{
		Paralel:      false,
		DependencyID: uuid.NewString(),
	})
}
