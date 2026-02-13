package notes

import (
	"context"

	"github.com/labstack/echo/v4"

	api "{{ .ModulePath }}/internal/modules/notes/api"
)

type Controller struct {
	service *Service
}

func NewController(e *echo.Echo, service *Service) *Controller {
	controller := &Controller{service: service}
	api.RegisterRoutes(e, controller)
	return controller
}

// NOTE:
// OpenAPI routes are registered via generated code (server.gen.go).
// Controller only implements the generated interface: notesHandlers.

// CreateNote handles POST /notes
func (c *Controller) CreateNote(
	ctx context.Context,
	req api.CreateNoteRequest,
) (api.NoteResponse, error) {
	return c.service.CreateNote(ctx, req)
}

// GetNote handles GET /notes/:id
func (c *Controller) GetNote(
	ctx context.Context,
) (api.NoteResponse, error) {
	return c.service.GetNote(ctx, req.ID)
}
