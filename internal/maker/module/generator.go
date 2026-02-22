package module

import (
	"bytes"
	"embed"
	"text/template"
)

//go:embed templates/**
var tplFS embed.FS

type data struct {
	Name             string
	Title            string
	ModuleImportPath string
}

func render(tplName string, d data) (string, error) {
	b, err := tplFS.ReadFile("templates/" + tplName)
	if err != nil {
		return "", err
	}

	t, err := template.New(tplName).Parse(string(b))
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, d); err != nil {
		return "", err
	}

	return buf.String(), nil
}
