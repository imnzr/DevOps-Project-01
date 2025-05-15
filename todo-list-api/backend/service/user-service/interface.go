package userservice

import (
	"context"

	web "github.com/imnzr/DevOps-Project-01/todo-list-api/backend/models/web/request"
	"github.com/imnzr/DevOps-Project-01/todo-list-api/backend/models/web/response"
)

type UserService interface {
	Create(ctx context.Context, request web.UserCreateRequest) response.UserResponse
	Update(ctx context.Context, request web.UserUpdateRequest) response.UserResponse
	Delete(ctx context.Context, userId int)
	FindById(ctx context.Context, userId int) response.UserResponse
	FindByAll(ctx context.Context) []response.UserResponse

	// Login service
	Login(ctx context.Context, request web.UserLoginRequest) response.UserResponse
}
