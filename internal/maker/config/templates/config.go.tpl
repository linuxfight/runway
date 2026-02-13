package config

import env "github.com/caarlos0/env/v10"

type {{ .Name }}Config struct {
	// Example:
	// Addr string `env:"{{ .Env }}_ADDR" envDefault:":8080"`
}

func New{{ .Name }}Cfg() (*{{ .Name }}Config, error) {
	cfg := {{ .Name }}Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
