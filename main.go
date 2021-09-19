package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"ujklm23/restful_api/config"
	"ujklm23/restful_api/controller"
	"ujklm23/restful_api/exception"
	"ujklm23/restful_api/helper"
	"ujklm23/restful_api/middleware"
	"ujklm23/restful_api/repository"
	"ujklm23/restful_api/service"
)

func main() {
	db := config.NewDB()
	validate := validator.New()
	noteRepository := repository.NewNoteRepositoryImpl()
	noteService := service.NewNoteServiceImpl(db, validate, noteRepository)
	noteController := controller.NewNoteController(noteService)

	router := httprouter.New()

	router.POST("/api/v1/notes", noteController.Create)
	router.PUT("/api/v1/notes/:noteId", noteController.Update)
	router.DELETE("/api/v1/notes/:noteId", noteController.Delete)
	router.GET("/api/v1/notes", noteController.FindAll)
	router.GET("/api/v1/notes/:noteId", noteController.FindById)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: middleware.NewMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
