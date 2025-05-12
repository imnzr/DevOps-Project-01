package implementation

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/imnzr/DevOps-Project-01/todo-list-api/backend/helper"
	"github.com/imnzr/DevOps-Project-01/todo-list-api/backend/models/domain"
	web "github.com/imnzr/DevOps-Project-01/todo-list-api/backend/models/web/request"
	"github.com/imnzr/DevOps-Project-01/todo-list-api/backend/models/web/response"
	"github.com/imnzr/DevOps-Project-01/todo-list-api/backend/repository/repo"
)

type UserServiceImplementation struct {
	// Repository untuk memanipulasi datanya
	UserRepository repo.UserRepository
	// Koneksi database
	DB *sql.DB
}

func (service *UserServiceImplementation) Create(ctx context.Context, request web.UserCreateRequest) response.UserResponse {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(fmt.Errorf("failed to begin transaction: %w", err))
	}
	defer helper.HandleTx(tx)

	user := domain.User{
		Email:    request.Username,
		Username: request.Email,
	}
	savedUser, err := service.UserRepository.Save(ctx, tx, user)
	if err != nil {
		panic(err)
	}

	return response.UserResponse{
		Id:       savedUser.Id,
		Username: savedUser.Username,
		Email:    savedUser.Email,
	}
}
func (service *UserServiceImplementation) Update(ctx context.Context, request web.UserUpdateRequest) response.UserResponse {
	panic("Impl here")
}
func (service *UserServiceImplementation) Delete(ctx context.Context, userId int) {
	panic("Impl  here")
}
func (service *UserServiceImplementation) FindById(ctx context.Context, userId int) response.UserResponse {
	panic("Impl here")
}
func (service *UserServiceImplementation) FindByAll(ctx context.Context) []response.UserResponse {
	panic("Impl here")
}
