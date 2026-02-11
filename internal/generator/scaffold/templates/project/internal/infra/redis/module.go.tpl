{{- if .Options.Infra.Redis }}
package redis

import "go.uber.org/fx"

var Module = fx.Module(
	"redis",
	fx.Provide(
		NewRedis,
	),
	fx.Invoke(
		RegisterLifecycle,
	),
)
{{- end }}