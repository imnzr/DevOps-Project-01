package implementation

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/imnzr/DevOps-Project-01/todo-list-api/backend/models/domain"
)

type UserRepositoryImpl struct{}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	query := "INSERT INTO user(username, email, password) VALUES(?, ?, ?)"
	result, err := tx.ExecContext(ctx, query, user.Username, user.Email, user.Password)
	if err != nil {
		return user, fmt.Errorf("failed to insert user: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return user, fmt.Errorf("failed to retrieve last insert ID: %w", err)
	}

	user.Id = int(id)
	return user, nil
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) (domain.User, error) {
	query := "UPDATE `user` SET username = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, user.Username)
	if err != nil {
		return user, fmt.Errorf("failed to update username: %w", err)
	}

	return user, nil
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.User) error {
	query := "DELETE FROM `user` WHERE id = ?"
	result, err := tx.ExecContext(ctx, query, user.Id)
	if err != nil {
		return fmt.Errorf("failed to delete user with ID %d: %w", user.Id, err)
	}

	RowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve rows affected for delete operations: %w", err)
	}
	if RowsAffected == 0 {
		return fmt.Errorf("no user found with ID %d", user.Id)
	}

	return nil
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	query := "SELECT id, email FROM user WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, userId)
	if err != nil {
		// error message here
	}

	defer rows.Close()

	user := domain.User{}

	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Email)
		if err != nil {
			// message error here
		}
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}

}

func (repository *UserRepositoryImpl) FindByAll(ctx context.Context, tx *sql.Tx) []domain.User {
	query := "SELECT id, username, email FROM user"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		// err handler here
	}

	defer rows.Close()

	var users []domain.User

	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.Id, &user.Username, &user.Email)
		if err != nil {
			// err handler here
		}
		users = append(users, user)
	}

	return users

}
