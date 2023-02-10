package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type EnvConfig struct {
	MySqlDbHost     string `envconfig:"MYSQL_DB_HOST" required:"true"`
	MySqlDbPort     string `envconfig:"MYSQL_DB_PORT" required:"true"`
	MySqlDbUsername string `envconfig:"MYSQL_DB_USERNAME" required:"true"`
	MySqlDbPassword string `envconfig:"MYSQL_DB_PASSWORD" required:"true"`
	MySqlDbName     string `envconfig:"MYSQL_DB_NAME" required:"true"`

	MySqlRetry     int           `envconfig:"MYSQL_RETRY" default:"3"`
	MySqlWaitSleep time.Duration `envconfig:"MYSQL_WAIT_SLEEP" default:"1s"`
}

func NewConfig() (*EnvConfig, error) {
	var config EnvConfig

	filename := os.Getenv("CONFIG_FILE")

	if filename == "" {
		filename = ".env"
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		if err := envconfig.Process("", &config); err != nil {
			return nil, errors.Wrap(err, "failed to read from env variable")
		}
	}

	if err := godotenv.Load(filename); err != nil {
		return nil, errors.Wrap(err, "failed to read from .env file")
	}

	if err := envconfig.Process("", &config); err != nil {
		return nil, errors.Wrap(err, "failed to read from env variable")
	}

	return &config, nil

}
