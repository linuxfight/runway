package config

import env "github.com/caarlos0/env/v10"

type HTTPConfig struct {
	PublicAddr   string `env:"HTTP_PUBLIC_BIND_ADDR"  envDefault:":8080"`
	ReadTimeout  int    `env:"HTTP_READ_TIMEOUT_SEC"  envDefault:"15"`
	WriteTimeout int    `env:"HTTP_WRITE_TIMEOUT_SEC" envDefault:"15"`
	IdleTimeout  int    `env:"HTTP_IDLE_TIMEOUT_SEC"  envDefault:"60"`
}

func NewHTTPCfg() (*HTTPConfig, error) {
	cfg := HTTPConfig{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
