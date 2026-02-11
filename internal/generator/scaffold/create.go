package scaffold

import (
	"bytes"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func CreateProject(root string, ctx ProjectContext) error {
	return fs.WalkDir(projectTemplates, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if !strings.HasSuffix(path, ".tpl") {
			return nil
		}

		target := strings.TrimSuffix(path, ".tpl")
		target = strings.ReplaceAll(target, "{{ .ProjectName }}", ctx.ProjectName)

		outPath := filepath.Join(root, strings.TrimPrefix(target, "templates/project/"))

		tpl, err := template.ParseFS(projectTemplates, path)
		if err != nil {
			return err
		}

		var buf bytes.Buffer
		if err := tpl.Execute(&buf, ctx); err != nil {
			return err
		}

		if strings.TrimSpace(buf.String()) == "" {
			return nil
		}

		if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
			return err
		}

		f, err := os.Create(outPath)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = f.Write(buf.Bytes())
		return err
	})
}
