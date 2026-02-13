package notes

import "context"

// Repository is responsible for data access.
// Later this can be backed by Ent.

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

type Note struct {
	ID      string
	Title   string
	Content string
}

func (r *Repository) CreateNote(
	ctx context.Context,
	title string,
	content string,
) (string, error) {
	// TODO: replace with Ent create
	return "note-id-1", nil
}

func (r *Repository) GetNote(
	ctx context.Context,
	id string,
) (*Note, error) {
	// TODO: replace with Ent query
	return &Note{
		ID:      id,
		Title:   "Example note",
		Content: "This is a demo note",
	}, nil
}
