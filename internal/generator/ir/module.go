package ir

type APIModule struct {
	Name   string
	Routes []Route
	DTOs   map[string]DTO
}

type Route struct {
	MethodExpr  string
	Path        string
	Middlewares []MiddlewareRef
	Request     *TypeRef
	Response    *TypeRef

	// derived
	PathParams []string
}

type MiddlewareRef struct {
	Expr string
}

type DTO struct {
	Name   string
	Fields []Field
}

type Field struct {
	Name     string
	Type     string
	Required bool
	Tags     map[string]string
}

type TypeRef struct {
	Name string
	Pkg  string
}
