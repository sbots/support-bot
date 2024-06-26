package env

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"time"
)

type Config struct {
	Domain string `envconfig:"DOMAIN"`
	Host   string `envconfig:"HOST" default:"0.0.0.0"`
	Port   string `envconfig:"PORT" default:"8080"`

	DB string `envconfig:"DATABASE_URL"`

	SecretKey       string        `envconfig:"AUTH_SECRET_KEY" default:"secret"`
	TokenIssuer     string        `envconfig:"AUTH_TOKEN_ISSUER" default:"support-bot-platform-test"`
	TokenExpiration time.Duration `envconfig:"AUTH_TOKEN_EXPIRATION" default:"5m"`
	LogLevel        string        `envconfig:"LOG_LEVEL" default:"DEBUG"`
	LogPrettify     bool          `envconfig:"LOG_PRETTIFY" default:"true"`

	ProductionMode bool `envconfig:"PRODUCTION_MODE"`
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
