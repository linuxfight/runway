{{- if .Options.Infra.Redis }}
package config

import env "github.com/caarlos0/env/v10"

type RedisConfig struct {
	Addr     string `env:"REDIS_ADDR" envDefault:"localhost:6379"`
	Password string `env:"REDIS_PASSWORD" envDefault:""`
	DB       int    `env:"REDIS_DB" envDefault:"0"`
}

func NewRedisCfg() (*RedisConfig, error) {
	cfg := RedisConfig{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
{{- end }}