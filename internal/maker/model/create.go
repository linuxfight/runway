package model

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Create(name string) error {
	file := filepath.Join("schema", strings.ToLower(name)+".go")

	if _, err := os.Stat(file); err == nil {
		return fmt.Errorf("model already exists: %s", file)
	}

	src, err := render("model.go.tpl", name)
	if err != nil {
		return err
	}

	return os.WriteFile(file, []byte(src), 0644)
}
