package archetype

import (
	_ "github.com/Ignaciojeria/einar/app/adapter/in/controller"
	_ "github.com/Ignaciojeria/einar/app/business"
	"github.com/Ignaciojeria/einar/app/shared/archetype/container"
	_ "github.com/Ignaciojeria/einar/app/shared/archetype/echo_server"
	"github.com/Ignaciojeria/einar/app/shared/config"
	_ "github.com/Ignaciojeria/einar/app/shared/archetype/resty"
	_ "github.com/Ignaciojeria/einar/app/adapter/out/client"
)

// ARCHETYPE CONFIGURATION
func Setup() error {
	if err := config.Setup(); err != nil {
		return err
	}

	if err := InjectInstallations(); err != nil {
		return err
	}

	if err := injectOutboundAdapters(); err != nil {
		return err
	}

	if err := injectInboundAdapters(); err != nil {
		return err
	}

	if !config.Installations.EnableHTTPServer {
		return nil
	}
	if err := container.HTTPServerContainer.LoadDependency(); err != nil {
		return err
	}
	return nil
}

func InjectInstallations() error {
	for _, v := range container.InstallationsContainer {
		if err := v.LoadDependency(); err != nil {
			return err
		}
	}
	return nil
}

// CUSTOM INITIALIZATION OF YOUR DOMAIN COMPONENTS
func injectOutboundAdapters() error {
	for _, v := range container.OutboundAdapterContainer {
		if err := v.LoadDependency(); err != nil {
			return err
		}
	}
	return nil
}

func injectInboundAdapters() error {
	for _, v := range container.InboundAdapterContainer {
		if err := v.LoadDependency(); err != nil {
			return err
		}
	}
	return nil
}
