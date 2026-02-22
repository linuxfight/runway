package api

type CreateNoteRequest struct {
	Title   string `json:"title" validate:"required,min=3,max=120"`
	Content string `json:"content"`
}

type GetNoteRequest struct {
	ID string `path:"id"`
}

type NoteResponse struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
