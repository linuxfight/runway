package config

import "go.uber.org/fx"

var Module = fx.Module(
	"config",
	fx.Provide(
		NewAppCfg,
		{{- if .Options.Infra.Postgres }}
		NewPostgresCfg,
		{{- end }}
		{{- if .Options.Infra.Redis }}
		NewRedisCfg,
		{{- end }}
		NewHTTPCfg,
		NewLogCfg,
	),
)
