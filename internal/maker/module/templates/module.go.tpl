package {{ .Name }}

import "go.uber.org/fx"

var Module = fx.Module(
	"{{ .Name }}",
	fx.Provide(
		NewRepository,
		NewService,
		NewController,
	),
	fx.Invoke(
		RegisterHTTP,
	)
)
