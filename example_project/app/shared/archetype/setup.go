package archetype

import (
	_ "archetype/app/shared/archetype/chi_server"
	"archetype/app/shared/archetype/container"

	_ "archetype/app/shared/archetype/pubsub"
	"archetype/app/shared/config"

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

	if err := injectComponents(); err != nil {
		return err
	}

	if err := container.HTTPServerContainer.LoadDependency(); err != nil {
		return err
	}

	return nil
}

func InjectInstallations() error {
	for _, v := range container.InstallationsContainer {
		if v.InjectionProps.Paralel {
			go v.LoadDependency()
		}
		if !v.InjectionProps.Paralel {
			if err := v.LoadDependency(); err != nil {
				return err
			}
		}
	}
	return nil
}

// CUSTOM INITIALIZATION OF YOUR DOMAIN COMPONENTS
func injectComponents() error {
	for _, v := range container.ComponentsContainer {
		if v.InjectionProps.Paralel {
			go v.LoadDependency()
		}
		if !v.InjectionProps.Paralel {
			if err := v.LoadDependency(); err != nil {
				return err
			}
		}
	}
	return nil
}
