package server

import (
	"bytes"
	"fmt"
	"text/template"
)

type loaderData struct {
	Pkg string
}

func buildLoaderSource(pkg string) string {
	tpl, err := template.ParseFS(tplFS, "templates/runway_loader.go.tpl")
	if err != nil {
		panic(fmt.Errorf("parse tpl: %w", err))
	}

	data := loaderData{
		Pkg: pkg,
	}
	var buf bytes.Buffer
	if err := tpl.Execute(&buf, data); err != nil {
		panic(fmt.Errorf("exec tpl: %w", err))
	}

	return buf.String()
}
