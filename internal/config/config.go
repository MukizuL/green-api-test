package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"go.uber.org/fx"
)

type Config struct {
	Addr string `env:"RUN_ADDRESS"`
}

func newConfig() (*Config, error) {
	cfg := &Config{}

	err := env.Parse(cfg)
	if err != nil {
		return nil, fmt.Errorf("error parsing env variables")
	}

	if cfg.Addr == "" {
		return nil, fmt.Errorf("missing required environment variable RUN_ADDRESS")
	}

	return cfg, nil
}

func Provide() fx.Option {
	return fx.Provide(newConfig)
}
