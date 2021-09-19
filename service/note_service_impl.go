package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"ujklm23/restful_api/helper"
	"ujklm23/restful_api/model/entity"
	"ujklm23/restful_api/model/web"
	"ujklm23/restful_api/repository"
)

type NoteServiceImpl struct {
	repository.NoteRepository
	*sql.DB
	*validator.Validate
}

func NewNoteServiceImpl(db *sql.DB, validate *validator.Validate, noteRepository repository.NoteRepository) NoteService {
	return &NoteServiceImpl{
		NoteRepository: noteRepository,
		DB:             db,
		Validate:       validate,
	}
}

func (service *NoteServiceImpl) Create(ctx context.Context, request web.NoteCreateRequest) web.NoteResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	myNote := entity.Note{
		Name:    request.Name,
		Content: request.Content,
	}

	note := service.NoteRepository.Create(ctx, tx, myNote)
	return helper.ToNoteResponse(note)
}

func (service *NoteServiceImpl) Update(ctx context.Context, request web.NoteUpdateRequest) web.NoteResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.NoteRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	myNote := entity.Note{
		Id:      request.Id,
		Name:    request.Name,
		Content: request.Content,
	}

	note := service.NoteRepository.Update(ctx, tx, myNote)
	return helper.ToNoteResponse(note)
}

func (service *NoteServiceImpl) Delete(ctx context.Context, requestId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.NoteRepository.FindById(ctx, tx, requestId)
	helper.PanicIfError(err)

	service.NoteRepository.Delete(ctx, tx, requestId)
}

func (service *NoteServiceImpl) FindAll(ctx context.Context) []web.NoteResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	notes := service.NoteRepository.FindAll(ctx, tx)
	return helper.ToNoteResponses(notes)
}

func (service *NoteServiceImpl) FindById(ctx context.Context, requestId int) web.NoteResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	note, err := service.NoteRepository.FindById(ctx, tx, requestId)
	helper.PanicIfError(err)

	return helper.ToNoteResponse(note)
}
