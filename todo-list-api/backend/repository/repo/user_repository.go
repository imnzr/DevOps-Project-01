package repo

import (
	"context"
	"database/sql"

	"github.com/imnzr/DevOps-Project-01/todo-list-api/backend/models/domain"
)

// QUERY DATABASE REPOSITORY

type UserRepository interface {
	// CRUD
	Save(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
	Update(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error)
	Delete(ctx context.Context, tx *sql.Tx, user domain.User) error
	FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error)
	FindByAll(ctx context.Context, tx *sql.Tx) []domain.User
}
