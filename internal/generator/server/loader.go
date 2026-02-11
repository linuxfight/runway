package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func LoadRoutes(pkgDir, importPath string) ([]RuntimeRoute, error) {
	tmpName := ".runway_loader.go"
	tmpFile := filepath.Join(pkgDir, tmpName)

	fmt.Println(tmpFile)

	code := buildLoaderSource(importPath)

	if err := os.WriteFile(tmpFile, []byte(code), 0644); err != nil {
		return nil, err
	}

	fmt.Println("File writed")

	// defer os.Remove(tmpFile)

	fmt.Println("go", "run", tmpName)
	fmt.Println(pkgDir)
	cmd := exec.Command("go", "run", tmpName)
	cmd.Dir = pkgDir

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	var routes []RuntimeRoute
	if err := json.Unmarshal(out.Bytes(), &routes); err != nil {
		return nil, err
	}

	return routes, nil
}
