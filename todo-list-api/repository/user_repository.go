package repository

import (
	"context"

	"github.com/imnzr/DevOps-Project-01/todo-list-api/models"
)

type UserRepository interface {
	Insert(ctx context.Context, user models.User) (models.User, error)
	GetById(ctx context.Context, id int64) (models.User, error)
	UpdateById(ctx context.Context, user models.User)
	DeleteById(ctx context.Context, id int64) error
}
