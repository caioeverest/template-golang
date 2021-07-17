package config

import (
	"log"

	"github.com/caarlos0/env/v6"
)

//Config contains all configuration variables of the system
type Config struct {
	Env      string `env:"ENVIRONMENT" envDefault:"UNKNOWN"`
	API      string `env:"APP_NAME" envDefault:"{{.RepositoryName}}"`
	Version  string `env:"APP_VERSION" envDefault:"UNKNOWN"`
	HTTPPort string `env:"HTTP_PORT" envDefault:"8080"`
}

//This are constant help create conditions to specific environments
const (
	Development = "DEVELOPMENT"
	Production  = "PRODUCTION"
	Test        = "TEST"
	Unknown     = "UNKNOWN"
)

//Get return the config loaded on the start of the application
func Get() Config {
	return Load()
}

//Load will load all the mapped configurations on the environment
func Load() Config {
	_global := Config{}
	if err := env.Parse(&_global); err != nil {
		log.Panicf("error parsing configs - %+v", err)
	}
	return _global
}
