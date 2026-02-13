package server

import (
	"bytes"
	"embed"
	"io"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

//go:embed templates/**
var tplFS embed.FS

type Generator struct {
	tpl *template.Template
}

func New() (*Generator, error) {
	b, err := tplFS.ReadFile("templates/server.gen.go.tpl")
	if err != nil {
		return nil, err
	}

	t, err := template.New("server").Funcs(sprig.FuncMap()).Parse(string(b))
	if err != nil {
		return nil, err
	}

	return &Generator{tpl: t}, nil
}

func (g *Generator) Generate(w io.Writer, data TemplateData) error {
	return g.tpl.Execute(w, data)
}

func (g *Generator) GenerateBytes(data TemplateData) ([]byte, error) {
	var buf bytes.Buffer
	err := g.Generate(&buf, data)
	return buf.Bytes(), err
}
