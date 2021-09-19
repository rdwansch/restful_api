package repository

import (
	"context"
	"database/sql"
	"ujklm23/restful_api/model/entity"
)

type NoteRepository interface {
	Create(ctx context.Context, tx *sql.Tx, myNote entity.Note) entity.Note
	Update(ctx context.Context, tx *sql.Tx, myNote entity.Note) entity.Note
	Delete(ctx context.Context, tx *sql.Tx, myNoteId int)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Note
	FindById(ctx context.Context, tx *sql.Tx, myNoteId int) (entity.Note, error)
}
