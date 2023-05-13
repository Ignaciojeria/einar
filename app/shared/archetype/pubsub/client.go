package pubsub

import (
	"archetype/app/shared/config"
	"archetype/app/shared/utils"
	"context"
	"sync"

	"cloud.google.com/go/pubsub"
	"github.com/rs/zerolog/log"
	"google.golang.org/api/option"
)

var Client *pubsub.Client

var once sync.Once

func Setup() error {

	projectId := config.GOOGLE_PROJECT_ID.Get()
	creds, err := utils.DecodeBase64(config.GOOGLE_APPLICATION_CRETENTIALS_B64.Get())
	if err != nil {
		log.Error().Err(err).Msg("error decoding GOOGLE_APPLICATION_CRETENTIALS_B64")
		return err
	}

	var ex error

	once.Do(func() {
		c, err := pubsub.NewClient(context.Background(), projectId, option.WithCredentialsJSON(creds))
		if err != nil {
			log.Error().Err(err).Msg("error getting pubsub client")
			ex = err
		}
		Client = c
	})

	return ex
}
