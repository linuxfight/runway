package runway

// Method is an HTTP method string, e.g. "GET", "POST".
// Compatible with net/http constants (http.MethodGet, etc.).
type Method string

type Route struct {
	Method   string
	Path     string
	Request  any
	Response any

	Summary     string
	Description string
	Tags        []string

	Raw bool
}

type RoutesProvider interface {
	Routes() map[string]Route
}
