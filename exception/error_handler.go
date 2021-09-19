package exception

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"ujklm23/restful_api/helper"
	"ujklm23/restful_api/model/web"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if notFoundErr(writer, request, err) {
		return
	} else if ValidateErr(writer, request, err) {
		return
	}

	internalServerErr(writer, request, err)
}

func internalServerErr(writer http.ResponseWriter, _ *http.Request, err interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		StatusCode: http.StatusInternalServerError,
		Status:     "INTERNAL SERVER ERROR",
		Data:       err,
	}

	helper.WriteJSONToBody(writer, webResponse)

}

func notFoundErr(writer http.ResponseWriter, _ *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			StatusCode: http.StatusBadRequest,
			Status:     "BAD REQUEST",
			Data:       exception.Error,
		}

		helper.WriteJSONToBody(writer, webResponse)
		return true
	}

	return false
}

func ValidateErr(writer http.ResponseWriter, _ *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			StatusCode: http.StatusBadRequest,
			Status:     "BAD REQUEST",
			Data:       exception.Error(),
		}

		helper.WriteJSONToBody(writer, webResponse)
		return true
	}
	return false
}
