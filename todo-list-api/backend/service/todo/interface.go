package todo

import (
	"context"

	web "github.com/imnzr/DevOps-Project-01/todo-list-api/backend/models/web/request"
	"github.com/imnzr/DevOps-Project-01/todo-list-api/backend/models/web/response"
)

type TodoService interface {
	Create(ctx context.Context, request web.TodoCreateRequest) (response.TodoResponse, error)
	UpdateTitle(ctx context.Context, request web.TodoUpdateRequest) (response.TodoResponse, error)
	UpdateDescription(ctx context.Context, request web.TodoUpdateRequest) (response.TodoResponse, error)
	Delete(ctx context.Context, todoId int)
	FindById(ctx context.Context, todoId int) (response.TodoResponse, error)
	FindByAll(ctx context.Context) []response.TodoResponse
}
