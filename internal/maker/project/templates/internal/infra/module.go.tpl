package infra

import (
	"go.uber.org/fx"
	{{ if .Options.Infra.Redis -}}
	"{{ .ModulePath }}/internal/infra/redis"
	{{ end }}
	{{ if .Options.Infra.Postgres -}}
	"{{ .ModulePath }}/internal/infra/postgres"
	{{ end }}
	"{{ .ModulePath }}/internal/infra/logger"
	
)

var Module = fx.Module(
	"infra",
	logger.Module,
	{{- if .Options.Infra.Postgres }}
	postgres.Module,
	{{- end }}
	{{- if .Options.Infra.Redis }}
	redis.Module,
	{{- end }}
)
