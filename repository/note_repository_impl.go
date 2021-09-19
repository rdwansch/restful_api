package repository

import (
	"context"
	"database/sql"
	"errors"
	"ujklm23/restful_api/helper"
	"ujklm23/restful_api/model/entity"
)

type NoteRepositoryImpl struct {
}

func NewNoteRepositoryImpl() NoteRepository {
	return &NoteRepositoryImpl{}
}

func (repository *NoteRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, myNote entity.Note) entity.Note {
	SQL := "INSERT INTO note(name, content) VALUES (?, ?)"
	result, err := tx.ExecContext(ctx, SQL, myNote.Name, myNote.Content)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	myNote.Id = int(id)
	return myNote
}

func (repository *NoteRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, myNote entity.Note) entity.Note {
	SQL := "UPDATE note SET name = ?, content = ? WHERE id = ?"
	result, err := tx.ExecContext(ctx, SQL, myNote.Name, myNote.Content, myNote.Id)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	myNote.Id = int(id)
	return myNote
}

func (repository *NoteRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, myNoteId int) {
	SQL := "DELETE FROM note WHERE id = ?"
	result, err := tx.ExecContext(ctx, SQL, myNoteId)
	helper.PanicIfError(err)

	_, err = result.RowsAffected()
	helper.PanicIfError(err)

}

func (repository *NoteRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Note {
	SQL := "SELECT name, content, id FROM note"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var notes []entity.Note

	if rows.Next() {
		note := entity.Note{}
		err := rows.Scan(&note.Name, &note.Content, &note.Id)
		helper.PanicIfError(err)

		notes = append(notes, note)
	}

	return notes
}

func (repository *NoteRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, myNoteId int) (entity.Note, error) {
	SQL := "SELECT name, content, id FROM note WHERE id = ?"
	row, err := tx.QueryContext(ctx, SQL, myNoteId)
	helper.PanicIfError(err)
	defer row.Close()

	note := entity.Note{}

	if row.Next() {
		err = row.Scan(&note.Name, &note.Content, &note.Id)
		helper.PanicIfError(err)
		return note, nil
	}

	return note, errors.New("note not found")
}
