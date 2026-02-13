package module

import (
	"bytes"
	"embed"
	"strings"
	"text/template"
)

//go:embed templates/**
var tplFS embed.FS

type data struct {
	Name  string
	Title string
}

func render(tplName, name string) (string, error) {
	b, err := tplFS.ReadFile("templates/" + tplName)
	if err != nil {
		return "", err
	}

	t, err := template.New(tplName).Parse(string(b))
	if err != nil {
		return "", err
	}

	d := data{
		Name:  name,
		Title: strings.ToUpper(name[:1]) + name[1:],
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, d); err != nil {
		return "", err
	}

	return buf.String(), nil
}
