package todorepository

import (
	"context"
	"database/sql"

	"github.com/imnzr/DevOps-Project-01/todo-list-api/backend/models/domain"
)

type TodoRepository interface {
	Save(ctx context.Context, tx *sql.Tx, todo domain.Todo) (domain.Todo, error)
	UpdateTitle(ctx context.Context, tx *sql.Tx, todo domain.Todo) (domain.Todo, error)
	UpdateDescription(ctx context.Context, tx *sql.Tx, todo domain.Todo) (domain.Todo, error)
	Delete(ctx context.Context, tx *sql.Tx, todo domain.Todo) error
	FindById(ctx context.Context, tx *sql.Tx, todoId int) (domain.Todo, error)
	FindByAll(ctx context.Context, tx *sql.Tx) ([]domain.Todo, error)
}
