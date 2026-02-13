package openapi

import (
	"fmt"
	"os"

	"github.com/cryingcatscloud/runway/internal/gomod"
	"github.com/getkin/kin-openapi/openapi3"
)

type Config struct {
	Title   string
	Version string
	OutFile string
}

func Generate(cfg Config) error {
	apiPkgs, err := FindAPIPackages()
	if err != nil {
		return err
	}

	mod, err := gomod.Discover(".")
	if err != nil {
		return err
	}

	doc := &openapi3.T{
		OpenAPI: "3.0.3",
		Info: &openapi3.Info{
			Title:   cfg.Title,
			Version: cfg.Version,
		},
		Paths: &openapi3.Paths{},
	}

	sb := NewSchemaBuilder()

	for _, pkgDir := range apiPkgs {
		importPath, err := mod.ImportPath(pkgDir)
		if err != nil {
			return err
		}

		routes, err := LoadRoutes(pkgDir, importPath)
		if err != nil {
			return fmt.Errorf("load routes (%s): %w", pkgDir, err)
		}

		for _, r := range routes {
			addRoute(doc, r, sb)
		}
	}

	doc.Components = &openapi3.Components{
		Schemas: sb.Components(),
	}

	data, err := yamlMarshal(doc)
	if err != nil {
		return err
	}

	out := cfg.OutFile
	if out == "" {
		out = "openapi.yaml"
	}

	return os.WriteFile(out, data, 0644)
}
