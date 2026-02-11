package internal

import (
	"{{ .ModulePath }}/internal/infra"

	"go.uber.org/fx"
)

var Module = fx.Module(
	"internal",
	infra.Module,
)
