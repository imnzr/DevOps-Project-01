package usercontroller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	web "github.com/imnzr/DevOps-Project-01/todo-list-api/backend/models/web/request"
	"github.com/imnzr/DevOps-Project-01/todo-list-api/backend/models/web/response"
	userservice "github.com/imnzr/DevOps-Project-01/todo-list-api/backend/service/user-service"
	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService userservice.UserService
}

func NewUserController(userService userservice.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}

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
		panic(fmt.Errorf("error here %w", err))
	}
}
func (controller UserControllerImpl) Update(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	userUpdateRequest := web.UserUpdateRequest{}
	err := decoder.Decode(&userUpdateRequest)
	if err != nil {
		panic(fmt.Errorf("error: %w", err))
	}
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		panic(fmt.Errorf("user not found: %w", err))
	}

	userUpdateRequest.Id = id

	userResponse := controller.UserService.Update(request.Context(), userUpdateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	writter.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writter)
	err = encoder.Encode(webResponse)
	if err != nil {
		panic(fmt.Errorf("error handler here %w", err))
	}
}
func (controller UserControllerImpl) Delete(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		panic(fmt.Errorf("user not found: %w", err))
	}

	controller.UserService.Delete(request.Context(), id)
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
func (controller UserControllerImpl) FindById(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	if err != nil {
		panic(fmt.Errorf("user not found: %w", err))
	}

	userResponse := controller.UserService.FindById(request.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	writter.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writter)
	err = encoder.Encode(webResponse)
	if err != nil {
		panic(fmt.Errorf("error handler here: %w", err))
	}
}
func (controller UserControllerImpl) FindByAll(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userResponse := controller.UserService.FindByAll(request.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	writter.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writter)
	err := encoder.Encode(webResponse)
	if err != nil {
		panic(fmt.Errorf("error handler here: %w", err))
	}
}

// Login implements UserController.
func (controller *UserControllerImpl) Login(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	loginRequest := web.UserLoginRequest{}
	err := decoder.Decode(&loginRequest)
	if err != nil {
		panic(fmt.Errorf("error here: %w", err))
	}

	userResponse := controller.UserService.Login(request.Context(), loginRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	writter.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writter)
	err = encoder.Encode(webResponse)
	if err != nil {
		panic(fmt.Errorf("error: %w", err))
	}

}
