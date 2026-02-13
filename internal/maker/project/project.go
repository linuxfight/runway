package project

import (
	"embed"
)

//go:embed templates/**
var projectTemplates embed.FS
