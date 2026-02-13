package config

import env "github.com/caarlos0/env/v10"

type AppConfig struct {
	Env string `env:"APP_ENV" envDefault:"production"`
}

func NewAppCfg() (*AppConfig, error) {
	cfg := AppConfig{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
