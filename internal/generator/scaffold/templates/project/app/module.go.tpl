package app

import (
	"go.uber.org/fx"
	"{{ .ModulePath }}/app/config"
	"{{ .ModulePath }}/app/http"
	"{{ .ModulePath }}/internal"
)

var Module = fx.Module(
	"app",
	config.Module,
	http.Module,
	internal.Module,
)
