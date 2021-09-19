package middleware

import "net/http"

type Middleware struct {
	http.Handler
}

func NewMiddleware(handler http.Handler) *Middleware {
	return &Middleware{handler}
}

func (middleware *Middleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	middleware.Handler.ServeHTTP(writer, request)
}
