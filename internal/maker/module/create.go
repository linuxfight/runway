package module

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cryingcatscloud/runway/internal/gomod"
)

func Create(name string) error {
	mod, err := gomod.Discover(".")
	if err != nil {
		return err
	}

	name = strings.ToLower(name)

	root := filepath.Join("internal", "modules", name)
	api := filepath.Join(root, "api")

	if _, err := os.Stat(root); err == nil {
		return fmt.Errorf("module already exists: %s", root)
	}

	if err := os.MkdirAll(api, 0755); err != nil {
		return err
	}

	files := map[string]string{
		filepath.Join(root, "module.go"):     "module.go.tpl",
		filepath.Join(root, "controller.go"): "controller.go.tpl",
		filepath.Join(root, "service.go"):    "service.go.tpl",
		filepath.Join(root, "repository.go"): "repository.go.tpl",
		filepath.Join(root, "register.go"):   "register.go.tpl",
		filepath.Join(api, "routes.go"):      "routes.go.tpl",
		filepath.Join(api, "model.go"):       "model.go.tpl",
	}

	d := data{
		Name:             name,
		Title:            strings.ToUpper(name[:1]) + name[1:],
		ModuleImportPath: mod.ModulePath,
	}

	for out, tpl := range files {
		src, err := render(tpl, d)
		if err != nil {
			return err
		}
		if err := os.WriteFile(out, []byte(src), 0644); err != nil {
			return err
		}
	}

	return nil
}
