package openapi

import (
	"fmt"
	"os"
	"path/filepath"
)

func FindAPIPackages() ([]string, error) {
	if _, err := os.Stat("go.mod"); err != nil {
		return nil, fmt.Errorf(
			"run from project root: expected go.mod and internal/modules/.../api/routes.go",
		)
	}

	modulesDir := filepath.Join("internal", "modules")

	if _, err := os.Stat(modulesDir); err != nil {
		return nil, fmt.Errorf(
			"run from project root: expected internal/modules/.../api/routes.go",
		)
	}

	entries, err := os.ReadDir(modulesDir)
	if err != nil {
		return nil, err
	}

	var result []string

	for _, e := range entries {
		if !e.IsDir() {
			continue
		}

		apiDir := filepath.Join(modulesDir, e.Name(), "api")
		routesFile := filepath.Join(apiDir, "routes.go")

		if _, err := os.Stat(routesFile); err == nil {
			result = append(result, apiDir)
		}
	}

	if len(result) == 0 {
		return nil, fmt.Errorf(
			"no routes found. run from project root: expected internal/modules/.../api/routes.go",
		)
	}

	return result, nil
}
