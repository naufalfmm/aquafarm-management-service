package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type EnvConfig struct {
	ServiceName string `envconfig:"SERVICE_NAME" default:"Aquafarm Management Service" required:"true"`
	ServicePort int    `envconfig:"SERVICE_PORT" default:"8090" required:"true"`

	MySqlDbAddress         string        `envconfig:"MYSQL_DB_ADDRESS" required:"true"`
	MySqlDbUsername        string        `envconfig:"MYSQL_DB_USERNAME" required:"true"`
	MySqlDbPassword        string        `envconfig:"MYSQL_DB_PASSWORD" required:"true"`
	MySqlDbName            string        `envconfig:"MYSQL_DB_NAME" required:"true"`
	MySqlMaxIdleConnection int           `envconfig:"MYSQL_MAX_IDLE_CONNECTION" default:"10"`
	MySqlMaxOpenConnection int           `envconfig:"MYSQL_MAX_OPEN_CONNECTION" default:"10"`
	MySqlConnMaxLifetime   time.Duration `envconfig:"MYSQL_CONNECTION_MAX_LIFE_TIME" default:"60s"`

	JwtPublicKey string `envconfig:"JWT_PUBLIC_KEY" required:"true"`
	JwtAlg       string `envconfig:"JWT_ALG" required:"true" default:"HS256"`
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
