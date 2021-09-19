package helper

import (
	"ujklm23/restful_api/model/entity"
	"ujklm23/restful_api/model/web"
)

func ToNoteResponses(notes []entity.Note) []web.NoteResponse {
	var noteResponses []web.NoteResponse

	for _, note := range notes {
		noteResponses = append(noteResponses, ToNoteResponse(note))
	}

	return noteResponses
}
