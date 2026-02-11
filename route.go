package runway

import echo "github.com/labstack/echo/v4"

// Method is an HTTP method string, e.g. "GET", "POST".
// Compatible with net/http constants (http.MethodGet, etc.).
type Method string

type Route struct {
	Method      string
	Path        string
	Request     any
	Response    any
	Middlewares []echo.MiddlewareFunc
}

type RoutesProvider interface {
	Routes() []Route
}

func MW(m ...echo.MiddlewareFunc) []echo.MiddlewareFunc { return m }
