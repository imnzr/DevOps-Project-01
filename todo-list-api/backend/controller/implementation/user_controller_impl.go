package implementation

import (
	"encoding/json"
	"fmt"
	"net/http"

	web "github.com/imnzr/DevOps-Project-01/todo-list-api/backend/models/web/request"
	"github.com/imnzr/DevOps-Project-01/todo-list-api/backend/models/web/response"
	"github.com/imnzr/DevOps-Project-01/todo-list-api/backend/service/service"
	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func (controller UserControllerImpl) Create(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	userCreateRequest := web.UserCreateRequest{}
	err := decoder.Decode(&userCreateRequest)
	if err != nil {
		// error handler here
	}

	userResponse := controller.UserService.Create(request.Context(), userCreateRequest)
	WebResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	writter.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writter)
	err = encoder.Encode(WebResponse)
	if err != nil {
		panic(fmt.Errorf("error hereL %w", err))
	}
}
func (controller UserControllerImpl) Update(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
}
func (controller UserControllerImpl) Delete(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
}
func (controller UserControllerImpl) FindById(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
}
func (controller UserControllerImpl) FindByAll(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
}
