package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Create(name string) error {
	name = strings.ToLower(name)

	out := filepath.Join("app", "config", name+".go")

	if _, err := os.Stat(out); err == nil {
		return fmt.Errorf("config already exists: %s", out)
	}

	src, err := render(name)
	if err != nil {
		return err
	}

	return os.WriteFile(out, []byte(src), 0644)
}
