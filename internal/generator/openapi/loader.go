package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func LoadRoutes(pkgDir, importPath string) ([]RouteMeta, error) {
	tmpName := "runway_loader.go"
	tmpFile := filepath.Join(pkgDir, tmpName)

	code, err := buildLoaderSource(importPath)
	if err != nil {
		return nil, err
	}

	if err := os.WriteFile(tmpFile, []byte(code), 0644); err != nil {
		return nil, err
	}
	defer os.Remove(tmpFile)

	cmd := exec.Command("go", "run", tmpName)
	cmd.Dir = pkgDir

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("openapi loader failed: %w", err)
	}

	var routes []RouteMeta
	if err := json.Unmarshal(out.Bytes(), &routes); err != nil {
		return nil, err
	}

	return routes, nil
}
