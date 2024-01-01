package resty

import (
	"github.com/Ignaciojeria/einar/app/shared/archetype/container"
	"github.com/Ignaciojeria/einar/app/shared/config"
	"github.com/go-resty/resty/v2"
)

var Client *resty.Client

func init() {
	config.Installations.EnableRestyClient = true
	LoadDependency()
}

func LoadDependency() container.LoadDependency {
	var dependency container.LoadDependency = func() error {
		//Customize your resty client here :
		Client = resty.New()
		return nil
	}
	container.InjectInstallation(dependency)
	return dependency
}
