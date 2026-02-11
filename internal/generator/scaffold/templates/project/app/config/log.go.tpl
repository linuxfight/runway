package config

import env "github.com/caarlos0/env/v10"

type LogConfig struct {
	Level string `env:"LOG_LEVEL" envDefault:"info"`
	JSON  bool   `env:"LOG_JSON" envDefault:"true"`
}

func NewLogCfg() (*LogConfig, error) {
	cfg := LogConfig{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
