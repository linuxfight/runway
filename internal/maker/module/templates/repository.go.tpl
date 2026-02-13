package {{ .Name }}

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}
