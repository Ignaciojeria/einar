package archetype

import (
	"archetype/app/shared/archetype/container"
	"archetype/app/shared/config"

	"github.com/rs/zerolog"
	_ "archetype/app/shared/archetype/chi_server"
	_ "archetype/app/shared/archetype/pubsub"
	_ "archetype/app/shared/archetype/firestore"
	_ "archetype/app/shared/archetype/redis"
	_ "archetype/app/shared/archetype/postgres"
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
