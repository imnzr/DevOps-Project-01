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
	tx, err := service.DB.Begin()
	if err != nil {
		panic(fmt.Errorf("failed to begin transaction: %w", err))
	}
	defer helper.HandleTx(tx)

	user, err := service.UserRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(fmt.Errorf("user not found: %w", err))
	}

	user.Username = request.Username
	user.Email = request.Email
	user.Password = request.Password

	userUpdate, err := service.UserRepository.Update(ctx, tx, user)
	if err != nil {
		panic(fmt.Errorf("failed to update user: %w", err))
	}

	return response.UserResponse{
		Username: userUpdate.Username,
		Email:    userUpdate.Email,
	}

}
func (service *UserServiceImplementation) Delete(ctx context.Context, userId int) {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(fmt.Errorf("failed to begin transaction: %w", err))
	}
	defer helper.HandleTx(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(fmt.Errorf("failed to delete user: %w", err))
	}

	service.UserRepository.Delete(ctx, tx, user)
}
func (service *UserServiceImplementation) FindById(ctx context.Context, userId int) response.UserResponse {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(fmt.Errorf("failed to begin transaction: %w", err))
	}
	defer helper.HandleTx(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(fmt.Errorf("user not found: %w", err))
	}

	return response.UserResponse{
		Username: user.Username,
		Email:    user.Email,
	}
}
func (service *UserServiceImplementation) FindByAll(ctx context.Context) []response.UserResponse {
	tx, err := service.DB.Begin()
	if err != nil {
		panic(fmt.Errorf("failed to begin transaction: %w", err))
	}
	defer helper.HandleTx(tx)

	users := service.UserRepository.FindByAll(ctx, tx)

	var userResponse []response.UserResponse
	for _, users := range users {
		userResponse = append(userResponse, response.UserResponse{
			Username: users.Username,
			Email:    users.Email,
		})
	}

	return userResponse
}
