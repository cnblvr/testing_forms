package config

import (
	"time"
)

type configEnv struct {
	HttpServerPort   string        `env:"HTTP_SERVER_PORT,notEmpty" envDefault:"8080"`
	HttpReadTimeout  time.Duration `env:"HTTP_READ_TIMEOUT" envDefault:"5m"`
	HttpWriteTimeout time.Duration `env:"HTTP_WRITE_TIMEOUT" envDefault:"5m"`
}

func (c Config) HttpServerPort() string {
	return c.env.HttpServerPort
}

func (c Config) HttpReadTimeout() time.Duration {
	return c.env.HttpReadTimeout
}

func (c Config) HttpWriteTimeout() time.Duration {
	return c.env.HttpWriteTimeout
}
