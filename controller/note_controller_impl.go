package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"ujklm23/restful_api/helper"
	"ujklm23/restful_api/model/web"
	"ujklm23/restful_api/service"
)

type NoteControllerImpl struct {
	service.NoteService
}

func NewNoteController(noteService service.NoteService) NoteController {
	return &NoteControllerImpl{noteService}
}

func (controller *NoteControllerImpl) Create(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	noteCreateRequest := web.NoteCreateRequest{}
	helper.ReadJSONFromRequest(request, &noteCreateRequest)

	noteResponse := controller.NoteService.Create(request.Context(), noteCreateRequest)

	webResponse := web.WebResponse{
		StatusCode: 200,
		Status:     "OK",
		Data:       noteResponse,
	}
	helper.WriteJSONToBody(writer, webResponse)
}

func (controller *NoteControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	noteUpdateRequest := web.NoteUpdateRequest{}
	helper.ReadJSONFromRequest(request, &noteUpdateRequest)

	noteId := params.ByName("noteId")
	id, err := strconv.Atoi(noteId)
	helper.PanicIfError(err)

	noteUpdateRequest.Id = id
	noteResponse := controller.NoteService.Update(request.Context(), noteUpdateRequest)
	noteResponse.Id = id

	webResponse := web.WebResponse{
		StatusCode: 200,
		Status:     "OK",
		Data:       noteResponse,
	}
	helper.WriteJSONToBody(writer, webResponse)
}

func (controller *NoteControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	noteId, err := strconv.Atoi(params.ByName("noteId"))
	helper.PanicIfError(err)
	controller.NoteService.Delete(request.Context(), noteId)

	data := "Success Delete id: " + strconv.Itoa(noteId)

	webResponse := web.WebResponse{
		StatusCode: 200,
		Status:     "OK",
		Data:       data,
	}
	helper.WriteJSONToBody(writer, webResponse)
}

func (controller *NoteControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	noteResponses := controller.NoteService.FindAll(request.Context())

	webResponse := web.WebResponse{
		StatusCode: 200,
		Status:     "OK",
		Data:       noteResponses,
	}
	helper.WriteJSONToBody(writer, webResponse)
}

func (controller *NoteControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	noteId, err := strconv.Atoi(params.ByName("noteId"))
	helper.PanicIfError(err)

	noteResponse := controller.NoteService.FindById(request.Context(), noteId)
	webResponse := web.WebResponse{
		StatusCode: 200,
		Status:     "OK",
		Data:       noteResponse,
	}
	helper.WriteJSONToBody(writer, webResponse)
}
