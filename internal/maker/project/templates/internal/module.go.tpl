package internal

import (
	"{{ .ModulePath }}/internal/infra"
	"{{ .ModulePath }}/internal/modules"

	"go.uber.org/fx"
)

var Module = fx.Module(
	"internal",
	infra.Module,
	modules.Module,
)
