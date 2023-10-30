package archetype

import (
	"my-project-name/app/shared/archetype/container"
	"my-project-name/app/shared/config"

	_ "my-project-name/app/shared/archetype/echo_server"

	_ "my-project-name/app/adapter/in/view/home"

	_ "my-project-name/app/adapter/in/view/content"
	_ "my-project-name/app/adapter/in/view/layout"

	"github.com/rs/zerolog"
)

// ARCHETYPE CONFIGURATION
func Setup() error {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.LevelFieldName = "severity"
	zerolog.TimestampFieldName = "timestamp"

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
