package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type ArchetypeConfiguration struct {
	//HTTP CLIENT IS ENABLED BY DEFAULT
	EnvironmentPath    string
	EnablePostgreSQLDB bool
	EnablePubSub       bool
	EnableFirestore    bool
	EnableHTTPServer   bool
	EnableRedis        bool
}

type Config string

// ARCHETYPE CONFIGURATION
const PORT Config = "PORT"
const COUNTRY Config = "COUNTRY"
const SERVICE Config = "SERVICE"
const ENV Config = "ENV"

const INTEGRATION_TESTS Config = "INTEGRATION_TESTS"

const GOOGLE_PROJECT_ID Config = "GOOGLE_PROJECT_ID"
const GOOGLE_APPLICATION_CRETENTIALS_B64 Config = "GOOGLE_APPLICATION_CRETENTIALS_B64"

const DATABASE_POSTGRES_HOSTNAME Config = "DATABASE_POSTGRES_HOSTNAME"
const DATABASE_POSTGRES_PORT Config = "DATABASE_POSTGRES_PORT"
const DATABASE_POSTGRES_NAME Config = "DATABASE_POSTGRES_NAME"
const DATABASE_POSTGRES_USERNAME Config = "DATABASE_POSTGRES_USERNAME"
const DATABASE_POSTGRES_PASSWORD Config = "DATABASE_POSTGRES_PASSWORD"
const OPENAI_API_KEY Config = "OPENAI_API_KEY"

// Redis configuration
const REDIS_ADDRESS Config = "REDIS_ADDRESS"
const REDIS_PASSWORD Config = "REDIS_PASSWORD"
const REDIS_DB Config = "REDIS_DB"

func (e Config) Get() string {
	return os.Getenv(string(e))
}

func Setup(cnf ArchetypeConfiguration) error {

	errs := []string{}

	if err := godotenv.Load(); err != nil {
		log.Error().Err(err).Msg(".env file not found getting environments from envgonsul")
	}

	// Check that all required environment variables are set
	requiredEnvVars := []Config{
		//ARCHETYPE CONFIGURATION
		PORT,
		COUNTRY,
		SERVICE,
	}

	if cnf.EnablePubSub || cnf.EnableFirestore {
		requiredEnvVars = append(requiredEnvVars, GOOGLE_PROJECT_ID)
		requiredEnvVars = append(requiredEnvVars, GOOGLE_APPLICATION_CRETENTIALS_B64)
	}

	if cnf.EnablePostgreSQLDB {
		requiredEnvVars = append(requiredEnvVars, DATABASE_POSTGRES_HOSTNAME)
		requiredEnvVars = append(requiredEnvVars, DATABASE_POSTGRES_PORT)
		requiredEnvVars = append(requiredEnvVars, DATABASE_POSTGRES_NAME)
		requiredEnvVars = append(requiredEnvVars, DATABASE_POSTGRES_USERNAME)
		requiredEnvVars = append(requiredEnvVars, DATABASE_POSTGRES_PASSWORD)
	}

	for _, envVar := range requiredEnvVars {
		value := envVar.Get()
		if value == "" {
			errs = append(errs, string(envVar))
		}
	}

	if len(errs) > 0 {
		log.Error().Strs("notFoundEnvironments", errs).Msg("error loading environment variables")
		return fmt.Errorf("error loading environment variables: %v", errs)
	}

	return nil
}
