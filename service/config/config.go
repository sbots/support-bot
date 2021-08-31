package config

import (
	"fmt"
	"os"
)

const (
	defaultHost = "0.0.0.0"
	defaultPort = "7777"
)

type Config struct {
	TestToken string
	Domain    string
	Host      string
	Port      string
	SecretKey []byte
}

func FromOS() (*Config, error) {
	token := os.Getenv("TEST_BOT_TOKEN")
	if len(token) < 1 {
		return nil, fmt.Errorf("test token is missing in the env")
	}
	domain := os.Getenv("DOMAIN")
	if len(domain) < 1 {
		return nil, fmt.Errorf("domain is missing in the env")
	}
	port := os.Getenv("PORT")
	if len(port) < 1 {
		port = defaultPort
	}
	secret := os.Getenv("SECRET_KEY")
	if len(secret) < 1 {
		return nil, fmt.Errorf("secred key is required")
	}

	return &Config{
		TestToken: token,
		Domain:    domain,
		Host:      defaultHost,
		Port:      port,
		SecretKey: []byte(secret),
	}, nil
}

func (c Config) GetAddr() string {
	return c.Host + ":" + c.Port
}
