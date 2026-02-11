{{- if .Options.Infra.Postgres }}
package config

import env "github.com/caarlos0/env/v10"

type PostgresConfig struct {
	Dsn                string `env:"POSTGRES_DSN,required"`
	MaxOpenConns       int    `env:"POSTGRES_MAX_OPEN_CONNS" envDefault:"50"`
	MaxIdleConns       int    `env:"POSTGRES_MAX_IDLE_CONNS" envDefault:"10"`
	MaxConnMaxLifetime int    `env:"POSTGRES_CONN_MAX_LIFETIME_MINUTES" envDefault:"30"`
}

func NewPostgresCfg() (*PostgresConfig, error) {
	cfg := PostgresConfig{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
{{- end }}