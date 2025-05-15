package todo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/imnzr/DevOps-Project-01/todo-list-api/backend/helper"
	web "github.com/imnzr/DevOps-Project-01/todo-list-api/backend/models/web/request"
	"github.com/imnzr/DevOps-Project-01/todo-list-api/backend/models/web/response"
	"github.com/imnzr/DevOps-Project-01/todo-list-api/backend/service/todo"
	"github.com/julienschmidt/httprouter"
)

type TodoControllerImpl struct {
	TodoService todo.TodoService
}

// Create implements TodoController.
func (controller *TodoControllerImpl) Create(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	todoCreateRequest := web.TodoCreateRequest{}
	err := decoder.Decode(&todoCreateRequest)
	if err != nil {
		// error handler here
	}

	todoResponse, err := controller.TodoService.Create(request.Context(), todoCreateRequest)
	if err != nil {
		// error handler here
	}

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   todoResponse,
	}

	writter.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writter)
	err = encoder.Encode(webResponse)
	if err != nil {
		// error handler here
	}
}

// Delete implements TodoController.
func (controller *TodoControllerImpl) Delete(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoId := params.ByName("todoId")
	id, err := strconv.Atoi(todoId)
	if err != nil {
		panic(fmt.Errorf("todo not found: %w", err))
	}

	controller.TodoService.Delete(request.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
	}

	writter.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writter)
	err = encoder.Encode(webResponse)
	if err != nil {
		panic(fmt.Errorf("error handler here: %s", err))
	}
}

// FindByAll implements TodoController.
func (controller *TodoControllerImpl) FindByAll(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	todoResponse := controller.TodoService.FindByAll(request.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   todoResponse,
	}

	writter.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writter)
	err := encoder.Encode(webResponse)
	if err != nil {
		panic(fmt.Errorf("error handler here: %w", err))
	}
}

// FindById implements TodoController.
func (controller *TodoControllerImpl) FindById(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	todoResponse, err := controller.TodoService.FindById(request.Context(), id)
	helper.PanicIfError(err)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   todoResponse,
	}
	writter.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writter)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

// UpdateByDescription implements TodoController.
func (controller *TodoControllerImpl) UpdateByDescription(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	todoUpdateRequest := web.TodoUpdateRequest{}
	err := decoder.Decode(&todoUpdateRequest)
	helper.PanicIfError(err)

	todoId := params.ByName("todoId")
	id, err := strconv.Atoi(todoId)
	helper.PanicIfError(err)

	todoUpdateRequest.Id = id

	todoResponse, err := controller.TodoService.UpdateDescription(request.Context(), todoUpdateRequest)
	helper.PanicIfError(err)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   todoResponse,
	}

	writter.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writter)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

// UpdateByTitle implements TodoController.
func (controller *TodoControllerImpl) UpdateByTitle(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	todoUpdateRequest := web.TodoUpdateRequest{}
	err := decoder.Decode(&todoUpdateRequest)
	helper.PanicIfError(err)

	todoId := params.ByName("todoId")
	id, err := strconv.Atoi(todoId)
	helper.PanicIfError(err)

	todoUpdateRequest.Id = id

	todoResponse, err := controller.TodoService.UpdateTitle(request.Context(), todoUpdateRequest)
	helper.PanicIfError(err)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   todoResponse,
	}

	writter.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writter)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func NewTodoController(todoService todo.TodoService) TodoController {
	return &TodoControllerImpl{
		TodoService: todoService,
	}
}
