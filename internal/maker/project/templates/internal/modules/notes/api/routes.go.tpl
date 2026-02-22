package api

import (
	"net/http"

	"github.com/cryingcatscloud/runway"
)

type Routes struct{}

func (Routes) Routes() map[string]runway.Route {
	return map[string]runway.Route{
		"CreateNote": {
			Method:      http.MethodPost,
			Path:        "/notes",
			Request:     CreateNoteRequest{},
			Response:    NoteResponse{},
			Summary:     "Create note",
			Description: "Creates a new note",
			Tags:        []string{"notes"},
		},
		"GetNote": {
			Method:   http.MethodGet,
			Path:     "/notes/:id",
			Request:  GetNoteRequest{},
			Response: NoteResponse{},
			Summary:  "Get note by ID",
			Tags:     []string{"notes"},
		},
	}
}
