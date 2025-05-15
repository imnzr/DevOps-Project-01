package todo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/imnzr/DevOps-Project-01/todo-list-api/backend/helper"
	"github.com/imnzr/DevOps-Project-01/todo-list-api/backend/models/domain"
	web "github.com/imnzr/DevOps-Project-01/todo-list-api/backend/models/web/request"
	"github.com/imnzr/DevOps-Project-01/todo-list-api/backend/models/web/response"
	"github.com/imnzr/DevOps-Project-01/todo-list-api/backend/repository/todo"
)

type TodoServiceImpl struct {
	TodoRepository todo.TodoRepository
	DB             *sql.DB
}

func NewTodoService(todoRepository todo.TodoRepository, DB *sql.DB) TodoService {
	return &TodoServiceImpl{
		TodoRepository: todoRepository,
		DB:             DB,
	}
}

// Create implements TodoService.
func (service *TodoServiceImpl) Create(ctx context.Context, request web.TodoCreateRequest) (response.TodoResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return response.TodoResponse{}, fmt.Errorf("failed to begin transcatio: %w", err)
	}
	defer helper.HandleTx(tx)

	todo := domain.Todo{
		Title:       request.Title,
		Description: request.Description,
	}

	todo, err = service.TodoRepository.Save(ctx, tx, todo)
	if err != nil {
		return response.TodoResponse{}, fmt.Errorf("failed to save todo (title: %s): %w", request.Title, err)
	}

	return response.TodoResponse{
		Title:       todo.Title,
		Description: todo.Description,
	}, nil
}

// Delete implements TodoService.
func (service *TodoServiceImpl) Delete(ctx context.Context, todoId int) {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(fmt.Errorf("failed to begin transcations: %w", err))
	}
	defer helper.HandleTx(tx)

	todo, err := service.TodoRepository.FindById(ctx, tx, todoId)
	if err != nil {
		panic(fmt.Errorf("todo not found: %w", err))
	}

	service.TodoRepository.Delete(ctx, tx, todo)
}

// FindByAll implements TodoService.
func (service *TodoServiceImpl) FindByAll(ctx context.Context) []response.TodoResponse {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(fmt.Errorf("failed to begin transcations: %w", err))
	}
	defer helper.HandleTx(tx)

	todos, err := service.TodoRepository.FindByAll(ctx, tx)
	if err != nil {
		// error handler here
	}

	var todoResponse []response.TodoResponse
	for _, todos := range todos {
		todoResponse = append(todoResponse, response.TodoResponse{
			Title:       todos.Title,
			Description: todos.Description,
		})
	}

	return todoResponse
}

// FindById implements TodoService.
func (service *TodoServiceImpl) FindById(ctx context.Context, todoId int) (response.TodoResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(fmt.Errorf("failed to begin transcations: %w", err))
	}
	defer helper.HandleTx(tx)

	todo, err := service.TodoRepository.FindById(ctx, tx, todoId)
	if err != nil {
		panic(fmt.Errorf("todo not found: %w", err))
	}

	return response.TodoResponse{
		Title:       todo.Title,
		Description: todo.Description,
	}, nil
}

// UpdateDescription implements TodoService.
func (service *TodoServiceImpl) UpdateDescription(ctx context.Context, request web.TodoUpdateRequest) (response.TodoResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return response.TodoResponse{}, fmt.Errorf("failed to begin transcatio: %w", err)
	}
	defer helper.HandleTx(tx)

	todo, err := service.TodoRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		// error handler here
	}

	todo.Description = request.Description

	userUpdateDescription, err := service.TodoRepository.UpdateDescription(ctx, tx, todo)
	if err != nil {
		// error handler here
	}

	return response.TodoResponse{
		Description: userUpdateDescription.Description,
	}, nil

}

// UpdateTitle implements TodoService.
func (service *TodoServiceImpl) UpdateTitle(ctx context.Context, request web.TodoUpdateRequest) (response.TodoResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return response.TodoResponse{}, fmt.Errorf("failed to begin transcatio: %w", err)
	}
	defer helper.HandleTx(tx)

	todo, err := service.TodoRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		// error handle here
	}

	todo.Title = request.Title

	userUpdateTitle, err := service.TodoRepository.UpdateTitle(ctx, tx, todo)
	if err != nil {
		// error handle here
	}

	return response.TodoResponse{
		Title: userUpdateTitle.Title,
	}, nil
}
