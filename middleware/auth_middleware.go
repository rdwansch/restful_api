package middleware

import (
	"net/http"
	"ujklm23/restful_api/helper"
	"ujklm23/restful_api/model/web"
)

type Middleware struct {
	http.Handler
}

func NewMiddleware(handler http.Handler) *Middleware {
	return &Middleware{handler}
}

func (middleware *Middleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("apikey") == "nyeri-sendi" {
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			StatusCode: http.StatusUnauthorized,
			Status:     "UNAUTHORIZED",
			Data:       "missing or invalid API key",
		}

		helper.WriteJSONToBody(writer, webResponse)
	}
}
