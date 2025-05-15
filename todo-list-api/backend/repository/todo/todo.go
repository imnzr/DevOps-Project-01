package todo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/imnzr/DevOps-Project-01/todo-list-api/backend/models/domain"
)

type TodoRepositoryImpl struct{}

// UpdateDescription implements TodoRepository.
func (t *TodoRepositoryImpl) UpdateDescription(ctx context.Context, tx *sql.Tx, todo domain.Todo) (domain.Todo, error) {
	query := "UPDATE `todos` SET description = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, todo.Description)
	if err != nil {
		return todo, fmt.Errorf("failed to update description")
	}

	return todo, nil
}

// UpdateTitle implements TodoRepository.
func (t *TodoRepositoryImpl) UpdateTitle(ctx context.Context, tx *sql.Tx, todo domain.Todo) (domain.Todo, error) {
	query := "UPDATE `todos` SET title = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, todo.Title)
	if err != nil {
		return todo, fmt.Errorf("failed to update title")
	}

	return todo, nil
}

// Delete implements TodoRepository.
func (t *TodoRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, todo domain.Todo) error {
	query := "DELETE FROM `todos` WHERE id = ?"
	result, err := tx.ExecContext(ctx, query, todo.Id)
	if err != nil {
		return fmt.Errorf("failed to delete user with ID: %w", err)
	}

	RowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve rows affected")
	}
	if RowsAffected == 0 {
		return fmt.Errorf("no todo found with ID %d", todo.Id)
	}

	return nil
}

// FindByAll implements TodoRepository.
func (t *TodoRepositoryImpl) FindByAll(ctx context.Context, tx *sql.Tx) ([]domain.Todo, error) {
	query := "SELECT id, title, description WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	var todos []domain.Todo

	for rows.Next() {
		todo := domain.Todo{}
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Description)
		if err != nil {
			return nil, fmt.Errorf("error scanning rows: %w", err)
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

// FindById implements TodoRepository.
func (t *TodoRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, todoId int) (domain.Todo, error) {
	query := "SELECT id, title, description FROM todos WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, todoId)
	if err != nil {
		return domain.Todo{}, fmt.Errorf("error executing query for todo id %d: %w", todoId, err)
	}

	defer rows.Close()

	todo := domain.Todo{}

	if rows.Next() {
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Description)
		if err != nil {
			return todo, fmt.Errorf("error scanning todo rows: %w", err)
		}
		return todo, nil
	} else {
		return domain.Todo{}, fmt.Errorf("todo with id %d not found", todoId)
	}
}

// Save implements TodoRepository.
func (t *TodoRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, todo domain.Todo) (domain.Todo, error) {
	query := "INSERT INTO todos(title, description,) VALUES(?, ?)"
	result, err := tx.ExecContext(ctx, query, todo.Title, todo.Description)
	if err != nil {
		return todo, fmt.Errorf("failed to insert todo: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return todo, fmt.Errorf("failed to retrieve todo: %w", err)
	}

	todo.Id = int(id)
	return todo, nil
}

func NewTodoRepository() TodoRepository {
	return &TodoRepositoryImpl{}
}
