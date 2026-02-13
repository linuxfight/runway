package api

type Create{{ .Title }}Request struct {
	Name string `json:"name" validate:"required,min=3,max=120"`
}

type {{ .Title }}Response struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
