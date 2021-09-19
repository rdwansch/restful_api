package web

type NoteUpdateRequest struct {
	Id      int    `validate:"required" json:"id"`
	Name    string `validate:"required" json:"name"`
	Content string `validate:"required" json:"content"`
}
