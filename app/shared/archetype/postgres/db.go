package postgres

import (
	"archetype/app/shared/archetype/container"
	"archetype/app/shared/config"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	config.Installations.EnablePostgreSQLDB = true
	container.InjectInstallation(func() error {
		username := config.DATABASE_POSTGRES_USERNAME.Get()
		pwd := config.DATABASE_POSTGRES_PASSWORD.Get()
		host := config.DATABASE_POSTGRES_HOSTNAME.Get()
		dbname := config.DATABASE_POSTGRES_NAME.Get()
		db, err := gorm.Open(postgres.Open("postgres://" + username + ":" + pwd + "@" + host + "/" + dbname + "?sslmode=disable"))
		if err != nil {
			log.Error().Err(err).Msg("error getting postgresql connection")
			return err
		}
		DB = db
		return nil
	}, container.InjectionProps{
		Paralel:      false,
		DependencyID: uuid.NewString(),
	})
}
