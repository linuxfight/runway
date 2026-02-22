package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var allowedMethods = map[string]struct{}{
	"GET":     {},
	"POST":    {},
	"PUT":     {},
	"PATCH":   {},
	"DELETE":  {},
	"HEAD":    {},
	"OPTIONS": {},
}

func LoadRoutes(pkgDir, importPath string) ([]RuntimeRoute, error) {
	tmpName := "runway_loader.go"
	tmpFile := filepath.Join(pkgDir, tmpName)

	code := buildLoaderSource(importPath)

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
		return nil, err
	}

	var routes []RuntimeRoute
	if err := json.Unmarshal(out.Bytes(), &routes); err != nil {
		return nil, err
	}

	seen := map[string]string{}

	for i := range routes {
		r := &routes[i]

		if err := validateMethod(r.Method); err != nil {
			return nil, fmt.Errorf("route %s: %w", r.Id, err)
		}

		if err := validatePath(r.Path); err != nil {
			return nil, fmt.Errorf("route %s: %w", r.Id, err)
		}

		key := strings.ToUpper(r.Method) + " " + normalizePath(r.Path)

		if prev, exists := seen[key]; exists {
			return nil, fmt.Errorf(
				"duplicate route: %s conflicts with %s (%s %s)",
				r.Id, prev, r.Method, r.Path,
			)
		}

		if r.Id == "" {
			return nil, fmt.Errorf("found route with empty ID")
		}

		ids := map[string]struct{}{}

		for i := range routes {
			r := &routes[i]

			if _, ok := ids[r.Id]; ok {
				return nil, fmt.Errorf("duplicate route id: %s", r.Id)
			}
			ids[r.Id] = struct{}{}
		}

		routes[i].PathParams = extractPathParams(routes[i].Path)

		seen[key] = r.Id

	}

	return routes, nil
}

func normalizePath(p string) string {
	parts := strings.Split(p, "/")
	for i, part := range parts {
		if strings.HasPrefix(part, ":") {
			parts[i] = ":*"
		}
	}
	return strings.Join(parts, "/")
}

func validateMethod(m string) error {
	m = strings.ToUpper(m)

	if _, ok := allowedMethods[m]; !ok {
		return fmt.Errorf("invalid HTTP method: %s", m)
	}
	return nil
}

func validatePath(path string) error {
	if path == "" {
		return fmt.Errorf("path is empty")
	}

	if !strings.HasPrefix(path, "/") {
		return fmt.Errorf("path must start with '/' (%s)", path)
	}

	if strings.Contains(path, "//") {
		return fmt.Errorf("path contains '//' (%s)", path)
	}

	parts := strings.Split(path, "/")
	seen := map[string]struct{}{}

	for _, p := range parts {
		if p == "" {
			continue
		}

		if strings.Contains(p, ":") {
			if !strings.HasPrefix(p, ":") {
				return fmt.Errorf("invalid param segment '%s' in path %s", p, path)
			}

			name := p[1:]
			if name == "" {
				return fmt.Errorf("empty param name in path %s", path)
			}

			if strings.Contains(name, ":") {
				return fmt.Errorf("invalid param name '%s' in path %s", name, path)
			}

			if _, exists := seen[name]; exists {
				return fmt.Errorf("duplicate param '%s' in path %s", name, path)
			}

			seen[name] = struct{}{}
		}
	}

	return nil
}

func extractPathParams(path string) []string {
	parts := strings.Split(path, "/")
	var params []string
	for _, p := range parts {
		if strings.HasPrefix(p, ":") && len(p) > 1 {
			params = append(params, p[1:])
		}
	}
	return params
}
