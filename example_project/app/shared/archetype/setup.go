package archetype

import (
	"archetype/app/shared/archetype/chi_server"
	"archetype/app/shared/archetype/container"
	"archetype/app/shared/archetype/firestore"
	"archetype/app/shared/archetype/postgres"
	"archetype/app/shared/archetype/pubsub"
	"archetype/app/shared/archetype/redis"
	"archetype/app/shared/config"

	"github.com/rs/zerolog"
)

// ARCHETYPE CONFIGURATION
func Setup(cfg config.ArchetypeConfiguration) error {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.LevelFieldName = "severity"
	zerolog.TimestampFieldName = "timestamp"

	if err := config.Setup(cfg); err != nil {
		return err
	}

	if cfg.EnablePubSub {
		if err := pubsub.Setup(); err != nil {
			return err
		}
	}

	if cfg.EnablePostgreSQLDB {
		if err := postgres.Setup(); err != nil {
			return err
		}
	}

	if cfg.EnableRedis {
		if err := redis.Setup(); err != nil {
			return err
		}
	}

	if cfg.EnableFirestore {
		if err := firestore.Setup(); err != nil {
			return err
		}
	}

	if err := dependencyInjection(); err != nil {
		return err
	}

	if cfg.EnableHTTPServer {
		if err := chi_server.Setup(); err != nil {
			return err
		}
	}
	return nil
}

// CUSTOM INITIALIZATION OF YOUR DOMAIN COMPONENTS
func dependencyInjection() error {
	for _, v := range container.Dependencies {
		if v.InjectionProps.Paralel {
			go v.Dependency()
		}
		if !v.InjectionProps.Paralel {
			if err := v.Dependency(); err != nil {
				return err
			}
		}
	}
	return nil
}
