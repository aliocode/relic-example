package app

import (
	"errors"

	"github.com/caarlos0/env"
)

var (
	configValidationError = errors.New("failed to validate config")
)

type Config struct {
	RelicAppName string `env:"RELIC_APP_NAME" envDefault:""`
	RelicLicense string `env:"RELIC_LICENSE" envDefault:""`
	LogLevel     string `env:"LOG_LEVEL" envDefault:"debug"`
	HttpPort     string `env:"HTTP_PORT" envDefault:"8080"`
	UserPsqlDSN  string `env:"USER_PSQL_DSN" envDefault:""`
}

func NewConfig() (Config, error) {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, cfg.validate()
}

func (c Config) validate() error {
	if len(c.RelicAppName) == 0 {
		return configValidationError
	}

	if len(c.RelicLicense) == 0 {
		return configValidationError
	}

	return nil
}
