package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DB struct {
		Host      string `envconfig:"DB_HOST"`
		Port      int    `envconfig:"DB_PORT"`
		User      string `envconfig:"DB_USER"`
		Pass      string `envconfig:"DB_PASS"`
		Name      string `envconfig:"DB_NAME"`
	}
	Server struct {
		RESTPort string `envconfig:"REST_PORT"`
	}
}

// Get to get defined configuration
func Get() Config {
	cfg := Config{}
	envconfig.MustProcess("", &cfg)
	return cfg
}
