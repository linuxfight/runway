package modules

import (
	"{{ .ModulePath }}/internal/modules/notes"

	"go.uber.org/fx"
)

var Module = fx.Module(
	"modules",
	notes.Module,
)
