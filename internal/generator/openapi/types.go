package openapi

type FieldMeta struct {
	Name     string    `json:"name"`
	Type     string    `json:"type"`
	JSON     string    `json:"json,omitempty"`
	Query    string    `json:"query,omitempty"`
	Param    string    `json:"param,omitempty"`
	Header   string    `json:"header,omitempty"`
	Validate string    `json:"validate,omitempty"`
	Nested   *TypeMeta `json:"nested,omitempty"`
}

type TypeMeta struct {
	Name   string      `json:"name"`
	Fields []FieldMeta `json:"fields,omitempty"`
}

type RouteMeta struct {
	ID          string   `json:"id"`
	Method      string   `json:"method"`
	Path        string   `json:"path"`
	Summary     string   `json:"summary,omitempty"`
	Description string   `json:"description,omitempty"`
	Tags        []string `json:"tags,omitempty"`

	Request  *TypeMeta `json:"request,omitempty"`
	Response *TypeMeta `json:"response,omitempty"`

	StatusCode int `json:"status_code"`
}
