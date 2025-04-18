package controllers

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/imnzr/DevOps-Project-01/todo-list-api/models"
)

type UserRepositoryController struct {
	DB *sql.DB
}

func (r *UserRepositoryController) Insert(ctx context.Context, user models.User) {
	query := "INSERT INTO user(username, email, password) VALUES(?, ?, ?)"
	result, err := r.DB.ExecContext(ctx, query)
	if err != nil {
		return user, fmt.Errorf("error inserting user: %w", err)
	}
}
func (r *UserRepositoryController) GetById(ctx context.Context, id int64) (models.User, error) {}
func (r *UserRepositoryController) UpdateById(ctx context.Context, user models.User)           {}
func (r *UserRepositoryController) DeleteById(ctx context.Context, id int64) error             {}
