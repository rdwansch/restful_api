package service

import (
	"context"
	"ujklm23/restful_api/model/web"
)

type NoteService interface {
	Create(ctx context.Context, request web.NoteCreateRequest) web.NoteResponse
	Update(ctx context.Context, request web.NoteUpdateRequest) web.NoteResponse
	Delete(ctx context.Context, requestId int)
	FindAll(ctx context.Context) []web.NoteResponse
	FindById(ctx context.Context, requestId int) web.NoteResponse
}
