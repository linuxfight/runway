package api

import (
	"net/http"

	"github.com/cryingcatscloud/runway"
)

type Routes struct{}

func (Routes) Routes() map[string]runway.Route {
	return map[string]runway.Route{
		"Create{{ .Title }}": {
			Method:   http.MethodPost,
			Path:     "/{{ .Name }}",
			Request:  Create{{ .Title }}Request{},
			Response: {{ .Title }}Response{},
		},
		"Get{{ .Title }}": {
			Method:   http.MethodGet,
			Path:     "/{{ .Name }}/:id",
			Response: {{ .Title }}Response{},
		},
	}
}
