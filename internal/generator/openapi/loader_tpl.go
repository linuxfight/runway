package openapi

import (
	"embed"
	"strings"
	"text/template"
)

//go:embed templates/runway_loader.go.tpl
var tplFS embed.FS

func buildLoaderSource(importPath string) (string, error) {
	tplBytes, err := tplFS.ReadFile("templates/runway_loader.go.tpl")
	if err != nil {
		return "", err
	}

	tpl, err := template.New("loader").Parse(string(tplBytes))
	if err != nil {
		return "", err
	}

	var buf strings.Builder

	err = tpl.Execute(&buf, struct {
		ImportPath string
	}{
		ImportPath: importPath,
	})
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
