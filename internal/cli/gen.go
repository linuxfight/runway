package cli

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cryingcatscloud/runway/internal/generator/server"
	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "gen [package-dir]",
	Short: "Generate server code from Routes provider",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		pkgDir := args[0]

		root, err := findModuleRoot(pkgDir)
		if err != nil {
			return err
		}

		modulePath, err := readModulePath(root)
		if err != nil {
			return err
		}

		importPath, err := buildImportPath(root, modulePath, pkgDir)
		if err != nil {
			return err
		}

		routes, err := server.LoadRoutes(pkgDir, importPath)
		if err != nil {
			return err
		}

		data := server.BuildTemplateData(routes, filepath.Base(pkgDir))

		gen, err := server.New()
		if err != nil {
			return err
		}

		out, err := gen.GenerateBytes(data)
		if err != nil {
			return err
		}

		outPath := filepath.Join(pkgDir, "server.gen.go")
		if err := os.WriteFile(outPath, out, 0644); err != nil {
			return err
		}

		fmt.Println("✔ generated:", outPath)
		return nil
	},
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

func buildImportPath(root, modulePath, pkgDir string) (string, error) {
	absPkg, err := filepath.Abs(pkgDir)
	if err != nil {
		return "", err
	}

	rel, err := filepath.Rel(root, absPkg)
	if err != nil {
		return "", err
	}

	return modulePath + "/" + filepath.ToSlash(rel), nil
}
