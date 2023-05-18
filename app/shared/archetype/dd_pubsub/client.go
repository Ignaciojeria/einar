package pubsub

import (
	"archetype/app/shared/archetype/container"
	"archetype/app/shared/config"
	"archetype/app/shared/utils"
	"context"

	"cloud.google.com/go/pubsub"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"google.golang.org/api/option"
)

var Client *pubsub.Client

func init() {
	config.Installations.EnablePubSub = true
	container.InjectInstallation(func() error {
		projectId := config.GOOGLE_PROJECT_ID.Get()
		creds, err := utils.DecodeBase64(config.GOOGLE_APPLICATION_CRETENTIALS_B64.Get())
		if err != nil {
			log.Error().Err(err).Msg("error decoding GOOGLE_APPLICATION_CRETENTIALS_B64")
			return err
		}
		c, err := pubsub.NewClient(context.Background(), projectId, option.WithCredentialsJSON(creds))
		if err != nil {
			log.Error().Err(err).Msg("error getting pubsub client")
			return err
		}
		Client = c
		return nil
	}, container.InjectionProps{
		Paralel:      false,
		DependencyID: uuid.NewString(),
	})
}
