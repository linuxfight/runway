package notes

import (
	"context"

	api "{{ .ModulePath }}/internal/modules/notes/api"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// Business logic lives here.
// Service must not depend on HTTP/Echo.

func (s *Service) CreateNote(
	ctx context.Context,
	req api.CreateNoteRequest,
) (api.NoteResponse, error) {
	id, err := s.repo.CreateNote(ctx, req.Title, req.Content)
	if err != nil {
		return api.NoteResponse{}, err
	}

	return api.NoteResponse{
		ID:      id,
		Title:   req.Title,
		Content: req.Content,
	}, nil
}

func (s *Service) GetNote(
	ctx context.Context,
	id string,
) (api.NoteResponse, error) {
	note, err := s.repo.GetNote(ctx, id)
	if err != nil {
		return api.NoteResponse{}, err
	}

	return api.NoteResponse{
		ID:      note.ID,
		Title:   note.Title,
		Content: note.Content,
	}, nil
}
