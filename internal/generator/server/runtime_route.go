package server

type RuntimeRoute struct {
	Method   string `json:"method"`
	Path     string `json:"path"`
	Id       string `json:"id"`
	Request  string `json:"request"`
	Response string `json:"response"`

	PathParams []string

	Summary     string   `json:"summary"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`

	Raw bool `json:"raw"`
}
