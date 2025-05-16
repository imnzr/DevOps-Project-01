package todocontroller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/imnzr/DevOps-Project-01/todo-list-api/backend/helper"
	web "github.com/imnzr/DevOps-Project-01/todo-list-api/backend/models/web/request"
	"github.com/imnzr/DevOps-Project-01/todo-list-api/backend/models/web/response"
	todoservice "github.com/imnzr/DevOps-Project-01/todo-list-api/backend/service/todo-service"
	"github.com/julienschmidt/httprouter"
)

type TodoControllerImpl struct {
	TodoService todoservice.TodoService
}

// Create implements TodoController.
func (controller *TodoControllerImpl) Create(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	todoCreateRequest := web.TodoCreateRequest{}
	err := decoder.Decode(&todoCreateRequest)
	if err != nil {
		http.Error(writter, "Invalid request body", http.StatusBadRequest)
		return
	}

	todoResponse, err := controller.TodoService.Create(request.Context(), todoCreateRequest)
	if err != nil {
		http.Error(writter, "Failed to create todo"+err.Error(), http.StatusInternalServerError)
		return
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
	todoId := params.ByName("todoId")
	id, err := strconv.Atoi(todoId)
	if err != nil {
		helper.WriteErrorResponse(writter, http.StatusBadRequest, "invalid user id: "+err.Error())
		return
	}

	todoResponse, err := controller.TodoService.FindById(request.Context(), id)
	if err != nil {
		// Deteksi error not found
		if strings.Contains(err.Error(), "not found") {
			helper.WriteErrorResponse(writter, http.StatusNotFound, err.Error())
			return
		}
		// Default : error internal
		helper.WriteErrorResponse(writter, http.StatusInternalServerError, "failed to retrieve todo: "+err.Error())
	}

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   todoResponse,
	}
	writter.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(writter).Encode(webResponse)
	if err != nil {
		helper.WriteErrorResponse(writter, http.StatusInternalServerError, "failed to encode response: "+err.Error())
	}
	// encoder := json.NewEncoder(writter)
	// err = encoder.Encode(webResponse)
	// helper.PanicIfError(err)
}

// UpdateByDescription implements TodoController.
func (controller *TodoControllerImpl) UpdateByDescription(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	todoUpdateRequest := web.TodoUpdateRequest{}
	if err := decoder.Decode(&todoUpdateRequest); err != nil {
		helper.WriteErrorResponse(writter, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	todoId := params.ByName("todoId")
	id, err := strconv.Atoi(todoId)
	if err != nil {
		helper.WriteErrorResponse(writter, http.StatusBadRequest, "invalid todo ID: "+err.Error())
		return
	}

	todoUpdateRequest.Id = id

	todoResponse, err := controller.TodoService.UpdateDescription(request.Context(), todoUpdateRequest)
	if err != nil {
		helper.WriteErrorResponse(writter, http.StatusInternalServerError, "failed to update description: "+err.Error())
		return
	}

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   todoResponse,
	}

	writter.Header().Add("Content-Type", "application/json")
	writter.Header().Add("Content-Type", "application/json")
	json.NewEncoder(writter).Encode(webResponse)
	// encoder := json.NewEncoder(writter)
	// err = encoder.Encode(webResponse)
	// helper.PanicIfError(err)
}

// UpdateByTitle implements TodoController.
func (controller *TodoControllerImpl) UpdateByTitle(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	todoUpdateRequest := web.TodoUpdateRequest{}
	if err := decoder.Decode(&todoUpdateRequest); err != nil {
		helper.WriteErrorResponse(writter, http.StatusBadRequest, "invalid request body: "+err.Error())
		return
	}

	todoId := params.ByName("todoId")
	id, err := strconv.Atoi(todoId)
	if err != nil {
		helper.WriteErrorResponse(writter, http.StatusBadRequest, "invalid todo ID: "+err.Error())
		return
	}

	todoUpdateRequest.Id = id

	todoResponse, err := controller.TodoService.UpdateTitle(request.Context(), todoUpdateRequest)
	if err != nil {
		helper.WriteErrorResponse(writter, http.StatusInternalServerError, "failed to update todo: "+err.Error())
		return
	}

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   todoResponse,
	}

	writter.Header().Add("Content-Type", "application/json")
	json.NewEncoder(writter).Encode(webResponse)
	// encoder := json.NewEncoder(writter)
	// err = encoder.Encode(webResponse)
	// helper.PanicIfError(err)
}

func NewTodoController(todoService todoservice.TodoService) TodoController {
	return &TodoControllerImpl{
		TodoService: todoService,
	}
}
