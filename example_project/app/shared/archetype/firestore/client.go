package firestore

import (
	"archetype/app/shared/config"
	"archetype/app/shared/utils"

	"github.com/rs/zerolog/log"

	"context"
	"sync"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var once sync.Once

var Client *firestore.Client

func Setup() error {
	ctx := context.Background()

	creds, err := utils.DecodeBase64(config.GOOGLE_APPLICATION_CRETENTIALS_B64.Get())
	if err != nil {
		log.Error().Err(err).Msg("error decoding GOOGLE_APPLICATION_CRETENTIALS_B64")
		return err
	}

	sa := option.WithCredentialsJSON(creds)

	app, err := firebase.NewApp(ctx, &firebase.Config{
		ProjectID: config.GOOGLE_PROJECT_ID.Get(),
	}, sa)

	if err != nil {
		log.Error().Err(err).Msg("error initializing firebase app")
		return err
	}

	c, err := app.Firestore(ctx)

	if err != nil {
		log.Error().Err(err).Msg("error getting firestore client")
		return err
	}

	Client = c
	return nil
}
