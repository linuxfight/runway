package openapi

import (
	"bytes"
	"strconv"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"gopkg.in/yaml.v3"
)

func convertPath(p string) string {
	parts := strings.Split(p, "/")
	for i, part := range parts {
		if strings.HasPrefix(part, ":") {
			parts[i] = "{" + part[1:] + "}"
		}
	}
	return strings.Join(parts, "/")
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

func strPtr(s string) *string { return &s }

func yamlMarshal(v any) ([]byte, error) {
	var buf bytes.Buffer

	enc := yaml.NewEncoder(&buf)
	enc.SetIndent(2)

	if err := enc.Encode(v); err != nil {
		return nil, err
	}

	if err := enc.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func isString(s *openapi3.Schema) bool {
	return s.Type != nil && len(*s.Type) > 0 && (*s.Type)[0] == "string"
}

func parseUint(v string) uint64 {
	n, _ := strconv.ParseUint(v, 10, 64)
	return n
}

func parseFloat(v string) float64 {
	n, _ := strconv.ParseFloat(v, 64)
	return n
}

func uint64Ptr(v uint64) *uint64    { return &v }
func float64Ptr(v float64) *float64 { return &v }
