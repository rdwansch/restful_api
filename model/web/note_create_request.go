package web

type NoteCreateRequest struct {
	Name    string `validate:"required" json:"name"`
	Content string `validate:"required" json:"content"`
}
