package implementation

import (
	"net/http"

	"github.com/imnzr/DevOps-Project-01/todo-list-api/backend/service/service"
	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func (controller UserControllerImpl) Create(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
}
func (controller UserControllerImpl) Update(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
}
func (controller UserControllerImpl) Delete(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
}
func (controller UserControllerImpl) FindById(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
}
func (controller UserControllerImpl) FindByAll(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
}
