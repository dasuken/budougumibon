package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

// add port and TODO_ENV for port
type config struct {
	Env  string `env:"TODO_ENV" envDefault:"dev"`
	Port int    `env:"PORT" envDefault:"18080"`
}

func New() (*config, error) {
	cfg := &config{}
	if err := env.Parse(cfg); err != nil {
		fmt.Printf("%+v\n", err)
		return nil, err
	}
	return cfg, nil
}
