package redis

import (
	"archetype/app/shared/config"
	"context"
	"strconv"
	"sync"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

var Client *redis.Client

var once sync.Once

func Setup() error {
	addr := config.REDIS_ADDRESS.Get()
	once.Do(func() {
		db, _ := strconv.Atoi(config.REDIS_DB.Get())
		Client = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: config.REDIS_PASSWORD.Get(),
			DB:       db,
		})
	})
	ping := Client.Ping(context.Background())
	if err := ping.Err(); err != nil {
		log.Error().Err(err).Msg("error on ping redis connection")
		return err
	}
	return nil
}
