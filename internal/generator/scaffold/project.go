package scaffold

import (
	"embed"
)

//go:embed templates/project/**
var projectTemplates embed.FS
