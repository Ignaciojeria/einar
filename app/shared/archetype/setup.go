package archetype

import (
	_ "github.com/Ignaciojeria/einar/app/adapter/in/cli"
	_ "github.com/Ignaciojeria/einar/app/shared/archetype/cmd"
	"github.com/Ignaciojeria/einar/app/shared/archetype/container"
	"github.com/Ignaciojeria/einar/app/shared/config"
	_ "github.com/Ignaciojeria/einar/app/business"
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
