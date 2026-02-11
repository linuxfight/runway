package ir

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var titleCaser = cases.Title(language.Und)

// HandlerMethod returns generated handler method name,
// e.g. GetUsersById, PostCourses
func (r Route) HandlerMethod() string {
	method := handlerMethodPrefix(r.MethodExpr)
	path := pathToMethod(r.Path)
	return method + path
}

// ---- helpers ----

func handlerMethodPrefix(expr string) string {
	base := methodBaseName(expr)

	fmt.Println(base)

	switch base {
	case "get", "post", "put", "patch", "delete", "head", "options":
		return titleCaser.String(base)
	default:
		// safe fallback, never panic generator
		return "Handle"
	}
}

// methodBaseName extracts last identifier from expression
// examples:
//
//	http.MethodGet -> get
//	myhttp.GET     -> get
//	GET            -> get
func methodBaseName(expr string) string {
	expr = strings.TrimSpace(expr)

	if i := strings.LastIndex(expr, "."); i >= 0 {
		expr = expr[i+1:]
	}

	return strings.ToLower(expr)
}

func pathToMethod(path string) string {
	parts := strings.Split(strings.Trim(path, "/"), "/")

	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if strings.HasPrefix(p, ":") {
			name := strings.TrimPrefix(p, ":")
			out = append(out, "By"+toExported(name))
		} else {
			out = append(out, toExported(p))
		}
	}

	return strings.Join(out, "")
}

func toExported(s string) string {
	parts := strings.FieldsFunc(s, func(r rune) bool {
		return r == '_' || r == '-'
	})

	for i := range parts {
		parts[i] = titleCaser.String(parts[i])
	}

	return strings.Join(parts, "")
}
