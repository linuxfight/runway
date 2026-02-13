package gomod

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Info struct {
	Root       string
	ModulePath string
}

func Discover(start string) (*Info, error) {
	root, err := findModuleRoot(start)
	if err != nil {
		return nil, err
	}

	mod, err := readModulePath(root)
	if err != nil {
		return nil, err
	}

	return &Info{
		Root:       root,
		ModulePath: mod,
	}, nil
}

func (m *Info) ImportPath(pkgDir string) (string, error) {
	absPkg, err := filepath.Abs(pkgDir)
	if err != nil {
		return "", err
	}

	rel, err := filepath.Rel(m.Root, absPkg)
	if err != nil {
		return "", err
	}

	return m.ModulePath + "/" + filepath.ToSlash(rel), nil
}

func findModuleRoot(start string) (string, error) {
	dir, _ := filepath.Abs(start)

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("go.mod not found")
		}
		dir = parent
	}
}

func readModulePath(root string) (string, error) {
	f, err := os.Open(filepath.Join(root, "go.mod"))
	if err != nil {
		return "", err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module")), nil
		}
	}

	return "", fmt.Errorf("module path not found")
}
