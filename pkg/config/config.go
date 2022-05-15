package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/pkg/errors"
)

type Config struct {
	env *configEnv
}

func New() (*Config, error) {
	c := &Config{
		env: new(configEnv),
	}

	if err := env.Parse(c.env); err != nil {
		return nil, errors.WithStack(err)
	}

	return c, nil
}
