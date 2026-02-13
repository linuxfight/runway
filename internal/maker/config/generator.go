package config

import (
	"bytes"
	"embed"
	"strings"
	"text/template"
)

//go:embed templates/**
var tplFS embed.FS

type templateData struct {
	Name string
	Env  string
}

func render(name string) (string, error) {
	b, err := tplFS.ReadFile("templates/config.go.tpl")
	if err != nil {
		return "", err
	}

	t, err := template.New("config").Parse(string(b))
	if err != nil {
		return "", err
	}

	data := templateData{
		Name: Title(name),
		Env:  strings.ToUpper(name),
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
