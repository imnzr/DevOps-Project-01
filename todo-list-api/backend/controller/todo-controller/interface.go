package todocontroller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type TodoController interface {
	Create(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateByTitle(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateByDescription(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByAll(writter http.ResponseWriter, request *http.Request, params httprouter.Params)
}
