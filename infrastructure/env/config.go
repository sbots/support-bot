package env

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	TestToken string `envconfig:"TEST_BOT_TOKEN" default:"test"`
	Domain    string `envconfig:"DOMAIN"`
	Host      string `envconfig:"HOST" default:"localhost"`
	Port      string `envconfig:"PORT" default:"8080"`
	SecretKey string `envconfig:"SECRET_KEY" default:"secret"`
}

func FromOS() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to read env parameters: %w", err)
	}
	return &cfg, nil
}

func (c Config) GetAddr() string {
	return c.Host + ":" + c.Port
}
