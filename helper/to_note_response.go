package helper

import (
	"ujklm23/restful_api/model/entity"
	"ujklm23/restful_api/model/web"
)



func ToNoteResponse(noteEntity entity.Note) web.NoteResponse {
	noteResponse := web.NoteResponse{
		Id:      noteEntity.Id,
		Name:    noteEntity.Name,
		Content: noteEntity.Content,
	}

	return noteResponse
}
