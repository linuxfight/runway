package server

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cryingcatscloud/runway/internal/generator/openapi"
	"github.com/cryingcatscloud/runway/internal/gomod"
)

func GenerateAll() error {
	apiPkgs, err := openapi.FindAPIPackages()
	if err != nil {
		return err
	}

	mod, err := gomod.Discover(".")
	if err != nil {
		return err
	}

	gen, err := New()
	if err != nil {
		return err
	}

	for _, pkgDir := range apiPkgs {
		if err := generateOne(gen, mod, pkgDir); err != nil {
			return err
		}
	}

	fmt.Println("✔ server generation completed")
	return nil
}

func generateOne(gen *Generator, mod *gomod.Info, pkgDir string) error {
	outPath := filepath.Join(pkgDir, "server.gen.go")

	importPath, err := mod.ImportPath(pkgDir)
	if err != nil {
		return err
	}

	_ = os.Remove(outPath)

	routes, err := LoadRoutes(pkgDir, importPath)
	if err != nil {
		return fmt.Errorf("load routes (%s): %w", pkgDir, err)
	}

	absPkg, err := filepath.Abs(pkgDir)
	if err != nil {
		return err
	}

	pkgName := filepath.Base(absPkg)
	moduleName := filepath.Base(filepath.Dir(absPkg))

	data := BuildTemplateData(routes, pkgName, moduleName)

	out, err := gen.GenerateBytes(data)
	if err != nil {
		return err
	}

	if err := os.WriteFile(outPath, out, 0644); err != nil {
		return err
	}

	fmt.Println("✔ server generated:", outPath)
	return nil
}
