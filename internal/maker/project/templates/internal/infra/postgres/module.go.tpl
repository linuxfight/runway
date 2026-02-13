{{- if .Options.Infra.Postgres -}}
package postgres

import "go.uber.org/fx"

var Module = fx.Module(
	"postgres",
	fx.Provide(
		NewSQLDB,
		{{- if .Options.ORM.Ent -}}
		NewEntClient,
		{{- end -}}
	),

	{{- if .Options.ORM.Ent -}}
	fx.Invoke(
		RegisterEntLifecycle,
	),
	{{- end -}}
)
{{- end -}}
