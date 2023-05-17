package redis

import (
	"archetype/app/shared/archetype/container"
	"archetype/app/shared/config"
	"context"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

var Client *redis.Client

func init() {
	config.Installations.EnableRedis = true
	container.InjectInstallation(func() error {
		addr := config.REDIS_ADDRESS.Get()
		db, err := strconv.Atoi(config.REDIS_DB.Get())
		if err != nil {
			fmt.Println(err)
			return err
		}
		Client = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: config.REDIS_PASSWORD.Get(),
			DB:       db,
		})

		ping := Client.Ping(context.Background())
		if err := ping.Err(); err != nil {
			log.Error().Err(err).Msg("error on ping redis connection")
			return err
		}
		return nil
	}, container.InjectionProps{
		Paralel:      false,
		DependencyID: uuid.NewString(),
	})
}
