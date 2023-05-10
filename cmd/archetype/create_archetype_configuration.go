package archetype

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func CreateArchetypeConfiguration(moduleName string) error {
	configDir := filepath.Join(moduleName, "app/config")
	err := os.MkdirAll(configDir, 0755)
	if err != nil {
		return fmt.Errorf("error creating config directory: %v", err)
	}

	configPath := filepath.Join(configDir, "config.go")
	configContent := `package config

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
}

type Config string

// ARCHETYPE CONFIGURATION
const PORT Config = "PORT"
const COUNTRY Config = "COUNTRY"
const DD_SERVICE Config = "DD_SERVICE"
const DD_VERSION Config = "DD_VERSION"
const DD_ENV Config = "DD_ENV"

const INTEGRATION_TESTS Config = "INTEGRATION_TESTS"

const GOOGLE_PROJECT_ID Config = "GOOGLE_PROJECT_ID"
const GOOGLE_APPLICATION_CRETENTIALS_B64 Config = "GOOGLE_APPLICATION_CRETENTIALS_B64"

const DATABASE_HOST Config = "database.postgres.hostName"
const DATABASE_PORT Config = "database.postgres.port"
const DATABASE_NAME Config = "database.postgres.db.name"
const DATABASE_USERNAME Config = "database.postgres.username"
const DATABASE_PWD Config = "database.postgres.pwd"

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
		DD_SERVICE,
		DD_VERSION,
	}

	if cnf.EnablePubSub || cnf.EnableFirestore {
		requiredEnvVars = append(requiredEnvVars, GOOGLE_PROJECT_ID)
		requiredEnvVars = append(requiredEnvVars, GOOGLE_APPLICATION_CRETENTIALS_B64)
	}

	if cnf.EnablePostgreSQLDB {
		requiredEnvVars = append(requiredEnvVars, DATABASE_HOST)
		requiredEnvVars = append(requiredEnvVars, DATABASE_PORT)
		requiredEnvVars = append(requiredEnvVars, DATABASE_NAME)
		requiredEnvVars = append(requiredEnvVars, DATABASE_USERNAME)
		requiredEnvVars = append(requiredEnvVars, DATABASE_PWD)
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
}`

	err = ioutil.WriteFile(configPath, []byte(configContent), 0644)
	if err != nil {
		err := fmt.Errorf("error writing config file: %v", err)
		fmt.Println(err)
		return err
	}

	fmt.Printf("Config file generated successfully at %s.\n", configPath)
	return nil
}
