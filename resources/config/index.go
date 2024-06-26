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

	MySqlDbHost            string        `envconfig:"MYSQL_DB_HOST" required:"true"`
	MySqlDbPort            string        `envconfig:"MYSQL_DB_PORT" required:"true"`
	MySqlDbUsername        string        `envconfig:"MYSQL_DB_USERNAME" required:"true"`
	MySqlDbPassword        string        `envconfig:"MYSQL_DB_PASSWORD" required:"true"`
	MySqlDbName            string        `envconfig:"MYSQL_DB_NAME" required:"true"`
	MySqlMaxIdleConnection int           `envconfig:"MYSQL_MAX_IDLE_CONNECTION" default:"10"`
	MySqlMaxOpenConnection int           `envconfig:"MYSQL_MAX_OPEN_CONNECTION" default:"10"`
	MySqlConnMaxLifetime   time.Duration `envconfig:"MYSQL_CONNECTION_MAX_LIFE_TIME" default:"60s"`

	MySqlLogMode          bool          `envconfig:"MYSQL_LOG_MODE" default:"false"`
	MySqlLogSlowThreshold time.Duration `envconfig:"MYSQL_LOG_SLOW_THRESHOLD"`

	MySqlRetry     int           `envconfig:"MYSQL_RETRY" default:"3"`
	MySqlWaitSleep time.Duration `envconfig:"MYSQL_WAIT_SLEEP" default:"1s"`

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
